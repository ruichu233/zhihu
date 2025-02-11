package svc

import (
	"context"
	v9 "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"zhihu/app/like/internal/config"
	"zhihu/app/like/model"
	"zhihu/pkg/db"
	"zhihu/pkg/mq"
	"zhihu/pkg/mq/redis"
	"zhihu/pkg/rdb"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *v9.Client
	MQP    mq.Producer
	MQC    mq.Consumer
}

func NewServiceContext(c config.Config) *ServiceContext {
	_db := db.InitMysql(&c.DB)
	if err := _db.AutoMigrate(&model.LikeRecord{}); err != nil {
		panic(err)
	}
	_rdb := rdb.InitRedis(&c.RDB)
	return &ServiceContext{
		Config: c,
		DB:     _db,
		RDB:    _rdb,
		MQP:    redis.NewProducer(context.Background(), _rdb),
		MQC:    redis.NewConsumer(context.Background(), _rdb, "like", "1", "1"),
	}
}
