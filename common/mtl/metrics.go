package mtl

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"net/http"
)

var Registry *prometheus.Registry // 注册中心

func InitMetric(serviceName, metricsPort, registryAddr string) (registry.Registry, *registry.Info) {
	Registry = prometheus.NewRegistry()
	//注册 Go 语言运行时的默认指标，例如：内存分配（go_memstats_alloc_bytes）GC 次数（go_gc_duration_seconds）协程数量（go_goroutines）
	Registry.MustRegister(collectors.NewGoCollector())
	//注册与当前进程相关的指标，例如：进程 CPU 使用率（process_cpu_seconds_total）进程内存使用（process_resident_memory_bytes
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	r, err := consul.NewConsulRegister(registryAddr)
	fmt.Println(r, err)
	// 定义服务注册信息
	addr, _ := net.ResolveTCPAddr("tcp", metricsPort) // 将端口号（metricsPort）解析为 TCP 地址对象
	registryInfo := &registry.Info{                   // 定义 Prometheus 服务的注册信息
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": serviceName},
	}
	_ = r.Register(registryInfo)

	// 注册一个服务关闭的钩子函数，用于在服务关闭时从 Consul 中注销该服务。
	server.RegisterShutdownHook(func() {
		r.Deregister(registryInfo)
	})
	// 暴露指标接口
	// promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}) 创建了一个 Prometheus 指标处理器，用于暴露指标。
	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	// 启动一个 HTTP 服务器，监听 metricsPort
	go http.ListenAndServe(metricsPort, nil)
	return r, registryInfo
}
