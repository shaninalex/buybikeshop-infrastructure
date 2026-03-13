package pkg

import (
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"context"
)

func ProvideCatalogServer() *CatalogServer {
	return &CatalogServer{}
}

type CatalogServer struct {
	catalogAdapter *CatalogAdapter
	pb.UnimplementedCatalogServer
}

func (c CatalogServer) ProductList(ctx context.Context, request *pb.ProductListRequest) (*pb.ProductListReply, error) {
	return c.catalogAdapter.ProductList(ctx, request)
}

func (c CatalogServer) ProductGet(ctx context.Context, request *pb.ProductGetRequest) (*pb.ProductGetReply, error) {
	return c.catalogAdapter.ProductGet(ctx, request)
}

func (c CatalogServer) ProductVariantList(ctx context.Context, request *pb.ProductVariantListRequest) (*pb.ProductVariantListReply, error) {
	return c.catalogAdapter.ProductVariantList(ctx, request)
}

func (c CatalogServer) ProductVariantGet(ctx context.Context, request *pb.ProductVariantGetRequest) (*pb.ProductVariantGetReply, error) {
	return c.catalogAdapter.ProductVariantGet(ctx, request)
}
