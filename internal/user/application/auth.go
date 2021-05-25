package application

import (
	"context"
	"grpc-go-templete/internal/user/helper"
	"grpc-go-templete/pkg/pb/user_pb"
)

func (s *usersGrpcServer) Login(ctx context.Context, in *user_pb.LoginRequest) (*user_pb.LoginResponse, error) {
	result, err := s.authService.Login(in.Username, in.Password)
	if err != nil {
		return nil, err
	}
	return &user_pb.LoginResponse{
		Profile: helper.ToUserGRPC(result.Profile),
		Token:   result.Token,
	}, nil
}