package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"net/http"
	authM "stouch_server/src/auth/model"
	"stouch_server/src/core"
)

var connMap = map[int64]*websocket.Conn{}
var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}}

func Send(ids []int64, message string) []int64 {
	var closeIds []int64
	for _, id := range ids {
		if conn, ok := connMap[id]; ok {
			if err := conn.WriteMessage(1, []byte(message)); err != nil {
				core.Logger.Error("write to websocket:", err)
			}
		} else {
			closeIds = append(closeIds, id)
		}
	}
	return closeIds
}

func handleConnection(c *gin.Context) {
	con, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		core.Logger.Error("upgrade:", err)
		return
	}
	defer con.Close()
	user := c.MustGet("user").(authM.User)
	connMap[user.Id] = con
	defer delete(connMap, user.Id)
	for {
		mt, message, err := con.ReadMessage()
		if err != nil {
			core.Logger.Error("read websocket message: ", err)
			break
		}
		core.Logger.Info(string(message))
		core.Logger.Error("error test:", string(message))
		err = con.WriteMessage(mt, []byte(" recv over: "+string(message)))
		if err != nil {
			core.Logger.Error("write to websocket:", err)
			break
		}
	}
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
