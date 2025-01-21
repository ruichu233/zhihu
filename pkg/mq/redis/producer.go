package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"zhihu/pkg/mq"
)

type Producer struct {
	ctx context.Context
	rdb *redis.Client
}

func (p *Producer) Publish(topic string, msg *mq.MsgEntity) error {
	p.rdb.XAdd(p.ctx, &redis.XAddArgs{
		Stream: topic,
		MaxLen: 10000,
		Values: msg.TransStructToMap(),
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
