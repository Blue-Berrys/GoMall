package email

import (
	"github.com/Blue-Berrys/GoMall/app/email/infra/mq"
	"github.com/Blue-Berrys/GoMall/app/email/infra/notify"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/email"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func ConsumerInit() {
	sub, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		//收到的消息是protobuf格式的
		var req email.EmailReq
		err := proto.Unmarshal(msg.Data, &req) //反序列化得到能看的消息
		if err != nil {
			klog.Error(err)
		}

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
