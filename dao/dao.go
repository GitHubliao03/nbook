package dao

import (
	"gorm.io/gorm"
)

type BookDao struct {
	Db gorm.DB
	gorm.Model
}

func NewDao(db gorm.DB) *BookDao {
	return &BookDao{Db: db}
}

// add

// update

// sort id price
//
type Res struct {
	Name  string
	Price string
}

// redis
// find first redis key is exits , not mysql set redis key
// add update mysql  set redis return
