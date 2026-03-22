package tdata

import (
	"buybikeshop/libs/go/persistance"
	"context"
	"database/sql"
	"fmt"
)

func ProvideTestConnector() *sql.DB {
	connectionString := fmt.Sprintf(
		"host=%s, port=%d user=%s password=%s dbname=%s sslmode=%s",
		"localhost",
		5432,
		"postgres",
		"postgres",
		"buybikeshop_test",
		"disable",
	)

	return persistance.Connect(connectionString)
}

func BDClear(ctx context.Context, db *sql.DB) {
	_, err := db.QueryContext(ctx, `
        DELETE FROM catalog.product_variants;
        DELETE FROM catalog.products;
        DELETE FROM catalog.brands;
        DELETE FROM catalog.categories;
    `)
	if err != nil {
		panic(err)
	}
}
