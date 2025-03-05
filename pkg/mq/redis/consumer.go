package redis

import (
	"context"
	"log"
	mq "zhihu/pkg/mq"
	"zhihu/pkg/utils"

	"github.com/redis/go-redis/v9"
)

type Consumer struct {
	ctx        context.Context
	cancel     context.CancelFunc
	rdb        *redis.Client
	topic      string // 消费的 topic
	groupID    string // 所属的消费者组
	consumerID string // 当前节点的消费者 id
}

// NewConsumer 初始化消费者
func NewConsumer(ctx context.Context, rdb *redis.Client, topic, groupID, consumerID string) *Consumer {
	ctx, cancel := context.WithCancel(ctx)
	return &Consumer{
		ctx:        ctx,
		cancel:     cancel,
		rdb:        rdb,
		topic:      topic,
		groupID:    groupID,
		consumerID: consumerID,
	}
}

// Run 启动消费者，使用 XREADGROUP 从消费者组读取消息
func (c *Consumer) Run(handler func(msg *mq.MsgEntity) error) {
	// 创建消费者组，如果已存在则忽略错误
	if err := c.rdb.XGroupCreateMkStream(c.ctx, c.topic, c.groupID, "$").Err(); err != nil && err != redis.Nil {
		log.Printf("Error creating group: %v", err)
	}

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		// 从 Stream 中读取消息
		messages, err := c.rdb.XReadGroup(c.ctx, &redis.XReadGroupArgs{
			Group:    c.groupID,
			Consumer: c.consumerID,
			Streams:  []string{c.topic, ">"},
			Count:    1,
			Block:    0,
			NoAck:    true,
		}).Result()
		if err != nil {
			log.Printf("Error reading from stream: %v", err)
			continue
		}

		// 处理读取到的消息
		for _, message := range messages {
			for _, msg := range message.Messages {
				var msgEntity mq.MsgEntity
				if err := utils.MapToStruct(msg.Values, &msgEntity); err != nil {
					log.Printf("Error mapping message to struct: %v", err)
					continue
				}

				// 调用处理函数
				if err := handler(&msgEntity); err != nil {
					log.Printf("Handler error for message %v: %v", msg.ID, err)
					continue
				}

				// 处理成功后确认消息
				if _, err := c.rdb.XAck(c.ctx, c.topic, c.groupID, msg.ID).Result(); err != nil {
					log.Printf("Failed to ack message %v: %v", msg.ID, err)
				}
			}
		}
	}
}

// Close 关闭消费者，停止消费
func (c *Consumer) Close() {
	c.cancel()
}

var _ mq.Consumer = (*Consumer)(nil)
