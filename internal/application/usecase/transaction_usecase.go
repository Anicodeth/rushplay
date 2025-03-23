package application

import (
	"errors"
	irepositories "rushplay/internal/domain/contracts/repository"
	entities "rushplay/internal/domain/entities"
	"time"
)

type TransactionUseCase struct {
	transactionRepo irepositories.ITransactionRepository
	userRepo        irepositories.IUserRepository
	betRepo         irepositories.IBetRepository
}

func NewTransactionUseCase(transactionRepo irepositories.ITransactionRepository, userRepo irepositories.IUserRepository, betRepo irepositories.IBetRepository) *TransactionUseCase {
	return &TransactionUseCase{
		transactionRepo: transactionRepo,
		userRepo:        userRepo,
		betRepo:         betRepo,
	}
}

func (u *TransactionUseCase) CreateTransaction(transaction *entities.Transaction) error {
	_, err := u.userRepo.GetUserByID(transaction.UserID)
	if err != nil {
		return err
	}

	// If transaction is related to a bet, validate bet exists
	if transaction.BetID != nil {
		bet, err := u.betRepo.GetBetByID(*transaction.BetID)
		if err != nil {
			return err
		}
		if bet.UserID != transaction.UserID {
			return errors.New("bet does not belong to user")
		}
	}

	transaction.Status = entities.TransactionStatusPending
	transaction.CreatedAt = time.Now()

	return u.transactionRepo.CreateTransaction(transaction)
}

func (u *TransactionUseCase) GetTransactionByID(id uint) (*entities.Transaction, error) {
	return u.transactionRepo.GetTransactionByID(id)
}

func (u *TransactionUseCase) GetTransactions() ([]*entities.Transaction, error) {
	return u.transactionRepo.GetTransactions()
}

func (u *TransactionUseCase) GetTransactionsByUserID(userID uint) ([]*entities.Transaction, error) {
	return u.transactionRepo.GetTransactionsByUserID(userID)
}

func (u *TransactionUseCase) GetTransactionsByBetID(betID uint) ([]*entities.Transaction, error) {
	return u.transactionRepo.GetTransactionsByBetID(betID)
}

func (u *TransactionUseCase) UpdateTransactionStatus(id uint, status entities.TransactionStatus) error {
	transaction, err := u.transactionRepo.GetTransactionByID(id)
	if err != nil {
		return err
	}

	if transaction.Status != entities.TransactionStatusPending {
		return errors.New("can only update status of pending transactions")
	}

	transaction.Status = status
	return u.transactionRepo.UpdateTransactionStatus(id, status)
}

func (u *TransactionUseCase) GetUserTransactionHistory(userID uint, limit, offset int) ([]*entities.Transaction, error) {
	return u.transactionRepo.GetUserTransactionHistory(userID, limit, offset)
}

func (u *TransactionUseCase) GetTransactionByReferenceID(referenceID string) (*entities.Transaction, error) {
	return u.transactionRepo.GetTransactionByReferenceID(referenceID)
}

func (u *TransactionUseCase) ProcessDeposit(transactionID uint) error {
	transaction, err := u.transactionRepo.GetTransactionByID(transactionID)
	if err != nil {
		return err
	}

	if transaction.Type != entities.TransactionTypeDeposit {
		return errors.New("transaction is not a deposit")
	}

	if transaction.Status != entities.TransactionStatusPending {
		return errors.New("transaction is not pending")
	}

	user, err := u.userRepo.GetUserByID(transaction.UserID)
	if err != nil {
		return err
	}

	user.Balance += transaction.Amount
	transaction.Status = entities.TransactionStatusCompleted

	if err := u.userRepo.UpdateUser(user); err != nil {
		return err
	}
	return u.transactionRepo.UpdateTransaction(transaction)
}

func (u *TransactionUseCase) ProcessWithdrawal(transactionID uint) error {
	transaction, err := u.transactionRepo.GetTransactionByID(transactionID)
	if err != nil {
		return err
	}

	if transaction.Type != entities.TransactionTypeWithdrawal {
		return errors.New("transaction is not a withdrawal")
	}

	if transaction.Status != entities.TransactionStatusPending {
		return errors.New("transaction is not pending")
	}

	user, err := u.userRepo.GetUserByID(transaction.UserID)
	if err != nil {
		return err
	}

	if user.Balance < transaction.Amount {
		return errors.New("insufficient balance")
	}

	user.Balance -= transaction.Amount
	transaction.Status = entities.TransactionStatusCompleted

	if err := u.userRepo.UpdateUser(user); err != nil {
		return err
	}
	return u.transactionRepo.UpdateTransaction(transaction)
}
