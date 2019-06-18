package controller

import (
	"github.com/kataras/iris"
	"imgo/conf"
	"time"
)
import "imgo/auth/model"

type UserController struct{
	Ctx iris.Context
}

func (c *UserController) GetHello() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}

func (c *UserController) GetInsert() interface{} {
	user := &model.User{Username: "kataras", Salt: "hash---", Password: "hashed", CreatedAt: time.Now()}
	conf.Orm.Insert(user)
	return user
}

func (c *UserController) GetGet(ctx iris.Context) interface{}{
	user := model.User{Id:1}
	/*
	user1 := model.User{}
	ctx.ReadJSON(&user1)
	fmt.Println(user1)
	*/
	res, err := conf.Cache.Value("user")
	if err == nil {
		user = *res.Data().(*model.User)
	} else {
		if ok, _ := conf.Orm.Get(&user); ok {
			conf.Cache.Add("user", 500*time.Second, &user)
		}
	}
	return user
}
