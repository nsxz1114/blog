package article

import (
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
	"github.com/nsxz1114/blog/service/search_ser"
)

type ArticleListRequest struct {
	models.PageInfo
}

func (a Article) ArticleList(c *gin.Context) {
	var req ArticleListRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		res.FailWithCode(res.CodeInvalidParam, c)
		return
	}
	articles := search_ser.SearchAllDocuments(req.PageInfo)
	res.OkWithList(articles, len(articles), c)
}
