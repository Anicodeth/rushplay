package application

import (
	"errors"
	irepositories "rushplay/internal/domain/contracts/repository"
	entities "rushplay/internal/domain/entities"
)

type UserUseCase struct {
	userRepo irepositories.IUserRepository
}

func NewUserUseCase(userRepo irepositories.IUserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (u *UserUseCase) RegisterUser(user *entities.User) error {
	existingUser, _ := u.userRepo.GetUserByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email already in use")
	}

	user.PasswordHash = "hashed_" + user.PasswordHash

	return u.userRepo.CreateUser(user)
}

func (u *UserUseCase) GetUserByID(id uint) (*entities.User, error) {
	return u.userRepo.GetUserByID(id)
}

func (u *UserUseCase) UpdateUser(user *entities.User) error {
	return u.userRepo.UpdateUser(user)
}

func (u *UserUseCase) DeleteUser(id uint) error {
	return u.userRepo.DeleteUser(id)
}
