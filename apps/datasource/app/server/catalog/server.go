package catalog

import (
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"context"
)

func ProvideServer(
	adapter *Adapter,
) *Server {
	return &Server{
		adapter: adapter,
	}
}

type Server struct {
	adapter *Adapter
	pb.UnimplementedCatalogServer
}

var _ pb.CatalogServer = &Server{}

func (c *Server) ProductSave(ctx context.Context, request *pb.ProductSaveRequest) (*pb.ProductSaveReply, error) {
	return c.adapter.ProductSave(ctx, request)
}

func (c *Server) ProductList(ctx context.Context, request *pb.ProductListRequest) (*pb.ProductListReply, error) {
	return c.adapter.ProductList(ctx, request)
}

func (c *Server) ProductGet(ctx context.Context, request *pb.ProductGetRequest) (*pb.ProductGetReply, error) {
	return c.adapter.ProductGet(ctx, request)
}

func (c *Server) ProductVariantList(ctx context.Context, request *pb.ProductVariantListRequest) (*pb.ProductVariantListReply, error) {
	return c.adapter.ProductVariantList(ctx, request)
}

func (c *Server) ProductVariantGet(ctx context.Context, request *pb.ProductVariantGetRequest) (*pb.ProductVariantGetReply, error) {
	return c.adapter.ProductVariantGet(ctx, request)
}

func (c *Server) BrandList(ctx context.Context, request *pb.BrandListRequest) (*pb.BrandListReply, error) {
	return c.adapter.BrandList(ctx, request)
}

func (c *Server) BrandSave(ctx context.Context, request *pb.BrandSaveRequest) (*pb.BrandSaveReply, error) {
	return c.adapter.BrandSave(ctx, request)
}

func (c *Server) BrandDelete(ctx context.Context, request *pb.BrandDeleteRequest) (*pb.BrandDeleteReply, error) {
	return c.adapter.BrandDelete(ctx, request)
}

func (c *Server) CategoryList(ctx context.Context, request *pb.CategoryListRequest) (*pb.CategoryListReply, error) {
	return c.adapter.CategoryList(ctx, request)
}

func (c *Server) CategorySave(ctx context.Context, request *pb.CategorySaveRequest) (*pb.CategorySaveReply, error) {
	return c.adapter.CategorySave(ctx, request)
}

func (c *Server) CategoryDelete(ctx context.Context, request *pb.CategoryDeleteRequest) (*pb.CategoryDeleteReply, error) {
	return c.adapter.CategoryDelete(ctx, request)
}
