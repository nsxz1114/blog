package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nsxz1114/blog/global"
)

type PayLoad struct {
	Account string `json:"Account"` // 账号
	Role    int    `json:"role"`    // 权限
	UserID  uint   `json:"user_id"` // 用户id
}

var MySecret []byte

type CustomClaims struct {
	PayLoad
	jwt.StandardClaims
}

type MyClaims struct {
	ID       int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// func GetToken(payload PayLoad) (string, error) {
// 	MySecret = []byte(global.Config.Jwt.Secret)
// 	claim := CustomClaims{
// 		payload,
// 		jwt.StandardClaims{
// 			ExpiresAt: jwt.At(time.Now().Add(time.Hour * 24 * time.Duration(global.Config.Jwt.Expires))), //过期时间
// 			Issuer:    global.Config.Jwt.Issuer,
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
// 	return token.SignedString(MySecret)
// }

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return global.Config.Jwt.Issuer, nil
	})
	if err != nil {
		return nil, err
	}

	if token.Valid { // 校验token是否有效
		return mc, nil
	}

	return nil, errors.New("invalid token")
}

// GenToken2 生成access token和refresh token
func GenToken2(payload PayLoad) (aToken, rToken string, err error) {
	// 生成Access Token
	claims := CustomClaims{
		PayLoad: payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Duration(time.Minute * 15)).Unix(), // 过期时间
			Issuer: global.Config.Jwt.Issuer, // 签发人
		},
	}

	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(global.Config.Jwt.Secret)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate access token: %w", err)
	}

	// 生成Refresh Token
	rClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(global.Config.Jwt.Expires) * time.Hour).Unix(), // 过期时间
		Issuer:    global.Config.Jwt.Issuer,                                                    // 签发人
	}

	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, rClaims).SignedString(global.Config.Jwt.Secret)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return aToken, rToken, nil
}

// RefreshToken 刷新AccessToken
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	// 验证刷新令牌
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return global.Config.Jwt.Secret, nil
	}

	// 解析并验证刷新令牌
	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
		return "", "", fmt.Errorf("invalid refresh token: %w", err)
	}

	// 定义存储解析出的声明信息的变量
	var claims CustomClaims

	// 解析访问令牌
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	if err != nil {
		// 如果访问令牌过期，生成新的访问令牌和刷新令牌
		if v, ok := err.(*jwt.ValidationError); ok && v.Errors&jwt.ValidationErrorExpired != 0 {
			newAToken, newRToken, err = GenToken2(claims.PayLoad)
			if err != nil {
				return "", "", fmt.Errorf("failed to generate new tokens: %w", err)
			}
			return newAToken, newRToken, nil
		}
		// 处理其他访问令牌解析错误
		return "", "", fmt.Errorf("invalid access token: %w", err)
	}

	// 如果访问令牌仍然有效，直接返回空值，表示不需要刷新
	return "", "", nil
}
