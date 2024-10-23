package router

import "github.com/nsxz1114/blog/api"

func (router RouterGroup) CaptchaRouter() {
	SystemApi := api.AppGroupApp.SystemApi
	router.GET("captcha", SystemApi.CaptchaCreate)
	router.GET("refreshtoken", SystemApi.RefreshToken)
}
