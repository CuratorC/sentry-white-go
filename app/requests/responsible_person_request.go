package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ResponsiblePersonRequest struct {
	Name  string `json:"name" valid:"name" form:"name"`
	Phone string `json:"phone" valid:"phone" form:"phone"`
}

func ResponsiblePersonSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":  []string{"required", "min_cn:2", "max_cn:255"},
		"phone": []string{"required", "min_cn:3", "max_cn:255"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 255 个字",
		},
		"phone": []string{
			"required:手机号为必填项",
			"min_cn:手机号长度需至少 3 个字",
			"max_cn:手机号长度不能超过 255 个字",
		},
	}
	return validate(data, rules, messages)
}
