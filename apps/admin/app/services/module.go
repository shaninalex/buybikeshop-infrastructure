package services

import (
	"buybikeshop/apps/admin/app/services/employee"

	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(employee.ProvideEmployeeService)

	return nil
}
