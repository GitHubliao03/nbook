package router

import (
	"github.com/gin-gonic/gin"
	"nbook/controller"
)

func StartRouter() {
	r := gin.Default()

	r.POST("/loginName", controller.Login)
	//r.POST("loginID",controller.Login)
	//r.POST("loginPhoneNumber",controller.Login)

	r.POST("/rigist", controller.Regist)

	r.POST("/checkUserName", controller.CheckUserName)

	r.GET("/book", controller.GetBooks)
	r.POST("/book", controller.AddBook)
	r.DELETE("/book", controller.DeleteBook)
	r.PUT("/book", controller.ModifyInventory)

	r.POST("/borrowBook", controller.BorrowBook)
	r.POST("/returnBook", controller.ReturnBook)

	r.Run("192.168.31.63:8888")
}
