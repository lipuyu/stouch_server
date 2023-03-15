package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stouch_server/src/auth/model"
	"stouch_server/src/common/livemsg"
	"stouch_server/src/common/re"
	"stouch_server/src/common/utils"
	"stouch_server/src/core"
	"stouch_server/src/websock/datalayer"
	"stouch_server/src/websock/service"
	"strconv"
)

func PostContentBy(c *gin.Context) {
	user := c.MustGet("user").(model.User)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	bookKey := datalayer.GetHaveBookContentKey(user.Id)
	if contentId, err := core.Redis.Get(bookKey).Result(); err == nil {
		contentId1, _ := strconv.ParseInt(contentId, 10, 64)
		core.Redis.SRem(datalayer.GetBookContentKey(contentId1), user.Id)
	}
	core.Redis.Set(bookKey, id, 0)
	core.Redis.SAdd(datalayer.GetBookContentKey(id), user.Id)

	var ids []int64
	if results, err := core.Redis.SMembers(datalayer.GetBookContentKey(id)).Result(); err == nil {
		for _, val := range results {
			if id, err := strconv.ParseInt(val, 10, 64); err == nil {
				ids = append(ids, id)
			}
		}
	}
	closeIds := service.Send(ids, &livemsg.LiveMsg{Code: livemsg.LiveCount, Data: livemsg.LiveCountMsg{Count: len(ids)}})
	if len(closeIds) != 0 {
		core.Redis.SRem(datalayer.GetBookContentKey(id), utils.TransIntsToInterface(closeIds)...)
	}
	c.JSON(http.StatusOK, re.Data(gin.H{"result": true}))
}
