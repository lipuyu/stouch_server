package model

import "time"

type User struct {
	ID        int64  // auto-increment by-default by xorm
	Version   string `xorm:"varchar(200)"`
	Salt      string
	Username  string
	Password  string    `xorm:"varchar(200)"`
	Languages string    `xorm:"varchar(200)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
