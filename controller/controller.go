package controller

import (
	"MrWang/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func AddbookView(c *gin.Context) {
	c.HTML(200, "add.html", nil)
}

func Addbook(c *gin.Context) {
	//前端填写书籍信息，点击提交，会发请求到这里
	//1.从请求中取出数据
	var book models.Book
	err := c.ShouldBind(&book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(book)
	//2.存入数据库 3.返回响应
	err = models.Addbook(&book)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
	} else {
		c.JSON(http.StatusOK, book)
	}
}

func FindbookBynameView(c *gin.Context) {
	c.HTML(200, "find.html", nil)
}

func FindbookByname(c *gin.Context) {
	var book []models.Book
	//前端填写书名，点击查询，发送书名到后端
	//1.从请求中取出书名
	bookname := c.PostForm("bookname")
	fmt.Println(bookname)
	//2.从数据库中查询，并返回book表单
	err := models.Findbook(&book, bookname)
	if err != nil {
		c.JSON(200, gin.H{"err": "出现错误"})
	} else {
		c.JSON(200, book)
	}
}

func FindbookByauthorView(c *gin.Context) {
	c.HTML(200, "findByauthor.html", nil)
}

func FindbookByauthor(c *gin.Context) {
	var book []models.Book
	//前端填写书名，点击查询，发送书名到后端
	//1.从请求中取出书名
	author := c.PostForm("author")
	fmt.Println(author)
	//2.从数据库中查询，并返回book表单
	err := models.FindbookByauthor(&book, author)
	if err != nil {
		c.JSON(200, gin.H{"err": "出现错误"})
	} else {
		c.JSON(200, book)
	}
}

func FindbookBytypeView(c *gin.Context) {
	c.HTML(200, "findBytype.html", nil)
}

func FindbookBytype(c *gin.Context) {
	var book []models.Book
	//前端填写书名，点击查询，发送书名到后端
	//1.从请求中取出书名
	ttype := c.PostForm("type")
	fmt.Println(ttype)
	//2.从数据库中查询，并返回book表单
	err := models.FindbookBytype(&book, ttype)
	if err != nil {
		c.JSON(200, gin.H{"err": "出现错误"})
	} else {
		c.JSON(200, book)
	}
}

func FindbookBypubView(c *gin.Context) {
	c.HTML(200, "findBypub.html", nil)
}

func FindbookBypub(c *gin.Context) {
	var book models.Book
	//前端填写书名，点击查询，发送书名到后端
	//1.从请求中取出书名
	publisher := c.PostForm("publisher")
	fmt.Println(publisher)
	//2.从数据库中查询，并返回book表单
	err := models.FindbookBypub(&book, publisher)
	if err != nil {
		c.JSON(200, gin.H{"err": "出现错误"})
	} else {
		c.HTML(200, "book.html", book)
	}
}

func ModifyView(c *gin.Context) {
	c.HTML(200, "mf.html", nil)
}

func Modify(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

func BorrowView(c *gin.Context) {
	c.HTML(200, "borrow.html", nil)
}

func RegisterBorrowView(c *gin.Context) {
	c.HTML(200, "b1.html", nil)
}

func RegisterBorrow(c *gin.Context) {
	//前端填写外借信息，点击提交，会发送请求到这里
	//1.从请求中把数据拿出来
	var borrow models.BorrowRecord
	c.ShouldBind(&borrow)
	fmt.Println(borrow)
	//2.存入数据库3.返回响应
	err := models.CreateBorrowRecord(&borrow)
	if err != nil {
		c.JSON(200, gin.H{"erroer": "存入数据库错误"})
	} else {
		c.JSON(200, borrow)
		//根据书的id在book表中-1
		err := models.Reduce(&borrow)
		if err != nil {
			c.JSON(200, gin.H{"error": "借书失败"})
		} else {
			c.JSON(200, gin.H{"status": "借书成功"})
		}
	}
}

func Print(c *gin.Context) {
	//查询borrow表内所有数据
	var borrow []models.BorrowRecord
	err := models.PrintBorrow(&borrow)
	if err != nil {
		c.JSON(200, gin.H{"error": "获取数据库内容错误"})
	} else {
		c.JSON(200, borrow)
	}
}

func StaBorrow(c *gin.Context) {
	var book []models.Book
	//从数据库中取出数据
	err := models.Status(&book)
	if err != nil {
		c.JSON(200, gin.H{"error": "取数据出错"})
	} else {
		c.JSON(200, book)
	}
}
