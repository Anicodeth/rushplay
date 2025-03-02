package infrastructure

import (
	entities "rushplay/internal/domain/entities"
	irepositories "rushplay/internal/domain/contracts/repository"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) irepositories.IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByID(id uint) (*entities.User, error) {
	user := new(entities.User)
	err := r.db.Where("id = ?", id).First(user).Error
	return user, err
}

func (r *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	user := new(entities.User)
	err := r.db.Where("email = ?", email).First(user).Error
	return user, err
}

func (r *UserRepository) UpdateUser(user *entities.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&entities.User{}, id).Error
}