package api

import (
	"buybikeshop/apps/office/app/api/controllers/catalog"
	"buybikeshop/apps/office/app/api/controllers/partners"

	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = partners.Module(c)
	_ = catalog.Module(c)

	_ = c.Provide(ProvideAPI)

	return nil
}
