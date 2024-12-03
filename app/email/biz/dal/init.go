package dal

import (
	"github.com/Blue-Berrys/GoMall/app/email/biz/dal/mysql"
	"github.com/Blue-Berrys/GoMall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
