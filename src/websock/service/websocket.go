package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"stouch_server/src/auth/model"
	"stouch_server/src/common/livemsg"
	"stouch_server/src/common/utils"
	"stouch_server/src/core"
	"sync"
)

var connMap = &sync.Map{}
var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}}

func closeAction(id int64) {
	if conn, ok := connMap.Load(id); ok {
		if err := conn.(*websocket.Conn).Close(); err != nil {
			core.Logger.Error(logrus.Fields{"id": id, "err": err})
		}
		connMap.Delete(id)
	}
	SendMessageToAll(livemsg.NewLiveMsg(livemsg.LiveCount, livemsg.LiveCountMsg{Count: utils.GetSyncMapLen(connMap)}))
	core.Logger.Info("websock connect is closed. userId: ", id)
}

func Send(ids []int64, message *livemsg.LiveMsg) []int64 {
	var closeIds []int64
	for _, id := range ids {
		if conn, ok := connMap.Load(id); ok {
			if messageJson, err := json.Marshal(message); err == nil {
				if err := conn.(*websocket.Conn).WriteMessage(1, messageJson); err != nil {
					core.Logger.Error("write to websocket:", err)
				}
			} else {
				core.Logger.Error("live message to json is err:", err)
			}
		} else {
			closeIds = append(closeIds, id)
		}
	}
	return closeIds
}

func SendMessageToAll(message *livemsg.LiveMsg) {
	connMap.Range(func(key any, value any) bool {
		jsonBytes, _ := json.Marshal(message)
		err := value.(*websocket.Conn).WriteMessage(1, jsonBytes)
		if err != nil {
		}
		return true
	})
}

func handleConnectionAll(c *gin.Context) {
	con, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		core.Logger.Error("upgrade:", err)
		return
	}
	user := c.MustGet("user").(model.User)
	connMap.Store(user.Id, con)
	defer closeAction(user.Id)
	SendMessageToAll(livemsg.NewLiveMsg(livemsg.LiveCount, livemsg.LiveCountMsg{Count: utils.GetSyncMapLen(connMap)}))
	for {
		mt, message, err := con.ReadMessage()
		if err != nil {
			core.Logger.Error("read websocket message: ", err)
			break
		}
		core.Logger.WithFields(logrus.Fields{"userId": user.Id}).Info(string(message))
		// ping pong 保活
		if string(message) == "ping" {
			err = con.WriteMessage(mt, []byte("pong"))
		}
		if err != nil {
			core.Logger.Error("write to websocket:", err)
			break
		}
	}
}

func AddWebsocketRoutes(rg *gin.RouterGroup) {
	rg.GET("", handleConnectionAll)
}
