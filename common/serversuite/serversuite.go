package serversuite

import (
	"github.com/Blue-Berrys/GoMall/common/mtl"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

// CommonServerSuite 是一个通用的服务配置结构体，主要用于配置与服务相关的元信息和选项。
type CommonServerSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

// 返回一组 server.Option，用于配置服务运行时的行为
func (s CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		server.WithMetaHandler(transmeta.ClientHTTP2Handler), //设置一个元数据处理器 (meta handler)，用来处理 HTTP2 客户端的元信息。
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ //配置服务的基本信息，通过 rpcinfo.EndpointBasicInfo 结构体传入
			ServiceName: s.CurrentServiceName,
		}),
		//配置服务的追踪器（Tracer），用于监控和记录服务的性能指标
		server.WithTracer(prometheus.NewServerTracer("", "",
			prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))),
		// "" 和 "" 通常是服务名称和环境标识的占位符，因为我们已经启动了自定义的metricsServer，所以这里置空。
		//prometheus.WithDisableServer(true)：禁用了服务器端的某些默认追踪功能（如自动记录某些指标）。
		//prometheus.WithRegistry(mtl.Registry)：指定 Prometheus 使用的注册表，用于存储和暴露自定义指标

		// 增加追踪服务
		server.WithSuite(tracing.NewServerSuite()),
	}

	r, err := consul.NewConsulRegister(s.RegistryAddr)
	if err != nil {
		klog.Fatal(err) //会立即终止程序执行
	}
	opts = append(opts, server.WithRegistry(r))

	return opts
}
