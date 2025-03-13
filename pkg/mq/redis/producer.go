package redis

import (
	"context"
	"fmt"
	"zhihu/pkg/mq"

	"github.com/redis/go-redis/v9"
)

type Producer struct {
	ctx context.Context
	rdb *redis.Client
}

func (p *Producer) Publish(topic string, msg *mq.MsgEntity) error {
	// 将结构体转换为 map
	mp, err := msg.TransStructToMap()
	if err != nil {
		return err
	}
	result, err := p.rdb.XAdd(p.ctx, &redis.XAddArgs{
		Stream: topic,
		MaxLen: 10000,
		Values: mp,
	}).Result()
	if err != nil {
		return err
	}
	fmt.Println("Message added successfully. ID:", result)
	return nil
}

func NewProducer(ctx context.Context, rdb *redis.Client) *Producer {
	return &Producer{
		ctx: ctx,
		rdb: rdb,
	}
}

var _ mq.Producer = (*Producer)(nil)
