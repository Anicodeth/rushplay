package infrastructure

import (
	irepositories "rushplay/internal/domain/contracts/repository"
	entities "rushplay/internal/domain/entities"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) irepositories.ITransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(transaction *entities.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *TransactionRepository) GetTransactionByID(id uint) (*entities.Transaction, error) {
	transaction := new(entities.Transaction)
	err := r.db.Where("id = ?", id).First(transaction).Error
	return transaction, err
}

func (r *TransactionRepository) GetTransactionsByUserID(userID uint) ([]*entities.Transaction, error) {
	var transactions []*entities.Transaction
	err := r.db.Where("user_id = ?", userID).Find(&transactions).Error
	return transactions, err
}

func (r *TransactionRepository) GetTransactionsByBetID(betID uint) ([]*entities.Transaction, error) {
	var transactions []*entities.Transaction
	err := r.db.Where("bet_id = ?", betID).Find(&transactions).Error
	return transactions, err
}

func (r *TransactionRepository) UpdateTransaction(transaction *entities.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r *TransactionRepository) DeleteTransaction(id uint) error {
	return r.db.Delete(&entities.Transaction{}, id).Error
}

func (r *TransactionRepository) UpdateTransactionStatus(id uint, status entities.TransactionStatus) error {
	return r.db.Model(&entities.Transaction{}).Where("id = ?", id).Update("status", status).Error
}

func (r *TransactionRepository) GetUserTransactionHistory(userID uint, limit, offset int) ([]*entities.Transaction, error) {
	var transactions []*entities.Transaction
	err := r.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&transactions).Error
	return transactions, err
}

func (r *TransactionRepository) GetTransactionByReferenceID(referenceID string) (*entities.Transaction, error) {
	transaction := new(entities.Transaction)
	err := r.db.Where("reference_id = ?", referenceID).First(transaction).Error
	return transaction, err
}

func (r *TransactionRepository) GetTransactions() ([]*entities.Transaction, error) {
	var transactions []*entities.Transaction
	err := r.db.Find(&transactions).Error
	return transactions, err
}
