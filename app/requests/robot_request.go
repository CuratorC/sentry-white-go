package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type RobotRequest struct {
	Name        string `json:"name" valid:"name" form:"name"`
	AccessToken string `json:"access_token" valid:"access_token" form:"access_token"`
	ProjectID   uint64 `json:"project_id" valid:"project_id" form:"project_id"`
}

func RobotSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":         []string{"required", "min_cn:2", "max_cn:255"},
		"access_token": []string{"required", "min_cn:3", "max_cn:255"},
		"project_id":   []string{"required", "id_in_oss:projects"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 255 个字",
		},
		"access_token": []string{
			"required:密钥为必填项",
			"min_cn:密钥长度需至少 3 个字",
			"max_cn:密钥长度不能超过 255 个字",
		},
		"project_id": []string{
			"required:项目 ID 为必填项",
			"id_in_oss:项目不存在",
		},
	}
	return validate(data, rules, messages)
}
