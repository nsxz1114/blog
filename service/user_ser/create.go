package user_ser

import (
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/ctype"
	"github.com/nsxz1114/blog/utils"
)

const Avatar = "/upload/avatar/default_avatar.jpg"

func CreateUser(nickname, account, password, ip string, role ctype.Role) (err error) {
	//判断用户是否存在
	var user models.UserModel
	err = global.DB.Where("nick_name=?", nickname).First(&user).Error
	if err == nil {
		global.Log.Errorf("昵称%s已存在", nickname)
		return err
	}
	//对密码进行加密
	hashpwd := utils.HashPassword(password)

	addr := utils.GetAddrByIp(ip)

	err = global.DB.Create(&models.UserModel{
		Avatar:   Avatar,
		Nickname: nickname,
		Account:  account,
		Password: hashpwd,
		Role:     role,
		Address:  addr,
	}).Error
	if err != nil {
		global.Log.Error("create user err", err)
		return err
	}
	return nil
}
