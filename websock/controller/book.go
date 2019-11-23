package controller

import (
	"github.com/kataras/iris"
	"stouch_server/auth/model"
	"stouch_server/common/er"
	"stouch_server/conf"
	"stouch_server/websock"
	"stouch_server/websock/datalayer"
	"strconv"
)

type BookController struct{
	Ctx iris.Context
}


func (c *BookController) PostContentBy(id int64) interface{}{
	user := c.Ctx.Values().Get("user").(model.User)
	bookKey := datalayer.GetHaveBookContentKey(user.Id)
	if contentId, err := conf.Redis.Get(bookKey).Result(); err == nil {
		contentId1, _ := strconv.ParseInt(contentId, 10, 64)
		conf.Redis.SRem(datalayer.GetBookContentKey(contentId1), user.Id)
	}
	conf.Redis.Set(bookKey, id, 0)
	conf.Redis.SAdd(datalayer.GetBookContentKey(id), user.Id)

	var ids []int64
	if results, err := conf.Redis.SMembers(datalayer.GetBookContentKey(id)).Result();  err == nil{
		for _, val := range results {
			if id, err := strconv.ParseInt(val, 10, 64); err == nil {
				ids = append(ids, id)
			}
		}
	}
	websock.Send(ids, strconv.Itoa(len(ids)) + "个人正在看这条内容，你可以与他们沟通。")
	return er.Data(map[string]bool{"result": true})
}
