// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: city_zone.proto

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

type CityZone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneId       string   `protobuf:"bytes,1,opt,name=ZoneId,proto3" json:"ZoneId,omitempty"`
	ZoneName     string   `protobuf:"bytes,2,opt,name=ZoneName,proto3" json:"ZoneName,omitempty"`
	Area         *Polygon `protobuf:"bytes,3,opt,name=Area,proto3" json:"Area,omitempty"`
	CurrentUsage string   `protobuf:"bytes,4,opt,name=CurrentUsage,proto3" json:"CurrentUsage,omitempty"`
	CreatedAt    string   `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt    string   `protobuf:"bytes,6,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt    string   `protobuf:"bytes,7,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
}

func (x *CityZone) Reset() {
	*x = CityZone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_zone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityZone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityZone) ProtoMessage() {}

func (x *CityZone) ProtoReflect() protoreflect.Message {
	mi := &file_city_zone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityZone.ProtoReflect.Descriptor instead.
func (*CityZone) Descriptor() ([]byte, []int) {
	return file_city_zone_proto_rawDescGZIP(), []int{0}
}

func (x *CityZone) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *CityZone) GetZoneName() string {
	if x != nil {
		return x.ZoneName
	}
	return ""
}

func (x *CityZone) GetArea() *Polygon {
	if x != nil {
		return x.Area
	}
	return nil
}

func (x *CityZone) GetCurrentUsage() string {
	if x != nil {
		return x.CurrentUsage
	}
	return ""
}

func (x *CityZone) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CityZone) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *CityZone) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type CityZoneCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneId       string   `protobuf:"bytes,1,opt,name=ZoneId,proto3" json:"ZoneId,omitempty"`
	ZoneName     string   `protobuf:"bytes,2,opt,name=ZoneName,proto3" json:"ZoneName,omitempty"`
	Area         *Polygon `protobuf:"bytes,3,opt,name=Area,proto3" json:"Area,omitempty"`
	CurrentUsage string   `protobuf:"bytes,4,opt,name=CurrentUsage,proto3" json:"CurrentUsage,omitempty"`
}

func (x *CityZoneCreate) Reset() {
	*x = CityZoneCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_zone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityZoneCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityZoneCreate) ProtoMessage() {}

func (x *CityZoneCreate) ProtoReflect() protoreflect.Message {
	mi := &file_city_zone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityZoneCreate.ProtoReflect.Descriptor instead.
func (*CityZoneCreate) Descriptor() ([]byte, []int) {
	return file_city_zone_proto_rawDescGZIP(), []int{1}
}

func (x *CityZoneCreate) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *CityZoneCreate) GetZoneName() string {
	if x != nil {
		return x.ZoneName
	}
	return ""
}

func (x *CityZoneCreate) GetArea() *Polygon {
	if x != nil {
		return x.Area
	}
	return nil
}

func (x *CityZoneCreate) GetCurrentUsage() string {
	if x != nil {
		return x.CurrentUsage
	}
	return ""
}

type CityZoneUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneName     string   `protobuf:"bytes,1,opt,name=ZoneName,proto3" json:"ZoneName,omitempty"`
	Area         *Polygon `protobuf:"bytes,2,opt,name=Area,proto3" json:"Area,omitempty"`
	CurrentUsage string   `protobuf:"bytes,3,opt,name=CurrentUsage,proto3" json:"CurrentUsage,omitempty"`
}

func (x *CityZoneUpdate) Reset() {
	*x = CityZoneUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_zone_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityZoneUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityZoneUpdate) ProtoMessage() {}

func (x *CityZoneUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_city_zone_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityZoneUpdate.ProtoReflect.Descriptor instead.
func (*CityZoneUpdate) Descriptor() ([]byte, []int) {
	return file_city_zone_proto_rawDescGZIP(), []int{2}
}

func (x *CityZoneUpdate) GetZoneName() string {
	if x != nil {
		return x.ZoneName
	}
	return ""
}

func (x *CityZoneUpdate) GetArea() *Polygon {
	if x != nil {
		return x.Area
	}
	return nil
}

func (x *CityZoneUpdate) GetCurrentUsage() string {
	if x != nil {
		return x.CurrentUsage
	}
	return ""
}

type CityZoneList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityZones []*CityZone `protobuf:"bytes,1,rep,name=CityZones,proto3" json:"CityZones,omitempty"`
}

func (x *CityZoneList) Reset() {
	*x = CityZoneList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_zone_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityZoneList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityZoneList) ProtoMessage() {}

func (x *CityZoneList) ProtoReflect() protoreflect.Message {
	mi := &file_city_zone_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityZoneList.ProtoReflect.Descriptor instead.
func (*CityZoneList) Descriptor() ([]byte, []int) {
	return file_city_zone_proto_rawDescGZIP(), []int{3}
}

func (x *CityZoneList) GetCityZones() []*CityZone {
	if x != nil {
		return x.CityZones
	}
	return nil
}

type CityZoneFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneId       string   `protobuf:"bytes,1,opt,name=ZoneId,proto3" json:"ZoneId,omitempty"`
	ZoneName     string   `protobuf:"bytes,2,opt,name=ZoneName,proto3" json:"ZoneName,omitempty"`
	Area         *Polygon `protobuf:"bytes,3,opt,name=Area,proto3" json:"Area,omitempty"`
	CurrentUsage string   `protobuf:"bytes,4,opt,name=CurrentUsage,proto3" json:"CurrentUsage,omitempty"`
}

func (x *CityZoneFilter) Reset() {
	*x = CityZoneFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_zone_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityZoneFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityZoneFilter) ProtoMessage() {}

func (x *CityZoneFilter) ProtoReflect() protoreflect.Message {
	mi := &file_city_zone_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityZoneFilter.ProtoReflect.Descriptor instead.
func (*CityZoneFilter) Descriptor() ([]byte, []int) {
	return file_city_zone_proto_rawDescGZIP(), []int{4}
}

func (x *CityZoneFilter) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *CityZoneFilter) GetZoneName() string {
	if x != nil {
		return x.ZoneName
	}
	return ""
}

func (x *CityZoneFilter) GetArea() *Polygon {
	if x != nil {
		return x.Area
	}
	return nil
}

func (x *CityZoneFilter) GetCurrentUsage() string {
	if x != nil {
		return x.CurrentUsage
	}
	return ""
}

var File_city_zone_proto protoreflect.FileDescriptor

var file_city_zone_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8,
	0x01, 0x0a, 0x08, 0x43, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x5a,
	0x6f, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x5a, 0x6f, 0x6e,
	0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x5a, 0x6f, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x5a, 0x6f, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2a, 0x0a, 0x04, 0x41, 0x72, 0x65, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6f,
	0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52, 0x04, 0x41, 0x72, 0x65, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x94, 0x01, 0x0a, 0x0e, 0x43, 0x69,
	0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x5a, 0x6f,
	0x6e, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x5a, 0x6f, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x5a, 0x6f, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2a, 0x0a, 0x04, 0x41, 0x72, 0x65, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x50,
	0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52, 0x04, 0x41, 0x72, 0x65, 0x61, 0x12, 0x22, 0x0a, 0x0c,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x7c, 0x0a, 0x0e, 0x43, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x5a, 0x6f, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x5a, 0x6f, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a,
	0x0a, 0x04, 0x41, 0x72, 0x65, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63,
	0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x6c,
	0x79, 0x67, 0x6f, 0x6e, 0x52, 0x04, 0x41, 0x72, 0x65, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x22, 0x45,
	0x0a, 0x0c, 0x43, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35,
	0x0a, 0x09, 0x43, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x09, 0x43, 0x69, 0x74, 0x79,
	0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x0e, 0x43, 0x69, 0x74, 0x79, 0x5a, 0x6f,
	0x6e, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x5a, 0x6f, 0x6e, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x5a, 0x6f, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x5a, 0x6f, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x04,
	0x41, 0x72, 0x65, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x6c, 0x79, 0x67,
	0x6f, 0x6e, 0x52, 0x04, 0x41, 0x72, 0x65, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa4, 0x02, 0x0a,
	0x0f, 0x43, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x44, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x69, 0x74, 0x79, 0x5a, 0x6f,
	0x6e, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x1a, 0x13, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x44, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f,
	0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x13, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70,
	0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x13,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x42,
	0x79, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x49, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43,
	0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70,
	0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1b, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c,
	0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x19, 0x5a, 0x17, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_zone_proto_rawDescOnce sync.Once
	file_city_zone_proto_rawDescData = file_city_zone_proto_rawDesc
)

func file_city_zone_proto_rawDescGZIP() []byte {
	file_city_zone_proto_rawDescOnce.Do(func() {
		file_city_zone_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_zone_proto_rawDescData)
	})
	return file_city_zone_proto_rawDescData
}

var file_city_zone_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_city_zone_proto_goTypes = []interface{}{
	(*CityZone)(nil),       // 0: city_planning.CityZone
	(*CityZoneCreate)(nil), // 1: city_planning.CityZoneCreate
	(*CityZoneUpdate)(nil), // 2: city_planning.CityZoneUpdate
	(*CityZoneList)(nil),   // 3: city_planning.CityZoneList
	(*CityZoneFilter)(nil), // 4: city_planning.CityZoneFilter
	(*Polygon)(nil),        // 5: city_planning.Polygon
	(*ById)(nil),           // 6: city_planning.ById
	(*Void)(nil),           // 7: city_planning.Void
}
var file_city_zone_proto_depIdxs = []int32{
	5, // 0: city_planning.CityZone.Area:type_name -> city_planning.Polygon
	5, // 1: city_planning.CityZoneCreate.Area:type_name -> city_planning.Polygon
	5, // 2: city_planning.CityZoneUpdate.Area:type_name -> city_planning.Polygon
	0, // 3: city_planning.CityZoneList.CityZones:type_name -> city_planning.CityZone
	5, // 4: city_planning.CityZoneFilter.Area:type_name -> city_planning.Polygon
	1, // 5: city_planning.CityZoneService.CreateCityZone:input_type -> city_planning.CityZoneCreate
	1, // 6: city_planning.CityZoneService.UpdateCityZone:input_type -> city_planning.CityZoneCreate
	6, // 7: city_planning.CityZoneService.DeleteCityZone:input_type -> city_planning.ById
	4, // 8: city_planning.CityZoneService.GetCityZone:input_type -> city_planning.CityZoneFilter
	7, // 9: city_planning.CityZoneService.CreateCityZone:output_type -> city_planning.Void
	7, // 10: city_planning.CityZoneService.UpdateCityZone:output_type -> city_planning.Void
	7, // 11: city_planning.CityZoneService.DeleteCityZone:output_type -> city_planning.Void
	3, // 12: city_planning.CityZoneService.GetCityZone:output_type -> city_planning.CityZoneList
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_city_zone_proto_init() }
func file_city_zone_proto_init() {
	if File_city_zone_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_city_zone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityZone); i {
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
		file_city_zone_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityZoneCreate); i {
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
		file_city_zone_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityZoneUpdate); i {
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
		file_city_zone_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityZoneList); i {
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
		file_city_zone_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityZoneFilter); i {
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
			RawDescriptor: file_city_zone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_zone_proto_goTypes,
		DependencyIndexes: file_city_zone_proto_depIdxs,
		MessageInfos:      file_city_zone_proto_msgTypes,
	}.Build()
	File_city_zone_proto = out.File
	file_city_zone_proto_rawDesc = nil
	file_city_zone_proto_goTypes = nil
	file_city_zone_proto_depIdxs = nil
}
