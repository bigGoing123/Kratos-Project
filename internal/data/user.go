package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	v1 "kratosTestApp/api/user/v1"
	"kratosTestApp/internal/biz"
	"kratosTestApp/internal/data/model"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *userRepo) GetAllUser(ctx context.Context, in *v1.NullRequest) (*v1.GetAllUserReply, error) {
	//从数据库中获取所有用户
	db := r.data.UserDb.Debug()
	var users []model.User
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	reply := &v1.GetAllUserReply{}
	copier.Copy(&reply.User, &users)
	return reply, nil
}

func (r *userRepo) CreateUser(ctx context.Context, user *model.User) (*v1.RegisterReply, error) {
	db := r.data.UserDb
	// Check if user already exists
	var existingUser model.User
	if user.Email != "" { //首先判断是否有email
		rowsAffected := db.Where("email = ?", user.Email).First(&existingUser).RowsAffected
		if rowsAffected > 0 {
			return nil, errors.New(401, "User already exists", "该邮箱已被注册")
		}
	}
	rowsAffected := db.Where("username = ?", user.Username).First(&existingUser).RowsAffected
	if rowsAffected > 0 {
		return nil, errors.New(401, "User already exists", "User already exists")
	}
	err := db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{Message: "User created successfully"}, nil

}
