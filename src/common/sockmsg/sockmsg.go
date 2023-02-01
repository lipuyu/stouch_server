package sockmsg

import (
	"time"
)

var (
	LiveCount = 1
)

type WebsockResult struct {
	Data interface{} `json:"data"`
	Type int         `json:"type"`
	Id   string      `json:"id"`
	Time time.Time   `json:"time"`
}
