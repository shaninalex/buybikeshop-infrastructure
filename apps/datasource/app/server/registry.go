package server

import (
	"buybikeshop/apps/datasource/app/server/catalog"
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"

	"go.uber.org/dig"
	"google.golang.org/grpc"
)

type Registry struct {
	Server *grpc.Server
}

type ApiDeps struct {
	dig.In

	CatalogServer *catalog.Server
}

func NewRegistry(deps ApiDeps, s *grpc.Server) *Registry {
	pb.RegisterCatalogServer(s, deps.CatalogServer)

	return &Registry{
		Server: s,
	}
}
