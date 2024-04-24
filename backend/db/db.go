package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	conString, ok := os.LookupEnv("PSQL_CON_STR")

	if !ok {
		panic("PSQL_CON_STR not found in env")
	}

	log.Println("con string: ", conString)
	db, err := gorm.Open(postgres.Open(conString), &gorm.Config{})

	return db, err
}
