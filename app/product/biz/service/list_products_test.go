package service

import (
	"context"
	"github.com/Blue-Berrys/GoMall/app/product/biz/dal/mysql"
	product "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
	"testing"
)

func TestListProducts_Run(t *testing.T) {

	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewListProductsService(ctx)
	// init req and assert value

	req := &product.ListProductReq{CategoryName: "T-Shirt"}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
