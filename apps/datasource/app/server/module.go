package server

import (
	"buybikeshop/apps/datasource/app/server/catalog"
	"buybikeshop/apps/datasource/app/server/employee"
	"buybikeshop/apps/datasource/app/server/partners"

	"go.uber.org/dig"
)

func InitServerModules(c *dig.Container) error {
	if err := catalog.Module(c); err != nil {
		return err
	}

	if err := partners.Module(c); err != nil {
		return err
	}

	if err := employee.Module(c); err != nil {
		return err
	}

	return nil
}
