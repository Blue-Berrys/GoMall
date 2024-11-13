package service

import (
	"context"
	common "github.com/Blue-Berrys/GoMall/app/frontend/hertz_gen/frontend/common"

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
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	var resp = make(map[string]any)
	resp["title"] = "Hot Sales"
	items := []map[string]any{
		{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/t-shirt-1.png"},
		{"Name": "T-shirt-2", "Price": 110, "Picture": "/static/image/t-shirt-1.png"},
		{"Name": "T-shirt-3", "Price": 120, "Picture": "/static/image/t-shirt-2.png"},
		{"Name": "T-shirt-4", "Price": 130, "Picture": "/static/image/t-shirt-2.png"},
	}
	resp["items"] = items
	return resp, nil
}
