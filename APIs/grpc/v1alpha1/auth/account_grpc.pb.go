// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/v1alpha1/auth/account.proto

package auth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Account_UserGet_FullMethodName          = "/apis.proto.v1alpha1.auth.Account/UserGet"
	Account_UserCreate_FullMethodName       = "/apis.proto.v1alpha1.auth.Account/UserCreate"
	Account_UserUpdate_FullMethodName       = "/apis.proto.v1alpha1.auth.Account/UserUpdate"
	Account_UserEditPassword_FullMethodName = "/apis.proto.v1alpha1.auth.Account/UserEditPassword"
	Account_UserEditUsername_FullMethodName = "/apis.proto.v1alpha1.auth.Account/UserEditUsername"
	Account_UserEditEMail_FullMethodName    = "/apis.proto.v1alpha1.auth.Account/UserEditEMail"
)

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountClient interface {
	UserGet(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserGetResponse, error)
	UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
	UserUpdate(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserUpdateResponse, error)
	UserEditPassword(ctx context.Context, in *UserEditPasswordRequest, opts ...grpc.CallOption) (*UserEditPasswordResponse, error)
	UserEditUsername(ctx context.Context, in *UserEditUsernameRequest, opts ...grpc.CallOption) (*UserEditUsernameResponse, error)
	UserEditEMail(ctx context.Context, in *UserEditEMailRequest, opts ...grpc.CallOption) (*UserEditEMailResponse, error)
}

type accountClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountClient(cc grpc.ClientConnInterface) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) UserGet(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserGetResponse, error) {
	out := new(UserGetResponse)
	err := c.cc.Invoke(ctx, Account_UserGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	out := new(UserCreateResponse)
	err := c.cc.Invoke(ctx, Account_UserCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) UserUpdate(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserUpdateResponse, error) {
	out := new(UserUpdateResponse)
	err := c.cc.Invoke(ctx, Account_UserUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) UserEditPassword(ctx context.Context, in *UserEditPasswordRequest, opts ...grpc.CallOption) (*UserEditPasswordResponse, error) {
	out := new(UserEditPasswordResponse)
	err := c.cc.Invoke(ctx, Account_UserEditPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) UserEditUsername(ctx context.Context, in *UserEditUsernameRequest, opts ...grpc.CallOption) (*UserEditUsernameResponse, error) {
	out := new(UserEditUsernameResponse)
	err := c.cc.Invoke(ctx, Account_UserEditUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) UserEditEMail(ctx context.Context, in *UserEditEMailRequest, opts ...grpc.CallOption) (*UserEditEMailResponse, error) {
	out := new(UserEditEMailResponse)
	err := c.cc.Invoke(ctx, Account_UserEditEMail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
// All implementations should embed UnimplementedAccountServer
// for forward compatibility
type AccountServer interface {
	UserGet(context.Context, *emptypb.Empty) (*UserGetResponse, error)
	UserCreate(context.Context, *UserCreateRequest) (*UserCreateResponse, error)
	UserUpdate(context.Context, *UserUpdateRequest) (*UserUpdateResponse, error)
	UserEditPassword(context.Context, *UserEditPasswordRequest) (*UserEditPasswordResponse, error)
	UserEditUsername(context.Context, *UserEditUsernameRequest) (*UserEditUsernameResponse, error)
	UserEditEMail(context.Context, *UserEditEMailRequest) (*UserEditEMailResponse, error)
}

// UnimplementedAccountServer should be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (UnimplementedAccountServer) UserGet(context.Context, *emptypb.Empty) (*UserGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGet not implemented")
}
func (UnimplementedAccountServer) UserCreate(context.Context, *UserCreateRequest) (*UserCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCreate not implemented")
}
func (UnimplementedAccountServer) UserUpdate(context.Context, *UserUpdateRequest) (*UserUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdate not implemented")
}
func (UnimplementedAccountServer) UserEditPassword(context.Context, *UserEditPasswordRequest) (*UserEditPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserEditPassword not implemented")
}
func (UnimplementedAccountServer) UserEditUsername(context.Context, *UserEditUsernameRequest) (*UserEditUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserEditUsername not implemented")
}
func (UnimplementedAccountServer) UserEditEMail(context.Context, *UserEditEMailRequest) (*UserEditEMailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserEditEMail not implemented")
}

// UnsafeAccountServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServer will
// result in compilation errors.
type UnsafeAccountServer interface {
	mustEmbedUnimplementedAccountServer()
}

func RegisterAccountServer(s grpc.ServiceRegistrar, srv AccountServer) {
	s.RegisterService(&Account_ServiceDesc, srv)
}

func _Account_UserGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).UserGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_UserGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).UserGet(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_UserCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).UserCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_UserCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).UserCreate(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_UserUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).UserUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_UserUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).UserUpdate(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_UserEditPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEditPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).UserEditPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_UserEditPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).UserEditPassword(ctx, req.(*UserEditPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_UserEditUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEditUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).UserEditUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_UserEditUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).UserEditUsername(ctx, req.(*UserEditUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_UserEditEMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEditEMailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).UserEditEMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_UserEditEMail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).UserEditEMail(ctx, req.(*UserEditEMailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Account_ServiceDesc is the grpc.ServiceDesc for Account service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Account_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apis.proto.v1alpha1.auth.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserGet",
			Handler:    _Account_UserGet_Handler,
		},
		{
			MethodName: "UserCreate",
			Handler:    _Account_UserCreate_Handler,
		},
		{
			MethodName: "UserUpdate",
			Handler:    _Account_UserUpdate_Handler,
		},
		{
			MethodName: "UserEditPassword",
			Handler:    _Account_UserEditPassword_Handler,
		},
		{
			MethodName: "UserEditUsername",
			Handler:    _Account_UserEditUsername_Handler,
		},
		{
			MethodName: "UserEditEMail",
			Handler:    _Account_UserEditEMail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1alpha1/auth/account.proto",
}
