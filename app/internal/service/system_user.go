package service

import (
	"context"

	pb "github.com/gopherxin/goadmin/api/system/user/service/v1"
)

type SystemUserService struct {
	pb.UnimplementedSystemUserServer
}

func NewSystemUserService() *SystemUserService {
	return &SystemUserService{}
}

func (s *SystemUserService) CreateSystemUser(ctx context.Context, req *pb.CreateSystemUserRequest) (*pb.CreateSystemUserReply, error) {

	return &pb.CreateSystemUserReply{}, nil
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
