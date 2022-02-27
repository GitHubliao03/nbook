package model

type Book struct {
	Bid       int    `gorm:"column:bid" json:"bid" form:"bid"`
	BName     string `gorm:"column:bname" json:"bname" form:"bname"`
	Author    string `gorm:"column:author" json:"author" form:"author"`
	Inventory int    `gorm:"column:inventory" json:"inventory" form:"inventory"`
}

func (b *Book) TableName() string {
	return "library.book"
}
