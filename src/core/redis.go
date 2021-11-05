package core

import (
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func loadRedis(c Configuration) {
	if Redis == nil {
		Redis = redis.NewClient(&redis.Options{
			Addr: c.Redis.Node,
			Password: c.Redis.Password,
			DB: c.Redis.Db,
		})
	}
}
