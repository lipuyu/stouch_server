package controller

import (
	"github.com/kataras/iris"
)

type StorageTokenController struct{
	Ctx iris.Context
}

func (c *StorageTokenController) Get() interface{}{
	return "faker_token"
}
