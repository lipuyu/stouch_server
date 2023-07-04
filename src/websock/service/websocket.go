package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"stouch_server/src/auth/model"
	"stouch_server/src/common/livemsg"
	"stouch_server/src/common/msghandler"
	"stouch_server/src/common/utils"
	"stouch_server/src/core"
	handler2 "stouch_server/src/live/handler"
	"stouch_server/src/live/msg"
	"stouch_server/src/websock/handler"
	"stouch_server/src/websock/livepool"
)

var msgHandlers = []msghandler.MsgHandler{
	handler.PingMsgHandler{},
	handler2.LiveMsgHandler{},
}

var connMap = livepool.GetConnMap()
var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}}

func handleConnectionAll(c *gin.Context) {
	con, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		core.Logger.Error("upgrade:", err)
		return
	}
	user := c.MustGet("user").(model.User)
	connMap.Store(user.Id, con)
	defer livepool.CloseAction(user.Id)
	livepool.SendMessageToAll(livemsg.NewLiveMsg(livemsg.LiveCount, msg.LiveCountMsg{Count: utils.GetSyncMapLen(connMap)}))
	for {
		mt, message, err := con.ReadMessage()
		if err != nil {
			core.Logger.Error("read websocket message: ", err)
			break
		}
		core.Logger.WithFields(logrus.Fields{"userId": user.Id}).Info(string(message))
		for _, val := range msgHandlers {
			if ok, backMsg := val.GetBackMsg(message); ok {
				err = con.WriteMessage(mt, backMsg)
				break
			}
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
