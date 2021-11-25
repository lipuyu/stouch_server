package storage

import (
	"github.com/gin-gonic/gin"
	"stouch_server/src/storage/controller"
)

func AddRoutes(rg *gin.RouterGroup) {
	rg.GET("/:id", controller.GetBy)
	rg.POST("", controller.Post)
	rg.POST("/editor", controller.PostEditor)
}
