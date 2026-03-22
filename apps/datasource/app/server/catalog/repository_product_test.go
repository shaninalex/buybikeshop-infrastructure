package catalog_test

import (
	"buybikeshop/apps/datasource/app/server/catalog"
	"buybikeshop/tdata"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func seedProduct(ctx context.Context, t *testing.T, db *sql.DB, title string) uint64 {
	var categoryId uint64
	err := db.QueryRowContext(ctx,
		`INSERT INTO catalog.categories (title) VALUES ($1) ON CONFLICT (title) DO UPDATE SET title = $1 RETURNING id`, "Test Category",
	).Scan(&categoryId)
	assert.Nil(t, err)

	var productId uint64
	err = db.QueryRowContext(ctx,
		`INSERT INTO catalog.products (title, category_id, description, short_description) VALUES ($1, $2, $3, $4) RETURNING id`,
		title, categoryId, title+" description", title+" short",
	).Scan(&productId)
	assert.Nil(t, err)
	return productId
}

func seedVariant(ctx context.Context, t *testing.T, db *sql.DB, productId uint64, title string, sku string, price float32) uint64 {
	var variantId uint64
	err := db.QueryRowContext(ctx,
		`INSERT INTO catalog.product_variants (product_id, title, description, sku, barcode, price, currency) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		productId, title, title+" description", sku, "BAR-"+sku, price, "eur",
	).Scan(&variantId)
	assert.Nil(t, err)
	return variantId
}

func Test_Datasource_ProductList(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)

	seedProduct(ctx, t, db, "Product A")
	seedProduct(ctx, t, db, "Product B")

	products, err := repo.ProductList(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(products))
}

func Test_Datasource_ProductListEmpty(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)

	products, err := repo.ProductList(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(products))
}

func Test_Datasource_ProductGet(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)

	productId := seedProduct(ctx, t, db, "Product A")
	seedVariant(ctx, t, db, productId, "Variant A", "SKU-A", 19.99)
	seedVariant(ctx, t, db, productId, "Variant B", "SKU-B", 29.99)

	product, err := repo.ProductGet(ctx, productId)
	assert.Nil(t, err)
	assert.Equal(t, productId, product.ID)
	assert.Equal(t, "Product A", product.Title)
	assert.Equal(t, "Product A description", product.Description)
	assert.Equal(t, "Product A short", product.ShortDescription)
	assert.Equal(t, 2, len(product.Variants))
}

func Test_Datasource_ProductGetNotFound(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)

	_, err := repo.ProductGet(ctx, 999999)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "no rows")
}

func Test_Datasource_ProductVariantList(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)

	pA := seedProduct(ctx, t, db, "Product A")
	pB := seedProduct(ctx, t, db, "Product B")
	seedVariant(ctx, t, db, pA, "Variant A1", "SKU-A1", 10.00)
	seedVariant(ctx, t, db, pA, "Variant A2", "SKU-A2", 20.00)
	seedVariant(ctx, t, db, pB, "Variant B1", "SKU-B1", 30.00)

	variants, err := repo.ProductVariantList(ctx, []uint64{pA, pB})
	assert.Nil(t, err)
	assert.Equal(t, 3, len(variants))

	variants, err = repo.ProductVariantList(ctx, []uint64{pA})
	assert.Nil(t, err)
	assert.Equal(t, 2, len(variants))
}

func Test_Datasource_ProductVariantListEmpty(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)

	variants, err := repo.ProductVariantList(ctx, []uint64{999999})
	assert.Nil(t, err)
	assert.Equal(t, 0, len(variants))
}

func Test_Datasource_ProductVariantGet(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)

	productId := seedProduct(ctx, t, db, "Product A")
	seedVariant(ctx, t, db, productId, "Variant A", "SKU-A", 49.99)

	variant, err := repo.ProductVariantGet(ctx, productId)
	assert.Nil(t, err)
	assert.Equal(t, productId, variant.ProductId)
	assert.Equal(t, "Variant A", variant.Title)
	assert.Equal(t, "Variant A description", variant.Description)
	assert.Equal(t, "SKU-A", variant.Sku)
	assert.Equal(t, "BAR-SKU-A", variant.Barcode)
}

func Test_Datasource_ProductVariantGetNotFound(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)

	_, err := repo.ProductVariantGet(ctx, 999999)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "no rows")
}
