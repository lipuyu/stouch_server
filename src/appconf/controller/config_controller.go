package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stouch_server/src/common/re"
)

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, re.Data(gin.H{"cdn": "https://stouch.oss-cn-beijing.aliyuncs.com/"}))
}
