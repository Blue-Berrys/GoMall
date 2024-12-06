package mtl

import (
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

func InitTracing(serviceName string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),           //现在用的是追加服务，暂时用非https来上报
		provider.WithEnableMetrics(false), //禁用一下指标，指标是prometheus生成的，暂时不需要otel生成
	)
	return p
}
