package controller

import (
	"imgo/conf"
	"time"
)
import "imgo/auth/model"

type UserController struct{}


func (c *UserController) GetHello() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}

func (c *UserController) GetInsert() interface{} {
	user := &model.User{Username: "kataras", Salt: "hash---", Password: "hashed", CreatedAt: time.Now()}
	imgo.Orm.Insert(user)
	return user
}

func (c *UserController) GetGet() interface{} {
	user := model.User{Id: 1}
	res, err := imgo.Cache.Value("user")
	if err == nil {
		user = *res.Data().(*model.User)
	} else {
		if ok, _ := imgo.Orm.Get(&user); ok {
			imgo.Cache.Add("user", 5*time.Second, &user)
		}
	}
	return user
}
