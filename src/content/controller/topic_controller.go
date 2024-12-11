package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	authModel "stouch_server/src/auth/model"
	"stouch_server/src/common/re"
	"stouch_server/src/content/service"
	"strconv"
)

func PostUnfocusTopic(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, "")
		return
	}
	service.UnfocusTopic(c.MustGet("user").(authModel.User).Id, id)
	c.JSON(http.StatusOK, re.Data(gin.H{"result": true}))
}

func PostFocusTopic(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, "")
		return
	}
	service.FocusTopic(c.MustGet("user").(authModel.User).Id, id)
	c.JSON(http.StatusOK, re.Data(gin.H{"result": true}))
}
