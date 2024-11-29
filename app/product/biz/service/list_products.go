package service

import (
	"context"
	"github.com/Blue-Berrys/GoMall/app/product/biz/dal/mysql"
	"github.com/Blue-Berrys/GoMall/app/product/biz/model"
	product "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductReq) (resp *product.ListProductResp, err error) {
	// Finish your business logic.
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)
	categories, err := categoryQuery.GetProductsByCategoryName(req.CategoryName)
	resp = &product.ListProductResp{}
	for _, category := range categories { // 首先遍历分类列表
		for _, p := range category.Products { //遍历分类包含的商品
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(p.ID),
				Picture:     p.Picture,
				Price:       p.Price,
				Description: p.Description,
				Name:        p.Name,
			})
		}
	}
	return resp, nil
	//resp *product.ListProductResp：只是声明了变量，未分配内存，初始值为 nil。
	//resp = &product.ListProductResp{}：分配了一个实例，初始化了 resp，可以安全地使用。
}
