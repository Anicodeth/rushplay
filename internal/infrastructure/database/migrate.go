package infrastructure

import (
	"log"

	"gorm.io/gorm"

	"rushplay/internal/domain"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&domain.User{})

	if err != nil {
		log.Printf("Error migrating database: %v \n", err)
	} else {
		log.Println("Database migrated")
	}

	return err
}
