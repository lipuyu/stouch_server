package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        int64     `json:"id"` // auto-increment by-default by xorm
	Salt      string	`xorm:"varchar(32)" json:"salt"`
	Username  string	`xorm:"varchar(32)" json:"username"`
	Password  string    `xorm:"varchar(32)" json:"password"`
	Mobile    string    `xorm:"varchar(16)" json:"mobile"`
	Avatar    int64     `xorm:"bigint(20)" json:"avatar"`
	Gender    int       `xorm:"int(11)" json:"gender"`
	CreatedAt time.Time `xorm:"created" json:"createdAt"`
	Birthday time.Time  `xorm:"datetime" json:"birthday"`
}

func getmd5(password string) string {
	uuid1, _ := uuid.NewUUID()
	h := md5.New()
	h.Write([]byte(uuid1.String() + password))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func (user *User) SetPassword(password string) {
	user.Password = getmd5(password)
}

func (user *User) Check(password string) bool {
	return user.Password == getmd5(password)
}


type Token struct {
	Id        int64     `json:"id"` // auto-increment by-default by xorm
	Ticket    string	`xorm:"varchar(32)" json:"ticket"`
	ValidTime int64     `xorm:"bigint(20)" json:"validTime"`
	UserId    int64     `xorm:"bigint(20)" json:"userId"`
	CreatedAt time.Time `xorm:"created" json:"createdAt"`
}