package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"imgo/auth/controller"
	"imgo/conf"
	"imgo/websocket"
)


func newApp() *iris.Application {
	app := iris.New()

	mvc.New(app).Handle(new(controller.UserController))
	return app
}

func main() {
	app := newApp()
	imgo.LoadOrm()
	imgo.LoadCache()
	websocket.SetupWebsocket(app)
	app.Run(iris.Addr(":8080"))
}
