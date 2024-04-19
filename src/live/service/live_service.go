package service

import (
	"stouch_server/src/core"
	"stouch_server/src/live/service/datalayer"
	"strconv"
)

type LiveService struct {
	UserId int64
}

// Focus 关注某人在线状态
func (l LiveService) Focus(userId int64) {
	if l.UserId != userId {
		core.Redis.SAdd(datalayer.GetFocusedKey(userId), l.UserId)
	}
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
	ids = append(ids, l.UserId)
	return ids
}

func (l LiveService) UnfocusUser(userId int64) {
	if l.UserId != userId {
		core.Redis.SRem(datalayer.GetFocusedKey(userId), l.UserId)
	}
}

// UnFocused 不被某人关注在线状态
func (l LiveService) UnFocused(userIds []int64) {
	var interfaceSlice []interface{}
	for _, id := range userIds {
		if l.UserId != id {
			interfaceSlice = append(interfaceSlice, id)
		}
	}
	core.Redis.SRem(datalayer.GetFocusedKey(l.UserId), interfaceSlice...)
}
