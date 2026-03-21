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

func Test_Datasource_CategorySave(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)

	c := &models.Category{Title: "Category"}
	c, err := repo.CategorySave(ctx, *c)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.NotEmpty(t, c.ID)
	assert.GreaterOrEqual(t, c.ID, uint64(1))
}

func Test_Datasource_CategoryUniqueName(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)

	_, _ = repo.CategorySave(ctx, models.Category{Title: "Category"})

	_, err := repo.CategorySave(ctx, models.Category{Title: "Category"})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "duplicate key value violates unique constraint")
}

func Test_Datasource_CategorySaveParent(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)

	cA := &models.Category{Title: "Category A"}
	cA, err := repo.CategorySave(ctx, *cA)
	assert.NotNil(t, cA)

	cB := &models.Category{Title: "Category B", ParentId: &cA.ID}
	cB, err = repo.CategorySave(ctx, *cB)

	assert.Nil(t, err)
	assert.NotEmpty(t, cA.ID)
	assert.NotNil(t, cB.ParentId)
	assert.Equal(t, cA.ID, *cB.ParentId)
	assert.GreaterOrEqual(t, cA.ID, uint64(1))
}

func Test_Datasource_CategoryUpdate(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)
	c := &models.Category{Title: "Category"}

	c, _ = repo.CategorySave(ctx, *c)
	c.Title = "Updated"
	c, err := repo.CategorySave(ctx, *c)
	assert.Nil(t, err)

	categories, err := repo.CategoryList(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(categories))
	assert.Equal(t, c.Title, categories[0].Title)
	assert.Equal(t, c.ID, categories[0].ID)
}

func Test_Datasource_CategoryDelete(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)

	brand := &models.Category{Title: "Category"}
	brand, _ = repo.CategorySave(ctx, *brand)

	err := repo.CategoryDelete(ctx, brand.ID)
	assert.Nil(t, err)

	brands, err := repo.CategoryList(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(brands))
}

func Test_Datasource_CategoryList(t *testing.T) {
	ctx := context.Background()
	db := tdata.ProvideTestConnector()
	repo := catalog.ProvideRepository(db)
	tdata.BDClear(ctx, db)

	_, _ = repo.CategorySave(ctx, models.Category{Title: "Category A"})
	_, _ = repo.CategorySave(ctx, models.Category{Title: "Category B"})
	_, _ = repo.CategorySave(ctx, models.Category{Title: "Category C"})

	categories, err := repo.CategoryList(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(categories))

	for _, c := range categories {
		assert.True(t, slices.Contains([]string{"Category A", "Category B", "Category C"}, c.Title))
	}
}
