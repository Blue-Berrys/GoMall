package service

import (
	"context"
	"github.com/Blue-Berrys/GoMall/app/frontend/infra/rpc"
	frontendUtils "github.com/Blue-Berrys/GoMall/app/frontend/utlis"

	cart "github.com/Blue-Berrys/GoMall/app/frontend/hertz_gen/frontend/cart"
	common "github.com/Blue-Berrys/GoMall/app/frontend/hertz_gen/frontend/common"
	rpccart "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp *common.Empty, err error) {
	// todo edit your code
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  uint32(req.ProductNum),
		},
	})
	if err != nil {
		return nil, err
	}
	return
}
