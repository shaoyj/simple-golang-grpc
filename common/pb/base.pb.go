// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// - protoc             v4.25.1
// source: base.proto

package pb

import (
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

var File_base_proto protoreflect.FileDescriptor

var file_base_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70,
	0x69, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x66, 0x62, 0x5f, 0x70, 0x62, 0x2f, 0x7a, 0x64, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5b, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x07, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x12, 0x0d, 0x2e, 0x66, 0x62, 0x5f, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x1a,
	0x11, 0x2e, 0x66, 0x62, 0x5f, 0x70, 0x62, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x42, 0x2e, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x72, 0x69, 0x6f, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x42, 0x0a, 0x47, 0x72, 0x70, 0x63, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x61, 0x70, 0x69, 0x5f, 0x62,
	0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_base_proto_goTypes = []interface{}{
	(*ComReq)(nil),     // 0: ComReq
	(*ByteResult)(nil), // 1: ByteResult
}
var file_base_proto_depIdxs = []int32{
	0, // 0: api_base.BaseService.execute:input_type -> ComReq
	1, // 1: api_base.BaseService.execute:output_type -> ByteResult
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_base_proto_init() }
func file_base_proto_init() {
	if File_base_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_base_proto_goTypes,
		DependencyIndexes: file_base_proto_depIdxs,
	}.Build()
	File_base_proto = out.File
	file_base_proto_rawDesc = nil
	file_base_proto_goTypes = nil
	file_base_proto_depIdxs = nil
}
