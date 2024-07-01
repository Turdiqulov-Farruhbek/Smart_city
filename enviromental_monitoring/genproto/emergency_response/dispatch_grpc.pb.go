// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: dispatch.proto

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
	ResourceDispatcheService_Create_FullMethodName  = "/emergency_response.ResourceDispatcheService/Create"
	ResourceDispatcheService_GetById_FullMethodName = "/emergency_response.ResourceDispatcheService/GetById"
	ResourceDispatcheService_GetAll_FullMethodName  = "/emergency_response.ResourceDispatcheService/GetAll"
	ResourceDispatcheService_Update_FullMethodName  = "/emergency_response.ResourceDispatcheService/Update"
	ResourceDispatcheService_Delete_FullMethodName  = "/emergency_response.ResourceDispatcheService/Delete"
)

// ResourceDispatcheServiceClient is the client API for ResourceDispatcheService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceDispatcheServiceClient interface {
	Create(ctx context.Context, in *DispatchesCreateReq, opts ...grpc.CallOption) (*DispatchesRes, error)
	GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*DispatchesFullRes, error)
	GetAll(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*DispatchesGetAllRes, error)
	Update(ctx context.Context, in *DispatchesUpdateReq, opts ...grpc.CallOption) (*DispatchesRes, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
}

type resourceDispatcheServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceDispatcheServiceClient(cc grpc.ClientConnInterface) ResourceDispatcheServiceClient {
	return &resourceDispatcheServiceClient{cc}
}

func (c *resourceDispatcheServiceClient) Create(ctx context.Context, in *DispatchesCreateReq, opts ...grpc.CallOption) (*DispatchesRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DispatchesRes)
	err := c.cc.Invoke(ctx, ResourceDispatcheService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceDispatcheServiceClient) GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*DispatchesFullRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DispatchesFullRes)
	err := c.cc.Invoke(ctx, ResourceDispatcheService_GetById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceDispatcheServiceClient) GetAll(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*DispatchesGetAllRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DispatchesGetAllRes)
	err := c.cc.Invoke(ctx, ResourceDispatcheService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceDispatcheServiceClient) Update(ctx context.Context, in *DispatchesUpdateReq, opts ...grpc.CallOption) (*DispatchesRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DispatchesRes)
	err := c.cc.Invoke(ctx, ResourceDispatcheService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceDispatcheServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, ResourceDispatcheService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceDispatcheServiceServer is the server API for ResourceDispatcheService service.
// All implementations must embed UnimplementedResourceDispatcheServiceServer
// for forward compatibility
type ResourceDispatcheServiceServer interface {
	Create(context.Context, *DispatchesCreateReq) (*DispatchesRes, error)
	GetById(context.Context, *ById) (*DispatchesFullRes, error)
	GetAll(context.Context, *Filter) (*DispatchesGetAllRes, error)
	Update(context.Context, *DispatchesUpdateReq) (*DispatchesRes, error)
	Delete(context.Context, *ById) (*Void, error)
	mustEmbedUnimplementedResourceDispatcheServiceServer()
}

// UnimplementedResourceDispatcheServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResourceDispatcheServiceServer struct {
}

func (UnimplementedResourceDispatcheServiceServer) Create(context.Context, *DispatchesCreateReq) (*DispatchesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedResourceDispatcheServiceServer) GetById(context.Context, *ById) (*DispatchesFullRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedResourceDispatcheServiceServer) GetAll(context.Context, *Filter) (*DispatchesGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedResourceDispatcheServiceServer) Update(context.Context, *DispatchesUpdateReq) (*DispatchesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedResourceDispatcheServiceServer) Delete(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedResourceDispatcheServiceServer) mustEmbedUnimplementedResourceDispatcheServiceServer() {
}

// UnsafeResourceDispatcheServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceDispatcheServiceServer will
// result in compilation errors.
type UnsafeResourceDispatcheServiceServer interface {
	mustEmbedUnimplementedResourceDispatcheServiceServer()
}

func RegisterResourceDispatcheServiceServer(s grpc.ServiceRegistrar, srv ResourceDispatcheServiceServer) {
	s.RegisterService(&ResourceDispatcheService_ServiceDesc, srv)
}

func _ResourceDispatcheService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatchesCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceDispatcheServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceDispatcheService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceDispatcheServiceServer).Create(ctx, req.(*DispatchesCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceDispatcheService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceDispatcheServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceDispatcheService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceDispatcheServiceServer).GetById(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceDispatcheService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceDispatcheServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceDispatcheService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceDispatcheServiceServer).GetAll(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceDispatcheService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatchesUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceDispatcheServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceDispatcheService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceDispatcheServiceServer).Update(ctx, req.(*DispatchesUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceDispatcheService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceDispatcheServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceDispatcheService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceDispatcheServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceDispatcheService_ServiceDesc is the grpc.ServiceDesc for ResourceDispatcheService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceDispatcheService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emergency_response.ResourceDispatcheService",
	HandlerType: (*ResourceDispatcheServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ResourceDispatcheService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _ResourceDispatcheService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ResourceDispatcheService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ResourceDispatcheService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ResourceDispatcheService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dispatch.proto",
}