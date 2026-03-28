package models

import (
	pb "buybikeshop/gen/grpc-buybikeshop-go/partners"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Role struct {
	Id   uint64
	Role string
}

type PartnerRole struct {
	RoleId    uint64
	PartnerId uint64
}

type PartnerType string

var (
	PartnerTypeCompany PartnerType = "company"
	PartnerTypePerson  PartnerType = "person"
)

type PartnerContact struct {
	Id        uint64
	Contacts  string
	PartnerId uint64
	CreatedAt time.Time
}

type Partner struct {
	Id         uint64
	Title      string
	Type       PartnerType
	Active     bool
	IsSupplier bool
	CreatedAt  time.Time
	Roles      []uint64
	Contacts   []*PartnerContact
}

func ToModelPartner(p *pb.Partner) *Partner {
	roles := make([]uint64, len(p.Roles))
	contacts := make([]*PartnerContact, len(p.Contacts))

	for i, role := range p.Roles {
		roles[i] = role
	}

	for i, contact := range p.Contacts {
		contacts[i] = &PartnerContact{
			Id:        contact.Id,
			Contacts:  contact.Contacts,
			CreatedAt: p.CreatedAt.AsTime(),
		}
	}

	return &Partner{
		Id:         p.Id,
		Title:      p.Title,
		Type:       PartnerType(p.Type),
		Active:     p.Active,
		IsSupplier: p.IsSupplier,
		Roles:      roles,
		Contacts:   contacts,
	}
}

func ToPbPartner(p *Partner) *pb.Partner {
	contacts := make([]*pb.PartnerContact, len(p.Contacts))

	roleIds := make([]uint64, len(p.Roles))
	for i, role := range p.Roles {
		roleIds[i] = role
	}

	for i, contact := range p.Contacts {
		contacts[i] = &pb.PartnerContact{
			Id:        contact.Id,
			Contacts:  contact.Contacts,
			PartnerId: p.Id,
			CreatedAt: timestamppb.New(p.CreatedAt),
		}
	}

	return &pb.Partner{
		Id:         p.Id,
		Title:      p.Title,
		Type:       string(p.Type),
		Active:     p.Active,
		IsSupplier: p.IsSupplier,
		Roles:      roleIds,
		Contacts:   contacts,
		CreatedAt:  timestamppb.New(p.CreatedAt),
	}
}

func ToPbPartners(p []*Partner) []*pb.Partner {
	partners := make([]*pb.Partner, len(p))
	for i, partner := range p {
		partners[i] = ToPbPartner(partner)
	}
	return partners
}
