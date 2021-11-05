package controller

import (
	"github.com/kataras/iris"
	model2 "stouch_server/src/auth/model"
	"stouch_server/src/common/base"
	"stouch_server/src/common/er"
	model3 "stouch_server/src/content/model"
	"stouch_server/src/core"
	"stouch_server/src/websock"
	"stouch_server/src/websock/datalayer"
	"strconv"
)

type ContentController struct {
	Ctx iris.Context
}

func (c *ContentController) Get() interface{} {
	topics := make([]model3.Topic, 0)
	if err := core.Orm.Limit(10, 0).Desc("id").Find(&topics); err == nil {
		return re.NewByData(map[string][]model3.Topic{"topics": topics})
	} else {
		return er.SourceNotExistError
	}
}

func (c *ContentController) Post() interface{} {
	topic := model3.Topic{}
	if err := c.Ctx.ReadJSON(&topic); err == nil {
		topic.UserId = c.Ctx.Values().Get("user").(model2.User).Id
		if _, err = core.Orm.Insert(&topic); err != nil {
		}
	} else {
		return er.JsonBodyError
	}
	return re.NewByData(map[string]model3.Topic{"topic": topic})
}

func (c *ContentController) GetBy(id int64) interface{} {
	topic := model3.Topic{Id: id}
	if ok, _ := core.Orm.Get(&topic); ok {
		return re.NewByData(map[string]model3.Topic{"topic": topic})
	} else {
		return er.SourceNotExistError
	}
}

func (c *ContentController) PostByComment(id int64) interface{} {
	jsonData := struct {Comment string `json:"comment"`}{}
	if err := c.Ctx.ReadJSON(&jsonData); err != nil {
		return er.ParamsError
	}
	var ids []int64
	if results, err := core.Redis.SMembers(datalayer.GetBookContentKey(id)).Result(); err == nil {
		for _, val := range results {
			if id, err := strconv.ParseInt(val, 10, 64); err == nil {
				ids = append(ids, id)
			}
		}
	}
	websock.Send(ids, jsonData.Comment)
	return re.NewByData(map[string]bool{"result": true})
}
