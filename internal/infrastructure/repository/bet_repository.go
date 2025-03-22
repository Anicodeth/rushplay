package infrastructure

import (
	irepositories "rushplay/internal/domain/contracts/repository"
	entities "rushplay/internal/domain/entities"

	"gorm.io/gorm"
)

type BetRepository struct {
	db *gorm.DB
}

func NewBetRepository(db *gorm.DB) irepositories.IBetRepository {
	return &BetRepository{db: db}
}

func (r *BetRepository) CreateBet(bet *entities.Bet) error {
	return r.db.Create(bet).Error
}

func (r *BetRepository) GetBetByID(id uint) (*entities.Bet, error) {
	bet := new(entities.Bet)
	err := r.db.Where("id = ?", id).First(bet).Error
	return bet, err
}

func (r *BetRepository) GetBets() ([]*entities.Bet, error) {
	var bets []*entities.Bet
	err := r.db.Find(&bets).Error
	return bets, err
}

func (r *BetRepository) GetBetsByUserID(userID uint) ([]*entities.Bet, error) {
	var bets []*entities.Bet
	err := r.db.Where("user_id = ?", userID).Find(&bets).Error
	return bets, err
}

func (r *BetRepository) GetBetsByGameID(gameID uint) ([]*entities.Bet, error) {
	var bets []*entities.Bet
	err := r.db.Where("game_id = ?", gameID).Find(&bets).Error
	return bets, err
}

func (r *BetRepository) UpdateBet(bet *entities.Bet) error {
	return r.db.Save(bet).Error
}

func (r *BetRepository) DeleteBet(id uint) error {
	return r.db.Delete(&entities.Bet{}, id).Error
}

func (r *BetRepository) UpdateBetStatus(id uint, status entities.BetStatus) error {
	return r.db.Model(&entities.Bet{}).Where("id = ?", id).Update("status", status).Error
}

func (r *BetRepository) GetUserTotalBets(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&entities.Bet{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

func (r *BetRepository) GetUserTotalWinnings(userID uint) (float64, error) {
	var total float64
	err := r.db.Model(&entities.Bet{}).
		Where("user_id = ? AND status = ?", userID, entities.BetStatusWon).
		Select("COALESCE(SUM(winnings), 0)").
		Scan(&total).Error
	return total, err
}

func (r *BetRepository) GetUserTotalBetsByGameID(userID uint, gameID uint) (int64, error) {
	var count int64
	err := r.db.Model(&entities.Bet{}).
		Where("user_id = ? AND game_id = ?", userID, gameID).
		Count(&count).Error
	return count, err
}

func (r *BetRepository) GetUserTotalWinningsByGameID(userID uint, gameID uint) (float64, error) {
	var total float64
	err := r.db.Model(&entities.Bet{}).
		Where("user_id = ? AND game_id = ? AND status = ?", userID, gameID, entities.BetStatusWon).
		Select("COALESCE(SUM(winnings), 0)").
		Scan(&total).Error
	return total, err
}
