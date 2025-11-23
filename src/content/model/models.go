package model

import "time"

type Topic struct {
	Id        int64     `json:"id"` // auto-increment by-default by xorm
	UserId    int64     `xorm:"bigint(20)" json:"userId"`
	Content   string    `xorm:"text" json:"content"`
	CreatedAt time.Time `xorm:"created" json:"createdAt"`
	UpdatedAt time.Time `xorm:"updated" json:"updatedAt"`
}

type Comment struct {
	Id        int64     `json:"id"` // auto-increment by-default by xorm
	UserId    int64     `xorm:"bigint(20)" json:"userId"`
	Content   string    `xorm:"text" json:"content"`
	CreatedAt time.Time `xorm:"created" json:"createdAt"`
	UpdatedAt time.Time `xorm:"updated" json:"updatedAt"`
}

type TopicLike struct {
	Id        int64     `json:"id"` // auto-increment by-default by xorm
	UserId    int64     `xorm:"bigint(20) index" json:"userId"`
	TopicId   int64     `xorm:"bigint(20) index" json:"topicId"`
	CreatedAt time.Time `xorm:"created" json:"createdAt"`
	UpdatedAt time.Time `xorm:"updated" json:"updatedAt"`
}
