package models

import (
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Product struct {
	ID               uint64
	Title            string
	CategoryId       uint64
	BrandId          uint64
	Description      string
	ShortDescription string
	Variants         []ProductVariant
	CreatedAt        time.Time
}

func ToProtoProduct(p *Product) *pb.Product {
	variants := []*pb.ProductVariant{}
	for _, variant := range p.Variants {
		variants = append(variants, ToProtoProductVariant(&variant))
	}

	product := &pb.Product{
		Id:               p.ID,
		Title:            p.Title,
		CategoryId:       p.CategoryId,
		BrandId:          p.BrandId,
		Description:      p.Description,
		ShortDescription: p.ShortDescription,
		CreatedAt:        timestamppb.New(p.CreatedAt),
		Variants:         variants,
	}
	return product
}
