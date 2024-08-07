package user

import "github.com/gopherxin/goadmin/app/internal/biz"

var _ biz.SystemUserRepo = (*userRepo)(nil)
