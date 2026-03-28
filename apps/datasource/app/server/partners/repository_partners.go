package partners

import (
	"buybikeshop/apps/datasource/app/models"
	"context"
	"database/sql"
	"errors"

	"github.com/doug-martin/goqu/v9"
)

var (
	ErrorPartnersUnableToSavePartner  = errors.New("unable to save partner")
	ErrorPartnerUnableToDeletePartner = errors.New("unable to delete partner")
)

type RepositoryPartners struct {
	db *sql.DB
}

func ProvideRepositoryPartners(db *sql.DB) *RepositoryPartners {
	return &RepositoryPartners{
		db: db,
	}
}

func (s RepositoryPartners) PartnersList(ctx context.Context) ([]*models.Partner, error) {
	q, _, err := goqu.From("partners.partners").
		Select("id", "active", "type", "title", "created_at").
		ToSQL()
	if err != nil {
		return nil, err
	}
	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	partners := map[uint64]*models.Partner{}
	for rows.Next() {
		var p models.Partner
		if err := rows.Scan(&p.Id, &p.Active, &p.Type, &p.Title, &p.CreatedAt); err != nil {
			return nil, err
		}
		partners[p.Id] = &p
	}

	// Roles
	if err = s.getPartnerRoles(ctx, partners); err != nil {
		return nil, err
	}

	// Contacts
	if err = s.getPartnerContacts(ctx, partners); err != nil {
		return nil, err
	}

	// Suppliers
	if err = s.getSuppliers(ctx, partners); err != nil {
		return nil, err
	}

	out := []*models.Partner{}
	for _, p := range partners {
		out = append(out, p)
	}
	return out, nil
}

func (s RepositoryPartners) PartnersSave(ctx context.Context, partner *models.Partner) (*models.Partner, error) {
	record := goqu.Record{
		"title":  partner.Title,
		"type":   partner.Type,
		"active": partner.Active,
	}

	query := ""
	var err error
	if partner.Id == 0 {
		query, _, err = goqu.Insert("partners.partners").Rows(record).Returning("partners.partners.id").ToSQL()
		if err != nil {
			return nil, err
		}
	} else {
		query, _, err = goqu.Update("catalog.products").Set(record).Where(goqu.Ex{"id": partner.Id}).
			Returning("catalog.products.id").
			ToSQL()
		if err != nil {
			return nil, err
		}
	}
	if query == "" {
		return nil, ErrorPartnersUnableToSavePartner
	}
	var id uint64
	err = s.db.QueryRowContext(ctx, query).Scan(&id)
	if err != nil {
		return nil, err
	}

	// TODO: save/update contacts
	// TODO: save/update roles
	// TODO: update supplier state ( create or delete one )

	partner.Id = id
	return partner, nil
}

func (s RepositoryPartners) PartnersDelete(ctx context.Context, id uint64) error {
	q, _, err := goqu.Delete("partners.partners").Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		return err
	}
	res, err := s.db.ExecContext(ctx, q)
	if err != nil {
		return err
	}
	r, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if r == 0 {
		return ErrorPartnerUnableToDeletePartner
	}
	return nil
}

func (s RepositoryPartners) getSuppliers(ctx context.Context, partners map[uint64]*models.Partner) error {
	q, _, _ := goqu.From("partners.suppliers").Select("partner_id").ToSQL()
	suppliersIds := []uint64{}
	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var partnerId uint64
		if err = rows.Scan(&partnerId); err != nil {
			return err
		}
		suppliersIds = append(suppliersIds, partnerId)
	}
	for _, p := range partners {
		for _, ps := range suppliersIds {
			if ps == p.Id {
				p.IsSupplier = true
			}
		}
	}

	return nil
}

func (s RepositoryPartners) getPartnerRoles(ctx context.Context, partners map[uint64]*models.Partner) error {
	q, _, _ := goqu.From("partners.partner_roles").Select("role_id", "partner_id").ToSQL()
	partnerRoles := []models.PartnerRole{}
	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var pr models.PartnerRole
		err = rows.Scan(&pr.RoleId, &pr.PartnerId)
		if err != nil {
			return err
		}
		partnerRoles = append(partnerRoles, pr)
	}
	for _, p := range partners {
		for _, pr := range partnerRoles {
			if pr.PartnerId == p.Id {
				p.Roles = append(p.Roles, pr.RoleId)
			}
		}
	}
	return nil
}

func (s RepositoryPartners) getPartnerContacts(ctx context.Context, partners map[uint64]*models.Partner) error {
	q, _, _ := goqu.From("partners.partner_contacts").
		Select("id", "contacts", "partner_id", "created_at").ToSQL()
	partnerContacts := []models.PartnerContact{}
	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var pc models.PartnerContact
		err = rows.Scan(&pc.Id, &pc.Contacts, &pc.PartnerId, &pc.CreatedAt)
		if err != nil {
			return err
		}
		partnerContacts = append(partnerContacts, pc)
	}
	for _, p := range partners {
		for _, pc := range partnerContacts {
			if pc.PartnerId == p.Id {
				p.Contacts = append(p.Contacts, &pc)
			}
		}
	}
	return nil
}
