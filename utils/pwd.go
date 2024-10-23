package utils

import (
	"github.com/nsxz1114/blog/global"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword  hash密码
func HashPassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		global.Log.Error("HashPassword err", err)
		return ""
	}
	return string(hash)
}

// CheckPassword  验证密码
func CheckPassword(hashPwd string, pwd string) bool {
	byteHash := []byte(hashPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(pwd))
	if err != nil {
		global.Log.Error("CheckPassword err", err)
		return false
	}
	return true
}
