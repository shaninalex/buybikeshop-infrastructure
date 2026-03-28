package catalog

import "go.uber.org/dig"

func Module(c *dig.Container) error {
	_ = c.Provide(NewController)

	return nil
}
