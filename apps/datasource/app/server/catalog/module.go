package catalog

import "go.uber.org/dig"

func Module(c *dig.Container) error {
	_ = c.Provide(ProvideCatalogAdapter)
	_ = c.Provide(ProvideCatalogServer)

	return nil
}
