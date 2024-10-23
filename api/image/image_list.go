package image

import (
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
	"github.com/nsxz1114/blog/service/search_ser"
	"go.uber.org/zap"
)

func (i Image) ImageList(c *gin.Context) {
	list, count, err := search_ser.SqlSearch(models.ImageModel{}, search_ser.Option{})
	if err != nil {
		global.Log.Error("SqlSearch err", zap.Error(err))
		res.FailWithMessage("图片加载失败", c)
		return
	}
	res.OkWithList(list, int(count), c)
}
