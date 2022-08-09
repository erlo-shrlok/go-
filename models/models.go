package models

import "MrWang/dao"

//Book Model
type Book struct {
	ID        int     `form:"id"`
	BookName  string  `form:"bookname"`
	Type      string  `form:"type"`
	Author    string  `form:"author"`
	Publisher string  `form:"publisher"`
	Price     float64 `form:"price"`
	Shuliang  int     `form:"shuliang"`
}

//Addbook
func Addbook(book *Book) (err error) {
	err = dao.DB.Create(&book).Error
	return
}

//Findbook
func Findbook(book *[]Book, bookname string) (err error) {
	err = dao.DB.Where("book_name=?", bookname).Find(&book).Error
	return
}

//FindbookByauthor
func FindbookByauthor(book *[]Book, author string) (err error) {
	err = dao.DB.Where("author=?", author).Find(&book).Error
	return
}

//FindbookBytype
func FindbookBytype(book *[]Book, ttype string) (err error) {
	err = dao.DB.Where("type=?", ttype).Find(&book).Error
	return
}

//FindbookBypub
func FindbookBypub(book *Book, publisher string) (err error) {
	err = dao.DB.Where("publisher=?", publisher).Find(&book).Error
	return
}

//BorrowRecord Model
type BorrowRecord struct {
	BorrowMan  string `form:"borrowman"`
	BorrowDate string `form:"borrowdate"`
	BorrowBook string `form:"borrowBook"`
}

func CreateBorrowRecord(borrow *BorrowRecord) (err error) {
	err = dao.DB.Create(&borrow).Error
	return
}

func Reduce(borrow *BorrowRecord) (err error) {
	err = dao.DB.Model(&Book{}).Where("book_name = ?", borrow.BorrowBook).Update("shuliang", "0").Error
	return
}

func PrintBorrow(borrow *[]BorrowRecord) (err error) {
	err = dao.DB.Find(&borrow).Error
	return
}

func Status(book *[]Book) (err error) {
	err = dao.DB.Select("book_name,price , shuliang").Find(&book).Error
	return
}
