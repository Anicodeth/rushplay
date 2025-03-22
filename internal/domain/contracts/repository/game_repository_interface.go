package domain

import (
	domain "rushplay/internal/domain/entities"
)

type IGameRepository interface {
	CreateGame(game *domain.Game) error
	GetGameByID(id uint) (*domain.Game, error)
	GetGames() ([]*domain.Game, error)
	GetGamesByType(gameType domain.GameType) ([]*domain.Game, error)
	GetActiveGames() ([]*domain.Game, error)
	UpdateGame(game *domain.Game) error
	DeleteGame(id uint) error
	UpdateGameStatus(id uint, status domain.GameStatus) error
}