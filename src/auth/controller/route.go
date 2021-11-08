package controller

import (
"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	rg.GET("", get)
	rg.GET("/user/:id", getBy)
	rg.POST("/signin", postSignin)
	rg.POST("/signup", postSignup)
	rg.GET("/verification/code", getVerificationCode)
	rg.POST("/code/check", postCodeCheck)
}
