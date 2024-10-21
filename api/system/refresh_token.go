package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/utils"
)

func (s System) RefreshToken(c *gin.Context) {
	// 从 Cookie 中获取 Refresh Token
	cookie, err := c.Request.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing refresh token"})
		return
	}

	// 使用 Refresh Token 获取新的 Access Token
	newAccessToken, newRefreshToken, err := utils.RefreshToken("", cookie.Value)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
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
	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}
