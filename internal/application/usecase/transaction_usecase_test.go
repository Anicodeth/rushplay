package application

import (
	"testing"

	entities "rushplay/internal/domain/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) CreateTransaction(transaction *entities.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockTransactionRepository) GetTransactionByID(id uint) (*entities.Transaction, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) GetTransactions() ([]*entities.Transaction, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) GetTransactionsByUserID(userID uint) ([]*entities.Transaction, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) GetTransactionsByBetID(betID uint) ([]*entities.Transaction, error) {
	args := m.Called(betID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) UpdateTransaction(transaction *entities.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockTransactionRepository) DeleteTransaction(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTransactionRepository) UpdateTransactionStatus(id uint, status entities.TransactionStatus) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockTransactionRepository) GetUserTransactionHistory(userID uint, limit, offset int) ([]*entities.Transaction, error) {
	args := m.Called(userID, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) GetTransactionByReferenceID(referenceID string) (*entities.Transaction, error) {
	args := m.Called(referenceID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Transaction), args.Error(1)
}

func TestCreateTransaction_Success(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockUserRepo := new(MockUserRepository)
	mockBetRepo := new(MockBetRepository)
	useCase := NewTransactionUseCase(mockTransactionRepo, mockUserRepo, mockBetRepo)

	user := &entities.User{
		ID: 1,
	}

	transaction := &entities.Transaction{
		ID:     1,
		UserID: 1,
		Amount: 100,
		Type:   entities.TransactionTypeDeposit,
	}

	mockUserRepo.On("GetUserByID", uint(1)).Return(user, nil)
	mockTransactionRepo.On("CreateTransaction", transaction).Return(nil)

	err := useCase.CreateTransaction(transaction)

	assert.NoError(t, err)
	assert.Equal(t, entities.TransactionStatusPending, transaction.Status)
	mockTransactionRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestCreateTransaction_WithBet_Success(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockUserRepo := new(MockUserRepository)
	mockBetRepo := new(MockBetRepository)
	useCase := NewTransactionUseCase(mockTransactionRepo, mockUserRepo, mockBetRepo)

	user := &entities.User{
		ID: 1,
	}

	bet := &entities.Bet{
		ID:     1,
		UserID: 1,
	}

	betID := uint(1)
	transaction := &entities.Transaction{
		ID:     1,
		UserID: 1,
		Amount: 100,
		Type:   entities.TransactionTypeBet,
		BetID:  &betID,
	}

	mockUserRepo.On("GetUserByID", uint(1)).Return(user, nil)
	mockBetRepo.On("GetBetByID", uint(1)).Return(bet, nil)
	mockTransactionRepo.On("CreateTransaction", transaction).Return(nil)

	err := useCase.CreateTransaction(transaction)

	assert.NoError(t, err)
	assert.Equal(t, entities.TransactionStatusPending, transaction.Status)
	mockTransactionRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockBetRepo.AssertExpectations(t)
}

func TestCreateTransaction_WithBet_NotBelongToUser(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockUserRepo := new(MockUserRepository)
	mockBetRepo := new(MockBetRepository)
	useCase := NewTransactionUseCase(mockTransactionRepo, mockUserRepo, mockBetRepo)

	user := &entities.User{
		ID: 1,
	}

	bet := &entities.Bet{
		ID:     1,
		UserID: 2,
	}

	betID := uint(1)
	transaction := &entities.Transaction{
		ID:     1,
		UserID: 1,
		Amount: 100,
		Type:   entities.TransactionTypeBet,
		BetID:  &betID,
	}

	mockUserRepo.On("GetUserByID", uint(1)).Return(user, nil)
	mockBetRepo.On("GetBetByID", uint(1)).Return(bet, nil)

	err := useCase.CreateTransaction(transaction)

	assert.Error(t, err)
	assert.Equal(t, "bet does not belong to user", err.Error())
	mockUserRepo.AssertExpectations(t)
	mockBetRepo.AssertExpectations(t)
}

func TestGetTransactionByID_Success(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockUserRepo := new(MockUserRepository)
	mockBetRepo := new(MockBetRepository)
	useCase := NewTransactionUseCase(mockTransactionRepo, mockUserRepo, mockBetRepo)

	transaction := &entities.Transaction{
		ID:     1,
		UserID: 1,
		Amount: 100,
		Type:   entities.TransactionTypeDeposit,
	}

	mockTransactionRepo.On("GetTransactionByID", uint(1)).Return(transaction, nil)

	result, err := useCase.GetTransactionByID(1)

	assert.NoError(t, err)
	assert.Equal(t, transaction, result)
	mockTransactionRepo.AssertExpectations(t)
}

func TestGetTransactions_Success(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockUserRepo := new(MockUserRepository)
	mockBetRepo := new(MockBetRepository)
	useCase := NewTransactionUseCase(mockTransactionRepo, mockUserRepo, mockBetRepo)

	transactions := []*entities.Transaction{
		{
			ID:     1,
			UserID: 1,
			Amount: 100,
			Type:   entities.TransactionTypeDeposit,
		},
		{
			ID:     2,
			UserID: 2,
			Amount: 200,
			Type:   entities.TransactionTypeWithdrawal,
		},
	}

	mockTransactionRepo.On("GetTransactions").Return(transactions, nil)

	result, err := useCase.GetTransactions()

	assert.NoError(t, err)
	assert.Equal(t, transactions, result)
	mockTransactionRepo.AssertExpectations(t)
}

func TestGetTransactionsByUserID_Success(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockUserRepo := new(MockUserRepository)
	mockBetRepo := new(MockBetRepository)
	useCase := NewTransactionUseCase(mockTransactionRepo, mockUserRepo, mockBetRepo)

	transactions := []*entities.Transaction{
		{
			ID:     1,
			UserID: 1,
			Amount: 100,
			Type:   entities.TransactionTypeDeposit,
		},
		{
			ID:     2,
			UserID: 1,
			Amount: 200,
			Type:   entities.TransactionTypeWithdrawal,
		},
	}

	mockTransactionRepo.On("GetTransactionsByUserID", uint(1)).Return(transactions, nil)

	result, err := useCase.GetTransactionsByUserID(1)

	assert.NoError(t, err)
	assert.Equal(t, transactions, result)
	mockTransactionRepo.AssertExpectations(t)
}

func TestGetTransactionsByBetID_Success(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockUserRepo := new(MockUserRepository)
	mockBetRepo := new(MockBetRepository)
	useCase := NewTransactionUseCase(mockTransactionRepo, mockUserRepo, mockBetRepo)

	transactions := []*entities.Transaction{
		{
			ID:     1,
			UserID: 1,
			BetID:  &[]uint{1}[0],
			Amount: 100,
			Type:   entities.TransactionTypeBet,
		},
	}

	mockTransactionRepo.On("GetTransactionsByBetID", uint(1)).Return(transactions, nil)

	result, err := useCase.GetTransactionsByBetID(1)

	assert.NoError(t, err)
	assert.Equal(t, transactions, result)
	mockTransactionRepo.AssertExpectations(t)
}

func TestUpdateTransactionStatus_Success(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockUserRepo := new(MockUserRepository)
	mockBetRepo := new(MockBetRepository)
	useCase := NewTransactionUseCase(mockTransactionRepo, mockUserRepo, mockBetRepo)

	transaction := &entities.Transaction{
		ID:     1,
		Status: entities.TransactionStatusPending,
	}

	mockTransactionRepo.On("GetTransactionByID", uint(1)).Return(transaction, nil)
	mockTransactionRepo.On("UpdateTransactionStatus", uint(1), entities.TransactionStatusCompleted).Return(nil)

	err := useCase.UpdateTransactionStatus(1, entities.TransactionStatusCompleted)

	assert.NoError(t, err)
	mockTransactionRepo.AssertExpectations(t)
}

func TestUpdateTransactionStatus_NotPending(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockUserRepo := new(MockUserRepository)
	mockBetRepo := new(MockBetRepository)
	useCase := NewTransactionUseCase(mockTransactionRepo, mockUserRepo, mockBetRepo)

	transaction := &entities.Transaction{
		ID:     1,
		Status: entities.TransactionStatusCompleted,
	}

	mockTransactionRepo.On("GetTransactionByID", uint(1)).Return(transaction, nil)

	err := useCase.UpdateTransactionStatus(1, entities.TransactionStatusFailed)

	assert.Error(t, err)
	assert.Equal(t, "can only update status of pending transactions", err.Error())
	mockTransactionRepo.AssertExpectations(t)
}

func TestGetUserTransactionHistory_Success(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockUserRepo := new(MockUserRepository)
	mockBetRepo := new(MockBetRepository)
	useCase := NewTransactionUseCase(mockTransactionRepo, mockUserRepo, mockBetRepo)

	transactions := []*entities.Transaction{
		{
			ID:     1,
			UserID: 1,
			Amount: 100,
			Type:   entities.TransactionTypeDeposit,
		},
		{
			ID:     2,
			UserID: 1,
			Amount: 200,
			Type:   entities.TransactionTypeWithdrawal,
		},
	}

	mockTransactionRepo.On("GetUserTransactionHistory", uint(1), 10, 0).Return(transactions, nil)

	result, err := useCase.GetUserTransactionHistory(1, 10, 0)

	assert.NoError(t, err)
	assert.Equal(t, transactions, result)
	mockTransactionRepo.AssertExpectations(t)
}

func TestGetTransactionByReferenceID_Success(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockUserRepo := new(MockUserRepository)
	mockBetRepo := new(MockBetRepository)
	useCase := NewTransactionUseCase(mockTransactionRepo, mockUserRepo, mockBetRepo)

	transaction := &entities.Transaction{
		ID:          1,
		UserID:      1,
		Amount:      100,
		Type:        entities.TransactionTypeDeposit,
		ReferenceID: "ref123",
	}

	mockTransactionRepo.On("GetTransactionByReferenceID", "ref123").Return(transaction, nil)

	result, err := useCase.GetTransactionByReferenceID("ref123")

	assert.NoError(t, err)
	assert.Equal(t, transaction, result)
	mockTransactionRepo.AssertExpectations(t)
}

func TestProcessDeposit_Success(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockUserRepo := new(MockUserRepository)
	mockBetRepo := new(MockBetRepository)
	useCase := NewTransactionUseCase(mockTransactionRepo, mockUserRepo, mockBetRepo)

	transaction := &entities.Transaction{
		ID:     1,
		UserID: 1,
		Amount: 100,
		Type:   entities.TransactionTypeDeposit,
		Status: entities.TransactionStatusPending,
	}

	user := &entities.User{
		ID:      1,
		Balance: 1000,
	}

	updatedUser := &entities.User{
		ID:      1,
		Balance: 1100,
	}

	updatedTransaction := &entities.Transaction{
		ID:     1,
		UserID: 1,
		Amount: 100,
		Type:   entities.TransactionTypeDeposit,
		Status: entities.TransactionStatusCompleted,
	}

	mockTransactionRepo.On("GetTransactionByID", uint(1)).Return(transaction, nil)
	mockUserRepo.On("GetUserByID", uint(1)).Return(user, nil)
	mockUserRepo.On("UpdateUser", updatedUser).Return(nil)
	mockTransactionRepo.On("UpdateTransaction", updatedTransaction).Return(nil)

	err := useCase.ProcessDeposit(1)

	assert.NoError(t, err)
	mockTransactionRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestProcessDeposit_NotDeposit(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockUserRepo := new(MockUserRepository)
	mockBetRepo := new(MockBetRepository)
	useCase := NewTransactionUseCase(mockTransactionRepo, mockUserRepo, mockBetRepo)

	transaction := &entities.Transaction{
		ID:     1,
		UserID: 1,
		Amount: 100,
		Type:   entities.TransactionTypeWithdrawal,
		Status: entities.TransactionStatusPending,
	}

	mockTransactionRepo.On("GetTransactionByID", uint(1)).Return(transaction, nil)

	err := useCase.ProcessDeposit(1)

	assert.Error(t, err)
	assert.Equal(t, "transaction is not a deposit", err.Error())
	mockTransactionRepo.AssertExpectations(t)
}

func TestProcessDeposit_NotPending(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockUserRepo := new(MockUserRepository)
	mockBetRepo := new(MockBetRepository)
	useCase := NewTransactionUseCase(mockTransactionRepo, mockUserRepo, mockBetRepo)

	transaction := &entities.Transaction{
		ID:     1,
		UserID: 1,
		Amount: 100,
		Type:   entities.TransactionTypeDeposit,
		Status: entities.TransactionStatusCompleted,
	}

	mockTransactionRepo.On("GetTransactionByID", uint(1)).Return(transaction, nil)

	err := useCase.ProcessDeposit(1)

	assert.Error(t, err)
	assert.Equal(t, "transaction is not pending", err.Error())
	mockTransactionRepo.AssertExpectations(t)
}

func TestProcessWithdrawal_Success(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockUserRepo := new(MockUserRepository)
	mockBetRepo := new(MockBetRepository)
	useCase := NewTransactionUseCase(mockTransactionRepo, mockUserRepo, mockBetRepo)

	transaction := &entities.Transaction{
		ID:     1,
		UserID: 1,
		Amount: 100,
		Type:   entities.TransactionTypeWithdrawal,
		Status: entities.TransactionStatusPending,
	}

	user := &entities.User{
		ID:      1,
		Balance: 1000,
	}

	updatedUser := &entities.User{
		ID:      1,
		Balance: 900,
	}

	updatedTransaction := &entities.Transaction{
		ID:     1,
		UserID: 1,
		Amount: 100,
		Type:   entities.TransactionTypeWithdrawal,
		Status: entities.TransactionStatusCompleted,
	}

	mockTransactionRepo.On("GetTransactionByID", uint(1)).Return(transaction, nil)
	mockUserRepo.On("GetUserByID", uint(1)).Return(user, nil)
	mockUserRepo.On("UpdateUser", updatedUser).Return(nil)
	mockTransactionRepo.On("UpdateTransaction", updatedTransaction).Return(nil)

	err := useCase.ProcessWithdrawal(1)

	assert.NoError(t, err)
	mockTransactionRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestProcessWithdrawal_NotWithdrawal(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockUserRepo := new(MockUserRepository)
	mockBetRepo := new(MockBetRepository)
	useCase := NewTransactionUseCase(mockTransactionRepo, mockUserRepo, mockBetRepo)

	transaction := &entities.Transaction{
		ID:     1,
		UserID: 1,
		Amount: 100,
		Type:   entities.TransactionTypeDeposit,
		Status: entities.TransactionStatusPending,
	}

	mockTransactionRepo.On("GetTransactionByID", uint(1)).Return(transaction, nil)

	err := useCase.ProcessWithdrawal(1)

	assert.Error(t, err)
	assert.Equal(t, "transaction is not a withdrawal", err.Error())
	mockTransactionRepo.AssertExpectations(t)
}

func TestProcessWithdrawal_NotPending(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockUserRepo := new(MockUserRepository)
	mockBetRepo := new(MockBetRepository)
	useCase := NewTransactionUseCase(mockTransactionRepo, mockUserRepo, mockBetRepo)

	transaction := &entities.Transaction{
		ID:     1,
		UserID: 1,
		Amount: 100,
		Type:   entities.TransactionTypeWithdrawal,
		Status: entities.TransactionStatusCompleted,
	}

	mockTransactionRepo.On("GetTransactionByID", uint(1)).Return(transaction, nil)

	err := useCase.ProcessWithdrawal(1)

	assert.Error(t, err)
	assert.Equal(t, "transaction is not pending", err.Error())
	mockTransactionRepo.AssertExpectations(t)
}

func TestProcessWithdrawal_InsufficientBalance(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockUserRepo := new(MockUserRepository)
	mockBetRepo := new(MockBetRepository)
	useCase := NewTransactionUseCase(mockTransactionRepo, mockUserRepo, mockBetRepo)

	transaction := &entities.Transaction{
		ID:     1,
		UserID: 1,
		Amount: 100,
		Type:   entities.TransactionTypeWithdrawal,
		Status: entities.TransactionStatusPending,
	}

	user := &entities.User{
		ID:      1,
		Balance: 50,
	}

	mockTransactionRepo.On("GetTransactionByID", uint(1)).Return(transaction, nil)
	mockUserRepo.On("GetUserByID", uint(1)).Return(user, nil)

	err := useCase.ProcessWithdrawal(1)

	assert.Error(t, err)
	assert.Equal(t, "insufficient balance", err.Error())
	mockTransactionRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}
