package image_ser

import (
	"fmt"
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models"
	"github.com/nsxz1114/blog/models/ctype"
	"github.com/nsxz1114/blog/utils"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

var WhiteList = []string{
	"jpg",
	"png",
	"jpeg",
	"ico",
	"tiff",
	"gif",
	"svg",
	"webp",
}

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

// ImageUploadService 文件上传
func ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fileName := file.Filename
	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, file.Filename)
	res.FileName = filePath
	fileType := ctype.Local
	filePath = "/" + filePath
	// 文件白名单判断
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, WhiteList) {
		res.Msg = "文件格式错误"
		return
	}

	// 判断文件大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小超过设定,当前大小为:%.2fMB,设定大小为:%dMB", size, global.Config.Upload.Size)
		return
	}

	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error("open fileObj err", zap.Error(err))
		return
	}

	byteData, err := io.ReadAll(fileObj)
	if err != nil {
		global.Log.Error("read file err", zap.Error(err))
		return
	}

	imageHash := utils.Md5(byteData)
	var image models.ImageModel
	err = global.DB.Where("hash = ?", imageHash).First(&image).Error
	if err == nil {
		res.Msg = "图片已存在"
		res.FileName = image.Path
		return
	}
	res.Msg = "图片上传成功"
	res.IsSuccess = true
	global.DB.Create(&models.ImageModel{
		Hash:      imageHash,
		Path:      filePath,
		Name:      fileName,
		ImageType: fileType,
	})
	return
}
