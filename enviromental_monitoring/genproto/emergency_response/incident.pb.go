// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: incident.proto

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

type IncidentsCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Location    string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status      string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	ReportedAt  string `protobuf:"bytes,5,opt,name=reported_at,json=reportedAt,proto3" json:"reported_at,omitempty"`
}

func (x *IncidentsCreateReq) Reset() {
	*x = IncidentsCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_incident_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncidentsCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentsCreateReq) ProtoMessage() {}

func (x *IncidentsCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_incident_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentsCreateReq.ProtoReflect.Descriptor instead.
func (*IncidentsCreateReq) Descriptor() ([]byte, []int) {
	return file_incident_proto_rawDescGZIP(), []int{0}
}

func (x *IncidentsCreateReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *IncidentsCreateReq) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *IncidentsCreateReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *IncidentsCreateReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *IncidentsCreateReq) GetReportedAt() string {
	if x != nil {
		return x.ReportedAt
	}
	return ""
}

type IncidentsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type        string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Location    string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Status      string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	ReportedAt  string `protobuf:"bytes,6,opt,name=reported_at,json=reportedAt,proto3" json:"reported_at,omitempty"`
}

func (x *IncidentsRes) Reset() {
	*x = IncidentsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_incident_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncidentsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentsRes) ProtoMessage() {}

func (x *IncidentsRes) ProtoReflect() protoreflect.Message {
	mi := &file_incident_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentsRes.ProtoReflect.Descriptor instead.
func (*IncidentsRes) Descriptor() ([]byte, []int) {
	return file_incident_proto_rawDescGZIP(), []int{1}
}

func (x *IncidentsRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *IncidentsRes) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *IncidentsRes) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *IncidentsRes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *IncidentsRes) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *IncidentsRes) GetReportedAt() string {
	if x != nil {
		return x.ReportedAt
	}
	return ""
}

type IncidentsGetByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Incident *IncidentsRes `protobuf:"bytes,1,opt,name=incident,proto3" json:"incident,omitempty"`
}

func (x *IncidentsGetByIdRes) Reset() {
	*x = IncidentsGetByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_incident_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncidentsGetByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentsGetByIdRes) ProtoMessage() {}

func (x *IncidentsGetByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_incident_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentsGetByIdRes.ProtoReflect.Descriptor instead.
func (*IncidentsGetByIdRes) Descriptor() ([]byte, []int) {
	return file_incident_proto_rawDescGZIP(), []int{2}
}

func (x *IncidentsGetByIdRes) GetIncident() *IncidentsRes {
	if x != nil {
		return x.Incident
	}
	return nil
}

type IncidentsGetAllRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Incidents []*IncidentsRes `protobuf:"bytes,1,rep,name=incidents,proto3" json:"incidents,omitempty"`
}

func (x *IncidentsGetAllRes) Reset() {
	*x = IncidentsGetAllRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_incident_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncidentsGetAllRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentsGetAllRes) ProtoMessage() {}

func (x *IncidentsGetAllRes) ProtoReflect() protoreflect.Message {
	mi := &file_incident_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentsGetAllRes.ProtoReflect.Descriptor instead.
func (*IncidentsGetAllRes) Descriptor() ([]byte, []int) {
	return file_incident_proto_rawDescGZIP(), []int{3}
}

func (x *IncidentsGetAllRes) GetIncidents() []*IncidentsRes {
	if x != nil {
		return x.Incidents
	}
	return nil
}

type IncidentsUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Incident *IncidentsCreateReq `protobuf:"bytes,2,opt,name=incident,proto3" json:"incident,omitempty"`
}

func (x *IncidentsUpdateReq) Reset() {
	*x = IncidentsUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_incident_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncidentsUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentsUpdateReq) ProtoMessage() {}

func (x *IncidentsUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_incident_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentsUpdateReq.ProtoReflect.Descriptor instead.
func (*IncidentsUpdateReq) Descriptor() ([]byte, []int) {
	return file_incident_proto_rawDescGZIP(), []int{4}
}

func (x *IncidentsUpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *IncidentsUpdateReq) GetIncident() *IncidentsCreateReq {
	if x != nil {
		return x.Incident
	}
	return nil
}

var File_incident_proto protoreflect.FileDescriptor

var file_incident_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x12, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0xa9, 0x01, 0x0a, 0x0c, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x53, 0x0a, 0x13, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x08, 0x69, 0x6e, 0x63, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6d, 0x65, 0x72,
	0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49,
	0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x52, 0x08, 0x69, 0x6e, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x22, 0x54, 0x0a, 0x12, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x73, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x09, 0x69,
	0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x52, 0x09, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x68, 0x0a, 0x12, 0x49,
	0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x42, 0x0a, 0x08, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x52, 0x08, 0x69, 0x6e, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x32, 0xa6, 0x03, 0x0a, 0x18, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65,
	0x6e, 0x63, 0x79, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x54, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x65,
	0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x27, 0x2e,
	0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x1a, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x26,
	0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x26, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x65, 0x6d, 0x65,
	0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3e,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x79,
	0x49, 0x64, 0x1a, 0x18, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x1d,
	0x5a, 0x1b, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6d, 0x65, 0x72, 0x67,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_incident_proto_rawDescOnce sync.Once
	file_incident_proto_rawDescData = file_incident_proto_rawDesc
)

func file_incident_proto_rawDescGZIP() []byte {
	file_incident_proto_rawDescOnce.Do(func() {
		file_incident_proto_rawDescData = protoimpl.X.CompressGZIP(file_incident_proto_rawDescData)
	})
	return file_incident_proto_rawDescData
}

var file_incident_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_incident_proto_goTypes = []interface{}{
	(*IncidentsCreateReq)(nil),  // 0: emergency_response.IncidentsCreateReq
	(*IncidentsRes)(nil),        // 1: emergency_response.IncidentsRes
	(*IncidentsGetByIdRes)(nil), // 2: emergency_response.IncidentsGetByIdRes
	(*IncidentsGetAllRes)(nil),  // 3: emergency_response.IncidentsGetAllRes
	(*IncidentsUpdateReq)(nil),  // 4: emergency_response.IncidentsUpdateReq
	(*ById)(nil),                // 5: emergency_response.ById
	(*Filter)(nil),              // 6: emergency_response.Filter
	(*Void)(nil),                // 7: emergency_response.Void
}
var file_incident_proto_depIdxs = []int32{
	1, // 0: emergency_response.IncidentsGetByIdRes.incident:type_name -> emergency_response.IncidentsRes
	1, // 1: emergency_response.IncidentsGetAllRes.incidents:type_name -> emergency_response.IncidentsRes
	0, // 2: emergency_response.IncidentsUpdateReq.incident:type_name -> emergency_response.IncidentsCreateReq
	0, // 3: emergency_response.EmergencyIncidentService.Create:input_type -> emergency_response.IncidentsCreateReq
	5, // 4: emergency_response.EmergencyIncidentService.GetById:input_type -> emergency_response.ById
	6, // 5: emergency_response.EmergencyIncidentService.GetAll:input_type -> emergency_response.Filter
	4, // 6: emergency_response.EmergencyIncidentService.Update:input_type -> emergency_response.IncidentsUpdateReq
	5, // 7: emergency_response.EmergencyIncidentService.Delete:input_type -> emergency_response.ById
	1, // 8: emergency_response.EmergencyIncidentService.Create:output_type -> emergency_response.IncidentsRes
	2, // 9: emergency_response.EmergencyIncidentService.GetById:output_type -> emergency_response.IncidentsGetByIdRes
	3, // 10: emergency_response.EmergencyIncidentService.GetAll:output_type -> emergency_response.IncidentsGetAllRes
	1, // 11: emergency_response.EmergencyIncidentService.Update:output_type -> emergency_response.IncidentsRes
	7, // 12: emergency_response.EmergencyIncidentService.Delete:output_type -> emergency_response.Void
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_incident_proto_init() }
func file_incident_proto_init() {
	if File_incident_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_incident_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncidentsCreateReq); i {
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
		file_incident_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncidentsRes); i {
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
		file_incident_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncidentsGetByIdRes); i {
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
		file_incident_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncidentsGetAllRes); i {
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
		file_incident_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncidentsUpdateReq); i {
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
			RawDescriptor: file_incident_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_incident_proto_goTypes,
		DependencyIndexes: file_incident_proto_depIdxs,
		MessageInfos:      file_incident_proto_msgTypes,
	}.Build()
	File_incident_proto = out.File
	file_incident_proto_rawDesc = nil
	file_incident_proto_goTypes = nil
	file_incident_proto_depIdxs = nil
}
