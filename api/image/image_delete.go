package image

import (
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
	"go.uber.org/zap"
)

func (i Image) ImageDelete(c *gin.Context) {
	var req models.RemoveRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.CodeInvalidParam, c)
		return
	}
	var imageList []models.ImageModel
	global.DB.Find(&imageList, req.IDList)
	err = global.DB.Delete(&imageList).Error
	if err != nil {
		global.Log.Error("Delete err", zap.Error(err))
		res.FailWithMessage("图片删除失败", c)
		return
	}
	res.Ok(c)

}
