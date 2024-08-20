// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: msg.proto

package __

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
	MessageClient_Upload_FullMethodName = "/msg.MessageClient/Upload"
	MessageClient_Pull_FullMethodName   = "/msg.MessageClient/Pull"
)

// MessageClientClient is the client API for MessageClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageClientClient interface {
	Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
	Pull(ctx context.Context, in *PullRequest, opts ...grpc.CallOption) (*PullResponse, error)
}

type messageClientClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageClientClient(cc grpc.ClientConnInterface) MessageClientClient {
	return &messageClientClient{cc}
}

func (c *messageClientClient) Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadResponse)
	err := c.cc.Invoke(ctx, MessageClient_Upload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClientClient) Pull(ctx context.Context, in *PullRequest, opts ...grpc.CallOption) (*PullResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PullResponse)
	err := c.cc.Invoke(ctx, MessageClient_Pull_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageClientServer is the server API for MessageClient service.
// All implementations must embed UnimplementedMessageClientServer
// for forward compatibility.
type MessageClientServer interface {
	Upload(context.Context, *UploadRequest) (*UploadResponse, error)
	Pull(context.Context, *PullRequest) (*PullResponse, error)
	mustEmbedUnimplementedMessageClientServer()
}

// UnimplementedMessageClientServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMessageClientServer struct{}

func (UnimplementedMessageClientServer) Upload(context.Context, *UploadRequest) (*UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedMessageClientServer) Pull(context.Context, *PullRequest) (*PullResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pull not implemented")
}
func (UnimplementedMessageClientServer) mustEmbedUnimplementedMessageClientServer() {}
func (UnimplementedMessageClientServer) testEmbeddedByValue()                       {}

// UnsafeMessageClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageClientServer will
// result in compilation errors.
type UnsafeMessageClientServer interface {
	mustEmbedUnimplementedMessageClientServer()
}

func RegisterMessageClientServer(s grpc.ServiceRegistrar, srv MessageClientServer) {
	// If the following call pancis, it indicates UnimplementedMessageClientServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MessageClient_ServiceDesc, srv)
}

func _MessageClient_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageClientServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageClient_Upload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageClientServer).Upload(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageClient_Pull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageClientServer).Pull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageClient_Pull_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageClientServer).Pull(ctx, req.(*PullRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageClient_ServiceDesc is the grpc.ServiceDesc for MessageClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msg.MessageClient",
	HandlerType: (*MessageClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _MessageClient_Upload_Handler,
		},
		{
			MethodName: "Pull",
			Handler:    _MessageClient_Pull_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msg.proto",
}
