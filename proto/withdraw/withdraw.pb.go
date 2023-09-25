// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: proto/withdraw/withdraw.proto

package withdraw

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type WithdrawBo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscordUserId   string `protobuf:"bytes,1,opt,name=discord_user_id,json=discordUserId,proto3" json:"discord_user_id,omitempty"`
	WithdrawAddress string `protobuf:"bytes,3,opt,name=withdraw_address,json=withdrawAddress,proto3" json:"withdraw_address,omitempty"`
}

func (x *WithdrawBo) Reset() {
	*x = WithdrawBo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_withdraw_withdraw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawBo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawBo) ProtoMessage() {}

func (x *WithdrawBo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_withdraw_withdraw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawBo.ProtoReflect.Descriptor instead.
func (*WithdrawBo) Descriptor() ([]byte, []int) {
	return file_proto_withdraw_withdraw_proto_rawDescGZIP(), []int{0}
}

func (x *WithdrawBo) GetDiscordUserId() string {
	if x != nil {
		return x.DiscordUserId
	}
	return ""
}

func (x *WithdrawBo) GetWithdrawAddress() string {
	if x != nil {
		return x.WithdrawAddress
	}
	return ""
}

var File_proto_withdraw_withdraw_proto protoreflect.FileDescriptor

var file_proto_withdraw_withdraw_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x2f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x0a, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x42, 0x6f, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10,
	0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x50, 0x0a, 0x0f, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x08, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x42, 0x6f, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_withdraw_withdraw_proto_rawDescOnce sync.Once
	file_proto_withdraw_withdraw_proto_rawDescData = file_proto_withdraw_withdraw_proto_rawDesc
)

func file_proto_withdraw_withdraw_proto_rawDescGZIP() []byte {
	file_proto_withdraw_withdraw_proto_rawDescOnce.Do(func() {
		file_proto_withdraw_withdraw_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_withdraw_withdraw_proto_rawDescData)
	})
	return file_proto_withdraw_withdraw_proto_rawDescData
}

var file_proto_withdraw_withdraw_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_withdraw_withdraw_proto_goTypes = []interface{}{
	(*WithdrawBo)(nil),             // 0: proto.WithdrawBo
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
}
var file_proto_withdraw_withdraw_proto_depIdxs = []int32{
	0, // 0: proto.WithdrawRequest.Withdraw:input_type -> proto.WithdrawBo
	1, // 1: proto.WithdrawRequest.Withdraw:output_type -> google.protobuf.StringValue
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_withdraw_withdraw_proto_init() }
func file_proto_withdraw_withdraw_proto_init() {
	if File_proto_withdraw_withdraw_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_withdraw_withdraw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawBo); i {
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
			RawDescriptor: file_proto_withdraw_withdraw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_withdraw_withdraw_proto_goTypes,
		DependencyIndexes: file_proto_withdraw_withdraw_proto_depIdxs,
		MessageInfos:      file_proto_withdraw_withdraw_proto_msgTypes,
	}.Build()
	File_proto_withdraw_withdraw_proto = out.File
	file_proto_withdraw_withdraw_proto_rawDesc = nil
	file_proto_withdraw_withdraw_proto_goTypes = nil
	file_proto_withdraw_withdraw_proto_depIdxs = nil
}