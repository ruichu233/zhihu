package svc

import (
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"zhihu/app/feed/feedclient"
	"zhihu/app/user/userclient"
	"zhihu/app/video/internal/config"
	"zhihu/app/video/internal/model"
	"zhihu/pkg/db"
	"zhihu/pkg/oss"
	"zhihu/pkg/rdb"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	RDB     *redis.Client
	OSS     *minio.Client
	UserRPC userclient.User
	FeedRPC feedclient.Feed
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormDb := db.InitMysql(&c.DBConf)
	if err := gormDb.AutoMigrate(&model.Video{}); err != nil {
		panic(err)
	}
	//userConn := zrpc.MustNewClient(zrpc.RpcClientConf{
	//	Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
	//		Hosts:              []string{"127.0.0.1:2379"},
	//		Key:                "user.rpc",
	//		User:               "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
	//		Pass:               "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
	//		CertFile:           "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
	//		CertKeyFile:        "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
	//		CACertFile:         "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
	//		InsecureSkipVerify: false, // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
	//	},
	//})
	//FeedConn := zrpc.MustNewClient(zrpc.RpcClientConf{
	//	Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
	//		Hosts:              []string{"127.0.0.1:2379"},
	//		Key:                "feed.rpc",
	//		User:               "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
	//		Pass:               "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
	//		CertFile:           "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
	//		CertKeyFile:        "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
	//		CACertFile:         "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
	//		InsecureSkipVerify: false, // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
	//	},
	//})

	return &ServiceContext{
		Config: c,
		DB:     db.InitMysql(&c.DBConf),
		RDB:    rdb.InitRedis(&c.RDBConf),
		OSS:    oss.InitMinio(&c.OSSConf),
		//UserRPC: userclient.NewUser(userConn),
		//FeedRPC: feedclient.NewFeed(FeedConn),
	}
}
