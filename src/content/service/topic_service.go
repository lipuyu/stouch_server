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
	core.Redis.SAdd(datalayer.GetBookContentKey(topicId), userId)
	if userStrIds, err := core.Redis.SMembers(datalayer.GetBookContentKey(topicId)).Result(); err == nil {
		msg := livemsg.NewLiveMsg(livemsg.VIEW_TOPIC, msg2.ViewTopicMsgR{Count: len(userStrIds)})
		livepool.Send(utils.StringsToInts(userStrIds), msg)
	}
}

func FocusTopic(userId int64, topicId int64) {
	core.Redis.SRem(datalayer.GetBookContentKey(topicId), userId)
	if userStrIds, err := core.Redis.SMembers(datalayer.GetBookContentKey(topicId)).Result(); err == nil {
		msg := livemsg.NewLiveMsg(livemsg.VIEW_TOPIC, msg2.ViewTopicMsgR{Count: len(userStrIds)})
		livepool.Send(utils.StringsToInts(userStrIds), msg)
	}
}
