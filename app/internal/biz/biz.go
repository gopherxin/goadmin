package biz

import (
	"context"
	"github.com/google/wire"
	"time"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewSystemUserUseCase)

// Transaction 事务
type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

// BaseModel 基础模型
type BaseModel struct {
	ID         int64 `gorm:"primary_key;auto_increment;not_null"`
	CreateTime time.Time
	UpdateTime time.Time
	//默认为空
	DeleteTime *time.Time
	Creator    int64
	Updater    int64
}
