package kafka

import (
	"context"
	"log"
	"zhihu/pkg/mq"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	ctx        context.Context
	cancel     context.CancelFunc
	client     *kafka.Reader
	topic      string // 消费的 topic
	groupID    string // 所属的消费者组
	consumerID string // 当前节点的消费者 id
}

func NewConsumer(ctx context.Context, client *kafka.Reader, topic, groupID, consumerID string) *Consumer {
	ctx, cancel := context.WithCancel(ctx)
	return &Consumer{
		ctx:        ctx,
		cancel:     cancel,
		client:     client,
		topic:      topic,
		groupID:    groupID,
		consumerID: consumerID,
	}
}

// Run 启动消费者，处理消息
func (c *Consumer) Run(handler func(msg *mq.MsgEntity) error) {
	log.Println("consumer start")
	defer log.Println("consumer stop")
	for {
		select {
		case <-c.ctx.Done():
			c.client.Close()
			return
		default:
		}
		msg, err := c.client.FetchMessage(c.ctx)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}
		var msgEntity mq.MsgEntity
		msgEntity.Val = string(msg.Value)
		if err := handler(&msgEntity); err != nil {
			log.Printf("Handler error for message %v: %v", msg.Offset, err)
			continue
		}
		if err := c.client.CommitMessages(c.ctx, msg); err != nil {
			log.Printf("Failed to commit message %v: %v", msg.Offset, err)
			continue
		}
	}
}
func (c *Consumer) Close() {
	c.cancel()
}
