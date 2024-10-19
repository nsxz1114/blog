package router

import (
	"github.com/nsxz1114/blog/api"
	"github.com/nsxz1114/blog/middleware"
)

func (router RouterGroup) CommentRouter() {
	commentRouter := router.Group("comment")
	commentApi := api.AppGroupApp.CommentApi
	commentRouter.POST("create", middleware.JwtAdmin(), commentApi.CommentCreate)
	commentRouter.DELETE("delete/:id", middleware.JwtAdmin(), commentApi.CommentDelete)
	commentRouter.GET("list/:id", commentApi.CommentList)
}
