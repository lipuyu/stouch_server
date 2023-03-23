package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCdn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"cdn": "https://stouch.oss-cn-beijing.aliyuncs.com/"})
}
