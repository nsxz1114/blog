package models

import (
	"github.com/nsxz1114/blog/global"
	"github.com/nsxz1114/blog/models/ctype"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

type ImageModel struct {
	MODEL
	Path      string          `gorm:"comment:图片路径" json:"path"`
	Hash      string          `gorm:"comment:图片的hash值" json:"hash"`
	Name      string          `gorm:"comment:图片名称" json:"name"`
	ImageType ctype.ImageType `gorm:"default:1;comment:图片的类型，1本地，2云端，默认是1" json:"image_type,string"`
}

func (b *ImageModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		err = os.Remove(b.Path[1:])
		if err != nil {
			global.Log.Error("delete local file failed", zap.Error(err))
			return nil
		}
	}
	return nil
}
