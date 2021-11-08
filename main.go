package main

import (
	"fmt"
	_ "github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"net/http"
	appconfCt "stouch_server/src/appconf/controller"
	authCt "stouch_server/src/auth/controller"
	contentCt "stouch_server/src/content/controller"
	"stouch_server/src/core"
	storageCt "stouch_server/src/storage/controller"
	bookCt "stouch_server/src/websock/controller"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	// 解决跨域问题
	allowAllOrigins := func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")                                         // <- HERE
		ctx.Header("Access-Control-Allow-Headers", "ticket, Content-Type, src, authorization") // <- HERE
		ctx.Header("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, PATCH, POST, PUT")
		if ctx.Method() != "OPTIONS" {
			ctx.Next()
		}
	}
	app.Use(allowAllOrigins)
	app.AllowMethods(iris.MethodOptions)
	app.StaticWeb("/web", "./static")
	mvc.New(app.Party("/storage/token")).Handle(new(storageCt.StorageTokenController))
	mvc.New(app.Party("/book")).Handle(new(bookCt.BookController))
	return app
}

func main() {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "http://airport.xiaorere.com/index.html")
	})
	r.Static("/static", "./resources/static")
	appconfCt.AddRoutes(r.Group("/appconf"))
	authCt.AddRoutes(r.Group("/auth"))
	contentCt.AddRoutes(r.Group("/content"))
	storageCt.AddRoutes(r.Group("/storage/picture"))
	// 定时任务
	go core.Run()
	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
