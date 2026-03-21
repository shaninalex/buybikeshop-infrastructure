package models

import pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"

type Brand struct {
	ID    uint64
	Title string
}

type Category struct {
	ID       uint64
	Title    string
	ParentId uint64
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
		ParentId: &p.ParentId,
	}
}
