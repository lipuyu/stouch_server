package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stouch_server/src/auth/model"
	"stouch_server/src/common/re"
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
		var user model.User
		if ticket == "" {
			user = model.User{}
		} else {
			token := model.Token{Ticket: ticket}
			if g, err := core.Orm.Get(&token); err == nil && g {
				user = model.User{Id: token.UserId}
				if c, err := core.Orm.Get(&user); err != nil || !c {
					user = model.User{}
				}
			} else {
				user = model.User{}
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
