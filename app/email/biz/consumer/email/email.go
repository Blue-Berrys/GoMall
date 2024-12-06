package email

import (
	"context"
	"github.com/Blue-Berrys/GoMall/app/email/infra/mq"
	"github.com/Blue-Berrys/GoMall/app/email/infra/notify"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/email"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
)

func ConsumerInit() {
	// 获取一个 Tracer 对象，用于创建和操作追踪（Trace）数据
	tracer := otel.Tracer("shop-nats-consumer")
	sub, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		//收到的消息是protobuf格式的
		var req email.EmailReq
		err := proto.Unmarshal(msg.Data, &req) //反序列化得到能看的消息
		if err != nil {
			klog.Error(err)
		}

		ctx := context.Background()
		// 获取 OpenTelemetry 的传播器（Propagator）。
		// 传播器的作用：
		// 从上游服务（或请求头部）中提取分布式追踪的上下文信息（如 Trace ID、Span ID）。
		// 将这些上下文信息注入到当前服务的 Context，使分布式追踪链路保持完整。
		ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(msg.Header))
		// 使用 Tracer 创建一个新的 Span
		// 继承了追踪上下文的 Context
		_, span := tracer.Start(ctx, "shop-email-consumer")
		defer span.End()
		// span: 一个表示追踪片段的对象
		//作用：
		//记录当前服务执行的一段操作（如消费消息、处理任务）。
		//如果追踪链路中存在父追踪（Parent Span），当前 Span 会自动关联到父追踪，形成完整的调用链。

		noopEmail := notify.NewNoopEmail()
		noopEmail.Send(&req)
	})
	if err != nil {
		panic(err)
	}
	// 服务关闭的时候要把订阅取消
	server.RegisterShutdownHook(func() {
		sub.Unsubscribe()
		mq.Nc.Close()
	})
}
