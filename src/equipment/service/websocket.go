package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"stouch_server/src/core"
	"stouch_server/src/equipment/model"
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
	devCode := c.Query("devCode")

	// 检查设备是否存在
	var device model.Device
	if has, err := core.Orm.Where("unique_id = ?", devCode).Get(&device); err != nil {
		core.Logger.Error("get device by code:", err)
	} else if !has {
		core.Logger.WithFields(logrus.Fields{"devCode": devCode}).Error("device not found")
		con.WriteMessage(1, []byte("device not found"))
		return
	}

	// websocket打开关闭动作
	defer CloseAction(devCode)
	OpenAction(devCode, con)

	for {
		mt, message, err := con.ReadMessage()
		if err != nil {
			core.Logger.Error("read websocket message: ", err)
			break
		}
		core.Logger.WithFields(logrus.Fields{"devCode": devCode, "type": "websocket receive"}).Info(string(message))
		// 处理handler信息
		if string(message) == "ping" {
			err = con.WriteMessage(mt, []byte("pong"))
		}

		if err != nil {
			core.Logger.Error("write to websocket:", err)
			break
		}
	}
}

func AddEquipmentRoutes(rg *gin.RouterGroup) {
	rg.GET("", handleConnectionAll)
}
