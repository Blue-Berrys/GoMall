package kafka

import (
	"encoding/json"
	"github.com/Blue-Berrys/GoMall/app/checkout/infra/rpc"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/checkout"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ReadPayment(respChan chan<- *checkout.CheckoutResp) {
	defer close(respChan) // 确保通道在函数退出时被关闭
	for {
		msg, err := readerPayment.ReadMessage(ctx)
		if err != nil {
			klog.Error(err)
			continue
		}
		klog.Infof("Received message: %s", msg.Value) // 添加这行来检查消息内容

		var orderMessage OrderMessage
		_ = json.Unmarshal(msg.Value, &orderMessage)

		//jsonData, err := json.MarshalIndent(orderMessage, "", "  ")
		//if err != nil {
		//	klog.Error(err)
		//	return
		//}
		//
		//// 打印格式化后的 JSON 字符串
		//fmt.Println(string(jsonData))
		orderMessage.CreditCard = &CreditCardInfo{} // 防止空指针
		paymentResult, err := rpc.PaymentClient.Charge(ctx, &payment.ChargeReq{
			UserId:  orderMessage.UserId,
			OrderId: orderMessage.OrderId,
			CreditCard: &payment.CreditCardInfo{
				CreditCardExpirationYear:  orderMessage.CreditCard.CreditCardExpirationYear,
				CreditCardNumber:          orderMessage.CreditCard.CreditCardNumber,
				CreditCardExpirationMonth: orderMessage.CreditCard.CreditCardExpirationMonth,
				CreditCardCvv:             orderMessage.CreditCard.CreditCardCvv,
			},
			Amount: orderMessage.TotalPrice,
		})
		if err != nil {
			klog.Error(err.Error())
		}
		klog.Info(paymentResult)
		resp := &checkout.CheckoutResp{OrderId: orderMessage.OrderId, TransactionId: paymentResult.TransactionId}
		respChan <- resp // 将 resp 发送到通道
	}
}
