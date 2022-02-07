package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LocalStorageRequest struct {
	AliyunAccessID          string `json:"aliyun_access_id" valid:"aliyun_access_id" form:"aliyun_access_id"`
	AliyunAccessSecret      string `json:"aliyun_access_secret" valid:"aliyun_access_secret" form:"aliyun_access_secret"`
	AliyunAccessOSSEndpoint string `json:"aliyun_access_oss_endpoint" valid:"aliyun_access_oss_endpoint" form:"aliyun_access_oss_endpoint"`
	AliyunAccessOSSBucket   string `json:"aliyun_access_oss_bucket" valid:"aliyun_access_oss_bucket" form:"aliyun_access_oss_bucket"`
}

func LocalStorageSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"aliyun_access_id":           []string{"required", "min_cn:1", "max_cn:255"},
		"aliyun_access_secret":       []string{"required", "min_cn:1", "max_cn:255"},
		"aliyun_access_oss_endpoint": []string{"required", "min_cn:1", "max_cn:255"},
		"aliyun_access_oss_bucket":   []string{"required", "min_cn:1", "max_cn:255"},
	}
	messages := govalidator.MapData{
		"aliyun_access_id": []string{
			"required:阿里云 AccessKey 为必填项",
			"min_cn:阿里云 AccessKey 长度需至少 1 个字",
			"max_cn:阿里云 AccessKey 长度不能超过 255 个字",
		},
		"aliyun_access_secret": []string{
			"required:阿里云 AccessSecret 为必填项",
			"min_cn:阿里云 AccessSecret 长度需至少 1 个字",
			"max_cn:阿里云 AccessSecret 长度不能超过 255 个字",
		},
		"aliyun_access_oss_endpoint": []string{
			"required:OSS endpoint 为必填项",
			"min_cn:OSS endpoint 长度需至少 1 个字",
			"max_cn:OSS endpoint 长度不能超过 255 个字",
		},
		"aliyun_access_oss_bucket": []string{
			"required:OSS Bucket 为必填项",
			"min_cn:OSS Bucket 长度需至少 1 个字",
			"max_cn:OSS Bucket 长度不能超过 255 个字",
		},
	}

	errs := validate(data, rules, messages)
	return errs
}
