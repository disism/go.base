// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/v1alpha1/auth/oauth2.proto

package auth

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
	OAuth2_ApplicationCreate_FullMethodName = "/apis.proto.v1alpha1.auth.OAuth2/ApplicationCreate"
)

// OAuth2Client is the client API for OAuth2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OAuth2Client interface {
	ApplicationCreate(ctx context.Context, in *ApplicationCreateRequest, opts ...grpc.CallOption) (*ApplicationCreateResponse, error)
}

type oAuth2Client struct {
	cc grpc.ClientConnInterface
}

func NewOAuth2Client(cc grpc.ClientConnInterface) OAuth2Client {
	return &oAuth2Client{cc}
}

func (c *oAuth2Client) ApplicationCreate(ctx context.Context, in *ApplicationCreateRequest, opts ...grpc.CallOption) (*ApplicationCreateResponse, error) {
	out := new(ApplicationCreateResponse)
	err := c.cc.Invoke(ctx, OAuth2_ApplicationCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OAuth2Server is the server API for OAuth2 service.
// All implementations should embed UnimplementedOAuth2Server
// for forward compatibility
type OAuth2Server interface {
	ApplicationCreate(context.Context, *ApplicationCreateRequest) (*ApplicationCreateResponse, error)
}

// UnimplementedOAuth2Server should be embedded to have forward compatible implementations.
type UnimplementedOAuth2Server struct {
}

func (UnimplementedOAuth2Server) ApplicationCreate(context.Context, *ApplicationCreateRequest) (*ApplicationCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplicationCreate not implemented")
}

// UnsafeOAuth2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OAuth2Server will
// result in compilation errors.
type UnsafeOAuth2Server interface {
	mustEmbedUnimplementedOAuth2Server()
}

func RegisterOAuth2Server(s grpc.ServiceRegistrar, srv OAuth2Server) {
	s.RegisterService(&OAuth2_ServiceDesc, srv)
}

func _OAuth2_ApplicationCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuth2Server).ApplicationCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OAuth2_ApplicationCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuth2Server).ApplicationCreate(ctx, req.(*ApplicationCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OAuth2_ServiceDesc is the grpc.ServiceDesc for OAuth2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OAuth2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apis.proto.v1alpha1.auth.OAuth2",
	HandlerType: (*OAuth2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApplicationCreate",
			Handler:    _OAuth2_ApplicationCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1alpha1/auth/oauth2.proto",
}
