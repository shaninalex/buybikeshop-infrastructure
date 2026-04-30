package bus

import (
	"context"
)

func ProvideRabbitMQBus(host string) Bus {
	return rabbitmqBus{}
}

type rabbitmqBus struct{}

var _ Bus = (*rabbitmqBus)(nil)

func (r rabbitmqBus) Dispatch(ctx context.Context, eventType EventType, data any) {
	//TODO implement me
	panic("implement me")
}

func (r rabbitmqBus) Subscribe(eventType EventType, c Callback) {
	//TODO implement me
	panic("implement me")
}
