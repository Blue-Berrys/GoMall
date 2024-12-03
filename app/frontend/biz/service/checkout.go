package service

import (
	"context"
	"github.com/Blue-Berrys/GoMall/app/frontend/infra/rpc"
	frontendUtils "github.com/Blue-Berrys/GoMall/app/frontend/utlis"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/cart"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"

	common "github.com/Blue-Berrys/GoMall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *common.Empty) (resp map[string]any, err error) {
	// todo edit your code
	// 展示要结算的商品列表
	var items []map[string]any
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	carts, err := rpc.CartClient.GetCart(h.Context, &cart.GetCartReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}
	var total float32
	for _, v := range carts.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: v.ProductId})
		if err != nil {
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]any{
			"Name":    p.Name,
			"Price":   strconv.FormatFloat(float64(p.Price), 'f', 2, 64), //'f': 使用定点表示法（如 123.45）
			"Picture": p.Picture,
			"Qty":     strconv.Itoa(int(v.Quantity)),
		})
		total += p.Price * float32(v.Quantity)
	}
	return utils.H{
		"title": "Checkout",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
