// Copyright 2014 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: bq_table.proto

package protos

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type BigQueryMessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies a name of table in BigQuery for the message.
	//
	// If not blank, indicates the message is a type of record to be stored into BigQuery.
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	// If true, BigQuery field names will default to a field's JSON name,
	// not its original/proto field name.
	UseJsonNames bool `protobuf:"varint,2,opt,name=use_json_names,json=useJsonNames,proto3" json:"use_json_names,omitempty"`
	// If set, adds defined extra fields to a JSON representation of the message.
	// Value format: "<field name>:<BigQuery field type>" for basic types
	// or "<field name>:RECORD:<protobuf type>" for message types.
	// "NULLABLE" by default, different mode may be set via optional suffix ":<mode>"
	ExtraFields []string `protobuf:"bytes,3,rep,name=extra_fields,json=extraFields,proto3" json:"extra_fields,omitempty"`
}

func (x *BigQueryMessageOptions) Reset() {
	*x = BigQueryMessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bq_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigQueryMessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigQueryMessageOptions) ProtoMessage() {}

func (x *BigQueryMessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_bq_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigQueryMessageOptions.ProtoReflect.Descriptor instead.
func (*BigQueryMessageOptions) Descriptor() ([]byte, []int) {
	return file_bq_table_proto_rawDescGZIP(), []int{0}
}

func (x *BigQueryMessageOptions) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *BigQueryMessageOptions) GetUseJsonNames() bool {
	if x != nil {
		return x.UseJsonNames
	}
	return false
}

func (x *BigQueryMessageOptions) GetExtraFields() []string {
	if x != nil {
		return x.ExtraFields
	}
	return nil
}

var file_bq_table_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*BigQueryMessageOptions)(nil),
		Field:         1021,
		Name:          "bq_schema.options",
		Tag:           "bytes,1021,opt,name=options",
		Filename:      "bq_table.proto",
	},
}

// Extension fields to descriptor.MessageOptions.
var (
	// BigQuery message schema generation options.
	//
	// The field number is a globally unique id for this option, assigned by
	// protobuf-global-extension-registry@google.com
	//
	// optional bq_schema.fieldMessageOptions options = 1021;
	E_BigqueryOpts = &file_bq_table_proto_extTypes[0]
)

var File_bq_table_proto protoreflect.FileDescriptor

var file_bq_table_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x71, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x67, 0x65, 0x6e, 0x5f, 0x62, 0x71, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x80, 0x01, 0x0a, 0x16, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x75,
	0x73, 0x65, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x72, 0x61, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x3a, 0x6c, 0x0a, 0x0d, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5f, 0x6f, 0x70, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfd, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x67, 0x65, 0x6e, 0x5f, 0x62, 0x71, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x42, 0x69,
	0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70,
	0x74, 0x73, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bq_table_proto_rawDescOnce sync.Once
	file_bq_table_proto_rawDescData = file_bq_table_proto_rawDesc
)

func file_bq_table_proto_rawDescGZIP() []byte {
	file_bq_table_proto_rawDescOnce.Do(func() {
		file_bq_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_bq_table_proto_rawDescData)
	})
	return file_bq_table_proto_rawDescData
}

var file_bq_table_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bq_table_proto_goTypes = []interface{}{
	(*BigQueryMessageOptions)(nil),    // 0: bq_schema.fieldMessageOptions
	(*descriptor.MessageOptions)(nil), // 1: google.protobuf.MessageOptions
}
var file_bq_table_proto_depIdxs = []int32{
	1, // 0: bq_schema.options:extendee -> google.protobuf.MessageOptions
	0, // 1: bq_schema.options:type_name -> bq_schema.fieldMessageOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bq_table_proto_init() }
func file_bq_table_proto_init() {
	if File_bq_table_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bq_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigQueryMessageOptions); i {
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
			RawDescriptor: file_bq_table_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_bq_table_proto_goTypes,
		DependencyIndexes: file_bq_table_proto_depIdxs,
		MessageInfos:      file_bq_table_proto_msgTypes,
		ExtensionInfos:    file_bq_table_proto_extTypes,
	}.Build()
	File_bq_table_proto = out.File
	file_bq_table_proto_rawDesc = nil
	file_bq_table_proto_goTypes = nil
	file_bq_table_proto_depIdxs = nil
}
