package server

import (
	"buybikeshop/apps/datasource/app/server/catalog"

	"go.uber.org/dig"
)

func InitServerModules(c *dig.Container) error {
	if err := catalog.Module(c); err != nil {
		return err
	}

	return nil
}
