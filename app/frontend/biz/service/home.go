package service

import (
	"context"
	common "github.com/Blue-Berrys/GoMall/app/frontend/hertz_gen/frontend/common"
	"github.com/Blue-Berrys/GoMall/app/frontend/infra/rpc"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	// todo edit your code
	//var resp = make(map[string]any)
	//resp["title"] = "Hot Sales"
	//items := []map[string]any{
	//	{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/t-shirt-1.png"},
	//	{"Name": "T-shirt-2", "Price": 110, "Picture": "/static/image/t-shirt-1.png"},
	//	{"Name": "T-shirt-3", "Price": 120, "Picture": "/static/image/t-shirt-2.png"},
	//	{"Name": "T-shirt-4", "Price": 130, "Picture": "/static/image/t-shirt-2.png"},
	//}
	//resp["items"] = items
	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductReq{})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title": "Hot Sale",
		"items": products.Products,
	}, nil
}
