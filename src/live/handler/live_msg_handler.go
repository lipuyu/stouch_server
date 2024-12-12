package handler

import (
	"stouch_server/src/common/livemsg"
	"stouch_server/src/live/msg"
	"stouch_server/src/websock/livepool"
)

type LiveStatusMsgHandler struct {
}

func (p LiveStatusMsgHandler) GetBackMsg(message *livemsg.LiveMsg) *livemsg.LiveMsg {
	liveStatusMsgR := message.Data.(msg.LiveStatusMsgR)
	liveStatusMsg := msg.LiveStatusMsg{}
	liveStatusMsg.UserId = liveStatusMsgR.UserId
	liveStatusMsg.Status = livepool.Online(liveStatusMsgR.UserId)
	resultMsg := livemsg.NewLiveMsg(livemsg.LIVE_STATUS, liveStatusMsg)
	return resultMsg
}
