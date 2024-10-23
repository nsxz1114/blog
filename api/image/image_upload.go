package image

import (
	"github.com/gin-gonic/gin"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models/res"
	"github.com/nsxz1114/blog/service/image_ser"
	"go.uber.org/zap"
	"io/fs"
	"os"
)

func (i Image) ImageUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		global.Log.Error("MultipartForm err", zap.Error(err))
		res.FailWithMessage("图片上传失败", c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("图片上传失败", c)
		return
	}

	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error("MkdirAll err", zap.Error(err))
			res.FailWithMessage("图片上传失败", c)
			return
		}
	}

	var resList []image_ser.FileUploadResponse

	for _, file := range fileList {
		serviceRes := image_ser.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		err = c.SaveUploadedFile(file, serviceRes.FileName)
		if err != nil {
			global.Log.Error("save file err", zap.Error(err))
			serviceRes.Msg = err.Error()
			serviceRes.IsSuccess = false
			resList = append(resList, serviceRes)
			continue
		}
		resList = append(resList, serviceRes)
	}
	res.OkWithData(resList, c)
}
