package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Db  *gorm.DB
	err error
)

func init() {

	// "用户名:密码@(127.0.0.1:端口号)/数据库名称?charset=utf8&parseTime=True&loc=Local"
	dns := "root:123456@(127.0.0.1:3306)/library?charset=utf8&parseTime=True&loc=Local"
	Db, err = gorm.Open("mysql", dns)
	Db.LogMode(true)
	if err != nil {
		panic(err)
	}
	//defer Db.Close()

}
