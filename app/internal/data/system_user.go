package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gopherxin/goadmin/app/internal/biz"
)

type systemUserRepo struct {
	data *Data
	log  *log.Helper
}

func NewSystemUserRepo(data *Data, logger log.Logger) biz.SystemUserRepo {
	return &systemUserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (u *systemUserRepo) CreateUser(ctx context.Context, m *biz.SystemUser) (int64, error) {
	user := &biz.SystemUser{
		Username: m.Username,
		Nickname: m.Nickname,
		Mobile:   m.Mobile,
		Password: m.Password,
	}
	tx := u.data.db.Create(&user)
	return user.ID, tx.Error
}
