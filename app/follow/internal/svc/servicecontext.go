package svc

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"zhihu/app/follow/internal/config"
	"zhihu/pkg/db"
	"zhihu/pkg/rdb"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     db.InitMysql(&c.DB),
		RDB:    rdb.InitRedis(&c.RDB),
	}
}
