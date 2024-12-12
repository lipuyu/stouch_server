package msghandler

import "stouch_server/src/common/livemsg"

type MsgHandler interface {
	GetBackMsg(msg *livemsg.LiveMsg) *livemsg.LiveMsg
}
