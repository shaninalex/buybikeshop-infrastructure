package observers

import (
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(ProvideEmployeeObserver)
	_ = c.Provide(ProvidePermissionObserver)

	return nil
}
