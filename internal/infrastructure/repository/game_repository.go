package infrastructure

import (
	irepositories "rushplay/internal/domain/contracts/repository"
	entities "rushplay/internal/domain/entities"

	"gorm.io/gorm"
)

type GameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) irepositories.IGameRepository {
	return &GameRepository{db: db}
}

func (r *GameRepository) CreateGame(game *entities.Game) error {
	return r.db.Create(game).Error
}

func (r *GameRepository) GetGameByID(id uint) (*entities.Game, error) {
	game := new(entities.Game)
	err := r.db.Where("id = ?", id).First(game).Error
	return game, err
}

func (r *GameRepository) GetGamesByType(gameType entities.GameType) ([]*entities.Game, error) {
	var games []*entities.Game
	err := r.db.Where("type = ?", gameType).Find(&games).Error
	return games, err
}

func (r *GameRepository) GetActiveGames() ([]*entities.Game, error) {
	var games []*entities.Game
	err := r.db.Where("status = ?", entities.GameStatusActive).Find(&games).Error
	return games, err
}

func (r *GameRepository) UpdateGame(game *entities.Game) error {
	return r.db.Save(game).Error
}

func (r *GameRepository) DeleteGame(id uint) error {
	return r.db.Delete(&entities.Game{}, id).Error
}

func (r *GameRepository) UpdateGameStatus(id uint, status entities.GameStatus) error {
	return r.db.Model(&entities.Game{}).Where("id = ?", id).Update("status", status).Error
}

func (r *GameRepository) GetGames() ([]*entities.Game, error) {
	var games []*entities.Game
	err := r.db.Find(&games).Error
	return games, err
}
