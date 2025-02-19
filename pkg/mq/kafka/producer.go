package kafka

import (
	"context"
	"zhihu/pkg/mq"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	ctx    context.Context
	client *kafka.Writer
}

func (p *Producer) Publish(topic string, msg *mq.MsgEntity) error {
	kafkaMsg := kafka.Message{
		Topic: topic,
		Key:   []byte(msg.Key),
		Value: []byte(msg.Val),
	}
	if err := p.client.WriteMessages(p.ctx, kafkaMsg); err != nil {
		return err
	}
	return nil
}

func NewProducer(ctx context.Context, client *kafka.Writer) *Producer {
	return &Producer{
		ctx:    ctx,
		client: client,
	}
}

var _ mq.Producer = (*Producer)(nil)
