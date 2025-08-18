package model

import (
	"time"
)

type Device struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name" xorm:"varchar(32) notnull"`
	Memo       string    `json:"memo" xorm:"text"`
	Type       string    `json:"type" xorm:"varchar(6) notnull"`
	UserId     int64     `json:"user_id" xorm:"default 0 notnull"`
	DeviceCode string    `json:"device_code" xorm:"varchar(64) unique notnull"`
	CreatedAt  time.Time `json:"created_at" xorm:"created"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"updated"`
}
