package livemsg

type LiveStatusMsg struct {
	Status bool  `json:"status"`
	UserId int64 `json:"userId"`
}
