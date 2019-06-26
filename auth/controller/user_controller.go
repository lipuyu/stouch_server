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

func (c *UserController) GetGet(ctx iris.Context) interface{}{
	user := model.User{Id:2}
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

func (c *UserController) PostGet(ctx iris.Context) interface{} {
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	user := model.User{Username: username}
	conf.Orm.Get(&user)
	return map[string]bool{"rt": user.Check(password)}
}

func (c *UserController) PostSignup(ctx iris.Context) interface{} {
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	user := &model.User{Username: username, CreatedAt: time.Now()}
	user.SetPassword(password)
	conf.Orm.Insert(user)
	return user
}
