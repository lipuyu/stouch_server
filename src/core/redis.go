package core

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
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
	if _, err := Redis.Get("test").Result(); err != nil && err.Error() != "redis: nil" {
		fmt.Println(err)
		os.Exit(-1)
	}
}
