package domain

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	FirstName    string    `gorm:"column:first_name;type:varchar(100);not null"`
	LastName     string    `gorm:"column:last_name;type:varchar(100);not null"`
	Email        string    `gorm:"uniqueIndex;type:varchar(255);not null"`
	PasswordHash string    `gorm:"column:password_hash;type:varchar(255);not null"`
	Balance      float64   `gorm:"type:decimal(18,2);default:0"`
	Role         string    `gorm:"type:varchar(20);not null;default:'player'"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`

	Bets         []Bet         `gorm:"foreignKey:UserID"`
	Transactions []Transaction `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}