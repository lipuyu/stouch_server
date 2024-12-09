package service

import (
	"stouch_server/src/common/msghandler"
	handler2 "stouch_server/src/live/handler"
	"stouch_server/src/websock/handler"
)

var msgHandlers = []msghandler.MsgHandler{
	handler.PingMsgHandler{},
	handler2.LiveStatusMsgHandler{},
}
