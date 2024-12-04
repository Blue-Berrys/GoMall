package service

import (
	"context"
	"github.com/Blue-Berrys/GoMall/app/email/infra/mq"
	email "github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/email"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type SendService struct {
	ctx context.Context
} // NewSendService new SendService
func NewSendService(ctx context.Context) *SendService {
	return &SendService{ctx: ctx}
}

// Run create note info
func (s *SendService) Run(req *email.EmailReq) (resp *email.EmailResp, err error) {
	// Finish your business logic.
	data, _ := proto.Marshal(&email.EmailReq{
		From:        req.From,
		To:          req.To,
		Content:     req.Content,
		ContentType: req.ContentType,
		Subject:     req.Subject,
	})
	msg := &nats.Msg{Subject: "email", Data: data} // Subject是主题，主题作为索引
	_ = mq.Nc.PublishMsg(msg)
	// NATS 是一种发布-订阅模式的消息系统，当你使用 mq.Nc.PublishMsg(msg) 发布消息到 email 主题时，
	// 所有订阅了该主题的消费者（如你的 ConsumerInit 函数中定义的订阅者）都会自动收到消息并执行对应的处理逻辑。
	return
}
