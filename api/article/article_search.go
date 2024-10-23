package article

import (
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/models/res"
	"github.com/nsxz1114/blog/service/search_ser"
)

type ArticleSearchRequest struct {
	Key string `json:"key" form:"key"`
}

func (a Article) ArticleSearch(c *gin.Context) {
	var req ArticleSearchRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		res.FailWithCode(res.CodeInvalidParam, c)
		return
	}
	fields := []string{"title", "id"}
	articles := search_ser.SearchDocumentMultiMatchByTitle(fields, req.Key)
	res.OkWithList(articles, len(articles), c)
}
