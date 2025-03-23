package application

import (
	"errors"
	"testing"

	entities "rushplay/internal/domain/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBetRepository struct {
	mock.Mock
}

func (m *MockBetRepository) CreateBet(bet *entities.Bet) error {
	args := m.Called(bet)
	return args.Error(0)
}

func (m *MockBetRepository) GetBetByID(id uint) (*entities.Bet, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Bet), args.Error(1)
}

func (m *MockBetRepository) GetBets() ([]*entities.Bet, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.Bet), args.Error(1)
}

func (m *MockBetRepository) GetBetsByUserID(userID uint) ([]*entities.Bet, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.Bet), args.Error(1)
}

func (m *MockBetRepository) GetBetsByGameID(gameID uint) ([]*entities.Bet, error) {
	args := m.Called(gameID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.Bet), args.Error(1)
}

func (m *MockBetRepository) GetUserTotalBets(userID uint) (int64, error) {
	args := m.Called(userID)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockBetRepository) GetUserTotalBetsByGameID(userID uint, gameID uint) (int64, error) {
	args := m.Called(userID, gameID)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockBetRepository) GetUserTotalWinnings(userID uint) (float64, error) {
	args := m.Called(userID)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockBetRepository) GetUserTotalWinningsByGameID(userID uint, gameID uint) (float64, error) {
	args := m.Called(userID, gameID)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockBetRepository) UpdateBet(bet *entities.Bet) error {
	args := m.Called(bet)
	return args.Error(0)
}

func (m *MockBetRepository) DeleteBet(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBetRepository) UpdateBetStatus(id uint, status entities.BetStatus) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func TestCreateBet_Success(t *testing.T) {
	mockBetRepo := new(MockBetRepository)
	mockGameRepo := new(MockGameRepository)
	mockUserRepo := new(MockUserRepository)
	useCase := NewBetUseCase(mockBetRepo, mockGameRepo, mockUserRepo)

	game := &entities.Game{
		ID:     1,
		Status: entities.GameStatusActive,
		MinBet: 10,
		MaxBet: 100,
	}

	user := &entities.User{
		ID:      1,
		Balance: 1000,
	}

	bet := &entities.Bet{
		ID:     1,
		UserID: 1,
		GameID: 1,
		Amount: 50,
	}

	mockGameRepo.On("GetGameByID", uint(1)).Return(game, nil)
	mockUserRepo.On("GetUserByID", uint(1)).Return(user, nil)
	mockBetRepo.On("CreateBet", bet).Return(nil)

	err := useCase.CreateBet(bet)

	assert.NoError(t, err)
	assert.Equal(t, entities.BetStatusPending, bet.Status)
	mockBetRepo.AssertExpectations(t)
	mockGameRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestCreateBet_GameNotActive(t *testing.T) {
	mockBetRepo := new(MockBetRepository)
	mockGameRepo := new(MockGameRepository)
	mockUserRepo := new(MockUserRepository)
	useCase := NewBetUseCase(mockBetRepo, mockGameRepo, mockUserRepo)

	game := &entities.Game{
		ID:     1,
		Status: entities.GameStatusInactive,
		MinBet: 10,
		MaxBet: 100,
	}

	bet := &entities.Bet{
		UserID: 1,
		GameID: 1,
		Amount: 50,
	}

	mockGameRepo.On("GetGameByID", uint(1)).Return(game, nil)

	err := useCase.CreateBet(bet)

	assert.Error(t, err)
	assert.Equal(t, "game is not active", err.Error())
	mockGameRepo.AssertExpectations(t)
}

func TestCreateBet_InvalidAmount(t *testing.T) {
	mockBetRepo := new(MockBetRepository)
	mockGameRepo := new(MockGameRepository)
	mockUserRepo := new(MockUserRepository)
	useCase := NewBetUseCase(mockBetRepo, mockGameRepo, mockUserRepo)

	game := &entities.Game{
		ID:     1,
		Status: entities.GameStatusActive,
		MinBet: 10,
		MaxBet: 100,
	}

	bet := &entities.Bet{
		UserID: 1,
		GameID: 1,
		Amount: 5,
	}

	mockGameRepo.On("GetGameByID", uint(1)).Return(game, nil)

	err := useCase.CreateBet(bet)

	assert.Error(t, err)
	assert.Equal(t, "bet amount is outside game limits", err.Error())
	mockGameRepo.AssertExpectations(t)
}

func TestCreateBet_InsufficientBalance(t *testing.T) {
	mockBetRepo := new(MockBetRepository)
	mockGameRepo := new(MockGameRepository)
	mockUserRepo := new(MockUserRepository)
	useCase := NewBetUseCase(mockBetRepo, mockGameRepo, mockUserRepo)

	game := &entities.Game{
		ID:     1,
		Status: entities.GameStatusActive,
		MinBet: 10,
		MaxBet: 100,
	}

	user := &entities.User{
		ID:      1,
		Balance: 5,
	}

	bet := &entities.Bet{
		UserID: 1,
		GameID: 1,
		Amount: 50,
	}

	mockGameRepo.On("GetGameByID", uint(1)).Return(game, nil)
	mockUserRepo.On("GetUserByID", uint(1)).Return(user, nil)

	err := useCase.CreateBet(bet)

	assert.Error(t, err)
	assert.Equal(t, "insufficient balance", err.Error())
	mockGameRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestGetBetByID_Success(t *testing.T) {
	mockBetRepo := new(MockBetRepository)
	mockGameRepo := new(MockGameRepository)
	mockUserRepo := new(MockUserRepository)
	useCase := NewBetUseCase(mockBetRepo, mockGameRepo, mockUserRepo)

	bet := &entities.Bet{
		ID:     1,
		UserID: 1,
		GameID: 1,
		Amount: 50,
	}

	mockBetRepo.On("GetBetByID", uint(1)).Return(bet, nil)

	result, err := useCase.GetBetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, bet, result)
	mockBetRepo.AssertExpectations(t)
}

func TestGetBets_Success(t *testing.T) {
	mockBetRepo := new(MockBetRepository)
	mockGameRepo := new(MockGameRepository)
	mockUserRepo := new(MockUserRepository)
	useCase := NewBetUseCase(mockBetRepo, mockGameRepo, mockUserRepo)

	bets := []*entities.Bet{
		{
			ID:     1,
			UserID: 1,
			GameID: 1,
			Amount: 50,
		},
		{
			ID:     2,
			UserID: 2,
			GameID: 2,
			Amount: 100,
		},
	}

	mockBetRepo.On("GetBets").Return(bets, nil)

	result, err := useCase.GetBets()

	assert.NoError(t, err)
	assert.Equal(t, bets, result)
	mockBetRepo.AssertExpectations(t)
}

func TestGetBetsByUserID_Success(t *testing.T) {
	mockBetRepo := new(MockBetRepository)
	mockGameRepo := new(MockGameRepository)
	mockUserRepo := new(MockUserRepository)
	useCase := NewBetUseCase(mockBetRepo, mockGameRepo, mockUserRepo)

	bets := []*entities.Bet{
		{
			ID:     1,
			UserID: 1,
			GameID: 1,
			Amount: 50,
		},
		{
			ID:     2,
			UserID: 1,
			GameID: 2,
			Amount: 100,
		},
	}

	mockBetRepo.On("GetBetsByUserID", uint(1)).Return(bets, nil)

	result, err := useCase.GetBetsByUserID(1)

	assert.NoError(t, err)
	assert.Equal(t, bets, result)
	mockBetRepo.AssertExpectations(t)
}

func TestUpdateBetStatus_Success(t *testing.T) {
	mockBetRepo := new(MockBetRepository)
	mockGameRepo := new(MockGameRepository)
	mockUserRepo := new(MockUserRepository)
	useCase := NewBetUseCase(mockBetRepo, mockGameRepo, mockUserRepo)

	bet := &entities.Bet{
		ID:     1,
		Status: entities.BetStatusPending,
	}

	updatedBet := &entities.Bet{
		ID:     1,
		Status: entities.BetStatusWon,
	}

	mockBetRepo.On("GetBetByID", uint(1)).Return(bet, nil)
	mockBetRepo.On("UpdateBet", updatedBet).Return(nil)

	err := useCase.UpdateBetStatus(1, entities.BetStatusWon)

	assert.NoError(t, err)
	mockBetRepo.AssertExpectations(t)
}

func TestUpdateBetStatus_NotPending(t *testing.T) {
	mockBetRepo := new(MockBetRepository)
	mockGameRepo := new(MockGameRepository)
	mockUserRepo := new(MockUserRepository)
	useCase := NewBetUseCase(mockBetRepo, mockGameRepo, mockUserRepo)

	bet := &entities.Bet{
		ID:     1,
		Status: entities.BetStatusWon,
	}

	mockBetRepo.On("GetBetByID", uint(1)).Return(bet, nil)

	err := useCase.UpdateBetStatus(1, entities.BetStatusLost)

	assert.Error(t, err)
	assert.Equal(t, "can only update status of pending bets", err.Error())
	mockBetRepo.AssertExpectations(t)
}

func TestProcessBetResult_Success_Won(t *testing.T) {
	mockBetRepo := new(MockBetRepository)
	mockGameRepo := new(MockGameRepository)
	mockUserRepo := new(MockUserRepository)
	useCase := NewBetUseCase(mockBetRepo, mockGameRepo, mockUserRepo)

	bet := &entities.Bet{
		ID:     1,
		UserID: 1,
		Amount: 50,
		Status: entities.BetStatusPending,
	}

	user := &entities.User{
		ID:      1,
		Balance: 1000,
	}

	updatedBet := &entities.Bet{
		ID:       1,
		UserID:   1,
		Amount:   50,
		Status:   entities.BetStatusWon,
		Winnings: 100,
	}

	updatedUser := &entities.User{
		ID:      1,
		Balance: 1100,
	}

	mockBetRepo.On("GetBetByID", uint(1)).Return(bet, nil)
	mockUserRepo.On("GetUserByID", uint(1)).Return(user, nil)
	mockBetRepo.On("UpdateBet", updatedBet).Return(nil)
	mockUserRepo.On("UpdateUser", updatedUser).Return(nil)

	err := useCase.ProcessBetResult(1, true)

	assert.NoError(t, err)
	mockBetRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestProcessBetResult_Success_Lost(t *testing.T) {
	mockBetRepo := new(MockBetRepository)
	mockGameRepo := new(MockGameRepository)
	mockUserRepo := new(MockUserRepository)
	useCase := NewBetUseCase(mockBetRepo, mockGameRepo, mockUserRepo)

	bet := &entities.Bet{
		ID:     1,
		UserID: 1,
		Amount: 50,
		Status: entities.BetStatusPending,
	}

	user := &entities.User{
		ID:      1,
		Balance: 1000,
	}

	updatedBet := &entities.Bet{
		ID:       1,
		UserID:   1,
		Amount:   50,
		Status:   entities.BetStatusLost,
		Winnings: 0,
	}

	mockBetRepo.On("GetBetByID", uint(1)).Return(bet, nil)
	mockUserRepo.On("GetUserByID", uint(1)).Return(user, nil)
	mockBetRepo.On("UpdateBet", updatedBet).Return(nil)
	mockUserRepo.On("UpdateUser", user).Return(nil)

	err := useCase.ProcessBetResult(1, false)

	assert.NoError(t, err)
	mockBetRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestProcessBetResult_NotPending(t *testing.T) {
	mockBetRepo := new(MockBetRepository)
	mockGameRepo := new(MockGameRepository)
	mockUserRepo := new(MockUserRepository)
	useCase := NewBetUseCase(mockBetRepo, mockGameRepo, mockUserRepo)

	bet := &entities.Bet{
		ID:     1,
		Status: entities.BetStatusWon,
	}

	mockBetRepo.On("GetBetByID", uint(1)).Return(bet, nil)

	err := useCase.ProcessBetResult(1, true)

	assert.Error(t, err)
	assert.Equal(t, "can only process pending bets", err.Error())
	mockBetRepo.AssertExpectations(t)
}

func TestProcessBetResult_Error(t *testing.T) {
	mockBetRepo := new(MockBetRepository)
	mockGameRepo := new(MockGameRepository)
	mockUserRepo := new(MockUserRepository)
	useCase := NewBetUseCase(mockBetRepo, mockGameRepo, mockUserRepo)

	bet := &entities.Bet{
		ID:     1,
		UserID: 1,
		Amount: 50,
		Status: entities.BetStatusPending,
	}

	user := &entities.User{
		ID:      1,
		Balance: 1000,
	}

	updatedBet := &entities.Bet{
		ID:       1,
		UserID:   1,
		Amount:   50,
		Status:   entities.BetStatusWon,
		Winnings: 100,
	}

	mockBetRepo.On("GetBetByID", uint(1)).Return(bet, nil)
	mockUserRepo.On("GetUserByID", uint(1)).Return(user, nil)
	mockBetRepo.On("UpdateBet", updatedBet).Return(errors.New("database error"))

	err := useCase.ProcessBetResult(1, true)

	assert.Error(t, err)
	assert.Equal(t, "database error", err.Error())
	mockBetRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}
