package rpc

import (
	"github.com/Blue-Berrys/GoMall/app/checkout/conf"
	checkoutUtils "github.com/Blue-Berrys/GoMall/app/checkout/utils"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/email/emailservice"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	EmailClient   emailservice.Client
	once          sync.Once
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
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoutUtils.MustHandleError(err)

	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	CartClient, err = cartservice.NewClient("cart", opts...)
	checkoutUtils.MustHandleError(err)
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoutUtils.MustHandleError(err)

	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	checkoutUtils.MustHandleError(err)
}

func initPaymentClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoutUtils.MustHandleError(err)

	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	//可扩展性强：你可以通过 opts 增加更多配置，例如负载均衡策略、自定义拦截器等。
	//细粒度控制：明确配置了传输协议（GRPC）和元数据处理器（HTTP/2）。
	//服务标识明确：通过 client.WithClientBasicInfo 设置服务名，方便服务端识别客户端。
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	checkoutUtils.MustHandleError(err)
}

func initOrderClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoutUtils.MustHandleError(err)

	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	OrderClient, err = orderservice.NewClient("order", opts...)
	checkoutUtils.MustHandleError(err)
}

func initEmailClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoutUtils.MustHandleError(err)

	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	EmailClient, err = emailservice.NewClient("email", opts...)
	checkoutUtils.MustHandleError(err)
}
