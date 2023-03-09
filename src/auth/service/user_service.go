package service

import (
	"stouch_server/src/auth/model"
	"stouch_server/src/core"
)

func GetUserByTicket(ticket string) (model.User, bool) {
	token := model.Token{Ticket: ticket}
	if has, err := core.Orm.Get(&token); err == nil && has {
		var user = model.User{Id: token.UserId}
		if has, err := core.Orm.Get(&user); err == nil && has {
			return user, true
		}
	}
	return model.User{}, false
}
