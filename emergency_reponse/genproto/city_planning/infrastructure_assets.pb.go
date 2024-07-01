// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: infrastructure_assets.proto

package city_planning

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

type InfrastuctureAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetId          string `protobuf:"bytes,1,opt,name=AssetId,proto3" json:"AssetId,omitempty"`
	AssetType        string `protobuf:"bytes,2,opt,name=AssetType,proto3" json:"AssetType,omitempty"`
	Location         *Point `protobuf:"bytes,3,opt,name=Location,proto3" json:"Location,omitempty"`
	InstallationDate string `protobuf:"bytes,4,opt,name=InstallationDate,proto3" json:"InstallationDate,omitempty"`
	CurrentStatus    string `protobuf:"bytes,5,opt,name=CurrentStatus,proto3" json:"CurrentStatus,omitempty"`
	UpdatedAt        string `protobuf:"bytes,6,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt        string `protobuf:"bytes,7,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
}

func (x *InfrastuctureAsset) Reset() {
	*x = InfrastuctureAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_assets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfrastuctureAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfrastuctureAsset) ProtoMessage() {}

func (x *InfrastuctureAsset) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_assets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfrastuctureAsset.ProtoReflect.Descriptor instead.
func (*InfrastuctureAsset) Descriptor() ([]byte, []int) {
	return file_infrastructure_assets_proto_rawDescGZIP(), []int{0}
}

func (x *InfrastuctureAsset) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *InfrastuctureAsset) GetAssetType() string {
	if x != nil {
		return x.AssetType
	}
	return ""
}

func (x *InfrastuctureAsset) GetLocation() *Point {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *InfrastuctureAsset) GetInstallationDate() string {
	if x != nil {
		return x.InstallationDate
	}
	return ""
}

func (x *InfrastuctureAsset) GetCurrentStatus() string {
	if x != nil {
		return x.CurrentStatus
	}
	return ""
}

func (x *InfrastuctureAsset) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *InfrastuctureAsset) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type InfrastuctureAssetCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetId          string `protobuf:"bytes,1,opt,name=AssetId,proto3" json:"AssetId,omitempty"`
	AssetType        string `protobuf:"bytes,2,opt,name=AssetType,proto3" json:"AssetType,omitempty"`
	Location         *Point `protobuf:"bytes,3,opt,name=Location,proto3" json:"Location,omitempty"`
	InstallationDate string `protobuf:"bytes,4,opt,name=InstallationDate,proto3" json:"InstallationDate,omitempty"`
	CurrentStatus    string `protobuf:"bytes,5,opt,name=CurrentStatus,proto3" json:"CurrentStatus,omitempty"`
}

func (x *InfrastuctureAssetCreate) Reset() {
	*x = InfrastuctureAssetCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_assets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfrastuctureAssetCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfrastuctureAssetCreate) ProtoMessage() {}

func (x *InfrastuctureAssetCreate) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_assets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfrastuctureAssetCreate.ProtoReflect.Descriptor instead.
func (*InfrastuctureAssetCreate) Descriptor() ([]byte, []int) {
	return file_infrastructure_assets_proto_rawDescGZIP(), []int{1}
}

func (x *InfrastuctureAssetCreate) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *InfrastuctureAssetCreate) GetAssetType() string {
	if x != nil {
		return x.AssetType
	}
	return ""
}

func (x *InfrastuctureAssetCreate) GetLocation() *Point {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *InfrastuctureAssetCreate) GetInstallationDate() string {
	if x != nil {
		return x.InstallationDate
	}
	return ""
}

func (x *InfrastuctureAssetCreate) GetCurrentStatus() string {
	if x != nil {
		return x.CurrentStatus
	}
	return ""
}

type InfrastuctureAssetUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetType        string `protobuf:"bytes,1,opt,name=AssetType,proto3" json:"AssetType,omitempty"`
	Location         *Point `protobuf:"bytes,2,opt,name=Location,proto3" json:"Location,omitempty"`
	InstallationDate string `protobuf:"bytes,3,opt,name=InstallationDate,proto3" json:"InstallationDate,omitempty"`
	CurrentStatus    string `protobuf:"bytes,4,opt,name=CurrentStatus,proto3" json:"CurrentStatus,omitempty"`
}

func (x *InfrastuctureAssetUpdate) Reset() {
	*x = InfrastuctureAssetUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_assets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfrastuctureAssetUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfrastuctureAssetUpdate) ProtoMessage() {}

func (x *InfrastuctureAssetUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_assets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfrastuctureAssetUpdate.ProtoReflect.Descriptor instead.
func (*InfrastuctureAssetUpdate) Descriptor() ([]byte, []int) {
	return file_infrastructure_assets_proto_rawDescGZIP(), []int{2}
}

func (x *InfrastuctureAssetUpdate) GetAssetType() string {
	if x != nil {
		return x.AssetType
	}
	return ""
}

func (x *InfrastuctureAssetUpdate) GetLocation() *Point {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *InfrastuctureAssetUpdate) GetInstallationDate() string {
	if x != nil {
		return x.InstallationDate
	}
	return ""
}

func (x *InfrastuctureAssetUpdate) GetCurrentStatus() string {
	if x != nil {
		return x.CurrentStatus
	}
	return ""
}

type InfrastuctureAssetList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Assets []*InfrastuctureAsset `protobuf:"bytes,1,rep,name=Assets,proto3" json:"Assets,omitempty"`
}

func (x *InfrastuctureAssetList) Reset() {
	*x = InfrastuctureAssetList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_assets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfrastuctureAssetList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfrastuctureAssetList) ProtoMessage() {}

func (x *InfrastuctureAssetList) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_assets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfrastuctureAssetList.ProtoReflect.Descriptor instead.
func (*InfrastuctureAssetList) Descriptor() ([]byte, []int) {
	return file_infrastructure_assets_proto_rawDescGZIP(), []int{3}
}

func (x *InfrastuctureAssetList) GetAssets() []*InfrastuctureAsset {
	if x != nil {
		return x.Assets
	}
	return nil
}

type InfrastuctureAssetFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetId          string `protobuf:"bytes,1,opt,name=AssetId,proto3" json:"AssetId,omitempty"`
	AssetType        string `protobuf:"bytes,2,opt,name=AssetType,proto3" json:"AssetType,omitempty"`
	Location         *Point `protobuf:"bytes,3,opt,name=Location,proto3" json:"Location,omitempty"`
	InstallationDate string `protobuf:"bytes,4,opt,name=InstallationDate,proto3" json:"InstallationDate,omitempty"`
	CurrentStatus    string `protobuf:"bytes,5,opt,name=CurrentStatus,proto3" json:"CurrentStatus,omitempty"`
}

func (x *InfrastuctureAssetFilter) Reset() {
	*x = InfrastuctureAssetFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_assets_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfrastuctureAssetFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfrastuctureAssetFilter) ProtoMessage() {}

func (x *InfrastuctureAssetFilter) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_assets_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfrastuctureAssetFilter.ProtoReflect.Descriptor instead.
func (*InfrastuctureAssetFilter) Descriptor() ([]byte, []int) {
	return file_infrastructure_assets_proto_rawDescGZIP(), []int{4}
}

func (x *InfrastuctureAssetFilter) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *InfrastuctureAssetFilter) GetAssetType() string {
	if x != nil {
		return x.AssetType
	}
	return ""
}

func (x *InfrastuctureAssetFilter) GetLocation() *Point {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *InfrastuctureAssetFilter) GetInstallationDate() string {
	if x != nil {
		return x.InstallationDate
	}
	return ""
}

func (x *InfrastuctureAssetFilter) GetCurrentStatus() string {
	if x != nil {
		return x.CurrentStatus
	}
	return ""
}

var File_infrastructure_assets_proto protoreflect.FileDescriptor

var file_infrastructure_assets_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63,
	0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x0c, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x02, 0x0a, 0x12, 0x49,
	0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd6, 0x01, 0x0a, 0x18, 0x49, 0x6e,
	0x66, 0x72, 0x61, 0x73, 0x74, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30,
	0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2a, 0x0a, 0x10, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x18, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x75, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a,
	0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2a, 0x0a, 0x10, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x53, 0x0a, 0x16, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x66, 0x72,
	0x61, 0x73, 0x74, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x06,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x22, 0xd6, 0x01, 0x0a, 0x18, 0x49, 0x6e, 0x66, 0x72, 0x61,
	0x73, 0x74, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a,
	0x10, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32,
	0x8b, 0x03, 0x0a, 0x19, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x75, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a,
	0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x27, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x66, 0x72, 0x61,
	0x73, 0x74, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x1a, 0x13, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x19, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x27, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70,
	0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x1a, 0x13, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00,
	0x12, 0x6b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x27, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x66, 0x72,
	0x61, 0x73, 0x74, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x1a, 0x25, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x19, 0x5a,
	0x17, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x5f,
	0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infrastructure_assets_proto_rawDescOnce sync.Once
	file_infrastructure_assets_proto_rawDescData = file_infrastructure_assets_proto_rawDesc
)

func file_infrastructure_assets_proto_rawDescGZIP() []byte {
	file_infrastructure_assets_proto_rawDescOnce.Do(func() {
		file_infrastructure_assets_proto_rawDescData = protoimpl.X.CompressGZIP(file_infrastructure_assets_proto_rawDescData)
	})
	return file_infrastructure_assets_proto_rawDescData
}

var file_infrastructure_assets_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_infrastructure_assets_proto_goTypes = []interface{}{
	(*InfrastuctureAsset)(nil),       // 0: city_planning.InfrastuctureAsset
	(*InfrastuctureAssetCreate)(nil), // 1: city_planning.InfrastuctureAssetCreate
	(*InfrastuctureAssetUpdate)(nil), // 2: city_planning.InfrastuctureAssetUpdate
	(*InfrastuctureAssetList)(nil),   // 3: city_planning.InfrastuctureAssetList
	(*InfrastuctureAssetFilter)(nil), // 4: city_planning.InfrastuctureAssetFilter
	(*Point)(nil),                    // 5: city_planning.Point
	(*ById)(nil),                     // 6: city_planning.ById
	(*Void)(nil),                     // 7: city_planning.Void
}
var file_infrastructure_assets_proto_depIdxs = []int32{
	5, // 0: city_planning.InfrastuctureAsset.Location:type_name -> city_planning.Point
	5, // 1: city_planning.InfrastuctureAssetCreate.Location:type_name -> city_planning.Point
	5, // 2: city_planning.InfrastuctureAssetUpdate.Location:type_name -> city_planning.Point
	0, // 3: city_planning.InfrastuctureAssetList.Assets:type_name -> city_planning.InfrastuctureAsset
	5, // 4: city_planning.InfrastuctureAssetFilter.Location:type_name -> city_planning.Point
	1, // 5: city_planning.InfrastuctureAssetService.CreateInfrastructureAsset:input_type -> city_planning.InfrastuctureAssetCreate
	1, // 6: city_planning.InfrastuctureAssetService.UpdateInfrastructureAsset:input_type -> city_planning.InfrastuctureAssetCreate
	6, // 7: city_planning.InfrastuctureAssetService.DeleteInfrastructureAsset:input_type -> city_planning.ById
	4, // 8: city_planning.InfrastuctureAssetService.GetInfrastructureAssets:input_type -> city_planning.InfrastuctureAssetFilter
	7, // 9: city_planning.InfrastuctureAssetService.CreateInfrastructureAsset:output_type -> city_planning.Void
	7, // 10: city_planning.InfrastuctureAssetService.UpdateInfrastructureAsset:output_type -> city_planning.Void
	7, // 11: city_planning.InfrastuctureAssetService.DeleteInfrastructureAsset:output_type -> city_planning.Void
	3, // 12: city_planning.InfrastuctureAssetService.GetInfrastructureAssets:output_type -> city_planning.InfrastuctureAssetList
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_infrastructure_assets_proto_init() }
func file_infrastructure_assets_proto_init() {
	if File_infrastructure_assets_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_infrastructure_assets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfrastuctureAsset); i {
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
		file_infrastructure_assets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfrastuctureAssetCreate); i {
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
		file_infrastructure_assets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfrastuctureAssetUpdate); i {
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
		file_infrastructure_assets_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfrastuctureAssetList); i {
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
		file_infrastructure_assets_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfrastuctureAssetFilter); i {
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
			RawDescriptor: file_infrastructure_assets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infrastructure_assets_proto_goTypes,
		DependencyIndexes: file_infrastructure_assets_proto_depIdxs,
		MessageInfos:      file_infrastructure_assets_proto_msgTypes,
	}.Build()
	File_infrastructure_assets_proto = out.File
	file_infrastructure_assets_proto_rawDesc = nil
	file_infrastructure_assets_proto_goTypes = nil
	file_infrastructure_assets_proto_depIdxs = nil
}
