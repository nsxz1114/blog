package article

import (
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
	"go.uber.org/zap"
)

func (a Article) ArticleDelete(c *gin.Context) {
	var req models.ArticleSearchRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		res.FailWithCode(res.CodeInvalidParam, c)
		return
	}
	err = models.DeleteDoc(req.ID)
	if err != nil {
		global.Log.Error("DeleteDoc err", zap.Error(err))
		res.FailWithMessage("文章删除失败", c)
		return
	}
	res.Ok(c)
}
