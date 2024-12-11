package datalayer

import "fmt"

func GetBookContentKey(key int64) string {
	return fmt.Sprintf("book_content_%d", key)
}

func GetBookTopicKey(key int64) string {
	return fmt.Sprintf("book_topic_%d", key)
}
