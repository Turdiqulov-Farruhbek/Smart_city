// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: maintance.proto

package transport

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

type MaintanceScheduleCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduleId string `protobuf:"bytes,1,opt,name=ScheduleId,proto3" json:"ScheduleId,omitempty"`
	RoadId     string `protobuf:"bytes,2,opt,name=RoadId,proto3" json:"RoadId,omitempty"`
	StartPoint string `protobuf:"bytes,3,opt,name=StartPoint,proto3" json:"StartPoint,omitempty"`
	EndPoint   string `protobuf:"bytes,4,opt,name=EndPoint,proto3" json:"EndPoint,omitempty"`
	StartDate  string `protobuf:"bytes,5,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate    string `protobuf:"bytes,6,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	Organizer  string `protobuf:"bytes,7,opt,name=Organizer,proto3" json:"Organizer,omitempty"`
	Status     string `protobuf:"bytes,8,opt,name=Status,proto3" json:"Status,omitempty"`
	TotalArea  string `protobuf:"bytes,9,opt,name=TotalArea,proto3" json:"TotalArea,omitempty"`
}

func (x *MaintanceScheduleCreate) Reset() {
	*x = MaintanceScheduleCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maintance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaintanceScheduleCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaintanceScheduleCreate) ProtoMessage() {}

func (x *MaintanceScheduleCreate) ProtoReflect() protoreflect.Message {
	mi := &file_maintance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaintanceScheduleCreate.ProtoReflect.Descriptor instead.
func (*MaintanceScheduleCreate) Descriptor() ([]byte, []int) {
	return file_maintance_proto_rawDescGZIP(), []int{0}
}

func (x *MaintanceScheduleCreate) GetScheduleId() string {
	if x != nil {
		return x.ScheduleId
	}
	return ""
}

func (x *MaintanceScheduleCreate) GetRoadId() string {
	if x != nil {
		return x.RoadId
	}
	return ""
}

func (x *MaintanceScheduleCreate) GetStartPoint() string {
	if x != nil {
		return x.StartPoint
	}
	return ""
}

func (x *MaintanceScheduleCreate) GetEndPoint() string {
	if x != nil {
		return x.EndPoint
	}
	return ""
}

func (x *MaintanceScheduleCreate) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *MaintanceScheduleCreate) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *MaintanceScheduleCreate) GetOrganizer() string {
	if x != nil {
		return x.Organizer
	}
	return ""
}

func (x *MaintanceScheduleCreate) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MaintanceScheduleCreate) GetTotalArea() string {
	if x != nil {
		return x.TotalArea
	}
	return ""
}

type MaintanceSchedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	RoadId     string `protobuf:"bytes,2,opt,name=RoadId,proto3" json:"RoadId,omitempty"`
	StartPoint string `protobuf:"bytes,3,opt,name=StartPoint,proto3" json:"StartPoint,omitempty"`
	EndPoint   string `protobuf:"bytes,4,opt,name=EndPoint,proto3" json:"EndPoint,omitempty"`
	StartTime  string `protobuf:"bytes,5,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime    string `protobuf:"bytes,6,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	Organizer  string `protobuf:"bytes,7,opt,name=Organizer,proto3" json:"Organizer,omitempty"`
	Status     string `protobuf:"bytes,8,opt,name=Status,proto3" json:"Status,omitempty"`
	TotalArea  int32  `protobuf:"varint,9,opt,name=TotalArea,proto3" json:"TotalArea,omitempty"`
	CreatedAt  string `protobuf:"bytes,10,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt  string `protobuf:"bytes,11,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt  string `protobuf:"bytes,12,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
}

func (x *MaintanceSchedule) Reset() {
	*x = MaintanceSchedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maintance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaintanceSchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaintanceSchedule) ProtoMessage() {}

func (x *MaintanceSchedule) ProtoReflect() protoreflect.Message {
	mi := &file_maintance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaintanceSchedule.ProtoReflect.Descriptor instead.
func (*MaintanceSchedule) Descriptor() ([]byte, []int) {
	return file_maintance_proto_rawDescGZIP(), []int{1}
}

func (x *MaintanceSchedule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MaintanceSchedule) GetRoadId() string {
	if x != nil {
		return x.RoadId
	}
	return ""
}

func (x *MaintanceSchedule) GetStartPoint() string {
	if x != nil {
		return x.StartPoint
	}
	return ""
}

func (x *MaintanceSchedule) GetEndPoint() string {
	if x != nil {
		return x.EndPoint
	}
	return ""
}

func (x *MaintanceSchedule) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *MaintanceSchedule) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *MaintanceSchedule) GetOrganizer() string {
	if x != nil {
		return x.Organizer
	}
	return ""
}

func (x *MaintanceSchedule) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MaintanceSchedule) GetTotalArea() int32 {
	if x != nil {
		return x.TotalArea
	}
	return 0
}

func (x *MaintanceSchedule) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *MaintanceSchedule) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *MaintanceSchedule) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type MaintanceScheduleFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDate string `protobuf:"bytes,1,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate   string `protobuf:"bytes,2,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	Organizer string `protobuf:"bytes,3,opt,name=Organizer,proto3" json:"Organizer,omitempty"`
	Status    string `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
	RoadId    string `protobuf:"bytes,5,opt,name=RoadId,proto3" json:"RoadId,omitempty"`
}

func (x *MaintanceScheduleFilter) Reset() {
	*x = MaintanceScheduleFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maintance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaintanceScheduleFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaintanceScheduleFilter) ProtoMessage() {}

func (x *MaintanceScheduleFilter) ProtoReflect() protoreflect.Message {
	mi := &file_maintance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaintanceScheduleFilter.ProtoReflect.Descriptor instead.
func (*MaintanceScheduleFilter) Descriptor() ([]byte, []int) {
	return file_maintance_proto_rawDescGZIP(), []int{2}
}

func (x *MaintanceScheduleFilter) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *MaintanceScheduleFilter) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *MaintanceScheduleFilter) GetOrganizer() string {
	if x != nil {
		return x.Organizer
	}
	return ""
}

func (x *MaintanceScheduleFilter) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MaintanceScheduleFilter) GetRoadId() string {
	if x != nil {
		return x.RoadId
	}
	return ""
}

type MaintanceScheduleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schedules []*MaintanceSchedule `protobuf:"bytes,1,rep,name=Schedules,proto3" json:"Schedules,omitempty"`
	Count     int32                `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *MaintanceScheduleList) Reset() {
	*x = MaintanceScheduleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maintance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaintanceScheduleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaintanceScheduleList) ProtoMessage() {}

func (x *MaintanceScheduleList) ProtoReflect() protoreflect.Message {
	mi := &file_maintance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaintanceScheduleList.ProtoReflect.Descriptor instead.
func (*MaintanceScheduleList) Descriptor() ([]byte, []int) {
	return file_maintance_proto_rawDescGZIP(), []int{3}
}

func (x *MaintanceScheduleList) GetSchedules() []*MaintanceSchedule {
	if x != nil {
		return x.Schedules
	}
	return nil
}

func (x *MaintanceScheduleList) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_maintance_proto protoreflect.FileDescriptor

var file_maintance_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x0c, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x02, 0x0a, 0x17, 0x4d,
	0x61, 0x69, 0x6e, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x61, 0x64, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x45, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x45, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x41, 0x72, 0x65, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x41, 0x72, 0x65, 0x61, 0x22, 0xdd, 0x02, 0x0a, 0x11, 0x4d, 0x61, 0x69, 0x6e, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x52, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x6f,
	0x61, 0x64, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x45, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x72, 0x65, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x72, 0x65, 0x61, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x9f, 0x01, 0x0a, 0x17, 0x4d, 0x61, 0x69, 0x6e, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x52, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x22, 0x69, 0x0a, 0x15, 0x4d, 0x61, 0x69, 0x6e,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x3a, 0x0a, 0x09, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x09, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x32, 0xf6, 0x02, 0x0a, 0x10, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x1c, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x1c, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x12, 0x0f, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x20, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_maintance_proto_rawDescOnce sync.Once
	file_maintance_proto_rawDescData = file_maintance_proto_rawDesc
)

func file_maintance_proto_rawDescGZIP() []byte {
	file_maintance_proto_rawDescOnce.Do(func() {
		file_maintance_proto_rawDescData = protoimpl.X.CompressGZIP(file_maintance_proto_rawDescData)
	})
	return file_maintance_proto_rawDescData
}

var file_maintance_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_maintance_proto_goTypes = []interface{}{
	(*MaintanceScheduleCreate)(nil), // 0: transport.MaintanceScheduleCreate
	(*MaintanceSchedule)(nil),       // 1: transport.MaintanceSchedule
	(*MaintanceScheduleFilter)(nil), // 2: transport.MaintanceScheduleFilter
	(*MaintanceScheduleList)(nil),   // 3: transport.MaintanceScheduleList
	(*ById)(nil),                    // 4: transport.ById
	(*Void)(nil),                    // 5: transport.Void
}
var file_maintance_proto_depIdxs = []int32{
	1, // 0: transport.MaintanceScheduleList.Schedules:type_name -> transport.MaintanceSchedule
	0, // 1: transport.MaintanceService.CreateMaintranceSchedule:input_type -> transport.MaintanceScheduleCreate
	0, // 2: transport.MaintanceService.UpdateMaintranceSchedule:input_type -> transport.MaintanceScheduleCreate
	4, // 3: transport.MaintanceService.DeleteMaintranceSchedule:input_type -> transport.ById
	2, // 4: transport.MaintanceService.GetAllMaintanceSchedules:input_type -> transport.MaintanceScheduleFilter
	1, // 5: transport.MaintanceService.CreateMaintranceSchedule:output_type -> transport.MaintanceSchedule
	1, // 6: transport.MaintanceService.UpdateMaintranceSchedule:output_type -> transport.MaintanceSchedule
	5, // 7: transport.MaintanceService.DeleteMaintranceSchedule:output_type -> transport.Void
	3, // 8: transport.MaintanceService.GetAllMaintanceSchedules:output_type -> transport.MaintanceScheduleList
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_maintance_proto_init() }
func file_maintance_proto_init() {
	if File_maintance_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_maintance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaintanceScheduleCreate); i {
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
		file_maintance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaintanceSchedule); i {
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
		file_maintance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaintanceScheduleFilter); i {
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
		file_maintance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaintanceScheduleList); i {
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
			RawDescriptor: file_maintance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_maintance_proto_goTypes,
		DependencyIndexes: file_maintance_proto_depIdxs,
		MessageInfos:      file_maintance_proto_msgTypes,
	}.Build()
	File_maintance_proto = out.File
	file_maintance_proto_rawDesc = nil
	file_maintance_proto_goTypes = nil
	file_maintance_proto_depIdxs = nil
}