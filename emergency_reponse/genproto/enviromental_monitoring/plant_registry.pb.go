// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: plant_registry.proto

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

type PlantRegistry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegistryId  string `protobuf:"bytes,1,opt,name=RegistryId,proto3" json:"RegistryId,omitempty"`
	SpaceId     string `protobuf:"bytes,2,opt,name=SpaceId,proto3" json:"SpaceId,omitempty"`
	SpeciesName string `protobuf:"bytes,3,opt,name=SpeciesName,proto3" json:"SpeciesName,omitempty"`
	Quantity    int32  `protobuf:"varint,4,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	PlantDate   string `protobuf:"bytes,5,opt,name=PlantDate,proto3" json:"PlantDate,omitempty"`
}

func (x *PlantRegistry) Reset() {
	*x = PlantRegistry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plant_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlantRegistry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlantRegistry) ProtoMessage() {}

func (x *PlantRegistry) ProtoReflect() protoreflect.Message {
	mi := &file_plant_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlantRegistry.ProtoReflect.Descriptor instead.
func (*PlantRegistry) Descriptor() ([]byte, []int) {
	return file_plant_registry_proto_rawDescGZIP(), []int{0}
}

func (x *PlantRegistry) GetRegistryId() string {
	if x != nil {
		return x.RegistryId
	}
	return ""
}

func (x *PlantRegistry) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *PlantRegistry) GetSpeciesName() string {
	if x != nil {
		return x.SpeciesName
	}
	return ""
}

func (x *PlantRegistry) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *PlantRegistry) GetPlantDate() string {
	if x != nil {
		return x.PlantDate
	}
	return ""
}

type PlantRegistryList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlantRegistries []*PlantRegistry `protobuf:"bytes,1,rep,name=PlantRegistries,proto3" json:"PlantRegistries,omitempty"`
}

func (x *PlantRegistryList) Reset() {
	*x = PlantRegistryList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plant_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlantRegistryList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlantRegistryList) ProtoMessage() {}

func (x *PlantRegistryList) ProtoReflect() protoreflect.Message {
	mi := &file_plant_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlantRegistryList.ProtoReflect.Descriptor instead.
func (*PlantRegistryList) Descriptor() ([]byte, []int) {
	return file_plant_registry_proto_rawDescGZIP(), []int{1}
}

func (x *PlantRegistryList) GetPlantRegistries() []*PlantRegistry {
	if x != nil {
		return x.PlantRegistries
	}
	return nil
}

type PlantRegistryFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegistryId  string `protobuf:"bytes,1,opt,name=RegistryId,proto3" json:"RegistryId,omitempty"`
	SpaceId     string `protobuf:"bytes,2,opt,name=SpaceId,proto3" json:"SpaceId,omitempty"`
	SpeciesName string `protobuf:"bytes,3,opt,name=SpeciesName,proto3" json:"SpeciesName,omitempty"`
	Quantity    int32  `protobuf:"varint,4,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	PlantDate   string `protobuf:"bytes,5,opt,name=PlantDate,proto3" json:"PlantDate,omitempty"`
}

func (x *PlantRegistryFilter) Reset() {
	*x = PlantRegistryFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plant_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlantRegistryFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlantRegistryFilter) ProtoMessage() {}

func (x *PlantRegistryFilter) ProtoReflect() protoreflect.Message {
	mi := &file_plant_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlantRegistryFilter.ProtoReflect.Descriptor instead.
func (*PlantRegistryFilter) Descriptor() ([]byte, []int) {
	return file_plant_registry_proto_rawDescGZIP(), []int{2}
}

func (x *PlantRegistryFilter) GetRegistryId() string {
	if x != nil {
		return x.RegistryId
	}
	return ""
}

func (x *PlantRegistryFilter) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *PlantRegistryFilter) GetSpeciesName() string {
	if x != nil {
		return x.SpeciesName
	}
	return ""
}

func (x *PlantRegistryFilter) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *PlantRegistryFilter) GetPlantDate() string {
	if x != nil {
		return x.PlantDate
	}
	return ""
}

var File_plant_registry_proto protoreflect.FileDescriptor

var file_plant_registry_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x1a,
	0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01,
	0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12,
	0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x65, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x22, 0x65, 0x0a, 0x11, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x0f, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0xab, 0x01, 0x0a,
	0x13, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x32, 0x9a, 0x03, 0x0a, 0x14, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x1a, 0x1d, 0x2e,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x70,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x1a, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00,
	0x12, 0x5e, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x26, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x1a,
	0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00,
	0x12, 0x55, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plant_registry_proto_rawDescOnce sync.Once
	file_plant_registry_proto_rawDescData = file_plant_registry_proto_rawDesc
)

func file_plant_registry_proto_rawDescGZIP() []byte {
	file_plant_registry_proto_rawDescOnce.Do(func() {
		file_plant_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_plant_registry_proto_rawDescData)
	})
	return file_plant_registry_proto_rawDescData
}

var file_plant_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_plant_registry_proto_goTypes = []interface{}{
	(*PlantRegistry)(nil),       // 0: enviromental_monitoring.PlantRegistry
	(*PlantRegistryList)(nil),   // 1: enviromental_monitoring.PlantRegistryList
	(*PlantRegistryFilter)(nil), // 2: enviromental_monitoring.PlantRegistryFilter
	(*ById)(nil),                // 3: enviromental_monitoring.ById
	(*Void)(nil),                // 4: enviromental_monitoring.Void
}
var file_plant_registry_proto_depIdxs = []int32{
	0, // 0: enviromental_monitoring.PlantRegistryList.PlantRegistries:type_name -> enviromental_monitoring.PlantRegistry
	0, // 1: enviromental_monitoring.PlantRegistryService.RegisterPlants:input_type -> enviromental_monitoring.PlantRegistry
	2, // 2: enviromental_monitoring.PlantRegistryService.GetPlantRegistries:input_type -> enviromental_monitoring.PlantRegistryFilter
	0, // 3: enviromental_monitoring.PlantRegistryService.UpdatePlantRegistry:input_type -> enviromental_monitoring.PlantRegistry
	3, // 4: enviromental_monitoring.PlantRegistryService.DeletePlantRegistry:input_type -> enviromental_monitoring.ById
	4, // 5: enviromental_monitoring.PlantRegistryService.RegisterPlants:output_type -> enviromental_monitoring.Void
	1, // 6: enviromental_monitoring.PlantRegistryService.GetPlantRegistries:output_type -> enviromental_monitoring.PlantRegistryList
	4, // 7: enviromental_monitoring.PlantRegistryService.UpdatePlantRegistry:output_type -> enviromental_monitoring.Void
	4, // 8: enviromental_monitoring.PlantRegistryService.DeletePlantRegistry:output_type -> enviromental_monitoring.Void
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_plant_registry_proto_init() }
func file_plant_registry_proto_init() {
	if File_plant_registry_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_plant_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlantRegistry); i {
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
		file_plant_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlantRegistryList); i {
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
		file_plant_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlantRegistryFilter); i {
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
			RawDescriptor: file_plant_registry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plant_registry_proto_goTypes,
		DependencyIndexes: file_plant_registry_proto_depIdxs,
		MessageInfos:      file_plant_registry_proto_msgTypes,
	}.Build()
	File_plant_registry_proto = out.File
	file_plant_registry_proto_rawDesc = nil
	file_plant_registry_proto_goTypes = nil
	file_plant_registry_proto_depIdxs = nil
}
