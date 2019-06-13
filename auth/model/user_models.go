package model

import "time"

type User struct {
	Id int64  `json:"id"` // auto-increment by-default by xorm
	Salt      string	`xorm:"varchar(32)" json:"salt"`
	Username  string	`xorm:"varchar(32)" json:"username"`
	Password  string    `xorm:"varchar(32)" json:"password"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
}
