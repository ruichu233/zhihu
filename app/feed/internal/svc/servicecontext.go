package svc

import (
	client "github.com/gorse-io/gorse-go"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"zhihu/app/feed/internal/config"
	"zhihu/app/follow/followclient"
	"zhihu/app/video/videoclient"
	"zhihu/pkg/rdb"
)

type ServiceContext struct {
	Config    config.Config
	Gorse     *client.GorseClient
	RDB       *redis.Client
	FollowRPC followclient.Follow
	VideoRPC  videoclient.Video
}

func NewServiceContext(c config.Config) *ServiceContext {
	videoConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "video.rpc",
		},
	})
	followConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "follow.rpc",
		},
	})
	gorse := client.NewGorseClient("http://127.0.0.1:8088", "api_key")
	return &ServiceContext{
		Config:    c,
		RDB:       rdb.InitRedis(&c.RDBConf),
		FollowRPC: followclient.NewFollow(followConn),
		VideoRPC:  videoclient.NewVideo(videoConn),
		Gorse:     gorse,
	}
}
