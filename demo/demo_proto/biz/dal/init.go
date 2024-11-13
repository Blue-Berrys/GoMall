package dal

import (
	"github.com/Blue-Berrys/GoMall/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
