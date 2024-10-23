package user

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models/ctype"
	"github.com/nsxz1114/blog/models/res"
	"github.com/nsxz1114/blog/service/user_ser"
	"github.com/nsxz1114/blog/utils"
	"go.uber.org/zap"
)

type UserCreateRequest struct {
	Nickname string     `json:"nick_name" binding:"required" msg:"请输入昵称"`
	Password string     `json:"password" binding:"required" msg:"请输入密码"`
	Role     ctype.Role `json:"role" binding:"required" msg:"请选择身份"`
}

func (u User) UserCreate(c *gin.Context) {
	var req UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithCode(res.CodeInvalidParam, c)
		return
	}
	account := utils.GenerateID()
	err := user_ser.CreateUser(req.Nickname, strconv.FormatInt(account, 10), req.Password, c.ClientIP(), req.Role)
	if err != nil {
		global.Log.Error("CreateUser err", zap.Error(err))
		res.FailWithMessage("用户创建失败", c)
		return
	}
	res.Ok(c)
}
