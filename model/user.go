package model

type Users struct {
	UID         int    `gorm:"column:uid" json:"uid" form:"uid"`
	UName       string `gorm:"column:uname" json:"uname" form:"uname"`
	PhoneNumber string `gorm:"column:phone_number" json:"phone_number" form:"phone_number"`
	UPassword   string `gorm:"column:upassword" json:"upassword" form:"upassword"`
}

func (u *Users) TableName() string {
	return "users"
}
