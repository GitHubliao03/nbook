package dao

import (
	"nbook/model"
	"nbook/utils"
	"time"
)

////CheckBorrowUserID 根据用户ID查找借阅记录
//func CheckBorrowUserID(uid int) ([]*model.Borrow, error) {
//	var borrows []*model.Borrow
//	err := utils.Db.Where("uid = ?", uid).Find(&borrows).Error
//	return borrows, err
//}
//
////CheckBorrowBookID 根据书ID查找借阅记录
//func CheckBorrowBookID(bid int) ([]*model.Borrow, error) {
//	var borrows []*model.Borrow
//	err := utils.Db.Where("bid = ?", bid).Find(&borrows).Error
//	return borrows, err
//}

//CheckBorrowUserIDAndBookID 根据用户ID以及图书ID查找未完成的借阅记录
func CheckBorrowUserIDAndBookID(uid int, bid int) (model.Borrow, error) {
	var borrows model.Borrow

	err := utils.Db.Where("uid = ? AND bid = ? AND real_date < borrow_date", uid, bid).First(&borrows).Error
	return borrows, err
}

//NewBorrow 新增借阅记录
func NewBorrow(uid int, bid int) error {

	borrowTime := time.Now()
	returnTime := borrowTime.AddDate(0, 1, 0)
	realTime := time.Date(1000, 1, 1, 0, 0, 0, 0, time.Now().Location())

	borrow := model.Borrow{
		UID:        uid,
		BID:        bid,
		BorrowDate: borrowTime,
		ReturnDate: returnTime,
		RealDate:   realTime,
	}

	err := utils.Db.Create(&borrow).Error
	return err
}

//FinishBorrow 完成借阅（还书）
func FinishBorrow(brid int) error {
	borrow := &model.Borrow{}
	err := utils.Db.Where("brid = ?", brid).Find(&borrow).Error

	err = utils.Db.Model(&borrow).Where("brid = ?", borrow.BrID).Update("real_date", time.Now()).Error
	return err
}
