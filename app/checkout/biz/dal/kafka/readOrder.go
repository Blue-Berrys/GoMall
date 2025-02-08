package kafka

import (
	"encoding/json"
	"github.com/Blue-Berrys/GoMall/app/checkout/infra/rpc"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/checkout"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ReadOrder() {
	// 读下单信息
	for {
		msg, err := readerOrder.ReadMessage(ctx)
		if err != nil {
			klog.Error(err)
			continue
		}
		var orderMessage OrderMessage
		_ = json.Unmarshal(msg.Value, &orderMessage)

		// 下单
		OrderResp, err := rpc.OrderClient.PlaceOrder(ctx, &order.PlaceOrderReq{
			Id:     orderMessage.OrderId,
			UserId: orderMessage.UserId,
			Email:  orderMessage.Email,
			Address: &checkout.Address{
				StreetAddress: orderMessage.Address.StreetAddress,
				City:          orderMessage.Address.City,
				State:         orderMessage.Address.State,
				Country:       orderMessage.Address.Country,
				ZipCode:       orderMessage.Address.ZipCode,
			},
			Items: orderMessage.Items,
		})
		if err != nil {
			klog.Error(err)
		}
		if OrderResp == nil {
			klog.Info("OrderResp is nil")
		}
		if err := readerOrder.CommitMessages(ctx, msg); err != nil {
			klog.Errorf("failed to commit messages: %v", err)
		}
	}

	defer readerOrder.Close()
}
