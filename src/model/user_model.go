package model

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
func (u *UserModel) Get(data map[string]string) {
	slaveDb.Table("user")
}


