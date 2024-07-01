// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: city_submodule/citizen/notification.proto

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

type NotificationPref struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreferenceId     string `protobuf:"bytes,1,opt,name=PreferenceId,proto3" json:"PreferenceId,omitempty"`
	CitizenId        string `protobuf:"bytes,2,opt,name=CitizenId,proto3" json:"CitizenId,omitempty"`
	NotificationType string `protobuf:"bytes,3,opt,name=NotificationType,proto3" json:"NotificationType,omitempty"`
	IsEnabled        bool   `protobuf:"varint,4,opt,name=IsEnabled,proto3" json:"IsEnabled,omitempty"`
}

func (x *NotificationPref) Reset() {
	*x = NotificationPref{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_submodule_citizen_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationPref) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationPref) ProtoMessage() {}

func (x *NotificationPref) ProtoReflect() protoreflect.Message {
	mi := &file_city_submodule_citizen_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationPref.ProtoReflect.Descriptor instead.
func (*NotificationPref) Descriptor() ([]byte, []int) {
	return file_city_submodule_citizen_notification_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationPref) GetPreferenceId() string {
	if x != nil {
		return x.PreferenceId
	}
	return ""
}

func (x *NotificationPref) GetCitizenId() string {
	if x != nil {
		return x.CitizenId
	}
	return ""
}

func (x *NotificationPref) GetNotificationType() string {
	if x != nil {
		return x.NotificationType
	}
	return ""
}

func (x *NotificationPref) GetIsEnabled() bool {
	if x != nil {
		return x.IsEnabled
	}
	return false
}

var File_city_submodule_citizen_notification_proto protoreflect.FileDescriptor

var file_city_submodule_citizen_notification_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2f, 0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x69, 0x74,
	0x69, 0x7a, 0x65, 0x6e, 0x1a, 0x23, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2f, 0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x10, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x66, 0x12, 0x22,
	0x0a, 0x0c, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x49, 0x64,
	0x12, 0x2a, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x49, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x49, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x32, 0xdf, 0x01, 0x0a, 0x1a, 0x43,
	0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x13, 0x53, 0x65, 0x74,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x66,
	0x12, 0x19, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x66, 0x1a, 0x0d, 0x2e, 0x63, 0x69,
	0x74, 0x69, 0x7a, 0x65, 0x6e, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x16,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x65, 0x66, 0x12, 0x19, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65,
	0x66, 0x1a, 0x0d, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x2e, 0x56, 0x6f, 0x69, 0x64,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x66, 0x12, 0x0d, 0x2e, 0x63,
	0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0d, 0x2e, 0x63, 0x69,
	0x74, 0x69, 0x7a, 0x65, 0x6e, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x69, 0x74, 0x69, 0x7a, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_submodule_citizen_notification_proto_rawDescOnce sync.Once
	file_city_submodule_citizen_notification_proto_rawDescData = file_city_submodule_citizen_notification_proto_rawDesc
)

func file_city_submodule_citizen_notification_proto_rawDescGZIP() []byte {
	file_city_submodule_citizen_notification_proto_rawDescOnce.Do(func() {
		file_city_submodule_citizen_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_submodule_citizen_notification_proto_rawDescData)
	})
	return file_city_submodule_citizen_notification_proto_rawDescData
}

var file_city_submodule_citizen_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_city_submodule_citizen_notification_proto_goTypes = []any{
	(*NotificationPref)(nil), // 0: citizen.NotificationPref
	(*ById)(nil),             // 1: citizen.ById
	(*Void)(nil),             // 2: citizen.Void
}
var file_city_submodule_citizen_notification_proto_depIdxs = []int32{
	0, // 0: citizen.CitizenNotificationService.SetNotificationPref:input_type -> citizen.NotificationPref
	0, // 1: citizen.CitizenNotificationService.UpdateNotificationPref:input_type -> citizen.NotificationPref
	1, // 2: citizen.CitizenNotificationService.DeleteNotificationPref:input_type -> citizen.ById
	2, // 3: citizen.CitizenNotificationService.SetNotificationPref:output_type -> citizen.Void
	2, // 4: citizen.CitizenNotificationService.UpdateNotificationPref:output_type -> citizen.Void
	2, // 5: citizen.CitizenNotificationService.DeleteNotificationPref:output_type -> citizen.Void
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_city_submodule_citizen_notification_proto_init() }
func file_city_submodule_citizen_notification_proto_init() {
	if File_city_submodule_citizen_notification_proto != nil {
		return
	}
	file_city_submodule_citizen_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_city_submodule_citizen_notification_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NotificationPref); i {
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
			RawDescriptor: file_city_submodule_citizen_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_submodule_citizen_notification_proto_goTypes,
		DependencyIndexes: file_city_submodule_citizen_notification_proto_depIdxs,
		MessageInfos:      file_city_submodule_citizen_notification_proto_msgTypes,
	}.Build()
	File_city_submodule_citizen_notification_proto = out.File
	file_city_submodule_citizen_notification_proto_rawDesc = nil
	file_city_submodule_citizen_notification_proto_goTypes = nil
	file_city_submodule_citizen_notification_proto_depIdxs = nil
}