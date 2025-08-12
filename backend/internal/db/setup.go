package db

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func SetupDatabase(db *gorm.DB) error {
	log.Println("Setting up the database...")
	err := Migrate(db)
	if err != nil {
		return err
	}
	log.Println("Running seeders...")
	if err := SeedTables(db); err != nil {
		return fmt.Errorf("seeding failed: %w", err)
	}
	return nil
}
