package service

import (
	"context"
	"fmt"
	"github.com/Blue-Berrys/GoMall/app/cart/biz/dal/mysql"
	"github.com/Blue-Berrys/GoMall/app/cart/biz/model"
	"github.com/Blue-Berrys/GoMall/app/cart/rpc"
	cart "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/cart"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	// 需要调用product服务的判断产品是否为空，参照frontend/infra/rpc设计
	fmt.Println(req.Item.ProductId)
	productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}
	if productResp == nil || productResp.Product.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(40004, "product not found")
	}

	cartItem := model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       req.Item.Quantity,
	}
	err = model.AddItem(s.ctx, mysql.DB, &cartItem)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(50000, err.Error())
	}
	return &cart.AddItemResp{}, nil
}
