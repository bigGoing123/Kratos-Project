package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
	v1 "kratosTestApp/api/user/v1"
	"kratosTestApp/internal/data/model"
)

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *model.User) (*v1.RegisterReply, error)
	GetAllUser(ctx context.Context, in *v1.NullRequest) (*v1.GetAllUserReply, error)
}

func (uc *UserUsecase) Register(ctx context.Context, in *v1.RegisterRequest) (*v1.RegisterReply, error) {
	//对密码进行加密
	u := in.User
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Username: u.Username,
		Password: string(hashedPassword),
		Email:    u.Email,
	}
	reply, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (uc *UserUsecase) GetAllUser(ctx context.Context, in *v1.NullRequest) (*v1.GetAllUserReply, error) {
	return uc.repo.GetAllUser(ctx, in)
}
