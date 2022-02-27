package dao

import (
	"nbook/model"
	"nbook/utils"
)

//CheckUserNameAndPassword 根据用户名和密码从数据库中查询一条记录
func CheckUserNameAndPassword(username string, password string) (*model.Users, error) {

	users := &model.Users{}
	//sql: select uid,uname,phong_number,password from users where uname = username and upassword = password;
	err := utils.Db.Where("uname = ? AND upassword >= ?", username, password).Find(&users).Error
	return users, err
}

//CheckUserName 根据用户名从数据库中查询一条记录
func CheckUserName(username string) (*model.Users, error) {
	users := &model.Users{}
	//sql: select uid,uname,phong_number,password from users where uname = username;
	err := utils.Db.Where("uname = ?", username).Find(&users).Error
	return users, err
}

//SaveUser 向数据库中插入用户信息
func SaveUser(username string, phone_number string, password string) error {
	//sql语句: insert into users(uname,phone_number,upassword) values(username,phone_number,password)
	user := model.Users{
		UName:       username,
		PhoneNumber: phone_number,
		UPassword:   password,
	}

	//通过Error对象检测，插入数据有没有成功，如果没有错误那就是数据写入成功了。
	err := utils.Db.Create(&user).Error
	return err

}
