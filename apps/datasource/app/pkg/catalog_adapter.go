package pkg

import (
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"context"
	"database/sql"
)

type CatalogAdapter struct {
	db *sql.DB
}

func (c CatalogAdapter) ProductList(ctx context.Context, request *pb.ProductListRequest) (*pb.ProductListReply, error) {
	return &pb.ProductListReply{}, nil
}

func (c CatalogAdapter) ProductGet(ctx context.Context, request *pb.ProductGetRequest) (*pb.ProductGetReply, error) {
	return &pb.ProductGetReply{}, nil
}

func (c CatalogAdapter) ProductVariantList(ctx context.Context, request *pb.ProductVariantListRequest) (*pb.ProductVariantListReply, error) {
	return &pb.ProductVariantListReply{}, nil
}

func (c CatalogAdapter) ProductVariantGet(ctx context.Context, request *pb.ProductVariantGetRequest) (*pb.ProductVariantGetReply, error) {
	return &pb.ProductVariantGetReply{}, nil
}
