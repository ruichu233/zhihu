package svc

import (
	"context"
	"zhihu/app/follow/internal/config"
	"zhihu/app/follow/model"
	"zhihu/pkg/db"
	"zhihu/pkg/mq"
	_kafka "zhihu/pkg/mq/kafka"
	"zhihu/pkg/rdb"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
	Prod   mq.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormDb := db.InitMysql(&c.DBConf)
	if err := gormDb.AutoMigrate(&model.Follow{}); err != nil {
		panic(err)
	}
	if err := gormDb.AutoMigrate(&model.FollowsCount{}); err != nil {
		panic(err)
	}
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"127.0.0.1:9092"},
		Topic:   "user_follow",
	})

	// kafkaConsumer := kafka.New(kafka.ReaderConfig{ // 创建消费者
	// 	Brokers:  []string{"127.0.0.1:9092"},
	// 	Topic:    "user_follow",
	// 	GroupID:  "",
	// 	MinBytes: 10e3, // 10KB
	// 	MaxBytes: 10e6, // 10MB
	// 	MaxWait:  time.Second * 1,
	// })
	return &ServiceContext{
		Config: c,
		DB:     db.InitMysql(&c.DBConf),
		RDB:    rdb.InitRedis(&c.RDBConf),
		Prod:   _kafka.NewProducer(context.Background(), writer),
	}
}
