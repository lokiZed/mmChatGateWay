package middleware

import (
	"github.com/zeromicro/go-zero/core/collection"
	"log"
	"time"
)

var ipCache *collection.Cache

func init() {
	expireTime := time.Duration(365*24) * time.Hour
	cache, err := collection.NewCache(expireTime, collection.WithName("ipCache"))
	if err != nil {
		log.Fatal("init ipCache fail")
	}
	ipCache = cache
}
