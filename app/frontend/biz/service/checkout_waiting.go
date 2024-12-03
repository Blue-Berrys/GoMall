package service

import (
	"context"
	"fmt"
	"github.com/Blue-Berrys/GoMall/app/frontend/hertz_gen/frontend/checkout"
	"github.com/Blue-Berrys/GoMall/app/frontend/infra/rpc"
	frontendUtils "github.com/Blue-Berrys/GoMall/app/frontend/utlis"
	rpccheckout "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/checkout"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	// todo edit your code
	// 真正的结算
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	p, err := rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    uint32(userId),
		Email:     req.Email,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Address: &rpccheckout.Address{
			Country:       req.Country,
			ZipCode:       req.Zipcode,
			City:          req.City,
			State:         req.Province,
			StreetAddress: req.Street,
		},
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CardNum,
			CreditCardExpirationYear:  req.ExpirationYear,
			CreditCardExpirationMonth: req.ExpirationMonth,
			CreditCardCvv:             string(req.Cvv),
		},
	})
	fmt.Println("p: ", p)
	fmt.Println(err)
	if err != nil {
		//klog.Error(err)
		return nil, err
	}
	return utils.H{
		"title":    "waiting",
		"redirect": "/checkout/result",
	}, nil
}
