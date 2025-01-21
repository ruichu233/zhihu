package svc

import (
	client "github.com/gorse-io/gorse-go"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"zhihu/app/video/internal/config"
	"zhihu/app/video/internal/model"
	"zhihu/pkg/db"
	"zhihu/pkg/oss"
	"zhihu/pkg/rdb"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
	OSS    *minio.Client
	Gorse  *client.GorseClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormDb := db.InitMysql(&c.DBConf)
	if err := gormDb.AutoMigrate(&model.Video{}); err != nil {
		panic(err)
	}
	gorse := client.NewGorseClient("http://127.0.0.1:8088", "api_key")

	return &ServiceContext{
		Config: c,
		DB:     db.InitMysql(&c.DBConf),
		RDB:    rdb.InitRedis(&c.RDBConf),
		OSS:    oss.InitMinio(&c.OSSConf),
		Gorse:  gorse,
	}
}
