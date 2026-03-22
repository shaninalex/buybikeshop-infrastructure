package catalog

import (
	"buybikeshop/apps/datasource/app/models"
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"context"
)

type Adapter struct {
	repository *Repository
}

func ProvideAdapter(catalogRepository *Repository) *Adapter {
	return &Adapter{
		repository: catalogRepository,
	}
}

func (c Adapter) ProductSave(ctx context.Context, request *pb.ProductSaveRequest) (*pb.ProductSaveReply, error) {
	p, err := c.repository.ProductSave(ctx, models.Product{
		ID:               request.Product.Id,
		Title:            request.Product.Title,
		CategoryId:       request.Product.CategoryId,
		BrandId:          request.Product.BrandId,
		Description:      request.Product.Description,
		ShortDescription: request.Product.ShortDescription,
		//Variants:         request.Product.Variants,
		CreatedAt: request.Product.CreatedAt.AsTime(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.ProductSaveReply{
		Product: models.ToProtoProduct(p),
	}, nil
}

func (c Adapter) ProductList(ctx context.Context, request *pb.ProductListRequest) (*pb.ProductListReply, error) {
	// TODO: process request parameters
	products, err := c.repository.ProductList(ctx)
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

func (c Adapter) ProductGet(ctx context.Context, request *pb.ProductGetRequest) (*pb.ProductGetReply, error) {
	p, err := c.repository.ProductGet(ctx, request.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.ProductGetReply{
		Product: models.ToProtoProduct(p),
	}, nil
}

func (c Adapter) ProductVariantList(ctx context.Context, request *pb.ProductVariantListRequest) (*pb.ProductVariantListReply, error) {
	variants, err := c.repository.ProductVariantList(ctx, request.GetProductIds())
	if err != nil {
		return nil, err
	}
	pbProductVariants := make([]*pb.ProductVariant, len(variants))
	for i, p := range variants {
		pbProductVariants[i] = models.ToProtoProductVariant(&p)
	}
	return &pb.ProductVariantListReply{
		Variants: pbProductVariants,
	}, nil
}

func (c Adapter) ProductVariantGet(ctx context.Context, request *pb.ProductVariantGetRequest) (*pb.ProductVariantGetReply, error) {
	variant, err := c.repository.ProductVariantGet(ctx, request.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.ProductVariantGetReply{
		Variant: models.ToProtoProductVariant(variant),
	}, nil
}

func (c Adapter) BrandList(ctx context.Context, request *pb.BrandListRequest) (*pb.BrandListReply, error) {
	brands, err := c.repository.BrandList(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.BrandListReply{
		Brands: models.ToProtoBrands(brands),
	}, nil
}

func (c Adapter) BrandSave(ctx context.Context, request *pb.BrandSaveRequest) (*pb.BrandSaveReply, error) {
	brand, err := c.repository.BrandSave(ctx, models.Brand{
		ID:    request.Brand.Id,
		Title: request.Brand.Title,
	})
	if err != nil {
		return nil, err
	}
	return &pb.BrandSaveReply{
		Brand: models.ToProtoBrand(*brand),
	}, nil
}

func (c Adapter) BrandDelete(ctx context.Context, request *pb.BrandDeleteRequest) (*pb.BrandDeleteReply, error) {
	err := c.repository.BrandDelete(ctx, request.GetId())
	if err != nil {
		return &pb.BrandDeleteReply{
			Status:  false,
			Message: err.Error(),
		}, err
	}
	return &pb.BrandDeleteReply{
		Status:  true,
		Message: "Success",
	}, nil
}

//goland:noinspection GoUnusedParameter
func (c Adapter) CategoryList(ctx context.Context, request *pb.CategoryListRequest) (*pb.CategoryListReply, error) {
	categories, err := c.repository.CategoryList(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.CategoryListReply{
		Brands: models.ToProtoCatalogs(categories),
	}, nil
}

func (c Adapter) CategorySave(ctx context.Context, request *pb.CategorySaveRequest) (*pb.CategorySaveReply, error) {
	category, err := c.repository.CategorySave(ctx, models.Category{
		ID:       request.Category.Id,
		Title:    request.Category.Title,
		ParentId: request.Category.ParentId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CategorySaveReply{
		Category: models.ToProtoCategory(category),
	}, nil
}

func (c Adapter) CategoryDelete(ctx context.Context, request *pb.CategoryDeleteRequest) (*pb.CategoryDeleteReply, error) {
	err := c.repository.CategoryDelete(ctx, request.GetId())
	if err != nil {
		return &pb.CategoryDeleteReply{
			Status:  false,
			Message: err.Error(),
		}, err
	}
	return &pb.CategoryDeleteReply{
		Status:  true,
		Message: "Success",
	}, nil
}
