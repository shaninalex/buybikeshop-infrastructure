package catalog

import (
	"buybikeshop/apps/datasource/app/models"
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type CatalogRepository struct {
	db *sql.DB
}

func ProvideCatalogRepository(db *sql.DB) *CatalogRepository {
	return &CatalogRepository{
		db: db,
	}
}

func (s *CatalogRepository) ProductList(ctx context.Context) ([]models.Product, error) {
	q, _, err := goqu.From("catalog.products").
		Select("id", "title", "description", "short_description", "created_at").
		ToSQL()
	if err != nil {
		return nil, err
	}
	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := []models.Product{}
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.ShortDescription, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *CatalogRepository) ProductGet(ctx context.Context, productId uint64) (*models.Product, error) {
	q, _, err := goqu.From("catalog.products").Select().
		Select("id", "title", "description", "short_description", "created_at").
		Where(goqu.C("id").Eq(productId)).
		ToSQL()
	if err != nil {
		return nil, err
	}
	p := models.Product{}
	row := s.db.QueryRowContext(ctx, q)
	if err := row.Scan(&p.ID, &p.Title, &p.Description, &p.ShortDescription, &p.CreatedAt); err != nil {
		return nil, err
	}
	return &p, nil
}
