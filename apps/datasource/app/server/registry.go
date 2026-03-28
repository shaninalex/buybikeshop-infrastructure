package server

import (
	"buybikeshop/apps/datasource/app/server/catalog"
	"buybikeshop/apps/datasource/app/server/partners"
	pbCatalog "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	pbPartners "buybikeshop/gen/grpc-buybikeshop-go/partners"

	"go.uber.org/dig"
	"google.golang.org/grpc"
)

type Registry struct {
	Server *grpc.Server
}

type ApiDeps struct {
	dig.In

	CatalogServer *catalog.Server
	PartnerServer *partners.Server
}

func NewRegistry(deps ApiDeps, s *grpc.Server) *Registry {
	pbCatalog.RegisterCatalogServer(s, deps.CatalogServer)
	pbPartners.RegisterPartnersServer(s, deps.PartnerServer)

	return &Registry{
		Server: s,
	}
}
