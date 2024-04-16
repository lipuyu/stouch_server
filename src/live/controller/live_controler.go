package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stouch_server/src/auth/model"
	"stouch_server/src/common/er"
	"stouch_server/src/common/re"
	"stouch_server/src/live/msg"
	"stouch_server/src/live/service"
)

func PostFocusUser(c *gin.Context) {
	msg := msg.LiveStatusMsgR{}
	if err := c.ShouldBindUri(&msg); err == nil {
		liveService := service.LiveService{UserId: msg.UserId}
		user := c.MustGet("user").(model.User)
		liveService.Focus(user.Id)
	} else {
		c.JSON(http.StatusOK, er.JsonBodyError)
		return
	}
	c.JSON(http.StatusOK, re.Data(gin.H{}))
}
