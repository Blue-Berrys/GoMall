package main

import (
	"context"
	user "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/user"
	"github.com/Blue-Berrys/GoMall/app/user/biz/service"
)

// EchoServiceImpl implements the last service interface defined in the IDL.
type EchoServiceImpl struct{}

// Register implements the EchoServiceImpl interface.
func (s *EchoServiceImpl) Register(ctx context.Context, req *user.RegisterRep) (resp *user.RegisterResp, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// Login implements the EchoServiceImpl interface.
func (s *EchoServiceImpl) Login(ctx context.Context, req *user.LoginRep) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}
