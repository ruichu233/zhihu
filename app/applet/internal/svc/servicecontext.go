package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zhihu/app/applet/internal/config"
	"zhihu/app/applet/internal/middleware"
	"zhihu/app/feed/feedclient"
	"zhihu/app/like/likeclient"
	"zhihu/app/user/userclient"
	"zhihu/app/video/videoclient"
	"zhihu/pkg/rdb"
)

type ServiceContext struct {
	Config              config.Config
	UserRPC             userclient.User
	Redis               *redis.Client
	VideoRPC            videoclient.Video
	LikeRPC             likeclient.Like
	FeedRPC             feedclient.Feed
	AuthMiddleware      rest.Middleware
	MustLoginMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	userConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "user.rpc",
		},
	})
	videoConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "video.rpc",
		},
	})
	return &ServiceContext{
		Config:              c,
		UserRPC:             userclient.NewUser(userConn),
		VideoRPC:            videoclient.NewVideo(videoConn),
		Redis:               rdb.InitRedis(&c.RDB),
		AuthMiddleware:      middleware.NewAuthMiddleware().Handle,
		MustLoginMiddleware: middleware.NewMustLoginMiddleware().Handle,
	}
}
