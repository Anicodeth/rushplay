// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: api/proto/bet.proto

package betpb

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
	BetService_CreateBet_FullMethodName                = "/bet.BetService/CreateBet"
	BetService_GetBet_FullMethodName                   = "/bet.BetService/GetBet"
	BetService_GetBets_FullMethodName                  = "/bet.BetService/GetBets"
	BetService_GetBetsByUserID_FullMethodName          = "/bet.BetService/GetBetsByUserID"
	BetService_GetBetsByGameID_FullMethodName          = "/bet.BetService/GetBetsByGameID"
	BetService_UpdateBetStatus_FullMethodName          = "/bet.BetService/UpdateBetStatus"
	BetService_ProcessBetResult_FullMethodName         = "/bet.BetService/ProcessBetResult"
	BetService_GetUserTotalBets_FullMethodName         = "/bet.BetService/GetUserTotalBets"
	BetService_GetUserTotalBetsByGameID_FullMethodName = "/bet.BetService/GetUserTotalBetsByGameID"
	BetService_GetUserTotalWinnings_FullMethodName     = "/bet.BetService/GetUserTotalWinnings"
)

// BetServiceClient is the client API for BetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BetServiceClient interface {
	CreateBet(ctx context.Context, in *CreateBetRequest, opts ...grpc.CallOption) (*CreateBetResponse, error)
	GetBet(ctx context.Context, in *GetBetRequest, opts ...grpc.CallOption) (*GetBetResponse, error)
	GetBets(ctx context.Context, in *GetBetsRequest, opts ...grpc.CallOption) (*GetBetsResponse, error)
	GetBetsByUserID(ctx context.Context, in *GetBetsByUserIDRequest, opts ...grpc.CallOption) (*GetBetsByUserIDResponse, error)
	GetBetsByGameID(ctx context.Context, in *GetBetsByGameIDRequest, opts ...grpc.CallOption) (*GetBetsByGameIDResponse, error)
	UpdateBetStatus(ctx context.Context, in *UpdateBetStatusRequest, opts ...grpc.CallOption) (*UpdateBetStatusResponse, error)
	ProcessBetResult(ctx context.Context, in *ProcessBetResultRequest, opts ...grpc.CallOption) (*ProcessBetResultResponse, error)
	GetUserTotalBets(ctx context.Context, in *GetUserTotalBetsRequest, opts ...grpc.CallOption) (*GetUserTotalBetsResponse, error)
	GetUserTotalBetsByGameID(ctx context.Context, in *GetUserTotalBetsByGameIDRequest, opts ...grpc.CallOption) (*GetUserTotalBetsByGameIDResponse, error)
	GetUserTotalWinnings(ctx context.Context, in *GetUserTotalWinningsRequest, opts ...grpc.CallOption) (*GetUserTotalWinningsResponse, error)
}

type betServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBetServiceClient(cc grpc.ClientConnInterface) BetServiceClient {
	return &betServiceClient{cc}
}

func (c *betServiceClient) CreateBet(ctx context.Context, in *CreateBetRequest, opts ...grpc.CallOption) (*CreateBetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBetResponse)
	err := c.cc.Invoke(ctx, BetService_CreateBet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *betServiceClient) GetBet(ctx context.Context, in *GetBetRequest, opts ...grpc.CallOption) (*GetBetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBetResponse)
	err := c.cc.Invoke(ctx, BetService_GetBet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *betServiceClient) GetBets(ctx context.Context, in *GetBetsRequest, opts ...grpc.CallOption) (*GetBetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBetsResponse)
	err := c.cc.Invoke(ctx, BetService_GetBets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *betServiceClient) GetBetsByUserID(ctx context.Context, in *GetBetsByUserIDRequest, opts ...grpc.CallOption) (*GetBetsByUserIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBetsByUserIDResponse)
	err := c.cc.Invoke(ctx, BetService_GetBetsByUserID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *betServiceClient) GetBetsByGameID(ctx context.Context, in *GetBetsByGameIDRequest, opts ...grpc.CallOption) (*GetBetsByGameIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBetsByGameIDResponse)
	err := c.cc.Invoke(ctx, BetService_GetBetsByGameID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *betServiceClient) UpdateBetStatus(ctx context.Context, in *UpdateBetStatusRequest, opts ...grpc.CallOption) (*UpdateBetStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBetStatusResponse)
	err := c.cc.Invoke(ctx, BetService_UpdateBetStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *betServiceClient) ProcessBetResult(ctx context.Context, in *ProcessBetResultRequest, opts ...grpc.CallOption) (*ProcessBetResultResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProcessBetResultResponse)
	err := c.cc.Invoke(ctx, BetService_ProcessBetResult_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *betServiceClient) GetUserTotalBets(ctx context.Context, in *GetUserTotalBetsRequest, opts ...grpc.CallOption) (*GetUserTotalBetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserTotalBetsResponse)
	err := c.cc.Invoke(ctx, BetService_GetUserTotalBets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *betServiceClient) GetUserTotalBetsByGameID(ctx context.Context, in *GetUserTotalBetsByGameIDRequest, opts ...grpc.CallOption) (*GetUserTotalBetsByGameIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserTotalBetsByGameIDResponse)
	err := c.cc.Invoke(ctx, BetService_GetUserTotalBetsByGameID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *betServiceClient) GetUserTotalWinnings(ctx context.Context, in *GetUserTotalWinningsRequest, opts ...grpc.CallOption) (*GetUserTotalWinningsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserTotalWinningsResponse)
	err := c.cc.Invoke(ctx, BetService_GetUserTotalWinnings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BetServiceServer is the server API for BetService service.
// All implementations must embed UnimplementedBetServiceServer
// for forward compatibility.
type BetServiceServer interface {
	CreateBet(context.Context, *CreateBetRequest) (*CreateBetResponse, error)
	GetBet(context.Context, *GetBetRequest) (*GetBetResponse, error)
	GetBets(context.Context, *GetBetsRequest) (*GetBetsResponse, error)
	GetBetsByUserID(context.Context, *GetBetsByUserIDRequest) (*GetBetsByUserIDResponse, error)
	GetBetsByGameID(context.Context, *GetBetsByGameIDRequest) (*GetBetsByGameIDResponse, error)
	UpdateBetStatus(context.Context, *UpdateBetStatusRequest) (*UpdateBetStatusResponse, error)
	ProcessBetResult(context.Context, *ProcessBetResultRequest) (*ProcessBetResultResponse, error)
	GetUserTotalBets(context.Context, *GetUserTotalBetsRequest) (*GetUserTotalBetsResponse, error)
	GetUserTotalBetsByGameID(context.Context, *GetUserTotalBetsByGameIDRequest) (*GetUserTotalBetsByGameIDResponse, error)
	GetUserTotalWinnings(context.Context, *GetUserTotalWinningsRequest) (*GetUserTotalWinningsResponse, error)
	mustEmbedUnimplementedBetServiceServer()
}

// UnimplementedBetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBetServiceServer struct{}

func (UnimplementedBetServiceServer) CreateBet(context.Context, *CreateBetRequest) (*CreateBetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBet not implemented")
}
func (UnimplementedBetServiceServer) GetBet(context.Context, *GetBetRequest) (*GetBetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBet not implemented")
}
func (UnimplementedBetServiceServer) GetBets(context.Context, *GetBetsRequest) (*GetBetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBets not implemented")
}
func (UnimplementedBetServiceServer) GetBetsByUserID(context.Context, *GetBetsByUserIDRequest) (*GetBetsByUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBetsByUserID not implemented")
}
func (UnimplementedBetServiceServer) GetBetsByGameID(context.Context, *GetBetsByGameIDRequest) (*GetBetsByGameIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBetsByGameID not implemented")
}
func (UnimplementedBetServiceServer) UpdateBetStatus(context.Context, *UpdateBetStatusRequest) (*UpdateBetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBetStatus not implemented")
}
func (UnimplementedBetServiceServer) ProcessBetResult(context.Context, *ProcessBetResultRequest) (*ProcessBetResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessBetResult not implemented")
}
func (UnimplementedBetServiceServer) GetUserTotalBets(context.Context, *GetUserTotalBetsRequest) (*GetUserTotalBetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTotalBets not implemented")
}
func (UnimplementedBetServiceServer) GetUserTotalBetsByGameID(context.Context, *GetUserTotalBetsByGameIDRequest) (*GetUserTotalBetsByGameIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTotalBetsByGameID not implemented")
}
func (UnimplementedBetServiceServer) GetUserTotalWinnings(context.Context, *GetUserTotalWinningsRequest) (*GetUserTotalWinningsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTotalWinnings not implemented")
}
func (UnimplementedBetServiceServer) mustEmbedUnimplementedBetServiceServer() {}
func (UnimplementedBetServiceServer) testEmbeddedByValue()                    {}

// UnsafeBetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BetServiceServer will
// result in compilation errors.
type UnsafeBetServiceServer interface {
	mustEmbedUnimplementedBetServiceServer()
}

func RegisterBetServiceServer(s grpc.ServiceRegistrar, srv BetServiceServer) {
	// If the following call pancis, it indicates UnimplementedBetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BetService_ServiceDesc, srv)
}

func _BetService_CreateBet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetServiceServer).CreateBet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BetService_CreateBet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetServiceServer).CreateBet(ctx, req.(*CreateBetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BetService_GetBet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetServiceServer).GetBet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BetService_GetBet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetServiceServer).GetBet(ctx, req.(*GetBetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BetService_GetBets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetServiceServer).GetBets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BetService_GetBets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetServiceServer).GetBets(ctx, req.(*GetBetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BetService_GetBetsByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBetsByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetServiceServer).GetBetsByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BetService_GetBetsByUserID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetServiceServer).GetBetsByUserID(ctx, req.(*GetBetsByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BetService_GetBetsByGameID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBetsByGameIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetServiceServer).GetBetsByGameID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BetService_GetBetsByGameID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetServiceServer).GetBetsByGameID(ctx, req.(*GetBetsByGameIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BetService_UpdateBetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetServiceServer).UpdateBetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BetService_UpdateBetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetServiceServer).UpdateBetStatus(ctx, req.(*UpdateBetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BetService_ProcessBetResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessBetResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetServiceServer).ProcessBetResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BetService_ProcessBetResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetServiceServer).ProcessBetResult(ctx, req.(*ProcessBetResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BetService_GetUserTotalBets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserTotalBetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetServiceServer).GetUserTotalBets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BetService_GetUserTotalBets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetServiceServer).GetUserTotalBets(ctx, req.(*GetUserTotalBetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BetService_GetUserTotalBetsByGameID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserTotalBetsByGameIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetServiceServer).GetUserTotalBetsByGameID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BetService_GetUserTotalBetsByGameID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetServiceServer).GetUserTotalBetsByGameID(ctx, req.(*GetUserTotalBetsByGameIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BetService_GetUserTotalWinnings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserTotalWinningsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetServiceServer).GetUserTotalWinnings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BetService_GetUserTotalWinnings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetServiceServer).GetUserTotalWinnings(ctx, req.(*GetUserTotalWinningsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BetService_ServiceDesc is the grpc.ServiceDesc for BetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bet.BetService",
	HandlerType: (*BetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBet",
			Handler:    _BetService_CreateBet_Handler,
		},
		{
			MethodName: "GetBet",
			Handler:    _BetService_GetBet_Handler,
		},
		{
			MethodName: "GetBets",
			Handler:    _BetService_GetBets_Handler,
		},
		{
			MethodName: "GetBetsByUserID",
			Handler:    _BetService_GetBetsByUserID_Handler,
		},
		{
			MethodName: "GetBetsByGameID",
			Handler:    _BetService_GetBetsByGameID_Handler,
		},
		{
			MethodName: "UpdateBetStatus",
			Handler:    _BetService_UpdateBetStatus_Handler,
		},
		{
			MethodName: "ProcessBetResult",
			Handler:    _BetService_ProcessBetResult_Handler,
		},
		{
			MethodName: "GetUserTotalBets",
			Handler:    _BetService_GetUserTotalBets_Handler,
		},
		{
			MethodName: "GetUserTotalBetsByGameID",
			Handler:    _BetService_GetUserTotalBetsByGameID_Handler,
		},
		{
			MethodName: "GetUserTotalWinnings",
			Handler:    _BetService_GetUserTotalWinnings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/bet.proto",
}
