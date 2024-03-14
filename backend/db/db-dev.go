//go:build prod

package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GetDB() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("../tmp/minitwit.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	return db, err
}
