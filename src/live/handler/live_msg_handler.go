package handler

import (
	"encoding/json"
	"stouch_server/src/common/livemsg"
	"stouch_server/src/live/msg"
	"stouch_server/src/websock/livepool"
)

type LiveStatusMsgHandler struct {
}

func (p LiveStatusMsgHandler) GetBackMsg(input []byte) (bool, []byte) {
	message := &livemsg.LiveMsg{}
	if err := json.Unmarshal(input, message); err != nil && message.Code == livemsg.LIVE_STATUS {
		liveStatusMsgR := message.Data.(msg.LiveStatusMsgR)
		liveStatusMsg := msg.LiveStatusMsg{}
		liveStatusMsg.UserId = liveStatusMsgR.UserId
		liveStatusMsg.Status = livepool.Online(liveStatusMsgR.UserId)
		resultMsg := livemsg.NewLiveMsg(livemsg.LIVE_STATUS, liveStatusMsg)

		if jsonByte, err := json.Marshal(resultMsg); err == nil {
			return true, jsonByte
		}
	}
	return false, nil
}
