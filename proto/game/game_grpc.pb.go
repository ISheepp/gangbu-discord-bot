// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: proto/game/game.proto

package game

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GameRequestClient is the client API for GameRequest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameRequestClient interface {
	// 创建一次bet
	CreateGame(ctx context.Context, in *GameCreateBo, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	// 获取最近5条游戏记录
	GetLastFiveGameHistoryByDiscordId(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*GameHistoryDtoSlice, error)
	// 从the graph获取最近5条游戏记录
	GetLastFiveGameHistoryByDiscordIdFromTheGraph(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*GameHistoryDtoSlice, error)
	// 发送消息
	SendCallbackMessage(ctx context.Context, in *CallbackMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type gameRequestClient struct {
	cc grpc.ClientConnInterface
}

func NewGameRequestClient(cc grpc.ClientConnInterface) GameRequestClient {
	return &gameRequestClient{cc}
}

func (c *gameRequestClient) CreateGame(ctx context.Context, in *GameCreateBo, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/proto.GameRequest/CreateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameRequestClient) GetLastFiveGameHistoryByDiscordId(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*GameHistoryDtoSlice, error) {
	out := new(GameHistoryDtoSlice)
	err := c.cc.Invoke(ctx, "/proto.GameRequest/GetLastFiveGameHistoryByDiscordId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameRequestClient) GetLastFiveGameHistoryByDiscordIdFromTheGraph(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*GameHistoryDtoSlice, error) {
	out := new(GameHistoryDtoSlice)
	err := c.cc.Invoke(ctx, "/proto.GameRequest/GetLastFiveGameHistoryByDiscordIdFromTheGraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameRequestClient) SendCallbackMessage(ctx context.Context, in *CallbackMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.GameRequest/SendCallbackMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameRequestServer is the server API for GameRequest service.
// All implementations must embed UnimplementedGameRequestServer
// for forward compatibility
type GameRequestServer interface {
	// 创建一次bet
	CreateGame(context.Context, *GameCreateBo) (*wrapperspb.StringValue, error)
	// 获取最近5条游戏记录
	GetLastFiveGameHistoryByDiscordId(context.Context, *wrapperspb.StringValue) (*GameHistoryDtoSlice, error)
	// 从the graph获取最近5条游戏记录
	GetLastFiveGameHistoryByDiscordIdFromTheGraph(context.Context, *wrapperspb.StringValue) (*GameHistoryDtoSlice, error)
	// 发送消息
	SendCallbackMessage(context.Context, *CallbackMessage) (*emptypb.Empty, error)
	mustEmbedUnimplementedGameRequestServer()
}

// UnimplementedGameRequestServer must be embedded to have forward compatible implementations.
type UnimplementedGameRequestServer struct {
}

func (UnimplementedGameRequestServer) CreateGame(context.Context, *GameCreateBo) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (UnimplementedGameRequestServer) GetLastFiveGameHistoryByDiscordId(context.Context, *wrapperspb.StringValue) (*GameHistoryDtoSlice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastFiveGameHistoryByDiscordId not implemented")
}
func (UnimplementedGameRequestServer) GetLastFiveGameHistoryByDiscordIdFromTheGraph(context.Context, *wrapperspb.StringValue) (*GameHistoryDtoSlice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastFiveGameHistoryByDiscordIdFromTheGraph not implemented")
}
func (UnimplementedGameRequestServer) SendCallbackMessage(context.Context, *CallbackMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCallbackMessage not implemented")
}
func (UnimplementedGameRequestServer) mustEmbedUnimplementedGameRequestServer() {}

// UnsafeGameRequestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameRequestServer will
// result in compilation errors.
type UnsafeGameRequestServer interface {
	mustEmbedUnimplementedGameRequestServer()
}

func RegisterGameRequestServer(s grpc.ServiceRegistrar, srv GameRequestServer) {
	s.RegisterService(&GameRequest_ServiceDesc, srv)
}

func _GameRequest_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameCreateBo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameRequestServer).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GameRequest/CreateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameRequestServer).CreateGame(ctx, req.(*GameCreateBo))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameRequest_GetLastFiveGameHistoryByDiscordId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameRequestServer).GetLastFiveGameHistoryByDiscordId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GameRequest/GetLastFiveGameHistoryByDiscordId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameRequestServer).GetLastFiveGameHistoryByDiscordId(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameRequest_GetLastFiveGameHistoryByDiscordIdFromTheGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameRequestServer).GetLastFiveGameHistoryByDiscordIdFromTheGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GameRequest/GetLastFiveGameHistoryByDiscordIdFromTheGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameRequestServer).GetLastFiveGameHistoryByDiscordIdFromTheGraph(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameRequest_SendCallbackMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameRequestServer).SendCallbackMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GameRequest/SendCallbackMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameRequestServer).SendCallbackMessage(ctx, req.(*CallbackMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// GameRequest_ServiceDesc is the grpc.ServiceDesc for GameRequest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameRequest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GameRequest",
	HandlerType: (*GameRequestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGame",
			Handler:    _GameRequest_CreateGame_Handler,
		},
		{
			MethodName: "GetLastFiveGameHistoryByDiscordId",
			Handler:    _GameRequest_GetLastFiveGameHistoryByDiscordId_Handler,
		},
		{
			MethodName: "GetLastFiveGameHistoryByDiscordIdFromTheGraph",
			Handler:    _GameRequest_GetLastFiveGameHistoryByDiscordIdFromTheGraph_Handler,
		},
		{
			MethodName: "SendCallbackMessage",
			Handler:    _GameRequest_SendCallbackMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/game/game.proto",
}
