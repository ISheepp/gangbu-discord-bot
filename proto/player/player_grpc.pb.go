// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: proto/player/player.proto

package player

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PlayerRequestClient is the client API for PlayerRequest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlayerRequestClient interface {
	ShowPlayerInfo(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*PlayerInfoVo, error)
}

type playerRequestClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayerRequestClient(cc grpc.ClientConnInterface) PlayerRequestClient {
	return &playerRequestClient{cc}
}

func (c *playerRequestClient) ShowPlayerInfo(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*PlayerInfoVo, error) {
	out := new(PlayerInfoVo)
	err := c.cc.Invoke(ctx, "/proto.PlayerRequest/ShowPlayerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerRequestServer is the server API for PlayerRequest service.
// All implementations must embed UnimplementedPlayerRequestServer
// for forward compatibility
type PlayerRequestServer interface {
	ShowPlayerInfo(context.Context, *wrapperspb.StringValue) (*PlayerInfoVo, error)
	mustEmbedUnimplementedPlayerRequestServer()
}

// UnimplementedPlayerRequestServer must be embedded to have forward compatible implementations.
type UnimplementedPlayerRequestServer struct {
}

func (UnimplementedPlayerRequestServer) ShowPlayerInfo(context.Context, *wrapperspb.StringValue) (*PlayerInfoVo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowPlayerInfo not implemented")
}
func (UnimplementedPlayerRequestServer) mustEmbedUnimplementedPlayerRequestServer() {}

// UnsafePlayerRequestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlayerRequestServer will
// result in compilation errors.
type UnsafePlayerRequestServer interface {
	mustEmbedUnimplementedPlayerRequestServer()
}

func RegisterPlayerRequestServer(s grpc.ServiceRegistrar, srv PlayerRequestServer) {
	s.RegisterService(&PlayerRequest_ServiceDesc, srv)
}

func _PlayerRequest_ShowPlayerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerRequestServer).ShowPlayerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PlayerRequest/ShowPlayerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerRequestServer).ShowPlayerInfo(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

// PlayerRequest_ServiceDesc is the grpc.ServiceDesc for PlayerRequest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlayerRequest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PlayerRequest",
	HandlerType: (*PlayerRequestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShowPlayerInfo",
			Handler:    _PlayerRequest_ShowPlayerInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/player/player.proto",
}
