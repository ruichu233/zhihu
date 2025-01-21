package svc

import (
	"context"
	v9 "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"zhihu/app/like/internal/config"
	"zhihu/pkg/db"
	"zhihu/pkg/mq"
	"zhihu/pkg/mq/redis"
	"zhihu/pkg/rdb"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *v9.Client
	MQ     mq.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {

	_rdb := rdb.InitRedis(&c.RDB)
	return &ServiceContext{
		Config: c,
		DB:     db.InitMysql(&c.DB),
		RDB:    _rdb,
		MQ:     redis.NewProducer(context.Background(), _rdb),
	}
}
