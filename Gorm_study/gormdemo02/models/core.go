package models

//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "root:123456@tcp(192.168.0.6:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true, //打印sql
	})
	// DB.Debug()
	if err != nil {
		fmt.Println(err)
	}
}
