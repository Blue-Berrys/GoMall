package mq

import (
	"github.com/Blue-Berrys/GoMall/app/email/conf"
	"github.com/nats-io/nats.go"
)

var (
	Nc  *nats.Conn
	err error
)

func Init() {
	Nc, err = nats.Connect(conf.GetConf().Nats.Address)
	if err != nil {
		panic(err)
		//这里初始化连接是mq核心的地方，如果它失败了，mq后面都不起作用，所以直接panic
	}
}
