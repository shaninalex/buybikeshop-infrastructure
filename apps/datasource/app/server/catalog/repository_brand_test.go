package catalog_test

import (
	"buybikeshop/apps/datasource/app/models"
	"buybikeshop/apps/datasource/app/server/catalog"
	"buybikeshop/tdata"
	"context"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Datasource_BrandSave(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)

	brand := &models.Brand{Title: "Brand"}
	brand, err := repo.BrandSave(ctx, *brand)
	assert.Nil(t, err)
	assert.NotEmpty(t, brand.ID)
	assert.GreaterOrEqual(t, brand.ID, uint64(1))
}

func Test_Datasource_BrandUniqueName(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)

	_, _ = repo.BrandSave(ctx, models.Brand{Title: "Brand"})

	_, err := repo.BrandSave(ctx, models.Brand{Title: "Brand"})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "duplicate key value violates unique constraint")
}

func Test_Datasource_BrandUpdate(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)
	brand := &models.Brand{Title: "Brand"}

	brand, _ = repo.BrandSave(ctx, *brand)
	brand.Title = "Updated"
	brand, err := repo.BrandSave(ctx, *brand)
	assert.Nil(t, err)

	brands, err := repo.BrandList(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(brands))
	assert.Equal(t, brand.Title, brands[0].Title)
	assert.Equal(t, brand.ID, brands[0].ID)
}

func Test_Datasource_BrandDelete(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)

	brand := &models.Brand{Title: "Brand"}
	brand, _ = repo.BrandSave(ctx, *brand)

	err := repo.BrandDelete(ctx, brand.ID)
	assert.Nil(t, err)

	brands, err := repo.BrandList(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(brands))
}

func Test_Datasource_BrandList(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)

	_, _ = repo.BrandSave(ctx, models.Brand{Title: "Brand A"})
	_, _ = repo.BrandSave(ctx, models.Brand{Title: "Brand B"})
	_, _ = repo.BrandSave(ctx, models.Brand{Title: "Brand C"})

	brands, err := repo.BrandList(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(brands))

	for _, brand := range brands {
		assert.True(t, slices.Contains([]string{"Brand A", "Brand B", "Brand C"}, brand.Title))
	}
}
