package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "github.com/gopherxin/goadmin/api/system/user/service/v1"
	"github.com/gopherxin/goadmin/app/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewSystemUserService)

type SystemUserService struct {
	pb.UnimplementedSystemUserServer
	service *biz.SystemUserUseCase
	log     *log.Helper
}
