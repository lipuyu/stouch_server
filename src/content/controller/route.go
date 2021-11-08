package controller

import "github.com/gin-gonic/gin"

func AddRoutes(rg *gin.RouterGroup) {
	rg.GET("/:id", getBy)
	rg.GET("", get)
	rg.POST("", post)
	rg.POST("/comment/:id", postByComment)
}
