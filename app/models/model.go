// Package models 模型通用属性和方法
package models

import (
	"encoding/json"
	"github.com/curatorc/cngf/logger"
	"sentry-white-go/app/handlers/oss"
	"time"

	"github.com/spf13/cast"
)

type IModel interface {
	GetStringID() string
}

// BaseModel 模型基类
type BaseModel struct {
	ID          uint64 `json:"id,omitempty"`
	Status      int8   `json:"status,omitempty"`
	AdminRemark string `json:"admin_remark,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// BaseCollection 模型集合
type BaseCollection struct {
	MaxID uint64 `json:"max_id,omitempty"`
}

// GetStringID 获取 ID 的字符串格式
func (bm BaseModel) GetStringID() string {
	return cast.ToString(bm.ID)
}

func UploadToOss(path string, collection interface{}, newModel IModel) {
	oss.Upload(path, collection)
	oss.Upload(path+"/"+newModel.GetStringID(), newModel)
}
func DeleteOnOss(path string, collection interface{}, newModel IModel) {
	oss.Upload(path, collection)
	oss.Delete(path + "/" + newModel.GetStringID())
}

// GetModelFromOSS 从 OSS 中获取单个模型信息
func GetModelFromOSS(path string, model interface{}) {
	if oss.IsExist(path) {
		response := oss.Get(oss.SignURL(path))
		err := json.Unmarshal([]byte(response), &model)
		logger.LogIf(err)
	}
}
