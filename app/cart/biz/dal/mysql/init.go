package mysql

import (
	"fmt"
	"github.com/Blue-Berrys/GoMall/app/cart/biz/model"
	"github.com/Blue-Berrys/GoMall/app/cart/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"))
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/cart?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
	//	os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT")) //test测试的时候要用这个
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.Cart{})
	if err != nil {
		panic(err)
	}
}
