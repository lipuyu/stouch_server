package livepooleq

import (
	"errors"
	"github.com/gorilla/websocket"
	"stouch_server/src/common/utils"
	"sync"
)

var connMap = &sync.Map{}

func GetEquipmentConnMap() *sync.Map {
	return connMap
}

func Send(devCode string, text string) error {
	if conn, ok := connMap.Load(devCode); ok {
		if err := conn.(*websocket.Conn).WriteMessage(websocket.TextMessage, []byte(text)); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("conn not exist")
	}
}

// Online 用户在线状态
func Online(devCode string) bool {
	_, ok := connMap.Load(devCode)
	return ok
}

func OnlineCount() int {
	return utils.GetSyncMapLen(connMap)
}
