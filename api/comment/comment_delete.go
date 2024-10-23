package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CommentIDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

func (comment Comment) CommentDelete(c *gin.Context) {
	var req CommentIDRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		res.FailWithCode(res.CodeInvalidParam, c)
		return
	}
	var commentModel models.CommentModel
	err = global.DB.Take(&commentModel, req.ID).Error
	if err != nil {
		global.Log.Error("Take err", zap.Error(err))
		res.FailWithMessage("删除失败", c)
		return
	}

	subCommentList := models.FindAllSubComment(commentModel)
	count := len(subCommentList) + 1

	if commentModel.ParentCommentID != nil {
		global.DB.Model(&models.CommentModel{}).
			Where("id = ?", *commentModel.ParentCommentID).
			Update("comment_count", gorm.Expr("comment_count - ?", count))
	}

	var deleteCommentIDList []uint
	for _, model := range subCommentList {
		deleteCommentIDList = append(deleteCommentIDList, model.ID)
	}
	deleteCommentIDList = append(deleteCommentIDList, commentModel.ID)
	for _, id := range deleteCommentIDList {
		global.DB.Model(models.CommentModel{}).Delete("id = ?", id)
	}
	res.Ok(c)
}
