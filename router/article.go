package router

import (
	"github.com/nsxz1114/blog/api"
	"github.com/nsxz1114/blog/middleware"
)

func (router RouterGroup) ArticleRouter() {
	articleApi := api.AppGroupApp.ArticleApi
	articleRouter := router.Group("article")
	router.GET("article/:id", articleApi.ArticleDetail)
	articleRouter.POST("create", middleware.JwtAdmin(), articleApi.ArticleCreate)
	articleRouter.GET("list", articleApi.ArticleList)
	articleRouter.DELETE("delete/:id", middleware.JwtAdmin(), articleApi.ArticleDelete)
	articleRouter.PUT("update", middleware.JwtAdmin(), articleApi.ArticleUpdate)
	articleRouter.GET("search", articleApi.ArticleSearch)
}
