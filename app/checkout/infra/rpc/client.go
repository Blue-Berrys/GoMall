package rpc

import (
	"github.com/Blue-Berrys/GoMall/app/checkout/conf"
	checkoutUtils "github.com/Blue-Berrys/GoMall/app/checkout/utils"
	"github.com/Blue-Berrys/GoMall/common/clientsuite"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/email/emailservice"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	EmailClient   emailservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
	err           error
)

func InitClient() {
	once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
		initEmailClient()
	})
}

func initCartClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServerName: ServiceName,
			RegistryAddr:      RegistryAddr,
		}),
	}

	CartClient, err = cartservice.NewClient("cart", opts...)
	checkoutUtils.MustHandleError(err)
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServerName: ServiceName,
			RegistryAddr:      RegistryAddr,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	checkoutUtils.MustHandleError(err)
}

func initPaymentClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServerName: ServiceName,
			RegistryAddr:      RegistryAddr,
		}),
	}

	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	checkoutUtils.MustHandleError(err)
}

func initOrderClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServerName: ServiceName,
			RegistryAddr:      RegistryAddr,
		}),
	}

	OrderClient, err = orderservice.NewClient("order", opts...)
	checkoutUtils.MustHandleError(err)
}

func initEmailClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServerName: ServiceName,
			RegistryAddr:      RegistryAddr,
		}),
	}

	EmailClient, err = emailservice.NewClient("email", opts...)
	checkoutUtils.MustHandleError(err)
}
