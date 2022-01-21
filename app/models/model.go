// Package models 模型通用属性和方法
package models

import (
	"time"

	"github.com/spf13/cast"
)

// BaseModel 模型基类
type BaseModel struct {
	ID          uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
	Status      int8   `gorm:"column:status;type:int;index;" json:"status,omitempty"`
	AdminRemark string `gorm:"column:admin_remark;default:null;" json:"admin_remark,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
	DeletedAt time.Time `gorm:"column:deleted_at;index;" json:"deleted_at,omitempty"`
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}
