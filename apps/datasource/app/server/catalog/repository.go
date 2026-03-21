package catalog

import (
	"buybikeshop/apps/datasource/app/models"
	"context"
	"database/sql"
	"errors"

	"github.com/doug-martin/goqu/v9"
)

type Repository struct {
	db *sql.DB
}

func ProvideRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (s *Repository) ProductList(ctx context.Context) ([]models.Product, error) {
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

func (s *Repository) ProductGet(ctx context.Context, productId uint64) (*models.Product, error) {
	q, _, err := goqu.From("catalog.products").
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
	variants, err := s.ProductVariantList(ctx, []uint64{productId})
	if err != nil {
		return nil, err
	}
	p.Variants = variants
	return &p, nil
}

func (s *Repository) ProductVariantList(ctx context.Context, productIds []uint64) ([]models.ProductVariant, error) {
	q, _, err := goqu.From("catalog.product_variants").Select().
		Select(
			"id",
			"product_id",
			"title",
			"description",
			"sku",
			"barcode",
			"price",
			"currency",
			"created_at",
			"updated_at",
		).
		Where(goqu.C("product_id").In(productIds)).
		ToSQL()
	if err != nil {
		return nil, err
	}
	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	productVariants := []models.ProductVariant{}
	for rows.Next() {
		var p models.ProductVariant
		err = rows.Scan(&p.Id, &p.ProductId, &p.Title, &p.Description, &p.Sku, &p.Barcode, &p.Price, &p.Currency, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		productVariants = append(productVariants, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return productVariants, nil
}

func (s *Repository) ProductVariantGet(ctx context.Context, productId uint64) (*models.ProductVariant, error) {
	q, _, err := goqu.From("catalog.product_variants").Select().
		Select(
			"id",
			"product_id",
			"title",
			"description",
			"sku",
			"barcode",
			"created_at",
			"updated_at",
		).
		Where(goqu.C("product_id").Eq(productId)).
		ToSQL()
	if err != nil {
		return nil, err
	}
	row := s.db.QueryRowContext(ctx, q)
	var p models.ProductVariant
	err = row.Scan(&p.Id, &p.ProductId, &p.Title, &p.Description, &p.Sku, &p.Barcode, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (s *Repository) BrandList(ctx context.Context) ([]models.Brand, error) {
	brands := []models.Brand{}
	q, _, err := goqu.From("catalog.brands").
		Select("id", "title").
		ToSQL()
	if err != nil {
		return nil, err
	}
	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b models.Brand
		if err = rows.Scan(&b.ID, &b.Title); err != nil {
			return nil, err
		}
		brands = append(brands, b)
	}
	return brands, nil
}

var (
	ErrorUnableToSaveBrand = errors.New("unable to save brand")
)

func (s *Repository) BrandSave(ctx context.Context, brand models.Brand) (*models.Brand, error) {
	query := ""
	var err error
	if brand.ID == 0 {
		query, _, err = goqu.Insert("catalog.brands").
			Rows(models.Brand{Title: "Greg"}).
			Returning("id").
			ToSQL()
		if err != nil {
			return nil, err
		}
	} else {
		query, _, err = goqu.Update("catalog.brands").
			Set(goqu.Record{"title": brand.Title}).
			Where(goqu.Ex{"id": brand.ID}).
			Returning("id").
			ToSQL()
		if err != nil {
			return nil, err
		}
	}
	if query == "" {
		return nil, ErrorUnableToSaveBrand
	}
	res, err := s.db.ExecContext(ctx, query)
	if err != nil {
		return nil, err
	}
	if id, err := res.LastInsertId(); err == nil {
		brand.ID = uint64(id)
	}

	return &brand, nil
}
