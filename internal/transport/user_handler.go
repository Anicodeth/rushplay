package transport

import (
	"context"
	"errors"
	"rushplay/api/generated/proto/userpb"
	iusecases "rushplay/internal/domain/contracts/usecase"
	entities "rushplay/internal/domain/entities"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	useCase iusecases.IUserUseCase
}

func NewUserHandler(useCase iusecases.IUserUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

func (h *UserHandler) RegisterUser(ctx context.Context, req *userpb.RegisterUserRequest) (*userpb.RegisterUserResponse, error) {
	user := &entities.User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		PasswordHash: req.Password,
	}

	err := h.useCase.RegisterUser(user)
	if err != nil {
		return nil, err
	}

	userPb := &userpb.User{
		Id:        uint64(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Balance:   user.Balance,
		Role:      user.Role,
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
	}

	return &userpb.RegisterUserResponse{User: userPb}, nil
}

func (h *UserHandler) LoginUser(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	if req.Email != "test@example.com" || req.Password != "password123" {
		return nil, errors.New("invalid credentials")
	}

	token := "mocked-jwt-token"

	return &userpb.LoginResponse{Token: token}, nil
}
