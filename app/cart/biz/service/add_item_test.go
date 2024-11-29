package service

import (
	"context"
	"github.com/Blue-Berrys/GoMall/app/cart/biz/dal/mysql"
	"github.com/Blue-Berrys/GoMall/app/cart/rpc"
	cart "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/cart"
	"github.com/joho/godotenv"
	"testing"
)

func TestAddItem_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewAddItemService(ctx)
	// init req and assert value
	rpc.InitClient()
	req := &cart.AddItemReq{
		UserId: 10,
		Item:   &cart.CartItem{ProductId: 1, Quantity: 5},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
