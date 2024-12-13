package handler

import "stouch_server/src/content/service"

type OfflineTopicHandler struct {
}

func (p OfflineTopicHandler) OffLineAction(userId int64) {
	service.OfflineTopicAction(userId)
}
