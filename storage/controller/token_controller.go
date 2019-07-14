package controller

import (
	"github.com/kataras/iris"
)

type StorageTokenController struct{
	Ctx iris.Context
}

func (c *StorageTokenController) Get(ctx iris.Context) interface{}{
	return "faker_token"
}