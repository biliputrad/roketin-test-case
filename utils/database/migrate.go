package database

import (
	"gorm.io/gorm"
	"test-case-roketin/models"
)

func MigrateTable(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Admin{}, &models.Movie{})

	return err
}
