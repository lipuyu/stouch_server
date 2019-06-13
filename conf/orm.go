package conf

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"imgo/auth/model"
)

var Orm *xorm.Engine

func LoadOrm(){
	if Orm == nil {
		var err error
		Orm, err = xorm.NewEngine(
			"mysql",
			"stouch:Jibuzhu123@rm-2zew4kr6drni3qkok.mysql.rds.aliyuncs.com/stouch",
			)
		if err != nil {
		}
		iris.RegisterOnInterrupt(func() {
			Orm.Close()
		})
		err = Orm.Sync2(new(model.User))
	}
}
