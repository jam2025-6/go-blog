package global

import (
	"time"

	"gorm.io/gorm"
)

type MODEL struct {
	ID        uint           `json:"id" gotm:"primarykey"` // 主键id
	CreatedAt time.Time      `json:"created_at"`           // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`           // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`       // 删除时间 软删除字段
}
