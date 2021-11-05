package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris"
	"math/rand"
	"net/http"
	model2 "stouch_server/src/auth/model"
	"stouch_server/src/common/base"
	"stouch_server/src/common/er"
	"stouch_server/src/common/utils"
	"stouch_server/src/core"
	"strconv"
	"time"
)

type UserController struct{
	Ctx iris.Context
}

func (c *UserController) Get(ctx iris.Context) interface{}{
	/*
	user1 := model.User{}
	ctx.ReadJSON(&user1)
	fmt.Println(user1)
	*/
	user := ctx.Values().Get("user").(model2.User)
	return re.NewByData(map[string]model2.User{"user": user})
}

func (c *UserController) PostSignin() interface{} {
	jsonData := map[string]string{"username": "", "password": ""}
	c.Ctx.ReadJSON(&jsonData)
	username, _ := jsonData["username"]
	password, _ := jsonData["password"]
	user := model2.User{Username: username}
	if ok, _ := core.Orm.Get(&user); ok {
		if user.Check(password){
			token := model2.Token{UserId: user.Id, Ticket: utils.GetUUID()}
			core.Orm.Insert(token)
			return re.NewByData(map[string]string{"ticket": token.Ticket})
		} else {
			return re.NewByError(er.PasswordError)
		}
	} else {
		return re.NewByError(er.UserNotExistError)
	}
}

func (c *UserController) PostSignup() interface{} {
	jsonData := map[string]string{"username": "", "password": ""}
	if err := c.Ctx.ReadJSON(&jsonData); err != nil {
		return re.NewByError(er.ParamsError)
	}
	username, _ := jsonData["username"]
	password, _ := jsonData["password"]
	user := &model2.User{Username: username, CreatedAt: time.Now()}
	user.SetPassword(password)
	var token model2.Token
	if has, _ := core.Orm.Get(&model2.User{Username: username}); has {
		return re.NewByError(er.UserNameRepeatError)
	}
	if  _, err := core.Orm.Insert(user); err == nil {
		token = model2.Token{UserId: user.Id, Ticket: utils.GetUUID()}
		core.Orm.Insert(token)
	} else {
	}
	return re.NewByData(map[string]string{"ticket": token.Ticket})
}

func GetBy(c *gin.Context){
	user := model2.User{}
	_ = c.ShouldBindUri(&user)
	if ok, _ := core.Orm.Get(&user); ok {
		c.JSON(http.StatusOK, re.NewByData(gin.H{"user": user}))
	} else  {
		c.JSON(http.StatusOK, re.NewByError(er.UserNotExistError))
	}
}

func (c *UserController) GetVerificationCode() interface{}{
	jsonData := struct {Mobile string `json:"mobile"`}{}
	if err := c.Ctx.ReadJSON(&jsonData); err != nil {
		return re.NewByError(er.ParamsError)
	}
	a := rand.Int63n(900000) + 100000
	code := model2.VerificationCode{
		Mobile: jsonData.Mobile,
		Code: strconv.FormatInt(a,10),
		ValidTime: time.Now().Unix(),
	}
	if _, err := core.Orm.Insert(code); err != nil {
	}
	// core.SendSMS(jsonData.Mobile, a)
	return re.NewByData(iris.Map{"result": true})
}

func (c *UserController) PostCodeCheck() interface{}{
	user := c.Ctx.Values().Get("user").(model2.User)
	jsonData := struct{Mobile string `json:"mobile"`; Code string `json:"code"`}{}
	if err := c.Ctx.ReadJSON(&jsonData); err != nil {
		return re.NewByError(er.ParamsError)
	}
	code := model2.VerificationCode{Mobile: jsonData.Mobile}
	if _, err := core.Orm.Desc("id").Get(&code); err != nil {
	}
	if code.Code == jsonData.Code && time.Now().Unix() - code.ValidTime < 300 {
		user.Mobile = jsonData.Mobile
		core.Orm.Id(user.Id).Cols("mobile").Update(&user)
		return re.NewByData(iris.Map{"result": true})
	} else {
		return re.NewByData(iris.Map{"result": false})
	}
}
