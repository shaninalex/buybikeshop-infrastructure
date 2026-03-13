package server

import (
	"buybikeshop/apps/datasource/app/server/catalog"
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"fmt"

	"go.uber.org/dig"
	"google.golang.org/grpc"
)

type PackageDeps struct {
	dig.In

	CatalogServer *catalog.CatalogServer
}

func NewDatasource(deps PackageDeps, grpcServer *grpc.Server) error {
	fmt.Println(" register NewDatasource with CatalogServer")
	pb.RegisterCatalogServer(grpcServer, deps.CatalogServer)
	return nil
}
