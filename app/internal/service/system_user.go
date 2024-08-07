package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gopherxin/goadmin/app/internal/biz"

	pb "github.com/gopherxin/goadmin/api/system/user/service/v1"
)

func NewSystemUserService(service *biz.SystemUserUseCase, logger log.Logger) *SystemUserService {
	return &SystemUserService{
		service: service,
		log:     log.NewHelper(logger),
	}
}

func (s *SystemUserService) CreateSystemUser(ctx context.Context, req *pb.CreateSystemUserRequest) (*pb.CreateSystemUserReply, error) {
	s.log.Infof("input data %v", req)
	user, err := s.service.CreateUser(ctx, &biz.SystemUser{
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		Password: req.Password,
		Username: req.Username,
	})
	fmt.Println("user", user)
	if err != nil {

	}
	return &pb.CreateSystemUserReply{}, err
}
func (s *SystemUserService) UpdateSystemUser(ctx context.Context, req *pb.UpdateSystemUserRequest) (*pb.UpdateSystemUserReply, error) {
	return &pb.UpdateSystemUserReply{}, nil
}
func (s *SystemUserService) DeleteSystemUser(ctx context.Context, req *pb.DeleteSystemUserRequest) (*pb.DeleteSystemUserReply, error) {
	return &pb.DeleteSystemUserReply{}, nil
}
func (s *SystemUserService) GetSystemUser(ctx context.Context, req *pb.GetSystemUserRequest) (*pb.GetSystemUserReply, error) {
	return &pb.GetSystemUserReply{}, nil
}
func (s *SystemUserService) ListSystemUser(ctx context.Context, req *pb.ListSystemUserRequest) (*pb.ListSystemUserReply, error) {

	return &pb.ListSystemUserReply{}, nil
}
