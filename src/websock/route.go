package websock

import (
	"github.com/gin-gonic/gin"
	"stouch_server/src/websock/controller"
)

func AddRoutes(rg *gin.RouterGroup) {
	rg.POST("/content/:id", controller.PostContentBy)
}
