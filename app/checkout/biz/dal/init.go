package dal

import (
	"github.com/Blue-Berrys/GoMall/app/checkout/biz/dal/kafka"
	"github.com/Blue-Berrys/GoMall/app/checkout/biz/dal/mysql"
	"github.com/Blue-Berrys/GoMall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
	kafka.Init()
}
