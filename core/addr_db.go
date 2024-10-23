package core

import (
	"github.com/cc14514/go-geoip2"
	geoip2db "github.com/cc14514/go-geoip2-db"
	"github.com/nsxz1114/blog/global"
	"go.uber.org/zap"
)

func InitAddrDB() *geoip2.DBReader {
	db, err := geoip2db.NewGeoipDbByStatik()
	if err != nil {
		global.Log.Fatal("InitAddrDB err", zap.Error(err))
	}
	return db
}
