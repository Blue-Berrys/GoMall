package service

import (
	"context"
	"fmt"
	"github.com/Blue-Berrys/GoMall/app/checkout/biz/dal/kafka"
	"github.com/Blue-Berrys/GoMall/app/checkout/infra/rpc"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/cart"
	checkout "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/checkout"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/email"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/order"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// 首先获取购物车商品
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})

	if err != nil {
		fmt.Println(req.UserId, cartResult)
		//klog.Info(req.UserId, cartResult)
		klog.Error(err.Error())
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Items == nil {
		klog.Error(err.Error())
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}

	// 重写一个批量rpc调用,只要做一次调用
	IdList := []uint32{}
	for _, cartItem := range cartResult.Items {
		IdList = append(IdList, cartItem.ProductId)
	}
	productsResp, err := rpc.ProductClient.BatchGetProduct(s.ctx, &product.BatchGetProductReq{
		Id: IdList,
	})
	if err != nil {
		klog.Error(err.Error())
		return nil, kerrors.NewGRPCBizStatusError(5005002, err.Error())
	}
	// 将产品列表转化为map,才能有索引
	products := make(map[uint32]*product.Product)
	for _, product := range productsResp.Product {
		products[product.Id] = product
	}

	//拿到每个商品的价格
	var total float32
	var ois []*order.OrderItem
	for _, cartItem := range cartResult.Items {
		prod := products[cartItem.ProductId]
		// 在循环内重复使用rpc调用，性能影响很大，需要在for循环外面一次性调用，然后遍历赋值
		if prod == nil {
			continue
		}
		cost := prod.Price * float32(cartItem.Quantity)
		total += cost
		ois = append(ois, &order.OrderItem{
			Item: &cart.CartItem{
				Quantity:  cartItem.Quantity,
				ProductId: cartItem.ProductId,
			},
			Cost: cost,
		})
	}

	address := kafka.Address{
		StreetAddress: req.Address.StreetAddress,
		City:          req.Address.City,
		Country:       req.Address.Country,
		State:         req.Address.State,
		ZipCode:       req.Address.ZipCode,
	}
	creditCardInfo := kafka.CreditCardInfo{
		CreditCardCvv:             req.CreditCard.CreditCardCvv,
		CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
		CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
		CreditCardNumber:          req.CreditCard.CreditCardNumber,
	}
	// 结账后放入kafka生产者，分发给下单服务、支付服务、邮件服务
	orderId, _ := uuid.NewRandom() // 结账时就有orderId
	checkoutRespChan := make(chan *checkout.CheckoutResp)
	go kafka.WriteOrder(orderId.String(), ois, req.UserId, req.Email, address, &creditCardInfo, total)

	go kafka.ReadOrder()

	go kafka.ReadPayment(checkoutRespChan)

	// 发送邮件消息
	_, err = rpc.EmailClient.Send(s.ctx, &email.EmailReq{
		From:        "from@example.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "You have just created an order in the CloudWeGo Shop",
		Content:     "You have just created an order in the CloudWeGo Shop",
	})
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	//清空购物车
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err.Error())
	}
	for resp := range checkoutRespChan {
		fmt.Printf("OrderId: %v TransId %v", resp.OrderId, resp.TransactionId)
		return resp, nil
	}
	return resp, nil
}
