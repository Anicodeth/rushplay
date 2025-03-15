package application

import (
	"testing"

	entities "rushplay/internal/domain/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByID(id uint) (*entities.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByEmail(email string) (*entities.User, error) {
	args := m.Called(email)
	user := args.Get(0)
	if user != nil {
		return user.(*entities.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) UpdateUser(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUser(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestRegisterUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	user := &entities.User{
		Email:        "test@example.com",
		PasswordHash: "password123",
	}

	mockRepo.On("GetUserByEmail", user.Email).Return(nil, nil)
	mockRepo.On("CreateUser", mock.Anything).Return(nil)

	err := useCase.RegisterUser(user)

	assert.NoError(t, err)
	assert.Equal(t, "hashed_password123", user.PasswordHash)

	mockRepo.AssertExpectations(t)
}

func TestRegisterUser_EmailAlreadyInUse(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	existingUser := &entities.User{Email: "test@example.com"}

	mockRepo.On("GetUserByEmail", existingUser.Email).Return(existingUser, nil)

	err := useCase.RegisterUser(existingUser)

	assert.Error(t, err)
	assert.Equal(t, "email already in use", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestGetUserByID_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	user := &entities.User{ID: 1, Email: "test@example.com"}

	mockRepo.On("GetUserByID", user.ID).Return(user, nil)

	result, err := useCase.GetUserByID(user.ID)

	assert.NoError(t, err)
	assert.Equal(t, user, result)

	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	user := &entities.User{ID: 1, Email: "test@example.com"}

	mockRepo.On("UpdateUser", user).Return(nil)

	err := useCase.UpdateUser(user)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	mockRepo.On("DeleteUser", uint(1)).Return(nil)

	err := useCase.DeleteUser(1)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
