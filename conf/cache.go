package imgo

import (
	"github.com/muesli/cache2go"
)

var Cache *cache2go.CacheTable

func LoadCache(){
	if Cache == nil {
		Cache = cache2go.Cache("myCache")
	}
}
