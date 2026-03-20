package models

import (
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProductVariant struct {
	Id              uint64
	ProductId       uint64
	InventoryItemId *uint64
	Title           string
	Description     string
	Sku             string
	Barcode         string
	Price           float32
	Currency        string
	CreatedAt       time.Time
	UpdatedAt       *time.Time
}

func ToProtoProductVariant(p *ProductVariant) *pb.ProductVariant {
	variant := &pb.ProductVariant{
		Id:              p.Id,
		ProductId:       p.ProductId,
		InventoryItemId: p.InventoryItemId,
		Title:           p.Title,
		Description:     p.Description,
		Sku:             p.Sku,
		Barcode:         p.Barcode,
		Price:           p.Price,
		Currency:        p.Currency,
		CreatedAt:       timestamppb.New(p.CreatedAt),
	}
	if p.UpdatedAt != nil {
		variant.UpdatedAt = timestamppb.New(*p.UpdatedAt)
	}
	return variant
}
