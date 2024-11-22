package service

import (
	"context"
	"errors"
	"github.com/Blue-Berrys/GoMall/app/user/biz/dal/mysql"
	"github.com/Blue-Berrys/GoMall/app/user/biz/model"
	user "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterRep) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" || req.PasswordConfirm == "" {
		return nil, errors.New("email or password is empty")
	}
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("password not match")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		Email:    req.Email,
		Password: string(passwordHash),
	}
	if err = model.Create(mysql.DB, newUser); err != nil {
		return nil, err
	}
	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
