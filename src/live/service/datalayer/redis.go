package datalayer

import "fmt"

func GetFocusedKey(key int64) string {
	return fmt.Sprintf("live_focused_%d", key)
}
