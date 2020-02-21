package utils

import (
	"log"

	"github.com/jmoiron/sqlx"
)

// GetDbConnection create a db connection
func GetDbConnection(driver string, connectionString string) *sqlx.DB {
	db, err := sqlx.Connect(driver, connectionString)
	if err != nil {
		log.Fatalf("Unable to establish connection: %v\n", err)
	}
	return db
}
