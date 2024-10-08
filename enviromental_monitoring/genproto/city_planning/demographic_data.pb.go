// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: demographic_data.proto

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

type DemographicData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataId       string  `protobuf:"bytes,1,opt,name=DataId,proto3" json:"DataId,omitempty"`
	ZoneId       string  `protobuf:"bytes,2,opt,name=ZoneId,proto3" json:"ZoneId,omitempty"`
	Population   int32   `protobuf:"varint,3,opt,name=Population,proto3" json:"Population,omitempty"`
	MedianAge    int32   `protobuf:"varint,4,opt,name=MedianAge,proto3" json:"MedianAge,omitempty"`
	MedianIncome float32 `protobuf:"fixed32,5,opt,name=MedianIncome,proto3" json:"MedianIncome,omitempty"`
	DataYear     string  `protobuf:"bytes,6,opt,name=DataYear,proto3" json:"DataYear,omitempty"`
	UpdatedAt    string  `protobuf:"bytes,7,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt    string  `protobuf:"bytes,8,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
}

func (x *DemographicData) Reset() {
	*x = DemographicData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demographic_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemographicData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemographicData) ProtoMessage() {}

func (x *DemographicData) ProtoReflect() protoreflect.Message {
	mi := &file_demographic_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemographicData.ProtoReflect.Descriptor instead.
func (*DemographicData) Descriptor() ([]byte, []int) {
	return file_demographic_data_proto_rawDescGZIP(), []int{0}
}

func (x *DemographicData) GetDataId() string {
	if x != nil {
		return x.DataId
	}
	return ""
}

func (x *DemographicData) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *DemographicData) GetPopulation() int32 {
	if x != nil {
		return x.Population
	}
	return 0
}

func (x *DemographicData) GetMedianAge() int32 {
	if x != nil {
		return x.MedianAge
	}
	return 0
}

func (x *DemographicData) GetMedianIncome() float32 {
	if x != nil {
		return x.MedianIncome
	}
	return 0
}

func (x *DemographicData) GetDataYear() string {
	if x != nil {
		return x.DataYear
	}
	return ""
}

func (x *DemographicData) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *DemographicData) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type DemographicDataCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataId       string  `protobuf:"bytes,1,opt,name=DataId,proto3" json:"DataId,omitempty"`
	ZoneId       string  `protobuf:"bytes,2,opt,name=ZoneId,proto3" json:"ZoneId,omitempty"`
	Population   int32   `protobuf:"varint,3,opt,name=Population,proto3" json:"Population,omitempty"`
	MedianAge    int32   `protobuf:"varint,4,opt,name=MedianAge,proto3" json:"MedianAge,omitempty"`
	MedianIncome float32 `protobuf:"fixed32,5,opt,name=MedianIncome,proto3" json:"MedianIncome,omitempty"`
	DataYear     string  `protobuf:"bytes,6,opt,name=DataYear,proto3" json:"DataYear,omitempty"`
}

func (x *DemographicDataCreate) Reset() {
	*x = DemographicDataCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demographic_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemographicDataCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemographicDataCreate) ProtoMessage() {}

func (x *DemographicDataCreate) ProtoReflect() protoreflect.Message {
	mi := &file_demographic_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemographicDataCreate.ProtoReflect.Descriptor instead.
func (*DemographicDataCreate) Descriptor() ([]byte, []int) {
	return file_demographic_data_proto_rawDescGZIP(), []int{1}
}

func (x *DemographicDataCreate) GetDataId() string {
	if x != nil {
		return x.DataId
	}
	return ""
}

func (x *DemographicDataCreate) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *DemographicDataCreate) GetPopulation() int32 {
	if x != nil {
		return x.Population
	}
	return 0
}

func (x *DemographicDataCreate) GetMedianAge() int32 {
	if x != nil {
		return x.MedianAge
	}
	return 0
}

func (x *DemographicDataCreate) GetMedianIncome() float32 {
	if x != nil {
		return x.MedianIncome
	}
	return 0
}

func (x *DemographicDataCreate) GetDataYear() string {
	if x != nil {
		return x.DataYear
	}
	return ""
}

type DemographicDataUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneId       string  `protobuf:"bytes,1,opt,name=ZoneId,proto3" json:"ZoneId,omitempty"`
	Population   int32   `protobuf:"varint,2,opt,name=Population,proto3" json:"Population,omitempty"`
	MedianAge    int32   `protobuf:"varint,3,opt,name=MedianAge,proto3" json:"MedianAge,omitempty"`
	MedianIncome float32 `protobuf:"fixed32,4,opt,name=MedianIncome,proto3" json:"MedianIncome,omitempty"`
	DataYear     string  `protobuf:"bytes,5,opt,name=DataYear,proto3" json:"DataYear,omitempty"`
}

func (x *DemographicDataUpdate) Reset() {
	*x = DemographicDataUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demographic_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemographicDataUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemographicDataUpdate) ProtoMessage() {}

func (x *DemographicDataUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_demographic_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemographicDataUpdate.ProtoReflect.Descriptor instead.
func (*DemographicDataUpdate) Descriptor() ([]byte, []int) {
	return file_demographic_data_proto_rawDescGZIP(), []int{2}
}

func (x *DemographicDataUpdate) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *DemographicDataUpdate) GetPopulation() int32 {
	if x != nil {
		return x.Population
	}
	return 0
}

func (x *DemographicDataUpdate) GetMedianAge() int32 {
	if x != nil {
		return x.MedianAge
	}
	return 0
}

func (x *DemographicDataUpdate) GetMedianIncome() float32 {
	if x != nil {
		return x.MedianIncome
	}
	return 0
}

func (x *DemographicDataUpdate) GetDataYear() string {
	if x != nil {
		return x.DataYear
	}
	return ""
}

type DemographicDataList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DemographicDatas []*DemographicData `protobuf:"bytes,1,rep,name=DemographicDatas,proto3" json:"DemographicDatas,omitempty"`
}

func (x *DemographicDataList) Reset() {
	*x = DemographicDataList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demographic_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemographicDataList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemographicDataList) ProtoMessage() {}

func (x *DemographicDataList) ProtoReflect() protoreflect.Message {
	mi := &file_demographic_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemographicDataList.ProtoReflect.Descriptor instead.
func (*DemographicDataList) Descriptor() ([]byte, []int) {
	return file_demographic_data_proto_rawDescGZIP(), []int{3}
}

func (x *DemographicDataList) GetDemographicDatas() []*DemographicData {
	if x != nil {
		return x.DemographicDatas
	}
	return nil
}

type DemographicDataFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataId       string  `protobuf:"bytes,1,opt,name=DataId,proto3" json:"DataId,omitempty"`
	ZoneId       string  `protobuf:"bytes,2,opt,name=ZoneId,proto3" json:"ZoneId,omitempty"`
	Population   int32   `protobuf:"varint,3,opt,name=Population,proto3" json:"Population,omitempty"`
	MedianAge    int32   `protobuf:"varint,4,opt,name=MedianAge,proto3" json:"MedianAge,omitempty"`
	MedianIncome float32 `protobuf:"fixed32,5,opt,name=MedianIncome,proto3" json:"MedianIncome,omitempty"`
	DataYear     string  `protobuf:"bytes,6,opt,name=DataYear,proto3" json:"DataYear,omitempty"`
}

func (x *DemographicDataFilter) Reset() {
	*x = DemographicDataFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demographic_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemographicDataFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemographicDataFilter) ProtoMessage() {}

func (x *DemographicDataFilter) ProtoReflect() protoreflect.Message {
	mi := &file_demographic_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemographicDataFilter.ProtoReflect.Descriptor instead.
func (*DemographicDataFilter) Descriptor() ([]byte, []int) {
	return file_demographic_data_proto_rawDescGZIP(), []int{4}
}

func (x *DemographicDataFilter) GetDataId() string {
	if x != nil {
		return x.DataId
	}
	return ""
}

func (x *DemographicDataFilter) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *DemographicDataFilter) GetPopulation() int32 {
	if x != nil {
		return x.Population
	}
	return 0
}

func (x *DemographicDataFilter) GetMedianAge() int32 {
	if x != nil {
		return x.MedianAge
	}
	return 0
}

func (x *DemographicDataFilter) GetMedianIncome() float32 {
	if x != nil {
		return x.MedianIncome
	}
	return 0
}

func (x *DemographicDataFilter) GetDataYear() string {
	if x != nil {
		return x.DataYear
	}
	return ""
}

var File_demographic_data_proto protoreflect.FileDescriptor

var file_demographic_data_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70,
	0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x61, 0x74,
	0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x61, 0x74, 0x61, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6f, 0x70,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x50,
	0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x6e, 0x41, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x6e, 0x41, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x6e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x6e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44,
	0x61, 0x74, 0x61, 0x59, 0x65, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44,
	0x61, 0x74, 0x61, 0x59, 0x65, 0x61, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xc5, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x44, 0x61, 0x74, 0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44,
	0x61, 0x74, 0x61, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x41, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x41, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x6e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x59, 0x65, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x59, 0x65, 0x61, 0x72, 0x22, 0xad, 0x01, 0x0a, 0x15,
	0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x41, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x41, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x6e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x59, 0x65, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x59, 0x65, 0x61, 0x72, 0x22, 0x61, 0x0a, 0x13, 0x44,
	0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x4a, 0x0a, 0x10, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69,
	0x63, 0x44, 0x61, 0x74, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63,
	0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6d,
	0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x52, 0x10, 0x44, 0x65,
	0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x73, 0x22, 0xc5,
	0x01, 0x0a, 0x15, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x44, 0x61,
	0x74, 0x61, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x61, 0x74, 0x61,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x61, 0x74, 0x61, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6f, 0x70, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x50, 0x6f,
	0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x6e, 0x41, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x6e, 0x41, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e,
	0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x6e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x61,
	0x74, 0x61, 0x59, 0x65, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x61,
	0x74, 0x61, 0x59, 0x65, 0x61, 0x72, 0x32, 0xec, 0x02, 0x0a, 0x16, 0x44, 0x65, 0x6d, 0x6f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x54, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x12, 0x24, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x1a, 0x13, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x24, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x13, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c,
	0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x12, 0x13, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c,
	0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64,
	0x22, 0x00, 0x12, 0x61, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x73, 0x12, 0x24, 0x2e, 0x63, 0x69, 0x74, 0x79,
	0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a,
	0x22, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x44, 0x65, 0x6d, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_demographic_data_proto_rawDescOnce sync.Once
	file_demographic_data_proto_rawDescData = file_demographic_data_proto_rawDesc
)

func file_demographic_data_proto_rawDescGZIP() []byte {
	file_demographic_data_proto_rawDescOnce.Do(func() {
		file_demographic_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_demographic_data_proto_rawDescData)
	})
	return file_demographic_data_proto_rawDescData
}

var file_demographic_data_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_demographic_data_proto_goTypes = []interface{}{
	(*DemographicData)(nil),       // 0: city_planning.DemographicData
	(*DemographicDataCreate)(nil), // 1: city_planning.DemographicDataCreate
	(*DemographicDataUpdate)(nil), // 2: city_planning.DemographicDataUpdate
	(*DemographicDataList)(nil),   // 3: city_planning.DemographicDataList
	(*DemographicDataFilter)(nil), // 4: city_planning.DemographicDataFilter
	(*ById)(nil),                  // 5: city_planning.ById
	(*Void)(nil),                  // 6: city_planning.Void
}
var file_demographic_data_proto_depIdxs = []int32{
	0, // 0: city_planning.DemographicDataList.DemographicDatas:type_name -> city_planning.DemographicData
	1, // 1: city_planning.DemographicDataService.CreateDemographicData:input_type -> city_planning.DemographicDataCreate
	1, // 2: city_planning.DemographicDataService.UpdateDemographicData:input_type -> city_planning.DemographicDataCreate
	5, // 3: city_planning.DemographicDataService.DeleteDemographicData:input_type -> city_planning.ById
	4, // 4: city_planning.DemographicDataService.GetDemographicDatas:input_type -> city_planning.DemographicDataFilter
	6, // 5: city_planning.DemographicDataService.CreateDemographicData:output_type -> city_planning.Void
	6, // 6: city_planning.DemographicDataService.UpdateDemographicData:output_type -> city_planning.Void
	6, // 7: city_planning.DemographicDataService.DeleteDemographicData:output_type -> city_planning.Void
	3, // 8: city_planning.DemographicDataService.GetDemographicDatas:output_type -> city_planning.DemographicDataList
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_demographic_data_proto_init() }
func file_demographic_data_proto_init() {
	if File_demographic_data_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_demographic_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemographicData); i {
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
		file_demographic_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemographicDataCreate); i {
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
		file_demographic_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemographicDataUpdate); i {
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
		file_demographic_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemographicDataList); i {
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
		file_demographic_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemographicDataFilter); i {
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
			RawDescriptor: file_demographic_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demographic_data_proto_goTypes,
		DependencyIndexes: file_demographic_data_proto_depIdxs,
		MessageInfos:      file_demographic_data_proto_msgTypes,
	}.Build()
	File_demographic_data_proto = out.File
	file_demographic_data_proto_rawDesc = nil
	file_demographic_data_proto_goTypes = nil
	file_demographic_data_proto_depIdxs = nil
}
