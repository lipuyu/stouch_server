package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	model2 "stouch_server/src/auth/model"
	"stouch_server/src/common/base"
	"stouch_server/src/common/er"
	model3 "stouch_server/src/content/model"
	"stouch_server/src/core"
	"stouch_server/src/websock"
	"stouch_server/src/websock/datalayer"
	"strconv"
)

func get(c *gin.Context) {
	topics := make([]model3.Topic, 0)
	if err := core.Orm.Limit(10, 0).Desc("id").Find(&topics); err == nil {
		c.JSON(http.StatusOK, re.NewByData(gin.H{"topics": topics}))
	} else {
		c.JSON(http.StatusOK, er.SourceNotExistError)
	}
}

func post(c *gin.Context) {
	topic := model3.Topic{}
	if err := c.ShouldBindJSON(&topic); err == nil {
		topic.UserId = c.MustGet("user").(model2.User).Id
		if _, err = core.Orm.Insert(&topic); err != nil {
		}
	} else {
		c.JSON(http.StatusOK, er.JsonBodyError)
		return
	}
	c.JSON(http.StatusOK, re.NewByData(gin.H{"topic": topic}))
}

func getBy(c *gin.Context) {
	topic := model3.Topic{}
	_ = c.ShouldBindUri(&topic)
	if ok, _ := core.Orm.Get(&topic); ok {
		c.JSON(http.StatusOK, re.NewByData(gin.H{"topic": topic}))
	} else {
		c.JSON(http.StatusOK, er.SourceNotExistError)
	}
}

func postByComment(c *gin.Context) {
	jsonData := struct {Comment string `json:"comment"`}{}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusOK, er.ParamsError)
	}
	var ids []int64
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, "")
		return
	}
	if results, err := core.Redis.SMembers(datalayer.GetBookContentKey(id)).Result(); err == nil {
		for _, val := range results {
			if id, err := strconv.ParseInt(val, 10, 64); err == nil {
				ids = append(ids, id)
			}
		}
	}
	websock.Send(ids, jsonData.Comment)
	c.JSON(http.StatusOK, re.NewByData(gin.H{"result": true}))
}
