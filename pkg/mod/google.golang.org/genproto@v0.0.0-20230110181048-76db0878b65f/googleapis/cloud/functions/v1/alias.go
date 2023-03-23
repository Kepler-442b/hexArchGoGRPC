// Copyright 2022 Google LLC
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

// Code generated by aliasgen. DO NOT EDIT.

// Package functions aliases all exported identifiers in package
// "cloud.google.com/go/functions/apiv1/functionspb".
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package functions

import (
	src "cloud.google.com/go/functions/apiv1/functionspb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/functions/apiv1/functionspb
const (
	CloudFunctionStatus_ACTIVE                              = src.CloudFunctionStatus_ACTIVE
	CloudFunctionStatus_CLOUD_FUNCTION_STATUS_UNSPECIFIED   = src.CloudFunctionStatus_CLOUD_FUNCTION_STATUS_UNSPECIFIED
	CloudFunctionStatus_DELETE_IN_PROGRESS                  = src.CloudFunctionStatus_DELETE_IN_PROGRESS
	CloudFunctionStatus_DEPLOY_IN_PROGRESS                  = src.CloudFunctionStatus_DEPLOY_IN_PROGRESS
	CloudFunctionStatus_OFFLINE                             = src.CloudFunctionStatus_OFFLINE
	CloudFunctionStatus_UNKNOWN                             = src.CloudFunctionStatus_UNKNOWN
	CloudFunction_ALLOW_ALL                                 = src.CloudFunction_ALLOW_ALL
	CloudFunction_ALLOW_INTERNAL_AND_GCLB                   = src.CloudFunction_ALLOW_INTERNAL_AND_GCLB
	CloudFunction_ALLOW_INTERNAL_ONLY                       = src.CloudFunction_ALLOW_INTERNAL_ONLY
	CloudFunction_ALL_TRAFFIC                               = src.CloudFunction_ALL_TRAFFIC
	CloudFunction_ARTIFACT_REGISTRY                         = src.CloudFunction_ARTIFACT_REGISTRY
	CloudFunction_CONTAINER_REGISTRY                        = src.CloudFunction_CONTAINER_REGISTRY
	CloudFunction_DOCKER_REGISTRY_UNSPECIFIED               = src.CloudFunction_DOCKER_REGISTRY_UNSPECIFIED
	CloudFunction_INGRESS_SETTINGS_UNSPECIFIED              = src.CloudFunction_INGRESS_SETTINGS_UNSPECIFIED
	CloudFunction_PRIVATE_RANGES_ONLY                       = src.CloudFunction_PRIVATE_RANGES_ONLY
	CloudFunction_VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED = src.CloudFunction_VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED
	HttpsTrigger_SECURE_ALWAYS                              = src.HttpsTrigger_SECURE_ALWAYS
	HttpsTrigger_SECURE_OPTIONAL                            = src.HttpsTrigger_SECURE_OPTIONAL
	HttpsTrigger_SECURITY_LEVEL_UNSPECIFIED                 = src.HttpsTrigger_SECURITY_LEVEL_UNSPECIFIED
	OperationType_CREATE_FUNCTION                           = src.OperationType_CREATE_FUNCTION
	OperationType_DELETE_FUNCTION                           = src.OperationType_DELETE_FUNCTION
	OperationType_OPERATION_UNSPECIFIED                     = src.OperationType_OPERATION_UNSPECIFIED
	OperationType_UPDATE_FUNCTION                           = src.OperationType_UPDATE_FUNCTION
)

// Deprecated: Please use vars in: cloud.google.com/go/functions/apiv1/functionspb
var (
	CloudFunctionStatus_name                        = src.CloudFunctionStatus_name
	CloudFunctionStatus_value                       = src.CloudFunctionStatus_value
	CloudFunction_DockerRegistry_name               = src.CloudFunction_DockerRegistry_name
	CloudFunction_DockerRegistry_value              = src.CloudFunction_DockerRegistry_value
	CloudFunction_IngressSettings_name              = src.CloudFunction_IngressSettings_name
	CloudFunction_IngressSettings_value             = src.CloudFunction_IngressSettings_value
	CloudFunction_VpcConnectorEgressSettings_name   = src.CloudFunction_VpcConnectorEgressSettings_name
	CloudFunction_VpcConnectorEgressSettings_value  = src.CloudFunction_VpcConnectorEgressSettings_value
	File_google_cloud_functions_v1_functions_proto  = src.File_google_cloud_functions_v1_functions_proto
	File_google_cloud_functions_v1_operations_proto = src.File_google_cloud_functions_v1_operations_proto
	HttpsTrigger_SecurityLevel_name                 = src.HttpsTrigger_SecurityLevel_name
	HttpsTrigger_SecurityLevel_value                = src.HttpsTrigger_SecurityLevel_value
	OperationType_name                              = src.OperationType_name
	OperationType_value                             = src.OperationType_value
)

// Request for the `CallFunction` method.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type CallFunctionRequest = src.CallFunctionRequest

// Response of `CallFunction` method.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type CallFunctionResponse = src.CallFunctionResponse

// Describes a Cloud Function that contains user computation executed in
// response to an event. It encapsulate function and triggers configurations.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type CloudFunction = src.CloudFunction

// Describes the current stage of a deployment.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type CloudFunctionStatus = src.CloudFunctionStatus

// Docker Registry to use for storing function Docker images.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type CloudFunction_DockerRegistry = src.CloudFunction_DockerRegistry
type CloudFunction_EventTrigger = src.CloudFunction_EventTrigger
type CloudFunction_HttpsTrigger = src.CloudFunction_HttpsTrigger

// Available ingress settings. This controls what traffic can reach the
// function. If unspecified, ALLOW_ALL will be used.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type CloudFunction_IngressSettings = src.CloudFunction_IngressSettings
type CloudFunction_SourceArchiveUrl = src.CloudFunction_SourceArchiveUrl
type CloudFunction_SourceRepository = src.CloudFunction_SourceRepository
type CloudFunction_SourceUploadUrl = src.CloudFunction_SourceUploadUrl

// Available egress settings. This controls what traffic is diverted through
// the VPC Access Connector resource. By default PRIVATE_RANGES_ONLY will be
// used.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type CloudFunction_VpcConnectorEgressSettings = src.CloudFunction_VpcConnectorEgressSettings

// CloudFunctionsServiceClient is the client API for CloudFunctionsService
// service. For semantics around ctx use and closing/ending streaming RPCs,
// please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type CloudFunctionsServiceClient = src.CloudFunctionsServiceClient

// CloudFunctionsServiceServer is the server API for CloudFunctionsService
// service.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type CloudFunctionsServiceServer = src.CloudFunctionsServiceServer

// Request for the `CreateFunction` method.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type CreateFunctionRequest = src.CreateFunctionRequest

// Request for the `DeleteFunction` method.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type DeleteFunctionRequest = src.DeleteFunctionRequest

// Describes EventTrigger, used to request events be sent from another
// service.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type EventTrigger = src.EventTrigger

// Describes the policy in case of function's execution failure. If empty,
// then defaults to ignoring failures (i.e. not retrying them).
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type FailurePolicy = src.FailurePolicy

// Describes the retry policy in case of function's execution failure. A
// function execution will be retried on any failure. A failed execution will
// be retried up to 7 days with an exponential backoff (capped at 10 seconds).
// Retried execution is charged as any other execution.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type FailurePolicy_Retry = src.FailurePolicy_Retry
type FailurePolicy_Retry_ = src.FailurePolicy_Retry_

// Request of `GenerateDownloadUrl` method.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type GenerateDownloadUrlRequest = src.GenerateDownloadUrlRequest

// Response of `GenerateDownloadUrl` method.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type GenerateDownloadUrlResponse = src.GenerateDownloadUrlResponse

// Request of `GenerateSourceUploadUrl` method.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type GenerateUploadUrlRequest = src.GenerateUploadUrlRequest

// Response of `GenerateSourceUploadUrl` method.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type GenerateUploadUrlResponse = src.GenerateUploadUrlResponse

// Request for the `GetFunction` method.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type GetFunctionRequest = src.GetFunctionRequest

// Describes HttpsTrigger, could be used to connect web hooks to function.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type HttpsTrigger = src.HttpsTrigger

// Available security level settings. This controls the methods to enforce
// security (HTTPS) on a URL. If unspecified, SECURE_OPTIONAL will be used.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type HttpsTrigger_SecurityLevel = src.HttpsTrigger_SecurityLevel

// Request for the `ListFunctions` method.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type ListFunctionsRequest = src.ListFunctionsRequest

// Response for the `ListFunctions` method.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type ListFunctionsResponse = src.ListFunctionsResponse

// Metadata describing an [Operation][google.longrunning.Operation]
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type OperationMetadataV1 = src.OperationMetadataV1

// A type of an operation.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type OperationType = src.OperationType

// Configuration for a secret environment variable. It has the information
// necessary to fetch the secret value from secret manager and expose it as an
// environment variable.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type SecretEnvVar = src.SecretEnvVar

// Configuration for a secret volume. It has the information necessary to
// fetch the secret value from secret manager and make it available as files
// mounted at the requested paths within the application container. Secret
// value is not a part of the configuration. Every filesystem read operation
// performs a lookup in secret manager to retrieve the secret value.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type SecretVolume = src.SecretVolume

// Configuration for a single version.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type SecretVolume_SecretVersion = src.SecretVolume_SecretVersion

// Describes SourceRepository, used to represent parameters related to source
// repository where a function is hosted.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type SourceRepository = src.SourceRepository

// UnimplementedCloudFunctionsServiceServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type UnimplementedCloudFunctionsServiceServer = src.UnimplementedCloudFunctionsServiceServer

// Request for the `UpdateFunction` method.
//
// Deprecated: Please use types in: cloud.google.com/go/functions/apiv1/functionspb
type UpdateFunctionRequest = src.UpdateFunctionRequest

// Deprecated: Please use funcs in: cloud.google.com/go/functions/apiv1/functionspb
func NewCloudFunctionsServiceClient(cc grpc.ClientConnInterface) CloudFunctionsServiceClient {
	return src.NewCloudFunctionsServiceClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/functions/apiv1/functionspb
func RegisterCloudFunctionsServiceServer(s *grpc.Server, srv CloudFunctionsServiceServer) {
	src.RegisterCloudFunctionsServiceServer(s, srv)
}