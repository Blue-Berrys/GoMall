package kafka

import (
	"encoding/json"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/segmentio/kafka-go"
)

func WriteOrder(orderId string, orders []*order.OrderItem, userId uint32, email string, address Address, CreditCard *CreditCardInfo, totalPrice float32) {
	// 一个订单存进kafka

	orderMessage := OrderMessage{
		OrderId:    orderId,
		UserId:     userId,
		Email:      email,
		Address:    address,
		Items:      orders,
		CreditCard: CreditCard,
		TotalPrice: totalPrice,
	}

	orderBytes, _ := json.Marshal(orderMessage)
	msg := kafka.Message{
		Key:   []byte(orderId),
		Value: orderBytes,
	}

	// 只尝试写入一次

	if err := writer.WriteMessages(ctx, msg); err != nil {
		klog.Errorf("Failed to write order message: %v", err)
		return
	}

	klog.Infof("Order Write Successful Id %s", orderId) // 使用 Infof 格式化日志
}
