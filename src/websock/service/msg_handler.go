package service

import (
	"stouch_server/src/common/livemsg"
	"stouch_server/src/common/msghandler"
	"stouch_server/src/live/handler"
)

var msgHandlerMap = map[int64]msghandler.MsgHandler{
	livemsg.LIVE_STATUS: handler.LiveStatusMsgHandler{},
}
