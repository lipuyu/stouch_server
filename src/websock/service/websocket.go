package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"stouch_server/src/auth/model"
	"stouch_server/src/common/livemsg"
	"stouch_server/src/core"
	"stouch_server/src/websock/conf"
)

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

	// websocket打开关闭动作
	defer CloseAction(user.Id)
	OpenAction(user.Id, con)

	for {
		mt, message, err := con.ReadMessage()
		if err != nil {
			core.Logger.Error("read websocket message: ", err)
			break
		}
		core.Logger.WithFields(logrus.Fields{"userId": user.Id, "type": "websocket receive"}).Info(string(message))
		// 处理handler信息
		if string(message) == "ping" {
			err = con.WriteMessage(mt, []byte("pong"))
		} else {
			msgObject := &livemsg.LiveMsg{}
			if err := json.Unmarshal(message, msgObject); err != nil {
				resultMsg := conf.MsgHandlerMap[msgObject.Code].GetBackMsg(msgObject)
				if jsonByte, err := json.Marshal(resultMsg); err == nil {
					err = con.WriteMessage(mt, jsonByte)
				}
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
