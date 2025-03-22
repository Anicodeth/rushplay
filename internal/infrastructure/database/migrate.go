package infrastructure

import (
	"log"

	"gorm.io/gorm"

	entities "rushplay/internal/domain/entities"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entities.User{},
		&entities.Game{},
		&entities.Bet{},
		&entities.Transaction{},
	)

	if err != nil {
		log.Printf("Error migrating database: %v\n", err)
		return err
	}

	log.Println("Database migrated successfully")
	return nil
}
