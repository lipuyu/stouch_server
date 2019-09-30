package controller

import (
	"github.com/kataras/iris"
	"stouch_server/common/er"
)

type AppConfController struct{
	Ctx iris.Context
}

func (c *AppConfController) GetCdn() interface{} {
	return er.Data(map[string]string{"cdn": "http://airport.xiaorere.com"})
}
