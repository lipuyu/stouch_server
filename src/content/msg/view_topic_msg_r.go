package msg

type ViewTopicMsgR struct {
	TopicId int64 `json:"topicId"`
	Count   int   `json:"count"`
}
