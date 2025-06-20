package main

import (
	"context"
	"fmt"
	"github.com/Blue-Berrys/GoMall/demo/demo_thrift/kitex_gen/api"
	"github.com/Blue-Berrys/GoMall/demo/demo_thrift/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
)

func main() {
	cli, err := echo.NewClient("demo_thrift", client.WithHostPorts("localhost:8888"),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "demo_thrift_client",
		}),
	)
	if err != nil {
		panic(err)
	}
	res, err := cli.Echo(context.TODO(), &api.Request{Message: "hello"})
	fmt.Println(res)
}
