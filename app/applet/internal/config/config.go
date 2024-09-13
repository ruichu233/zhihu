package config

import (
	"github.com/zeromicro/go-zero/rest"
	"zhihu/pkg/db"
	"zhihu/pkg/rdb"
)

type Config struct {
	rest.RestConf
	DB  db.Conf  `json:"db" yaml:"db"`
	RDB rdb.Conf `json:"rdb" yaml:"rdb"`
}
