package category

import (
	"context"

	"github.com/Blue-Berrys/GoMall/app/frontend/biz/service"
	"github.com/Blue-Berrys/GoMall/app/frontend/biz/utils"
	"github.com/Blue-Berrys/GoMall/app/frontend/hertz_gen/frontend/category"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Category .
// @router /category/:category [GET]
func Category(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category.CategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCategoryService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)

	c.HTML(consts.StatusOK, "category", resp)
	//模版名称是category，关联category.templ
	//resp拿到的数据也要渲染上去
}
