package model

import (
	"crypto/md5"
	"encoding/hex"
	"stouch_server/src/common/utils"
	"stouch_server/src/core"
	"time"
)

type User struct {
	Id        int64     `json:"id"` // auto-increment by-default by xorm
	Salt      string    `xorm:"varchar(32)" json:"-"`
	Username  string    `xorm:"varchar(32) notnull unique" json:"username"`
	Password  string    `xorm:"varchar(32)" json:"-"`
	Mobile    string    `xorm:"varchar(16)" json:"-"`
	Avatar    string    `xorm:"varchar(36)" json:"avatar"`
	Gender    int       `xorm:"int(11)" json:"gender"`
	CreatedAt time.Time `xorm:"created" json:"createdAt"`
	Birthday  time.Time `xorm:"datetime" json:"birthday"`
}

func getMd5(password string, salt string) string {
	h := md5.New()
	h.Write([]byte(salt + password))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func (user *User) SetPassword(password string) {
	user.Salt = utils.GetUUID()
	user.Password = getMd5(password, user.Salt)
}

func (user *User) Check(password string) bool {
	if core.Config.Application.Mode == "debug" {
		return password == "1234"
	}
	return user.Password == getMd5(password, user.Salt)
}

func (user *User) GetCode(mobile string) bool {
	return true
}

func (user *User) CheckCode(mobile string) bool {
	return true
}

type Token struct {
	Id        int64     `json:"id"` // auto-increment by-default by xorm
	Ticket    string    `xorm:"varchar(32)" json:"ticket"`
	ValidTime int64     `xorm:"bigint(20)" json:"validTime"`
	UserId    int64     `xorm:"bigint(20)" json:"userId"`
	CreatedAt time.Time `xorm:"created" json:"createdAt"`
}

type VerificationCode struct {
	Id        int64  `json:"id"` // auto-increment by-default by xorm
	Mobile    string `xorm:"varchar(16)" json:"mobile"`
	Code      string `xorm:"varchar(8)" json:"code"`
	ValidTime int64  `xorm:"bigint(20)" json:"validTime"`
}
