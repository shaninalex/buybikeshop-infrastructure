package partners

import (
	"buybikeshop/apps/datasource/app/models"
	"context"
	"database/sql"
	"errors"

	"github.com/doug-martin/goqu/v9"
)

var (
	ErrorPartnersUnableToSaveRole  = errors.New("unable to save role")
	ErrorPartnerUnableToDeleteRole = errors.New("unable to delete role")
)

type RepositoryRoles struct {
	db *sql.DB
}

func ProvideRepositoryRoles(db *sql.DB) *RepositoryRoles {
	return &RepositoryRoles{
		db: db,
	}
}

func (s RepositoryRoles) RolesGet(ctx context.Context) ([]models.PartnerRole, error) {
	q, _, _ := goqu.From("partners.roles").Select("id", "role").ToSQL()
	result, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	var roles []models.PartnerRole
	for result.Next() {
		var role models.PartnerRole
		if err = result.Scan(&role.Id, &role.Role); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (s RepositoryRoles) RolesSave(ctx context.Context, role *models.PartnerRole) (*models.PartnerRole, error) {
	record := goqu.Record{
		"role": role.Role,
	}

	query := ""
	var err error
	if role.Id == 0 {
		query, _, err = goqu.Insert("partners.roles").Rows(record).Returning("partners.roles.id").ToSQL()
		if err != nil {
			return nil, err
		}
	} else {
		query, _, err = goqu.Update("partners.roles").Set(record).Where(goqu.Ex{"id": role.Id}).
			Returning("partners.roles.id").
			ToSQL()
		if err != nil {
			return nil, err
		}
	}
	if query == "" {
		return nil, ErrorPartnersUnableToSaveRole
	}
	var id uint64
	err = s.db.QueryRowContext(ctx, query).Scan(&id)
	if err != nil {
		return nil, err
	}

	role.Id = id
	return role, nil
}

func (s RepositoryRoles) RolesDelete(ctx context.Context, id uint64) error {
	q, _, err := goqu.Delete("partners.roles").Where(goqu.Ex{"id": id}).ToSQL()
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
		return ErrorPartnerUnableToDeleteRole
	}
	return nil
}
