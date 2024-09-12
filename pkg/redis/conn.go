package redis

import (
	"context"
	"log"
)
import redis "github.com/redis/go-redis/v9"

type Options struct {
	Addr     string
	Password string
	DB       int
}

func InitRedis(opt *Options) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     opt.Addr,
		Password: opt.Password,
		DB:       opt.DB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("redis connect error: %v", err)

	}
	return rdb
}
