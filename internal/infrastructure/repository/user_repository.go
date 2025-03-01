package infrastructure

import (
	"rushplay/internal/domain/contracts/repository"
	"rushplay/internal/domain/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}