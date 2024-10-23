package message

import (
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
	"github.com/nsxz1114/blog/utils"
	"go.uber.org/zap"
)

type MessageCreateRequest struct {
	RevUserID uint   `json:"rec_user_id" binding:"required"`
	Content   string `json:"content" binding:"required"`
}

func (m Message) MessageCreate(c *gin.Context) {
	var req MessageCreateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.CodeInvalidParam, c)
		return
	}
	var sendUser, recUser models.UserModel

	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)
	err = global.DB.Take(&sendUser, claims.UserID).Error
	if err != nil {
		global.Log.Error("Take err", zap.Error(err))
		res.FailWithMessage("消息发送失败", c)
		return
	}
	err = global.DB.Take(&recUser, req.RevUserID).Error
	if err != nil {
		global.Log.Error("Take err", zap.Error(err))
		res.FailWithMessage("消息发送失败", c)
		return
	}
	err = global.DB.Create(&models.MessageModel{
		SendUserID:       sendUser.ID,
		SendUserNickName: sendUser.Nickname,
		SendUserAvatar:   sendUser.Avatar,
		RecUserID:        req.RevUserID,
		RecUserNickName:  recUser.Nickname,
		RecUserAvatar:    recUser.Avatar,
		IsRead:           false,
		Content:          req.Content,
	}).Error
	if err != nil {
		global.Log.Error("Take err", zap.Error(err))
		res.FailWithMessage("消息发送失败", c)
		return
	}
	res.Ok(c)
}
