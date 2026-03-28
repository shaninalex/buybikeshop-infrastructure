package partners

import (
	"buybikeshop/apps/datasource/app/models"
	"context"
	"database/sql"
	"errors"
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
	var roles []*models.Partner
	return roles, nil
}

func (s RepositoryPartners) PartnersSave(ctx context.Context, role *models.Partner) (*models.Partner, error) {
	return nil, nil
}

func (s RepositoryPartners) PartnersDelete(ctx context.Context, id uint64) error {

	return nil
}
