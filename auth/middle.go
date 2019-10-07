package auth

import (
	"github.com/kataras/iris"
	"stouch_server/auth/model"
	"stouch_server/common/er"
	"stouch_server/conf"
	"strings"
)

func Before(ctx iris.Context) {
	/*
	shareInformation := "this is a sharable information between handlers"
	requestPath := ctx.Path()
	println("Before the indexHandler or contactHandler: " + requestPath)
	ctx.Values().Set("info", shareInformation)
	*/
	ticket := ctx.GetHeader("ticket")
	var user model.User
	if ticket == "" {
		user = model.User{}
	} else {
		token := model.Token{Ticket: ticket}
		if g, err := conf.Orm.Get(&token);  err == nil && g {
			user = model.User{Id: token.UserId}
			if c, err := conf.Orm.Get(&user); err != nil || !c{
				user = model.User{}
			}
		} else {
			user = model.User{}
		}
	}
	ctx.Values().Set("user", user)
	app := ctx.GetHeader("app")
	path := ctx.GetCurrentRoute().ResolvePath()
	if (app == "stouch" && user.Id != 0) || path == "/storage/token" || path == "/" ||
		strings.HasPrefix(path, "/web/") || strings.HasPrefix(path, "/websocket"){
		ctx.Next()
	} else {
		ctx.JSON(er.AppError)
	}
}
