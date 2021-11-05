package controller

import (
	"github.com/kataras/iris"
	"stouch_server/src/auth/model"
	"stouch_server/src/common/base"
	"stouch_server/src/common/utils"
	"stouch_server/src/core"
	"stouch_server/src/websock"
	datalayer2 "stouch_server/src/websock/datalayer"
	"strconv"
)

type BookController struct{
	Ctx iris.Context
}


func (c *BookController) PostContentBy(id int64) interface{}{
	user := c.Ctx.Values().Get("user").(model.User)
	bookKey := datalayer2.GetHaveBookContentKey(user.Id)
	if contentId, err := core.Redis.Get(bookKey).Result(); err == nil {
		contentId1, _ := strconv.ParseInt(contentId, 10, 64)
		core.Redis.SRem(datalayer2.GetBookContentKey(contentId1), user.Id)
	}
	core.Redis.Set(bookKey, id, 0)
	core.Redis.SAdd(datalayer2.GetBookContentKey(id), user.Id)

	var ids []int64
	if results, err := core.Redis.SMembers(datalayer2.GetBookContentKey(id)).Result();  err == nil{
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
	return re.NewByData(iris.Map{"result": true})
}
