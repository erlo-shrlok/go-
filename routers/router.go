package routers

import (
	"MrWang/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//创建静态文件的映射
	r.Static("/xxx", "static")
	//解析文件
	r.LoadHTMLGlob("template/*")
	//主页
	r.GET("/", controller.IndexHandler)
	//添加图书
	r.GET("/add", controller.AddbookView)
	r.POST("/add", controller.Addbook)
	//查询图书
	cGroup := r.Group("c")
	{
		//按书名
		cGroup.GET("/find", controller.FindbookBynameView)
		cGroup.POST("/find", controller.FindbookByname)
		//按作者
		cGroup.GET("/findByauthor", controller.FindbookByauthorView)
		cGroup.POST("/findByauthor", controller.FindbookByauthor)
		//按类别
		cGroup.GET("/findBytype", controller.FindbookBytypeView)
		cGroup.POST("/findBytype", controller.FindbookBytype)
		//按出版社
		cGroup.GET("/findBypub", controller.FindbookBypubView)
		cGroup.POST("/findBypub", controller.FindbookBypub)
	}
	//修改信息
	r.GET("/mf", controller.ModifyView)
	r.PUT("/mf", controller.Modify)
	//外借
	bGroup := r.Group("b")
	{
		//外借索引页
		bGroup.GET("/borrow", controller.BorrowView)
		//登记外借
		bGroup.GET("/b1", controller.RegisterBorrowView)
		bGroup.POST("/b1", controller.RegisterBorrow)
		//打印外借情况
		bGroup.GET("/b2", controller.Print)
	}
	//统计信息（购买金额、册数）
	r.GET("/t", controller.StaBorrow)
	return r
}
