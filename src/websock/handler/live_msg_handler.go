package handler

import (
	"encoding/json"
	"stouch_server/src/common/livemsg"
	"stouch_server/src/websock/livepool"
)

type LiveMsgHandler struct {
}

func (p LiveMsgHandler) GetBackMsg(input []byte) (bool, []byte) {
	msg := &livemsg.LiveMsg{}
	if err := json.Unmarshal(input, msg); err != nil && msg.Code == livemsg.LiveStatus {
		liveStatusMsgR := msg.Data.(livemsg.LiveStatusMsgR)

		liveStatusMsg := livemsg.LiveStatusMsg{}
		liveStatusMsg.UserId = liveStatusMsgR.UserId
		liveStatusMsg.Status = livepool.Online(liveStatusMsgR.UserId)
		resultMsg := livemsg.NewLiveMsg(livemsg.LiveStatus, liveStatusMsg)

		if jsonByte, err := json.Marshal(resultMsg); err == nil {
			return true, jsonByte
		}
	}
	return false, nil
}
