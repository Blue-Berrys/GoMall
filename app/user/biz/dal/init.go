package dal

import (
	"github.com/Blue-Berrys/GoMall/app/user/biz/dal/mysql"
	"github.com/Blue-Berrys/GoMall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
