package model

import (
	"fmt"
	"stouch_server/src/core"
)

func init() {
	if err := core.Orm.Sync2(new(Topic), new(Comment), new(TopicLike)); err != nil {
		fmt.Println(err)
	}
}
