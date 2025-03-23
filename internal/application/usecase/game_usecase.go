package application

import (
	"errors"
	irepositories "rushplay/internal/domain/contracts/repository"
	entities "rushplay/internal/domain/entities"
)

type GameUseCase struct {
	gameRepo irepositories.IGameRepository
}

func NewGameUseCase(gameRepo irepositories.IGameRepository) *GameUseCase {
	return &GameUseCase{
		gameRepo: gameRepo,
	}
}

func (u *GameUseCase) CreateGame(game *entities.Game) error {
	if game.MinBet <= 0 || game.MaxBet <= 0 {
		return errors.New("bet amounts must be positive")
	}
	if game.MinBet > game.MaxBet {
		return errors.New("minimum bet cannot be greater than maximum bet")
	}
	return u.gameRepo.CreateGame(game)
}

func (u *GameUseCase) GetGameByID(id uint) (*entities.Game, error) {
	return u.gameRepo.GetGameByID(id)
}

func (u *GameUseCase) GetGames() ([]*entities.Game, error) {
	return u.gameRepo.GetGames()
}

func (u *GameUseCase) GetGamesByType(gameType entities.GameType) ([]*entities.Game, error) {
	return u.gameRepo.GetGamesByType(gameType)
}

func (u *GameUseCase) GetActiveGames() ([]*entities.Game, error) {
	return u.gameRepo.GetActiveGames()
}

func (u *GameUseCase) UpdateGame(game *entities.Game) error {
	if game.MinBet <= 0 || game.MaxBet <= 0 {
		return errors.New("bet amounts must be positive")
	}
	if game.MinBet > game.MaxBet {
		return errors.New("minimum bet cannot be greater than maximum bet")
	}
	return u.gameRepo.UpdateGame(game)
}

func (u *GameUseCase) DeleteGame(id uint) error {
	return u.gameRepo.DeleteGame(id)
}

func (u *GameUseCase) UpdateGameStatus(id uint, status entities.GameStatus) error {
	return u.gameRepo.UpdateGameStatus(id, status)
}
