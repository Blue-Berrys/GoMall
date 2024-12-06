package rpc

import (
	"github.com/Blue-Berrys/GoMall/app/frontend/conf"
	frontendUtils "github.com/Blue-Berrys/GoMall/app/frontend/utlis"
	"github.com/Blue-Berrys/GoMall/common/clientsuite"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/user/echoservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	UserClient     echoservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
	ServiceName    = frontendUtils.ServiceName
	RegistryAddr   = conf.GetConf().Hertz.RegistryAddr
	err            error
)

func Init() {
	once.Do(func() {
		iniUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
	})
}

func iniUserClient() {
	UserClient, err = echoservice.NewClient("user", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServerName: ServiceName,
		RegistryAddr:      RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServerName: ServiceName,
		RegistryAddr:      RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServerName: ServiceName,
		RegistryAddr:      RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServerName: ServiceName,
		RegistryAddr:      RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServerName: ServiceName,
		RegistryAddr:      RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}
