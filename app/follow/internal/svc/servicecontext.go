package svc

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"zhihu/app/follow/internal/config"
	"zhihu/app/follow/model"
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
	if err := gormDb.AutoMigrate(&model.Follow{}); err != nil {
		panic(err)
	}
	if err := gormDb.AutoMigrate(&model.FollowsCount{}); err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		DB:     db.InitMysql(&c.DBConf),
		RDB:    rdb.InitRedis(&c.RDBConf),
	}
}
