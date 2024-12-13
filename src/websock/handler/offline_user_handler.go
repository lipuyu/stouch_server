package handler

import (
	"stouch_server/src/common/livemsg"
	"stouch_server/src/live/msg"
	"stouch_server/src/live/service"
	"stouch_server/src/websock/livepool"
)

type OfflineUserHandler struct {
}

func (p OfflineUserHandler) OffLineAction(userId int64) {
	livepool.SendMessageToAll(livemsg.NewLiveMsg(livemsg.LIVE_COUNT, msg.LiveCountMsg{Count: livepool.OnlineCount()}))
	liveMsg := livemsg.NewLiveMsg(livemsg.LIVE_STATUS, msg.LiveStatusMsg{Status: false, UserId: userId})
	ids := livepool.Send(service.LiveService{UserId: userId}.GetFoucsMeIds(), liveMsg)
	service.LiveService{UserId: userId}.UnFocused(ids)
}
