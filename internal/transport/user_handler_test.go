package transport

import (
	"context"
	"rushplay/api/generated/proto/userpb"
	entities "rushplay/internal/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) RegisterUser(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserUseCase) GetUserByID(id uint) (*entities.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserUseCase) UpdateUser(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserUseCase) DeleteUser(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestRegisterUser(t *testing.T) {
	mockUseCase := new(MockUserUseCase)
	handler := NewUserHandler(mockUseCase)

	tests := []struct {
		name          string
		request       *userpb.RegisterUserRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful registration",
			request: &userpb.RegisterUserRequest{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john@example.com",
				Password:  "password123",
			},
			mockSetup: func() {
				mockUseCase.On("RegisterUser", mock.MatchedBy(func(user *entities.User) bool {
					return user.FirstName == "John" && user.LastName == "Doe" && user.Email == "john@example.com"
				})).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "registration error",
			request: &userpb.RegisterUserRequest{
				FirstName: "Jane",
				LastName:  "Doe",
				Email:     "jane@example.com",
				Password:  "password123",
			},
			mockSetup: func() {
				mockUseCase.On("RegisterUser", mock.MatchedBy(func(user *entities.User) bool {
					return user.FirstName == "Jane" && user.LastName == "Doe" && user.Email == "jane@example.com"
				})).Return(assert.AnError)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.RegisterUser(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.request.FirstName, resp.User.FirstName)
				assert.Equal(t, tt.request.LastName, resp.User.LastName)
				assert.Equal(t, tt.request.Email, resp.User.Email)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestLoginUser(t *testing.T) {
	handler := NewUserHandler(nil) // LoginUser doesn't use the useCase

	tests := []struct {
		name          string
		request       *userpb.LoginRequest
		expectedError bool
		expectedToken string
	}{
		{
			name: "successful login",
			request: &userpb.LoginRequest{
				Email:    "test@example.com",
				Password: "password123",
			},
			expectedError: false,
			expectedToken: "mocked-jwt-token",
		},
		{
			name: "invalid credentials",
			request: &userpb.LoginRequest{
				Email:    "wrong@example.com",
				Password: "wrongpassword",
			},
			expectedError: true,
			expectedToken: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			resp, err := handler.LoginUser(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.expectedToken, resp.Token)
			}
		})
	}
}
