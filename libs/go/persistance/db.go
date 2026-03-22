package persistance

import (
	"buybikeshop/libs/go/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ProvideDB(config *config.Config) *sql.DB {
	connectionString := fmt.Sprintf(
		"host=%s, port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.String("database.host"),
		config.Int("database.port"),
		config.String("database.user"),
		config.String("database.password"),
		config.String("database.dbname"),
		config.String("database.sslmode"),
	)
	return Connect(connectionString)
}

func Connect(connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic("failed to connect database")
	}

	if err = db.Ping(); err != nil {
		panic("failed to ping database: " + err.Error())
	}
	return db
}
