package imgo

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"imgo/auth/model"
)

var Orm *xorm.Engine

func GetOrm(){
	if Orm == nil {
		var err error
		Orm, err = xorm.NewEngine("mysql", "root:@/imgo")
		if err != nil {
		}
		iris.RegisterOnInterrupt(func() {
			Orm.Close()
		})
		err = Orm.Sync2(new(model.User))
	}
}
