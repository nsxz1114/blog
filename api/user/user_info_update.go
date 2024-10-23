package user

import (
	"strings"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
	"github.com/nsxz1114/blog/utils"
	"go.uber.org/zap"
)

type UserInfoUpdateRequest struct {
	NickName string `json:"nick_name" structs:"nick_name"`
	Avatar   string `json:"avatar" structs:"avatar"`
	Email    string `json:"email"  structs:"email"`
}

func (u User) UserInfoUpdate(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)
	var req UserInfoUpdateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.CodeInvalidParam, c)
		return
	}
	var newMap = map[string]interface{}{}
	maps := structs.Map(req)
	for key, v := range maps {
		if val, ok := v.(string); ok && strings.TrimSpace(val) != "" {
			newMap[key] = val
		}
	}
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		global.Log.Error("Take err", zap.Error(err))
		res.FailWithMessage("更新失败", c)
		return
	}
	err = global.DB.Model(&user).Updates(newMap).Error
	if err != nil {
		global.Log.Error("Updates err", zap.Error(err))
		res.FailWithMessage("更新失败", c)
		return
	}
	res.Ok(c)
}
