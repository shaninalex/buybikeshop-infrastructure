package api

import (
	"buybikeshop/apps/admin/app/api/controllers/employee"

	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = employee.Module(c)

	_ = c.Provide(ProvideAPI)

	return nil
}
