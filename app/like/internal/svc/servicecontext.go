package svc

import (
	"context"
	"zhihu/app/like/internal/config"
	"zhihu/app/like/model"
	"zhihu/pkg/db"
	"zhihu/pkg/mq"
	_redis "zhihu/pkg/mq/redis"
	"zhihu/pkg/rdb"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
	MQP    mq.Producer
	MQC    mq.Consumer
}

func NewServiceContext(c config.Config) *ServiceContext {
	_db := db.InitMysql(&c.DBConf)
	if err := _db.AutoMigrate(&model.LikeRecord{}); err != nil {
		panic(err)
	}
	_rdb := rdb.InitRedis(&c.RDBConf)
	return &ServiceContext{
		Config: c,
		DB:     _db,
		RDB:    _rdb,
		MQP:    _redis.NewProducer(context.Background(), _rdb),
		MQC:    _redis.NewConsumer(context.Background(), _rdb, "like", "1", "1"),
	}
}
