package router

import (
	"github.com/nsxz1114/blog/api"
	"github.com/nsxz1114/blog/middleware"
)

func (router RouterGroup) UserRouter() {
	userApi := api.AppGroupApp.UserApi
	userRouter := router.Group("user")
	userRouter.GET("info", middleware.JwtAuth(), userApi.Userinfo)
	userRouter.POST("create", userApi.UserCreate)
	userRouter.POST("login", userApi.UserLogin)
	userRouter.POST("logout", middleware.JwtAuth(), userApi.UserLogout)
	userRouter.POST("update", middleware.JwtAuth(), userApi.UserInfoUpdate)
}
