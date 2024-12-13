package service

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"stouch_server/src/core"
	"stouch_server/src/websock/conf"
	handler2 "stouch_server/src/websock/handler"
	"stouch_server/src/websock/livepool"
)

// OpenAction 打开websocket处理函数
func OpenAction(id int64, con *websocket.Conn) {
	connMap := livepool.GetConnMap()
	connMap.Store(id, con)
	handler2.OnlineUserHandler{}.OnLineAction(id)
	core.Logger.Info("websock connect is open. userId: ", id)
}

// CloseAction 结束websocket处理函数
func CloseAction(id int64) {
	connMap := livepool.GetConnMap()
	if conn, ok := connMap.Load(id); ok {
		if err := conn.(*websocket.Conn).Close(); err != nil {
			core.Logger.Error(logrus.Fields{"id": id, "err": err})
		}
		connMap.Delete(id)
	}
	for _, offlineHandler := range conf.OfflineHandlers {
		offlineHandler.OffLineAction(id)
	}
	core.Logger.Info("websock connect is closed. userId: ", id)
}
