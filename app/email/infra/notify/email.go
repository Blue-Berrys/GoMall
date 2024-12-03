package notify

import (
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/email"
	"github.com/kr/pretty"
)

// 模拟一个发送邮件

type NoopEmail struct {
}

func (e *NoopEmail) Send(req *email.EmailReq) {
	pretty.Printf("%v\n", req)
	// pretty 是一个第三方库，常用于格式化输出复杂的 Go 数据结构
}

func NewNoopEmail() NoopEmail {
	return NoopEmail{}
}
