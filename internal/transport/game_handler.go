package transport

import (
	"context"
	"rushplay/api/generated/proto/gamepb"
	iusecases "rushplay/internal/domain/contracts/usecase"
	entities "rushplay/internal/domain/entities"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type GameHandler struct {
	gamepb.UnimplementedGameServiceServer
	useCase iusecases.IGameUseCase
}

func NewGameHandler(useCase iusecases.IGameUseCase) *GameHandler {
	return &GameHandler{useCase: useCase}
}

func (h *GameHandler) CreateGame(ctx context.Context, req *gamepb.CreateGameRequest) (*gamepb.CreateGameResponse, error) {
	game := &entities.Game{
		Name:        req.Name,
		Description: req.Description,
		Type:        entities.GameType(req.Type),
		MinBet:      req.MinBet,
		MaxBet:      req.MaxBet,
	}

	err := h.useCase.CreateGame(game)
	if err != nil {
		return nil, err
	}

	gamePb := &gamepb.Game{
		Id:          uint64(game.ID),
		Name:        game.Name,
		Description: game.Description,
		Type:        string(game.Type),
		Status:      string(game.Status),
		MinBet:      game.MinBet,
		MaxBet:      game.MaxBet,
		CreatedAt:   timestamppb.New(game.CreatedAt),
		UpdatedAt:   timestamppb.New(game.UpdatedAt),
	}

	return &gamepb.CreateGameResponse{Game: gamePb}, nil
}

func (h *GameHandler) GetGame(ctx context.Context, req *gamepb.GetGameRequest) (*gamepb.GetGameResponse, error) {
	game, err := h.useCase.GetGameByID(uint(req.Id))
	if err != nil {
		return nil, err
	}

	gamePb := &gamepb.Game{
		Id:          uint64(game.ID),
		Name:        game.Name,
		Description: game.Description,
		Type:        string(game.Type),
		Status:      string(game.Status),
		MinBet:      game.MinBet,
		MaxBet:      game.MaxBet,
		CreatedAt:   timestamppb.New(game.CreatedAt),
		UpdatedAt:   timestamppb.New(game.UpdatedAt),
	}

	return &gamepb.GetGameResponse{Game: gamePb}, nil
}

func (h *GameHandler) GetGames(ctx context.Context, req *gamepb.GetGamesRequest) (*gamepb.GetGamesResponse, error) {
	games, err := h.useCase.GetGames()
	if err != nil {
		return nil, err
	}

	if games == nil {
		return &gamepb.GetGamesResponse{
			Games: []*gamepb.Game{},
		}, nil
	}

	var gamesPb []*gamepb.Game
	for _, game := range games {
		gamesPb = append(gamesPb, &gamepb.Game{
			Id:          uint64(game.ID),
			Name:        game.Name,
			Description: game.Description,
			Type:        string(game.Type),
			Status:      string(game.Status),
			MinBet:      game.MinBet,
			MaxBet:      game.MaxBet,
			CreatedAt:   timestamppb.New(game.CreatedAt),
			UpdatedAt:   timestamppb.New(game.UpdatedAt),
		})
	}

	return &gamepb.GetGamesResponse{
		Games: gamesPb,
	}, nil
}

func (h *GameHandler) GetGamesByType(ctx context.Context, req *gamepb.GetGamesByTypeRequest) (*gamepb.GetGamesByTypeResponse, error) {
	games, err := h.useCase.GetGamesByType(entities.GameType(req.Type))
	if err != nil {
		return nil, err
	}

	if games == nil {
		return &gamepb.GetGamesByTypeResponse{
			Games: []*gamepb.Game{},
		}, nil
	}

	var gamesPb []*gamepb.Game
	for _, game := range games {
		gamesPb = append(gamesPb, &gamepb.Game{
			Id:          uint64(game.ID),
			Name:        game.Name,
			Description: game.Description,
			Type:        string(game.Type),
			Status:      string(game.Status),
			MinBet:      game.MinBet,
			MaxBet:      game.MaxBet,
			CreatedAt:   timestamppb.New(game.CreatedAt),
			UpdatedAt:   timestamppb.New(game.UpdatedAt),
		})
	}

	return &gamepb.GetGamesByTypeResponse{
		Games: gamesPb,
	}, nil
}

func (h *GameHandler) GetActiveGames(ctx context.Context, req *gamepb.GetActiveGamesRequest) (*gamepb.GetActiveGamesResponse, error) {
	games, err := h.useCase.GetActiveGames()
	if err != nil {
		return nil, err
	}

	if games == nil {
		return &gamepb.GetActiveGamesResponse{
			Games: []*gamepb.Game{},
		}, nil
	}

	var gamesPb []*gamepb.Game
	for _, game := range games {
		gamesPb = append(gamesPb, &gamepb.Game{
			Id:          uint64(game.ID),
			Name:        game.Name,
			Description: game.Description,
			Type:        string(game.Type),
			Status:      string(game.Status),
			MinBet:      game.MinBet,
			MaxBet:      game.MaxBet,
			CreatedAt:   timestamppb.New(game.CreatedAt),
			UpdatedAt:   timestamppb.New(game.UpdatedAt),
		})
	}

	return &gamepb.GetActiveGamesResponse{
		Games: gamesPb,
	}, nil
}

func (h *GameHandler) UpdateGame(ctx context.Context, req *gamepb.UpdateGameRequest) (*gamepb.UpdateGameResponse, error) {
	game := &entities.Game{
		ID:          uint(req.Id),
		Name:        req.Name,
		Description: req.Description,
		Type:        entities.GameType(req.Type),
		MinBet:      req.MinBet,
		MaxBet:      req.MaxBet,
	}

	err := h.useCase.UpdateGame(game)
	if err != nil {
		return nil, err
	}

	gamePb := &gamepb.Game{
		Id:          uint64(game.ID),
		Name:        game.Name,
		Description: game.Description,
		Type:        string(game.Type),
		Status:      string(game.Status),
		MinBet:      game.MinBet,
		MaxBet:      game.MaxBet,
		CreatedAt:   timestamppb.New(game.CreatedAt),
		UpdatedAt:   timestamppb.New(game.UpdatedAt),
	}

	return &gamepb.UpdateGameResponse{Game: gamePb}, nil
}

func (h *GameHandler) DeleteGame(ctx context.Context, req *gamepb.DeleteGameRequest) (*gamepb.DeleteGameResponse, error) {
	err := h.useCase.DeleteGame(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &gamepb.DeleteGameResponse{}, nil
}

func (h *GameHandler) UpdateGameStatus(ctx context.Context, req *gamepb.UpdateGameStatusRequest) (*gamepb.UpdateGameStatusResponse, error) {
	err := h.useCase.UpdateGameStatus(uint(req.Id), entities.GameStatus(req.Status))
	if err != nil {
		return nil, err
	}

	game, err := h.useCase.GetGameByID(uint(req.Id))
	if err != nil {
		return nil, err
	}

	gamePb := &gamepb.Game{
		Id:          uint64(game.ID),
		Name:        game.Name,
		Description: game.Description,
		Type:        string(game.Type),
		Status:      string(game.Status),
		MinBet:      game.MinBet,
		MaxBet:      game.MaxBet,
		CreatedAt:   timestamppb.New(game.CreatedAt),
		UpdatedAt:   timestamppb.New(game.UpdatedAt),
	}

	return &gamepb.UpdateGameStatusResponse{Game: gamePb}, nil
}
