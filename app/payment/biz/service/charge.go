package service

import (
	"context"
	"github.com/Blue-Berrys/GoMall/app/payment/biz/dal/mysql"
	"github.com/Blue-Berrys/GoMall/app/payment/biz/model"
	payment "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    req.CreditCard.CreditCardCvv,
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}
	err = card.Validate(true) // true代表可以使用测试卡号
	if err != nil {
		klog.Error(err.Error())
	}
	transactionId, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005002, err.Error())
	}
	err = model.CreatePaymentLog(s.ctx, mysql.DB, &model.PaymentLog{
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		TransactionId: transactionId.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})
	//fmt.Println(err)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005002, err.Error())
	}
	return &payment.ChargeResp{TransactionId: transactionId.String()}, nil
}
