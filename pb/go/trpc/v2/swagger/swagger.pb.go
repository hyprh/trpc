// Tencent is pleased to support the open source community by making tRPC available.
// Copyright (C) 2023 THL A29 Limited, a Tencent company. All rights reserved.
// If you have downloaded a copy of the tRPC source code from Tencent,
// please note that tRPC source code is licensed under the Apache 2.0 License,
// A copy of the Apache 2.0 License is included in this file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.27.1
// source: trpc/v2/swagger/swagger.proto

package swagger

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// to gen swagger json
type SwaggerRule struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,50103,opt,name=title,proto3" json:"title,omitempty"`
	Method        string                 `protobuf:"bytes,50104,opt,name=method,proto3" json:"method,omitempty"`
	Description   string                 `protobuf:"bytes,50105,opt,name=description,proto3" json:"description,omitempty"`
	Params        []*SwaggerParam        `protobuf:"bytes,50106,rep,name=params,proto3" json:"params,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SwaggerRule) Reset() {
	*x = SwaggerRule{}
	mi := &file_trpc_v2_swagger_swagger_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SwaggerRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwaggerRule) ProtoMessage() {}

func (x *SwaggerRule) ProtoReflect() protoreflect.Message {
	mi := &file_trpc_v2_swagger_swagger_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwaggerRule.ProtoReflect.Descriptor instead.
func (*SwaggerRule) Descriptor() ([]byte, []int) {
	return file_trpc_v2_swagger_swagger_proto_rawDescGZIP(), []int{0}
}

func (x *SwaggerRule) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SwaggerRule) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *SwaggerRule) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SwaggerRule) GetParams() []*SwaggerParam {
	if x != nil {
		return x.Params
	}
	return nil
}

type SwaggerParam struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Required      bool                   `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
	Default       string                 `protobuf:"bytes,3,opt,name=default,proto3" json:"default,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SwaggerParam) Reset() {
	*x = SwaggerParam{}
	mi := &file_trpc_v2_swagger_swagger_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SwaggerParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwaggerParam) ProtoMessage() {}

func (x *SwaggerParam) ProtoReflect() protoreflect.Message {
	mi := &file_trpc_v2_swagger_swagger_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwaggerParam.ProtoReflect.Descriptor instead.
func (*SwaggerParam) Descriptor() ([]byte, []int) {
	return file_trpc_v2_swagger_swagger_proto_rawDescGZIP(), []int{1}
}

func (x *SwaggerParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SwaggerParam) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *SwaggerParam) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

var file_trpc_v2_swagger_swagger_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*SwaggerRule)(nil),
		Field:         1156,
		Name:          "trpc.v2.swagger",
		Tag:           "bytes,1156,opt,name=swagger",
		Filename:      "trpc/v2/swagger/swagger.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional trpc.v2.SwaggerRule swagger = 1156;
	E_Swagger = &file_trpc_v2_swagger_swagger_proto_extTypes[0]
)

var File_trpc_v2_swagger_swagger_proto protoreflect.FileDescriptor

const file_trpc_v2_swagger_swagger_proto_rawDesc = "" +
	"\n" +
	"\x1dtrpc/v2/swagger/swagger.proto\x12\atrpc.v2\x1a google/protobuf/descriptor.proto\"\x94\x01\n" +
	"\vSwaggerRule\x12\x16\n" +
	"\x05title\x18\xb7\x87\x03 \x01(\tR\x05title\x12\x18\n" +
	"\x06method\x18\xb8\x87\x03 \x01(\tR\x06method\x12\"\n" +
	"\vdescription\x18\xb9\x87\x03 \x01(\tR\vdescription\x12/\n" +
	"\x06params\x18\xba\x87\x03 \x03(\v2\x15.trpc.v2.SwaggerParamR\x06params\"X\n" +
	"\fSwaggerParam\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1a\n" +
	"\brequired\x18\x02 \x01(\bR\brequired\x12\x18\n" +
	"\adefault\x18\x03 \x01(\tR\adefault:O\n" +
	"\aswagger\x12\x1e.google.protobuf.MethodOptions\x18\x84\t \x01(\v2\x14.trpc.v2.SwaggerRuleR\aswaggerBT\n" +
	"\x1dcom.tencent.trpc.protobuf.extZ3trpc.group/trpc/trpc-protocol/pb/go/trpc/v2/swaggerb\x06proto3"

var (
	file_trpc_v2_swagger_swagger_proto_rawDescOnce sync.Once
	file_trpc_v2_swagger_swagger_proto_rawDescData []byte
)

func file_trpc_v2_swagger_swagger_proto_rawDescGZIP() []byte {
	file_trpc_v2_swagger_swagger_proto_rawDescOnce.Do(func() {
		file_trpc_v2_swagger_swagger_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_trpc_v2_swagger_swagger_proto_rawDesc), len(file_trpc_v2_swagger_swagger_proto_rawDesc)))
	})
	return file_trpc_v2_swagger_swagger_proto_rawDescData
}

var file_trpc_v2_swagger_swagger_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_trpc_v2_swagger_swagger_proto_goTypes = []any{
	(*SwaggerRule)(nil),                // 0: trpc.v2.SwaggerRule
	(*SwaggerParam)(nil),               // 1: trpc.v2.SwaggerParam
	(*descriptorpb.MethodOptions)(nil), // 2: google.protobuf.MethodOptions
}
var file_trpc_v2_swagger_swagger_proto_depIdxs = []int32{
	1, // 0: trpc.v2.SwaggerRule.params:type_name -> trpc.v2.SwaggerParam
	2, // 1: trpc.v2.swagger:extendee -> google.protobuf.MethodOptions
	0, // 2: trpc.v2.swagger:type_name -> trpc.v2.SwaggerRule
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_trpc_v2_swagger_swagger_proto_init() }
func file_trpc_v2_swagger_swagger_proto_init() {
	if File_trpc_v2_swagger_swagger_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_trpc_v2_swagger_swagger_proto_rawDesc), len(file_trpc_v2_swagger_swagger_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_trpc_v2_swagger_swagger_proto_goTypes,
		DependencyIndexes: file_trpc_v2_swagger_swagger_proto_depIdxs,
		MessageInfos:      file_trpc_v2_swagger_swagger_proto_msgTypes,
		ExtensionInfos:    file_trpc_v2_swagger_swagger_proto_extTypes,
	}.Build()
	File_trpc_v2_swagger_swagger_proto = out.File
	file_trpc_v2_swagger_swagger_proto_goTypes = nil
	file_trpc_v2_swagger_swagger_proto_depIdxs = nil
}
