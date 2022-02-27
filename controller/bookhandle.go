package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"nbook/dao"
	"nbook/model"
	"nbook/utils"
	"net/http"
	"strconv"
)

//GetBooks 获取所有图书
func GetBooks(c *gin.Context) {

	book, err := dao.GetBooks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get books",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"book": book,
		})
	}

}

//AddBook 添加图书
func AddBook(c *gin.Context) {

	bname := c.PostForm("bname")
	author := c.PostForm("author")
	inventory, _ := strconv.Atoi(c.PostForm("inventory"))

	if book, _ := dao.CheckBookName(bname); book.Bid > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The book already exists",
		})
	} else {
		err := dao.SaveBook(bname, author, inventory)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to add book",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Success to add book",
			})
		}
	}
}

//DeleteBook 删除图书
func DeleteBook(c *gin.Context) {
	bname := c.PostForm("bname")
	err := dao.DeleteBookName(bname)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Success to delete",
		})
	}
}

//ModifyInventory 修改库存
func ModifyInventory(c *gin.Context) {

	bid, _ := strconv.Atoi(c.PostForm("bid"))
	increment, _ := strconv.Atoi(c.PostForm("increment"))

	err := dao.UpdateInventory(bid, increment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Success to modify inventory",
		})
	}
}

//BorrowBook 根据用户ID以及图书ID新增一条借阅记录
func BorrowBook(c *gin.Context) {

	uid, _ := strconv.Atoi(c.PostForm("uid"))
	bid, _ := strconv.Atoi(c.PostForm("bid"))

	_, err := dao.CheckBorrowUserIDAndBookID(uid, bid)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "无法借阅与尚未归还的书相同的书",
		})
		return
	}

	//else if err != nil && err != gorm.ErrRecordNotFound {
	//	c.JSON(http.StatusBadGateway, gin.H{
	//		"error": "服务器内部错误",
	//	})
	//	return
	//}

	var book model.Book
	err = utils.Db.Where("bid = ?", bid).First(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else if book.Inventory < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "inventory is not enough",
		})
		return
	}

	err = dao.UpdateInventory(bid, -1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = dao.NewBorrow(uid, bid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Success to borrow",
		})
		return
	}

}

//ReturnBook 根据用户ID以及图书ID完结借阅
func ReturnBook(c *gin.Context) {

	uid, _ := strconv.Atoi(c.PostForm("uid"))
	bid, _ := strconv.Atoi(c.PostForm("bid"))

	borrow, err := dao.CheckBorrowUserIDAndBookID(uid, bid)
	if err != nil && err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "没有相关借阅信息",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "服务器内部错误",
		})
	}

	err = dao.UpdateInventory(bid, 1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "服务器内部错误",
		})
		return
	}

	err = dao.FinishBorrow(borrow.BrID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Success to return book",
		})
	}
}
