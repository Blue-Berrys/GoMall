package dal

import (
	"github.com/Blue-Berrys/GoMall/app/frontend/biz/dal/mysql"
	"github.com/Blue-Berrys/GoMall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
