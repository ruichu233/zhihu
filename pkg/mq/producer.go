package mq

type Producer interface {
	Publish(topic string, msg *MsgEntity) error
}
