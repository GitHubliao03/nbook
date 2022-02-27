package model

import (
	"time"
)

type Borrow struct {
	BrID       int       `gorm:"column:brid" json:"brid" form:"brid"`
	UID        int       `gorm:"column:uid" json:"uid" form:"uid"`
	BID        int       `gorm:"column:bid" json:"bid" form:"bid"`
	BorrowDate time.Time `gorm:"column:borrow_date" json:"borrow_date" form:"borrow_date"`
	ReturnDate time.Time `gorm:"column:return_date" json:"return_date" form:"return_date"`
	RealDate   time.Time `gorm:"column:real_date" json:"real_date" form:"real_date"`
}

func (br *Borrow) TableName() string {
	return "borrow"
}
