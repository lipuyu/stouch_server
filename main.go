package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"stouch_server/auth"
	"stouch_server/auth/controller"
	"stouch_server/conf"
	"stouch_server/websock"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	mvc.New(app).Handle(new(controller.UserController))
	return app
}

func main() {
	app := newApp()
	conf.LoadLog(app)
	websock.SetupWebsocket(app) // websocket 服务
	/*
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML("<b>Resource Not found</b>")
	})*/
	app.UseGlobal(auth.Before)
	go conf.Run()
	app.Run(iris.Addr("0.0.0.0:8080"), iris.WithConfiguration(conf.Config))
}
