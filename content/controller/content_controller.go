package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"stouch_server/common/er"
	"stouch_server/conf"
	"stouch_server/content/model"
	"stouch_server/websock"
	"stouch_server/websock/datalayer"
	"strconv"
)

type ContentController struct{
	Ctx iris.Context
}

func (c *ContentController) Get() interface{} {
	topics := make([]model.Topic, 0)
	if err := conf.Orm.Find(&topics); err == nil {
		return er.Data(map[string][]model.Topic{"topics": topics})
	} else {
		return er.SourceNotExistError
	}
}

func (c *ContentController) Post() interface{} {
	topic := model.Topic{}
	if err := c.Ctx.ReadJSON(&topic); err == nil {
		if _, err = conf.Orm.Insert(&topic); err != nil {
			conf.Logger.Error(err)
		}
	} else {
		conf.Logger.Error(err)
		return er.JsonBodyError
	}
	return er.Data(map[string]model.Topic{"topic": topic})
}

func (c *ContentController) GetBy(id int64) interface{} {
	topic := model.Topic{Id: id}
	if ok, _ := conf.Orm.Get(&topic); ok {
		return er.Data(map[string]model.Topic{"topic": topic})
	} else  {
		return er.SourceNotExistError
	}
}

func (c *ContentController) PostByComment(id int64) interface{} {
	var ids []int64
	if results, err := conf.Redis.SMembers(datalayer.GetBookContentKey(id)).Result();  err == nil{
		for _, val := range results {
			if id, err := strconv.ParseInt(val, 10, 64); err == nil {
				ids = append(ids, id)
			}
		}
	}
	websock.Send(ids, c.Ctx.PostValue("comment"))
	fmt.Println(ids, c.Ctx.PostValue("comment"))
	return er.Data(map[string]bool{"result": true})
}
