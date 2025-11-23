package service

import (
	"stouch_server/src/core"
	"stouch_server/src/equipment/livepooleq"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// OpenAction 打开websocket处理函数
func OpenAction(devCode string, con *websocket.Conn) {
	connMap := livepooleq.GetEquipmentConnMap()
	connMap.Store(devCode, con)
	core.Logger.Info("websock connect is open. userId: ", devCode)
}

// CloseAction 结束websocket处理函数
func CloseAction(devCode string) {
	connMap := livepooleq.GetEquipmentConnMap()
	if conn, ok := connMap.Load(devCode); ok {
		if err := conn.(*websocket.Conn).Close(); err != nil {
			core.Logger.Error(logrus.Fields{"devCode": devCode, "err": err})
		}
		connMap.Delete(devCode)
	}
	core.Logger.Info("websock connect is closed. userId: ", devCode)
}
