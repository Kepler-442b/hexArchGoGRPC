// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: google/logging/type/log_severity.proto

package ltype

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// The severity of the event described in a log entry, expressed as one of the
// standard severity levels listed below.  For your reference, the levels are
// assigned the listed numeric values. The effect of using numeric values other
// than those listed is undefined.
//
// You can filter for log entries by severity.  For example, the following
// filter expression will match log entries with severities `INFO`, `NOTICE`,
// and `WARNING`:
//
//     severity > DEBUG AND severity <= WARNING
//
// If you are writing log entries, you should map other severity encodings to
// one of these standard levels. For example, you might map all of Java's FINE,
// FINER, and FINEST levels to `LogSeverity.DEBUG`. You can preserve the
// original severity level in the log entry payload if you wish.
type LogSeverity int32

const (
	// (0) The log entry has no assigned severity level.
	LogSeverity_DEFAULT LogSeverity = 0
	// (100) Debug or trace information.
	LogSeverity_DEBUG LogSeverity = 100
	// (200) Routine information, such as ongoing status or performance.
	LogSeverity_INFO LogSeverity = 200
	// (300) Normal but significant events, such as start up, shut down, or
	// a configuration change.
	LogSeverity_NOTICE LogSeverity = 300
	// (400) Warning events might cause problems.
	LogSeverity_WARNING LogSeverity = 400
	// (500) Error events are likely to cause problems.
	LogSeverity_ERROR LogSeverity = 500
	// (600) Critical events cause more severe problems or outages.
	LogSeverity_CRITICAL LogSeverity = 600
	// (700) A person must take an action immediately.
	LogSeverity_ALERT LogSeverity = 700
	// (800) One or more systems are unusable.
	LogSeverity_EMERGENCY LogSeverity = 800
)

// Enum value maps for LogSeverity.
var (
	LogSeverity_name = map[int32]string{
		0:   "DEFAULT",
		100: "DEBUG",
		200: "INFO",
		300: "NOTICE",
		400: "WARNING",
		500: "ERROR",
		600: "CRITICAL",
		700: "ALERT",
		800: "EMERGENCY",
	}
	LogSeverity_value = map[string]int32{
		"DEFAULT":   0,
		"DEBUG":     100,
		"INFO":      200,
		"NOTICE":    300,
		"WARNING":   400,
		"ERROR":     500,
		"CRITICAL":  600,
		"ALERT":     700,
		"EMERGENCY": 800,
	}
)

func (x LogSeverity) Enum() *LogSeverity {
	p := new(LogSeverity)
	*p = x
	return p
}

func (x LogSeverity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogSeverity) Descriptor() protoreflect.EnumDescriptor {
	return file_google_logging_type_log_severity_proto_enumTypes[0].Descriptor()
}

func (LogSeverity) Type() protoreflect.EnumType {
	return &file_google_logging_type_log_severity_proto_enumTypes[0]
}

func (x LogSeverity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogSeverity.Descriptor instead.
func (LogSeverity) EnumDescriptor() ([]byte, []int) {
	return file_google_logging_type_log_severity_proto_rawDescGZIP(), []int{0}
}

var File_google_logging_type_log_severity_proto protoreflect.FileDescriptor

var file_google_logging_type_log_severity_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x82, 0x01, 0x0a, 0x0b,
	0x4c, 0x6f, 0x67, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x44,
	0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x55,
	0x47, 0x10, 0x64, 0x12, 0x09, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0xc8, 0x01, 0x12, 0x0b,
	0x0a, 0x06, 0x4e, 0x4f, 0x54, 0x49, 0x43, 0x45, 0x10, 0xac, 0x02, 0x12, 0x0c, 0x0a, 0x07, 0x57,
	0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x90, 0x03, 0x12, 0x0a, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0xf4, 0x03, 0x12, 0x0d, 0x0a, 0x08, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41,
	0x4c, 0x10, 0xd8, 0x04, 0x12, 0x0a, 0x0a, 0x05, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x10, 0xbc, 0x05,
	0x12, 0x0e, 0x0a, 0x09, 0x45, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x59, 0x10, 0xa0, 0x06,
	0x42, 0x9f, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x42, 0x10, 0x4c, 0x6f,
	0x67, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x3b, 0x6c, 0x74, 0x79, 0x70, 0x65, 0xaa, 0x02, 0x19, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x2e, 0x54, 0x79, 0x70, 0x65, 0xca, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x54, 0x79,
	0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_logging_type_log_severity_proto_rawDescOnce sync.Once
	file_google_logging_type_log_severity_proto_rawDescData = file_google_logging_type_log_severity_proto_rawDesc
)

func file_google_logging_type_log_severity_proto_rawDescGZIP() []byte {
	file_google_logging_type_log_severity_proto_rawDescOnce.Do(func() {
		file_google_logging_type_log_severity_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_logging_type_log_severity_proto_rawDescData)
	})
	return file_google_logging_type_log_severity_proto_rawDescData
}

var file_google_logging_type_log_severity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_logging_type_log_severity_proto_goTypes = []interface{}{
	(LogSeverity)(0), // 0: google.logging.type.LogSeverity
}
var file_google_logging_type_log_severity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_logging_type_log_severity_proto_init() }
func file_google_logging_type_log_severity_proto_init() {
	if File_google_logging_type_log_severity_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_logging_type_log_severity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_logging_type_log_severity_proto_goTypes,
		DependencyIndexes: file_google_logging_type_log_severity_proto_depIdxs,
		EnumInfos:         file_google_logging_type_log_severity_proto_enumTypes,
	}.Build()
	File_google_logging_type_log_severity_proto = out.File
	file_google_logging_type_log_severity_proto_rawDesc = nil
	file_google_logging_type_log_severity_proto_goTypes = nil
	file_google_logging_type_log_severity_proto_depIdxs = nil
}
