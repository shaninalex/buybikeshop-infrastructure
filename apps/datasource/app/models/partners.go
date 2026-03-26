package models

import "time"

type PartnerRole struct {
	Id   uint64
	Role string
}

type PartnerType string

var (
	PartnerTypeCompany PartnerType = "company"
	PartnerTypePerson  PartnerType = "person"
)

type PartnerContact struct {
	Id        uint64
	Contacts  string
	CreatedAt time.Time
}

type Partner struct {
	Id          uint64
	Title       string
	Type        PartnerType
	Active      bool
	IsSupplier  bool
	CreatedAt   time.Time
	PartnerRole []PartnerRole
	Contacts    []PartnerContact
}
