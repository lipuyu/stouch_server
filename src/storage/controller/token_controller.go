package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context) {
	c.JSON(http.StatusOK,"faker_token")
}
