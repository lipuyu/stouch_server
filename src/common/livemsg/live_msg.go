package livemsg

import (
	"time"
)

type LiveMsg struct {
	Code int64       `json:"code"`
	Time int64       `json:"time"`
	Data interface{} `json:"data"`
}

func NewLiveMsg(code int64, data interface{}) *LiveMsg {
	return &LiveMsg{
		Code: code,
		Time: time.Now().UnixMilli(),
		Data: data,
	}
}
