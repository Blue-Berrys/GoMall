package dal

import (
	"github.com/Blue-Berrys/GoMall/app/order/biz/dal/mysql"
	"github.com/Blue-Berrys/GoMall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
