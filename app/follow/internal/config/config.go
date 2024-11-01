package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zhihu/pkg/db"
	"zhihu/pkg/rdb"
)

type Config struct {
	zrpc.RpcServerConf
	DB  db.Conf
	RDB rdb.Conf
}
