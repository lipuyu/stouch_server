package offlinehandler

type OfflineHandler interface {
	OffLineAction(userId int64)
}
