package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"stouch_server/common/utils"
	"time"
)

type User struct {
	Id        int64     `json:"id"` // auto-increment by-default by xorm
	Salt      string	`xorm:"varchar(32)" json:"-"`
	Username  string	`xorm:"varchar(32) notnull unique" json:"username"`
	Password  string    `xorm:"varchar(32)" json:"-"`
	Mobile    string    `xorm:"varchar(16)" json:"mobile"`
	Avatar    int64     `xorm:"bigint(20)" json:"avatar"`
	Gender    int       `xorm:"int(11)" json:"gender"`
	CreatedAt time.Time `xorm:"created" json:"createdAt"`
	Birthday time.Time  `xorm:"datetime" json:"birthday"`
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
	fmt.Println(user.Password, getMd5(password, user.Salt))
	fmt.Println(user.Password, getMd5(password, user.Salt))
	return user.Password == getMd5(password, user.Salt)
}

type Token struct {
	Id        int64     `json:"id"` // auto-increment by-default by xorm
	Ticket    string	`xorm:"varchar(32)" json:"ticket"`
	ValidTime int64     `xorm:"bigint(20)" json:"validTime"`
	UserId    int64     `xorm:"bigint(20)" json:"userId"`
	CreatedAt time.Time `xorm:"created" json:"createdAt"`
}