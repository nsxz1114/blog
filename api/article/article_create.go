package article

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/res"
	"github.com/nsxz1114/blog/service/search_ser"
	"github.com/nsxz1114/blog/utils"
)

type ArticleRequest struct {
	Title    string `json:"title"`
	Abstract string `json:"abstract"`
	Category string `json:"category"`
	Content  string `json:"content" `
	CoverID  uint   `json:"cover_id"`
}

func (a Article) ArticleCreate(c *gin.Context) {
	var req ArticleRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(res.CodeInvalidParam, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*utils.CustomClaims)
	userId := claims.UserID
	html, err := utils.ConvertMarkdownToHTML(req.Content)
	if err != nil {
		res.FailWithMessage("文章发布失败", c)
		return
	}
	content, err := utils.ConvertHTMLToMarkdown(html)
	if err != nil {
		res.FailWithMessage("文章发布失败", c)
		return
	}

	if req.CoverID == 0 {
		var imageIDList []uint
		global.DB.Model(models.ImageModel{}).Select("id").Scan(&imageIDList)
		if len(imageIDList) == 0 {
			res.FailWithMessage("暂无图片", c)
			return
		}
		rand.New(rand.NewSource(time.Now().UnixNano()))
		req.CoverID = imageIDList[rand.Intn(len(imageIDList))]
	}

	var coverUrl string
	err = global.DB.Model(models.ImageModel{}).Where("id = ?", req.CoverID).Select("path").Scan(&coverUrl).Error
	if err != nil {
		res.FailWithMessage("文章发布失败", c)
		return
	}
	var user models.UserModel
	err = global.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		res.FailWithMessage("文章发布失败", c)
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	id := strconv.FormatInt(utils.GenerateID(), 10)
	article := models.Article{
		ID:         id,
		Title:      req.Title,
		Abstract:   req.Abstract,
		Category:   req.Category,
		Content:    content,
		CoverID:    req.CoverID,
		CoverURL:   coverUrl,
		UserID:     userId,
		CreatedAt:  now,
		UpdatedAt:  now,
		UserName:   user.Nickname,
		UserAvatar: user.Avatar,
	}
	exist := search_ser.DocIsExistByTitle(req.Title)
	if exist {
		res.FailWithMessage("文章已存在", c)
		return
	}
	err = article.CreateDoc()
	if err != nil {
		res.FailWithMessage("文章发布失败", c)
		return
	}
	res.OkWithMessage("文章发布成功", c)
}
