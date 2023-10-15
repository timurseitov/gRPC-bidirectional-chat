// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: chat.proto

package chatserver

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Services_ChatService_FullMethodName = "/chatserver.Services/ChatService"
)

// ServicesClient is the client API for Services service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicesClient interface {
	ChatService(ctx context.Context, opts ...grpc.CallOption) (Services_ChatServiceClient, error)
}

type servicesClient struct {
	cc grpc.ClientConnInterface
}

func NewServicesClient(cc grpc.ClientConnInterface) ServicesClient {
	return &servicesClient{cc}
}

func (c *servicesClient) ChatService(ctx context.Context, opts ...grpc.CallOption) (Services_ChatServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &Services_ServiceDesc.Streams[0], Services_ChatService_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &servicesChatServiceClient{stream}
	return x, nil
}

type Services_ChatServiceClient interface {
	Send(*FromClient) error
	Recv() (*FromServer, error)
	grpc.ClientStream
}

type servicesChatServiceClient struct {
	grpc.ClientStream
}

func (x *servicesChatServiceClient) Send(m *FromClient) error {
	return x.ClientStream.SendMsg(m)
}

func (x *servicesChatServiceClient) Recv() (*FromServer, error) {
	m := new(FromServer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServicesServer is the server API for Services service.
// All implementations should embed UnimplementedServicesServer
// for forward compatibility
type ServicesServer interface {
	ChatService(Services_ChatServiceServer) error
}

// UnimplementedServicesServer should be embedded to have forward compatible implementations.
type UnimplementedServicesServer struct {
}

func (UnimplementedServicesServer) ChatService(Services_ChatServiceServer) error {
	return status.Errorf(codes.Unimplemented, "method ChatService not implemented")
}

// UnsafeServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicesServer will
// result in compilation errors.
type UnsafeServicesServer interface {
	mustEmbedUnimplementedServicesServer()
}

func RegisterServicesServer(s grpc.ServiceRegistrar, srv ServicesServer) {
	s.RegisterService(&Services_ServiceDesc, srv)
}

func _Services_ChatService_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServicesServer).ChatService(&servicesChatServiceServer{stream})
}

type Services_ChatServiceServer interface {
	Send(*FromServer) error
	Recv() (*FromClient, error)
	grpc.ServerStream
}

type servicesChatServiceServer struct {
	grpc.ServerStream
}

func (x *servicesChatServiceServer) Send(m *FromServer) error {
	return x.ServerStream.SendMsg(m)
}

func (x *servicesChatServiceServer) Recv() (*FromClient, error) {
	m := new(FromClient)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Services_ServiceDesc is the grpc.ServiceDesc for Services service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Services_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatserver.Services",
	HandlerType: (*ServicesServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ChatService",
			Handler:       _Services_ChatService_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}
