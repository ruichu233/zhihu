package svc

import (
	"context"
	"zhihu/app/comment/internal/config"
	"zhihu/app/comment/model"
	"zhihu/pkg/db"
	"zhihu/pkg/mq"
	_redis "zhihu/pkg/mq/redis"
	"zhihu/pkg/rdb"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	RDB      *redis.Client
	Producer mq.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormDb := db.InitMysql(&c.DBConf)
	err := gormDb.AutoMigrate(&model.Comments{}, &model.CommentCount{})
	if err != nil {
		panic(err)
	}
	_redisC := rdb.InitRedis(&c.RDBConf)
	return &ServiceContext{
		Config:   c,
		DB:       gormDb,
		RDB:      _redisC,
		Producer: _redis.NewProducer(context.Background(), _redisC),
	}
}
