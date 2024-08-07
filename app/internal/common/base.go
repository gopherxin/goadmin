package common

import "time"

// BaseModel 基础模型
type BaseModel struct {
	ID         int64     `gorm:"primary_key;auto_increment;not_null"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP;not null"`
	UpdateTime *time.Time
	//默认为空
	DeleteTime *time.Time
	Creator    int64
	Updater    int64
}
