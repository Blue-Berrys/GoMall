package service

import (
	"context"
	"testing"
	product "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product"
)

func TestBatchGetProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewBatchGetProductService(ctx)
	// init req and assert value

	req := &product.BatchGetProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
