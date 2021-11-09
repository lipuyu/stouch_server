package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stouch_server/src/auth/model"
	"stouch_server/src/common/base"
	"stouch_server/src/common/utils"
	"stouch_server/src/core"
	"stouch_server/src/websock"
	datalayer2 "stouch_server/src/websock/datalayer"
	"strconv"
)

func postContentBy(c *gin.Context) {
	user := c.MustGet("user").(model.User)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	bookKey := datalayer2.GetHaveBookContentKey(user.Id)
	if contentId, err := core.Redis.Get(bookKey).Result(); err == nil {
		contentId1, _ := strconv.ParseInt(contentId, 10, 64)
		core.Redis.SRem(datalayer2.GetBookContentKey(contentId1), user.Id)
	}
	core.Redis.Set(bookKey, id, 0)
	core.Redis.SAdd(datalayer2.GetBookContentKey(id), user.Id)

	var ids []int64
	if results, err := core.Redis.SMembers(datalayer2.GetBookContentKey(id)).Result(); err == nil {
		for _, val := range results {
			if id, err := strconv.ParseInt(val, 10, 64); err == nil {
				ids = append(ids, id)
			}
		}
	}
	closeIds := websock.Send(ids, strconv.Itoa(len(ids)) + "个人正在看这条内容，你可以与他们沟通。")
	if len(closeIds) != 0 {
		core.Redis.SRem(datalayer2.GetBookContentKey(id), utils.TransIntsToInterface(closeIds)...)
	}
	c.JSON(http.StatusOK, re.NewByData(gin.H{"result": true}))
}
