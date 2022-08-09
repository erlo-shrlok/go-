package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	err error
	DB  *gorm.DB
)

func InitMySQL() (err error) {
	u := "root:root@(localhost)/lib?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", u)
	if err != nil {
		return
	}
	return DB.DB().Ping() //测试连通性
}
