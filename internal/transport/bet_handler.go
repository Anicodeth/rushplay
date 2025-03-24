package transport

import (
	"context"
	"rushplay/api/generated/proto/betpb"
	iusecases "rushplay/internal/domain/contracts/usecase"
	entities "rushplay/internal/domain/entities"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type BetHandler struct {
	betpb.UnimplementedBetServiceServer
	useCase iusecases.IBetUseCase
}

func NewBetHandler(useCase iusecases.IBetUseCase) *BetHandler {
	return &BetHandler{useCase: useCase}
}

func (h *BetHandler) CreateBet(ctx context.Context, req *betpb.CreateBetRequest) (*betpb.CreateBetResponse, error) {
	bet := &entities.Bet{
		UserID: uint(req.UserId),
		GameID: uint(req.GameId),
		Amount: req.Amount,
	}

	err := h.useCase.CreateBet(bet)
	if err != nil {
		return nil, err
	}

	betPb := &betpb.Bet{
		Id:           uint64(bet.ID),
		UserId:       uint64(bet.UserID),
		GameId:       uint64(bet.GameID),
		Amount:       bet.Amount,
		Status:       string(bet.Status),
		PotentialWin: bet.Winnings,
		CreatedAt:    timestamppb.New(bet.CreatedAt),
		UpdatedAt:    timestamppb.New(bet.UpdatedAt),
	}

	return &betpb.CreateBetResponse{Bet: betPb}, nil
}

func (h *BetHandler) GetBet(ctx context.Context, req *betpb.GetBetRequest) (*betpb.GetBetResponse, error) {
	bet, err := h.useCase.GetBetByID(uint(req.Id))
	if err != nil {
		return nil, err
	}

	betPb := &betpb.Bet{
		Id:           uint64(bet.ID),
		UserId:       uint64(bet.UserID),
		GameId:       uint64(bet.GameID),
		Amount:       bet.Amount,
		Status:       string(bet.Status),
		PotentialWin: bet.Winnings,
		CreatedAt:    timestamppb.New(bet.CreatedAt),
		UpdatedAt:    timestamppb.New(bet.UpdatedAt),
	}

	return &betpb.GetBetResponse{Bet: betPb}, nil
}

func (h *BetHandler) GetBets(ctx context.Context, req *betpb.GetBetsRequest) (*betpb.GetBetsResponse, error) {
	bets, err := h.useCase.GetBets()
	if err != nil {
		return nil, err
	}

	betsPb := make([]*betpb.Bet, len(bets))
	for i, bet := range bets {
		betsPb[i] = &betpb.Bet{
			Id:           uint64(bet.ID),
			UserId:       uint64(bet.UserID),
			GameId:       uint64(bet.GameID),
			Amount:       bet.Amount,
			Status:       string(bet.Status),
			PotentialWin: bet.Winnings,
			CreatedAt:    timestamppb.New(bet.CreatedAt),
			UpdatedAt:    timestamppb.New(bet.UpdatedAt),
		}
	}

	return &betpb.GetBetsResponse{Bets: betsPb}, nil
}

func (h *BetHandler) GetBetsByUserID(ctx context.Context, req *betpb.GetBetsByUserIDRequest) (*betpb.GetBetsByUserIDResponse, error) {
	bets, err := h.useCase.GetBetsByUserID(uint(req.UserId))
	if err != nil {
		return nil, err
	}

	betsPb := make([]*betpb.Bet, len(bets))
	for i, bet := range bets {
		betsPb[i] = &betpb.Bet{
			Id:           uint64(bet.ID),
			UserId:       uint64(bet.UserID),
			GameId:       uint64(bet.GameID),
			Amount:       bet.Amount,
			Status:       string(bet.Status),
			PotentialWin: bet.Winnings,
			CreatedAt:    timestamppb.New(bet.CreatedAt),
			UpdatedAt:    timestamppb.New(bet.UpdatedAt),
		}
	}

	return &betpb.GetBetsByUserIDResponse{Bets: betsPb}, nil
}

func (h *BetHandler) GetBetsByGameID(ctx context.Context, req *betpb.GetBetsByGameIDRequest) (*betpb.GetBetsByGameIDResponse, error) {
	bets, err := h.useCase.GetBetsByGameID(uint(req.GameId))
	if err != nil {
		return nil, err
	}

	betsPb := make([]*betpb.Bet, len(bets))
	for i, bet := range bets {
		betsPb[i] = &betpb.Bet{
			Id:           uint64(bet.ID),
			UserId:       uint64(bet.UserID),
			GameId:       uint64(bet.GameID),
			Amount:       bet.Amount,
			Status:       string(bet.Status),
			PotentialWin: bet.Winnings,
			CreatedAt:    timestamppb.New(bet.CreatedAt),
			UpdatedAt:    timestamppb.New(bet.UpdatedAt),
		}
	}

	return &betpb.GetBetsByGameIDResponse{Bets: betsPb}, nil
}

func (h *BetHandler) UpdateBetStatus(ctx context.Context, req *betpb.UpdateBetStatusRequest) (*betpb.UpdateBetStatusResponse, error) {
	err := h.useCase.UpdateBetStatus(uint(req.Id), entities.BetStatus(req.Status))
	if err != nil {
		return nil, err
	}

	bet, err := h.useCase.GetBetByID(uint(req.Id))
	if err != nil {
		return nil, err
	}

	betPb := &betpb.Bet{
		Id:           uint64(bet.ID),
		UserId:       uint64(bet.UserID),
		GameId:       uint64(bet.GameID),
		Amount:       bet.Amount,
		Status:       string(bet.Status),
		PotentialWin: bet.Winnings,
		CreatedAt:    timestamppb.New(bet.CreatedAt),
		UpdatedAt:    timestamppb.New(bet.UpdatedAt),
	}

	return &betpb.UpdateBetStatusResponse{Bet: betPb}, nil
}

func (h *BetHandler) ProcessBetResult(ctx context.Context, req *betpb.ProcessBetResultRequest) (*betpb.ProcessBetResultResponse, error) {
	result := iusecases.BetResult{
		Won:     req.Result == "WON",
		Details: req.Result,
	}
	err := h.useCase.ProcessBetResult(uint(req.Id), result)
	if err != nil {
		return nil, err
	}

	bet, err := h.useCase.GetBetByID(uint(req.Id))
	if err != nil {
		return nil, err
	}

	betPb := &betpb.Bet{
		Id:           uint64(bet.ID),
		UserId:       uint64(bet.UserID),
		GameId:       uint64(bet.GameID),
		Amount:       bet.Amount,
		Status:       string(bet.Status),
		PotentialWin: bet.Winnings,
		CreatedAt:    timestamppb.New(bet.CreatedAt),
		UpdatedAt:    timestamppb.New(bet.UpdatedAt),
	}

	return &betpb.ProcessBetResultResponse{Bet: betPb}, nil
}

func (h *BetHandler) GetUserTotalBets(ctx context.Context, req *betpb.GetUserTotalBetsRequest) (*betpb.GetUserTotalBetsResponse, error) {
	total, err := h.useCase.GetUserTotalBets(uint(req.UserId))
	if err != nil {
		return nil, err
	}

	return &betpb.GetUserTotalBetsResponse{Total: int32(total)}, nil
}

func (h *BetHandler) GetUserTotalBetsByGameID(ctx context.Context, req *betpb.GetUserTotalBetsByGameIDRequest) (*betpb.GetUserTotalBetsByGameIDResponse, error) {
	total, err := h.useCase.GetUserTotalBetsByGameID(uint(req.UserId), uint(req.GameId))
	if err != nil {
		return nil, err
	}

	return &betpb.GetUserTotalBetsByGameIDResponse{Total: int32(total)}, nil
}

func (h *BetHandler) GetUserTotalWinnings(ctx context.Context, req *betpb.GetUserTotalWinningsRequest) (*betpb.GetUserTotalWinningsResponse, error) {
	total, err := h.useCase.GetUserTotalWinnings(uint(req.UserId))
	if err != nil {
		return nil, err
	}

	return &betpb.GetUserTotalWinningsResponse{Total: total}, nil
}
