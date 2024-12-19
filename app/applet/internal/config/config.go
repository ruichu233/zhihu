package config

import (
	"github.com/zeromicro/go-zero/rest"
	"zhihu/pkg/rdb"
)

type Config struct {
	rest.RestConf
	RDB     rdb.Conf `json:"rdb" yaml:"rdb"`
	AuthKey string
}
