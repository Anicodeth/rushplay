package domain

import (
	domain "rushplay/internal/domain/entities"
)

type IGameUseCase interface {
	CreateGame(game *domain.Game) error
	GetGameByID(id uint) (*domain.Game, error)
	GetGames() ([]*domain.Game, error)
	GetGamesByType(gameType domain.GameType) ([]*domain.Game, error)
	GetActiveGames() ([]*domain.Game, error)
	UpdateGame(game *domain.Game) error
	DeleteGame(id uint) error
	UpdateGameStatus(id uint, status domain.GameStatus) error
	ValidateGameBet(gameID uint, amount float64) error
	GetGameStatistics(gameID uint) (*GameStatistics, error)
}

type GameStatistics struct {
	TotalBets     int64   `json:"total_bets"`
	TotalWinnings float64 `json:"total_winnings"`
	AverageBet    float64 `json:"average_bet"`
	WinRate       float64 `json:"win_rate"`
}
