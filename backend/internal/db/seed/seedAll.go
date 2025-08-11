package seed

import (
	"fmt"
	"log"
	"reflect"
	config "schule/config/seed"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SeedAll(db *gorm.DB) error {
	silentDB := db.Session(&gorm.Session{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	for _, entry := range config.Seeds() {
		t := reflect.TypeOf(entry).Elem()
		tableName := t.Name()

		idField := t.Field(0).Name

		idValue := reflect.ValueOf(entry).Elem().Field(0).Interface()

		existing := reflect.New(t).Interface()
		err := silentDB.First(existing, fmt.Sprintf("%s = ?", idField), idValue).Error

		if err == gorm.ErrRecordNotFound {
			if err := db.Create(entry).Error; err != nil {
				return fmt.Errorf("error inserting %s with ID %v: %w", tableName, idValue, err)
			}
			log.Printf("Created %s with ID %v", tableName, idValue)
		} else if err != nil {
			return fmt.Errorf("error checking %s with ID %v: %w", tableName, idValue, err)
		}
	}

	log.Println("Database seeding completed successfully")
	return nil
}