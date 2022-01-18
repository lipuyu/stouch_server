package sockmsg

import "time"

type WebsockResult struct {
	Data interface{} `json:"data"`
	Type int `json:"type"`
	Id int `json:"id"`
	Time time.Time `json:"time"`
}
