package dal

import (
	"github.com/Blue-Berrys/GoMall/demo/demo_thrift/biz/dal/mysql"
	"github.com/Blue-Berrys/GoMall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
