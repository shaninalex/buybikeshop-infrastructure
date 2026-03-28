package catalog

import (
	"buybikeshop/apps/datasource/app/models"
	"context"
	"database/sql"
	"errors"

	"github.com/doug-martin/goqu/v9"
)

var (
	ErrorCatalogUnableToSaveBrand      = errors.New("unable to save brand")
	ErrorCatalogUnableToDeleteBrand    = errors.New("unable to delete brand")
	ErrorCatalogUnableToSaveCategory   = errors.New("unable to save category")
	ErrorCatalogUnableToDeleteCategory = errors.New("unable to delete category")
	ErrorCatalogUnableToSaveProduct    = errors.New("unable to save catalog")
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

func (s *Repository) ProductSave(ctx context.Context, product models.Product) (*models.Product, error) {
	record := goqu.Record{
		"title":             product.Title,
		"category_id":       product.CategoryId,
		"brand_id":          product.BrandId,
		"description":       product.Description,
		"short_description": product.ShortDescription,
	}

	query := ""
	var err error
	if product.ID == 0 {
		query, _, err = goqu.Insert("catalog.products").Rows(record).Returning("catalog.products.id").ToSQL()
		if err != nil {
			return nil, err
		}
	} else {
		query, _, err = goqu.Update("catalog.products").Set(record).Where(goqu.Ex{"id": product.ID}).
			Returning("catalog.products.id").
			ToSQL()
		if err != nil {
			return nil, err
		}
	}
	if query == "" {
		return nil, ErrorCatalogUnableToSaveProduct
	}
	var id uint64
	err = s.db.QueryRowContext(ctx, query).Scan(&id)
	if err != nil {
		return nil, err
	}

	product.ID = id

	return &product, nil
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
		Select("id", "product_id", "title", "description", "sku", "barcode", "price", "currency", "created_at", "updated_at").
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
		Select("id", "product_id", "title", "description", "sku", "barcode", "created_at", "updated_at").
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
	q, _, err := goqu.From("catalog.brands").Select("id", "title").ToSQL()
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

func (s *Repository) BrandSave(ctx context.Context, brand models.Brand) (*models.Brand, error) {
	record := goqu.Record{"title": brand.Title}
	query := ""
	var err error
	if brand.ID == 0 {
		query, _, err = goqu.Insert("catalog.brands").Rows(record).Returning("catalog.brands.id").ToSQL()
		if err != nil {
			return nil, err
		}
	} else {
		query, _, err = goqu.Update("catalog.brands").Set(record).Where(goqu.Ex{"id": brand.ID}).
			Returning("catalog.brands.id").
			ToSQL()
		if err != nil {
			return nil, err
		}
	}
	if query == "" {
		return nil, ErrorCatalogUnableToSaveBrand
	}
	var id uint64
	err = s.db.QueryRowContext(ctx, query).Scan(&id)
	if err != nil {
		return nil, err
	}

	brand.ID = id

	return &brand, nil
}

func (s *Repository) BrandDelete(ctx context.Context, id uint64) error {
	q, _, err := goqu.Delete("catalog.brands").Where(goqu.Ex{"id": id}).ToSQL()
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
		return ErrorCatalogUnableToDeleteBrand
	}
	return nil
}

func (s *Repository) CategoryList(ctx context.Context) ([]models.Category, error) {
	categories := []models.Category{}
	q, _, err := goqu.From("catalog.categories").Select("id", "title", "parent_id").ToSQL()
	if err != nil {
		return nil, err
	}
	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b models.Category
		if err = rows.Scan(&b.ID, &b.Title, &b.ParentId); err != nil {
			return nil, err
		}
		categories = append(categories, b)
	}
	return categories, nil
}

func (s *Repository) CategorySave(ctx context.Context, category models.Category) (*models.Category, error) {
	record := goqu.Record{"title": category.Title, "parent_id": category.ParentId}
	query := ""
	var err error
	if category.ID == 0 {
		query, _, err = goqu.Insert("catalog.categories").Rows(record).Returning("catalog.categories.id").ToSQL()
		if err != nil {
			return nil, err
		}
	} else {
		query, _, err = goqu.Update("catalog.categories").Set(record).Where(goqu.Ex{"id": category.ID}).
			Returning("catalog.categories.id").
			ToSQL()
		if err != nil {
			return nil, err
		}
	}
	if query == "" {
		return nil, ErrorCatalogUnableToSaveCategory
	}
	var id uint64
	err = s.db.QueryRowContext(ctx, query).Scan(&id)
	if err != nil {
		return nil, err
	}

	category.ID = id

	return &category, nil
}

func (s *Repository) CategoryDelete(ctx context.Context, id uint64) error {
	q, _, err := goqu.Delete("catalog.categories").Where(goqu.Ex{"id": id}).ToSQL()
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
		return ErrorCatalogUnableToDeleteCategory
	}
	return nil
}
