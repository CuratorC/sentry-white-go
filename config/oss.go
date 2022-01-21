// Package config 站点配置信息
package config

import "github.com/curatorc/cngf/config"

func init() {
	config.Add("oss", func() map[string]interface{} {
		return map[string]interface{}{

			"access_key_id":     config.Env("ALIYUN_ACCESS_ID"),
			"access_key_secret": config.Env("ALIYUN_ACCESS_SECRET"),
			"endpoint":          config.Env("OSS_ENDPOINT", "https://oss-cn-hangzhou.aliyuncs.com"),
			"bucket":            config.Env("OSS_BUCKET", "sentry-white-api"),
		}
	})
}
