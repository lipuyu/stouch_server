package core

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
)

var Orm *xorm.Engine

func loadOrm(c Configuration) {
	if Orm == nil {
		var err error
		if Orm, err = xorm.NewEngine(c.Database.DriverName, c.Database.Url); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}
	if _, err := Orm.QueryString("select 1"); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
