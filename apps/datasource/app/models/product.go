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
	CreatedAt        time.Time
}

func ToProtoProduct(p *Product) *pb.Product {
	return &pb.Product{
		Id:               p.ID,
		Title:            p.Title,
		Description:      p.Description,
		ShortDescription: p.ShortDescription,
		CreatedAt:        timestamppb.New(p.CreatedAt),
	}
}
