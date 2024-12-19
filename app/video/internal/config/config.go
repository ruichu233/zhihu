package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zhihu/pkg/db"
	"zhihu/pkg/oss"
	"zhihu/pkg/rdb"
)

type Config struct {
	zrpc.RpcServerConf
	DBConf  db.Conf
	RDBConf rdb.Conf
	OSSConf oss.Conf
	WorkId  uint16
}
