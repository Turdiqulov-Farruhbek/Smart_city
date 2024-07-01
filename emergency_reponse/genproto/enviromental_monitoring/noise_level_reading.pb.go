// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: noise_level_reading.proto

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

type NoiseLevelReading struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReadingId    string  `protobuf:"bytes,1,opt,name=ReadingId,proto3" json:"ReadingId,omitempty"`
	ZoneId       string  `protobuf:"bytes,2,opt,name=ZoneId,proto3" json:"ZoneId,omitempty"`
	DecibelLevel float32 `protobuf:"fixed32,3,opt,name=DecibelLevel,proto3" json:"DecibelLevel,omitempty"`
	Time         string  `protobuf:"bytes,4,opt,name=Time,proto3" json:"Time,omitempty"`
}

func (x *NoiseLevelReading) Reset() {
	*x = NoiseLevelReading{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noise_level_reading_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoiseLevelReading) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoiseLevelReading) ProtoMessage() {}

func (x *NoiseLevelReading) ProtoReflect() protoreflect.Message {
	mi := &file_noise_level_reading_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoiseLevelReading.ProtoReflect.Descriptor instead.
func (*NoiseLevelReading) Descriptor() ([]byte, []int) {
	return file_noise_level_reading_proto_rawDescGZIP(), []int{0}
}

func (x *NoiseLevelReading) GetReadingId() string {
	if x != nil {
		return x.ReadingId
	}
	return ""
}

func (x *NoiseLevelReading) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *NoiseLevelReading) GetDecibelLevel() float32 {
	if x != nil {
		return x.DecibelLevel
	}
	return 0
}

func (x *NoiseLevelReading) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

type NoiseLevelReadingList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoiseLevelReadings []*NoiseLevelReading `protobuf:"bytes,1,rep,name=NoiseLevelReadings,proto3" json:"NoiseLevelReadings,omitempty"`
}

func (x *NoiseLevelReadingList) Reset() {
	*x = NoiseLevelReadingList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noise_level_reading_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoiseLevelReadingList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoiseLevelReadingList) ProtoMessage() {}

func (x *NoiseLevelReadingList) ProtoReflect() protoreflect.Message {
	mi := &file_noise_level_reading_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoiseLevelReadingList.ProtoReflect.Descriptor instead.
func (*NoiseLevelReadingList) Descriptor() ([]byte, []int) {
	return file_noise_level_reading_proto_rawDescGZIP(), []int{1}
}

func (x *NoiseLevelReadingList) GetNoiseLevelReadings() []*NoiseLevelReading {
	if x != nil {
		return x.NoiseLevelReadings
	}
	return nil
}

type ZoneFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneId string `protobuf:"bytes,1,opt,name=ZoneId,proto3" json:"ZoneId,omitempty"`
	Time   string `protobuf:"bytes,2,opt,name=Time,proto3" json:"Time,omitempty"`
}

func (x *ZoneFilter) Reset() {
	*x = ZoneFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noise_level_reading_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneFilter) ProtoMessage() {}

func (x *ZoneFilter) ProtoReflect() protoreflect.Message {
	mi := &file_noise_level_reading_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneFilter.ProtoReflect.Descriptor instead.
func (*ZoneFilter) Descriptor() ([]byte, []int) {
	return file_noise_level_reading_proto_rawDescGZIP(), []int{2}
}

func (x *ZoneFilter) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *ZoneFilter) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

var File_noise_level_reading_proto protoreflect.FileDescriptor

var file_noise_level_reading_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6e, 0x6f, 0x69, 0x73, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x72, 0x65,
	0x61, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x11, 0x4e, 0x6f, 0x69, 0x73, 0x65, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x44, 0x65, 0x63, 0x69, 0x62, 0x65, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x44, 0x65, 0x63, 0x69, 0x62, 0x65, 0x6c, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x73, 0x0a, 0x15, 0x4e, 0x6f, 0x69, 0x73, 0x65, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x5a, 0x0a, 0x12, 0x4e, 0x6f, 0x69, 0x73, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x6f, 0x69, 0x73, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x12, 0x4e, 0x6f, 0x69, 0x73, 0x65, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x38, 0x0a, 0x0a, 0x5a,
	0x6f, 0x6e, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x5a, 0x6f, 0x6e,
	0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xf1, 0x01, 0x0a, 0x18, 0x4e, 0x6f, 0x69, 0x73, 0x65, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x66, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x69, 0x73,
	0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x2a, 0x2e,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x6f, 0x69, 0x73, 0x65, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x1a, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x69, 0x73, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x23, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x5a, 0x6f, 0x6e,
	0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x4e, 0x6f, 0x69, 0x73, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_noise_level_reading_proto_rawDescOnce sync.Once
	file_noise_level_reading_proto_rawDescData = file_noise_level_reading_proto_rawDesc
)

func file_noise_level_reading_proto_rawDescGZIP() []byte {
	file_noise_level_reading_proto_rawDescOnce.Do(func() {
		file_noise_level_reading_proto_rawDescData = protoimpl.X.CompressGZIP(file_noise_level_reading_proto_rawDescData)
	})
	return file_noise_level_reading_proto_rawDescData
}

var file_noise_level_reading_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_noise_level_reading_proto_goTypes = []interface{}{
	(*NoiseLevelReading)(nil),     // 0: enviromental_monitoring.NoiseLevelReading
	(*NoiseLevelReadingList)(nil), // 1: enviromental_monitoring.NoiseLevelReadingList
	(*ZoneFilter)(nil),            // 2: enviromental_monitoring.ZoneFilter
	(*Void)(nil),                  // 3: enviromental_monitoring.Void
}
var file_noise_level_reading_proto_depIdxs = []int32{
	0, // 0: enviromental_monitoring.NoiseLevelReadingList.NoiseLevelReadings:type_name -> enviromental_monitoring.NoiseLevelReading
	0, // 1: enviromental_monitoring.NoiseLevelReadingService.CreateNoiseLevelReading:input_type -> enviromental_monitoring.NoiseLevelReading
	2, // 2: enviromental_monitoring.NoiseLevelReadingService.GetNoiseLevelReading:input_type -> enviromental_monitoring.ZoneFilter
	3, // 3: enviromental_monitoring.NoiseLevelReadingService.CreateNoiseLevelReading:output_type -> enviromental_monitoring.Void
	1, // 4: enviromental_monitoring.NoiseLevelReadingService.GetNoiseLevelReading:output_type -> enviromental_monitoring.NoiseLevelReadingList
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_noise_level_reading_proto_init() }
func file_noise_level_reading_proto_init() {
	if File_noise_level_reading_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_noise_level_reading_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoiseLevelReading); i {
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
		file_noise_level_reading_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoiseLevelReadingList); i {
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
		file_noise_level_reading_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneFilter); i {
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
			RawDescriptor: file_noise_level_reading_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_noise_level_reading_proto_goTypes,
		DependencyIndexes: file_noise_level_reading_proto_depIdxs,
		MessageInfos:      file_noise_level_reading_proto_msgTypes,
	}.Build()
	File_noise_level_reading_proto = out.File
	file_noise_level_reading_proto_rawDesc = nil
	file_noise_level_reading_proto_goTypes = nil
	file_noise_level_reading_proto_depIdxs = nil
}
