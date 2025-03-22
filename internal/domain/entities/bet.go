package domain

import "time"

type BetStatus string

const (
	BetStatusPending   BetStatus = "PENDING"
	BetStatusWon       BetStatus = "WON"
	BetStatusLost      BetStatus = "LOST"
	BetStatusCancelled BetStatus = "CANCELLED"
)

type Bet struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uint      `gorm:"not null"`
	GameID    uint      `gorm:"not null"`
	Amount    float64   `gorm:"type:decimal(18,2);not null"`
	Status    BetStatus `gorm:"type:varchar(20);not null;default:'PENDING'"`
	Winnings  float64   `gorm:"type:decimal(18,2);default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User *User `gorm:"foreignKey:UserID;references:ID"`
	Game *Game `gorm:"foreignKey:GameID;references:ID"`
}

func (Bet) TableName() string {
	return "bets"
}
