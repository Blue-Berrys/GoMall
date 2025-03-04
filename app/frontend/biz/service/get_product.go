package service

import (
	"context"
	"fmt"
	"github.com/Blue-Berrys/GoMall/app/frontend/infra/rpc"
	"github.com/cloudwego/hertz/pkg/common/utils"

	product "github.com/Blue-Berrys/GoMall/app/frontend/hertz_gen/frontend/product"
	rpcproduct "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	// todo edit your code
	p, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: req.Id})
	fmt.Println(p)
	if err != nil {
		return nil, err
	}
	return utils.H{"item": p.Product}, err
}
