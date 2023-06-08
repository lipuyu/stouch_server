package handler

type MsgHandler interface {
	GetBackMsg(input string) string
}
