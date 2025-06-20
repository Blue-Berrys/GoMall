package main

import (
	"context"
	product "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product"
	"github.com/Blue-Berrys/GoMall/app/product/biz/service"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductReq) (resp *product.ListProductResp, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}

// BatchGetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) BatchGetProduct(ctx context.Context, req *product.BatchGetProductReq) (resp *product.BatchGetProductResp, err error) {
	resp, err = service.NewBatchGetProductService(ctx).Run(req)

	return resp, err
}
