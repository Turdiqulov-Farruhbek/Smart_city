// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: dispatch.proto

package emergency_response

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DispatchesCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncidentId   string `protobuf:"bytes,1,opt,name=incident_id,json=incidentId,proto3" json:"incident_id,omitempty"`
	ResourceId   string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	DispatchedAt string `protobuf:"bytes,3,opt,name=dispatched_at,json=dispatchedAt,proto3" json:"dispatched_at,omitempty"`
	ArrivedAt    string `protobuf:"bytes,4,opt,name=arrived_at,json=arrivedAt,proto3" json:"arrived_at,omitempty"`
}

func (x *DispatchesCreateReq) Reset() {
	*x = DispatchesCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchesCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchesCreateReq) ProtoMessage() {}

func (x *DispatchesCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_dispatch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchesCreateReq.ProtoReflect.Descriptor instead.
func (*DispatchesCreateReq) Descriptor() ([]byte, []int) {
	return file_dispatch_proto_rawDescGZIP(), []int{0}
}

func (x *DispatchesCreateReq) GetIncidentId() string {
	if x != nil {
		return x.IncidentId
	}
	return ""
}

func (x *DispatchesCreateReq) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *DispatchesCreateReq) GetDispatchedAt() string {
	if x != nil {
		return x.DispatchedAt
	}
	return ""
}

func (x *DispatchesCreateReq) GetArrivedAt() string {
	if x != nil {
		return x.ArrivedAt
	}
	return ""
}

type DispatchesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IncidentId   string `protobuf:"bytes,2,opt,name=incident_id,json=incidentId,proto3" json:"incident_id,omitempty"`
	ResourceId   string `protobuf:"bytes,3,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	DispatchedAt string `protobuf:"bytes,4,opt,name=dispatched_at,json=dispatchedAt,proto3" json:"dispatched_at,omitempty"`
	ArrivedAt    string `protobuf:"bytes,5,opt,name=arrived_at,json=arrivedAt,proto3" json:"arrived_at,omitempty"`
}

func (x *DispatchesRes) Reset() {
	*x = DispatchesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchesRes) ProtoMessage() {}

func (x *DispatchesRes) ProtoReflect() protoreflect.Message {
	mi := &file_dispatch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchesRes.ProtoReflect.Descriptor instead.
func (*DispatchesRes) Descriptor() ([]byte, []int) {
	return file_dispatch_proto_rawDescGZIP(), []int{1}
}

func (x *DispatchesRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DispatchesRes) GetIncidentId() string {
	if x != nil {
		return x.IncidentId
	}
	return ""
}

func (x *DispatchesRes) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *DispatchesRes) GetDispatchedAt() string {
	if x != nil {
		return x.DispatchedAt
	}
	return ""
}

func (x *DispatchesRes) GetArrivedAt() string {
	if x != nil {
		return x.ArrivedAt
	}
	return ""
}

type DispatchesFullRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DispatchedAt string        `protobuf:"bytes,2,opt,name=dispatched_at,json=dispatchedAt,proto3" json:"dispatched_at,omitempty"`
	ArrivedAt    string        `protobuf:"bytes,3,opt,name=arrived_at,json=arrivedAt,proto3" json:"arrived_at,omitempty"`
	Incident     *IncidentsRes `protobuf:"bytes,4,opt,name=incident,proto3" json:"incident,omitempty"`
	Resource     *ResourcesRes `protobuf:"bytes,5,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *DispatchesFullRes) Reset() {
	*x = DispatchesFullRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchesFullRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchesFullRes) ProtoMessage() {}

func (x *DispatchesFullRes) ProtoReflect() protoreflect.Message {
	mi := &file_dispatch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchesFullRes.ProtoReflect.Descriptor instead.
func (*DispatchesFullRes) Descriptor() ([]byte, []int) {
	return file_dispatch_proto_rawDescGZIP(), []int{2}
}

func (x *DispatchesFullRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DispatchesFullRes) GetDispatchedAt() string {
	if x != nil {
		return x.DispatchedAt
	}
	return ""
}

func (x *DispatchesFullRes) GetArrivedAt() string {
	if x != nil {
		return x.ArrivedAt
	}
	return ""
}

func (x *DispatchesFullRes) GetIncident() *IncidentsRes {
	if x != nil {
		return x.Incident
	}
	return nil
}

func (x *DispatchesFullRes) GetResource() *ResourcesRes {
	if x != nil {
		return x.Resource
	}
	return nil
}

type DispatchesGetByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dispatch *DispatchesRes `protobuf:"bytes,1,opt,name=dispatch,proto3" json:"dispatch,omitempty"`
}

func (x *DispatchesGetByIdRes) Reset() {
	*x = DispatchesGetByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatch_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchesGetByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchesGetByIdRes) ProtoMessage() {}

func (x *DispatchesGetByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_dispatch_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchesGetByIdRes.ProtoReflect.Descriptor instead.
func (*DispatchesGetByIdRes) Descriptor() ([]byte, []int) {
	return file_dispatch_proto_rawDescGZIP(), []int{3}
}

func (x *DispatchesGetByIdRes) GetDispatch() *DispatchesRes {
	if x != nil {
		return x.Dispatch
	}
	return nil
}

type DispatchesGetAllRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dispatches []*DispatchesRes `protobuf:"bytes,1,rep,name=dispatches,proto3" json:"dispatches,omitempty"`
}

func (x *DispatchesGetAllRes) Reset() {
	*x = DispatchesGetAllRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatch_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchesGetAllRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchesGetAllRes) ProtoMessage() {}

func (x *DispatchesGetAllRes) ProtoReflect() protoreflect.Message {
	mi := &file_dispatch_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchesGetAllRes.ProtoReflect.Descriptor instead.
func (*DispatchesGetAllRes) Descriptor() ([]byte, []int) {
	return file_dispatch_proto_rawDescGZIP(), []int{4}
}

func (x *DispatchesGetAllRes) GetDispatches() []*DispatchesRes {
	if x != nil {
		return x.Dispatches
	}
	return nil
}

type DispatchesUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Dispatch *DispatchesCreateReq `protobuf:"bytes,2,opt,name=dispatch,proto3" json:"dispatch,omitempty"`
}

func (x *DispatchesUpdateReq) Reset() {
	*x = DispatchesUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dispatch_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchesUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchesUpdateReq) ProtoMessage() {}

func (x *DispatchesUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_dispatch_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchesUpdateReq.ProtoReflect.Descriptor instead.
func (*DispatchesUpdateReq) Descriptor() ([]byte, []int) {
	return file_dispatch_proto_rawDescGZIP(), []int{5}
}

func (x *DispatchesUpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DispatchesUpdateReq) GetDispatch() *DispatchesCreateReq {
	if x != nil {
		return x.Dispatch
	}
	return nil
}

var File_dispatch_proto protoreflect.FileDescriptor

var file_dispatch_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x34, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x01, 0x0a, 0x13, 0x44, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x72, 0x69, 0x76, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x72, 0x69,
	0x76, 0x65, 0x64, 0x41, 0x74, 0x22, 0xa5, 0x01, 0x0a, 0x0d, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x63, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x72, 0x72, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x72, 0x69, 0x76, 0x65, 0x64, 0x41, 0x74, 0x22, 0xe3, 0x01,
	0x0a, 0x11, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x46, 0x75, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x72, 0x69,
	0x76, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72,
	0x72, 0x69, 0x76, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3c, 0x0a, 0x08, 0x69, 0x6e, 0x63, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6d, 0x65, 0x72,
	0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49,
	0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x52, 0x08, 0x69, 0x6e, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65,
	0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0x55, 0x0a, 0x14, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x08, 0x64,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x52, 0x08, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x22, 0x58, 0x0a, 0x13, 0x44, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x12, 0x41, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63,
	0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x73, 0x22, 0x6a, 0x0a, 0x13, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x08, 0x64,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x52, 0x08, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x32, 0xab, 0x03, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e,
	0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x21, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x18, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x28, 0x2e, 0x65, 0x6d, 0x65, 0x72,
	0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12,
	0x1a, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x27, 0x2e, 0x65, 0x6d,
	0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x27, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x65, 0x6d, 0x65, 0x72,
	0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3e,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x79,
	0x49, 0x64, 0x1a, 0x18, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x1d,
	0x5a, 0x1b, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6d, 0x65, 0x72, 0x67,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dispatch_proto_rawDescOnce sync.Once
	file_dispatch_proto_rawDescData = file_dispatch_proto_rawDesc
)

func file_dispatch_proto_rawDescGZIP() []byte {
	file_dispatch_proto_rawDescOnce.Do(func() {
		file_dispatch_proto_rawDescData = protoimpl.X.CompressGZIP(file_dispatch_proto_rawDescData)
	})
	return file_dispatch_proto_rawDescData
}

var file_dispatch_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_dispatch_proto_goTypes = []any{
	(*DispatchesCreateReq)(nil),  // 0: emergency_response.DispatchesCreateReq
	(*DispatchesRes)(nil),        // 1: emergency_response.DispatchesRes
	(*DispatchesFullRes)(nil),    // 2: emergency_response.DispatchesFullRes
	(*DispatchesGetByIdRes)(nil), // 3: emergency_response.DispatchesGetByIdRes
	(*DispatchesGetAllRes)(nil),  // 4: emergency_response.DispatchesGetAllRes
	(*DispatchesUpdateReq)(nil),  // 5: emergency_response.DispatchesUpdateReq
	(*IncidentsRes)(nil),         // 6: emergency_response.IncidentsRes
	(*ResourcesRes)(nil),         // 7: emergency_response.ResourcesRes
	(*ById)(nil),                 // 8: emergency_response.ById
	(*Filter)(nil),               // 9: emergency_response.Filter
	(*Void)(nil),                 // 10: emergency_response.Void
}
var file_dispatch_proto_depIdxs = []int32{
	6,  // 0: emergency_response.DispatchesFullRes.incident:type_name -> emergency_response.IncidentsRes
	7,  // 1: emergency_response.DispatchesFullRes.resource:type_name -> emergency_response.ResourcesRes
	1,  // 2: emergency_response.DispatchesGetByIdRes.dispatch:type_name -> emergency_response.DispatchesRes
	1,  // 3: emergency_response.DispatchesGetAllRes.dispatches:type_name -> emergency_response.DispatchesRes
	0,  // 4: emergency_response.DispatchesUpdateReq.dispatch:type_name -> emergency_response.DispatchesCreateReq
	0,  // 5: emergency_response.ResourceDispatchService.Create:input_type -> emergency_response.DispatchesCreateReq
	8,  // 6: emergency_response.ResourceDispatchService.GetById:input_type -> emergency_response.ById
	9,  // 7: emergency_response.ResourceDispatchService.GetAll:input_type -> emergency_response.Filter
	5,  // 8: emergency_response.ResourceDispatchService.Update:input_type -> emergency_response.DispatchesUpdateReq
	8,  // 9: emergency_response.ResourceDispatchService.Delete:input_type -> emergency_response.ById
	1,  // 10: emergency_response.ResourceDispatchService.Create:output_type -> emergency_response.DispatchesRes
	3,  // 11: emergency_response.ResourceDispatchService.GetById:output_type -> emergency_response.DispatchesGetByIdRes
	4,  // 12: emergency_response.ResourceDispatchService.GetAll:output_type -> emergency_response.DispatchesGetAllRes
	1,  // 13: emergency_response.ResourceDispatchService.Update:output_type -> emergency_response.DispatchesRes
	10, // 14: emergency_response.ResourceDispatchService.Delete:output_type -> emergency_response.Void
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_dispatch_proto_init() }
func file_dispatch_proto_init() {
	if File_dispatch_proto != nil {
		return
	}
	file_common4_proto_init()
	file_resource_proto_init()
	file_incident_res_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dispatch_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DispatchesCreateReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dispatch_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DispatchesRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dispatch_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DispatchesFullRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dispatch_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DispatchesGetByIdRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dispatch_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DispatchesGetAllRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dispatch_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DispatchesUpdateReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dispatch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dispatch_proto_goTypes,
		DependencyIndexes: file_dispatch_proto_depIdxs,
		MessageInfos:      file_dispatch_proto_msgTypes,
	}.Build()
	File_dispatch_proto = out.File
	file_dispatch_proto_rawDesc = nil
	file_dispatch_proto_goTypes = nil
	file_dispatch_proto_depIdxs = nil
}
