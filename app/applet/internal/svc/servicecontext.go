package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zhihu/app/applet/internal/config"
	"zhihu/app/applet/internal/middleware"
	"zhihu/app/user/userclient"
	"zhihu/pkg/rdb"
)

type ServiceContext struct {
	Config              config.Config
	UserRPC             userclient.User
	Redis               *redis.Client
	AuthMiddleware      rest.Middleware
	MustLoginMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	userRPC := zrpc.MustNewClient(c.UserRPC)
	return &ServiceContext{
		Config:              c,
		UserRPC:             userclient.NewUser(userRPC),
		Redis:               rdb.InitRedis(&c.RDB),
		AuthMiddleware:      middleware.NewAuthMiddleware().Handle,
		MustLoginMiddleware: middleware.NewMustLoginMiddleware().Handle,
	}
}
