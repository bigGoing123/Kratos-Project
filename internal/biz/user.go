package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
	v1 "kratosTestApp/api/user/v1"
	"kratosTestApp/internal/conf"
	"kratosTestApp/internal/data/model"
	"kratosTestApp/internal/pkg/middlewire"
)

type UserUsecase struct {
	repo     UserRepo
	log      *log.Helper
	authConf *conf.Auth
}

func NewUserUsecase(repo UserRepo, logger log.Logger, auth *conf.Auth) *UserUsecase {
	return &UserUsecase{
		repo:     repo,
		log:      log.NewHelper(logger),
		authConf: auth,
	}
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *model.User) (*v1.NullReply, error)
	GetAllUser(ctx context.Context, in *v1.NullRequest) (*v1.GetAllUserReply, error)
	FindByUsername(ctx context.Context, user *model.User) (*model.User, error)
}

func (uc *UserUsecase) Register(ctx context.Context, in *v1.RegisterRequest) (*v1.NullReply, error) {
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

func (uc *UserUsecase) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	u := in.User
	user := &model.User{
		Username: u.Username,
		Email:    u.Email,
	}
	user1, err := uc.repo.FindByUsername(ctx, user)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user1.Password), []byte(u.Password))
	if err != nil {
		return nil, errors.New(401, "Password is incorrect", "密码错误")
	}
	secret := uc.authConf.Jwt.Secret
	token := middlewire.GenerateToken(secret, user.Username)
	return &v1.LoginReply{Token: token}, nil
}
