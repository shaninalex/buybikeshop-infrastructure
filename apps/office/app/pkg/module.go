package pkg

import (
	"buybikeshop/apps/office/app/pkg/connector"

	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	if err := c.Provide(connector.ProvideDatasourceClient); err != nil {
		return err
	}

	return nil
}
