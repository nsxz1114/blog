package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	ctype "github.com/nsxz1114/blog/models/ctype"
	"github.com/nsxz1114/blog/models/res"
	"github.com/nsxz1114/blog/utils"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization Header 中的 Access Token
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			res.FailWithMessage("未登录", c) // 未登录
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			res.FailWithMessage("无效的token", c)
			c.Abort()
			return
		}

		accessToken := parts[1]

		// 解析 Access Token
		mc, err := utils.ParseToken(accessToken)
		if err != nil {
			// 如果是 Access Token 过期，尝试使用 Refresh Token 刷新
			if v, ok := err.(*jwt.ValidationError); ok && v.Errors == jwt.ValidationErrorExpired {
				// 从 Cookie 中获取 Refresh Token
				rTokenCookie, err := c.Request.Cookie("refresh_token")
				if err != nil {
					res.FailWithMessage("未登录", c) // 缺少 Refresh Token
					c.Abort()
					return
				}

				// 使用 Refresh Token 刷新 Access Token
				newAToken, newRToken, err := utils.RefreshToken(accessToken, rTokenCookie.Value)
				if err != nil {
					res.FailWithMessage("无效的token", c)
					c.Abort()
					return
				}

				// 刷新成功，设置新的 Refresh Token
				if newRToken != "" {
					http.SetCookie(c.Writer, &http.Cookie{
						Name:     "refresh_token",
						Value:    newRToken,
						HttpOnly: true,
						Secure:   true,
						SameSite: http.SameSiteStrictMode,
						Path:     "/",
					})
				}

				// 返回新的 Access Token 给前端
				res.OkWithData(newAToken, c)
				c.Abort() // 阻止后续操作，因为已经返回新 token
				return
			}

			// 其他 Token 解析错误
			res.FailWithMessage("无效的token", c)
			c.Abort()
			return
		}

		// Access Token 有效，保存用户信息到上下文
		c.Set("claims", mc)

		// 继续处理下一个中间件或请求处理函数
		c.Next()
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			res.FailWithMessage("未登录", c)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			res.FailWithMessage("无效的token", c)
			c.Abort()
			return
		}
		accessToken := parts[1]
		mc, err := utils.ParseToken(accessToken)
		if err != nil {
			if v, ok := err.(*jwt.ValidationError); ok && v.Errors == jwt.ValidationErrorExpired {
				rTokenCookie, err := c.Request.Cookie("refresh_token")
				if err != nil {
					res.FailWithMessage("未登录", c) // 缺少 Refresh Token
					c.Abort()
					return
				}
				newAToken, newRToken, err := utils.RefreshToken(accessToken, rTokenCookie.Value)
				if err != nil {
					res.FailWithMessage("无效的token", c)
					c.Abort()
					return
				}
				if newRToken != "" {
					http.SetCookie(c.Writer, &http.Cookie{
						Name:     "refresh_token",
						Value:    newRToken,
						HttpOnly: true,
						Secure:   true,
						SameSite: http.SameSiteStrictMode,
						Path:     "/",
					})
				}
				res.OkWithData(newAToken, c)
				c.Abort()
				return
			}
			res.FailWithMessage("无效的token", c)
			c.Abort()
			return
		}
		if mc.PayLoad.Role != int(ctype.PermissionAdmin) {
			res.FailWithMessage("权限错误", c)
			c.Abort()
			return
		}
		c.Set("claims", mc)
		c.Next()
	}
}
