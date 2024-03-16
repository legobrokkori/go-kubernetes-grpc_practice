// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: service_simple_bank.proto

package pb

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

// SimplaBankClient is the client API for SimplaBank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimplaBankClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
}

type simplaBankClient struct {
	cc grpc.ClientConnInterface
}

func NewSimplaBankClient(cc grpc.ClientConnInterface) SimplaBankClient {
	return &simplaBankClient{cc}
}

func (c *simplaBankClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/pb.SimplaBank/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simplaBankClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, "/pb.SimplaBank/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimplaBankServer is the server API for SimplaBank service.
// All implementations must embed UnimplementedSimplaBankServer
// for forward compatibility
type SimplaBankServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	mustEmbedUnimplementedSimplaBankServer()
}

// UnimplementedSimplaBankServer must be embedded to have forward compatible implementations.
type UnimplementedSimplaBankServer struct {
}

func (UnimplementedSimplaBankServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedSimplaBankServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedSimplaBankServer) mustEmbedUnimplementedSimplaBankServer() {}

// UnsafeSimplaBankServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimplaBankServer will
// result in compilation errors.
type UnsafeSimplaBankServer interface {
	mustEmbedUnimplementedSimplaBankServer()
}

func RegisterSimplaBankServer(s grpc.ServiceRegistrar, srv SimplaBankServer) {
	s.RegisterService(&SimplaBank_ServiceDesc, srv)
}

func _SimplaBank_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimplaBankServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SimplaBank/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimplaBankServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimplaBank_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimplaBankServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SimplaBank/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimplaBankServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SimplaBank_ServiceDesc is the grpc.ServiceDesc for SimplaBank service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimplaBank_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SimplaBank",
	HandlerType: (*SimplaBankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _SimplaBank_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _SimplaBank_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_simple_bank.proto",
}
