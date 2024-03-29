// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: auth_service.proto

package auth_service_pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	// Upon succesful login user receives a refresh_token.
	// When it expires or is revoked user has to login again.
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	// SignOut revokes user's active refresh_token.
	SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetAccessToken(ctx context.Context, in *GetAccessTokenRequest, opts ...grpc.CallOption) (*GetAccessTokenResponse, error)
	TranslateAccessToken(ctx context.Context, in *TranslateAccessTokenRequest, opts ...grpc.CallOption) (*TranslateAccessTokenResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/AuthService/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/AuthService/SignOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetAccessToken(ctx context.Context, in *GetAccessTokenRequest, opts ...grpc.CallOption) (*GetAccessTokenResponse, error) {
	out := new(GetAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/AuthService/GetAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) TranslateAccessToken(ctx context.Context, in *TranslateAccessTokenRequest, opts ...grpc.CallOption) (*TranslateAccessTokenResponse, error) {
	out := new(TranslateAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/AuthService/TranslateAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	// Upon succesful login user receives a refresh_token.
	// When it expires or is revoked user has to login again.
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	// SignOut revokes user's active refresh_token.
	SignOut(context.Context, *SignOutRequest) (*empty.Empty, error)
	GetAccessToken(context.Context, *GetAccessTokenRequest) (*GetAccessTokenResponse, error)
	TranslateAccessToken(context.Context, *TranslateAccessTokenRequest) (*TranslateAccessTokenResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedAuthServiceServer) SignOut(context.Context, *SignOutRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
func (UnimplementedAuthServiceServer) GetAccessToken(context.Context, *GetAccessTokenRequest) (*GetAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessToken not implemented")
}
func (UnimplementedAuthServiceServer) TranslateAccessToken(context.Context, *TranslateAccessTokenRequest) (*TranslateAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranslateAccessToken not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SignOut(ctx, req.(*SignOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/GetAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetAccessToken(ctx, req.(*GetAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_TranslateAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).TranslateAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthService/TranslateAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).TranslateAccessToken(ctx, req.(*TranslateAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _AuthService_SignIn_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _AuthService_SignOut_Handler,
		},
		{
			MethodName: "GetAccessToken",
			Handler:    _AuthService_GetAccessToken_Handler,
		},
		{
			MethodName: "TranslateAccessToken",
			Handler:    _AuthService_TranslateAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_service.proto",
}
