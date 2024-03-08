//go:build prod

package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	conString := os.Getenv("PSQL_CON_STR")
	db, err := gorm.Open(postgres.Open(conString), &gorm.Config{})

	return db, err
}
