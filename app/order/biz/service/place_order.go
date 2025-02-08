package service

import (
	"context"
	"fmt"
	"github.com/Blue-Berrys/GoMall/app/order/biz/dal/mysql"
	"github.com/Blue-Berrys/GoMall/app/order/biz/model"
	order "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	if len(req.Items) == 0 {
		err = kerrors.NewGRPCBizStatusError(500001, "items is empty")
		return nil, err
	}

	// 收到订单请求之后，要分发给支付，邮件通知

	//涉及到两个表的操作，要用事务,全部执行成功才会执行
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		//orderId, _ := uuid.NewRandom()

		o := &model.Order{ // 单条记录创建推荐指针，会自增ID
			OrderId:      req.Id,
			UserId:       req.UserId,
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}
		if req.Address != nil {
			o.Consignee.StreetAddress = req.Address.StreetAddress
			o.Consignee.State = req.Address.State
			o.Consignee.ZipCode = req.Address.ZipCode
			o.Consignee.City = req.Address.City
			o.Consignee.Country = req.Address.Country
		}
		if err = tx.Model(&model.Order{}).Create(o).Error; err != nil {
			fmt.Println("ORDER TABLE not insert")
			klog.Errorf("Failed to create OrderItem: %v", err)
			return err
		}
		//订单表写入成功，写订单子表
		var items []model.OrderItem
		for _, v := range req.Items {
			items = append(items, model.OrderItem{
				ProductId:    v.Item.ProductId,
				Quantity:     v.Item.Quantity,
				Cost:         v.Cost,
				OrderIdRefer: o.OrderId, //外键，一定要写，不然会报错
			})
		}
		// 当批量创建时,推荐使用非指针切片,传递指针切片可能会导致混乱，因为每个指针都指向不同的对象
		if err := tx.Model(&model.OrderItem{}).Create(items).Error; err != nil {
			klog.Errorf("Failed to create OrderItem: %v", err)
			return err
		}
		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: req.Id,
			},
		}
		return nil
	})
	return
}
