package svc

import (
	"context"

	"github.com/redis/go-redis/v9"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"zhihu/app/notification/internal/config"
	"zhihu/app/notification/model"
	"zhihu/pkg/db"
	"zhihu/pkg/mq"
	_redis "zhihu/pkg/mq/redis"
	"zhihu/pkg/rdb"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	RDB      *redis.Client
	Consumer mq.Consumer
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormDb := db.InitMysql(&c.DBConf)
	if err := gormDb.AutoMigrate(&model.Notification{}); err != nil {
		panic(err)
	}
	_rdb := rdb.InitRedis(&c.RDBConf)
	return &ServiceContext{
		Config:   c,
		DB:       gormDb,
		RDB:      _rdb,
		Consumer: _redis.NewConsumer(context.Background(), _rdb, "notify_topic", "notify_topic_group_1", "notify_topic_group_1_consumer_1"),
	}
}
