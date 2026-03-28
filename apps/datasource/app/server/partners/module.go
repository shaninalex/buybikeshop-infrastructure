package partners

import "go.uber.org/dig"

func Module(c *dig.Container) error {
	_ = c.Provide(ProvideRepositoryRoles)
	_ = c.Provide(ProvideRepositoryPartners)

	_ = c.Provide(ProvideAdapter)
	_ = c.Provide(ProvideServer)

	return nil
}
