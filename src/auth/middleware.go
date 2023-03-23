package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stouch_server/src/auth/service"
	"stouch_server/src/common/er"
	"stouch_server/src/common/re"
	"stouch_server/src/common/utils"
	"stouch_server/src/core"
	"strings"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		ticket, _ := c.Cookie("ticket")
		app, _ := c.Cookie("app")

		if strings.HasPrefix(path, "/static/") || path == "/" {
			c.Next()
			return
		}

		if app != "stouch" {
			c.Abort()
			c.JSON(http.StatusOK, re.Error(er.AppError))
			return
		}

		// URL 拦截
		if c.Request.Method == "OPTIONS" || utils.PathMatch(core.Config.PathWhiteList, path) {
			c.Next()
			return
		}

		if user, has := service.GetUserByTicket(ticket); has {
			c.Set("user", user)
		} else {
			c.Abort()
			c.JSON(http.StatusOK, re.Error(er.UnLoginError))
			return
		}
	}
}
