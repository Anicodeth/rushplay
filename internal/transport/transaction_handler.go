package transport

import (
	"context"
	"rushplay/api/generated/proto/transactionpb"
	iusecases "rushplay/internal/domain/contracts/usecase"
	entities "rushplay/internal/domain/entities"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type TransactionHandler struct {
	transactionpb.UnimplementedTransactionServiceServer
	useCase iusecases.ITransactionUseCase
}

func NewTransactionHandler(useCase iusecases.ITransactionUseCase) *TransactionHandler {
	return &TransactionHandler{useCase: useCase}
}

func (h *TransactionHandler) CreateTransaction(ctx context.Context, req *transactionpb.CreateTransactionRequest) (*transactionpb.CreateTransactionResponse, error) {
	var betID *uint
	if req.BetId != nil {
		u := uint(*req.BetId)
		betID = &u
	}

	transaction := &entities.Transaction{
		UserID: uint(req.UserId),
		Type:   entities.TransactionType(req.Type),
		Amount: req.Amount,
		BetID:  betID,
	}

	err := h.useCase.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}

	transactionPb := &transactionpb.Transaction{
		Id:          uint64(transaction.ID),
		UserId:      uint64(transaction.UserID),
		Type:        string(transaction.Type),
		Status:      string(transaction.Status),
		Amount:      transaction.Amount,
		ReferenceId: transaction.ReferenceID,
		BetId:       uint64Ptr(transaction.BetID),
		CreatedAt:   timestamppb.New(transaction.CreatedAt),
		UpdatedAt:   timestamppb.New(transaction.UpdatedAt),
	}

	return &transactionpb.CreateTransactionResponse{Transaction: transactionPb}, nil
}

func (h *TransactionHandler) GetTransaction(ctx context.Context, req *transactionpb.GetTransactionRequest) (*transactionpb.GetTransactionResponse, error) {
	transaction, err := h.useCase.GetTransactionByID(uint(req.Id))
	if err != nil {
		return nil, err
	}

	transactionPb := &transactionpb.Transaction{
		Id:          uint64(transaction.ID),
		UserId:      uint64(transaction.UserID),
		Type:        string(transaction.Type),
		Status:      string(transaction.Status),
		Amount:      transaction.Amount,
		ReferenceId: transaction.ReferenceID,
		BetId:       uint64Ptr(transaction.BetID),
		CreatedAt:   timestamppb.New(transaction.CreatedAt),
		UpdatedAt:   timestamppb.New(transaction.UpdatedAt),
	}

	return &transactionpb.GetTransactionResponse{Transaction: transactionPb}, nil
}

func (h *TransactionHandler) GetTransactions(ctx context.Context, req *transactionpb.GetTransactionsRequest) (*transactionpb.GetTransactionsResponse, error) {
	transactions, err := h.useCase.GetTransactions()
	if err != nil {
		return nil, err
	}

	transactionsPb := make([]*transactionpb.Transaction, len(transactions))
	for i, transaction := range transactions {
		transactionsPb[i] = &transactionpb.Transaction{
			Id:          uint64(transaction.ID),
			UserId:      uint64(transaction.UserID),
			Type:        string(transaction.Type),
			Status:      string(transaction.Status),
			Amount:      transaction.Amount,
			ReferenceId: transaction.ReferenceID,
			BetId:       uint64Ptr(transaction.BetID),
			CreatedAt:   timestamppb.New(transaction.CreatedAt),
			UpdatedAt:   timestamppb.New(transaction.UpdatedAt),
		}
	}

	return &transactionpb.GetTransactionsResponse{Transactions: transactionsPb}, nil
}

func (h *TransactionHandler) GetTransactionsByUserID(ctx context.Context, req *transactionpb.GetTransactionsByUserIDRequest) (*transactionpb.GetTransactionsByUserIDResponse, error) {
	transactions, err := h.useCase.GetTransactionsByUserID(uint(req.UserId))
	if err != nil {
		return nil, err
	}

	transactionsPb := make([]*transactionpb.Transaction, len(transactions))
	for i, transaction := range transactions {
		transactionsPb[i] = &transactionpb.Transaction{
			Id:          uint64(transaction.ID),
			UserId:      uint64(transaction.UserID),
			Type:        string(transaction.Type),
			Status:      string(transaction.Status),
			Amount:      transaction.Amount,
			ReferenceId: transaction.ReferenceID,
			BetId:       uint64Ptr(transaction.BetID),
			CreatedAt:   timestamppb.New(transaction.CreatedAt),
			UpdatedAt:   timestamppb.New(transaction.UpdatedAt),
		}
	}

	return &transactionpb.GetTransactionsByUserIDResponse{Transactions: transactionsPb}, nil
}

func (h *TransactionHandler) GetTransactionsByBetID(ctx context.Context, req *transactionpb.GetTransactionsByBetIDRequest) (*transactionpb.GetTransactionsByBetIDResponse, error) {
	transactions, err := h.useCase.GetTransactionsByBetID(uint(req.BetId))
	if err != nil {
		return nil, err
	}

	transactionsPb := make([]*transactionpb.Transaction, len(transactions))
	for i, transaction := range transactions {
		transactionsPb[i] = &transactionpb.Transaction{
			Id:          uint64(transaction.ID),
			UserId:      uint64(transaction.UserID),
			Type:        string(transaction.Type),
			Status:      string(transaction.Status),
			Amount:      transaction.Amount,
			ReferenceId: transaction.ReferenceID,
			BetId:       uint64Ptr(transaction.BetID),
			CreatedAt:   timestamppb.New(transaction.CreatedAt),
			UpdatedAt:   timestamppb.New(transaction.UpdatedAt),
		}
	}

	return &transactionpb.GetTransactionsByBetIDResponse{Transactions: transactionsPb}, nil
}

func (h *TransactionHandler) UpdateTransactionStatus(ctx context.Context, req *transactionpb.UpdateTransactionStatusRequest) (*transactionpb.UpdateTransactionStatusResponse, error) {
	err := h.useCase.UpdateTransactionStatus(uint(req.Id), entities.TransactionStatus(req.Status))
	if err != nil {
		return nil, err
	}

	transaction, err := h.useCase.GetTransactionByID(uint(req.Id))
	if err != nil {
		return nil, err
	}

	transactionPb := &transactionpb.Transaction{
		Id:          uint64(transaction.ID),
		UserId:      uint64(transaction.UserID),
		Type:        string(transaction.Type),
		Status:      string(transaction.Status),
		Amount:      transaction.Amount,
		ReferenceId: transaction.ReferenceID,
		BetId:       uint64Ptr(transaction.BetID),
		CreatedAt:   timestamppb.New(transaction.CreatedAt),
		UpdatedAt:   timestamppb.New(transaction.UpdatedAt),
	}

	return &transactionpb.UpdateTransactionStatusResponse{Transaction: transactionPb}, nil
}

func (h *TransactionHandler) GetUserTransactionHistory(ctx context.Context, req *transactionpb.GetUserTransactionHistoryRequest) (*transactionpb.GetUserTransactionHistoryResponse, error) {
	transactions, err := h.useCase.GetUserTransactionHistory(uint(req.UserId), int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	transactionsPb := make([]*transactionpb.Transaction, len(transactions))
	for i, transaction := range transactions {
		transactionsPb[i] = &transactionpb.Transaction{
			Id:          uint64(transaction.ID),
			UserId:      uint64(transaction.UserID),
			Type:        string(transaction.Type),
			Status:      string(transaction.Status),
			Amount:      transaction.Amount,
			ReferenceId: transaction.ReferenceID,
			BetId:       uint64Ptr(transaction.BetID),
			CreatedAt:   timestamppb.New(transaction.CreatedAt),
			UpdatedAt:   timestamppb.New(transaction.UpdatedAt),
		}
	}

	return &transactionpb.GetUserTransactionHistoryResponse{Transactions: transactionsPb}, nil
}

func (h *TransactionHandler) GetTransactionByReferenceID(ctx context.Context, req *transactionpb.GetTransactionByReferenceIDRequest) (*transactionpb.GetTransactionByReferenceIDResponse, error) {
	transaction, err := h.useCase.GetTransactionByReferenceID(req.ReferenceId)
	if err != nil {
		return nil, err
	}

	transactionPb := &transactionpb.Transaction{
		Id:          uint64(transaction.ID),
		UserId:      uint64(transaction.UserID),
		Type:        string(transaction.Type),
		Status:      string(transaction.Status),
		Amount:      transaction.Amount,
		ReferenceId: transaction.ReferenceID,
		BetId:       uint64Ptr(transaction.BetID),
		CreatedAt:   timestamppb.New(transaction.CreatedAt),
		UpdatedAt:   timestamppb.New(transaction.UpdatedAt),
	}

	return &transactionpb.GetTransactionByReferenceIDResponse{Transaction: transactionPb}, nil
}

func (h *TransactionHandler) ProcessDeposit(ctx context.Context, req *transactionpb.ProcessDepositRequest) (*transactionpb.ProcessDepositResponse, error) {
	transaction, err := h.useCase.ProcessDeposit(uint(req.Id), 0, "")
	if err != nil {
		return nil, err
	}

	transactionPb := &transactionpb.Transaction{
		Id:          uint64(transaction.ID),
		UserId:      uint64(transaction.UserID),
		Type:        string(transaction.Type),
		Status:      string(transaction.Status),
		Amount:      transaction.Amount,
		ReferenceId: transaction.ReferenceID,
		BetId:       uint64Ptr(transaction.BetID),
		CreatedAt:   timestamppb.New(transaction.CreatedAt),
		UpdatedAt:   timestamppb.New(transaction.UpdatedAt),
	}

	return &transactionpb.ProcessDepositResponse{Transaction: transactionPb}, nil
}

func (h *TransactionHandler) ProcessWithdrawal(ctx context.Context, req *transactionpb.ProcessWithdrawalRequest) (*transactionpb.ProcessWithdrawalResponse, error) {
	transaction, err := h.useCase.ProcessWithdrawal(uint(req.Id), 0, "")
	if err != nil {
		return nil, err
	}

	transactionPb := &transactionpb.Transaction{
		Id:          uint64(transaction.ID),
		UserId:      uint64(transaction.UserID),
		Type:        string(transaction.Type),
		Status:      string(transaction.Status),
		Amount:      transaction.Amount,
		ReferenceId: transaction.ReferenceID,
		BetId:       uint64Ptr(transaction.BetID),
		CreatedAt:   timestamppb.New(transaction.CreatedAt),
		UpdatedAt:   timestamppb.New(transaction.UpdatedAt),
	}

	return &transactionpb.ProcessWithdrawalResponse{Transaction: transactionPb}, nil
}

func (h *TransactionHandler) GetUserBalance(ctx context.Context, req *transactionpb.GetUserBalanceRequest) (*transactionpb.GetUserBalanceResponse, error) {
	balance, err := h.useCase.GetUserBalance(uint(req.UserId))
	if err != nil {
		return nil, err
	}

	return &transactionpb.GetUserBalanceResponse{Balance: balance}, nil
}

// Helper functions
func uintPtr(v uint64) *uint {
	if v == 0 {
		return nil
	}
	u := uint(v)
	return &u
}

func uint64Ptr(v *uint) *uint64 {
	if v == nil {
		return nil
	}
	u := uint64(*v)
	return &u
}
