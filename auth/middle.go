package auth

import (
	"fmt"
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
	token := model.Token{Ticket: ctx.GetHeader("ticket")}
	g, err := conf.Orm.Get(&token)
	if err != nil {
		conf.Logger.Error(g, err)
	}
	user := model.User{Id:token.UserId}
	c, err := conf.Orm.Get(&user)
	if err != nil {
		conf.Logger.Error(c, err)
	}
	ctx.Values().Set("user", user)
	app := ctx.GetHeader("app")
	path := ctx.GetCurrentRoute().ResolvePath()
	fmt.Println(path, strings.HasPrefix(path, "web/"))
	if app == "stouch" || path == "/storage/token" || path == "/" || strings.HasPrefix(path, "/web/") ||
		strings.HasPrefix(path, "/websocket"){
		ctx.Next()
	} else {
		ctx.JSON(er.AppError)
	}
}
