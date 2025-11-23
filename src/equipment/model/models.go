package model

import (
	"time"
)

type Device struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name" xorm:"varchar(32) notnull"`
	Memo       string    `json:"memo" xorm:"text"`
	Type       string    `json:"type" xorm:"varchar(6) notnull"`
	UserId     int64     `json:"userId" xorm:"default 0 notnull"`
	DeviceCode string    `json:"deviceCode" xorm:"varchar(64) unique notnull"`
	CreatedAt  time.Time `json:"created" xorm:"created"`
	UpdatedAt  time.Time `json:"updated" xorm:"updated"`
}

type FlowerDeviceInfo struct {
	Id          int64     `json:"id"`
	DeviceCode  string    `json:"deviceCode" xorm:"varchar(64) unique notnull"`
	MinHumidity float64   `json:"minHumidity" xorm:"float"`
	MaxHumidity float64   `json:"maxHumidity" xorm:"float"`
	CreatedAt   time.Time `json:"created" xorm:"created"`
	UpdatedAt   time.Time `json:"updated" xorm:"updated"`
}

type DeviceHumidity struct {
	Id         int64     `json:"id"`
	DeviceCode string    `json:"device_code" xorm:"varchar(64) notnull"`
	Humidity   float64   `json:"humidity" xorm:"notnull"`
	CreatedAt  time.Time `json:"created_at" xorm:"created"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"updated"`
}

type DeviceTemperature struct {
	Id          int64     `json:"id"`
	DeviceCode  string    `json:"device_code" xorm:"varchar(64) notnull"`
	Temperature float64   `json:"temperature" xorm:"notnull"`
	CreatedAt   time.Time `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"updated"`
}
