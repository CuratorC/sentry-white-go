// Package bootstrap 启动程序功能
package bootstrap

import (
	"fmt"
	"github.com/curatorc/cngf/cache"
	"github.com/curatorc/cngf/config"
)

// SetupCache 缓存
func SetupCache() {
	// 初始化缓存专用的 redis client, 使用专属缓存 DB
	cacheDriver := config.GetString("cache.driver", "file")
	var ds cache.Store
	if cacheDriver == "file" {
		ds = cache.NewFileStore()
	} else {
		ds = cache.NewRedisStore(
			fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
			config.GetString("redis.username"),
			config.GetString("redis.password"),
			config.GetInt("redis.database_cache"),
		)
	}
	cache.InitWithCacheStore(ds)
}
