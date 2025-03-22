package domain

import "time"

type TransactionType string

const (
	TransactionTypeDeposit    TransactionType = "DEPOSIT"
	TransactionTypeWithdrawal TransactionType = "WITHDRAWAL"
	TransactionTypeBet        TransactionType = "BET"
	TransactionTypeWin        TransactionType = "WIN"
	TransactionTypeBonus      TransactionType = "BONUS"
	TransactionTypeRefund     TransactionType = "REFUND"
)

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "PENDING"
	TransactionStatusCompleted TransactionStatus = "COMPLETED"
	TransactionStatusFailed    TransactionStatus = "FAILED"
	TransactionStatusCancelled TransactionStatus = "CANCELLED"
)

type Transaction struct {
	ID          uint             `gorm:"primaryKey;autoIncrement"`
	UserID      uint             `gorm:"not null"`
	Type        TransactionType  `gorm:"type:varchar(20);not null"`
	Status      TransactionStatus `gorm:"type:varchar(20);not null;default:'PENDING'"`
	Amount      float64          `gorm:"type:decimal(18,2);not null"`
	Balance     float64          `gorm:"type:decimal(18,2);not null"` // User's balance after transaction
	ReferenceID string           `gorm:"type:varchar(255)"` // External reference (e.g., payment provider ID)
	Description string           `gorm:"type:text"`
	CreatedAt   time.Time        `gorm:"autoCreateTime"`
	UpdatedAt   time.Time        `gorm:"autoUpdateTime"`
	BetID       *uint            `gorm:"index"`

	User        *User            `gorm:"foreignKey:UserID;references:ID"`
	Bet         *Bet             `gorm:"foreignKey:BetID;references:ID"`
}

func (Transaction) TableName() string {
	return "transactions"
}