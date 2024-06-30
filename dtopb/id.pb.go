// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: dtopb/id.proto

package dtopb

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

type IDString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IDString) Reset() {
	*x = IDString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dtopb_id_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDString) ProtoMessage() {}

func (x *IDString) ProtoReflect() protoreflect.Message {
	mi := &file_dtopb_id_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDString.ProtoReflect.Descriptor instead.
func (*IDString) Descriptor() ([]byte, []int) {
	return file_dtopb_id_proto_rawDescGZIP(), []int{0}
}

func (x *IDString) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type IDNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IDNumber) Reset() {
	*x = IDNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dtopb_id_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDNumber) ProtoMessage() {}

func (x *IDNumber) ProtoReflect() protoreflect.Message {
	mi := &file_dtopb_id_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDNumber.ProtoReflect.Descriptor instead.
func (*IDNumber) Descriptor() ([]byte, []int) {
	return file_dtopb_id_proto_rawDescGZIP(), []int{1}
}

func (x *IDNumber) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_dtopb_id_proto protoreflect.FileDescriptor

var file_dtopb_id_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x74, 0x6f, 0x70, 0x62, 0x2f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x64, 0x74, 0x6f, 0x70, 0x62, 0x22, 0x1a, 0x0a, 0x08, 0x49, 0x44, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x08, 0x49, 0x44, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x42,
	0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6f,
	0x66, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x64, 0x74, 0x6f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dtopb_id_proto_rawDescOnce sync.Once
	file_dtopb_id_proto_rawDescData = file_dtopb_id_proto_rawDesc
)

func file_dtopb_id_proto_rawDescGZIP() []byte {
	file_dtopb_id_proto_rawDescOnce.Do(func() {
		file_dtopb_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_dtopb_id_proto_rawDescData)
	})
	return file_dtopb_id_proto_rawDescData
}

var file_dtopb_id_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dtopb_id_proto_goTypes = []interface{}{
	(*IDString)(nil), // 0: dtopb.IDString
	(*IDNumber)(nil), // 1: dtopb.IDNumber
}
var file_dtopb_id_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dtopb_id_proto_init() }
func file_dtopb_id_proto_init() {
	if File_dtopb_id_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dtopb_id_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDString); i {
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
		file_dtopb_id_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDNumber); i {
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
			RawDescriptor: file_dtopb_id_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dtopb_id_proto_goTypes,
		DependencyIndexes: file_dtopb_id_proto_depIdxs,
		MessageInfos:      file_dtopb_id_proto_msgTypes,
	}.Build()
	File_dtopb_id_proto = out.File
	file_dtopb_id_proto_rawDesc = nil
	file_dtopb_id_proto_goTypes = nil
	file_dtopb_id_proto_depIdxs = nil
}