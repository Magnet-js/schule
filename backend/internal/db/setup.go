package db

import (
	"log"

	"gorm.io/gorm"
)

func SetupDatabase(db *gorm.DB) error {
	log.Println("Setting up the database...")
	err := Migrate(db)
	if err != nil {
		return err
	}
	return nil
}