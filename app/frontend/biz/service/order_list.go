package service

import (
	"context"
	"fmt"
	"github.com/Blue-Berrys/GoMall/app/frontend/infra/rpc"
	"github.com/Blue-Berrys/GoMall/app/frontend/types"
	frontendUtils "github.com/Blue-Berrys/GoMall/app/frontend/utlis"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/order"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"time"

	common "github.com/Blue-Berrys/GoMall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	orderResp, err := rpc.OrderClient.ListOrder(h.Context, &order.ListOrderReq{
		UserId: uint32(userId),
	})
	fmt.Println(orderResp)
	// 在实际生产中，会通过组装productId，去批量调取rpc获取商品信息，组装成map后返回组装商品
	var list []*types.Order
	for _, v := range orderResp.Orders { // v是每一笔订单
		var total float32
		var items []types.OrderItem
		for _, item := range v.Item {
			total += item.Cost
			p, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{
				Id: item.Item.ProductId,
			})
			if err != nil {
				return nil, err
			}
			if p == nil || p.Product == nil {
				continue
			}
			items = append(items, types.OrderItem{
				ProductName: p.Product.Name,
				Picture:     p.Product.Picture,
				Qty:         item.Item.Quantity,
				Cost:        item.Cost,
			})
		}

		list = append(list, &types.Order{
			OrderId:     v.OrderId,
			CreatedDate: time.Unix(int64(v.CreateAt), 0).Format("2006-01-02 15:04:05"),
			Cost:        total,
			Items:       items,
		})
	}
	return utils.H{
		"title":  "Order",
		"orders": list,
	}, nil
}
