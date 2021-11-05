package controller

import (
"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	rg.GET("/user/:id", GetBy)
	rg.POST("/signin", PostSignin)
	rg.POST("/signup", PostSignup)
	rg.GET("/verification/code", GetVerificationCode)
	rg.POST("/code/check", PostCodeCheck)
}
