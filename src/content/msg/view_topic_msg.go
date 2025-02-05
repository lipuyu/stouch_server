package msg

type ViewTopicMsg struct {
	TopicId int64 `json:"topicId"`
	Count   int   `json:"count"`
}
