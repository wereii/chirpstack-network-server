// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: downlink_frame.proto

package storage

import (
	gw "github.com/wereii/chirpstack-api/go/v3/gw"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DownlinkFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token.
	Token uint32 `protobuf:"varint,1,opt,name=token,proto3" json:"token,omitempty"`
	// DevEUI.
	DevEui []byte `protobuf:"bytes,2,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	// Device queue item ID.
	DeviceQueueItemId int64 `protobuf:"varint,3,opt,name=device_queue_item_id,json=deviceQueueItemId,proto3" json:"device_queue_item_id,omitempty"`
	// Multicast Group ID.
	MulticastGroupId []byte `protobuf:"bytes,4,opt,name=multicast_group_id,json=multicastGroupId,proto3" json:"multicast_group_id,omitempty"`
	// Multicast queue item ID.
	MulticastQueueItemId int64 `protobuf:"varint,5,opt,name=multicast_queue_item_id,json=multicastQueueItemId,proto3" json:"multicast_queue_item_id,omitempty"`
	// Downlink frames.
	DownlinkFrame *gw.DownlinkFrame `protobuf:"bytes,6,opt,name=downlink_frame,json=downlinkFrame,proto3" json:"downlink_frame,omitempty"`
	// Routing Profile ID.
	RoutingProfileId []byte `protobuf:"bytes,7,opt,name=routing_profile_id,json=routingProfileId,proto3" json:"routing_profile_id,omitempty"`
	// Encrypted FOpts (LoRaWAN 1.1).
	EncryptedFopts bool `protobuf:"varint,8,opt,name=encrypted_fopts,json=encryptedFopts,proto3" json:"encrypted_fopts,omitempty"`
	// Network session encryption key (for FOpts).
	NwkSEncKey []byte `protobuf:"bytes,9,opt,name=nwk_s_enc_key,json=nwkSEncKey,proto3" json:"nwk_s_enc_key,omitempty"`
}

func (x *DownlinkFrame) Reset() {
	*x = DownlinkFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_downlink_frame_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownlinkFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownlinkFrame) ProtoMessage() {}

func (x *DownlinkFrame) ProtoReflect() protoreflect.Message {
	mi := &file_downlink_frame_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownlinkFrame.ProtoReflect.Descriptor instead.
func (*DownlinkFrame) Descriptor() ([]byte, []int) {
	return file_downlink_frame_proto_rawDescGZIP(), []int{0}
}

func (x *DownlinkFrame) GetToken() uint32 {
	if x != nil {
		return x.Token
	}
	return 0
}

func (x *DownlinkFrame) GetDevEui() []byte {
	if x != nil {
		return x.DevEui
	}
	return nil
}

func (x *DownlinkFrame) GetDeviceQueueItemId() int64 {
	if x != nil {
		return x.DeviceQueueItemId
	}
	return 0
}

func (x *DownlinkFrame) GetMulticastGroupId() []byte {
	if x != nil {
		return x.MulticastGroupId
	}
	return nil
}

func (x *DownlinkFrame) GetMulticastQueueItemId() int64 {
	if x != nil {
		return x.MulticastQueueItemId
	}
	return 0
}

func (x *DownlinkFrame) GetDownlinkFrame() *gw.DownlinkFrame {
	if x != nil {
		return x.DownlinkFrame
	}
	return nil
}

func (x *DownlinkFrame) GetRoutingProfileId() []byte {
	if x != nil {
		return x.RoutingProfileId
	}
	return nil
}

func (x *DownlinkFrame) GetEncryptedFopts() bool {
	if x != nil {
		return x.EncryptedFopts
	}
	return false
}

func (x *DownlinkFrame) GetNwkSEncKey() []byte {
	if x != nil {
		return x.NwkSEncKey
	}
	return nil
}

var File_downlink_frame_proto protoreflect.FileDescriptor

var file_downlink_frame_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a,
	0x0b, 0x67, 0x77, 0x2f, 0x67, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x03, 0x0a,
	0x0d, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x64, 0x65, 0x76, 0x45, 0x75, 0x69, 0x12, 0x2f, 0x0a,
	0x14, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x2c,
	0x0a, 0x12, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x63, 0x61, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x17,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x63, 0x61, 0x73, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x5f,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x77,
	0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x0d,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a,
	0x12, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x65,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x6f, 0x70, 0x74, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x46,
	0x6f, 0x70, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0d, 0x6e, 0x77, 0x6b, 0x5f, 0x73, 0x5f, 0x65, 0x6e,
	0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6e, 0x77, 0x6b,
	0x53, 0x45, 0x6e, 0x63, 0x4b, 0x65, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_downlink_frame_proto_rawDescOnce sync.Once
	file_downlink_frame_proto_rawDescData = file_downlink_frame_proto_rawDesc
)

func file_downlink_frame_proto_rawDescGZIP() []byte {
	file_downlink_frame_proto_rawDescOnce.Do(func() {
		file_downlink_frame_proto_rawDescData = protoimpl.X.CompressGZIP(file_downlink_frame_proto_rawDescData)
	})
	return file_downlink_frame_proto_rawDescData
}

var file_downlink_frame_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_downlink_frame_proto_goTypes = []interface{}{
	(*DownlinkFrame)(nil),    // 0: storage.DownlinkFrame
	(*gw.DownlinkFrame)(nil), // 1: gw.DownlinkFrame
}
var file_downlink_frame_proto_depIdxs = []int32{
	1, // 0: storage.DownlinkFrame.downlink_frame:type_name -> gw.DownlinkFrame
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_downlink_frame_proto_init() }
func file_downlink_frame_proto_init() {
	if File_downlink_frame_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_downlink_frame_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownlinkFrame); i {
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
			RawDescriptor: file_downlink_frame_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_downlink_frame_proto_goTypes,
		DependencyIndexes: file_downlink_frame_proto_depIdxs,
		MessageInfos:      file_downlink_frame_proto_msgTypes,
	}.Build()
	File_downlink_frame_proto = out.File
	file_downlink_frame_proto_rawDesc = nil
	file_downlink_frame_proto_goTypes = nil
	file_downlink_frame_proto_depIdxs = nil
}
