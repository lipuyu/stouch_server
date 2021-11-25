package model

import (
	"fmt"
	"stouch_server/src/core"
)

func init(){
	if err := core.Orm.Sync2(new(User), new(Token), new(VerificationCode)); err != nil {
		fmt.Println(err)
	}
}
