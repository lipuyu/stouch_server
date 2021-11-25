package model

import "time"

type Picture struct {
	Id        int64     `json:"id"` // auto-increment by-default by xorm
	Md5       string	`xorm:"varchar(32) notnull unique" json:"md5"`
	Format    string	`xorm:"varchar(8)" json:"format"`
	Width     int       `xorm:"int(11)" json:"width"`
	Height    int       `xorm:"int(11)" json:"height"`
	Size      int64     `xorm:"int(11)" json:"size"`
	CreatedAt time.Time `xorm:"created" json:"-"`
}