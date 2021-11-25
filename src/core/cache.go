package core

import (
	"github.com/muesli/cache2go"
)

var Cache *cache2go.CacheTable

func loadCache(){
	if Cache == nil {
		Cache = cache2go.Cache("myCache")
	}
}
