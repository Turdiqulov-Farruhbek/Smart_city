// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: city_submodule/citizen/citizen.proto

package citizen

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

type Citizen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CitizenId   string `protobuf:"bytes,1,opt,name=CitizenId,proto3" json:"CitizenId,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	FirstName   string `protobuf:"bytes,3,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName    string `protobuf:"bytes,4,opt,name=LastName,proto3" json:"LastName,omitempty"`
	DateOfBirth string `protobuf:"bytes,5,opt,name=DateOfBirth,proto3" json:"DateOfBirth,omitempty"`
	Address     string `protobuf:"bytes,6,opt,name=Address,proto3" json:"Address,omitempty"`
	PhoneNumber string `protobuf:"bytes,7,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Email       string `protobuf:"bytes,8,opt,name=Email,proto3" json:"Email,omitempty"`
	CreateAt    string `protobuf:"bytes,9,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
	UpdateAt    string `protobuf:"bytes,10,opt,name=UpdateAt,proto3" json:"UpdateAt,omitempty"`
	DeleteAt    string `protobuf:"bytes,11,opt,name=DeleteAt,proto3" json:"DeleteAt,omitempty"`
}

func (x *Citizen) Reset() {
	*x = Citizen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_submodule_citizen_citizen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Citizen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Citizen) ProtoMessage() {}

func (x *Citizen) ProtoReflect() protoreflect.Message {
	mi := &file_city_submodule_citizen_citizen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Citizen.ProtoReflect.Descriptor instead.
func (*Citizen) Descriptor() ([]byte, []int) {
	return file_city_submodule_citizen_citizen_proto_rawDescGZIP(), []int{0}
}

func (x *Citizen) GetCitizenId() string {
	if x != nil {
		return x.CitizenId
	}
	return ""
}

func (x *Citizen) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Citizen) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Citizen) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Citizen) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *Citizen) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Citizen) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Citizen) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Citizen) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

func (x *Citizen) GetUpdateAt() string {
	if x != nil {
		return x.UpdateAt
	}
	return ""
}

func (x *Citizen) GetDeleteAt() string {
	if x != nil {
		return x.DeleteAt
	}
	return ""
}

type CitizenCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CitizenId   string `protobuf:"bytes,1,opt,name=CitizenId,proto3" json:"CitizenId,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	FirstName   string `protobuf:"bytes,3,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName    string `protobuf:"bytes,4,opt,name=LastName,proto3" json:"LastName,omitempty"`
	DateOfBirth string `protobuf:"bytes,5,opt,name=DateOfBirth,proto3" json:"DateOfBirth,omitempty"`
	Address     string `protobuf:"bytes,6,opt,name=Address,proto3" json:"Address,omitempty"`
	PhoneNumber string `protobuf:"bytes,7,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Email       string `protobuf:"bytes,8,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *CitizenCreate) Reset() {
	*x = CitizenCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_submodule_citizen_citizen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CitizenCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CitizenCreate) ProtoMessage() {}

func (x *CitizenCreate) ProtoReflect() protoreflect.Message {
	mi := &file_city_submodule_citizen_citizen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CitizenCreate.ProtoReflect.Descriptor instead.
func (*CitizenCreate) Descriptor() ([]byte, []int) {
	return file_city_submodule_citizen_citizen_proto_rawDescGZIP(), []int{1}
}

func (x *CitizenCreate) GetCitizenId() string {
	if x != nil {
		return x.CitizenId
	}
	return ""
}

func (x *CitizenCreate) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CitizenCreate) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CitizenCreate) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CitizenCreate) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *CitizenCreate) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CitizenCreate) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CitizenCreate) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CitizenModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	FirstName   string `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName    string `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	DateOfBirth string `protobuf:"bytes,4,opt,name=DateOfBirth,proto3" json:"DateOfBirth,omitempty"`
	Address     string `protobuf:"bytes,5,opt,name=Address,proto3" json:"Address,omitempty"`
	PhoneNumber string `protobuf:"bytes,6,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Email       string `protobuf:"bytes,7,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *CitizenModel) Reset() {
	*x = CitizenModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_submodule_citizen_citizen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CitizenModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CitizenModel) ProtoMessage() {}

func (x *CitizenModel) ProtoReflect() protoreflect.Message {
	mi := &file_city_submodule_citizen_citizen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CitizenModel.ProtoReflect.Descriptor instead.
func (*CitizenModel) Descriptor() ([]byte, []int) {
	return file_city_submodule_citizen_citizen_proto_rawDescGZIP(), []int{2}
}

func (x *CitizenModel) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CitizenModel) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CitizenModel) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CitizenModel) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *CitizenModel) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CitizenModel) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CitizenModel) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_city_submodule_citizen_citizen_proto protoreflect.FileDescriptor

var file_city_submodule_citizen_citizen_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2f, 0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x2f, 0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x1a,
	0x23, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f,
	0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x02, 0x0a, 0x07, 0x43, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72,
	0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x22, 0xf3, 0x01, 0x0a, 0x0d, 0x43, 0x69, 0x74,
	0x69, 0x7a, 0x65, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x69,
	0x74, 0x69, 0x7a, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43,
	0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x61,
	0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xd4,
	0x01, 0x0a, 0x0c, 0x43, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69,
	0x72, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0x83, 0x02, 0x0a, 0x0e, 0x43, 0x69, 0x74, 0x69, 0x7a, 0x65,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x43, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x12, 0x16, 0x2e, 0x63, 0x69,
	0x74, 0x69, 0x7a, 0x65, 0x6e, 0x2e, 0x43, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x1a, 0x10, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x2e, 0x43, 0x69,
	0x74, 0x69, 0x7a, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x69,
	0x74, 0x69, 0x7a, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0d, 0x2e, 0x63,
	0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x63, 0x69,
	0x74, 0x69, 0x7a, 0x65, 0x6e, 0x2e, 0x43, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x22, 0x00, 0x12,
	0x42, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x7a, 0x65,
	0x6e, 0x2e, 0x43, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a,
	0x10, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x2e, 0x43, 0x69, 0x74, 0x69, 0x7a, 0x65,
	0x6e, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x69, 0x74,
	0x69, 0x7a, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0d, 0x2e, 0x63, 0x69,
	0x74, 0x69, 0x7a, 0x65, 0x6e, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0d, 0x2e, 0x63, 0x69, 0x74,
	0x69, 0x7a, 0x65, 0x6e, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_submodule_citizen_citizen_proto_rawDescOnce sync.Once
	file_city_submodule_citizen_citizen_proto_rawDescData = file_city_submodule_citizen_citizen_proto_rawDesc
)

func file_city_submodule_citizen_citizen_proto_rawDescGZIP() []byte {
	file_city_submodule_citizen_citizen_proto_rawDescOnce.Do(func() {
		file_city_submodule_citizen_citizen_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_submodule_citizen_citizen_proto_rawDescData)
	})
	return file_city_submodule_citizen_citizen_proto_rawDescData
}

var file_city_submodule_citizen_citizen_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_city_submodule_citizen_citizen_proto_goTypes = []any{
	(*Citizen)(nil),       // 0: citizen.Citizen
	(*CitizenCreate)(nil), // 1: citizen.CitizenCreate
	(*CitizenModel)(nil),  // 2: citizen.CitizenModel
	(*ById)(nil),          // 3: citizen.ById
	(*Void)(nil),          // 4: citizen.Void
}
var file_city_submodule_citizen_citizen_proto_depIdxs = []int32{
	1, // 0: citizen.CitizenService.RegisterCitizen:input_type -> citizen.CitizenCreate
	3, // 1: citizen.CitizenService.GetCitizenProfile:input_type -> citizen.ById
	1, // 2: citizen.CitizenService.UpdateCitizenProfile:input_type -> citizen.CitizenCreate
	3, // 3: citizen.CitizenService.DeleteCitizenProfile:input_type -> citizen.ById
	0, // 4: citizen.CitizenService.RegisterCitizen:output_type -> citizen.Citizen
	0, // 5: citizen.CitizenService.GetCitizenProfile:output_type -> citizen.Citizen
	0, // 6: citizen.CitizenService.UpdateCitizenProfile:output_type -> citizen.Citizen
	4, // 7: citizen.CitizenService.DeleteCitizenProfile:output_type -> citizen.Void
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_city_submodule_citizen_citizen_proto_init() }
func file_city_submodule_citizen_citizen_proto_init() {
	if File_city_submodule_citizen_citizen_proto != nil {
		return
	}
	file_city_submodule_citizen_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_city_submodule_citizen_citizen_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Citizen); i {
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
		file_city_submodule_citizen_citizen_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CitizenCreate); i {
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
		file_city_submodule_citizen_citizen_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CitizenModel); i {
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
			RawDescriptor: file_city_submodule_citizen_citizen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_submodule_citizen_citizen_proto_goTypes,
		DependencyIndexes: file_city_submodule_citizen_citizen_proto_depIdxs,
		MessageInfos:      file_city_submodule_citizen_citizen_proto_msgTypes,
	}.Build()
	File_city_submodule_citizen_citizen_proto = out.File
	file_city_submodule_citizen_citizen_proto_rawDesc = nil
	file_city_submodule_citizen_citizen_proto_goTypes = nil
	file_city_submodule_citizen_citizen_proto_depIdxs = nil
}