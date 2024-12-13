package service

import (
	"stouch_server/src/common/livemsg"
	"stouch_server/src/common/utils"
	msg2 "stouch_server/src/content/msg"
	"stouch_server/src/content/service/datalayer"
	"stouch_server/src/core"
	"stouch_server/src/websock/livepool"
)

func UnfocusTopic(userId int64, topicId int64) {
	key := datalayer.GetBookTopicKey(topicId)
	core.Redis.SRem(key, userId)
	sendTopicFocusUserCount(key, topicId)
}

func FocusTopic(userId int64, topicId int64) {
	key := datalayer.GetBookTopicKey(topicId)
	core.Redis.SAdd(key, userId)
	userTopicKey := datalayer.GetUserTopicKey(userId)
	core.Redis.SAdd(userTopicKey, topicId)
	sendTopicFocusUserCount(key, topicId)
}

func sendTopicFocusUserCount(key string, topicId int64) {
	if userStrIds, err := core.Redis.SMembers(key).Result(); err == nil {
		msg := livemsg.NewLiveMsg(livemsg.VIEW_TOPIC, msg2.ViewTopicMsgR{TopicId: topicId, Count: len(userStrIds)})
		closeIds := livepool.Send(utils.StringsToInts(userStrIds), msg)
		if len(closeIds) != 0 {
			core.Redis.SRem(key, utils.TransIntsToInterface(closeIds)...)
		}
	}
}
