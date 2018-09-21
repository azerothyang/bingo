package model

import (
	"log"
)

type UserModel struct {
	baseModel
}

type User struct {
	UserId       int64
	UserName     string
	UserPassword string
	CreateTime   string
}

//用户表
const UserTable = "t_user"

//获取用户列表
func (u *UserModel) GetList() *[]User {
	var users []User
	rows, err := slaveDb.Table(UserTable).Select("*").Rows()
	defer rows.Close() //使用defer优化语法，但是也会降低性能
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var user User
		slaveDb.ScanRows(rows, &user)
		users = append(users, user)
	}
	return &users
}
