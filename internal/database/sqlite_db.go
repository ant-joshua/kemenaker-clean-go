package database

import (
	"clean_go/internal/domains"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SqliteConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&domains.Product{})
	if err != nil {
		return nil
	}

	return db
}
