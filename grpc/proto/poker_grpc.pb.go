// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: poker.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Poker_GetNuts_FullMethodName       = "/poker.Poker/GetNuts"
	Poker_LotsOfReplies_FullMethodName = "/poker.Poker/LotsOfReplies"
)

// PokerClient is the client API for Poker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PokerClient interface {
	GetNuts(ctx context.Context, in *GetNutsRequest, opts ...grpc.CallOption) (*GetNutsResponse, error)
	LotsOfReplies(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Poker_LotsOfRepliesClient, error)
}

type pokerClient struct {
	cc grpc.ClientConnInterface
}

func NewPokerClient(cc grpc.ClientConnInterface) PokerClient {
	return &pokerClient{cc}
}

func (c *pokerClient) GetNuts(ctx context.Context, in *GetNutsRequest, opts ...grpc.CallOption) (*GetNutsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNutsResponse)
	err := c.cc.Invoke(ctx, Poker_GetNuts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokerClient) LotsOfReplies(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Poker_LotsOfRepliesClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Poker_ServiceDesc.Streams[0], Poker_LotsOfReplies_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &pokerLotsOfRepliesClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Poker_LotsOfRepliesClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type pokerLotsOfRepliesClient struct {
	grpc.ClientStream
}

func (x *pokerLotsOfRepliesClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PokerServer is the server API for Poker service.
// All implementations must embed UnimplementedPokerServer
// for forward compatibility
type PokerServer interface {
	GetNuts(context.Context, *GetNutsRequest) (*GetNutsResponse, error)
	LotsOfReplies(*HelloRequest, Poker_LotsOfRepliesServer) error
	mustEmbedUnimplementedPokerServer()
}

// UnimplementedPokerServer must be embedded to have forward compatible implementations.
type UnimplementedPokerServer struct {
}

func (UnimplementedPokerServer) GetNuts(context.Context, *GetNutsRequest) (*GetNutsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNuts not implemented")
}
func (UnimplementedPokerServer) LotsOfReplies(*HelloRequest, Poker_LotsOfRepliesServer) error {
	return status.Errorf(codes.Unimplemented, "method LotsOfReplies not implemented")
}
func (UnimplementedPokerServer) mustEmbedUnimplementedPokerServer() {}

// UnsafePokerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PokerServer will
// result in compilation errors.
type UnsafePokerServer interface {
	mustEmbedUnimplementedPokerServer()
}

func RegisterPokerServer(s grpc.ServiceRegistrar, srv PokerServer) {
	s.RegisterService(&Poker_ServiceDesc, srv)
}

func _Poker_GetNuts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNutsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokerServer).GetNuts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Poker_GetNuts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokerServer).GetNuts(ctx, req.(*GetNutsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poker_LotsOfReplies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PokerServer).LotsOfReplies(m, &pokerLotsOfRepliesServer{ServerStream: stream})
}

type Poker_LotsOfRepliesServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type pokerLotsOfRepliesServer struct {
	grpc.ServerStream
}

func (x *pokerLotsOfRepliesServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Poker_ServiceDesc is the grpc.ServiceDesc for Poker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Poker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "poker.Poker",
	HandlerType: (*PokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNuts",
			Handler:    _Poker_GetNuts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LotsOfReplies",
			Handler:       _Poker_LotsOfReplies_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "poker.proto",
}
