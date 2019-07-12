package auth

import (
	"github.com/kataras/iris"
	"stouch_server/auth/model"
	"stouch_server/common/error_response"
	"stouch_server/conf"
)

func Before(ctx iris.Context) {
	/*
	shareInformation := "this is a sharable information between handlers"
	requestPath := ctx.Path()
	println("Before the indexHandler or contactHandler: " + requestPath)
	ctx.Values().Set("info", shareInformation)
	*/
	token := model.Token{Ticket: ctx.GetHeader("ticket")}
	_, err := conf.Orm.Get(&token)
	if err == nil {}
	user := model.User{Id:token.UserId}
	c, err := conf.Orm.Get(&user)
	if err == nil { print(c) }
	ctx.Values().Set("user", user)
	app := ctx.Request().Header.Get("app")
	if app == "stouch" {
		ctx.Next()
	} else {
		error_response.NoError.Msg = "app error"
		ctx.JSON(error_response.NoError)
	}
}
