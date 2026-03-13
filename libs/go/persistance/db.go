package persistance

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DatabaseConfig interface {
	GetDatabaseUrl() string
}

func ProvideDB(config DatabaseConfig) *sql.DB {
	connectionString := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	return connect(connectionString)
}

func connect(connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
