// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: accore/v1/phone-service.proto

package v1

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	pb "github.com/zofs/protogen/accore/pb"
	dtopb "github.com/zofs/protogen/dtopb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_accore_v1_phone_service_proto protoreflect.FileDescriptor

var file_accore_v1_phone_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x63, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x61, 0x63, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x62,
	0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x64, 0x74,
	0x6f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xea, 0x03, 0x0a, 0x0c, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5a, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x11, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x73, 0x22, 0x27, 0x92, 0x41, 0x0e, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x5b, 0x0a, 0x03,
	0x41, 0x64, 0x64, 0x12, 0x10, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2a, 0x92,
	0x41, 0x0e, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x59, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x09, 0x2e, 0x64, 0x74, 0x6f, 0x70, 0x62, 0x2e, 0x49, 0x44, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2c, 0x92, 0x41, 0x0e, 0x62, 0x0c, 0x0a, 0x0a, 0x0a,
	0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a,
	0x13, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5d, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x12, 0x09, 0x2e, 0x64, 0x74, 0x6f, 0x70, 0x62, 0x2e, 0x49, 0x44, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2c, 0x92, 0x41, 0x0e, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06,
	0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x32, 0x13,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x67, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x2f, 0x92, 0x41, 0x0e,
	0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x42, 0xcc, 0x01, 0x92,
	0x41, 0xa4, 0x01, 0x12, 0x47, 0x0a, 0x0d, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x20, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2a, 0x31, 0x0a, 0x0b, 0x4d, 0x49, 0x54, 0x20, 0x4c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x72, 0x6d, 0x73,
	0x75, 0x62, 0x65, 0x6b, 0x74, 0x69, 0x2e, 0x6d, 0x69, 0x74, 0x2d, 0x6c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x5a, 0x59, 0x0a, 0x57,
	0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x3c,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6f, 0x66, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65,
	0x6e, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_accore_v1_phone_service_proto_goTypes = []interface{}{
	(*empty.Empty)(nil), // 0: google.protobuf.Empty
	(*pb.Phone)(nil),    // 1: accore.pb.Phone
	(*dtopb.ID)(nil),    // 2: dtopb.ID
	(*pb.Phones)(nil),   // 3: accore.pb.Phones
}
var file_accore_v1_phone_service_proto_depIdxs = []int32{
	0, // 0: accore.v1.PhoneService.List:input_type -> google.protobuf.Empty
	1, // 1: accore.v1.PhoneService.Add:input_type -> accore.pb.Phone
	2, // 2: accore.v1.PhoneService.Delete:input_type -> dtopb.ID
	2, // 3: accore.v1.PhoneService.SetPrimary:input_type -> dtopb.ID
	0, // 4: accore.v1.PhoneService.GetPrimary:input_type -> google.protobuf.Empty
	3, // 5: accore.v1.PhoneService.List:output_type -> accore.pb.Phones
	0, // 6: accore.v1.PhoneService.Add:output_type -> google.protobuf.Empty
	0, // 7: accore.v1.PhoneService.Delete:output_type -> google.protobuf.Empty
	0, // 8: accore.v1.PhoneService.SetPrimary:output_type -> google.protobuf.Empty
	1, // 9: accore.v1.PhoneService.GetPrimary:output_type -> accore.pb.Phone
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_accore_v1_phone_service_proto_init() }
func file_accore_v1_phone_service_proto_init() {
	if File_accore_v1_phone_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_accore_v1_phone_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accore_v1_phone_service_proto_goTypes,
		DependencyIndexes: file_accore_v1_phone_service_proto_depIdxs,
	}.Build()
	File_accore_v1_phone_service_proto = out.File
	file_accore_v1_phone_service_proto_rawDesc = nil
	file_accore_v1_phone_service_proto_goTypes = nil
	file_accore_v1_phone_service_proto_depIdxs = nil
}
