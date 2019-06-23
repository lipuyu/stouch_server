package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"imgo/auth/controller"
	"imgo/auth"
	"imgo/conf"
	"os"
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
	env := "test"
	if len(os.Args) >= 2 {
		env = os.Args[1]
	}
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML("<b>Resource Not found</b>")
	})
	app.UseGlobal(auth.Before)
	config := iris.YAML("./conf/source/" + env + ".yml")
	conf.LoadAll(config)
	go conf.Run()
	app.Run(iris.Addr("0.0.0.0:8080"), iris.WithConfiguration(config))
}
