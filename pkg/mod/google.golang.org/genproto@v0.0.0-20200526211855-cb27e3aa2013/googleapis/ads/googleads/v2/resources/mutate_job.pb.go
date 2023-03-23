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
// source: google/ads/googleads/v2/resources/mutate_job.proto

package resources

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// A list of mutates being processed asynchronously. The mutates are uploaded
// by the user. The mutates themselves aren't readable and the results of the
// job can only be read using MutateJobService.ListMutateJobResults.
type MutateJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the mutate job.
	// Mutate job resource names have the form:
	//
	// `customers/{customer_id}/mutateJobs/{mutate_job_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. ID of this mutate job.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. The next sequence token to use when adding operations. Only set when the
	// mutate job status is PENDING.
	NextAddSequenceToken *wrappers.StringValue `protobuf:"bytes,3,opt,name=next_add_sequence_token,json=nextAddSequenceToken,proto3" json:"next_add_sequence_token,omitempty"`
	// Output only. Contains additional information about this mutate job.
	Metadata *MutateJob_MutateJobMetadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Output only. Status of this mutate job.
	Status enums.MutateJobStatusEnum_MutateJobStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.MutateJobStatusEnum_MutateJobStatus" json:"status,omitempty"`
	// Output only. The resource name of the long-running operation that can be used to poll
	// for completion. Only set when the mutate job status is RUNNING or DONE.
	LongRunningOperation *wrappers.StringValue `protobuf:"bytes,6,opt,name=long_running_operation,json=longRunningOperation,proto3" json:"long_running_operation,omitempty"`
}

func (x *MutateJob) Reset() {
	*x = MutateJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_resources_mutate_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateJob) ProtoMessage() {}

func (x *MutateJob) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_resources_mutate_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateJob.ProtoReflect.Descriptor instead.
func (*MutateJob) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_resources_mutate_job_proto_rawDescGZIP(), []int{0}
}

func (x *MutateJob) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateJob) GetId() *wrappers.Int64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *MutateJob) GetNextAddSequenceToken() *wrappers.StringValue {
	if x != nil {
		return x.NextAddSequenceToken
	}
	return nil
}

func (x *MutateJob) GetMetadata() *MutateJob_MutateJobMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MutateJob) GetStatus() enums.MutateJobStatusEnum_MutateJobStatus {
	if x != nil {
		return x.Status
	}
	return enums.MutateJobStatusEnum_UNSPECIFIED
}

func (x *MutateJob) GetLongRunningOperation() *wrappers.StringValue {
	if x != nil {
		return x.LongRunningOperation
	}
	return nil
}

// Additional information about the mutate job. This message is also used as
// metadata returned in mutate job Long Running Operations.
type MutateJob_MutateJobMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The time when this mutate job was created.
	// Formatted as yyyy-mm-dd hh:mm:ss. Example: "2018-03-05 09:15:00"
	CreationDateTime *wrappers.StringValue `protobuf:"bytes,1,opt,name=creation_date_time,json=creationDateTime,proto3" json:"creation_date_time,omitempty"`
	// Output only. The time when this mutate job was completed.
	// Formatted as yyyy-MM-dd HH:mm:ss. Example: "2018-03-05 09:16:00"
	CompletionDateTime *wrappers.StringValue `protobuf:"bytes,2,opt,name=completion_date_time,json=completionDateTime,proto3" json:"completion_date_time,omitempty"`
	// Output only. The fraction (between 0.0 and 1.0) of mutates that have been processed.
	// This is empty if the job hasn't started running yet.
	EstimatedCompletionRatio *wrappers.DoubleValue `protobuf:"bytes,3,opt,name=estimated_completion_ratio,json=estimatedCompletionRatio,proto3" json:"estimated_completion_ratio,omitempty"`
	// Output only. The number of mutate operations in the mutate job.
	OperationCount *wrappers.Int64Value `protobuf:"bytes,4,opt,name=operation_count,json=operationCount,proto3" json:"operation_count,omitempty"`
	// Output only. The number of mutate operations executed by the mutate job.
	// Present only if the job has started running.
	ExecutedOperationCount *wrappers.Int64Value `protobuf:"bytes,5,opt,name=executed_operation_count,json=executedOperationCount,proto3" json:"executed_operation_count,omitempty"`
}

func (x *MutateJob_MutateJobMetadata) Reset() {
	*x = MutateJob_MutateJobMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_resources_mutate_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateJob_MutateJobMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateJob_MutateJobMetadata) ProtoMessage() {}

func (x *MutateJob_MutateJobMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_resources_mutate_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateJob_MutateJobMetadata.ProtoReflect.Descriptor instead.
func (*MutateJob_MutateJobMetadata) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_resources_mutate_job_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MutateJob_MutateJobMetadata) GetCreationDateTime() *wrappers.StringValue {
	if x != nil {
		return x.CreationDateTime
	}
	return nil
}

func (x *MutateJob_MutateJobMetadata) GetCompletionDateTime() *wrappers.StringValue {
	if x != nil {
		return x.CompletionDateTime
	}
	return nil
}

func (x *MutateJob_MutateJobMetadata) GetEstimatedCompletionRatio() *wrappers.DoubleValue {
	if x != nil {
		return x.EstimatedCompletionRatio
	}
	return nil
}

func (x *MutateJob_MutateJobMetadata) GetOperationCount() *wrappers.Int64Value {
	if x != nil {
		return x.OperationCount
	}
	return nil
}

func (x *MutateJob_MutateJobMetadata) GetExecutedOperationCount() *wrappers.Int64Value {
	if x != nil {
		return x.ExecutedOperationCount
	}
	return nil
}

var File_google_ads_googleads_v2_resources_mutate_job_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v2_resources_mutate_job_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6a, 0x6f,
	0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x08, 0x0a, 0x09, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x4f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xe0,
	0x41, 0x05, 0xfa, 0x41, 0x24, 0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x58, 0x0a, 0x17, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x14, 0x6e,
	0x65, 0x78, 0x74, 0x41, 0x64, 0x64, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x5f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x4a, 0x6f, 0x62, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x5f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4a,
	0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x57, 0x0a, 0x16, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x72, 0x75,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x14, 0x6c, 0x6f, 0x6e, 0x67, 0x52, 0x75,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0xc1,
	0x03, 0x0a, 0x11, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x4f, 0x0a, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x53, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5f, 0x0a, 0x1a, 0x65, 0x73,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x18, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x49, 0x0a, 0x0f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x5a, 0x0a, 0x18, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x64, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x16, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x3a, 0x55, 0xea, 0x41, 0x52, 0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x2c, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x7d, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x2f, 0x7b, 0x6d, 0x75,
	0x74, 0x61, 0x74, 0x65, 0x5f, 0x6a, 0x6f, 0x62, 0x7d, 0x42, 0xfb, 0x01, 0x0a, 0x25, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x42, 0x0e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x32, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x32, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea,
	0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x3a, 0x3a, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v2_resources_mutate_job_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v2_resources_mutate_job_proto_rawDescData = file_google_ads_googleads_v2_resources_mutate_job_proto_rawDesc
)

func file_google_ads_googleads_v2_resources_mutate_job_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v2_resources_mutate_job_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v2_resources_mutate_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v2_resources_mutate_job_proto_rawDescData)
	})
	return file_google_ads_googleads_v2_resources_mutate_job_proto_rawDescData
}

var file_google_ads_googleads_v2_resources_mutate_job_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v2_resources_mutate_job_proto_goTypes = []interface{}{
	(*MutateJob)(nil),                              // 0: google.ads.googleads.v2.resources.MutateJob
	(*MutateJob_MutateJobMetadata)(nil),            // 1: google.ads.googleads.v2.resources.MutateJob.MutateJobMetadata
	(*wrappers.Int64Value)(nil),                    // 2: google.protobuf.Int64Value
	(*wrappers.StringValue)(nil),                   // 3: google.protobuf.StringValue
	(enums.MutateJobStatusEnum_MutateJobStatus)(0), // 4: google.ads.googleads.v2.enums.MutateJobStatusEnum.MutateJobStatus
	(*wrappers.DoubleValue)(nil),                   // 5: google.protobuf.DoubleValue
}
var file_google_ads_googleads_v2_resources_mutate_job_proto_depIdxs = []int32{
	2,  // 0: google.ads.googleads.v2.resources.MutateJob.id:type_name -> google.protobuf.Int64Value
	3,  // 1: google.ads.googleads.v2.resources.MutateJob.next_add_sequence_token:type_name -> google.protobuf.StringValue
	1,  // 2: google.ads.googleads.v2.resources.MutateJob.metadata:type_name -> google.ads.googleads.v2.resources.MutateJob.MutateJobMetadata
	4,  // 3: google.ads.googleads.v2.resources.MutateJob.status:type_name -> google.ads.googleads.v2.enums.MutateJobStatusEnum.MutateJobStatus
	3,  // 4: google.ads.googleads.v2.resources.MutateJob.long_running_operation:type_name -> google.protobuf.StringValue
	3,  // 5: google.ads.googleads.v2.resources.MutateJob.MutateJobMetadata.creation_date_time:type_name -> google.protobuf.StringValue
	3,  // 6: google.ads.googleads.v2.resources.MutateJob.MutateJobMetadata.completion_date_time:type_name -> google.protobuf.StringValue
	5,  // 7: google.ads.googleads.v2.resources.MutateJob.MutateJobMetadata.estimated_completion_ratio:type_name -> google.protobuf.DoubleValue
	2,  // 8: google.ads.googleads.v2.resources.MutateJob.MutateJobMetadata.operation_count:type_name -> google.protobuf.Int64Value
	2,  // 9: google.ads.googleads.v2.resources.MutateJob.MutateJobMetadata.executed_operation_count:type_name -> google.protobuf.Int64Value
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v2_resources_mutate_job_proto_init() }
func file_google_ads_googleads_v2_resources_mutate_job_proto_init() {
	if File_google_ads_googleads_v2_resources_mutate_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v2_resources_mutate_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateJob); i {
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
		file_google_ads_googleads_v2_resources_mutate_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateJob_MutateJobMetadata); i {
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
			RawDescriptor: file_google_ads_googleads_v2_resources_mutate_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v2_resources_mutate_job_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v2_resources_mutate_job_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v2_resources_mutate_job_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v2_resources_mutate_job_proto = out.File
	file_google_ads_googleads_v2_resources_mutate_job_proto_rawDesc = nil
	file_google_ads_googleads_v2_resources_mutate_job_proto_goTypes = nil
	file_google_ads_googleads_v2_resources_mutate_job_proto_depIdxs = nil
}
