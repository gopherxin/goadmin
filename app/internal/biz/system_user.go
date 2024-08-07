package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

// User 用户
type SystemUser struct {
	Id       int64
	Username string
	Nickname string
	Mobile   string
	Password string
	BaseModel
}

type SystemUserRepo interface {
	CreateUser(ctx context.Context, u *SystemUser) (int64, error)
	GetUser(ctx context.Context, id int64) (*SystemUser, error)
}

type SystemUserUseCase struct {
	repo SystemUserRepo
	log  *log.Helper
	tm   Transaction
}

func NewSystemUserUseCase(repo SystemUserRepo, logger log.Logger) *SystemUserUseCase {
	return &SystemUserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "goadmin/system/user"))}
}
