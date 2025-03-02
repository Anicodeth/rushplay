package domain

import (
	domain "rushplay/internal/domain/entities"
)

type IUserUseCase interface {
	RegisterUser(user *domain.User) error
	GetUserByID(id uint) (*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id uint) error
}
