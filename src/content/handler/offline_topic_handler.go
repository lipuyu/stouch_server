package handler

import (
	"stouch_server/src/common/utils"
	"stouch_server/src/content/service"
	"stouch_server/src/content/service/datalayer"
	"stouch_server/src/core"
)

type OfflineTopicHandler struct {
}

func (p OfflineTopicHandler) OffLineAction(userId int64) {
	userTopicKey := datalayer.GetUserTopicKey(userId)
	if userStrIds, err := core.Redis.SMembers(userTopicKey).Result(); err == nil {
		for _, topicId := range utils.StringsToInts(userStrIds) {
			service.UnfocusTopic(userId, topicId)
		}
	}
	core.Redis.Del(userTopicKey)
}
