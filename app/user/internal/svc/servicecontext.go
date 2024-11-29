package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"zhihu/app/user/internal/config"
	"zhihu/app/user/model"
	"zhihu/app/user/userclient"
	"zhihu/app/video/videoclient"
	"zhihu/pkg/db"
	"zhihu/pkg/rdb"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	RDB      *redis.Client
	VideoRPC videoclient.Video
	UserRPC  userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormDb := db.InitMysql(&c.DBConf)
	err := gormDb.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	videoConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
			Hosts:              []string{"127.0.0.1:2379"},
			Key:                "video.rpc",
			User:               "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
			Pass:               "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
			CertFile:           "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
			CertKeyFile:        "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
			CACertFile:         "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
			InsecureSkipVerify: false, // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
		},
	})

	return &ServiceContext{
		Config:   c,
		DB:       db.InitMysql(&c.DBConf),
		RDB:      rdb.InitRedis(&c.RDBConf),
		VideoRPC: videoclient.NewVideo(videoConn),
	}
}
