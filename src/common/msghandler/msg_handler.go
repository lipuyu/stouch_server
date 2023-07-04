package msghandler

type MsgHandler interface {
	GetBackMsg(input []byte) (bool, []byte)
}
