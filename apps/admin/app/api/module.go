package api

import (
	"buybikeshop/apps/admin/app/api/employee"

	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = employee.Module(c)

	_ = c.Provide(ProvideAPI)

	return nil
}
