package controller

import (
	"net/http"
	"stouch_server/src/common/re"
	"time"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, re.Data(gin.H{
		"cdn":        "https://stouch.oss-cn-beijing.aliyuncs.com/",
		"expireTime": time.Now().Unix() + 24*60*60,
	}))
}
