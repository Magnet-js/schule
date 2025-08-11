package db

import (
	"schule/internal/db/seed"

	"gorm.io/gorm"
)

func SeedTables(db *gorm.DB) error {
	return seed.SeedAll(db)
}