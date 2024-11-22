package image

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"io/fs"
	"mime/multipart"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (i *Image) ImageUpload(c *gin.Context) {
	// 1. 获取上传文件
	form, err := c.MultipartForm()
	if err != nil {
		global.Log.Error("获取MultipartForm失败", zap.Error(err))
		res.Fail(c, res.CodeInternalError)
		return
	}

	fileList, ok := form.File["images"]
	if !ok || len(fileList) == 0 {
		res.Fail(c, res.CodeValidationFail)
		return
	}

	// 2. 确保上传目录存在
	if err := ensureUploadDir(global.Config.Upload.Path); err != nil {
		global.Log.Error("创建上传目录失败", zap.Error(err))
		res.Fail(c, res.CodeInternalError)
		return
	}

	// 3. 并发处理文件上传
	var (
		wg      sync.WaitGroup
		resList []models.UploadResponse
		mutex   sync.Mutex
	)

	for _, file := range fileList {
		wg.Add(1)
		go func(file *multipart.FileHeader) {
			defer wg.Done()

			// 处理单个文件上传
			serviceRes := processFileUpload(c, file)

			mutex.Lock()
			resList = append(resList, serviceRes)
			mutex.Unlock()
		}(file)
	}
	wg.Wait()

	res.Success(c, resList)
}

// 确保上传目录存在
func ensureUploadDir(path string) error {
	if _, err := os.ReadDir(path); err != nil {
		return os.MkdirAll(path, fs.ModePerm)
	}
	return nil
}

// 处理单个文件上传
func processFileUpload(c *gin.Context, file *multipart.FileHeader) models.UploadResponse {
	serviceRes := (&models.ImageModel{}).Upload(file)
	if !serviceRes.IsSuccess {
		return serviceRes
	}

	if err := c.SaveUploadedFile(file, serviceRes.FileName); err != nil {
		global.Log.Error("保存上传文件失败",
			zap.String("filename", file.Filename),
			zap.Error(err))
		return models.UploadResponse{
			IsSuccess: false,
			Msg:       "文件保存失败",
		}
	}

	return serviceRes
}