package auth

import (
	"github.com/gin-gonic/gin"
	"stouch_server/src/auth/controller"
)

func AddRoutes(rg *gin.RouterGroup) {
	rg.GET("", controller.Get)
	rg.GET("/:id", controller.GetBy)
	rg.POST("/login", controller.PostLoginIn)
	rg.POST("/signup", controller.PostSignup)
	rg.GET("/verification/code", controller.GetVerificationCode)
	rg.POST("/code/check", controller.PostCodeCheck)
}
