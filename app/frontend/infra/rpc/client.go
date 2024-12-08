package rpc

import (
	"context"
	"github.com/Blue-Berrys/GoMall/app/frontend/conf"
	frontendUtils "github.com/Blue-Berrys/GoMall/app/frontend/utlis"
	"github.com/Blue-Berrys/GoMall/common/clientsuite"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/user/echoservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	consulclient "github.com/kitex-contrib/config-consul/client"
	"github.com/kitex-contrib/config-consul/consul"
	"sync"
)

var (
	UserClient      echoservice.Client
	ProductClient   productcatalogservice.Client
	CartClient      cartservice.Client
	CheckoutClient  checkoutservice.Client
	OrderClient     orderservice.Client
	once            sync.Once
	ServiceName     = frontendUtils.ServiceName
	RegistryAddr    = conf.GetConf().Hertz.RegistryAddr
	err             error
	consulClient, _ = consul.NewClient(consul.Options{})
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
	//NewCBSuite：创建一个熔断器套件实例，用于管理多个熔断器
	//func(ri rpcinfo.RPCInfo) string 一个用于生成熔断键的函数, 通过 RPCInfo2Key，从 rpcinfo.RPCInfo（RPC 调用信息）生成唯一的熔断键
	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return circuitbreak.RPCInfo2Key(ri) // 返回一个字符串，包括fromServiceName,toServiceName,method，这三个之间用斜杆拼接起来
	})
	cbs.UpdateServiceCBConfig("frontend/product/GetProduct", circuitbreak.CBConfig{
		Enable:    true, // 开启熔断器
		ErrRate:   0.5,  // 错误率阈值：50%, 错误率 = 错误请求数 / 总请求数
		MinSample: 2,    // 最小请求数：至少 2 次请求后才开始统计错误率
	})
	ProductClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServerName: ServiceName,
		RegistryAddr:      RegistryAddr,
	}), client.WithCircuitBreaker(cbs), client.WithFallback(
		fallback.NewFallbackPolicy(
			fallback.UnwrapHelper(func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
				if err == nil {
					return resp, nil // 如果没错误，直接返回原始的
				}
				methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
				if methodName != "ListProducts" { // 如果当前调用的方法不是 ListProducts，直接返回原始的响应
					return resp, nil
				}
				// 如果是 ListProducts 方法且调用失败，使用降级策略，返回固定的产品数据
				return &product.ListProductResp{
					Products: []*product.Product{
						{
							Price:       6.6,
							Id:          3,
							Picture:     "/static/image/t-shirt.png",
							Name:        "T-Shirt",
							Description: "CloudWeGo T-shirt",
						},
					},
				}, nil
			}),
		),
	), client.WithSuite(consulclient.NewSuite("product", ServiceName, consulClient)))
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
