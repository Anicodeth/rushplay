package domain

import (
	domain "rushplay/internal/domain/entities"
)

type ITransactionRepository interface {
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
}
