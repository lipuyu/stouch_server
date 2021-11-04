package main

import (
	"fmt"
	_ "github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	appconfCt "stouch_server/appconf/controller"
	authCt "stouch_server/auth/controller"
	"stouch_server/conf"
	contentCt "stouch_server/content/controller"
	storageCt "stouch_server/storage/controller"
	bookCt "stouch_server/websock/controller"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	// 解决跨域问题
	allowAllOrigins := func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")                                         // <- HERE
		ctx.Header("Access-Control-Allow-Headers", "ticket, Content-Type, app, authorization") // <- HERE
		ctx.Header("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, PATCH, POST, PUT")
		if ctx.Method() != "OPTIONS" {
			ctx.Next()
		}
	}
	app.Use(allowAllOrigins)
	app.AllowMethods(iris.MethodOptions)

	app.Get("/", func(ctx iris.Context) { ctx.Redirect("http://airport.xiaorere.com/index.html") })
	app.StaticWeb("/web", "./static")
	mvc.New(app.Party("/user")).Handle(new(authCt.UserController))
	mvc.New(app.Party("/content")).Handle(new(contentCt.ContentController))
	mvc.New(app.Party("/storage/token")).Handle(new(storageCt.StorageTokenController))
	mvc.New(app.Party("/storage/picture")).Handle(new(storageCt.PictureController))
	mvc.New(app.Party("/book")).Handle(new(bookCt.BookController))
	return app
}

func main() {
	r := gin.Default()
	appconfCt.AddRoutes(r.Group("/appconf"))
	go conf.Run()
	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
