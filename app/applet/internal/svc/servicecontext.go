package svc

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zhihu/app/applet/internal/config"
	"zhihu/app/applet/internal/middleware"
	"zhihu/app/feed/feedclient"
	"zhihu/app/like/likeclient"
	"zhihu/app/user/pb/user"
	"zhihu/app/user/userclient"
	"zhihu/app/video/pb/video"
	"zhihu/app/video/videoclient"
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

	//likeConn := zrpc.MustNewClient(zrpc.RpcClientConf{
	//	Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
	//		Hosts: []string{"127.0.0.1:2379"},
	//		Key:   "like.rpc",
	//	},
	//})

	//feedConn := zrpc.MustNewClient(zrpc.RpcClientConf{
	//	Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
	//		Hosts: []string{"127.0.0.1:2379"},
	//		Key:   "feed.rpc",
	//	},
	//})
	videoc := video.NewVideoClient(videoConn.Conn())
	ping, err := videoc.PublishVideo(context.Background(), &video.PublishRequest{
		AuthorId:    1,
		CoverUrl:    "",
		Description: "",
		Title:       "",
		VideoUrl:    "",
	})
	if err != nil {
		return nil
	}
	logx.Infof("ping: %d", ping.VideoId)
	return &ServiceContext{
		Config:   c,
		UserRPC:  user.NewUserClient(userConn.Conn()),
		VideoRPC: videoc,
		//LikeRPC:             likeclient.NewLike(likeConn),
		//FeedRPC:             feedclient.NewFeed(feedConn),
		AuthMiddleware:      middleware.NewAuthMiddleware().Handle,
		MustLoginMiddleware: middleware.NewMustLoginMiddleware().Handle,
	}
}
