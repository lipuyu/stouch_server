package handler

import (
	"stouch_server/src/common/livemsg"
	"stouch_server/src/live/msg"
	"stouch_server/src/live/service"
	"stouch_server/src/websock/livepool"
)

type OnlineUserHandler struct {
}

func (p OnlineUserHandler) OnLineAction(userId int64) {
	livepool.SendMessageToAll(livemsg.NewLiveMsg(livemsg.LIVE_COUNT, msg.LiveCountMsg{Count: livepool.OnlineCount()}))
	liveMsg := livemsg.NewLiveMsg(livemsg.LIVE_STATUS, msg.LiveStatusMsg{Status: true, UserId: userId})
	livepool.Send(service.LiveService{UserId: userId}.GetFoucsMeIds(), liveMsg)
}
