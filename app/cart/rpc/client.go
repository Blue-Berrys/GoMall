package rpc

import (
	"github.com/Blue-Berrys/GoMall/app/cart/conf"
	cartUtils "github.com/Blue-Berrys/GoMall/app/cart/utils"
	"github.com/Blue-Berrys/GoMall/common/clientsuite"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
	err           error
)

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServerName: ServiceName,
			RegistryAddr:      RegistryAddr,
		}),
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartUtils.MustHandleError(err)
}
