// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: city_submodule/energy_management/energy_meter.proto

package energy_management

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
	EnenrgyMeterService_CreateEnergyMeter_FullMethodName = "/energy_management.EnenrgyMeterService/CreateEnergyMeter"
	EnenrgyMeterService_UpdateEnergyMeter_FullMethodName = "/energy_management.EnenrgyMeterService/UpdateEnergyMeter"
	EnenrgyMeterService_DeleteEnergyMeter_FullMethodName = "/energy_management.EnenrgyMeterService/DeleteEnergyMeter"
	EnenrgyMeterService_GetEnergyMeter_FullMethodName    = "/energy_management.EnenrgyMeterService/GetEnergyMeter"
)

// EnenrgyMeterServiceClient is the client API for EnenrgyMeterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnenrgyMeterServiceClient interface {
	CreateEnergyMeter(ctx context.Context, in *EnergyMeterCreate, opts ...grpc.CallOption) (*EnergyMeter, error)
	UpdateEnergyMeter(ctx context.Context, in *EnergyMeterUpdate, opts ...grpc.CallOption) (*EnergyMeter, error)
	DeleteEnergyMeter(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
	GetEnergyMeter(ctx context.Context, in *ById, opts ...grpc.CallOption) (*GetEnergyWithBuildings, error)
}

type enenrgyMeterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEnenrgyMeterServiceClient(cc grpc.ClientConnInterface) EnenrgyMeterServiceClient {
	return &enenrgyMeterServiceClient{cc}
}

func (c *enenrgyMeterServiceClient) CreateEnergyMeter(ctx context.Context, in *EnergyMeterCreate, opts ...grpc.CallOption) (*EnergyMeter, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EnergyMeter)
	err := c.cc.Invoke(ctx, EnenrgyMeterService_CreateEnergyMeter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enenrgyMeterServiceClient) UpdateEnergyMeter(ctx context.Context, in *EnergyMeterUpdate, opts ...grpc.CallOption) (*EnergyMeter, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EnergyMeter)
	err := c.cc.Invoke(ctx, EnenrgyMeterService_UpdateEnergyMeter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enenrgyMeterServiceClient) DeleteEnergyMeter(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, EnenrgyMeterService_DeleteEnergyMeter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enenrgyMeterServiceClient) GetEnergyMeter(ctx context.Context, in *ById, opts ...grpc.CallOption) (*GetEnergyWithBuildings, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEnergyWithBuildings)
	err := c.cc.Invoke(ctx, EnenrgyMeterService_GetEnergyMeter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnenrgyMeterServiceServer is the server API for EnenrgyMeterService service.
// All implementations must embed UnimplementedEnenrgyMeterServiceServer
// for forward compatibility
type EnenrgyMeterServiceServer interface {
	CreateEnergyMeter(context.Context, *EnergyMeterCreate) (*EnergyMeter, error)
	UpdateEnergyMeter(context.Context, *EnergyMeterUpdate) (*EnergyMeter, error)
	DeleteEnergyMeter(context.Context, *ById) (*Void, error)
	GetEnergyMeter(context.Context, *ById) (*GetEnergyWithBuildings, error)
	mustEmbedUnimplementedEnenrgyMeterServiceServer()
}

// UnimplementedEnenrgyMeterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEnenrgyMeterServiceServer struct {
}

func (UnimplementedEnenrgyMeterServiceServer) CreateEnergyMeter(context.Context, *EnergyMeterCreate) (*EnergyMeter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEnergyMeter not implemented")
}
func (UnimplementedEnenrgyMeterServiceServer) UpdateEnergyMeter(context.Context, *EnergyMeterUpdate) (*EnergyMeter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEnergyMeter not implemented")
}
func (UnimplementedEnenrgyMeterServiceServer) DeleteEnergyMeter(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEnergyMeter not implemented")
}
func (UnimplementedEnenrgyMeterServiceServer) GetEnergyMeter(context.Context, *ById) (*GetEnergyWithBuildings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEnergyMeter not implemented")
}
func (UnimplementedEnenrgyMeterServiceServer) mustEmbedUnimplementedEnenrgyMeterServiceServer() {}

// UnsafeEnenrgyMeterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnenrgyMeterServiceServer will
// result in compilation errors.
type UnsafeEnenrgyMeterServiceServer interface {
	mustEmbedUnimplementedEnenrgyMeterServiceServer()
}

func RegisterEnenrgyMeterServiceServer(s grpc.ServiceRegistrar, srv EnenrgyMeterServiceServer) {
	s.RegisterService(&EnenrgyMeterService_ServiceDesc, srv)
}

func _EnenrgyMeterService_CreateEnergyMeter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnergyMeterCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnenrgyMeterServiceServer).CreateEnergyMeter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnenrgyMeterService_CreateEnergyMeter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnenrgyMeterServiceServer).CreateEnergyMeter(ctx, req.(*EnergyMeterCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnenrgyMeterService_UpdateEnergyMeter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnergyMeterUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnenrgyMeterServiceServer).UpdateEnergyMeter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnenrgyMeterService_UpdateEnergyMeter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnenrgyMeterServiceServer).UpdateEnergyMeter(ctx, req.(*EnergyMeterUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnenrgyMeterService_DeleteEnergyMeter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnenrgyMeterServiceServer).DeleteEnergyMeter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnenrgyMeterService_DeleteEnergyMeter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnenrgyMeterServiceServer).DeleteEnergyMeter(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnenrgyMeterService_GetEnergyMeter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnenrgyMeterServiceServer).GetEnergyMeter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnenrgyMeterService_GetEnergyMeter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnenrgyMeterServiceServer).GetEnergyMeter(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

// EnenrgyMeterService_ServiceDesc is the grpc.ServiceDesc for EnenrgyMeterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EnenrgyMeterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "energy_management.EnenrgyMeterService",
	HandlerType: (*EnenrgyMeterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEnergyMeter",
			Handler:    _EnenrgyMeterService_CreateEnergyMeter_Handler,
		},
		{
			MethodName: "UpdateEnergyMeter",
			Handler:    _EnenrgyMeterService_UpdateEnergyMeter_Handler,
		},
		{
			MethodName: "DeleteEnergyMeter",
			Handler:    _EnenrgyMeterService_DeleteEnergyMeter_Handler,
		},
		{
			MethodName: "GetEnergyMeter",
			Handler:    _EnenrgyMeterService_GetEnergyMeter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "city_submodule/energy_management/energy_meter.proto",
}
