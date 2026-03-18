package catalog

import (
	"buybikeshop/apps/datasource/app/models"
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"context"
)

type CatalogAdapter struct {
	catalogRepository *CatalogRepository
}

func ProvideCatalogAdapter(catalogRepository *CatalogRepository) *CatalogAdapter {
	return &CatalogAdapter{
		catalogRepository: catalogRepository,
	}
}

func (c CatalogAdapter) ProductList(ctx context.Context, request *pb.ProductListRequest) (*pb.ProductListReply, error) {
	// TODO: process request parameters
	products, err := c.catalogRepository.ProductList(ctx)
	if err != nil {
		return nil, err
	}

	pbProducts := make([]*pb.Product, len(products))
	for i, p := range products {
		pbProducts[i] = models.ToProtoProduct(&p)
	}

	return &pb.ProductListReply{
		Products: pbProducts,
	}, nil
}

func (c CatalogAdapter) ProductGet(ctx context.Context, request *pb.ProductGetRequest) (*pb.ProductGetReply, error) {
	p, err := c.catalogRepository.ProductGet(ctx, request.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.ProductGetReply{
		Product: models.ToProtoProduct(p),
	}, nil
}

func (c CatalogAdapter) ProductVariantList(ctx context.Context, request *pb.ProductVariantListRequest) (*pb.ProductVariantListReply, error) {
	return &pb.ProductVariantListReply{}, nil
}

func (c CatalogAdapter) ProductVariantGet(ctx context.Context, request *pb.ProductVariantGetRequest) (*pb.ProductVariantGetReply, error) {
	return &pb.ProductVariantGetReply{}, nil
}
