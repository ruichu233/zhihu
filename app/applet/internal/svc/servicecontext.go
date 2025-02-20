package svc

import (
	"zhihu/app/applet/internal/config"
	"zhihu/app/applet/internal/middleware"
	"zhihu/app/chat/chatclient"
	"zhihu/app/comment/commentclient"
	"zhihu/app/comment/pb/comment"
	"zhihu/app/feed/feedclient"
	"zhihu/app/follow/followclient"
	"zhihu/app/like/likeclient"
	"zhihu/app/like/pb/like"
	"zhihu/app/user/pb/user"
	"zhihu/app/user/userclient"
	"zhihu/app/video/pb/video"
	"zhihu/app/video/videoclient"

	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	UserRPC             userclient.User
	VideoRPC            videoclient.Video
	LikeRPC             likeclient.Like
	FeedRPC             feedclient.Feed
	CommentRPC          commentclient.Comment
	FollowRPC           followclient.Follow
	ChatRPC             chatclient.Chat
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

	likeConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "like.rpc",
		},
	})

	feedConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "feed.rpc",
		},
	})
	commentConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "comment.rpc",
		},
	})
	followConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "follow.rpc",
		},
	})
	chatConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
			Hosts: []string{"IP_ADDRESS:2379"},
			Key:   "chat.rpc",
		},
	})
	return &ServiceContext{
		Config:              c,
		UserRPC:             user.NewUserClient(userConn.Conn()),
		VideoRPC:            video.NewVideoClient(videoConn.Conn()),
		LikeRPC:             like.NewLikeClient(likeConn.Conn()),
		CommentRPC:          comment.NewCommentClient(commentConn.Conn()),
		FeedRPC:             feedclient.NewFeed(feedConn),
		FollowRPC:           followclient.NewFollow(followConn),
		ChatRPC:             chatclient.NewChat(chatConn),
		AuthMiddleware:      middleware.NewAuthMiddleware().Handle,
		MustLoginMiddleware: middleware.NewMustLoginMiddleware().Handle,
	}
}
