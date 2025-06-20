package home

import (
	"context"

	"github.com/Blue-Berrys/GoMall/app/frontend/biz/service"
	"github.com/Blue-Berrys/GoMall/app/frontend/biz/utils"
	common "github.com/Blue-Berrys/GoMall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Home .
// @router / [GET]
func Home(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	//resp := &home.Empty{}
	resp, err := service.NewHomeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp = utils.WarpResponse(ctx, c, resp)
	c.HTML(consts.StatusOK, "home.templ", resp) //还会加载resp的数据

	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
