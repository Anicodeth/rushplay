package transport

import (
	"context"
	"rushplay/api/generated/proto/gamepb"
	domain "rushplay/internal/domain/contracts/usecase"
	entities "rushplay/internal/domain/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGameUseCase struct {
	mock.Mock
}

func (m *MockGameUseCase) CreateGame(game *entities.Game) error {
	args := m.Called(game)
	return args.Error(0)
}

func (m *MockGameUseCase) GetGameByID(id uint) (*entities.Game, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Game), args.Error(1)
}

func (m *MockGameUseCase) GetGames() ([]*entities.Game, error) {
	args := m.Called()
	return args.Get(0).([]*entities.Game), args.Error(1)
}

func (m *MockGameUseCase) GetGamesByType(gameType entities.GameType) ([]*entities.Game, error) {
	args := m.Called(gameType)
	return args.Get(0).([]*entities.Game), args.Error(1)
}

func (m *MockGameUseCase) GetActiveGames() ([]*entities.Game, error) {
	args := m.Called()
	return args.Get(0).([]*entities.Game), args.Error(1)
}

func (m *MockGameUseCase) UpdateGame(game *entities.Game) error {
	args := m.Called(game)
	return args.Error(0)
}

func (m *MockGameUseCase) DeleteGame(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockGameUseCase) UpdateGameStatus(id uint, status entities.GameStatus) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockGameUseCase) GetGameStatistics(id uint) (*domain.GameStatistics, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.GameStatistics), args.Error(1)
}

func (m *MockGameUseCase) ValidateGameBet(gameID uint, amount float64) error {
	args := m.Called(gameID, amount)
	return args.Error(0)
}

func TestCreateGame(t *testing.T) {
	mockUseCase := new(MockGameUseCase)
	handler := NewGameHandler(mockUseCase)

	tests := []struct {
		name          string
		request       *gamepb.CreateGameRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful creation",
			request: &gamepb.CreateGameRequest{
				Name:        "Test Game",
				Description: "Test Description",
				Type:        "SLOT",
				MinBet:      1.0,
				MaxBet:      100.0,
			},
			mockSetup: func() {
				mockUseCase.On("CreateGame", mock.MatchedBy(func(game *entities.Game) bool {
					return game.Name == "Test Game" && game.Description == "Test Description" &&
						game.Type == entities.GameType("SLOT") && game.MinBet == 1.0 && game.MaxBet == 100.0
				})).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "creation error",
			request: &gamepb.CreateGameRequest{
				Name:        "Error Game",
				Description: "Error Description",
				Type:        "SLOT",
				MinBet:      1.0,
				MaxBet:      100.0,
			},
			mockSetup: func() {
				mockUseCase.On("CreateGame", mock.MatchedBy(func(game *entities.Game) bool {
					return game.Name == "Error Game" && game.Description == "Error Description"
				})).Return(assert.AnError)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.CreateGame(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.request.Name, resp.Game.Name)
				assert.Equal(t, tt.request.Description, resp.Game.Description)
				assert.Equal(t, tt.request.Type, resp.Game.Type)
				assert.Equal(t, tt.request.MinBet, resp.Game.MinBet)
				assert.Equal(t, tt.request.MaxBet, resp.Game.MaxBet)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestGetGame(t *testing.T) {
	mockUseCase := new(MockGameUseCase)
	handler := NewGameHandler(mockUseCase)

	mockGame := &entities.Game{
		ID:          1,
		Name:        "Test Game",
		Description: "Test Description",
		Type:        entities.GameType("SLOT"),
		Status:      entities.GameStatus("ACTIVE"),
		MinBet:      1.0,
		MaxBet:      100.0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tests := []struct {
		name          string
		request       *gamepb.GetGameRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful retrieval",
			request: &gamepb.GetGameRequest{
				Id: 1,
			},
			mockSetup: func() {
				mockUseCase.On("GetGameByID", uint(1)).Return(mockGame, nil)
			},
			expectedError: false,
		},
		{
			name: "retrieval error",
			request: &gamepb.GetGameRequest{
				Id: 2,
			},
			mockSetup: func() {
				mockUseCase.On("GetGameByID", uint(2)).Return((*entities.Game)(nil), assert.AnError)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.GetGame(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, mockGame.Name, resp.Game.Name)
				assert.Equal(t, mockGame.Description, resp.Game.Description)
				assert.Equal(t, string(mockGame.Type), resp.Game.Type)
				assert.Equal(t, string(mockGame.Status), resp.Game.Status)
				assert.Equal(t, mockGame.MinBet, resp.Game.MinBet)
				assert.Equal(t, mockGame.MaxBet, resp.Game.MaxBet)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestGetGames(t *testing.T) {
	mockUseCase := new(MockGameUseCase)
	handler := NewGameHandler(mockUseCase)

	mockGames := []*entities.Game{
		{
			ID:          1,
			Name:        "Game 1",
			Description: "Description 1",
			Type:        entities.GameType("SLOT"),
			Status:      entities.GameStatus("ACTIVE"),
			MinBet:      1.0,
			MaxBet:      100.0,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "Game 2",
			Description: "Description 2",
			Type:        entities.GameType("TABLE"),
			Status:      entities.GameStatus("ACTIVE"),
			MinBet:      5.0,
			MaxBet:      500.0,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	tests := []struct {
		name          string
		request       *gamepb.GetGamesRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name:    "successful retrieval",
			request: &gamepb.GetGamesRequest{},
			mockSetup: func() {
				mockUseCase.On("GetGames").Return(mockGames, nil)
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.GetGames(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Len(t, resp.Games, 2)
				assert.Equal(t, mockGames[0].Name, resp.Games[0].Name)
				assert.Equal(t, mockGames[1].Name, resp.Games[1].Name)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestGetGamesByType(t *testing.T) {
	mockUseCase := new(MockGameUseCase)
	handler := NewGameHandler(mockUseCase)

	mockGames := []*entities.Game{
		{
			ID:          1,
			Name:        "Slot Game 1",
			Description: "Slot Description 1",
			Type:        entities.GameType("SLOT"),
			Status:      entities.GameStatus("ACTIVE"),
			MinBet:      1.0,
			MaxBet:      100.0,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	tests := []struct {
		name          string
		request       *gamepb.GetGamesByTypeRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful retrieval",
			request: &gamepb.GetGamesByTypeRequest{
				Type: "SLOT",
			},
			mockSetup: func() {
				mockUseCase.On("GetGamesByType", entities.GameType("SLOT")).Return(mockGames, nil)
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.GetGamesByType(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Len(t, resp.Games, 1)
				assert.Equal(t, mockGames[0].Name, resp.Games[0].Name)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestGetActiveGames(t *testing.T) {
	mockUseCase := new(MockGameUseCase)
	handler := NewGameHandler(mockUseCase)

	mockGames := []*entities.Game{
		{
			ID:          1,
			Name:        "Active Game 1",
			Description: "Active Description 1",
			Type:        entities.GameType("SLOT"),
			Status:      entities.GameStatus("ACTIVE"),
			MinBet:      1.0,
			MaxBet:      100.0,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	tests := []struct {
		name          string
		request       *gamepb.GetActiveGamesRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name:    "successful retrieval",
			request: &gamepb.GetActiveGamesRequest{},
			mockSetup: func() {
				mockUseCase.On("GetActiveGames").Return(mockGames, nil)
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.GetActiveGames(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Len(t, resp.Games, 1)
				assert.Equal(t, mockGames[0].Name, resp.Games[0].Name)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestUpdateGame(t *testing.T) {
	mockUseCase := new(MockGameUseCase)
	handler := NewGameHandler(mockUseCase)

	tests := []struct {
		name          string
		request       *gamepb.UpdateGameRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful update",
			request: &gamepb.UpdateGameRequest{
				Id:          1,
				Name:        "Updated Game",
				Description: "Updated Description",
				Type:        "SLOT",
				MinBet:      2.0,
				MaxBet:      200.0,
			},
			mockSetup: func() {
				mockUseCase.On("UpdateGame", mock.MatchedBy(func(game *entities.Game) bool {
					return game.ID == 1 && game.Name == "Updated Game" && game.Description == "Updated Description" &&
						game.Type == entities.GameType("SLOT") && game.MinBet == 2.0 && game.MaxBet == 200.0
				})).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "update error",
			request: &gamepb.UpdateGameRequest{
				Id:          2,
				Name:        "Error Game",
				Description: "Error Description",
				Type:        "SLOT",
				MinBet:      1.0,
				MaxBet:      100.0,
			},
			mockSetup: func() {
				mockUseCase.On("UpdateGame", mock.MatchedBy(func(game *entities.Game) bool {
					return game.ID == 2 && game.Name == "Error Game"
				})).Return(assert.AnError)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.UpdateGame(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.request.Name, resp.Game.Name)
				assert.Equal(t, tt.request.Description, resp.Game.Description)
				assert.Equal(t, tt.request.Type, resp.Game.Type)
				assert.Equal(t, tt.request.MinBet, resp.Game.MinBet)
				assert.Equal(t, tt.request.MaxBet, resp.Game.MaxBet)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestDeleteGame(t *testing.T) {
	mockUseCase := new(MockGameUseCase)
	handler := NewGameHandler(mockUseCase)

	tests := []struct {
		name          string
		request       *gamepb.DeleteGameRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful deletion",
			request: &gamepb.DeleteGameRequest{
				Id: 1,
			},
			mockSetup: func() {
				mockUseCase.On("DeleteGame", uint(1)).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "deletion error",
			request: &gamepb.DeleteGameRequest{
				Id: 2,
			},
			mockSetup: func() {
				mockUseCase.On("DeleteGame", uint(2)).Return(assert.AnError)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.DeleteGame(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestUpdateGameStatus(t *testing.T) {
	mockUseCase := new(MockGameUseCase)
	handler := NewGameHandler(mockUseCase)

	mockGame := &entities.Game{
		ID:          1,
		Name:        "Test Game",
		Description: "Test Description",
		Type:        entities.GameType("SLOT"),
		Status:      entities.GameStatus("INACTIVE"),
		MinBet:      1.0,
		MaxBet:      100.0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tests := []struct {
		name          string
		request       *gamepb.UpdateGameStatusRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful status update",
			request: &gamepb.UpdateGameStatusRequest{
				Id:     1,
				Status: "ACTIVE",
			},
			mockSetup: func() {
				mockUseCase.On("UpdateGameStatus", uint(1), entities.GameStatus("ACTIVE")).Return(nil)
				mockUseCase.On("GetGameByID", uint(1)).Return(mockGame, nil)
			},
			expectedError: false,
		},
		{
			name: "status update error",
			request: &gamepb.UpdateGameStatusRequest{
				Id:     2,
				Status: "ACTIVE",
			},
			mockSetup: func() {
				mockUseCase.On("UpdateGameStatus", uint(2), entities.GameStatus("ACTIVE")).Return(assert.AnError)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.UpdateGameStatus(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, mockGame.Name, resp.Game.Name)
				assert.Equal(t, string(mockGame.Status), resp.Game.Status)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}
