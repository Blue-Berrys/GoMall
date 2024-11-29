package service

import (
	"context"
	"github.com/Blue-Berrys/GoMall/app/product/biz/dal/mysql"
	product "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
	"testing"
)

func TestGetProduct_Run(t *testing.T) {
	//var dbName string
	//mysql.DB.Raw("SELECT DATABASE()").Scan(&dbName)
	//log.Println("Connected to database:", dbName)
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewGetProductService(ctx)
	// init req and assert value

	req := &product.GetProductReq{
		Id: 1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
