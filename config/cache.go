// Package config 站点配置信息
package config

import "github.com/curatorc/cngf/config"

func init() {
	config.Add("cache", func() map[string]interface{} {
		return map[string]interface{}{
			// 缓存驱动
			"driver": config.Env("CACHE_DRIVER", "file"),
		}
	})
}
