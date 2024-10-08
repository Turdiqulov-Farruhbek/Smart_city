// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: facility.proto

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
	EmergencyFacilityService_Create_FullMethodName  = "/emergency_response.EmergencyFacilityService/Create"
	EmergencyFacilityService_GetById_FullMethodName = "/emergency_response.EmergencyFacilityService/GetById"
	EmergencyFacilityService_GetAll_FullMethodName  = "/emergency_response.EmergencyFacilityService/GetAll"
	EmergencyFacilityService_Update_FullMethodName  = "/emergency_response.EmergencyFacilityService/Update"
	EmergencyFacilityService_Delete_FullMethodName  = "/emergency_response.EmergencyFacilityService/Delete"
)

// EmergencyFacilityServiceClient is the client API for EmergencyFacilityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmergencyFacilityServiceClient interface {
	Create(ctx context.Context, in *FacilitiesCreateReq, opts ...grpc.CallOption) (*FacilitiesRes, error)
	GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*FacilitiesGetByIdRes, error)
	GetAll(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*FacilitiesGetAllRes, error)
	Update(ctx context.Context, in *FacilitiesUpdateReq, opts ...grpc.CallOption) (*FacilitiesRes, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
}

type emergencyFacilityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmergencyFacilityServiceClient(cc grpc.ClientConnInterface) EmergencyFacilityServiceClient {
	return &emergencyFacilityServiceClient{cc}
}

func (c *emergencyFacilityServiceClient) Create(ctx context.Context, in *FacilitiesCreateReq, opts ...grpc.CallOption) (*FacilitiesRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FacilitiesRes)
	err := c.cc.Invoke(ctx, EmergencyFacilityService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emergencyFacilityServiceClient) GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*FacilitiesGetByIdRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FacilitiesGetByIdRes)
	err := c.cc.Invoke(ctx, EmergencyFacilityService_GetById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emergencyFacilityServiceClient) GetAll(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*FacilitiesGetAllRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FacilitiesGetAllRes)
	err := c.cc.Invoke(ctx, EmergencyFacilityService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emergencyFacilityServiceClient) Update(ctx context.Context, in *FacilitiesUpdateReq, opts ...grpc.CallOption) (*FacilitiesRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FacilitiesRes)
	err := c.cc.Invoke(ctx, EmergencyFacilityService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emergencyFacilityServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, EmergencyFacilityService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmergencyFacilityServiceServer is the server API for EmergencyFacilityService service.
// All implementations must embed UnimplementedEmergencyFacilityServiceServer
// for forward compatibility
type EmergencyFacilityServiceServer interface {
	Create(context.Context, *FacilitiesCreateReq) (*FacilitiesRes, error)
	GetById(context.Context, *ById) (*FacilitiesGetByIdRes, error)
	GetAll(context.Context, *Filter) (*FacilitiesGetAllRes, error)
	Update(context.Context, *FacilitiesUpdateReq) (*FacilitiesRes, error)
	Delete(context.Context, *ById) (*Void, error)
	mustEmbedUnimplementedEmergencyFacilityServiceServer()
}

// UnimplementedEmergencyFacilityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmergencyFacilityServiceServer struct {
}

func (UnimplementedEmergencyFacilityServiceServer) Create(context.Context, *FacilitiesCreateReq) (*FacilitiesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedEmergencyFacilityServiceServer) GetById(context.Context, *ById) (*FacilitiesGetByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedEmergencyFacilityServiceServer) GetAll(context.Context, *Filter) (*FacilitiesGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedEmergencyFacilityServiceServer) Update(context.Context, *FacilitiesUpdateReq) (*FacilitiesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedEmergencyFacilityServiceServer) Delete(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedEmergencyFacilityServiceServer) mustEmbedUnimplementedEmergencyFacilityServiceServer() {
}

// UnsafeEmergencyFacilityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmergencyFacilityServiceServer will
// result in compilation errors.
type UnsafeEmergencyFacilityServiceServer interface {
	mustEmbedUnimplementedEmergencyFacilityServiceServer()
}

func RegisterEmergencyFacilityServiceServer(s grpc.ServiceRegistrar, srv EmergencyFacilityServiceServer) {
	s.RegisterService(&EmergencyFacilityService_ServiceDesc, srv)
}

func _EmergencyFacilityService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FacilitiesCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmergencyFacilityServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmergencyFacilityService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmergencyFacilityServiceServer).Create(ctx, req.(*FacilitiesCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmergencyFacilityService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmergencyFacilityServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmergencyFacilityService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmergencyFacilityServiceServer).GetById(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmergencyFacilityService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmergencyFacilityServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmergencyFacilityService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmergencyFacilityServiceServer).GetAll(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmergencyFacilityService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FacilitiesUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmergencyFacilityServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmergencyFacilityService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmergencyFacilityServiceServer).Update(ctx, req.(*FacilitiesUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmergencyFacilityService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmergencyFacilityServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmergencyFacilityService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmergencyFacilityServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

// EmergencyFacilityService_ServiceDesc is the grpc.ServiceDesc for EmergencyFacilityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmergencyFacilityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emergency_response.EmergencyFacilityService",
	HandlerType: (*EmergencyFacilityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EmergencyFacilityService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _EmergencyFacilityService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _EmergencyFacilityService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EmergencyFacilityService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EmergencyFacilityService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "facility.proto",
}
