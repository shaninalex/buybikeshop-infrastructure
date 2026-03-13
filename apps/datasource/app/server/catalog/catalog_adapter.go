package catalog

import (
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"context"
)

type CatalogAdapter struct {
}

func ProvideCatalogAdapter() *CatalogAdapter {
	return &CatalogAdapter{}
}

func (c CatalogAdapter) ProductList(ctx context.Context, request *pb.ProductListRequest) (*pb.ProductListReply, error) {
	return &pb.ProductListReply{
		Products: []*pb.Product{
			{
				Id: 1,
			},
		},
	}, nil
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
