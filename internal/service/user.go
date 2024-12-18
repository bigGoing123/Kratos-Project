package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratosTestApp/api/user/v1"
	"kratosTestApp/internal/biz"
)

type UserService struct {
	v1.UnimplementedUserServer
	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Register(ctx context.Context, in *v1.RegisterRequest) (*v1.NullReply, error) {
	return s.uc.Register(ctx, in)
}

func (s *UserService) GetAllUser(ctx context.Context, in *v1.NullRequest) (*v1.GetAllUserReply, error) {
	return s.uc.GetAllUser(ctx, in)
}
func (s *UserService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	return s.uc.Login(ctx, in)
}
