package message

import (
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
	"go.uber.org/zap"
)

func (m Message) MessageDelete(c *gin.Context) {
	var req models.MessageSearchRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		res.FailWithCode(res.CodeInvalidParam, c)
		return
	}
	var message models.MessageModel
	err = global.DB.Where("id=?", req.ID).First(&message).Error
	if err != nil {
		global.Log.Error("id err", zap.Error(err))
		res.FailWithMessage("删除失败", c)
		return
	}
	global.DB.Model(&message).Delete("id=?", req.ID)
	res.Ok(c)
}
