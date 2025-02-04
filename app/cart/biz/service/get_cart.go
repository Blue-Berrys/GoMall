package service

import (
	"context"
	"github.com/Blue-Berrys/GoMall/app/cart/biz/dal/mysql"
	"github.com/Blue-Berrys/GoMall/app/cart/biz/model"
	cart "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	list, err := model.GetCartByUserId(s.ctx, mysql.DB, req.UserId)
	var items []*cart.CartItem
	for _, item := range list {
		items = append(items, &cart.CartItem{
			ProductId: item.ProductId,
			Quantity:  item.Qty,
		})
	}
	//fmt.Println("items:", items)
	return &cart.GetCartResp{Items: items}, nil
}
