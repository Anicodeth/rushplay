package domain

import (
	domain "rushplay/internal/domain/entities"
)

type ITransactionUseCase interface {
	CreateTransaction(transaction *domain.Transaction) error
	GetTransactionByID(id uint) (*domain.Transaction, error)
	GetTransactions() ([]*domain.Transaction, error)
	GetTransactionsByUserID(userID uint) ([]*domain.Transaction, error)
	GetTransactionsByBetID(betID uint) ([]*domain.Transaction, error)
	UpdateTransaction(transaction *domain.Transaction) error
	DeleteTransaction(id uint) error
	UpdateTransactionStatus(id uint, status domain.TransactionStatus) error
	GetUserTransactionHistory(userID uint, limit, offset int) ([]*domain.Transaction, error)
	GetTransactionByReferenceID(referenceID string) (*domain.Transaction, error)
	ProcessDeposit(userID uint, amount float64, referenceID string) (*domain.Transaction, error)
	ProcessWithdrawal(userID uint, amount float64, referenceID string) (*domain.Transaction, error)
	ProcessBetTransaction(userID uint, betID uint, amount float64) (*domain.Transaction, error)
	ProcessWinTransaction(userID uint, betID uint, amount float64) (*domain.Transaction, error)
	GetUserBalance(userID uint) (float64, error)
	ValidateTransaction(transaction *domain.Transaction) error
}
