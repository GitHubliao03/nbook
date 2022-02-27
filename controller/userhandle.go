package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"nbook/dao"
	"nbook/model"
	"net/http"
)

func Login(c *gin.Context) {

	user := &model.Users{}
	var err error

	user.UName = c.PostForm("user_name")
	user.UPassword = c.PostForm("user_password")

	user, err = dao.CheckUserNameAndPassword(user.UName, user.UPassword)
	if user.UID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"uid":          user.UID,
			"uname":        user.UName,
			"phone_number": user.PhoneNumber,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
}

func Regist(c *gin.Context) {

	user := &model.Users{}
	var err error

	userName := c.PostForm("user_name")
	phoneNumber := c.PostForm("phone_number")
	userPassword := c.PostForm("user_password")

	user, err = dao.CheckUserName(user.UName)
	if user.UID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"err": "用户已经存在"})
	} else {
		err = dao.SaveUser(userName, phoneNumber, userPassword)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "succeed",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}
	}

}

//CheckUserName 发送Ajax验证用户名
func CheckUserName(c *gin.Context) {

	userName := c.PostForm("user_name")
	_, err := dao.CheckUserName(userName)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	} else if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "用户名不可用",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "用户名可用",
		})
	}

}
