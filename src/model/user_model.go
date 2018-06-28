package model

import (
	"common/orm"
	"common/sqlbuilder"
)

type UserModel struct {
	baseModel
}

type User struct {
	Id      int64  `json:"id,omitempty"`
	LoginName    string `json:"login_name,omitempty"`
	Pwd       string `json:"pwd,omitempty"`
}


//用户表
const UserTable = "users"

//新建
func (u *UserModel) Add(data map[string]string) (int64, error) {
	cols, values := u.initMapData(data)
	sqlStr := (&sqlbuilder.SqlBuilder{}).Insert(UserTable, cols).GetSql()
	db := orm.NewOrm()
	stmt, err := db.Raw(sqlStr).Prepare() //stmt要关闭
	res, err := stmt.Exec(*values...)
	stmt.Close()
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

//通过loginName获取用户信息
func (u *UserModel) GetByLoginName(loginName string, cols string) (*User, error) {
	sqlStr := (&sqlbuilder.SqlBuilder{}).
		Select(cols, UserTable).
		Where("login_name", "=").
		GetSql()
	var user = User{}
	db := orm.NewOrm()
	db.Using("slave")
	err := db.Raw(sqlStr, loginName).QueryRow(&user)
	//这里把noRows当做没有错误, 只是没有返回数据
	if err != nil && err != orm.ErrNoRows{
		return &user, err
	}
	return &user, nil
}

//通过登录名删除用户
func (u *UserModel)DeleteUserByLoginName(loginName string) (int64, error) {
	sqlStr := (&sqlbuilder.SqlBuilder{}).
		Delete(UserTable).
		Where("login_name", "=").
		GetSql()
	db := orm.NewOrm()
	stmt, err := db.Raw(sqlStr).Prepare()
	defer stmt.Close()
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(loginName)
	return res.RowsAffected()
}

