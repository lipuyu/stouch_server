package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"stouch_server/src/appconf"
	"stouch_server/src/auth"
	"stouch_server/src/content"
	"stouch_server/src/core"
	"stouch_server/src/core/middlewares"
	"stouch_server/src/storage"
	"stouch_server/src/websock"
	"stouch_server/src/websock/service"
	"time"
)

func main() {
	gin.SetMode(core.Config.Application.Mode)
	r := gin.New()
	r.Use(middlewares.Log(), gin.Recovery(), middlewares.Cors(), auth.Middleware())

	// 加路由
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "static/index.html")
	})

	r.Static("/static", "./resources/static")
	appconf.AddRoutes(r.Group("/appconf"))
	auth.AddRoutes(r.Group("/user"))
	content.AddRoutes(r.Group("/content"))
	storage.AddRoutes(r.Group("/storage/picture"))
	websock.AddRoutes(r.Group("/book"))
	service.AddWebsocketRoutes(r.Group("/websocket"))

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
