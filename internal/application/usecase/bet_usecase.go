package application

import (
	"errors"
	irepositories "rushplay/internal/domain/contracts/repository"
	entities "rushplay/internal/domain/entities"
)

type BetUseCase struct {
	betRepo  irepositories.IBetRepository
	gameRepo irepositories.IGameRepository
	userRepo irepositories.IUserRepository
}

func NewBetUseCase(betRepo irepositories.IBetRepository, gameRepo irepositories.IGameRepository, userRepo irepositories.IUserRepository) *BetUseCase {
	return &BetUseCase{
		betRepo:  betRepo,
		gameRepo: gameRepo,
		userRepo: userRepo,
	}
}

func (u *BetUseCase) CreateBet(bet *entities.Bet) error {
	game, err := u.gameRepo.GetGameByID(bet.GameID)
	if err != nil {
		return err
	}

	if game.Status != entities.GameStatusActive {
		return errors.New("game is not active")
	}

	if bet.Amount < game.MinBet || bet.Amount > game.MaxBet {
		return errors.New("bet amount is outside game limits")
	}

	user, err := u.userRepo.GetUserByID(bet.UserID)
	if err != nil {
		return err
	}

	if user.Balance < bet.Amount {
		return errors.New("insufficient balance")
	}

	bet.Status = entities.BetStatusPending

	return u.betRepo.CreateBet(bet)
}

func (u *BetUseCase) GetBetByID(id uint) (*entities.Bet, error) {
	return u.betRepo.GetBetByID(id)
}

func (u *BetUseCase) GetBets() ([]*entities.Bet, error) {
	return u.betRepo.GetBets()
}

func (u *BetUseCase) GetBetsByUserID(userID uint) ([]*entities.Bet, error) {
	return u.betRepo.GetBetsByUserID(userID)
}

func (u *BetUseCase) UpdateBetStatus(id uint, status entities.BetStatus) error {
	bet, err := u.betRepo.GetBetByID(id)
	if err != nil {
		return err
	}

	if bet.Status != entities.BetStatusPending {
		return errors.New("can only update status of pending bets")
	}

	bet.Status = status
	return u.betRepo.UpdateBet(bet)
}

func (u *BetUseCase) ProcessBetResult(betID uint, won bool) error {
	bet, err := u.betRepo.GetBetByID(betID)
	if err != nil {
		return err
	}

	if bet.Status != entities.BetStatusPending {
		return errors.New("can only process pending bets")
	}

	user, err := u.userRepo.GetUserByID(bet.UserID)
	if err != nil {
		return err
	}

	if won {
		bet.Status = entities.BetStatusWon
		bet.Winnings = bet.Amount * 2
		user.Balance += bet.Winnings
	} else {
		bet.Status = entities.BetStatusLost
		bet.Winnings = 0
	}

	if err := u.betRepo.UpdateBet(bet); err != nil {
		return err
	}
	return u.userRepo.UpdateUser(user)
}
