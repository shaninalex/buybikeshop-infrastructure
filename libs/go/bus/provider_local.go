package bus

import (
	"context"
	"fmt"
)

func ProvideLocalBus() Bus {
	return localBus{
		subscribers: make(map[EventType][]Callback),
	}
}

type localBus struct {
	subscribers map[EventType][]Callback
}

var _ Bus = (*localBus)(nil)

func (r localBus) Dispatch(ctx context.Context, eventType EventType, data any) {
	callbacks, ok := r.subscribers[eventType]
	if !ok {
		return
	}
	for _, cb := range callbacks {
		if err := cb(ctx, data); err != nil {
			// TODO: log error
			fmt.Println(err)
		}
	}
}

func (r localBus) Subscribe(eventType EventType, c Callback) {
	r.subscribers[eventType] = append(r.subscribers[eventType], c)
}
