package svc

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"zhihu/app/user/internal/config"
	"zhihu/app/user/model"
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
	err := gormDb.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		DB:     db.InitMysql(&c.DBConf),
		RDB:    rdb.InitRedis(&c.RDBConf),
	}
}
