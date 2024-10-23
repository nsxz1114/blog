package api

import (
	"github.com/nsxz1114/blog/api/article"
	"github.com/nsxz1114/blog/api/comment"
	"github.com/nsxz1114/blog/api/image"
	"github.com/nsxz1114/blog/api/message"
	"github.com/nsxz1114/blog/api/system"
	"github.com/nsxz1114/blog/api/user"
)

type AppGroup struct {
	SystemApi  system.System
	UserApi    user.User
	ImageApi   image.Image
	ArticleApi article.Article
	CommentApi comment.Comment
	MessageApi message.Message
}

var AppGroupApp = new(AppGroup)
