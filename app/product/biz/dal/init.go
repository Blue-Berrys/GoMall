package dal

import (
	"github.com/Blue-Berrys/GoMall/app/product/biz/dal/mysql"
	"github.com/Blue-Berrys/GoMall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
