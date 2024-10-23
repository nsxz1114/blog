package article

import (
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
	"github.com/nsxz1114/blog/service/search_ser"
	"go.uber.org/zap"
)

func (a Article) ArticleDetail(c *gin.Context) {
	var req models.ArticleSearchRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		res.FailWithCode(res.CodeInvalidParam, c)
		return
	}
	article, err := search_ser.GetDocumentById(req.ID)
	if err != nil {
		global.Log.Error("GetDocumentById err", zap.Error(err))
		res.FailWithMessage("找不到该文章", c)
		return
	}
	res.OkWithData(article, c)
}
