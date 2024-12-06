package mysql

import (
	"fmt"
	"github.com/Blue-Berrys/GoMall/app/product/biz/model"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/plugin/opentelemetry/tracing"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	//dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"))
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/product?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT")) //test测试的时候要用这个
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" { // 如果不是线上环境
		needDemoData := !DB.Migrator().HasTable(&model.Product{}) // 不存在表则为1
		//fmt.Println(needDemoData)
		err = DB.AutoMigrate(&model.Product{}, &model.Category{}) // 不止会创建这两个表，还会创建中间表，如果表不存在，会创建表。 如果表已存在但结构不同，会自动更新表的结构。
		if err != nil {
			fmt.Println(6666)
			klog.Error(err)
		}
		if needDemoData {
			DB.Exec("INSERT INTO `product`.`category` VALUES (1,'2023-12-06 15:05:06','2023-12-06 15:05:06','T-Shirt','T-Shirt'),(2,'2023-12-06 15:05:06','2023-12-06 15:05:06','Sticker','Sticker')")
			DB.Exec("INSERT INTO `product`.`product` VALUES ( 1, '2023-12-06 15:26:19', '2023-12-09 22:29:10', 'Notebook', 'The cloudwego notebook is a highly efficient and feature-rich notebook designed to meet all your note-taking needs. ', '/static/image/notebook.png', 9.90 ), ( 2, '2023-12-06 15:26:19', '2023-12-09 22:29:10', 'Mouse-Pad', 'The cloudwego mouse pad is a premium-grade accessory designed to enhance your computer usage experience. ', '/static/image/mouse-pad.png', 8.80 ), ( 3, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-1.png', 6.60 ), ( 4, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-1.png', 2.20 ), ( 5, '2023-12-06 15:26:19', '2023-12-09 22:32:35', 'Sweatshirt', 'The cloudwego Sweatshirt is a cozy and fashionable garment that provides warmth and style during colder weather.', '/static/image/t-shirt-2.png', 1.10 ), ( 6, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-2.png', 1.80 ), ( 7, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'mascot', 'The cloudwego mascot is a charming and captivating representation of the brand, designed to bring joy and a playful spirit to any environment.', '/static/image/logo.jpg', 4.80 )")
			DB.Exec("INSERT INTO `product`.`product_category` (product_id,category_id) VALUES ( 1, 2 ), ( 2, 2 ), ( 3, 1 ), ( 4, 1 ), ( 5, 1 ), ( 6, 1 ),( 7, 2 )")
		}
	}
}
