package db

import (
	"schule/internal/db/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Role{},
		&models.Form{},
		&models.Submission{},
	)
}