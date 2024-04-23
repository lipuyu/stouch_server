package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stouch_server/src/auth/model"
	"stouch_server/src/common/er"
	"stouch_server/src/common/re"
	"stouch_server/src/live/msg"
	"stouch_server/src/live/service"
	"stouch_server/src/websock/livepool"
)

func PostFocusUser(c *gin.Context) {
	msgR := msg.LiveStatusMsgR{}
	if err := c.ShouldBindJSON(&msgR); err == nil {
		user := c.MustGet("user").(model.User)
		liveService := service.LiveService{UserId: user.Id}
		liveService.Focus(msgR.UserId)
	} else {
		c.JSON(http.StatusOK, re.Error(er.JsonBodyError))
		return
	}
	c.JSON(http.StatusOK, re.Data(gin.H{"status": livepool.Online(msgR.UserId)}))
}

func PostUnfocusUser(c *gin.Context) {
	msgR := msg.LiveStatusMsgR{}
	if err := c.ShouldBindJSON(&msgR); err == nil {
		user := c.MustGet("user").(model.User)
		liveService := service.LiveService{UserId: user.Id}
		liveService.Focus(msgR.UserId)
	} else {
		c.JSON(http.StatusOK, re.Error(er.JsonBodyError))
		return
	}
	c.JSON(http.StatusOK, re.Data(gin.H{}))
}
