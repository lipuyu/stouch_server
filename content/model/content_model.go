package model

import "time"

type Topic struct {
	Id int64  `json:"id"` // auto-increment by-default by xorm
	Content  string	`xorm:"text"`
	CreatedAt time.Time `xorm:"created"`
}
