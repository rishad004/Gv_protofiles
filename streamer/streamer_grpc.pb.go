// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: streamer.proto

package streamer_service

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
	StreamerService_Registration_FullMethodName = "/streamer.StreamerService/Registration"
	StreamerService_ChannelView_FullMethodName  = "/streamer.StreamerService/ChannelView"
	StreamerService_EditChannel_FullMethodName  = "/streamer.StreamerService/EditChannel"
)

// StreamerServiceClient is the client API for StreamerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamerServiceClient interface {
	Registration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	ChannelView(ctx context.Context, in *Verification, opts ...grpc.CallOption) (*ChannelResponse, error)
	EditChannel(ctx context.Context, in *EditRequest, opts ...grpc.CallOption) (*EditResponse, error)
}

type streamerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamerServiceClient(cc grpc.ClientConnInterface) StreamerServiceClient {
	return &streamerServiceClient{cc}
}

func (c *streamerServiceClient) Registration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, StreamerService_Registration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamerServiceClient) ChannelView(ctx context.Context, in *Verification, opts ...grpc.CallOption) (*ChannelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChannelResponse)
	err := c.cc.Invoke(ctx, StreamerService_ChannelView_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamerServiceClient) EditChannel(ctx context.Context, in *EditRequest, opts ...grpc.CallOption) (*EditResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EditResponse)
	err := c.cc.Invoke(ctx, StreamerService_EditChannel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamerServiceServer is the server API for StreamerService service.
// All implementations must embed UnimplementedStreamerServiceServer
// for forward compatibility.
type StreamerServiceServer interface {
	Registration(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
	ChannelView(context.Context, *Verification) (*ChannelResponse, error)
	EditChannel(context.Context, *EditRequest) (*EditResponse, error)
	mustEmbedUnimplementedStreamerServiceServer()
}

// UnimplementedStreamerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStreamerServiceServer struct{}

func (UnimplementedStreamerServiceServer) Registration(context.Context, *RegistrationRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registration not implemented")
}
func (UnimplementedStreamerServiceServer) ChannelView(context.Context, *Verification) (*ChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChannelView not implemented")
}
func (UnimplementedStreamerServiceServer) EditChannel(context.Context, *EditRequest) (*EditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditChannel not implemented")
}
func (UnimplementedStreamerServiceServer) mustEmbedUnimplementedStreamerServiceServer() {}
func (UnimplementedStreamerServiceServer) testEmbeddedByValue()                         {}

// UnsafeStreamerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamerServiceServer will
// result in compilation errors.
type UnsafeStreamerServiceServer interface {
	mustEmbedUnimplementedStreamerServiceServer()
}

func RegisterStreamerServiceServer(s grpc.ServiceRegistrar, srv StreamerServiceServer) {
	// If the following call pancis, it indicates UnimplementedStreamerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StreamerService_ServiceDesc, srv)
}

func _StreamerService_Registration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamerServiceServer).Registration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamerService_Registration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamerServiceServer).Registration(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamerService_ChannelView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Verification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamerServiceServer).ChannelView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamerService_ChannelView_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamerServiceServer).ChannelView(ctx, req.(*Verification))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamerService_EditChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamerServiceServer).EditChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamerService_EditChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamerServiceServer).EditChannel(ctx, req.(*EditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StreamerService_ServiceDesc is the grpc.ServiceDesc for StreamerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "streamer.StreamerService",
	HandlerType: (*StreamerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Registration",
			Handler:    _StreamerService_Registration_Handler,
		},
		{
			MethodName: "ChannelView",
			Handler:    _StreamerService_ChannelView_Handler,
		},
		{
			MethodName: "EditChannel",
			Handler:    _StreamerService_EditChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "streamer.proto",
}
