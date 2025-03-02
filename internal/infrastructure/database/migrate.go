package infrastructure

import (
	"log"

	"gorm.io/gorm"

	entities "rushplay/internal/domain/entities"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&entities.User{})

	if err != nil {
		log.Printf("Error migrating database: %v \n", err)
	} else {
		log.Println("Database migrated")
	}

	return err
}
