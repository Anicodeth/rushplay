package transport

import (
	"context"
	"rushplay/api/generated/proto/transactionpb"
	entities "rushplay/internal/domain/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionUseCase struct {
	mock.Mock
}

func (m *MockTransactionUseCase) CreateTransaction(transaction *entities.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockTransactionUseCase) GetTransactionByID(id uint) (*entities.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Transaction), args.Error(1)
}

func (m *MockTransactionUseCase) GetTransactions() ([]*entities.Transaction, error) {
	args := m.Called()
	return args.Get(0).([]*entities.Transaction), args.Error(1)
}

func (m *MockTransactionUseCase) GetTransactionsByUserID(userID uint) ([]*entities.Transaction, error) {
	args := m.Called(userID)
	return args.Get(0).([]*entities.Transaction), args.Error(1)
}

func (m *MockTransactionUseCase) GetTransactionsByBetID(betID uint) ([]*entities.Transaction, error) {
	args := m.Called(betID)
	return args.Get(0).([]*entities.Transaction), args.Error(1)
}

func (m *MockTransactionUseCase) UpdateTransaction(transaction *entities.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockTransactionUseCase) DeleteTransaction(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTransactionUseCase) UpdateTransactionStatus(id uint, status entities.TransactionStatus) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockTransactionUseCase) GetUserTransactionHistory(userID uint, limit, offset int) ([]*entities.Transaction, error) {
	args := m.Called(userID, limit, offset)
	return args.Get(0).([]*entities.Transaction), args.Error(1)
}

func (m *MockTransactionUseCase) GetTransactionByReferenceID(referenceID string) (*entities.Transaction, error) {
	args := m.Called(referenceID)
	return args.Get(0).(*entities.Transaction), args.Error(1)
}

func (m *MockTransactionUseCase) ProcessDeposit(userID uint, amount float64, referenceID string) (*entities.Transaction, error) {
	args := m.Called(userID, amount, referenceID)
	return args.Get(0).(*entities.Transaction), args.Error(1)
}

func (m *MockTransactionUseCase) ProcessWithdrawal(userID uint, amount float64, referenceID string) (*entities.Transaction, error) {
	args := m.Called(userID, amount, referenceID)
	return args.Get(0).(*entities.Transaction), args.Error(1)
}

func (m *MockTransactionUseCase) ProcessBetTransaction(userID uint, betID uint, amount float64) (*entities.Transaction, error) {
	args := m.Called(userID, betID, amount)
	return args.Get(0).(*entities.Transaction), args.Error(1)
}

func (m *MockTransactionUseCase) ProcessWinTransaction(userID uint, betID uint, amount float64) (*entities.Transaction, error) {
	args := m.Called(userID, betID, amount)
	return args.Get(0).(*entities.Transaction), args.Error(1)
}

func (m *MockTransactionUseCase) GetUserBalance(userID uint) (float64, error) {
	args := m.Called(userID)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockTransactionUseCase) ValidateTransaction(transaction *entities.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func TestCreateTransaction(t *testing.T) {
	mockUseCase := new(MockTransactionUseCase)
	handler := NewTransactionHandler(mockUseCase)

	tests := []struct {
		name          string
		request       *transactionpb.CreateTransactionRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful creation",
			request: &transactionpb.CreateTransactionRequest{
				UserId: 1,
				Type:   "DEPOSIT",
				Amount: 100.0,
			},
			mockSetup: func() {
				mockUseCase.On("CreateTransaction", mock.MatchedBy(func(transaction *entities.Transaction) bool {
					return transaction.UserID == 1 && transaction.Type == entities.TransactionType("DEPOSIT") && transaction.Amount == 100.0
				})).Return(nil)
			},
			expectedError: false,
		},
		{
			name: "creation error",
			request: &transactionpb.CreateTransactionRequest{
				UserId: 2,
				Type:   "WITHDRAWAL",
				Amount: 200.0,
			},
			mockSetup: func() {
				mockUseCase.On("CreateTransaction", mock.MatchedBy(func(transaction *entities.Transaction) bool {
					return transaction.UserID == 2 && transaction.Type == entities.TransactionType("WITHDRAWAL") && transaction.Amount == 200.0
				})).Return(assert.AnError)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.CreateTransaction(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.request.UserId, resp.Transaction.UserId)
				assert.Equal(t, tt.request.Type, resp.Transaction.Type)
				assert.Equal(t, tt.request.Amount, resp.Transaction.Amount)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestGetTransaction(t *testing.T) {
	mockUseCase := new(MockTransactionUseCase)
	handler := NewTransactionHandler(mockUseCase)

	mockTransaction := &entities.Transaction{
		ID:          1,
		UserID:      1,
		Type:        entities.TransactionType("DEPOSIT"),
		Status:      entities.TransactionStatus("PENDING"),
		Amount:      100.0,
		ReferenceID: "REF123",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tests := []struct {
		name          string
		request       *transactionpb.GetTransactionRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful retrieval",
			request: &transactionpb.GetTransactionRequest{
				Id: 1,
			},
			mockSetup: func() {
				mockUseCase.On("GetTransactionByID", uint(1)).Return(mockTransaction, nil)
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.GetTransaction(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, uint64(mockTransaction.UserID), resp.Transaction.UserId)
				assert.Equal(t, string(mockTransaction.Type), resp.Transaction.Type)
				assert.Equal(t, string(mockTransaction.Status), resp.Transaction.Status)
				assert.Equal(t, mockTransaction.Amount, resp.Transaction.Amount)
				assert.Equal(t, mockTransaction.ReferenceID, resp.Transaction.ReferenceId)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestGetTransactions(t *testing.T) {
	mockUseCase := new(MockTransactionUseCase)
	handler := NewTransactionHandler(mockUseCase)

	mockTransactions := []*entities.Transaction{
		{
			ID:          1,
			UserID:      1,
			Type:        entities.TransactionType("DEPOSIT"),
			Status:      entities.TransactionStatus("PENDING"),
			Amount:      100.0,
			ReferenceID: "REF123",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			UserID:      2,
			Type:        entities.TransactionType("WITHDRAWAL"),
			Status:      entities.TransactionStatus("COMPLETED"),
			Amount:      200.0,
			ReferenceID: "REF456",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	tests := []struct {
		name          string
		request       *transactionpb.GetTransactionsRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name:    "successful retrieval",
			request: &transactionpb.GetTransactionsRequest{},
			mockSetup: func() {
				mockUseCase.On("GetTransactions").Return(mockTransactions, nil)
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.GetTransactions(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Len(t, resp.Transactions, 2)
				assert.Equal(t, uint64(mockTransactions[0].UserID), resp.Transactions[0].UserId)
				assert.Equal(t, uint64(mockTransactions[1].UserID), resp.Transactions[1].UserId)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestGetTransactionsByUserID(t *testing.T) {
	mockUseCase := new(MockTransactionUseCase)
	handler := NewTransactionHandler(mockUseCase)

	mockTransactions := []*entities.Transaction{
		{
			ID:          1,
			UserID:      1,
			Type:        entities.TransactionType("DEPOSIT"),
			Status:      entities.TransactionStatus("PENDING"),
			Amount:      100.0,
			ReferenceID: "REF123",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	tests := []struct {
		name          string
		request       *transactionpb.GetTransactionsByUserIDRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful retrieval",
			request: &transactionpb.GetTransactionsByUserIDRequest{
				UserId: 1,
			},
			mockSetup: func() {
				mockUseCase.On("GetTransactionsByUserID", uint(1)).Return(mockTransactions, nil)
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.GetTransactionsByUserID(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Len(t, resp.Transactions, 1)
				assert.Equal(t, uint64(mockTransactions[0].UserID), resp.Transactions[0].UserId)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestGetTransactionsByBetID(t *testing.T) {
	mockUseCase := new(MockTransactionUseCase)
	handler := NewTransactionHandler(mockUseCase)

	mockTransactions := []*entities.Transaction{
		{
			ID:          1,
			UserID:      1,
			BetID:       uintPtr(1),
			Type:        entities.TransactionType("BET"),
			Status:      entities.TransactionStatus("COMPLETED"),
			Amount:      100.0,
			ReferenceID: "REF123",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	tests := []struct {
		name          string
		request       *transactionpb.GetTransactionsByBetIDRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful retrieval",
			request: &transactionpb.GetTransactionsByBetIDRequest{
				BetId: 1,
			},
			mockSetup: func() {
				mockUseCase.On("GetTransactionsByBetID", uint(1)).Return(mockTransactions, nil)
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.GetTransactionsByBetID(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Len(t, resp.Transactions, 1)
				assert.Equal(t, uint64(*mockTransactions[0].BetID), *resp.Transactions[0].BetId)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestUpdateTransactionStatus(t *testing.T) {
	mockUseCase := new(MockTransactionUseCase)
	handler := NewTransactionHandler(mockUseCase)

	mockTransaction := &entities.Transaction{
		ID:          1,
		UserID:      1,
		Type:        entities.TransactionType("DEPOSIT"),
		Status:      entities.TransactionStatus("COMPLETED"),
		Amount:      100.0,
		ReferenceID: "REF123",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tests := []struct {
		name          string
		request       *transactionpb.UpdateTransactionStatusRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful status update",
			request: &transactionpb.UpdateTransactionStatusRequest{
				Id:     1,
				Status: "COMPLETED",
			},
			mockSetup: func() {
				mockUseCase.On("UpdateTransactionStatus", uint(1), entities.TransactionStatus("COMPLETED")).Return(nil)
				mockUseCase.On("GetTransactionByID", uint(1)).Return(mockTransaction, nil)
			},
			expectedError: false,
		},
		{
			name: "status update error",
			request: &transactionpb.UpdateTransactionStatusRequest{
				Id:     2,
				Status: "FAILED",
			},
			mockSetup: func() {
				mockUseCase.On("UpdateTransactionStatus", uint(2), entities.TransactionStatus("FAILED")).Return(assert.AnError)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.UpdateTransactionStatus(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, string(mockTransaction.Status), resp.Transaction.Status)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestGetUserTransactionHistory(t *testing.T) {
	mockUseCase := new(MockTransactionUseCase)
	handler := NewTransactionHandler(mockUseCase)

	mockTransactions := []*entities.Transaction{
		{
			ID:          1,
			UserID:      1,
			Type:        entities.TransactionType("DEPOSIT"),
			Status:      entities.TransactionStatus("COMPLETED"),
			Amount:      100.0,
			ReferenceID: "REF123",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	tests := []struct {
		name          string
		request       *transactionpb.GetUserTransactionHistoryRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful retrieval",
			request: &transactionpb.GetUserTransactionHistoryRequest{
				UserId: 1,
				Limit:  10,
				Offset: 0,
			},
			mockSetup: func() {
				mockUseCase.On("GetUserTransactionHistory", uint(1), 10, 0).Return(mockTransactions, nil)
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.GetUserTransactionHistory(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Len(t, resp.Transactions, 1)
				assert.Equal(t, uint64(mockTransactions[0].UserID), resp.Transactions[0].UserId)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestGetTransactionByReferenceID(t *testing.T) {
	mockUseCase := new(MockTransactionUseCase)
	handler := NewTransactionHandler(mockUseCase)

	mockTransaction := &entities.Transaction{
		ID:          1,
		UserID:      1,
		Type:        entities.TransactionType("DEPOSIT"),
		Status:      entities.TransactionStatus("COMPLETED"),
		Amount:      100.0,
		ReferenceID: "REF123",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tests := []struct {
		name          string
		request       *transactionpb.GetTransactionByReferenceIDRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful retrieval",
			request: &transactionpb.GetTransactionByReferenceIDRequest{
				ReferenceId: "REF123",
			},
			mockSetup: func() {
				mockUseCase.On("GetTransactionByReferenceID", "REF123").Return(mockTransaction, nil)
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.GetTransactionByReferenceID(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, mockTransaction.ReferenceID, resp.Transaction.ReferenceId)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestProcessDeposit(t *testing.T) {
	mockUseCase := new(MockTransactionUseCase)
	handler := NewTransactionHandler(mockUseCase)

	mockTransaction := &entities.Transaction{
		ID:          1,
		UserID:      1,
		Type:        entities.TransactionType("DEPOSIT"),
		Status:      entities.TransactionStatus("COMPLETED"),
		Amount:      100.0,
		ReferenceID: "REF123",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tests := []struct {
		name          string
		request       *transactionpb.ProcessDepositRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful deposit processing",
			request: &transactionpb.ProcessDepositRequest{
				Id: 1,
			},
			mockSetup: func() {
				mockUseCase.On("ProcessDeposit", uint(1), float64(0), "").Return(mockTransaction, nil)
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.ProcessDeposit(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, uint64(mockTransaction.UserID), resp.Transaction.UserId)
				assert.Equal(t, string(mockTransaction.Type), resp.Transaction.Type)
				assert.Equal(t, mockTransaction.Amount, resp.Transaction.Amount)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestProcessWithdrawal(t *testing.T) {
	mockUseCase := new(MockTransactionUseCase)
	handler := NewTransactionHandler(mockUseCase)

	mockTransaction := &entities.Transaction{
		ID:          1,
		UserID:      1,
		Type:        entities.TransactionType("WITHDRAWAL"),
		Status:      entities.TransactionStatus("COMPLETED"),
		Amount:      100.0,
		ReferenceID: "REF123",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tests := []struct {
		name          string
		request       *transactionpb.ProcessWithdrawalRequest
		mockSetup     func()
		expectedError bool
	}{
		{
			name: "successful withdrawal processing",
			request: &transactionpb.ProcessWithdrawalRequest{
				Id: 1,
			},
			mockSetup: func() {
				mockUseCase.On("ProcessWithdrawal", uint(1), float64(0), "").Return(mockTransaction, nil)
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.ProcessWithdrawal(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, uint64(mockTransaction.UserID), resp.Transaction.UserId)
				assert.Equal(t, string(mockTransaction.Type), resp.Transaction.Type)
				assert.Equal(t, mockTransaction.Amount, resp.Transaction.Amount)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}

func TestGetUserBalance(t *testing.T) {
	mockUseCase := new(MockTransactionUseCase)
	handler := NewTransactionHandler(mockUseCase)

	tests := []struct {
		name            string
		request         *transactionpb.GetUserBalanceRequest
		mockSetup       func()
		expectedError   bool
		expectedBalance float64
	}{
		{
			name: "successful balance retrieval",
			request: &transactionpb.GetUserBalanceRequest{
				UserId: 1,
			},
			mockSetup: func() {
				mockUseCase.On("GetUserBalance", uint(1)).Return(1000.0, nil)
			},
			expectedError:   false,
			expectedBalance: 1000.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			ctx := context.Background()

			resp, err := handler.GetUserBalance(ctx, tt.request)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.expectedBalance, resp.Balance)
			}

			mockUseCase.AssertExpectations(t)
		})
	}
}