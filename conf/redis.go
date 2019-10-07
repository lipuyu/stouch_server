package conf

import (
	"github.com/go-redis/redis"
	"github.com/kataras/iris"
	"stouch_server/common/utils"
)

var Redis *redis.Client


func loadRedis(c iris.Configuration) {
	password := c.Other["RedisPassword"]
	if Redis == nil {
		Redis = redis.NewClient(&redis.Options{
			Addr: c.Other["RedisAddr"].(string),
			Password: utils.If(password != nil, password, "").(string),
			DB: c.Other["RedisDB"].(int),
		})
	}
}
