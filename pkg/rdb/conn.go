package rdb

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

type Conf struct {
	Addr     string `json:"addr,omitempty" yaml:"addr"`
	Password string `json:"password,omitempty" yaml:"password"`
	DB       int    `json:"db,omitempty" yaml:"db"`
}

func InitRedis(c *Conf) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("rdb connect error: %v", err)

	}
	return rdb
}
