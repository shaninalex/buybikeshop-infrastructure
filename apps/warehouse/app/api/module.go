package api

import (
	"buybikeshop/apps/warehouse/app/api/controllers/inventory"

	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = inventory.Module(c)

	_ = c.Provide(ProvideAPI)

	return nil
}
