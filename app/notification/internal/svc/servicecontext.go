package svc

import (
	"github.com/redis/go-redis/v9"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"zhihu/app/notification/internal/config"
	"zhihu/app/notification/model"
	"zhihu/pkg/db"
	"zhihu/pkg/rdb"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormDb := db.InitMysql(&c.DBConf)
	if err := gormDb.AutoMigrate(&model.Notification{}); err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		DB:     gormDb,
		RDB:    rdb.InitRedis(&c.RDBConf),
	}
}
