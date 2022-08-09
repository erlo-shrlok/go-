package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type BookList struct {
}

//向数据库中添加图书
func (b BookList) Add() {

}

func (b BookList) Modify() {

}

func (b BookList) query() {

}

func (b BookList) count() {

}

type Book struct {
	ID        int     `form:"id"`
	BookName  string  `form:"bookname"`
	Type      string  `form:"type"`
	Author    string  `form:"author"`
	Publisher string  `form:"publisher"`
	Price     float64 `form:"price"`
	Shuliang  int     `form:"shuliang"`
}

type ItBook struct {
	Book
}

type OtherBook struct {
	Book
}

type BorrowRecord struct {
	BorrowMan  string `form:"borrowman"`
	BorrowDate string `form:"borrowdate"`
	BorrowBook string `form:"borrowBook"`
}

type BorrowList struct {
}

func (b BorrowList) Add() {

}

func (b BorrowList) Remove() {

}

func (b BorrowList) Print() {

}

var (
	err error
	DB  *gorm.DB
)

func main() {
	r := gin.Default()
	//连接数据库
	u := "root:root@(localhost)/lib?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", u)
	if err != nil {
		panic(err)
		return
	}
	defer DB.Close()
	//创建Book表
	DB.AutoMigrate(&Book{})
	DB.AutoMigrate(&BorrowRecord{})
	//创建静态文件的映射
	r.Static("/xxx", "static")
	//解析文件
	r.LoadHTMLGlob("template/*")
	//主页
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	//添加图书
	r.GET("/add", func(c *gin.Context) {
		c.HTML(200, "add.html", nil)
	})
	r.POST("/add", func(c *gin.Context) {
		//前端填写书籍信息，点击提交，会发请求到这里
		//1.从请求中取出数据
		var book Book
		err := c.ShouldBind(&book)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(book)
		//2.存入数据库 3.返回响应
		if err = DB.Create(&book).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		} else {
			c.JSON(http.StatusOK, book)
		}
	})
	//查询图书
	cGroup := r.Group("c")
	{
		//按书名
		cGroup.GET("/find", func(c *gin.Context) {
			c.HTML(200, "find.html", nil)
		})
		cGroup.POST("/find", func(c *gin.Context) {
			var book []*Book
			//前端填写书名，点击查询，发送书名到后端
			//1.从请求中取出书名
			bookname := c.PostForm("bookname")
			fmt.Println(bookname)
			//2.从数据库中查询，并返回book表单
			if err = DB.Where("book_name=?", bookname).Find(&book).Error; err != nil {
				c.JSON(200, gin.H{"err": "出现错误"})
			} else {
				c.JSON(200, book)
			}
		})
		//按作者
		cGroup.GET("/findByauthor", func(c *gin.Context) {
			c.HTML(200, "findByauthor.html", nil)
		})
		cGroup.POST("/findByauthor", func(c *gin.Context) {
			var book []*Book
			//前端填写书名，点击查询，发送书名到后端
			//1.从请求中取出书名
			author := c.PostForm("author")
			fmt.Println(author)
			//2.从数据库中查询，并返回book表单
			if err = DB.Where("author=?", author).Find(&book).Error; err != nil {
				c.JSON(200, gin.H{"err": "出现错误"})
			} else {
				c.JSON(200, book)
			}
		})
		//按类别
		cGroup.GET("/findBytype", func(c *gin.Context) {
			c.HTML(200, "findBytype.html", nil)
		})
		cGroup.POST("/findBytype", func(c *gin.Context) {
			var book []*Book
			//前端填写书名，点击查询，发送书名到后端
			//1.从请求中取出书名
			ttype := c.PostForm("type")
			fmt.Println(ttype)
			//2.从数据库中查询，并返回book表单
			if err = DB.Where("type=?", ttype).Find(&book).Error; err != nil {
				c.JSON(200, gin.H{"err": "出现错误"})
			} else {
				c.JSON(200, book)
			}
		})
		//按出版社
		cGroup.GET("/findBypub", func(c *gin.Context) {
			c.HTML(200, "findBypub.html", nil)
		})
		cGroup.POST("/findBypub", func(c *gin.Context) {
			var book Book
			//前端填写书名，点击查询，发送书名到后端
			//1.从请求中取出书名
			publisher := c.PostForm("publisher")
			fmt.Println(publisher)
			//2.从数据库中查询，并返回book表单
			if err = DB.Where("publisher=?", publisher).Find(&book).Error; err != nil {
				c.JSON(200, gin.H{"err": "出现错误"})
			} else {
				c.HTML(200, "book.html", book)
			}
		})
	}
	//修改信息
	r.GET("/mf", func(c *gin.Context) {
		c.HTML(200, "mf.html", nil)
	})
	r.PUT("/mf", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	//外借
	bGroup := r.Group("b")
	{
		//外借索引页
		bGroup.GET("/borrow", func(c *gin.Context) {
			c.HTML(200, "borrow.html", nil)
		})
		//登记外借
		bGroup.GET("/b1", func(c *gin.Context) {
			c.HTML(200, "b1.html", nil)
		})
		bGroup.POST("/b1", func(c *gin.Context) {
			//前端填写外借信息，点击提交，会发送请求到这里
			//1.从请求中把数据拿出来
			var borrow BorrowRecord
			c.ShouldBind(&borrow)
			fmt.Println(borrow)
			//2.存入数据库3.返回响应
			if err = DB.Create(&borrow).Error; err != nil {
				c.JSON(200, gin.H{"erroer": "存入数据库错误"})
			} else {
				c.JSON(200, borrow)
				//根据书的id在book表中-1
				if DB.Model(&Book{}).Where("book_name = ?", borrow.BorrowBook).Update("shuliang", "0"); err != nil {
					c.JSON(200, gin.H{"error": "借书失败"})
				} else {
					c.JSON(200, gin.H{"status": "借书成功"})
				}
			}
		})
		//打印外借情况
		bGroup.GET("/b2", func(c *gin.Context) {
			//查询borrow表内所有数据
			var borrow []*BorrowRecord
			if err := DB.Find(&borrow).Error; err != nil {
				c.JSON(200, gin.H{"error": "获取数据库内容错误"})
			} else {
				c.JSON(200, borrow)
			}
		})
	}
	//统计信息（购买金额、册数）
	r.GET("/t", func(c *gin.Context) {
		var book []*Book
		//从数据库中取出数据
		if err := DB.Select("book_name,price , shuliang").Find(&book).Error; err != nil {
			c.JSON(200, gin.H{"error": "取数据出错"})
		} else {
			c.JSON(200, book)
		}
	})
	r.Run(":9000")
}
