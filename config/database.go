package config

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	const (
		filename = "database.db"
	)
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("[DATABASE]: Cannot initialize database: %s", err.Error()))
	}

	return db
}
