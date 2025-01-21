package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zhihu/pkg/rdb"
)

type Config struct {
	zrpc.RpcServerConf
	//DBConf  db.Conf
	RDBConf rdb.Conf
	WorkId  uint16
}
