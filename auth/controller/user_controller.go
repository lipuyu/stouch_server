package controller

import (
	"github.com/kataras/iris"
	"stouch_server/auth/model"
	"stouch_server/common/er"
	"stouch_server/common/utils"
	"stouch_server/conf"
	"time"
)

type UserController struct{
	Ctx iris.Context
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
	return er.NoError.SetData(map[string]model.User{"user": user})
}

func (c *UserController) PostSignin() interface{} {
	username := c.Ctx.PostValue("username")
	password := c.Ctx.PostValue("password")
	user := model.User{Username: username}
	if ok, _ := conf.Orm.Get(&user); ok {
		if user.Check(password){
			return map[string]bool{"rt": true}
		} else {
			return er.PasswordError
		}
	} else {
		return er.UserNotExistError
	}
}

func (c *UserController) PostSignup() interface{} {
	username := c.Ctx.PostValue("username")
	password := c.Ctx.PostValue("password")
	user := &model.User{Username: username, CreatedAt: time.Now()}
	user.SetPassword(password)
	var token model.Token
	if  _, err := conf.Orm.Insert(user); err == nil {
		token = model.Token{UserId: user.Id, Ticket: utils.GetUUID()}
		conf.Orm.Insert(token)
	} else {
		conf.Logger.Error(err)
	}
	return er.NoError.SetData(map[string]string{"ticket": token.Ticket})
}

func (c *UserController) GetBy(id int64) interface{}{
	user := model.User{Id: id}
	if ok, _ := conf.Orm.Get(&user); ok {
		return er.NoError.SetData(map[string]model.User{"user": user})
	} else  {
		return er.UserNotExistError
	}
}
