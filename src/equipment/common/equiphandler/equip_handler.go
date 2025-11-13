package equiphandler

import "stouch_server/src/common/livemsg"

type DeviceHandler interface {
	RecordMsg(msg *livemsg.LiveMsg) *livemsg.LiveMsg
}
