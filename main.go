package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"imgo/auth/conf"
	"imgo/auth/controller"
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
	imgo.GetOrm()
	app.Run(iris.Addr(":8080"))
}
