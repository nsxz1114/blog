package models

import "github.com/nsxz1114/blog/models/ctype"

type UserModel struct {
	MODEL    `json:",select(info)"`
	Nickname string     `json:"nick_name,select(c|info)" gorm:"comment:昵称"`
	Account  string     `json:"account" gorm:"comment:账号"`
	Password string     `json:"password" gorm:"comment:密码"`
	Avatar   string     `json:"avatar,select(c|info)" gorm:"comment:头像"`
	Email    string     `json:"email,select(info)" gorm:"comment:邮箱"`
	Address  string     `json:"address,select(c|info)" gorm:"comment:地址"`
	Token    string     `json:"token"`
	Role     ctype.Role `json:"role,string,select(info)" gorm:"comment:身份,1管理员,2用户"`
}
