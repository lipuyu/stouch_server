package conf

import (
	"github.com/go-redis/redis"
	"github.com/kataras/iris"
)

var Redis *redis.Client

func If(ok bool, a interface{}, b interface{}) interface{} {
	if ok {
		return a
	} else {
		return b
	}
}

func loadRedis(c iris.Configuration) {
	password := c.Other["RedisPassword"]
	if Redis == nil {
		Redis = redis.NewClient(&redis.Options{
			Addr: c.Other["RedisAddr"].(string),
			Password: If(password != nil, password, "").(string),
			DB: c.Other["RedisDB"].(int),
		})
	}
}
