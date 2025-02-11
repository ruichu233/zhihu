package svc

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"zhihu/app/comment/internal/config"
	"zhihu/app/comment/model"
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
	err := gormDb.AutoMigrate(&model.Comments{}, &model.CommentCount{})
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		DB:     gormDb,
		RDB:    rdb.InitRedis(&c.RDBConf),
	}
}
