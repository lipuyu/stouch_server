package content

import (
	"github.com/gin-gonic/gin"
	"stouch_server/src/content/controller"
)

func AddRoutes(rg *gin.RouterGroup) {
	rg.GET("/:id", controller.GetBy)
	rg.GET("", controller.Get)
	rg.POST("", controller.Post)
	rg.POST("/comment/:id", controller.PostByComment)
}
