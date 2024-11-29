package rpc

import (
	"fmt"
	"github.com/Blue-Berrys/GoMall/app/cart/conf"
	cartUtils "github.com/Blue-Berrys/GoMall/app/cart/utils"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	//r, err := consul.NewConsulResolver("127.0.0.1:8500")
	cartUtils.MustHandleError(err)
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	cartUtils.MustHandleError(err)
	fmt.Println(err)
}
