package svc

import (
	"zhihu/app/chat/internal/config"
	"zhihu/app/chat/model"
	"zhihu/pkg/db"
	"zhihu/pkg/rdb"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormDb := db.InitMysql(c.DBConf)
	// 自动迁移
	if err := gormDb.AutoMigrate(&model.Message{}); err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		DB:     gormDb,
		RDB:    rdb.InitRedis(c.RDBConf),
	}
}
