// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: traffic.proto

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

type TarfficCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConditionId  string `protobuf:"bytes,1,opt,name=ConditionId,proto3" json:"ConditionId,omitempty"`
	RoadId       string `protobuf:"bytes,2,opt,name=RoadId,proto3" json:"RoadId,omitempty"`
	TrafficLevel string `protobuf:"bytes,3,opt,name=TrafficLevel,proto3" json:"TrafficLevel,omitempty"`
	Description  string `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	CreatedAt    string `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt    string `protobuf:"bytes,6,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt    string `protobuf:"bytes,7,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
}

func (x *TarfficCondition) Reset() {
	*x = TarfficCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traffic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TarfficCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TarfficCondition) ProtoMessage() {}

func (x *TarfficCondition) ProtoReflect() protoreflect.Message {
	mi := &file_traffic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TarfficCondition.ProtoReflect.Descriptor instead.
func (*TarfficCondition) Descriptor() ([]byte, []int) {
	return file_traffic_proto_rawDescGZIP(), []int{0}
}

func (x *TarfficCondition) GetConditionId() string {
	if x != nil {
		return x.ConditionId
	}
	return ""
}

func (x *TarfficCondition) GetRoadId() string {
	if x != nil {
		return x.RoadId
	}
	return ""
}

func (x *TarfficCondition) GetTrafficLevel() string {
	if x != nil {
		return x.TrafficLevel
	}
	return ""
}

func (x *TarfficCondition) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TarfficCondition) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *TarfficCondition) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *TarfficCondition) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type TarfficConditionCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConditionId  string `protobuf:"bytes,1,opt,name=ConditionId,proto3" json:"ConditionId,omitempty"`
	RoadId       string `protobuf:"bytes,2,opt,name=RoadId,proto3" json:"RoadId,omitempty"`
	TrafficLevel string `protobuf:"bytes,3,opt,name=TrafficLevel,proto3" json:"TrafficLevel,omitempty"`
}

func (x *TarfficConditionCreate) Reset() {
	*x = TarfficConditionCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traffic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TarfficConditionCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TarfficConditionCreate) ProtoMessage() {}

func (x *TarfficConditionCreate) ProtoReflect() protoreflect.Message {
	mi := &file_traffic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TarfficConditionCreate.ProtoReflect.Descriptor instead.
func (*TarfficConditionCreate) Descriptor() ([]byte, []int) {
	return file_traffic_proto_rawDescGZIP(), []int{1}
}

func (x *TarfficConditionCreate) GetConditionId() string {
	if x != nil {
		return x.ConditionId
	}
	return ""
}

func (x *TarfficConditionCreate) GetRoadId() string {
	if x != nil {
		return x.RoadId
	}
	return ""
}

func (x *TarfficConditionCreate) GetTrafficLevel() string {
	if x != nil {
		return x.TrafficLevel
	}
	return ""
}

var File_traffic_proto protoreflect.FileDescriptor

var file_traffic_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x37, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a, 0x10, 0x54, 0x61,
	0x72, 0x66, 0x66, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x52, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x76, 0x0a, 0x16, 0x54, 0x61, 0x72, 0x66,
	0x66, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x32, 0x8e, 0x02, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x55, 0x70, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x54, 0x61, 0x72, 0x66, 0x66, 0x69, 0x63,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a,
	0x1b, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x54, 0x61, 0x72, 0x66,
	0x66, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x45,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x54, 0x61, 0x72, 0x66, 0x66, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x21, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x54, 0x61, 0x72, 0x66,
	0x66, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x1a, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x54,
	0x61, 0x72, 0x66, 0x66, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x00, 0x42, 0x15, 0x5a, 0x13, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_traffic_proto_rawDescOnce sync.Once
	file_traffic_proto_rawDescData = file_traffic_proto_rawDesc
)

func file_traffic_proto_rawDescGZIP() []byte {
	file_traffic_proto_rawDescOnce.Do(func() {
		file_traffic_proto_rawDescData = protoimpl.X.CompressGZIP(file_traffic_proto_rawDescData)
	})
	return file_traffic_proto_rawDescData
}

var file_traffic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_traffic_proto_goTypes = []any{
	(*TarfficCondition)(nil),       // 0: transport.TarfficCondition
	(*TarfficConditionCreate)(nil), // 1: transport.TarfficConditionCreate
	(*ById)(nil),                   // 2: transport.ById
}
var file_traffic_proto_depIdxs = []int32{
	1, // 0: transport.TrafficService.SetUpTrafficCondition:input_type -> transport.TarfficConditionCreate
	2, // 1: transport.TrafficService.GetTrafficCondition:input_type -> transport.ById
	1, // 2: transport.TrafficService.UpdateTrafficCondition:input_type -> transport.TarfficConditionCreate
	0, // 3: transport.TrafficService.SetUpTrafficCondition:output_type -> transport.TarfficCondition
	0, // 4: transport.TrafficService.GetTrafficCondition:output_type -> transport.TarfficCondition
	0, // 5: transport.TrafficService.UpdateTrafficCondition:output_type -> transport.TarfficCondition
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_traffic_proto_init() }
func file_traffic_proto_init() {
	if File_traffic_proto != nil {
		return
	}
	file_common7_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_traffic_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TarfficCondition); i {
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
		file_traffic_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TarfficConditionCreate); i {
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
			RawDescriptor: file_traffic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_traffic_proto_goTypes,
		DependencyIndexes: file_traffic_proto_depIdxs,
		MessageInfos:      file_traffic_proto_msgTypes,
	}.Build()
	File_traffic_proto = out.File
	file_traffic_proto_rawDesc = nil
	file_traffic_proto_goTypes = nil
	file_traffic_proto_depIdxs = nil
}
