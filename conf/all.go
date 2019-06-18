package conf

import "github.com/kataras/iris"

func LoadAll(c iris.Configuration) {
	loadRedis(c)
	loadCache()
	loadOrm(c)
}
