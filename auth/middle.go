package auth

import "github.com/kataras/iris"

func Before(ctx iris.Context) {
	/*
	shareInformation := "this is a sharable information between handlers"
	requestPath := ctx.Path()
	println("Before the indexHandler or contactHandler: " + requestPath)
	ctx.Values().Set("info", shareInformation)
	*/
	app := ctx.Request().Header.Get("app")
	if app == "stouch" {
		ctx.Next()
	} else {
		ctx.JSON(map[string]string{"error": "app error"})
	}
}
