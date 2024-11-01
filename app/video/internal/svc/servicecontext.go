package svc

import (
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"zhihu/app/feed/pb/feed"
	"zhihu/app/user/pb/user"
	"zhihu/app/video/internal/config"
	"zhihu/pkg/db"
	"zhihu/pkg/oss"
	"zhihu/pkg/rdb"
)

type ServiceContext struct {
	Config     config.Config
	DB         *gorm.DB
	RDB        *redis.Client
	OSS        *minio.Client
	UserClient user.UserClient
	FeedRPC    feed.FeedClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
			Hosts:              []string{"127.0.0.1:2379"},
			Key:                "greet.rpc",
			User:               "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
			Pass:               "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
			CertFile:           "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
			CertKeyFile:        "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
			CACertFile:         "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
			InsecureSkipVerify: false, // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
		},
	})

	userClient := user.NewUserClient(conn.Conn())
	return &ServiceContext{
		Config:     c,
		DB:         db.InitMysql(&c.DBConf),
		RDB:        rdb.InitRedis(&c.RDBConf),
		OSS:        oss.InitMinio(&c.OSSConf),
		UserClient: userClient,
	}
}
