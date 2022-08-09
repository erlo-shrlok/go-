package main

import (
	"MrWang/dao"
	"MrWang/models"
	"MrWang/routers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//type BookList struct {
//}
//
////向数据库中添加图书
//func (b BookList) Add() {
//
//}
//
//func (b BookList) Modify() {
//
//}
//
//func (b BookList) query() {
//
//}
//
//func (b BookList) count() {
//
//}
//
//
//type ItBook struct {
//	Book
//}
//
//type OtherBook struct {
//	Book
//}
//
//
//type BorrowList struct {
//}
//
//func (b BorrowList) Add() {
//
//}
//
//func (b BorrowList) Remove() {
//
//}
//
//func (b BorrowList) Print() {
//
//}

func main() {
	//创建sql数据库lib
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()
	//模型绑定
	dao.DB.AutoMigrate(&models.Book{})
	dao.DB.AutoMigrate(&models.BorrowRecord{})

	r := routers.SetupRouter()
	r.Run(":9000")
}
