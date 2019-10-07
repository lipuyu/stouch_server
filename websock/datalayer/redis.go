package datalayer

import "fmt"

func GetBookContentKey(key int64) string {
	return fmt.Sprintf("book_content_%d", key)
}

func GetHaveBookContentKey(key int64) string {
	return fmt.Sprintf("have_book_content_%d", key)
}
