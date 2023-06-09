package handler

type PingMsgHandler struct {
}

func (p PingMsgHandler) GetBackMsg(input []byte) (bool, []byte) {
	if string(input) == "ping" {
		return true, []byte("pong")
	} else {
		return false, nil
	}
}
