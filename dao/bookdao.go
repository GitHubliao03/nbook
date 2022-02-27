package dao

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"nbook/model"
	"nbook/utils"
)

// GetBooks 获取数据库中所有的图书
func GetBooks() ([]*model.Book, error) {

	var book []*model.Book
	err := utils.Db.Find(&book).Error
	if err != nil {
		fmt.Println("获取所有图书失败")
	}

	return book, err
}

//CheckBookNameAndAuthor 根据书名与作者名从数据库查询数据
func CheckBookNameAndAuthor(bname string, author string) (*model.Book, error) {
	book := &model.Book{}
	err := utils.Db.Where("bname = ? AND author >= ?", bname, author).Find(&book).Error
	return book, err
}

//CheckBookName 根据书名从数据库中查询数据
func CheckBookName(bname string) (*model.Book, error) {
	book := &model.Book{}
	err := utils.Db.Where("bname = ?", bname).Find(&book).Error
	return book, err
}

//SaveBook 向数据库中插入书籍信息
func SaveBook(bname string, author string, inventory int) error {
	book := model.Book{
		BName:     bname,
		Author:    author,
		Inventory: inventory,
	}

	//通过Error对象检测，插入数据有没有成功，如果没有错误那就是数据写入成功了。
	err := utils.Db.Create(book).Error
	return err
}

//DeleteBookName 根据书名删除数据信息
func DeleteBookName(bname string) error {

	book := model.Book{}
	utils.Db.Where("bname = ?", bname).Take(&book)
	err := utils.Db.Delete(&book).Error
	return err

}

//UpdateInventory 根据传入参数更改库存
func UpdateInventory(bid int, increment int) error {

	book := &model.Book{}
	err := utils.Db.Where("bid = ?", bid).Find(&book).Error
	if err != nil {
		return err
	} else if book.Inventory+increment < 0 {
		err = errors.New("not enough books")
		return err
	}

	err = utils.Db.Model(&book).Where("bid = ?", bid).Update("inventory", gorm.Expr("inventory + ?", increment)).Error
	return err
}
