package requests

import (
	"github.com/curatorc/cngf/logger"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ProjectRequest struct {
	Name                string   `json:"name" valid:"name" form:"name"`
	SubstituteName      string   `json:"substitute_name" valid:"substitute_name" form:"substitute_name"`
	OriginalID          uint64   `json:"original_id" valid:"original_id" form:"original_id"`
	ResponsiblePeopleID []uint64 `json:"responsible_people_id" valid:"responsible_people_id" form:"responsible_people_id"`
}

func ProjectSave(data interface{}, c *gin.Context) map[string][]string {

	rpID, _ := c.GetPostForm("responsible_people_id")
	logger.DebugString("project_save", "c", rpID)
	rules := govalidator.MapData{
		"name":                  []string{"required", "min_cn:2", "max_cn:255"},
		"substitute_name":       []string{"required", "min_cn:3", "max_cn:255"},
		"original_id":           []string{"required", "id_in_oss:originals"},
		"responsible_people_id": []string{"required", "slice"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 255 个字",
		},
		"substitute_name": []string{
			"required:别称为必填项",
			"min_cn:别称长度需至少 2 个字",
			"max_cn:别称长度不能超过 255 个字",
		},
		"original_id": []string{
			"required:组织 ID 为必填项",
			"id_in_oss:组织不存在",
		},
		"responsible_people_id": []string{
			"required:负责人为必填项",
			"slice:负责人必须为一个数组",
		},
	}
	return validate(data, rules, messages)
}
