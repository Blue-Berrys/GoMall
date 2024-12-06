package clientsuite

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonClientSuite struct {
	CurrentServerName string
	RegistryAddr      string
}

func (s CommonClientSuite) Options() []client.Option {
	opts := []client.Option{
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServerName,
		}),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithTransportProtocol(transport.GRPC),
		//可扩展性强：你可以通过 opts 增加更多配置，例如负载均衡策略、自定义拦截器等。
		//细粒度控制：明确配置了传输协议（GRPC）和元数据处理器（HTTP/2）。
		//服务标识明确：通过 client.WithClientBasicInfo 设置服务名，方便服务端识别客户端。

		//增加追踪服务
		client.WithSuite(tracing.NewClientSuite()),
	}

	// consul 解析放这里
	r, err := consul.NewConsulResolver(s.RegistryAddr)
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))

	return opts
}
