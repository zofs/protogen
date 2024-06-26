// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: accore/v1/user-service.proto

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/zofs/protogen/accore/pb"
	dtopb "github.com/zofs/protogen/dtopb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserService_Register_FullMethodName             = "/accore.v1.UserService/Register"
	UserService_Login_FullMethodName                = "/accore.v1.UserService/Login"
	UserService_GoogleOauth_FullMethodName          = "/accore.v1.UserService/GoogleOauth"
	UserService_GoogleOauthAuthorize_FullMethodName = "/accore.v1.UserService/GoogleOauthAuthorize"
	UserService_Update_FullMethodName               = "/accore.v1.UserService/Update"
	UserService_First_FullMethodName                = "/accore.v1.UserService/First"
	UserService_Me_FullMethodName                   = "/accore.v1.UserService/Me"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Register(ctx context.Context, in *pb.UserRegister, opts ...grpc.CallOption) (*pb.UserInfo, error)
	Login(ctx context.Context, in *pb.UserLogin, opts ...grpc.CallOption) (*pb.LoginInfo, error)
	GoogleOauth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	GoogleOauthAuthorize(ctx context.Context, in *pb.OauthCode, opts ...grpc.CallOption) (*empty.Empty, error)
	Update(ctx context.Context, in *pb.UserUpdate, opts ...grpc.CallOption) (*pb.UserInfo, error)
	First(ctx context.Context, in *dtopb.ID, opts ...grpc.CallOption) (*pb.UserInfo, error)
	Me(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*pb.UserInfo, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Register(ctx context.Context, in *pb.UserRegister, opts ...grpc.CallOption) (*pb.UserInfo, error) {
	out := new(pb.UserInfo)
	err := c.cc.Invoke(ctx, UserService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *pb.UserLogin, opts ...grpc.CallOption) (*pb.LoginInfo, error) {
	out := new(pb.LoginInfo)
	err := c.cc.Invoke(ctx, UserService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GoogleOauth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, UserService_GoogleOauth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GoogleOauthAuthorize(ctx context.Context, in *pb.OauthCode, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, UserService_GoogleOauthAuthorize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *pb.UserUpdate, opts ...grpc.CallOption) (*pb.UserInfo, error) {
	out := new(pb.UserInfo)
	err := c.cc.Invoke(ctx, UserService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) First(ctx context.Context, in *dtopb.ID, opts ...grpc.CallOption) (*pb.UserInfo, error) {
	out := new(pb.UserInfo)
	err := c.cc.Invoke(ctx, UserService_First_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Me(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*pb.UserInfo, error) {
	out := new(pb.UserInfo)
	err := c.cc.Invoke(ctx, UserService_Me_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Register(context.Context, *pb.UserRegister) (*pb.UserInfo, error)
	Login(context.Context, *pb.UserLogin) (*pb.LoginInfo, error)
	GoogleOauth(context.Context, *empty.Empty) (*empty.Empty, error)
	GoogleOauthAuthorize(context.Context, *pb.OauthCode) (*empty.Empty, error)
	Update(context.Context, *pb.UserUpdate) (*pb.UserInfo, error)
	First(context.Context, *dtopb.ID) (*pb.UserInfo, error)
	Me(context.Context, *empty.Empty) (*pb.UserInfo, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Register(context.Context, *pb.UserRegister) (*pb.UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *pb.UserLogin) (*pb.LoginInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) GoogleOauth(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoogleOauth not implemented")
}
func (UnimplementedUserServiceServer) GoogleOauthAuthorize(context.Context, *pb.OauthCode) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoogleOauthAuthorize not implemented")
}
func (UnimplementedUserServiceServer) Update(context.Context, *pb.UserUpdate) (*pb.UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserServiceServer) First(context.Context, *dtopb.ID) (*pb.UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method First not implemented")
}
func (UnimplementedUserServiceServer) Me(context.Context, *empty.Empty) (*pb.UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Me not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.UserRegister)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*pb.UserRegister))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.UserLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*pb.UserLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GoogleOauth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GoogleOauth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GoogleOauth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GoogleOauth(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GoogleOauthAuthorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.OauthCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GoogleOauthAuthorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GoogleOauthAuthorize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GoogleOauthAuthorize(ctx, req.(*pb.OauthCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.UserUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*pb.UserUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_First_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtopb.ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).First(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_First_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).First(ctx, req.(*dtopb.ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Me_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Me(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Me_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Me(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accore.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "GoogleOauth",
			Handler:    _UserService_GoogleOauth_Handler,
		},
		{
			MethodName: "GoogleOauthAuthorize",
			Handler:    _UserService_GoogleOauthAuthorize_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "First",
			Handler:    _UserService_First_Handler,
		},
		{
			MethodName: "Me",
			Handler:    _UserService_Me_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accore/v1/user-service.proto",
}
