package application

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	irepositories "rushplay/internal/domain/contracts/repository"
	iusecases "rushplay/internal/domain/contracts/usecase"
	entities "rushplay/internal/domain/entities"
)

type UserUseCase struct {
	userRepo  irepositories.IUserRepository
	secretKey string
}

func NewUserUseCase(userRepo irepositories.IUserRepository) iusecases.IUserUseCase {
	var secretKey string = os.Getenv("SECRET_KEY")

	return &UserUseCase{
		userRepo:  userRepo,
		secretKey: secretKey,
	}
}

func (u *UserUseCase) RegisterUser(user *entities.User) error {
	_, err := u.userRepo.GetUserByEmail(user.Email)

	if err == nil {
		return errors.New("email already in use")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("error checking existing user: %w", err)
	}

	hashedPassword, err := hashPassword(user.PasswordHash)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user.PasswordHash = hashedPassword

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

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
