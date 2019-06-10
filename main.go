package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"imgo/auth/controller"
	"imgo/conf"
	"imgo/websock"
)


func newApp() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	mvc.New(app).Handle(new(controller.UserController))
	return app
}

func main() {
	go conf.Run()
	app := newApp()
	conf.LoadOrm()
	conf.LoadCache()
	conf.LoadRedis()
	conf.LoadLog(app)
	websock.SetupWebsocket(app)
	app.Run(iris.Addr(":8080"))
}
