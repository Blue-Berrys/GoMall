package main

import (
	"fmt"
	"github.com/Blue-Berrys/GoMall/demo/demo_proto/biz/dal"
	"github.com/Blue-Berrys/GoMall/demo/demo_proto/biz/dal/mysql"
	"github.com/Blue-Berrys/GoMall/demo/demo_proto/model"
	"github.com/joho/godotenv"
)

// 测试数据库连接
func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	dal.Init()
	//mysql.DB.Create(&model.User{Email: "demo@example.com", Password: "aaa"})
	mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").Update("password", "222")

	row := model.User{}
	mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").First(&row)
	fmt.Println(row)

	mysql.DB.Where("email = ?", "demo@example.com").Delete(&model.User{})
}
