package pkg

import (
	"buybikeshop/apps/warehouse/app/pkg/connector"

	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(connector.ProvideDatasourceClient)

	return nil
}
