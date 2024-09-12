package config

import (
	"github.com/zeromicro/go-zero/rest"
	"zhihu/pkg/mysql"
	"zhihu/pkg/redis"
)

type Config struct {
	rest.RestConf
	DB  mysql.Conf `json:"db" yaml:"db"`
	RDB redis.Conf `json:"rdb" yaml:"rdb"`
}
