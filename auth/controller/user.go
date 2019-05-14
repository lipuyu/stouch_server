package controller

import (
	"imgo/auth/conf"
	"time"
)
import "imgo/auth/model"

type UserController struct{}


func (c *UserController) GetHello() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}

func (c *UserController) GetInsert() interface{} {
	user := &model.User{Username: "kataras", Salt: "hash---", Password: "hashed", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	imgo.Orm.Insert(user)
	return user
}

func (c *UserController) GetGet() interface{} {
	user := model.User{ID: 1}
	if ok, _ := imgo.Orm.Get(&user); ok {
		return user
	}
	return user
}
