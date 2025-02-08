package main

import (
	"context"
	"github.com/Blue-Berrys/GoMall/app/product/biz/dal"
	"github.com/Blue-Berrys/GoMall/common/mtl"
	"github.com/Blue-Berrys/GoMall/common/serversuite"
	"github.com/joho/godotenv"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"time"

	"github.com/Blue-Berrys/GoMall/app/product/conf"
	"github.com/Blue-Berrys/GoMall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	MetricsPort  = conf.GetConf().Kitex.MetricsPort
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func main() {
	//_ = godotenv.Load("/opt/gomall/product/.env")
	_ = godotenv.Load(".env")
	mtl.InitMetric(ServiceName, MetricsPort, RegistryAddr)
	p := mtl.InitTracing(ServiceName)
	defer p.Shutdown(context.Background()) //会把链路数据上传完再关闭
	dal.Init()
	opts := kitexInit()

	svr := productcatalogservice.NewServer(new(ProductCatalogServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	//consul
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	// server.WithSuite是配置suite用的
	// server.WithSuite suite 必须实现 Options() 方法即 CommonServerSuite.Options() 方法并将返回的选项追加到服务配置中
	opts = append(opts, server.WithSuite(serversuite.CommonServerSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
