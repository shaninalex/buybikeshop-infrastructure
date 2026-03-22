package models

import pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"

type Brand struct {
	ID    uint64 `db:"id"`
	Title string `db:"title"`
}

type Category struct {
	ID       uint64  `db:"id"`
	Title    string  `db:"title"`
	ParentId *uint64 `db:"parent_id"`
}

func ToProtoBrand(p Brand) *pb.Brand {
	return &pb.Brand{
		Id:    p.ID,
		Title: p.Title,
	}
}

func ToProtoBrands(p []Brand) []*pb.Brand {
	brands := make([]*pb.Brand, len(p))
	for _, b := range p {
		brands = append(brands, ToProtoBrand(b))
	}
	return brands
}

func ToProtoCategory(p *Category) *pb.Category {
	return &pb.Category{
		Id:       p.ID,
		Title:    p.Title,
		ParentId: p.ParentId,
	}
}

func ToProtoCatalogs(p []Category) []*pb.Category {
	brands := make([]*pb.Category, len(p))
	for _, b := range p {
		brands = append(brands, ToProtoCategory(&b))
	}
	return brands
}
