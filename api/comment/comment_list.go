package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
)

type CommentListRequest struct {
	ID string `form:"id" uri:"id" json:"id"`
}

func (comment Comment) CommentList(c *gin.Context) {
	var req CommentListRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		res.FailWithCode(res.CodeInvalidParam, c)
		return
	}
	commentList := models.FindArticleComment(req.ID)
	data := filter.Select("c", commentList)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list := make([]models.CommentModel, 0)
		res.OkWithList(list, 0, c)
		return
	}
	res.OkWithList(data, len(commentList), c)
}
