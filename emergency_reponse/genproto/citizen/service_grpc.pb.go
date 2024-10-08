// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: service.proto

package citizen

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
	CitizenRequestService_CreateServiceRequest_FullMethodName = "/citizen.CitizenRequestService/CreateServiceRequest"
	CitizenRequestService_GetCitizenRequests_FullMethodName   = "/citizen.CitizenRequestService/GetCitizenRequests"
	CitizenRequestService_UpdateServiceRequest_FullMethodName = "/citizen.CitizenRequestService/UpdateServiceRequest"
	CitizenRequestService_DeleteServiceRequest_FullMethodName = "/citizen.CitizenRequestService/DeleteServiceRequest"
)

// CitizenRequestServiceClient is the client API for CitizenRequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CitizenRequestServiceClient interface {
	CreateServiceRequest(ctx context.Context, in *ServiceReqCreate, opts ...grpc.CallOption) (*ServiceReq, error)
	GetCitizenRequests(ctx context.Context, in *ById, opts ...grpc.CallOption) (*ServiceReqList, error)
	UpdateServiceRequest(ctx context.Context, in *ServiceReqCreate, opts ...grpc.CallOption) (*ServiceReq, error)
	DeleteServiceRequest(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
}

type citizenRequestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCitizenRequestServiceClient(cc grpc.ClientConnInterface) CitizenRequestServiceClient {
	return &citizenRequestServiceClient{cc}
}

func (c *citizenRequestServiceClient) CreateServiceRequest(ctx context.Context, in *ServiceReqCreate, opts ...grpc.CallOption) (*ServiceReq, error) {
	out := new(ServiceReq)
	err := c.cc.Invoke(ctx, CitizenRequestService_CreateServiceRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *citizenRequestServiceClient) GetCitizenRequests(ctx context.Context, in *ById, opts ...grpc.CallOption) (*ServiceReqList, error) {
	out := new(ServiceReqList)
	err := c.cc.Invoke(ctx, CitizenRequestService_GetCitizenRequests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *citizenRequestServiceClient) UpdateServiceRequest(ctx context.Context, in *ServiceReqCreate, opts ...grpc.CallOption) (*ServiceReq, error) {
	out := new(ServiceReq)
	err := c.cc.Invoke(ctx, CitizenRequestService_UpdateServiceRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *citizenRequestServiceClient) DeleteServiceRequest(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, CitizenRequestService_DeleteServiceRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CitizenRequestServiceServer is the server API for CitizenRequestService service.
// All implementations must embed UnimplementedCitizenRequestServiceServer
// for forward compatibility
type CitizenRequestServiceServer interface {
	CreateServiceRequest(context.Context, *ServiceReqCreate) (*ServiceReq, error)
	GetCitizenRequests(context.Context, *ById) (*ServiceReqList, error)
	UpdateServiceRequest(context.Context, *ServiceReqCreate) (*ServiceReq, error)
	DeleteServiceRequest(context.Context, *ById) (*Void, error)
	mustEmbedUnimplementedCitizenRequestServiceServer()
}

// UnimplementedCitizenRequestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCitizenRequestServiceServer struct {
}

func (UnimplementedCitizenRequestServiceServer) CreateServiceRequest(context.Context, *ServiceReqCreate) (*ServiceReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServiceRequest not implemented")
}
func (UnimplementedCitizenRequestServiceServer) GetCitizenRequests(context.Context, *ById) (*ServiceReqList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCitizenRequests not implemented")
}
func (UnimplementedCitizenRequestServiceServer) UpdateServiceRequest(context.Context, *ServiceReqCreate) (*ServiceReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServiceRequest not implemented")
}
func (UnimplementedCitizenRequestServiceServer) DeleteServiceRequest(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServiceRequest not implemented")
}
func (UnimplementedCitizenRequestServiceServer) mustEmbedUnimplementedCitizenRequestServiceServer() {}

// UnsafeCitizenRequestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CitizenRequestServiceServer will
// result in compilation errors.
type UnsafeCitizenRequestServiceServer interface {
	mustEmbedUnimplementedCitizenRequestServiceServer()
}

func RegisterCitizenRequestServiceServer(s grpc.ServiceRegistrar, srv CitizenRequestServiceServer) {
	s.RegisterService(&CitizenRequestService_ServiceDesc, srv)
}

func _CitizenRequestService_CreateServiceRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceReqCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CitizenRequestServiceServer).CreateServiceRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CitizenRequestService_CreateServiceRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CitizenRequestServiceServer).CreateServiceRequest(ctx, req.(*ServiceReqCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _CitizenRequestService_GetCitizenRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CitizenRequestServiceServer).GetCitizenRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CitizenRequestService_GetCitizenRequests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CitizenRequestServiceServer).GetCitizenRequests(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _CitizenRequestService_UpdateServiceRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceReqCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CitizenRequestServiceServer).UpdateServiceRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CitizenRequestService_UpdateServiceRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CitizenRequestServiceServer).UpdateServiceRequest(ctx, req.(*ServiceReqCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _CitizenRequestService_DeleteServiceRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CitizenRequestServiceServer).DeleteServiceRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CitizenRequestService_DeleteServiceRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CitizenRequestServiceServer).DeleteServiceRequest(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

// CitizenRequestService_ServiceDesc is the grpc.ServiceDesc for CitizenRequestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CitizenRequestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "citizen.CitizenRequestService",
	HandlerType: (*CitizenRequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateServiceRequest",
			Handler:    _CitizenRequestService_CreateServiceRequest_Handler,
		},
		{
			MethodName: "GetCitizenRequests",
			Handler:    _CitizenRequestService_GetCitizenRequests_Handler,
		},
		{
			MethodName: "UpdateServiceRequest",
			Handler:    _CitizenRequestService_UpdateServiceRequest_Handler,
		},
		{
			MethodName: "DeleteServiceRequest",
			Handler:    _CitizenRequestService_DeleteServiceRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
