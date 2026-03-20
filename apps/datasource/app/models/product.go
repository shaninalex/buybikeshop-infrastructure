package models

import (
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Product struct {
	ID               uint64
	Title            string
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

	return &pb.Product{
		Id:               p.ID,
		Title:            p.Title,
		Description:      p.Description,
		ShortDescription: p.ShortDescription,
		CreatedAt:        timestamppb.New(p.CreatedAt),
		Variants:         variants,
	}
}
