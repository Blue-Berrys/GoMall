package service

import (
	"context"
	"github.com/Blue-Berrys/GoMall/app/order/biz/dal/mysql"
	"github.com/Blue-Berrys/GoMall/app/order/biz/model"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/cart"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/checkout"
	order "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	list_orders, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(500001, err.Error())
	}
	var orders []*order.Order
	for _, v := range list_orders {
		var items []*order.OrderItem
		for _, oi := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: oi.ProductId,
					Quantity:  oi.Quantity,
				},
				Cost: oi.Cost,
			})
		}
		orders = append(orders, &order.Order{
			OrderId:      v.OrderId,
			Userid:       v.UserId,
			UserCurrency: v.UserCurrency,
			Email:        v.Consignee.Email,
			Address: &checkout.Address{
				StreetAddress: v.Consignee.StreetAddress,
				State:         v.Consignee.State,
				City:          v.Consignee.City,
				ZipCode:       v.Consignee.ZipCode,
				Country:       v.Consignee.Country,
			},
			Item: items,
		})
	}
	resp = &order.ListOrderResp{
		Orders: orders,
	}
	return
}
