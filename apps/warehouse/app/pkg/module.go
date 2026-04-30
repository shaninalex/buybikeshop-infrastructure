package pkg

import (
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	//_ = c.Provide(connector.ProvideDatasourceClient)

	return nil
}
