package svc

import (
	redis "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"zhihu/app/applet/internal/config"
	"zhihu/pkg/mysql"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     mysql.InitMysql(&c.DB),
		Redis:  redis.InitRedis(&c.RDB),
	}
}
