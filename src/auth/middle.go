package auth

import (
	"fmt"
	"github.com/kataras/iris"
	model2 "stouch_server/src/auth/model"
	"stouch_server/src/common/er"
	"stouch_server/src/core"
	"strings"
)

func Before(ctx iris.Context) {
	/*
		shareInformation := "this is a sharable information between handlers"
		requestPath := ctx.Path()
		println("Before the indexHandler or contactHandler: " + requestPath)
		ctx.Values().Set("info", shareInformation)
	*/
	path := ctx.GetCurrentRoute().ResolvePath()
	ticket := ctx.GetHeader("ticket")
	app := ctx.GetHeader("src")

	// websocket 信息进行特殊处理
	if strings.HasPrefix(path, "/websocket") {
		ticket, _ = ctx.URLParams()["ticket"]
		app, _ = ctx.URLParams()["src"]
	}
	fmt.Println(ticket, app, "end")

	// 读取user信息
	var user model2.User
	if ticket == "" {
		user = model2.User{}
	} else {
		token := model2.Token{Ticket: ticket}
		if g, err := core.Orm.Get(&token); err == nil && g {
			user = model2.User{Id: token.UserId}
			if c, err := core.Orm.Get(&user); err != nil || !c {
				user = model2.User{}
			}
		} else {
			user = model2.User{}
		}
	}
	ctx.Values().Set("user", user)

	// URL 拦截
	if (app == "stouch" && user.Id != 0) || path == "/storage/token" || path == "/storage/picture/editor" ||
		path == "/" || ctx.Method() == "OPTIONS" || strings.HasPrefix(path, "/user/sign") ||
		strings.HasPrefix(path, "/web/") {
		ctx.Next()
	} else {
		ctx.JSON(er.AppError)
	}
}
