// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: recycling_center.proto

package enviromental_monitoring

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

type RecyclingCenterCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CenterId          string `protobuf:"bytes,1,opt,name=CenterId,proto3" json:"CenterId,omitempty"`
	CenterName        string `protobuf:"bytes,2,opt,name=CenterName,proto3" json:"CenterName,omitempty"`
	Address           string `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	AcceptedMaterials string `protobuf:"bytes,4,opt,name=AcceptedMaterials,proto3" json:"AcceptedMaterials,omitempty"`
}

func (x *RecyclingCenterCreate) Reset() {
	*x = RecyclingCenterCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recycling_center_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecyclingCenterCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecyclingCenterCreate) ProtoMessage() {}

func (x *RecyclingCenterCreate) ProtoReflect() protoreflect.Message {
	mi := &file_recycling_center_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecyclingCenterCreate.ProtoReflect.Descriptor instead.
func (*RecyclingCenterCreate) Descriptor() ([]byte, []int) {
	return file_recycling_center_proto_rawDescGZIP(), []int{0}
}

func (x *RecyclingCenterCreate) GetCenterId() string {
	if x != nil {
		return x.CenterId
	}
	return ""
}

func (x *RecyclingCenterCreate) GetCenterName() string {
	if x != nil {
		return x.CenterName
	}
	return ""
}

func (x *RecyclingCenterCreate) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RecyclingCenterCreate) GetAcceptedMaterials() string {
	if x != nil {
		return x.AcceptedMaterials
	}
	return ""
}

type RecyclingCenter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CenterId          string `protobuf:"bytes,1,opt,name=CenterId,proto3" json:"CenterId,omitempty"`
	CenterName        string `protobuf:"bytes,2,opt,name=CenterName,proto3" json:"CenterName,omitempty"`
	Address           string `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	AcceptedMaterials string `protobuf:"bytes,4,opt,name=AcceptedMaterials,proto3" json:"AcceptedMaterials,omitempty"`
	CreatedAt         string `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt         string `protobuf:"bytes,6,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt         string `protobuf:"bytes,7,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
}

func (x *RecyclingCenter) Reset() {
	*x = RecyclingCenter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recycling_center_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecyclingCenter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecyclingCenter) ProtoMessage() {}

func (x *RecyclingCenter) ProtoReflect() protoreflect.Message {
	mi := &file_recycling_center_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecyclingCenter.ProtoReflect.Descriptor instead.
func (*RecyclingCenter) Descriptor() ([]byte, []int) {
	return file_recycling_center_proto_rawDescGZIP(), []int{1}
}

func (x *RecyclingCenter) GetCenterId() string {
	if x != nil {
		return x.CenterId
	}
	return ""
}

func (x *RecyclingCenter) GetCenterName() string {
	if x != nil {
		return x.CenterName
	}
	return ""
}

func (x *RecyclingCenter) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RecyclingCenter) GetAcceptedMaterials() string {
	if x != nil {
		return x.AcceptedMaterials
	}
	return ""
}

func (x *RecyclingCenter) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *RecyclingCenter) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *RecyclingCenter) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type RecyclingCenterList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecyclingCenters []*RecyclingCenter `protobuf:"bytes,1,rep,name=RecyclingCenters,proto3" json:"RecyclingCenters,omitempty"`
}

func (x *RecyclingCenterList) Reset() {
	*x = RecyclingCenterList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recycling_center_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecyclingCenterList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecyclingCenterList) ProtoMessage() {}

func (x *RecyclingCenterList) ProtoReflect() protoreflect.Message {
	mi := &file_recycling_center_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecyclingCenterList.ProtoReflect.Descriptor instead.
func (*RecyclingCenterList) Descriptor() ([]byte, []int) {
	return file_recycling_center_proto_rawDescGZIP(), []int{2}
}

func (x *RecyclingCenterList) GetRecyclingCenters() []*RecyclingCenter {
	if x != nil {
		return x.RecyclingCenters
	}
	return nil
}

type RecyclingCenterFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CenterId          string `protobuf:"bytes,1,opt,name=CenterId,proto3" json:"CenterId,omitempty"`
	CenterName        string `protobuf:"bytes,2,opt,name=CenterName,proto3" json:"CenterName,omitempty"`
	Address           string `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	AcceptedMaterials string `protobuf:"bytes,4,opt,name=AcceptedMaterials,proto3" json:"AcceptedMaterials,omitempty"`
}

func (x *RecyclingCenterFilter) Reset() {
	*x = RecyclingCenterFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recycling_center_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecyclingCenterFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecyclingCenterFilter) ProtoMessage() {}

func (x *RecyclingCenterFilter) ProtoReflect() protoreflect.Message {
	mi := &file_recycling_center_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecyclingCenterFilter.ProtoReflect.Descriptor instead.
func (*RecyclingCenterFilter) Descriptor() ([]byte, []int) {
	return file_recycling_center_proto_rawDescGZIP(), []int{3}
}

func (x *RecyclingCenterFilter) GetCenterId() string {
	if x != nil {
		return x.CenterId
	}
	return ""
}

func (x *RecyclingCenterFilter) GetCenterName() string {
	if x != nil {
		return x.CenterName
	}
	return ""
}

func (x *RecyclingCenterFilter) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RecyclingCenterFilter) GetAcceptedMaterials() string {
	if x != nil {
		return x.AcceptedMaterials
	}
	return ""
}

type RecyclingCenterModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CenterName        string `protobuf:"bytes,1,opt,name=CenterName,proto3" json:"CenterName,omitempty"`
	Address           string `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	AcceptedMaterials string `protobuf:"bytes,3,opt,name=AcceptedMaterials,proto3" json:"AcceptedMaterials,omitempty"`
}

func (x *RecyclingCenterModel) Reset() {
	*x = RecyclingCenterModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recycling_center_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecyclingCenterModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecyclingCenterModel) ProtoMessage() {}

func (x *RecyclingCenterModel) ProtoReflect() protoreflect.Message {
	mi := &file_recycling_center_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecyclingCenterModel.ProtoReflect.Descriptor instead.
func (*RecyclingCenterModel) Descriptor() ([]byte, []int) {
	return file_recycling_center_proto_rawDescGZIP(), []int{4}
}

func (x *RecyclingCenterModel) GetCenterName() string {
	if x != nil {
		return x.CenterName
	}
	return ""
}

func (x *RecyclingCenterModel) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RecyclingCenterModel) GetAcceptedMaterials() string {
	if x != nil {
		return x.AcceptedMaterials
	}
	return ""
}

var File_recycling_center_proto protoreflect.FileDescriptor

var file_recycling_center_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x36, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9b, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x2c, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x65, 0x64, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x22, 0xef,
	0x01, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x65, 0x64, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x6b, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x10, 0x52, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x22, 0x9b, 0x01,
	0x0a, 0x15, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a,
	0x11, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x65, 0x64, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x7e, 0x0a, 0x14, 0x52,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a,
	0x11, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x65, 0x64, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x32, 0xbc, 0x03, 0x0a, 0x16,
	0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12,
	0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a,
	0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00,
	0x12, 0x68, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x79,
	0x49, 0x64, 0x1a, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x22, 0x00, 0x12, 0x75, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x69, 0x6e, 0x67, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2e, 0x2e, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x2c, 0x2e, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recycling_center_proto_rawDescOnce sync.Once
	file_recycling_center_proto_rawDescData = file_recycling_center_proto_rawDesc
)

func file_recycling_center_proto_rawDescGZIP() []byte {
	file_recycling_center_proto_rawDescOnce.Do(func() {
		file_recycling_center_proto_rawDescData = protoimpl.X.CompressGZIP(file_recycling_center_proto_rawDescData)
	})
	return file_recycling_center_proto_rawDescData
}

var file_recycling_center_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_recycling_center_proto_goTypes = []any{
	(*RecyclingCenterCreate)(nil), // 0: enviromental_monitoring.RecyclingCenterCreate
	(*RecyclingCenter)(nil),       // 1: enviromental_monitoring.RecyclingCenter
	(*RecyclingCenterList)(nil),   // 2: enviromental_monitoring.RecyclingCenterList
	(*RecyclingCenterFilter)(nil), // 3: enviromental_monitoring.RecyclingCenterFilter
	(*RecyclingCenterModel)(nil),  // 4: enviromental_monitoring.RecyclingCenterModel
	(*ById)(nil),                  // 5: enviromental_monitoring.ById
	(*Void)(nil),                  // 6: enviromental_monitoring.Void
}
var file_recycling_center_proto_depIdxs = []int32{
	1, // 0: enviromental_monitoring.RecyclingCenterList.RecyclingCenters:type_name -> enviromental_monitoring.RecyclingCenter
	0, // 1: enviromental_monitoring.RecyclingCenterService.CreateRecyclingCenter:input_type -> enviromental_monitoring.RecyclingCenterCreate
	0, // 2: enviromental_monitoring.RecyclingCenterService.UpdateRecyclingCenter:input_type -> enviromental_monitoring.RecyclingCenterCreate
	5, // 3: enviromental_monitoring.RecyclingCenterService.DeleteRecyclingCenter:input_type -> enviromental_monitoring.ById
	3, // 4: enviromental_monitoring.RecyclingCenterService.GetRecyclingCenters:input_type -> enviromental_monitoring.RecyclingCenterFilter
	6, // 5: enviromental_monitoring.RecyclingCenterService.CreateRecyclingCenter:output_type -> enviromental_monitoring.Void
	6, // 6: enviromental_monitoring.RecyclingCenterService.UpdateRecyclingCenter:output_type -> enviromental_monitoring.Void
	6, // 7: enviromental_monitoring.RecyclingCenterService.DeleteRecyclingCenter:output_type -> enviromental_monitoring.Void
	2, // 8: enviromental_monitoring.RecyclingCenterService.GetRecyclingCenters:output_type -> enviromental_monitoring.RecyclingCenterList
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_recycling_center_proto_init() }
func file_recycling_center_proto_init() {
	if File_recycling_center_proto != nil {
		return
	}
	file_common6_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_recycling_center_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RecyclingCenterCreate); i {
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
		file_recycling_center_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RecyclingCenter); i {
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
		file_recycling_center_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RecyclingCenterList); i {
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
		file_recycling_center_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RecyclingCenterFilter); i {
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
		file_recycling_center_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RecyclingCenterModel); i {
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
			RawDescriptor: file_recycling_center_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recycling_center_proto_goTypes,
		DependencyIndexes: file_recycling_center_proto_depIdxs,
		MessageInfos:      file_recycling_center_proto_msgTypes,
	}.Build()
	File_recycling_center_proto = out.File
	file_recycling_center_proto_rawDesc = nil
	file_recycling_center_proto_goTypes = nil
	file_recycling_center_proto_depIdxs = nil
}