package mq

type Consumer interface {
	Run(handler func(msg *MsgEntity) error)
	Close()
}
