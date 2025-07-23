package livepool

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"stouch_server/src/common/livemsg"
	"stouch_server/src/common/utils"
	"stouch_server/src/core"
	"sync"
)

var connMap = &sync.Map{}

func GetConnMap() *sync.Map {
	return connMap
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

func SendStringToAll(message string) {
	connMap.Range(func(key any, value any) bool {
		err := value.(*websocket.Conn).WriteMessage(1, []byte(message))
		if err != nil {
		}
		return true
	})
}

// Send 发送消息
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

// Online 用户在线状态
func Online(userId int64) bool {
	_, ok := connMap.Load(userId)
	return ok
}

func OnlineCount() int {
	return utils.GetSyncMapLen(connMap)
}
