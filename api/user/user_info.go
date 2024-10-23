package user

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
	"github.com/nsxz1114/blog/utils"
	"go.uber.org/zap"
)

func (u User) Userinfo(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)

	var user models.UserModel
	if err := global.DB.First(&user, claims.UserID).Error; err != nil {
		global.Log.Error("search err", zap.Error(err))
		res.FailWithMessage("获取用户信息失败", c)
		return
	}
	data := filter.Select("info", user)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		user := make([]models.UserModel, 0)
		res.OkWithData(user, c)
		return
	}
	res.OkWithData(data, c)
}
