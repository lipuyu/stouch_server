package livepool

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"stouch_server/src/common/livemsg"
	"stouch_server/src/common/utils"
	"stouch_server/src/core"
	"stouch_server/src/live/msg"
	"stouch_server/src/live/service"
	"sync"
)

var connMap = &sync.Map{}

// Online 用户在线状态
func Online(userId int64) bool {
	_, ok := connMap.Load(userId)
	return ok
}

func GetConnMap() *sync.Map {
	return connMap
}

// SendMessageToAll 发送消息给所有人
func SendMessageToAll(message *livemsg.LiveMsg) {
	connMap.Range(func(key any, value any) bool {
		jsonBytes, _ := json.Marshal(message)
		err := value.(*websocket.Conn).WriteMessage(1, jsonBytes)
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

// OpenAction 打开websocket处理函数
func OpenAction(id int64) {
	SendMessageToAll(livemsg.NewLiveMsg(livemsg.LIVE_COUNT, msg.LiveCountMsg{Count: utils.GetSyncMapLen(connMap)}))
	liveMsg := livemsg.NewLiveMsg(livemsg.LIVE_STATUS, msg.LiveStatusMsg{Status: true, UserId: id})
	Send(service.LiveService{UserId: id}.GetFoucsMeIds(), liveMsg)
	core.Logger.Info("websock connect is open. userId: ", id)
}

// CloseAction 结束websocket处理函数
func CloseAction(id int64) {
	if conn, ok := connMap.Load(id); ok {
		if err := conn.(*websocket.Conn).Close(); err != nil {
			core.Logger.Error(logrus.Fields{"id": id, "err": err})
		}
		connMap.Delete(id)
	}
	ConnMap := connMap
	count := utils.GetSyncMapLen(ConnMap)
	fmt.Println(count)
	SendMessageToAll(livemsg.NewLiveMsg(livemsg.LIVE_COUNT, msg.LiveCountMsg{Count: utils.GetSyncMapLen(ConnMap)}))
	liveMsg := livemsg.NewLiveMsg(livemsg.LIVE_STATUS, msg.LiveStatusMsg{Status: false, UserId: id})
	Send(service.LiveService{UserId: id}.GetFoucsMeIds(), liveMsg)
	core.Logger.Info("websock connect is closed. userId: ", id)
}
