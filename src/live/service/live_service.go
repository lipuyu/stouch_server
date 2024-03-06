package service

import (
	"stouch_server/src/core"
	"stouch_server/src/live/service/datalayer"
	"strconv"
)

type LiveService struct {
	UserId int64
}

func (l LiveService) Focus(userId int64) {
	core.Redis.SAdd(datalayer.GetFocusedKey(l.UserId), userId)
}

func (l LiveService) GetFoucsMeIds() []int64 {
	var ids []int64
	if results, err := core.Redis.SMembers(datalayer.GetFocusedKey(l.UserId)).Result(); err == nil {
		for _, val := range results {
			if id, err := strconv.ParseInt(val, 10, 64); err == nil {
				ids = append(ids, id)
			}
		}
	}
	return ids
}
