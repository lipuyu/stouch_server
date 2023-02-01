package controller

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"stouch_server/src/auth/model"
	"stouch_server/src/common/er"
	"stouch_server/src/common/re"
	"stouch_server/src/common/utils"
	"stouch_server/src/core"
	"strconv"
	"time"
)

func Get(c *gin.Context) {
	/*
		user1 := model.User{}
		ctx.ReadJSON(&user1)
		fmt.Println(user1)
	*/
	user := c.MustGet("user")
	c.JSON(http.StatusOK, re.Data(gin.H{"user": user}))
}

func PostSignin(c *gin.Context) {
	jsonData := map[string]string{"username": "", "password": ""}
	_ = c.ShouldBindJSON(&jsonData)
	username, _ := jsonData["username"]
	password, _ := jsonData["password"]
	user := model.User{Username: username}
	if ok, _ := core.Orm.Get(&user); ok {
		if user.Check(password) {
			token := model.Token{UserId: user.Id, Ticket: utils.GetUUID()}
			core.Orm.Insert(token)
			c.JSON(http.StatusOK, re.Data(gin.H{"ticket": token.Ticket}))
		} else {
			c.JSON(http.StatusOK, re.Error(er.PasswordError))
		}
	} else {
		c.JSON(http.StatusOK, re.Error(er.UserNotExistError))
	}
}

func PostSignup(c *gin.Context) {
	jsonData := map[string]string{"username": "", "password": ""}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusOK, re.Error(er.ParamsError))
	}
	username, _ := jsonData["username"]
	password, _ := jsonData["password"]
	user := &model.User{Username: username, CreatedAt: time.Now()}
	user.SetPassword(password)
	var token model.Token
	if has, _ := core.Orm.Get(&model.User{Username: username}); has {
		c.JSON(http.StatusOK, re.Error(er.UserNameRepeatError))
		return
	}
	if _, err := core.Orm.Insert(user); err == nil {
		token = model.Token{UserId: user.Id, Ticket: utils.GetUUID()}
		core.Orm.Insert(token)
	} else {
	}
	c.JSON(http.StatusOK, re.Data(gin.H{"ticket": token.Ticket}))
}

func GetBy(c *gin.Context) {
	user := model.User{}
	_ = c.ShouldBindUri(&user)
	if ok, _ := core.Orm.Get(&user); ok {
		c.JSON(http.StatusOK, re.Data(gin.H{"user": user}))
	} else {
		c.JSON(http.StatusOK, re.Error(er.UserNotExistError))
	}
}

func GetVerificationCode(c *gin.Context) {
	jsonData := struct {
		Mobile string `json:"mobile"`
	}{}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusOK, re.Error(er.ParamsError))
		return
	}
	a := rand.Int63n(900000) + 100000
	code := model.VerificationCode{
		Mobile:    jsonData.Mobile,
		Code:      strconv.FormatInt(a, 10),
		ValidTime: time.Now().Unix(),
	}
	if _, err := core.Orm.Insert(code); err != nil {
	}
	core.SendSMS(jsonData.Mobile, a)
	c.JSON(http.StatusOK, re.Data(gin.H{"result": true}))
}

func PostCodeCheck(c *gin.Context) {
	user := c.MustGet("user").(model.User)
	jsonData := struct {
		Mobile string `json:"mobile"`
		Code   string `json:"code"`
	}{}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusOK, re.Error(er.ParamsError))
		return
	}
	code := model.VerificationCode{Mobile: jsonData.Mobile}
	if _, err := core.Orm.Desc("id").Get(&code); err != nil {
	}
	if code.Code == jsonData.Code && time.Now().Unix()-code.ValidTime < 300 {
		user.Mobile = jsonData.Mobile
		core.Orm.Id(user.Id).Cols("mobile").Update(&user)
		c.JSON(http.StatusOK, re.Data(gin.H{"result": true}))
	} else {
		c.JSON(http.StatusOK, re.Data(gin.H{"result": false}))
	}
}
