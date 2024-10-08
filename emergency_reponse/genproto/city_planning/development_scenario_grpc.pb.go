// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: development_scenario.proto

package city_planning

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
	DevelopmentScenarioService_CreateScenario_FullMethodName = "/city_planning.DevelopmentScenarioService/CreateScenario"
	DevelopmentScenarioService_GetScenarios_FullMethodName   = "/city_planning.DevelopmentScenarioService/GetScenarios"
	DevelopmentScenarioService_UpdateScenario_FullMethodName = "/city_planning.DevelopmentScenarioService/UpdateScenario"
	DevelopmentScenarioService_DeleteScenario_FullMethodName = "/city_planning.DevelopmentScenarioService/DeleteScenario"
)

// DevelopmentScenarioServiceClient is the client API for DevelopmentScenarioService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DevelopmentScenarioServiceClient interface {
	CreateScenario(ctx context.Context, in *ScenarioCreate, opts ...grpc.CallOption) (*Void, error)
	GetScenarios(ctx context.Context, in *ScenarioFilter, opts ...grpc.CallOption) (*ScenarioList, error)
	UpdateScenario(ctx context.Context, in *ScenarioCreate, opts ...grpc.CallOption) (*Void, error)
	DeleteScenario(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
}

type developmentScenarioServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDevelopmentScenarioServiceClient(cc grpc.ClientConnInterface) DevelopmentScenarioServiceClient {
	return &developmentScenarioServiceClient{cc}
}

func (c *developmentScenarioServiceClient) CreateScenario(ctx context.Context, in *ScenarioCreate, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, DevelopmentScenarioService_CreateScenario_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developmentScenarioServiceClient) GetScenarios(ctx context.Context, in *ScenarioFilter, opts ...grpc.CallOption) (*ScenarioList, error) {
	out := new(ScenarioList)
	err := c.cc.Invoke(ctx, DevelopmentScenarioService_GetScenarios_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developmentScenarioServiceClient) UpdateScenario(ctx context.Context, in *ScenarioCreate, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, DevelopmentScenarioService_UpdateScenario_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *developmentScenarioServiceClient) DeleteScenario(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, DevelopmentScenarioService_DeleteScenario_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevelopmentScenarioServiceServer is the server API for DevelopmentScenarioService service.
// All implementations must embed UnimplementedDevelopmentScenarioServiceServer
// for forward compatibility
type DevelopmentScenarioServiceServer interface {
	CreateScenario(context.Context, *ScenarioCreate) (*Void, error)
	GetScenarios(context.Context, *ScenarioFilter) (*ScenarioList, error)
	UpdateScenario(context.Context, *ScenarioCreate) (*Void, error)
	DeleteScenario(context.Context, *ById) (*Void, error)
	mustEmbedUnimplementedDevelopmentScenarioServiceServer()
}

// UnimplementedDevelopmentScenarioServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDevelopmentScenarioServiceServer struct {
}

func (UnimplementedDevelopmentScenarioServiceServer) CreateScenario(context.Context, *ScenarioCreate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateScenario not implemented")
}
func (UnimplementedDevelopmentScenarioServiceServer) GetScenarios(context.Context, *ScenarioFilter) (*ScenarioList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScenarios not implemented")
}
func (UnimplementedDevelopmentScenarioServiceServer) UpdateScenario(context.Context, *ScenarioCreate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScenario not implemented")
}
func (UnimplementedDevelopmentScenarioServiceServer) DeleteScenario(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteScenario not implemented")
}
func (UnimplementedDevelopmentScenarioServiceServer) mustEmbedUnimplementedDevelopmentScenarioServiceServer() {
}

// UnsafeDevelopmentScenarioServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DevelopmentScenarioServiceServer will
// result in compilation errors.
type UnsafeDevelopmentScenarioServiceServer interface {
	mustEmbedUnimplementedDevelopmentScenarioServiceServer()
}

func RegisterDevelopmentScenarioServiceServer(s grpc.ServiceRegistrar, srv DevelopmentScenarioServiceServer) {
	s.RegisterService(&DevelopmentScenarioService_ServiceDesc, srv)
}

func _DevelopmentScenarioService_CreateScenario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScenarioCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevelopmentScenarioServiceServer).CreateScenario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DevelopmentScenarioService_CreateScenario_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevelopmentScenarioServiceServer).CreateScenario(ctx, req.(*ScenarioCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevelopmentScenarioService_GetScenarios_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScenarioFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevelopmentScenarioServiceServer).GetScenarios(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DevelopmentScenarioService_GetScenarios_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevelopmentScenarioServiceServer).GetScenarios(ctx, req.(*ScenarioFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevelopmentScenarioService_UpdateScenario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScenarioCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevelopmentScenarioServiceServer).UpdateScenario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DevelopmentScenarioService_UpdateScenario_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevelopmentScenarioServiceServer).UpdateScenario(ctx, req.(*ScenarioCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevelopmentScenarioService_DeleteScenario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevelopmentScenarioServiceServer).DeleteScenario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DevelopmentScenarioService_DeleteScenario_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevelopmentScenarioServiceServer).DeleteScenario(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

// DevelopmentScenarioService_ServiceDesc is the grpc.ServiceDesc for DevelopmentScenarioService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DevelopmentScenarioService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "city_planning.DevelopmentScenarioService",
	HandlerType: (*DevelopmentScenarioServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateScenario",
			Handler:    _DevelopmentScenarioService_CreateScenario_Handler,
		},
		{
			MethodName: "GetScenarios",
			Handler:    _DevelopmentScenarioService_GetScenarios_Handler,
		},
		{
			MethodName: "UpdateScenario",
			Handler:    _DevelopmentScenarioService_UpdateScenario_Handler,
		},
		{
			MethodName: "DeleteScenario",
			Handler:    _DevelopmentScenarioService_DeleteScenario_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "development_scenario.proto",
}
