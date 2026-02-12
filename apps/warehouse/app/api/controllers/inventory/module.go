package inventory

import "go.uber.org/dig"

func Module(c *dig.Container) error {
	_ = c.Provide(NewInventoryController)
	return nil
}
