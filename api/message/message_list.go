package message

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
	"github.com/nsxz1114/blog/utils"
	"go.uber.org/zap"
)

func (m Message) MessageList(c *gin.Context) {
	var req models.MessageSearchRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		res.FailWithCode(res.CodeInvalidParam, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)
	fmt.Println(claims.UserID, req.ID)
	var messageList []models.MessageModel
	err = global.DB.Order("created_at asc").
		Find(&messageList, "(send_user_id = ? and rec_user_id = ?) or (rec_user_id = ? and send_user_id = ?)",
			claims.UserID, req.ID, claims.UserID, req.ID).Error
	if err != nil {
		global.Log.Error("Find err", zap.Error(err))
		res.FailWithError(err, &messageList, c)
		return
	}
	res.OkWithList(messageList, len(messageList), c)
}
