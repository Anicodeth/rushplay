package domain

import "time"

type GameType string

const (
	GameTypeSlots     GameType = "SLOTS"
	GameTypeBlackjack GameType = "BLACKJACK"
	GameTypeRoulette  GameType = "ROULETTE"
	GameTypePoker     GameType = "POKER"
	GameTypeCraps     GameType = "CRAPS"
	GameTypeBaccarat  GameType = "BACCARAT"
)

type GameStatus string

const (
	GameStatusActive      GameStatus = "ACTIVE"
	GameStatusInactive    GameStatus = "INACTIVE"
	GameStatusMaintenance GameStatus = "MAINTENANCE"
)

type Game struct {
	ID          uint       `gorm:"primaryKey"`
	Name        string     `gorm:"type:varchar(100);not null"`
	Type        GameType   `gorm:"type:varchar(20);not null"`
	Status      GameStatus `gorm:"type:varchar(20);not null;default:'ACTIVE'"`
	MinBet      float64    `gorm:"type:decimal(18,2);not null"`
	MaxBet      float64    `gorm:"type:decimal(18,2);not null"`
	Description string     `gorm:"type:text"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`

	Bets []*Bet `gorm:"foreignKey:GameID;references:ID"`
}

func (Game) TableName() string {
	return "games"
}
