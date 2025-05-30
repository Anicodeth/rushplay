// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: api/proto/transaction.proto

package transactionpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TransactionService_CreateTransaction_FullMethodName           = "/transaction.TransactionService/CreateTransaction"
	TransactionService_GetTransaction_FullMethodName              = "/transaction.TransactionService/GetTransaction"
	TransactionService_GetTransactions_FullMethodName             = "/transaction.TransactionService/GetTransactions"
	TransactionService_GetTransactionsByUserID_FullMethodName     = "/transaction.TransactionService/GetTransactionsByUserID"
	TransactionService_GetTransactionsByBetID_FullMethodName      = "/transaction.TransactionService/GetTransactionsByBetID"
	TransactionService_UpdateTransactionStatus_FullMethodName     = "/transaction.TransactionService/UpdateTransactionStatus"
	TransactionService_GetUserTransactionHistory_FullMethodName   = "/transaction.TransactionService/GetUserTransactionHistory"
	TransactionService_GetTransactionByReferenceID_FullMethodName = "/transaction.TransactionService/GetTransactionByReferenceID"
	TransactionService_ProcessDeposit_FullMethodName              = "/transaction.TransactionService/ProcessDeposit"
	TransactionService_ProcessWithdrawal_FullMethodName           = "/transaction.TransactionService/ProcessWithdrawal"
	TransactionService_GetUserBalance_FullMethodName              = "/transaction.TransactionService/GetUserBalance"
)

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionServiceClient interface {
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error)
	GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error)
	GetTransactionsByUserID(ctx context.Context, in *GetTransactionsByUserIDRequest, opts ...grpc.CallOption) (*GetTransactionsByUserIDResponse, error)
	GetTransactionsByBetID(ctx context.Context, in *GetTransactionsByBetIDRequest, opts ...grpc.CallOption) (*GetTransactionsByBetIDResponse, error)
	UpdateTransactionStatus(ctx context.Context, in *UpdateTransactionStatusRequest, opts ...grpc.CallOption) (*UpdateTransactionStatusResponse, error)
	GetUserTransactionHistory(ctx context.Context, in *GetUserTransactionHistoryRequest, opts ...grpc.CallOption) (*GetUserTransactionHistoryResponse, error)
	GetTransactionByReferenceID(ctx context.Context, in *GetTransactionByReferenceIDRequest, opts ...grpc.CallOption) (*GetTransactionByReferenceIDResponse, error)
	ProcessDeposit(ctx context.Context, in *ProcessDepositRequest, opts ...grpc.CallOption) (*ProcessDepositResponse, error)
	ProcessWithdrawal(ctx context.Context, in *ProcessWithdrawalRequest, opts ...grpc.CallOption) (*ProcessWithdrawalResponse, error)
	GetUserBalance(ctx context.Context, in *GetUserBalanceRequest, opts ...grpc.CallOption) (*GetUserBalanceResponse, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, TransactionService_CreateTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransactionResponse)
	err := c.cc.Invoke(ctx, TransactionService_GetTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransactionsResponse)
	err := c.cc.Invoke(ctx, TransactionService_GetTransactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetTransactionsByUserID(ctx context.Context, in *GetTransactionsByUserIDRequest, opts ...grpc.CallOption) (*GetTransactionsByUserIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransactionsByUserIDResponse)
	err := c.cc.Invoke(ctx, TransactionService_GetTransactionsByUserID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetTransactionsByBetID(ctx context.Context, in *GetTransactionsByBetIDRequest, opts ...grpc.CallOption) (*GetTransactionsByBetIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransactionsByBetIDResponse)
	err := c.cc.Invoke(ctx, TransactionService_GetTransactionsByBetID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) UpdateTransactionStatus(ctx context.Context, in *UpdateTransactionStatusRequest, opts ...grpc.CallOption) (*UpdateTransactionStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTransactionStatusResponse)
	err := c.cc.Invoke(ctx, TransactionService_UpdateTransactionStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetUserTransactionHistory(ctx context.Context, in *GetUserTransactionHistoryRequest, opts ...grpc.CallOption) (*GetUserTransactionHistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserTransactionHistoryResponse)
	err := c.cc.Invoke(ctx, TransactionService_GetUserTransactionHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetTransactionByReferenceID(ctx context.Context, in *GetTransactionByReferenceIDRequest, opts ...grpc.CallOption) (*GetTransactionByReferenceIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransactionByReferenceIDResponse)
	err := c.cc.Invoke(ctx, TransactionService_GetTransactionByReferenceID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ProcessDeposit(ctx context.Context, in *ProcessDepositRequest, opts ...grpc.CallOption) (*ProcessDepositResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProcessDepositResponse)
	err := c.cc.Invoke(ctx, TransactionService_ProcessDeposit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ProcessWithdrawal(ctx context.Context, in *ProcessWithdrawalRequest, opts ...grpc.CallOption) (*ProcessWithdrawalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProcessWithdrawalResponse)
	err := c.cc.Invoke(ctx, TransactionService_ProcessWithdrawal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetUserBalance(ctx context.Context, in *GetUserBalanceRequest, opts ...grpc.CallOption) (*GetUserBalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserBalanceResponse)
	err := c.cc.Invoke(ctx, TransactionService_GetUserBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
// All implementations must embed UnimplementedTransactionServiceServer
// for forward compatibility.
type TransactionServiceServer interface {
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error)
	GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error)
	GetTransactionsByUserID(context.Context, *GetTransactionsByUserIDRequest) (*GetTransactionsByUserIDResponse, error)
	GetTransactionsByBetID(context.Context, *GetTransactionsByBetIDRequest) (*GetTransactionsByBetIDResponse, error)
	UpdateTransactionStatus(context.Context, *UpdateTransactionStatusRequest) (*UpdateTransactionStatusResponse, error)
	GetUserTransactionHistory(context.Context, *GetUserTransactionHistoryRequest) (*GetUserTransactionHistoryResponse, error)
	GetTransactionByReferenceID(context.Context, *GetTransactionByReferenceIDRequest) (*GetTransactionByReferenceIDResponse, error)
	ProcessDeposit(context.Context, *ProcessDepositRequest) (*ProcessDepositResponse, error)
	ProcessWithdrawal(context.Context, *ProcessWithdrawalRequest) (*ProcessWithdrawalResponse, error)
	GetUserBalance(context.Context, *GetUserBalanceRequest) (*GetUserBalanceResponse, error)
	mustEmbedUnimplementedTransactionServiceServer()
}

// UnimplementedTransactionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTransactionServiceServer struct{}

func (UnimplementedTransactionServiceServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
}
func (UnimplementedTransactionServiceServer) GetTransactionsByUserID(context.Context, *GetTransactionsByUserIDRequest) (*GetTransactionsByUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionsByUserID not implemented")
}
func (UnimplementedTransactionServiceServer) GetTransactionsByBetID(context.Context, *GetTransactionsByBetIDRequest) (*GetTransactionsByBetIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionsByBetID not implemented")
}
func (UnimplementedTransactionServiceServer) UpdateTransactionStatus(context.Context, *UpdateTransactionStatusRequest) (*UpdateTransactionStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransactionStatus not implemented")
}
func (UnimplementedTransactionServiceServer) GetUserTransactionHistory(context.Context, *GetUserTransactionHistoryRequest) (*GetUserTransactionHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTransactionHistory not implemented")
}
func (UnimplementedTransactionServiceServer) GetTransactionByReferenceID(context.Context, *GetTransactionByReferenceIDRequest) (*GetTransactionByReferenceIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionByReferenceID not implemented")
}
func (UnimplementedTransactionServiceServer) ProcessDeposit(context.Context, *ProcessDepositRequest) (*ProcessDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessDeposit not implemented")
}
func (UnimplementedTransactionServiceServer) ProcessWithdrawal(context.Context, *ProcessWithdrawalRequest) (*ProcessWithdrawalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessWithdrawal not implemented")
}
func (UnimplementedTransactionServiceServer) GetUserBalance(context.Context, *GetUserBalanceRequest) (*GetUserBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBalance not implemented")
}
func (UnimplementedTransactionServiceServer) mustEmbedUnimplementedTransactionServiceServer() {}
func (UnimplementedTransactionServiceServer) testEmbeddedByValue()                            {}

// UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceServer will
// result in compilation errors.
type UnsafeTransactionServiceServer interface {
	mustEmbedUnimplementedTransactionServiceServer()
}

func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer) {
	// If the following call pancis, it indicates UnimplementedTransactionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TransactionService_ServiceDesc, srv)
}

func _TransactionService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_CreateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_GetTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_GetTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetTransactions(ctx, req.(*GetTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetTransactionsByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetTransactionsByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_GetTransactionsByUserID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetTransactionsByUserID(ctx, req.(*GetTransactionsByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetTransactionsByBetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsByBetIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetTransactionsByBetID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_GetTransactionsByBetID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetTransactionsByBetID(ctx, req.(*GetTransactionsByBetIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_UpdateTransactionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTransactionStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).UpdateTransactionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_UpdateTransactionStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).UpdateTransactionStatus(ctx, req.(*UpdateTransactionStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetUserTransactionHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserTransactionHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetUserTransactionHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_GetUserTransactionHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetUserTransactionHistory(ctx, req.(*GetUserTransactionHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetTransactionByReferenceID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionByReferenceIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetTransactionByReferenceID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_GetTransactionByReferenceID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetTransactionByReferenceID(ctx, req.(*GetTransactionByReferenceIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_ProcessDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ProcessDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_ProcessDeposit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ProcessDeposit(ctx, req.(*ProcessDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_ProcessWithdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessWithdrawalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ProcessWithdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_ProcessWithdrawal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ProcessWithdrawal(ctx, req.(*ProcessWithdrawalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetUserBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetUserBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_GetUserBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetUserBalance(ctx, req.(*GetUserBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionService_ServiceDesc is the grpc.ServiceDesc for TransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTransaction",
			Handler:    _TransactionService_CreateTransaction_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _TransactionService_GetTransaction_Handler,
		},
		{
			MethodName: "GetTransactions",
			Handler:    _TransactionService_GetTransactions_Handler,
		},
		{
			MethodName: "GetTransactionsByUserID",
			Handler:    _TransactionService_GetTransactionsByUserID_Handler,
		},
		{
			MethodName: "GetTransactionsByBetID",
			Handler:    _TransactionService_GetTransactionsByBetID_Handler,
		},
		{
			MethodName: "UpdateTransactionStatus",
			Handler:    _TransactionService_UpdateTransactionStatus_Handler,
		},
		{
			MethodName: "GetUserTransactionHistory",
			Handler:    _TransactionService_GetUserTransactionHistory_Handler,
		},
		{
			MethodName: "GetTransactionByReferenceID",
			Handler:    _TransactionService_GetTransactionByReferenceID_Handler,
		},
		{
			MethodName: "ProcessDeposit",
			Handler:    _TransactionService_ProcessDeposit_Handler,
		},
		{
			MethodName: "ProcessWithdrawal",
			Handler:    _TransactionService_ProcessWithdrawal_Handler,
		},
		{
			MethodName: "GetUserBalance",
			Handler:    _TransactionService_GetUserBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/transaction.proto",
}
