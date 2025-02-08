package service

import (
	"context"
	"github.com/Blue-Berrys/GoMall/app/product/biz/dal/mysql"
	"github.com/Blue-Berrys/GoMall/app/product/biz/dal/redis"
	"github.com/Blue-Berrys/GoMall/app/product/biz/model"
	product "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type BatchGetProductService struct {
	ctx context.Context
} // NewBatchGetProductService new BatchGetProductService
func NewBatchGetProductService(ctx context.Context) *BatchGetProductService {
	return &BatchGetProductService{ctx: ctx}
}

// Run create note info
func (s *BatchGetProductService) Run(req *product.BatchGetProductReq) (resp *product.BatchGetProductResp, err error) {
	// Finish your business logic.
	resp = &product.BatchGetProductResp{}
	products := []*product.Product{}
	for productId := range req.Id {
		if productId == 0 {
			klog.Infof("one product id is 0")
			continue
		}
		productQuery := model.NewCachedProductQuery(s.ctx, mysql.DB, redis.RedisClient)
		p, err := productQuery.GetById(productId)
		if err != nil {
			klog.Infof("get product by id err:%v", err)
			continue
		}
		oneProduct := product.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Picture:     p.Picture,
			Description: p.Description,
			Price:       p.Price,
		}
		products = append(products, &oneProduct)
	}
	resp = &product.BatchGetProductResp{
		Product: products,
	}
	return resp, nil
}
