package conf

import (
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func LoadRedis() {
	if Redis == nil {
		Redis = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	}
}
