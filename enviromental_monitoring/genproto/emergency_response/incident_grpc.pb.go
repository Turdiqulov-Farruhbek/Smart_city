// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: incident.proto

package emergency_response

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
	EmergencyIncidentService_Create_FullMethodName  = "/emergency_response.EmergencyIncidentService/Create"
	EmergencyIncidentService_GetById_FullMethodName = "/emergency_response.EmergencyIncidentService/GetById"
	EmergencyIncidentService_GetAll_FullMethodName  = "/emergency_response.EmergencyIncidentService/GetAll"
	EmergencyIncidentService_Update_FullMethodName  = "/emergency_response.EmergencyIncidentService/Update"
	EmergencyIncidentService_Delete_FullMethodName  = "/emergency_response.EmergencyIncidentService/Delete"
)

// EmergencyIncidentServiceClient is the client API for EmergencyIncidentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmergencyIncidentServiceClient interface {
	Create(ctx context.Context, in *IncidentsCreateReq, opts ...grpc.CallOption) (*IncidentsRes, error)
	GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*IncidentsGetByIdRes, error)
	GetAll(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*IncidentsGetAllRes, error)
	Update(ctx context.Context, in *IncidentsUpdateReq, opts ...grpc.CallOption) (*IncidentsRes, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
}

type emergencyIncidentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmergencyIncidentServiceClient(cc grpc.ClientConnInterface) EmergencyIncidentServiceClient {
	return &emergencyIncidentServiceClient{cc}
}

func (c *emergencyIncidentServiceClient) Create(ctx context.Context, in *IncidentsCreateReq, opts ...grpc.CallOption) (*IncidentsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IncidentsRes)
	err := c.cc.Invoke(ctx, EmergencyIncidentService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emergencyIncidentServiceClient) GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*IncidentsGetByIdRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IncidentsGetByIdRes)
	err := c.cc.Invoke(ctx, EmergencyIncidentService_GetById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emergencyIncidentServiceClient) GetAll(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*IncidentsGetAllRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IncidentsGetAllRes)
	err := c.cc.Invoke(ctx, EmergencyIncidentService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emergencyIncidentServiceClient) Update(ctx context.Context, in *IncidentsUpdateReq, opts ...grpc.CallOption) (*IncidentsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IncidentsRes)
	err := c.cc.Invoke(ctx, EmergencyIncidentService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emergencyIncidentServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, EmergencyIncidentService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmergencyIncidentServiceServer is the server API for EmergencyIncidentService service.
// All implementations must embed UnimplementedEmergencyIncidentServiceServer
// for forward compatibility
type EmergencyIncidentServiceServer interface {
	Create(context.Context, *IncidentsCreateReq) (*IncidentsRes, error)
	GetById(context.Context, *ById) (*IncidentsGetByIdRes, error)
	GetAll(context.Context, *Filter) (*IncidentsGetAllRes, error)
	Update(context.Context, *IncidentsUpdateReq) (*IncidentsRes, error)
	Delete(context.Context, *ById) (*Void, error)
	mustEmbedUnimplementedEmergencyIncidentServiceServer()
}

// UnimplementedEmergencyIncidentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmergencyIncidentServiceServer struct {
}

func (UnimplementedEmergencyIncidentServiceServer) Create(context.Context, *IncidentsCreateReq) (*IncidentsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedEmergencyIncidentServiceServer) GetById(context.Context, *ById) (*IncidentsGetByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedEmergencyIncidentServiceServer) GetAll(context.Context, *Filter) (*IncidentsGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedEmergencyIncidentServiceServer) Update(context.Context, *IncidentsUpdateReq) (*IncidentsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedEmergencyIncidentServiceServer) Delete(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedEmergencyIncidentServiceServer) mustEmbedUnimplementedEmergencyIncidentServiceServer() {
}

// UnsafeEmergencyIncidentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmergencyIncidentServiceServer will
// result in compilation errors.
type UnsafeEmergencyIncidentServiceServer interface {
	mustEmbedUnimplementedEmergencyIncidentServiceServer()
}

func RegisterEmergencyIncidentServiceServer(s grpc.ServiceRegistrar, srv EmergencyIncidentServiceServer) {
	s.RegisterService(&EmergencyIncidentService_ServiceDesc, srv)
}

func _EmergencyIncidentService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentsCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmergencyIncidentServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmergencyIncidentService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmergencyIncidentServiceServer).Create(ctx, req.(*IncidentsCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmergencyIncidentService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmergencyIncidentServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmergencyIncidentService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmergencyIncidentServiceServer).GetById(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmergencyIncidentService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmergencyIncidentServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmergencyIncidentService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmergencyIncidentServiceServer).GetAll(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmergencyIncidentService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentsUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmergencyIncidentServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmergencyIncidentService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmergencyIncidentServiceServer).Update(ctx, req.(*IncidentsUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmergencyIncidentService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmergencyIncidentServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmergencyIncidentService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmergencyIncidentServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

// EmergencyIncidentService_ServiceDesc is the grpc.ServiceDesc for EmergencyIncidentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmergencyIncidentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emergency_response.EmergencyIncidentService",
	HandlerType: (*EmergencyIncidentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EmergencyIncidentService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _EmergencyIncidentService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _EmergencyIncidentService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EmergencyIncidentService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EmergencyIncidentService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "incident.proto",
}
