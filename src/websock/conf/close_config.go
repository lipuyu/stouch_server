package conf

import (
	"stouch_server/src/common/offlinehandler"
	"stouch_server/src/content/handler"
	handler2 "stouch_server/src/websock/handler"
)

var OfflineHandlers = []offlinehandler.OfflineHandler{
	handler2.OfflineUserHandler{},
	handler.OfflineTopicHandler{},
}
