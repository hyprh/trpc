// Tencent is pleased to support the open source community by making tRPC available.
// Copyright (C) 2023 THL A29 Limited, a Tencent company. All rights reserved.
// If you have downloaded a copy of the tRPC source code from Tencent,
// please note that tRPC source code is licensed under the Apache 2.0 License,
// A copy of the Apache 2.0 License is included in this file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.27.1
// source: trpc/proto/trpc_options.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_trpc_proto_trpc_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50001,
		Name:          "trpc.alias",
		Tag:           "bytes,50001,opt,name=alias",
		Filename:      "trpc/proto/trpc_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50050,
		Name:          "trpc.go_tag",
		Tag:           "bytes,50050,opt,name=go_tag",
		Filename:      "trpc/proto/trpc_options.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional string alias = 50001;
	E_Alias = &file_trpc_proto_trpc_options_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string go_tag = 50050;
	E_GoTag = &file_trpc_proto_trpc_options_proto_extTypes[1]
)

var File_trpc_proto_trpc_options_proto protoreflect.FileDescriptor

const file_trpc_proto_trpc_options_proto_rawDesc = "" +
	"\n" +
	"\x1dtrpc/proto/trpc_options.proto\x12\x04trpc\x1a google/protobuf/descriptor.proto:6\n" +
	"\x05alias\x12\x1e.google.protobuf.MethodOptions\x18ц\x03 \x01(\tR\x05alias:6\n" +
	"\x06go_tag\x12\x1d.google.protobuf.FieldOptions\x18\x82\x87\x03 \x01(\tR\x05goTagBO\n" +
	"\x1dcom.tencent.trpc.protobuf.extZ.trpc.group/trpc/trpc-protocol/pb/go/trpc/protob\x06proto3"

var file_trpc_proto_trpc_options_proto_goTypes = []any{
	(*descriptorpb.MethodOptions)(nil), // 0: google.protobuf.MethodOptions
	(*descriptorpb.FieldOptions)(nil),  // 1: google.protobuf.FieldOptions
}
var file_trpc_proto_trpc_options_proto_depIdxs = []int32{
	0, // 0: trpc.alias:extendee -> google.protobuf.MethodOptions
	1, // 1: trpc.go_tag:extendee -> google.protobuf.FieldOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_trpc_proto_trpc_options_proto_init() }
func file_trpc_proto_trpc_options_proto_init() {
	if File_trpc_proto_trpc_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_trpc_proto_trpc_options_proto_rawDesc), len(file_trpc_proto_trpc_options_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_trpc_proto_trpc_options_proto_goTypes,
		DependencyIndexes: file_trpc_proto_trpc_options_proto_depIdxs,
		ExtensionInfos:    file_trpc_proto_trpc_options_proto_extTypes,
	}.Build()
	File_trpc_proto_trpc_options_proto = out.File
	file_trpc_proto_trpc_options_proto_goTypes = nil
	file_trpc_proto_trpc_options_proto_depIdxs = nil
}
