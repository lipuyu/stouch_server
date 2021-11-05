package controller

import (
"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	rg.GET("/user/:id", GetBy)
}
