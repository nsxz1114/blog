package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models/res"
	"github.com/nsxz1114/blog/utils"
	"go.uber.org/zap"
)

func (s System) RefreshToken(c *gin.Context) {
	// 从 Cookie 中获取 Refresh Token
	cookie, err := c.Request.Cookie("refresh_token")
	if err != nil {
		global.Log.Error("refresh_token err", zap.Error(err))
		res.FailWithCode(res.CodeNeedLogin, c)
		return
	}

	// 使用 Refresh Token 获取新的 Access Token
	newAccessToken, newRefreshToken, err := utils.RefreshToken("", cookie.Value)
	if err != nil {
		global.Log.Error("refresh_token err", zap.Error(err))
		res.FailWithCode(res.CodeInvalidToken, c)
		return
	}

	// 如果生成了新的 Refresh Token，更新 Cookie
	if newRefreshToken != "" {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "refresh_token",
			Value:    newRefreshToken,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
		})
	}

	// 返回新的 Access Token
	res.OkWithData(newAccessToken, c)

}
