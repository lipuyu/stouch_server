package handler

type PingMsgHandler struct {
}

func (p PingMsgHandler) GetBackMsg(input string) string {
	if input == "ping" {
		return "pong"
	} else {
		return ""
	}
}
