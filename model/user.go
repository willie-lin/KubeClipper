package model

import (
	"KubeClipper/initDB"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username	string	`json:"username"`
	Passwoed	string	`json:"passwoed"`
	Email		string	`json:"email"`
	Nickname 	string	`json:"nickname"`
	Address 	string	`json:"address"`
}


// 注册用户，插入数据库

func (user User) Insert()  bool {

	initDB.Db.Create(&user)

	if initDB.Db.Error == nil {
		return false
	}
	return false
}

// 查询用户名

func (user User) QueryUsername() User  {
	initDB.Db.First(&user, user.Username)
	return user
}

