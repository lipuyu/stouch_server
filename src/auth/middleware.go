package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	model2 "stouch_server/src/auth/model"
	re "stouch_server/src/common/base"
	"stouch_server/src/common/er"
	"stouch_server/src/core"
	"strings"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		ticket := c.GetHeader("ticket")
		app := c.GetHeader("app")

		// websocket 信息进行特殊处理
		if strings.HasPrefix(path, "/websocket") {
			ticket = c.Query("ticket")
			app = c.Query("app")
		}

		// 读取user信息
		var user model2.User
		if ticket == "" {
			user = model2.User{}
		} else {
			token := model2.Token{Ticket: ticket}
			if g, err := core.Orm.Get(&token); err == nil && g {
				user = model2.User{Id: token.UserId}
				if c, err := core.Orm.Get(&user); err != nil || !c {
					user = model2.User{}
				}
			} else {
				user = model2.User{}
			}
		}
		c.Set("user", user)

		// URL 拦截
		if (app == "stouch" && user.Id != 0) || path == "/storage/token" || path == "/storage/picture/editor" ||
			path == "/" || c.Request.Method == "OPTIONS" || strings.HasPrefix(path, "/user/sign") ||
			strings.HasPrefix(path, "/static/") {
			c.Next()
		} else {
			c.Abort()
			c.JSON(http.StatusOK, re.NewByError(er.AppError))
			return
		}
	}
}
