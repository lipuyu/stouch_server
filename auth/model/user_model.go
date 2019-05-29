package model

import "time"

type User struct {
	Id int64  `json:"id"` // auto-increment by-default by xorm
	Salt      string	`xorm:"varchar(32)"`
	Username  string	`xorm:"varchar(32)"`
	Password  string    `xorm:"varchar(32)"`
	CreatedAt time.Time `xorm:"created"`
}
