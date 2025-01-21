package svc

import (
	"github.com/redis/go-redis/v9"
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
	//videoConn := zrpc.MustNewClient(zrpc.RpcClientConf{
	//	Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
	//		Hosts: []string{"127.0.0.1:2379"},
	//		Key:   "video.rpc",
	//	},
	//})

	return &ServiceContext{
		Config: c,
		DB:     db.InitMysql(&c.DBConf),
		RDB:    rdb.InitRedis(&c.RDBConf),
		//VideoRPC: videoclient.NewVideo(videoConn),
	}
}
