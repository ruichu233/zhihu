package config

import (
	"zhihu/pkg/db"
	"zhihu/pkg/rdb"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DBConf  *db.Conf
	RDBConf *rdb.Conf
}
