package controller

import "github.com/gin-gonic/gin"

func AddRoutes(rg *gin.RouterGroup) {
	rg.GET("/:id", getBy)
	rg.POST("", post)
	rg.POST("/editor", postEditor)
}
