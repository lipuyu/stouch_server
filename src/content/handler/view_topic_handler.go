package handler

import (
	"encoding/json"
	"stouch_server/src/common/livemsg"
)

type ViewTopicHandler struct {
}

func (p ViewTopicHandler) GetBackMsg(input []byte) (bool, []byte) {
	message := &livemsg.LiveMsg{}
	if err := json.Unmarshal(input, message); err != nil && message.Code == livemsg.VIEW_TOPIC {
	}
	return false, nil
}
