package service

import (
	"context"
	"github.com/Blue-Berrys/GoMall/app/frontend/infra/rpc"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/user"
	"github.com/hertz-contrib/sessions"

	auth "github.com/Blue-Berrys/GoMall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/Blue-Berrys/GoMall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterRep) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userResp, err := rpc.UserClient.Register(h.Context, &user.RegisterRep{
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	})
	if err != nil {
		return nil, err
	}
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", userResp.UserId)
	//session.Set("user_id", 1)
	//fmt.Println("Register.userResp.UserId: ", userResp.UserId)
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return
}
