// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: proposal_feedback.proto

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
	ProposalFeedbackService_CreateProposalFeedback_FullMethodName = "/city_planning.ProposalFeedbackService/CreateProposalFeedback"
	ProposalFeedbackService_UpdateProposalFeedback_FullMethodName = "/city_planning.ProposalFeedbackService/UpdateProposalFeedback"
	ProposalFeedbackService_DeleteProposalFeedback_FullMethodName = "/city_planning.ProposalFeedbackService/DeleteProposalFeedback"
	ProposalFeedbackService_GetProposalFeedbacks_FullMethodName   = "/city_planning.ProposalFeedbackService/GetProposalFeedbacks"
)

// ProposalFeedbackServiceClient is the client API for ProposalFeedbackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProposalFeedbackServiceClient interface {
	CreateProposalFeedback(ctx context.Context, in *ProposalFeedbackCreate, opts ...grpc.CallOption) (*Void, error)
	UpdateProposalFeedback(ctx context.Context, in *ProposalFeedbackCreate, opts ...grpc.CallOption) (*Void, error)
	DeleteProposalFeedback(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
	GetProposalFeedbacks(ctx context.Context, in *FeedbackFilter, opts ...grpc.CallOption) (*ProposalFeedbackList, error)
}

type proposalFeedbackServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProposalFeedbackServiceClient(cc grpc.ClientConnInterface) ProposalFeedbackServiceClient {
	return &proposalFeedbackServiceClient{cc}
}

func (c *proposalFeedbackServiceClient) CreateProposalFeedback(ctx context.Context, in *ProposalFeedbackCreate, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, ProposalFeedbackService_CreateProposalFeedback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proposalFeedbackServiceClient) UpdateProposalFeedback(ctx context.Context, in *ProposalFeedbackCreate, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, ProposalFeedbackService_UpdateProposalFeedback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proposalFeedbackServiceClient) DeleteProposalFeedback(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, ProposalFeedbackService_DeleteProposalFeedback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proposalFeedbackServiceClient) GetProposalFeedbacks(ctx context.Context, in *FeedbackFilter, opts ...grpc.CallOption) (*ProposalFeedbackList, error) {
	out := new(ProposalFeedbackList)
	err := c.cc.Invoke(ctx, ProposalFeedbackService_GetProposalFeedbacks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProposalFeedbackServiceServer is the server API for ProposalFeedbackService service.
// All implementations must embed UnimplementedProposalFeedbackServiceServer
// for forward compatibility
type ProposalFeedbackServiceServer interface {
	CreateProposalFeedback(context.Context, *ProposalFeedbackCreate) (*Void, error)
	UpdateProposalFeedback(context.Context, *ProposalFeedbackCreate) (*Void, error)
	DeleteProposalFeedback(context.Context, *ById) (*Void, error)
	GetProposalFeedbacks(context.Context, *FeedbackFilter) (*ProposalFeedbackList, error)
	mustEmbedUnimplementedProposalFeedbackServiceServer()
}

// UnimplementedProposalFeedbackServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProposalFeedbackServiceServer struct {
}

func (UnimplementedProposalFeedbackServiceServer) CreateProposalFeedback(context.Context, *ProposalFeedbackCreate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProposalFeedback not implemented")
}
func (UnimplementedProposalFeedbackServiceServer) UpdateProposalFeedback(context.Context, *ProposalFeedbackCreate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProposalFeedback not implemented")
}
func (UnimplementedProposalFeedbackServiceServer) DeleteProposalFeedback(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProposalFeedback not implemented")
}
func (UnimplementedProposalFeedbackServiceServer) GetProposalFeedbacks(context.Context, *FeedbackFilter) (*ProposalFeedbackList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProposalFeedbacks not implemented")
}
func (UnimplementedProposalFeedbackServiceServer) mustEmbedUnimplementedProposalFeedbackServiceServer() {
}

// UnsafeProposalFeedbackServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProposalFeedbackServiceServer will
// result in compilation errors.
type UnsafeProposalFeedbackServiceServer interface {
	mustEmbedUnimplementedProposalFeedbackServiceServer()
}

func RegisterProposalFeedbackServiceServer(s grpc.ServiceRegistrar, srv ProposalFeedbackServiceServer) {
	s.RegisterService(&ProposalFeedbackService_ServiceDesc, srv)
}

func _ProposalFeedbackService_CreateProposalFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposalFeedbackCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProposalFeedbackServiceServer).CreateProposalFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProposalFeedbackService_CreateProposalFeedback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProposalFeedbackServiceServer).CreateProposalFeedback(ctx, req.(*ProposalFeedbackCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProposalFeedbackService_UpdateProposalFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposalFeedbackCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProposalFeedbackServiceServer).UpdateProposalFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProposalFeedbackService_UpdateProposalFeedback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProposalFeedbackServiceServer).UpdateProposalFeedback(ctx, req.(*ProposalFeedbackCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProposalFeedbackService_DeleteProposalFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProposalFeedbackServiceServer).DeleteProposalFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProposalFeedbackService_DeleteProposalFeedback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProposalFeedbackServiceServer).DeleteProposalFeedback(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProposalFeedbackService_GetProposalFeedbacks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedbackFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProposalFeedbackServiceServer).GetProposalFeedbacks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProposalFeedbackService_GetProposalFeedbacks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProposalFeedbackServiceServer).GetProposalFeedbacks(ctx, req.(*FeedbackFilter))
	}
	return interceptor(ctx, in, info, handler)
}

// ProposalFeedbackService_ServiceDesc is the grpc.ServiceDesc for ProposalFeedbackService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProposalFeedbackService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "city_planning.ProposalFeedbackService",
	HandlerType: (*ProposalFeedbackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProposalFeedback",
			Handler:    _ProposalFeedbackService_CreateProposalFeedback_Handler,
		},
		{
			MethodName: "UpdateProposalFeedback",
			Handler:    _ProposalFeedbackService_UpdateProposalFeedback_Handler,
		},
		{
			MethodName: "DeleteProposalFeedback",
			Handler:    _ProposalFeedbackService_DeleteProposalFeedback_Handler,
		},
		{
			MethodName: "GetProposalFeedbacks",
			Handler:    _ProposalFeedbackService_GetProposalFeedbacks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proposal_feedback.proto",
}