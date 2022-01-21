package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"sentry-white-go/app/requests/validators"
)

type OriginalRequest struct {
	// Name        string `valid:"name" json:"name"`
	// Description string `valid:"description" json:"description,omitempty"`
	Name        string `json:"name" valid:"name" form:"name"`
	AccountName string `json:"account_name" valid:"account_name" form:"account_name"`
	Password    string `json:"password" valid:"password" form:"password"`
}

func OriginalSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name": []string{"required", "min_cn:1", "max_cn:20"},
		// "name":         []string{"required", "min_cn:1", "max_cn:20", "not_exists:originals,name"},
		"account_name": []string{"required", "min_cn:3", "max_cn:255"},
		"password":     []string{"required", "min_cn:3", "max_cn:255"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 1 个字",
			"max_cn:名称长度不能超过 20 个字",
		},
		"account_name": []string{
			"required:账号名为必填项",
			"min_cn:账号长度需至少 3 个字",
			"max_cn:账号长度不能超过 255 个字",
		},
		"password": []string{
			"required:密码为必填项",
			"min_cn:密码长度需至少 3 个字",
			"max_cn:密码长度不能超过 255 个字",
		},
	}

	errs := validate(data, rules, messages)

	errs = validators.ValidatePasswordConfirm("123456", "123456", errs)
	return errs
}
