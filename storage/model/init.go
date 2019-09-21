package model

import "stouch_server/conf"

func init(){
	conf.Orm.Sync2(new(Picture))
}
