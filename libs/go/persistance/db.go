package persistance

import (
	"buybikeshop/libs/go/config"
	"database/sql"

	_ "github.com/lib/pq"
)

func ProvideDB(config *config.Config) *sql.DB {
	connectionString := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	return connect(connectionString)
}

func connect(connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.Ping(); err != nil {
		panic("failed to ping database: " + err.Error())
	}

	return db
}
