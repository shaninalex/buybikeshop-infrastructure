package observers

import (
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	if err := c.Provide(ProvideEmployeeObserver); err != nil {
		return err
	}
	if err := c.Provide(ProvidePermissionObserver); err != nil {
		return err
	}

	return c.Invoke(func(*EmployeeObserver, *PermissionObserver) {})
}
