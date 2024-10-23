package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/api/system"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
	"github.com/nsxz1114/blog/utils"
	"go.uber.org/zap"
)

type UserLoginRequest struct {
	Account   string `json:"account"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captcha_id"`
}

func (u User) UserLogin(c *gin.Context) {
	var req UserLoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.CodeInvalidParam, c)
		return
	}
	if req.Captcha != "" && req.CaptchaId != "" && system.Store.Verify(req.CaptchaId, req.Captcha, true) {
		var user models.UserModel
		err = global.DB.Take(&user, "account=?", req.Account).Error
		if err != nil {
			global.Log.Error("Take err", zap.Error(err))
			res.FailWithMessage("用户名或密码错误", c)
			return
		}
		check := utils.CheckPassword(user.Password, req.Password)
		if !check {
			res.FailWithMessage("用户名或密码错误", c)
			return
		}
		aToken, rToken, err := utils.GenToken2(utils.PayLoad{
			Account: req.Account,
			UserID:  user.ID,
			Role:    int(user.Role)})
		if err != nil {
			global.Log.Error("token生成失败", zap.Error(err))
			res.FailWithMessage("登录失败", c)
			return
		}
		// 设置 Refresh Token 到 Cookie
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "refresh_token",
			Value:    rToken,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
		})
		res.OkWithData(aToken, c)
		return
	}
	res.FailWithMessage("验证码错误", c)
}
