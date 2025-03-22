package domain

import (
	domain "rushplay/internal/domain/entities"
)

type IBetUseCase interface {
	CreateBet(bet *domain.Bet) error
	GetBetByID(id uint) (*domain.Bet, error)
	GetBets() ([]*domain.Bet, error)
	GetBetsByUserID(userID uint) ([]*domain.Bet, error)
	GetBetsByGameID(gameID uint) ([]*domain.Bet, error)
	UpdateBet(bet *domain.Bet) error
	DeleteBet(id uint) error
	UpdateBetStatus(id uint, status domain.BetStatus) error
	GetUserTotalBets(userID uint) (int64, error)
	GetUserTotalWinnings(userID uint) (float64, error)
	GetUserTotalBetsByGameID(userID uint, gameID uint) (int64, error)
	GetUserTotalWinningsByGameID(userID uint, gameID uint) (float64, error)
	PlaceBet(userID uint, gameID uint, amount float64, details string) (*domain.Bet, error)
	ProcessBetResult(betID uint, result BetResult) error
	GetUserBettingHistory(userID uint, limit, offset int) ([]*domain.Bet, error)
}

type BetResult struct {
	Won        bool    `json:"won"`
	Amount     float64 `json:"amount"`
	Details    string  `json:"details"`
	Multiplier float64 `json:"multiplier"`
}
