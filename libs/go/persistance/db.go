package persistance

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DatabaseConfig interface {
	GetDatabaseUrl() string
}

func ProvideDB(config DatabaseConfig) *sql.DB {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.GetDatabaseUrl(), config.GetDatabasePort(), config., password, dbname)
	return connect(connectionString)
}

func connect(connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
