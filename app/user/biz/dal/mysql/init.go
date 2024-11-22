package mysql

import (
	"fmt"
	"github.com/Blue-Berrys/GoMall/app/user/biz/model"
	"github.com/Blue-Berrys/GoMall/app/user/conf"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"))
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/user?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
	//	os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT")) 测试的时候要用这个
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
