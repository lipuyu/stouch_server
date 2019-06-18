package conf

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"imgo/auth/model"
)

var Orm *xorm.Engine

func loadOrm(c iris.Configuration){
	if Orm == nil {
		var err error
		Orm, err = xorm.NewEngine(
			"mysql",
			//"stouch:Jibuzhu123@tcp(rm-2zew4kr6drni3qkok.mysql.rds.aliyuncs.com:3306)/stouch",
			c.Other["DataSourceName"].(string),
			)
		if err != nil {
			fmt.Print(err)
		} else {
			iris.RegisterOnInterrupt(func() {
				err = Orm.Close()
			})
			err = Orm.Sync2(new(model.User))
		}
	}
}
