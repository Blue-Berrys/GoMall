package rpc

import (
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/user/echoservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"sync"
)

var (
	userClient echoservice.Client
	once       sync.Once
)

func Init() {
	once.Do(func() {

	})
}

func iniUserClient() {
	r, err := consul.NewConsulRegister("127.0.0.1:8500")
	if err != nil {
		hlog.Fatal(err)
	}
	Userclient, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
}
