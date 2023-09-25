// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: proto/player/player.proto

package player

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlayerInfoVo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DiscordUserId string                 `protobuf:"bytes,2,opt,name=discord_user_id,json=discordUserId,proto3" json:"discord_user_id,omitempty"`
	WalletAddress string                 `protobuf:"bytes,3,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
	WalletValue   int64                  `protobuf:"varint,4,opt,name=wallet_value,json=walletValue,proto3" json:"wallet_value,omitempty"`
	CreateAt      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
}

func (x *PlayerInfoVo) Reset() {
	*x = PlayerInfoVo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_player_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInfoVo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfoVo) ProtoMessage() {}

func (x *PlayerInfoVo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_player_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfoVo.ProtoReflect.Descriptor instead.
func (*PlayerInfoVo) Descriptor() ([]byte, []int) {
	return file_proto_player_player_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerInfoVo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PlayerInfoVo) GetDiscordUserId() string {
	if x != nil {
		return x.DiscordUserId
	}
	return ""
}

func (x *PlayerInfoVo) GetWalletAddress() string {
	if x != nil {
		return x.WalletAddress
	}
	return ""
}

func (x *PlayerInfoVo) GetWalletValue() int64 {
	if x != nil {
		return x.WalletValue
	}
	return 0
}

func (x *PlayerInfoVo) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

var File_proto_player_player_proto protoreflect.FileDescriptor

var file_proto_player_player_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x56, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x32,
	0x56, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x45, 0x0a, 0x0e, 0x53, 0x68, 0x6f, 0x77, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x56, 0x6f, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_player_player_proto_rawDescOnce sync.Once
	file_proto_player_player_proto_rawDescData = file_proto_player_player_proto_rawDesc
)

func file_proto_player_player_proto_rawDescGZIP() []byte {
	file_proto_player_player_proto_rawDescOnce.Do(func() {
		file_proto_player_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_player_player_proto_rawDescData)
	})
	return file_proto_player_player_proto_rawDescData
}

var file_proto_player_player_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_player_player_proto_goTypes = []interface{}{
	(*PlayerInfoVo)(nil),           // 0: proto.PlayerInfoVo
	(*timestamppb.Timestamp)(nil),  // 1: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
}
var file_proto_player_player_proto_depIdxs = []int32{
	1, // 0: proto.PlayerInfoVo.create_at:type_name -> google.protobuf.Timestamp
	2, // 1: proto.PlayerRequest.ShowPlayerInfo:input_type -> google.protobuf.StringValue
	0, // 2: proto.PlayerRequest.ShowPlayerInfo:output_type -> proto.PlayerInfoVo
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_player_player_proto_init() }
func file_proto_player_player_proto_init() {
	if File_proto_player_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_player_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInfoVo); i {
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
			RawDescriptor: file_proto_player_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_player_player_proto_goTypes,
		DependencyIndexes: file_proto_player_player_proto_depIdxs,
		MessageInfos:      file_proto_player_player_proto_msgTypes,
	}.Build()
	File_proto_player_player_proto = out.File
	file_proto_player_player_proto_rawDesc = nil
	file_proto_player_player_proto_goTypes = nil
	file_proto_player_player_proto_depIdxs = nil
}