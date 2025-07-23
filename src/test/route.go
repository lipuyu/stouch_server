package test

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	rg.POST("/high2", Hign2)
	rg.POST("/low2", Low2)
}
