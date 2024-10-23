package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
	"github.com/nsxz1114/blog/service/search_ser"
	"github.com/nsxz1114/blog/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CommentCreateRequest struct {
	ArticleID       string `json:"article_id" binding:"required" msg:"请选择文章"`
	Content         string `json:"content" binding:"required" msg:"请输入评论内容"`
	ParentCommentID *uint  `json:"parent_comment_id"`
}

func (comment Comment) CommentCreate(c *gin.Context) {
	var req CommentCreateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.CodeInvalidParam, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)
	exist := search_ser.DocIsExistById(req.ArticleID)
	if exist == false {
		res.FailWithMessage("文章不存在", c)
		return
	}

	if req.ParentCommentID != nil {
		var parentComment models.CommentModel
		err = global.DB.Take(&parentComment, req.ParentCommentID).Error
		if err != nil {
			global.Log.Error("Take err", zap.Error(err))
			res.FailWithMessage("评论失败", c)
			return
		}
		global.DB.Model(&parentComment).Update("comment_count", gorm.Expr("comment_count + 1"))
	}
	err = global.DB.Create(&models.CommentModel{
		ParentCommentID: req.ParentCommentID,
		Content:         req.Content,
		ArticleID:       req.ArticleID,
		UserID:          claims.UserID,
	}).Error
	if err != nil {
		global.Log.Error("Create err", zap.Error(err))
		res.FailWithMessage("评论失败", c)
		return
	}
	res.Ok(c)
}
