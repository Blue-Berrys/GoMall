package service

import (
	"context"
	"github.com/Blue-Berrys/GoMall/app/checkout/infra/rpc"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/cart"
	checkout "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/checkout"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/email"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/order"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/payment"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
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
		klog.Error(err.Error())
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Items == nil {
		klog.Error(err.Error())
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}

	//拿到每个商品的价格
	var total float32
	var ois []*order.OrderItem
	for _, cartItem := range cartResult.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: cartItem.ProductId})
		// 在循环内重复使用rpc调用，性能影响很大，需要在for循环外面一次性调用，然后遍历赋值
		if resultErr != nil {
			klog.Error(err.Error())
			return nil, resultErr
		}
		if productResp.Product == nil {
			continue
		}
		cost := productResp.Product.Price * float32(cartItem.Quantity)
		total += cost
		ois = append(ois, &order.OrderItem{
			Item: &cart.CartItem{
				Quantity:  cartItem.Quantity,
				ProductId: cartItem.ProductId,
			},
			Cost: cost,
		})
	}

	var orderId string
	orderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId: req.UserId,
		Email:  req.Email,
		Address: &checkout.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		},
		Items: ois,
	})
	if err != nil {
		klog.Error(err.Error())
		return nil, kerrors.NewGRPCBizStatusError(5004002, err.Error())
	}
	if orderResp != nil && orderResp.Order != nil {
		orderId = orderResp.Order.OrderId
	}
	//支付
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, &payment.ChargeReq{
		UserId:     req.UserId,
		OrderId:    orderId,
		CreditCard: req.CreditCard,
		Amount:     total,
	})
	if err != nil {
		klog.Error(err.Error())
		return nil, err
	}
	klog.Info(paymentResult)
	resp = &checkout.CheckoutResp{OrderId: orderId, TransactionId: paymentResult.TransactionId}

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
	return resp, nil
}
