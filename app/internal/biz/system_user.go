package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gopherxin/goadmin/app/internal/common"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

// SystemUser  用户
type SystemUser struct {
	common.BaseModel
	Mobile   string
	Nickname string
	Password string
	Username string
}

type SystemUserRepo interface {
	CreateUser(ctx context.Context, u *SystemUser) (int64, error)
}

type SystemUserUseCase struct {
	repo SystemUserRepo
	log  *log.Helper
	tm   Transaction
}

func NewSystemUserUseCase(repo SystemUserRepo, logger log.Logger) *SystemUserUseCase {
	return &SystemUserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "goadmin/system/user"))}
}

func (su *SystemUserUseCase) CreateUser(ctx context.Context, u *SystemUser) (int64, error) {
	id, err := su.repo.CreateUser(ctx, u)
	return id, err
}
