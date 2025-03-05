package redis

import (
	"context"
	"zhihu/pkg/mq"

	"github.com/redis/go-redis/v9"
)

type Producer struct {
	ctx context.Context
	rdb *redis.Client
}

func (p *Producer) Publish(topic string, msg *mq.MsgEntity) error {
	p.rdb.XAdd(p.ctx, &redis.XAddArgs{
		Stream: topic,
		MaxLen: 10000,
		Values: msg.Val,
	})
	return nil
}

func NewProducer(ctx context.Context, rdb *redis.Client) *Producer {
	return &Producer{
		ctx: ctx,
		rdb: rdb,
	}
}

var _ mq.Producer = (*Producer)(nil)
