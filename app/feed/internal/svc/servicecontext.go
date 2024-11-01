package svc

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"zhihu/app/feed/internal/config"
	"zhihu/app/follow/followclient"
)

type ServiceContext struct {
	Config    config.Config
	DB        *gorm.DB
	RDB       *redis.Client
	FollowRPC followclient.Follow
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
