package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCdn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"cdn": "http://airport.xiaorere.com"})
}
