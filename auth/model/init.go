package model

import "stouch_server/conf"

func init(){
	conf.Orm.Sync2(new(User))
	conf.Orm.Sync2(new(Token))
}
