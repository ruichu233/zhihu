package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zhihu/pkg/rdb"
)

type Config struct {
	rest.RestConf
	RDB     rdb.Conf `json:"rdb" yaml:"rdb"`
	UserRPC zrpc.RpcClientConf
	AuthKey string
}
