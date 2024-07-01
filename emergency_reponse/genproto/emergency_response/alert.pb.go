// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: alert.proto

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

type AlertsCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Message      string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	AffectedArea string `protobuf:"bytes,3,opt,name=affected_area,json=affectedArea,proto3" json:"affected_area,omitempty"`
	IssuedAt     string `protobuf:"bytes,4,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
}

func (x *AlertsCreateReq) Reset() {
	*x = AlertsCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertsCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertsCreateReq) ProtoMessage() {}

func (x *AlertsCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertsCreateReq.ProtoReflect.Descriptor instead.
func (*AlertsCreateReq) Descriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{0}
}

func (x *AlertsCreateReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AlertsCreateReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AlertsCreateReq) GetAffectedArea() string {
	if x != nil {
		return x.AffectedArea
	}
	return ""
}

func (x *AlertsCreateReq) GetIssuedAt() string {
	if x != nil {
		return x.IssuedAt
	}
	return ""
}

type AlertsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type         string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Message      string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	AffectedArea string `protobuf:"bytes,4,opt,name=affected_area,json=affectedArea,proto3" json:"affected_area,omitempty"`
	IssuedAt     string `protobuf:"bytes,5,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
}

func (x *AlertsRes) Reset() {
	*x = AlertsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertsRes) ProtoMessage() {}

func (x *AlertsRes) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertsRes.ProtoReflect.Descriptor instead.
func (*AlertsRes) Descriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{1}
}

func (x *AlertsRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AlertsRes) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AlertsRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AlertsRes) GetAffectedArea() string {
	if x != nil {
		return x.AffectedArea
	}
	return ""
}

func (x *AlertsRes) GetIssuedAt() string {
	if x != nil {
		return x.IssuedAt
	}
	return ""
}

type AlertsGetByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alert *AlertsRes `protobuf:"bytes,1,opt,name=alert,proto3" json:"alert,omitempty"`
}

func (x *AlertsGetByIdRes) Reset() {
	*x = AlertsGetByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertsGetByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertsGetByIdRes) ProtoMessage() {}

func (x *AlertsGetByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertsGetByIdRes.ProtoReflect.Descriptor instead.
func (*AlertsGetByIdRes) Descriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{2}
}

func (x *AlertsGetByIdRes) GetAlert() *AlertsRes {
	if x != nil {
		return x.Alert
	}
	return nil
}

type AlertsGetAllRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alerts []*AlertsRes `protobuf:"bytes,1,rep,name=alerts,proto3" json:"alerts,omitempty"`
}

func (x *AlertsGetAllRes) Reset() {
	*x = AlertsGetAllRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertsGetAllRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertsGetAllRes) ProtoMessage() {}

func (x *AlertsGetAllRes) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertsGetAllRes.ProtoReflect.Descriptor instead.
func (*AlertsGetAllRes) Descriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{3}
}

func (x *AlertsGetAllRes) GetAlerts() []*AlertsRes {
	if x != nil {
		return x.Alerts
	}
	return nil
}

type AlertsUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Alert *AlertsCreateReq `protobuf:"bytes,2,opt,name=alert,proto3" json:"alert,omitempty"`
}

func (x *AlertsUpdateReq) Reset() {
	*x = AlertsUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertsUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertsUpdateReq) ProtoMessage() {}

func (x *AlertsUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertsUpdateReq.ProtoReflect.Descriptor instead.
func (*AlertsUpdateReq) Descriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{4}
}

func (x *AlertsUpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AlertsUpdateReq) GetAlert() *AlertsCreateReq {
	if x != nil {
		return x.Alert
	}
	return nil
}

var File_alert_proto protoreflect.FileDescriptor

var file_alert_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x65,
	0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x81, 0x01, 0x0a, 0x0f, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x72,
	0x65, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x41, 0x72, 0x65, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x8b, 0x01, 0x0a, 0x09, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x72, 0x65, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x41, 0x72, 0x65, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x47, 0x0a, 0x10, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x52, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x22, 0x48, 0x0a, 0x0f, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x35, 0x0a,
	0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x52, 0x06, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x22, 0x5c, 0x0a, 0x0f, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e,
	0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x52, 0x05, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x32, 0x88, 0x03, 0x0a, 0x0c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e,
	0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18,
	0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x24, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1a, 0x2e, 0x65, 0x6d, 0x65,
	0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x23, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e,
	0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4e, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65,
	0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x65,
	0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3e, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65,
	0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x79, 0x49,
	0x64, 0x1a, 0x18, 0x2e, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x1d, 0x5a,
	0x1b, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6d, 0x65, 0x72, 0x67, 0x65,
	0x6e, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alert_proto_rawDescOnce sync.Once
	file_alert_proto_rawDescData = file_alert_proto_rawDesc
)

func file_alert_proto_rawDescGZIP() []byte {
	file_alert_proto_rawDescOnce.Do(func() {
		file_alert_proto_rawDescData = protoimpl.X.CompressGZIP(file_alert_proto_rawDescData)
	})
	return file_alert_proto_rawDescData
}

var file_alert_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_alert_proto_goTypes = []interface{}{
	(*AlertsCreateReq)(nil),  // 0: emergency_response.AlertsCreateReq
	(*AlertsRes)(nil),        // 1: emergency_response.AlertsRes
	(*AlertsGetByIdRes)(nil), // 2: emergency_response.AlertsGetByIdRes
	(*AlertsGetAllRes)(nil),  // 3: emergency_response.AlertsGetAllRes
	(*AlertsUpdateReq)(nil),  // 4: emergency_response.AlertsUpdateReq
	(*ById)(nil),             // 5: emergency_response.ById
	(*Filter)(nil),           // 6: emergency_response.Filter
	(*Void)(nil),             // 7: emergency_response.Void
}
var file_alert_proto_depIdxs = []int32{
	1, // 0: emergency_response.AlertsGetByIdRes.alert:type_name -> emergency_response.AlertsRes
	1, // 1: emergency_response.AlertsGetAllRes.alerts:type_name -> emergency_response.AlertsRes
	0, // 2: emergency_response.AlertsUpdateReq.alert:type_name -> emergency_response.AlertsCreateReq
	0, // 3: emergency_response.AlertService.Create:input_type -> emergency_response.AlertsCreateReq
	5, // 4: emergency_response.AlertService.GetById:input_type -> emergency_response.ById
	6, // 5: emergency_response.AlertService.GetAll:input_type -> emergency_response.Filter
	4, // 6: emergency_response.AlertService.Update:input_type -> emergency_response.AlertsUpdateReq
	5, // 7: emergency_response.AlertService.Delete:input_type -> emergency_response.ById
	1, // 8: emergency_response.AlertService.Create:output_type -> emergency_response.AlertsRes
	2, // 9: emergency_response.AlertService.GetById:output_type -> emergency_response.AlertsGetByIdRes
	3, // 10: emergency_response.AlertService.GetAll:output_type -> emergency_response.AlertsGetAllRes
	1, // 11: emergency_response.AlertService.Update:output_type -> emergency_response.AlertsRes
	7, // 12: emergency_response.AlertService.Delete:output_type -> emergency_response.Void
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_alert_proto_init() }
func file_alert_proto_init() {
	if File_alert_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_alert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertsCreateReq); i {
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
		file_alert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertsRes); i {
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
		file_alert_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertsGetByIdRes); i {
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
		file_alert_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertsGetAllRes); i {
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
		file_alert_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertsUpdateReq); i {
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
			RawDescriptor: file_alert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_alert_proto_goTypes,
		DependencyIndexes: file_alert_proto_depIdxs,
		MessageInfos:      file_alert_proto_msgTypes,
	}.Build()
	File_alert_proto = out.File
	file_alert_proto_rawDesc = nil
	file_alert_proto_goTypes = nil
	file_alert_proto_depIdxs = nil
}
