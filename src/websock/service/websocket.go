package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"net/http"
	"stouch_server/src/core"
	"sync"
)

var connMap = sync.Map{}
var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}}

func Send(ids []int64, message string) []int64 {
	var closeIds []int64
	for _, id := range ids {
		if conn, ok := connMap.Load(id); ok {
			if err := conn.(*websocket.Conn).WriteMessage(1, []byte(message)); err != nil {
				core.Logger.Error("write to websocket:", err)
			}
		} else {
			closeIds = append(closeIds, id)
		}
	}
	return closeIds
}

var connSet = set.New(set.ThreadSafe)

func handleConnectionAll(c *gin.Context) {
	con, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		core.Logger.Error("upgrade:", err)
		return
	}
	connSet.Add(con)
	defer con.Close()
	defer connSet.Remove(con)
	for {
		mt, message, err := con.ReadMessage()
		if err != nil {
			core.Logger.Error("read websocket message: ", err)
			break
		}
		core.Logger.Info(string(message))
		connSet.Each(func(conTmp interface{}) bool {
			_ = conTmp.(*websocket.Conn).WriteMessage(mt, message)
			return true
		})
		if err != nil {
			core.Logger.Error("write to websocket:", err)
			break
		}
	}
}

func AddWebsocketRoutes(rg *gin.RouterGroup) {
	rg.GET("", handleConnectionAll)
}
