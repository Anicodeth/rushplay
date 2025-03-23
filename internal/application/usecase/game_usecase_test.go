package application

import (
	"testing"

	entities "rushplay/internal/domain/entities"

	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGameRepository struct {
	mock.Mock
}

func (m *MockGameRepository) CreateGame(game *entities.Game) error {
	args := m.Called(game)
	return args.Error(0)
}

func (m *MockGameRepository) GetGameByID(id uint) (*entities.Game, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Game), args.Error(1)
}

func (m *MockGameRepository) GetGames() ([]*entities.Game, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.Game), args.Error(1)
}

func (m *MockGameRepository) GetGamesByType(gameType entities.GameType) ([]*entities.Game, error) {
	args := m.Called(gameType)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.Game), args.Error(1)
}

func (m *MockGameRepository) GetActiveGames() ([]*entities.Game, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.Game), args.Error(1)
}

func (m *MockGameRepository) UpdateGame(game *entities.Game) error {
	args := m.Called(game)
	return args.Error(0)
}

func (m *MockGameRepository) DeleteGame(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockGameRepository) UpdateGameStatus(id uint, status entities.GameStatus) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func TestCreateGame_Success(t *testing.T) {
	mockRepo := new(MockGameRepository)
	useCase := NewGameUseCase(mockRepo)

	game := &entities.Game{
		Name:   "Test Game",
		Type:   entities.GameTypeSlots,
		Status: entities.GameStatusActive,
		MinBet: 10,
		MaxBet: 100,
	}

	mockRepo.On("CreateGame", game).Return(nil)

	err := useCase.CreateGame(game)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateGame_InvalidBetAmounts(t *testing.T) {
	mockRepo := new(MockGameRepository)
	useCase := NewGameUseCase(mockRepo)

	tests := []struct {
		name    string
		game    *entities.Game
		wantErr string
	}{
		{
			name: "negative min bet",
			game: &entities.Game{
				MinBet: -10,
				MaxBet: 100,
			},
			wantErr: "bet amounts must be positive",
		},
		{
			name: "negative max bet",
			game: &entities.Game{
				MinBet: 10,
				MaxBet: -100,
			},
			wantErr: "bet amounts must be positive",
		},
		{
			name: "min bet greater than max bet",
			game: &entities.Game{
				MinBet: 100,
				MaxBet: 10,
			},
			wantErr: "minimum bet cannot be greater than maximum bet",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := useCase.CreateGame(tt.game)
			assert.Error(t, err)
			assert.Equal(t, tt.wantErr, err.Error())
		})
	}
}

func TestGetGameByID_Success(t *testing.T) {
	mockRepo := new(MockGameRepository)
	useCase := NewGameUseCase(mockRepo)

	game := &entities.Game{
		ID:     1,
		Name:   "Test Game",
		Type:   entities.GameTypeSlots,
		Status: entities.GameStatusActive,
	}

	mockRepo.On("GetGameByID", uint(1)).Return(game, nil)

	result, err := useCase.GetGameByID(1)

	assert.NoError(t, err)
	assert.Equal(t, game, result)
	mockRepo.AssertExpectations(t)
}

func TestGetGames_Success(t *testing.T) {
	mockRepo := new(MockGameRepository)
	useCase := NewGameUseCase(mockRepo)

	games := []*entities.Game{
		{
			ID:     1,
			Name:   "Game 1",
			Type:   entities.GameTypeSlots,
			Status: entities.GameStatusActive,
		},
		{
			ID:     2,
			Name:   "Game 2",
			Type:   entities.GameTypePlinko,
			Status: entities.GameStatusActive,
		},
	}

	mockRepo.On("GetGames").Return(games, nil)

	result, err := useCase.GetGames()

	assert.NoError(t, err)
	assert.Equal(t, games, result)
	mockRepo.AssertExpectations(t)
}

func TestGetGamesByType_Success(t *testing.T) {
	mockRepo := new(MockGameRepository)
	useCase := NewGameUseCase(mockRepo)

	games := []*entities.Game{
		{
			ID:     1,
			Name:   "Slot Game 1",
			Type:   entities.GameTypeSlots,
			Status: entities.GameStatusActive,
		},
		{
			ID:     2,
			Name:   "Slot Game 2",
			Type:   entities.GameTypeSlots,
			Status: entities.GameStatusActive,
		},
	}

	mockRepo.On("GetGamesByType", entities.GameTypeSlots).Return(games, nil)

	result, err := useCase.GetGamesByType(entities.GameTypeSlots)

	assert.NoError(t, err)
	assert.Equal(t, games, result)
	mockRepo.AssertExpectations(t)
}

func TestGetActiveGames_Success(t *testing.T) {
	mockRepo := new(MockGameRepository)
	useCase := NewGameUseCase(mockRepo)

	games := []*entities.Game{
		{
			ID:     1,
			Name:   "Active Game 1",
			Status: entities.GameStatusActive,
		},
		{
			ID:     2,
			Name:   "Active Game 2",
			Status: entities.GameStatusActive,
		},
	}

	mockRepo.On("GetActiveGames").Return(games, nil)

	result, err := useCase.GetActiveGames()

	assert.NoError(t, err)
	assert.Equal(t, games, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateGame_Success(t *testing.T) {
	mockRepo := new(MockGameRepository)
	useCase := NewGameUseCase(mockRepo)

	game := &entities.Game{
		ID:     1,
		Name:   "Test Game",
		MinBet: 20,
		MaxBet: 200,
	}

	mockRepo.On("UpdateGame", game).Return(nil)

	err := useCase.UpdateGame(game)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateGame_InvalidBetAmounts(t *testing.T) {
	mockRepo := new(MockGameRepository)
	useCase := NewGameUseCase(mockRepo)

	tests := []struct {
		name    string
		game    *entities.Game
		wantErr string
	}{
		{
			name: "negative min bet",
			game: &entities.Game{
				MinBet: -10,
				MaxBet: 100,
			},
			wantErr: "bet amounts must be positive",
		},
		{
			name: "negative max bet",
			game: &entities.Game{
				MinBet: 10,
				MaxBet: -100,
			},
			wantErr: "bet amounts must be positive",
		},
		{
			name: "min bet greater than max bet",
			game: &entities.Game{
				MinBet: 100,
				MaxBet: 10,
			},
			wantErr: "minimum bet cannot be greater than maximum bet",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := useCase.UpdateGame(tt.game)
			assert.Error(t, err)
			assert.Equal(t, tt.wantErr, err.Error())
		})
	}
}

func TestDeleteGame_Success(t *testing.T) {
	mockRepo := new(MockGameRepository)
	useCase := NewGameUseCase(mockRepo)

	mockRepo.On("DeleteGame", uint(1)).Return(nil)

	err := useCase.DeleteGame(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateGameStatus_Success(t *testing.T) {
	mockRepo := new(MockGameRepository)
	useCase := NewGameUseCase(mockRepo)

	mockRepo.On("UpdateGameStatus", uint(1), entities.GameStatusInactive).Return(nil)

	err := useCase.UpdateGameStatus(1, entities.GameStatusInactive)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetGameByID_NotFound(t *testing.T) {
	mockRepo := new(MockGameRepository)
	useCase := NewGameUseCase(mockRepo)

	mockRepo.On("GetGameByID", uint(1)).Return(nil, errors.New("game not found"))

	result, err := useCase.GetGameByID(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "game not found", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestGetGames_Error(t *testing.T) {
	mockRepo := new(MockGameRepository)
	useCase := NewGameUseCase(mockRepo)

	mockRepo.On("GetGames").Return(nil, errors.New("database error"))

	result, err := useCase.GetGames()

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "database error", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestGetGamesByType_Error(t *testing.T) {
	mockRepo := new(MockGameRepository)
	useCase := NewGameUseCase(mockRepo)

	mockRepo.On("GetGamesByType", entities.GameTypeSlots).Return(nil, errors.New("database error"))

	result, err := useCase.GetGamesByType(entities.GameTypeSlots)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "database error", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestGetActiveGames_Error(t *testing.T) {
	mockRepo := new(MockGameRepository)
	useCase := NewGameUseCase(mockRepo)

	mockRepo.On("GetActiveGames").Return(nil, errors.New("database error"))

	result, err := useCase.GetActiveGames()

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "database error", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestUpdateGame_Error(t *testing.T) {
	mockRepo := new(MockGameRepository)
	useCase := NewGameUseCase(mockRepo)

	game := &entities.Game{
		ID:     1,
		Name:   "Test Game",
		MinBet: 20,
		MaxBet: 200,
	}

	mockRepo.On("UpdateGame", game).Return(errors.New("database error"))

	err := useCase.UpdateGame(game)

	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestDeleteGame_Error(t *testing.T) {
	mockRepo := new(MockGameRepository)
	useCase := NewGameUseCase(mockRepo)

	mockRepo.On("DeleteGame", uint(1)).Return(errors.New("database error"))

	err := useCase.DeleteGame(1)

	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestUpdateGameStatus_Error(t *testing.T) {
	mockRepo := new(MockGameRepository)
	useCase := NewGameUseCase(mockRepo)

	mockRepo.On("UpdateGameStatus", uint(1), entities.GameStatusInactive).Return(errors.New("database error"))

	err := useCase.UpdateGameStatus(1, entities.GameStatusInactive)

	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestCreateGame_Error(t *testing.T) {
	mockRepo := new(MockGameRepository)
	useCase := NewGameUseCase(mockRepo)

	game := &entities.Game{
		Name:   "Test Game",
		Type:   entities.GameTypeSlots,
		Status: entities.GameStatusActive,
		MinBet: 10,
		MaxBet: 100,
	}

	mockRepo.On("CreateGame", game).Return(errors.New("database error"))

	err := useCase.CreateGame(game)

	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
	mockRepo.AssertExpectations(t)
}
