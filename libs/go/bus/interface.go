package bus

import (
	"buybikeshop/libs/go/config"
	"context"
)

type Bus interface {
	Dispatch(ctx context.Context, eventType EventType, data any)
	Subscribe(eventType EventType, c Callback)
}

func ProvideEventBus(config *config.Config) Bus {
	host := config.String("rabbitmq.host")
	if host != "" {
		return ProvideRabbitMQBus(host)
	}
	return ProvideLocalBus()
}
