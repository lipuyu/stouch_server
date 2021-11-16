package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	appconfCt "stouch_server/src/appconf/controller"
	"stouch_server/src/auth"
	authCt "stouch_server/src/auth/controller"
	contentCt "stouch_server/src/content/controller"
	"stouch_server/src/core"
	"stouch_server/src/core/middlewares"
	storageCt "stouch_server/src/storage/controller"
	"stouch_server/src/websock"
	bookCt "stouch_server/src/websock/controller"
	"time"
)

func main() {
	gin.SetMode(core.Config.Application.Mode)
	r := gin.New()
	r.Use(middlewares.Log(), gin.Recovery(), middlewares.Cors(), auth.Middleware())
	// 加路由
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "http://airport.xiaorere.com/index.html")
	})

	r.Static("/static", "./resources/static")
	appconfCt.AddRoutes(r.Group("/appconf"))
	authCt.AddRoutes(r.Group("/auth"))
	contentCt.AddRoutes(r.Group("/content"))
	storageCt.AddRoutes(r.Group("/storage/picture"))
	bookCt.AddRoutes(r.Group("/book"))
	websock.AddRoutes(r.Group("/websocket"))

	// 定时任务
	go core.Run()
	s := &http.Server{
		Addr:           core.Config.Application.Addr,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 主程序
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
