package bus

import (
	"context"
)

type Callback func(ctx context.Context, data any) error

type Event struct {
	EventType EventType `json:"event_type"`
	Data      any       `json:"data"`
}

type EventType string

var (
	EmployeeCreatedEventType EventType = "EMPLOYEE_CREATED"
)
