// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.14.0
// source: user-service.proto

package upbapi

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
	UserService_CreateUser_FullMethodName      = "/upbapi.UserService/CreateUser"
	UserService_GetUser_FullMethodName         = "/upbapi.UserService/GetUser"
	UserService_CreateInterest_FullMethodName  = "/upbapi.UserService/CreateInterest"
	UserService_GetInterest_FullMethodName     = "/upbapi.UserService/GetInterest"
	UserService_GetAllInterests_FullMethodName = "/upbapi.UserService/GetAllInterests"
	UserService_CreateCity_FullMethodName      = "/upbapi.UserService/CreateCity"
	UserService_GetCity_FullMethodName         = "/upbapi.UserService/GetCity"
	UserService_GetAllCities_FullMethodName    = "/upbapi.UserService/GetAllCities"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	CreateInterest(ctx context.Context, in *CreateInterestRequest, opts ...grpc.CallOption) (*CreateInterestResponse, error)
	GetInterest(ctx context.Context, in *GetInterestRequest, opts ...grpc.CallOption) (*GetInterestResponse, error)
	GetAllInterests(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllInterestsResponse, error)
	CreateCity(ctx context.Context, in *CreateCityRequest, opts ...grpc.CallOption) (*CreateCityResponse, error)
	GetCity(ctx context.Context, in *GetCityRequest, opts ...grpc.CallOption) (*GetCityResponse, error)
	GetAllCities(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllCitiesResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, UserService_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateInterest(ctx context.Context, in *CreateInterestRequest, opts ...grpc.CallOption) (*CreateInterestResponse, error) {
	out := new(CreateInterestResponse)
	err := c.cc.Invoke(ctx, UserService_CreateInterest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetInterest(ctx context.Context, in *GetInterestRequest, opts ...grpc.CallOption) (*GetInterestResponse, error) {
	out := new(GetInterestResponse)
	err := c.cc.Invoke(ctx, UserService_GetInterest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllInterests(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllInterestsResponse, error) {
	out := new(GetAllInterestsResponse)
	err := c.cc.Invoke(ctx, UserService_GetAllInterests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateCity(ctx context.Context, in *CreateCityRequest, opts ...grpc.CallOption) (*CreateCityResponse, error) {
	out := new(CreateCityResponse)
	err := c.cc.Invoke(ctx, UserService_CreateCity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCity(ctx context.Context, in *GetCityRequest, opts ...grpc.CallOption) (*GetCityResponse, error) {
	out := new(GetCityResponse)
	err := c.cc.Invoke(ctx, UserService_GetCity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllCities(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllCitiesResponse, error) {
	out := new(GetAllCitiesResponse)
	err := c.cc.Invoke(ctx, UserService_GetAllCities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	CreateInterest(context.Context, *CreateInterestRequest) (*CreateInterestResponse, error)
	GetInterest(context.Context, *GetInterestRequest) (*GetInterestResponse, error)
	GetAllInterests(context.Context, *Empty) (*GetAllInterestsResponse, error)
	CreateCity(context.Context, *CreateCityRequest) (*CreateCityResponse, error)
	GetCity(context.Context, *GetCityRequest) (*GetCityResponse, error)
	GetAllCities(context.Context, *Empty) (*GetAllCitiesResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) CreateInterest(context.Context, *CreateInterestRequest) (*CreateInterestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInterest not implemented")
}
func (UnimplementedUserServiceServer) GetInterest(context.Context, *GetInterestRequest) (*GetInterestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInterest not implemented")
}
func (UnimplementedUserServiceServer) GetAllInterests(context.Context, *Empty) (*GetAllInterestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllInterests not implemented")
}
func (UnimplementedUserServiceServer) CreateCity(context.Context, *CreateCityRequest) (*CreateCityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCity not implemented")
}
func (UnimplementedUserServiceServer) GetCity(context.Context, *GetCityRequest) (*GetCityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCity not implemented")
}
func (UnimplementedUserServiceServer) GetAllCities(context.Context, *Empty) (*GetAllCitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCities not implemented")
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

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateInterest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateInterest(ctx, req.(*CreateInterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetInterest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetInterest(ctx, req.(*GetInterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllInterests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllInterests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetAllInterests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllInterests(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateCity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateCity(ctx, req.(*CreateCityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetCity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCity(ctx, req.(*GetCityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllCities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllCities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetAllCities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllCities(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "upbapi.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "CreateInterest",
			Handler:    _UserService_CreateInterest_Handler,
		},
		{
			MethodName: "GetInterest",
			Handler:    _UserService_GetInterest_Handler,
		},
		{
			MethodName: "GetAllInterests",
			Handler:    _UserService_GetAllInterests_Handler,
		},
		{
			MethodName: "CreateCity",
			Handler:    _UserService_CreateCity_Handler,
		},
		{
			MethodName: "GetCity",
			Handler:    _UserService_GetCity_Handler,
		},
		{
			MethodName: "GetAllCities",
			Handler:    _UserService_GetAllCities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user-service.proto",
}
