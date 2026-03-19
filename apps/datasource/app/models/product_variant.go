package models

import (
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProductVariant struct {
	Id              uint64
	ProductId       uint64
	InventoryItemId uint64
	Title           string
	Description     string
	Sku             string
	Barcode         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func ToProtoProductVariant(p *ProductVariant) *pb.ProductVariant {
	return &pb.ProductVariant{
		Id:              p.Id,
		ProductId:       p.ProductId,
		InventoryItemId: p.InventoryItemId,
		Title:           p.Title,
		Description:     p.Description,
		Sku:             p.Sku,
		Barcode:         p.Barcode,
		UpdatedAt:       timestamppb.New(p.UpdatedAt),
		CreatedAt:       timestamppb.New(p.CreatedAt),
	}
}
