package handler

type MsgHandler interface {
	GetBackMsg(input []byte) (bool, []byte)
}
