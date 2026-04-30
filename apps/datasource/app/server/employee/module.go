package employee

import "go.uber.org/dig"

func Module(c *dig.Container) error {
	_ = c.Provide(ProvideRepository)
	_ = c.Provide(ProvideAdapter)
	_ = c.Provide(ProvideServer)

	return nil
}
