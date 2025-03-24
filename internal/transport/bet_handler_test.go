package transport

import (
	"context"
	"rushplay/api/generated/proto/betpb"
	domain "rushplay/internal/domain/contracts/usecase"
	entities "rushplay/internal/domain/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBetUseCase struct {
	mock.Mock
}

func (m *MockBetUseCase) CreateBet(bet *entities.Bet) error {
	args := m.Called(bet)
	return args.Error(0)
}

func (m *MockBetUseCase) GetBetByID(id uint) (*entities.Bet, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Bet), args.Error(1)
}

func (m *MockBetUseCase) GetBets() ([]*entities.Bet, error) {
	args := m.Called()
	return args.Get(0).([]*entities.Bet), args.Error(1)
}

func (m *MockBetUseCase) GetBetsByUserID(userID uint) ([]*entities.Bet, error) {
	args := m.Called(userID)
	return args.Get(0).([]*entities.Bet), args.Error(1)
}

func (m *MockBetUseCase) GetBetsByGameID(gameID uint) ([]*entities.Bet, error) {
	args := m.Called(gameID)
	return args.Get(0).([]*entities.Bet), args.Error(1)
}

func (m *MockBetUseCase) UpdateBetStatus(id uint, status entities.BetStatus) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockBetUseCase) ProcessBetResult(id uint, result domain.BetResult) error {
	args := m.Called(id, result)
	return args.Error(0)
}

func (m *MockBetUseCase) GetUserTotalBets(userID uint) (int64, error) {
	args := m.Called(userID)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockBetUseCase) GetUserTotalBetsByGameID(userID uint, gameID uint) (int64, error) {
	args := m.Called(userID, gameID)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockBetUseCase) GetUserTotalWinnings(userID uint) (float64, error) {
	args := m.Called(userID)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockBetUseCase) GetUserTotalWinningsByGameID(userID uint, gameID uint) (float64, error) {
	args := m.Called(userID, gameID)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockBetUseCase) DeleteBet(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBetUseCase) GetUserBettingHistory(userID uint, page int, pageSize int) ([]*entities.Bet, error) {
	args := m.Called(userID, page, pageSize)
	return args.Get(0).([]*entities.Bet), args.Error(1)
}

func (m *MockBetUseCase) PlaceBet(userID uint, gameID uint, amount float64, betType string) (*entities.Bet, error) {
	args := m.Called(userID, gameID, amount, betType)
	return args.Get(0).(*entities.Bet), args.Error(1)
}

func (m *MockBetUseCase) UpdateBet(bet *entities.Bet) error {
	args := m.Called(bet)
	return args.Error(0)
}

func TestCreateBet(t *testing.T) {
	mockUseCase := new(MockBetUseCase)
	handler := NewBetHandler(mockUseCase)

	tests := []struct {
		name          string
		request       *betpb.CreateBetRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful creation",
			request: &betpb.CreateBetRequest{
				UserId: 1,
				GameId: 1,
				Amount: 100.0,
			},
			mockSetup: func() {
				mockUseCase.On("CreateBet", mock.MatchedBy(func(bet *entities.Bet) bool {
					return bet.UserID == 1 && bet.GameID == 1 && bet.Amount == 100.0
				})).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "creation error",
			request: &betpb.CreateBetRequest{
				UserId: 2,
				GameId: 2,
				Amount: 200.0,
			},
			mockSetup: func() {
				mockUseCase.On("CreateBet", mock.MatchedBy(func(bet *entities.Bet) bool {
					return bet.UserID == 2 && bet.GameID == 2 && bet.Amount == 200.0
				})).Return(assert.AnError)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.CreateBet(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.request.UserId, resp.Bet.UserId)
				assert.Equal(t, tt.request.GameId, resp.Bet.GameId)
				assert.Equal(t, tt.request.Amount, resp.Bet.Amount)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestGetBet(t *testing.T) {
	mockUseCase := new(MockBetUseCase)
	handler := NewBetHandler(mockUseCase)

	mockBet := &entities.Bet{
		ID:        1,
		UserID:    1,
		GameID:    1,
		Amount:    100.0,
		Status:    entities.BetStatus("PENDING"),
		Winnings:  200.0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tests := []struct {
		name          string
		request       *betpb.GetBetRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful retrieval",
			request: &betpb.GetBetRequest{
				Id: 1,
			},
			mockSetup: func() {
				mockUseCase.On("GetBetByID", uint(1)).Return(mockBet, nil)
			},
			expectedError: false,
		},
		{
			name: "retrieval error",
			request: &betpb.GetBetRequest{
				Id: 2,
			},
			mockSetup: func() {
				mockUseCase.On("GetBetByID", uint(2)).Return((*entities.Bet)(nil), assert.AnError)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.GetBet(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, uint64(mockBet.UserID), resp.Bet.UserId)
				assert.Equal(t, uint64(mockBet.GameID), resp.Bet.GameId)
				assert.Equal(t, mockBet.Amount, resp.Bet.Amount)
				assert.Equal(t, string(mockBet.Status), resp.Bet.Status)
				assert.Equal(t, mockBet.Winnings, resp.Bet.PotentialWin)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestGetBets(t *testing.T) {
	mockUseCase := new(MockBetUseCase)
	handler := NewBetHandler(mockUseCase)

	mockBets := []*entities.Bet{
		{
			ID:        1,
			UserID:    1,
			GameID:    1,
			Amount:    100.0,
			Status:    entities.BetStatus("PENDING"),
			Winnings:  200.0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			UserID:    2,
			GameID:    2,
			Amount:    200.0,
			Status:    entities.BetStatus("WON"),
			Winnings:  400.0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	tests := []struct {
		name          string
		request       *betpb.GetBetsRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name:    "successful retrieval",
			request: &betpb.GetBetsRequest{},
			mockSetup: func() {
				mockUseCase.On("GetBets").Return(mockBets, nil)
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.GetBets(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Len(t, resp.Bets, 2)
				assert.Equal(t, uint64(mockBets[0].UserID), resp.Bets[0].UserId)
				assert.Equal(t, uint64(mockBets[1].UserID), resp.Bets[1].UserId)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestGetBetsByUserID(t *testing.T) {
	mockUseCase := new(MockBetUseCase)
	handler := NewBetHandler(mockUseCase)

	mockBets := []*entities.Bet{
		{
			ID:        1,
			UserID:    1,
			GameID:    1,
			Amount:    100.0,
			Status:    entities.BetStatus("PENDING"),
			Winnings:  200.0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	tests := []struct {
		name          string
		request       *betpb.GetBetsByUserIDRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful retrieval",
			request: &betpb.GetBetsByUserIDRequest{
				UserId: 1,
			},
			mockSetup: func() {
				mockUseCase.On("GetBetsByUserID", uint(1)).Return(mockBets, nil)
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.GetBetsByUserID(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Len(t, resp.Bets, 1)
				assert.Equal(t, uint64(mockBets[0].UserID), resp.Bets[0].UserId)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestGetBetsByGameID(t *testing.T) {
	mockUseCase := new(MockBetUseCase)
	handler := NewBetHandler(mockUseCase)

	mockBets := []*entities.Bet{
		{
			ID:        1,
			UserID:    1,
			GameID:    1,
			Amount:    100.0,
			Status:    entities.BetStatus("PENDING"),
			Winnings:  200.0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	tests := []struct {
		name          string
		request       *betpb.GetBetsByGameIDRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful retrieval",
			request: &betpb.GetBetsByGameIDRequest{
				GameId: 1,
			},
			mockSetup: func() {
				mockUseCase.On("GetBetsByGameID", uint(1)).Return(mockBets, nil)
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.GetBetsByGameID(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Len(t, resp.Bets, 1)
				assert.Equal(t, uint64(mockBets[0].GameID), resp.Bets[0].GameId)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestUpdateBetStatus(t *testing.T) {
	mockUseCase := new(MockBetUseCase)
	handler := NewBetHandler(mockUseCase)

	mockBet := &entities.Bet{
		ID:        1,
		UserID:    1,
		GameID:    1,
		Amount:    100.0,
		Status:    entities.BetStatus("WON"),
		Winnings:  200.0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tests := []struct {
		name          string
		request       *betpb.UpdateBetStatusRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful status update",
			request: &betpb.UpdateBetStatusRequest{
				Id:     1,
				Status: "WON",
			},
			mockSetup: func() {
				mockUseCase.On("UpdateBetStatus", uint(1), entities.BetStatus("WON")).Return(nil)
				mockUseCase.On("GetBetByID", uint(1)).Return(mockBet, nil)
			},
			expectedError: false,
		},
		{
			name: "status update error",
			request: &betpb.UpdateBetStatusRequest{
				Id:     2,
				Status: "LOST",
			},
			mockSetup: func() {
				mockUseCase.On("UpdateBetStatus", uint(2), entities.BetStatus("LOST")).Return(assert.AnError)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.UpdateBetStatus(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, string(mockBet.Status), resp.Bet.Status)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestProcessBetResult(t *testing.T) {
	mockUseCase := new(MockBetUseCase)
	handler := NewBetHandler(mockUseCase)

	mockBet := &entities.Bet{
		ID:        1,
		UserID:    1,
		GameID:    1,
		Amount:    100.0,
		Status:    entities.BetStatus("WON"),
		Winnings:  200.0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tests := []struct {
		name          string
		request       *betpb.ProcessBetResultRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful result processing",
			request: &betpb.ProcessBetResultRequest{
				Id:     1,
				Result: "WON",
			},
			mockSetup: func() {
				mockUseCase.On("ProcessBetResult", uint(1), domain.BetResult{
					Won:     true,
					Details: "WON",
				}).Return(nil)
				mockUseCase.On("GetBetByID", uint(1)).Return(mockBet, nil)
			},
			expectedError: false,
		},
		{
			name: "result processing error",
			request: &betpb.ProcessBetResultRequest{
				Id:     2,
				Result: "LOST",
			},
			mockSetup: func() {
				mockUseCase.On("ProcessBetResult", uint(2), domain.BetResult{
					Won:     false,
					Details: "LOST",
				}).Return(assert.AnError)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.ProcessBetResult(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, string(mockBet.Status), resp.Bet.Status)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestGetUserTotalBets(t *testing.T) {
	mockUseCase := new(MockBetUseCase)
	handler := NewBetHandler(mockUseCase)

	tests := []struct {
		name          string
		request       *betpb.GetUserTotalBetsRequest
		mockSetup     func()
		expectedError bool
		expectedTotal int32
	}{
		{
			name: "successful retrieval",
			request: &betpb.GetUserTotalBetsRequest{
				UserId: 1,
			},
			mockSetup: func() {
				mockUseCase.On("GetUserTotalBets", uint(1)).Return(int64(5), nil)
			},
			expectedError: false,
			expectedTotal: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.GetUserTotalBets(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.expectedTotal, resp.Total)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestGetUserTotalBetsByGameID(t *testing.T) {
	mockUseCase := new(MockBetUseCase)
	handler := NewBetHandler(mockUseCase)

	tests := []struct {
		name          string
		request       *betpb.GetUserTotalBetsByGameIDRequest
		mockSetup     func()
		expectedError bool
		expectedTotal int32
	}{
		{
			name: "successful retrieval",
			request: &betpb.GetUserTotalBetsByGameIDRequest{
				UserId: 1,
				GameId: 1,
			},
			mockSetup: func() {
				mockUseCase.On("GetUserTotalBetsByGameID", uint(1), uint(1)).Return(int64(3), nil)
			},
			expectedError: false,
			expectedTotal: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.GetUserTotalBetsByGameID(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.expectedTotal, resp.Total)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestGetUserTotalWinnings(t *testing.T) {
	mockUseCase := new(MockBetUseCase)
	handler := NewBetHandler(mockUseCase)

	tests := []struct {
		name          string
		request       *betpb.GetUserTotalWinningsRequest
		mockSetup     func()
		expectedError bool
		expectedTotal float64
	}{
		{
			name: "successful retrieval",
			request: &betpb.GetUserTotalWinningsRequest{
				UserId: 1,
			},
			mockSetup: func() {
				mockUseCase.On("GetUserTotalWinnings", uint(1)).Return(1000.0, nil)
			},
			expectedError: false,
			expectedTotal: 1000.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.GetUserTotalWinnings(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.expectedTotal, resp.Total)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}
