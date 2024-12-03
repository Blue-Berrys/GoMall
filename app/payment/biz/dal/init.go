package dal

import (
	"github.com/Blue-Berrys/GoMall/app/payment/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
