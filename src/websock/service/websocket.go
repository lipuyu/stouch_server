package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	authM "stouch_server/src/auth/model"
)

var connMap = map[int64]*websocket.Conn{}

var upgrader = websocket.Upgrader{}

func Send(ids []int64, message string) []int64 {
	var closeIds []int64
	for _, id := range ids {
		if conn, ok := connMap[id]; ok {
			if err := conn.WriteMessage(1, []byte(message)); err != nil {
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
		log.Print("upgrade:", err)
		return
	}
	defer con.Close()
	for {
		mt, message, err := con.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = con.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
	user := c.MustGet("user").(authM.User)
	connMap[user.Id] = con
	defer delete(connMap, user.Id)
}

func AddWebsocketRoutes(rg *gin.RouterGroup) {
	rg.GET("", handleConnection)
}
