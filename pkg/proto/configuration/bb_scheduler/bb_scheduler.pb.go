// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.23.1
// source: pkg/proto/configuration/bb_scheduler/bb_scheduler.proto

package bb_scheduler

import (
	v2 "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	scheduler "github.com/buildbarn/bb-remote-execution/pkg/proto/configuration/scheduler"
	auth "github.com/buildbarn/bb-storage/pkg/proto/configuration/auth"
	blobstore "github.com/buildbarn/bb-storage/pkg/proto/configuration/blobstore"
	global "github.com/buildbarn/bb-storage/pkg/proto/configuration/global"
	grpc "github.com/buildbarn/bb-storage/pkg/proto/configuration/grpc"
	http "github.com/buildbarn/bb-storage/pkg/proto/configuration/http"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ApplicationConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminHttpServers                  []*http.ServerConfiguration              `protobuf:"bytes,19,rep,name=admin_http_servers,json=adminHttpServers,proto3" json:"admin_http_servers,omitempty"`
	AdminRoutePrefix                  string                                   `protobuf:"bytes,22,opt,name=admin_route_prefix,json=adminRoutePrefix,proto3" json:"admin_route_prefix,omitempty"`
	ClientGrpcServers                 []*grpc.ServerConfiguration              `protobuf:"bytes,3,rep,name=client_grpc_servers,json=clientGrpcServers,proto3" json:"client_grpc_servers,omitempty"`
	WorkerGrpcServers                 []*grpc.ServerConfiguration              `protobuf:"bytes,4,rep,name=worker_grpc_servers,json=workerGrpcServers,proto3" json:"worker_grpc_servers,omitempty"`
	BrowserUrl                        string                                   `protobuf:"bytes,5,opt,name=browser_url,json=browserUrl,proto3" json:"browser_url,omitempty"`
	ContentAddressableStorage         *blobstore.BlobAccessConfiguration       `protobuf:"bytes,6,opt,name=content_addressable_storage,json=contentAddressableStorage,proto3" json:"content_addressable_storage,omitempty"`
	MaximumMessageSizeBytes           int64                                    `protobuf:"varint,7,opt,name=maximum_message_size_bytes,json=maximumMessageSizeBytes,proto3" json:"maximum_message_size_bytes,omitempty"`
	Global                            *global.Configuration                    `protobuf:"bytes,8,opt,name=global,proto3" json:"global,omitempty"`
	BuildQueueStateGrpcServers        []*grpc.ServerConfiguration              `protobuf:"bytes,11,rep,name=build_queue_state_grpc_servers,json=buildQueueStateGrpcServers,proto3" json:"build_queue_state_grpc_servers,omitempty"`
	PredeclaredPlatformQueues         []*PredeclaredPlatformQueueConfiguration `protobuf:"bytes,12,rep,name=predeclared_platform_queues,json=predeclaredPlatformQueues,proto3" json:"predeclared_platform_queues,omitempty"`
	ExecuteAuthorizer                 *auth.AuthorizerConfiguration            `protobuf:"bytes,15,opt,name=execute_authorizer,json=executeAuthorizer,proto3" json:"execute_authorizer,omitempty"`
	ModifyDrainsAuthorizer            *auth.AuthorizerConfiguration            `protobuf:"bytes,20,opt,name=modify_drains_authorizer,json=modifyDrainsAuthorizer,proto3" json:"modify_drains_authorizer,omitempty"`
	KillOperationsAuthorizer          *auth.AuthorizerConfiguration            `protobuf:"bytes,21,opt,name=kill_operations_authorizer,json=killOperationsAuthorizer,proto3" json:"kill_operations_authorizer,omitempty"`
	ActionRouter                      *scheduler.ActionRouterConfiguration     `protobuf:"bytes,16,opt,name=action_router,json=actionRouter,proto3" json:"action_router,omitempty"`
	InitialSizeClassCache             *blobstore.BlobAccessConfiguration       `protobuf:"bytes,17,opt,name=initial_size_class_cache,json=initialSizeClassCache,proto3" json:"initial_size_class_cache,omitempty"`
	PlatformQueueWithNoWorkersTimeout *durationpb.Duration                     `protobuf:"bytes,18,opt,name=platform_queue_with_no_workers_timeout,json=platformQueueWithNoWorkersTimeout,proto3" json:"platform_queue_with_no_workers_timeout,omitempty"`
}

func (x *ApplicationConfiguration) Reset() {
	*x = ApplicationConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationConfiguration) ProtoMessage() {}

func (x *ApplicationConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationConfiguration.ProtoReflect.Descriptor instead.
func (*ApplicationConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_rawDescGZIP(), []int{0}
}

func (x *ApplicationConfiguration) GetAdminHttpServers() []*http.ServerConfiguration {
	if x != nil {
		return x.AdminHttpServers
	}
	return nil
}

func (x *ApplicationConfiguration) GetAdminRoutePrefix() string {
	if x != nil {
		return x.AdminRoutePrefix
	}
	return ""
}

func (x *ApplicationConfiguration) GetClientGrpcServers() []*grpc.ServerConfiguration {
	if x != nil {
		return x.ClientGrpcServers
	}
	return nil
}

func (x *ApplicationConfiguration) GetWorkerGrpcServers() []*grpc.ServerConfiguration {
	if x != nil {
		return x.WorkerGrpcServers
	}
	return nil
}

func (x *ApplicationConfiguration) GetBrowserUrl() string {
	if x != nil {
		return x.BrowserUrl
	}
	return ""
}

func (x *ApplicationConfiguration) GetContentAddressableStorage() *blobstore.BlobAccessConfiguration {
	if x != nil {
		return x.ContentAddressableStorage
	}
	return nil
}

func (x *ApplicationConfiguration) GetMaximumMessageSizeBytes() int64 {
	if x != nil {
		return x.MaximumMessageSizeBytes
	}
	return 0
}

func (x *ApplicationConfiguration) GetGlobal() *global.Configuration {
	if x != nil {
		return x.Global
	}
	return nil
}

func (x *ApplicationConfiguration) GetBuildQueueStateGrpcServers() []*grpc.ServerConfiguration {
	if x != nil {
		return x.BuildQueueStateGrpcServers
	}
	return nil
}

func (x *ApplicationConfiguration) GetPredeclaredPlatformQueues() []*PredeclaredPlatformQueueConfiguration {
	if x != nil {
		return x.PredeclaredPlatformQueues
	}
	return nil
}

func (x *ApplicationConfiguration) GetExecuteAuthorizer() *auth.AuthorizerConfiguration {
	if x != nil {
		return x.ExecuteAuthorizer
	}
	return nil
}

func (x *ApplicationConfiguration) GetModifyDrainsAuthorizer() *auth.AuthorizerConfiguration {
	if x != nil {
		return x.ModifyDrainsAuthorizer
	}
	return nil
}

func (x *ApplicationConfiguration) GetKillOperationsAuthorizer() *auth.AuthorizerConfiguration {
	if x != nil {
		return x.KillOperationsAuthorizer
	}
	return nil
}

func (x *ApplicationConfiguration) GetActionRouter() *scheduler.ActionRouterConfiguration {
	if x != nil {
		return x.ActionRouter
	}
	return nil
}

func (x *ApplicationConfiguration) GetInitialSizeClassCache() *blobstore.BlobAccessConfiguration {
	if x != nil {
		return x.InitialSizeClassCache
	}
	return nil
}

func (x *ApplicationConfiguration) GetPlatformQueueWithNoWorkersTimeout() *durationpb.Duration {
	if x != nil {
		return x.PlatformQueueWithNoWorkersTimeout
	}
	return nil
}

type PredeclaredPlatformQueueConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceNamePrefix                        string                 `protobuf:"bytes,1,opt,name=instance_name_prefix,json=instanceNamePrefix,proto3" json:"instance_name_prefix,omitempty"`
	Platform                                  *v2.Platform           `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`
	MaximumSizeClass                          uint32                 `protobuf:"varint,3,opt,name=maximum_size_class,json=maximumSizeClass,proto3" json:"maximum_size_class,omitempty"`
	WorkerInvocationStickinessLimits          []*durationpb.Duration `protobuf:"bytes,5,rep,name=worker_invocation_stickiness_limits,json=workerInvocationStickinessLimits,proto3" json:"worker_invocation_stickiness_limits,omitempty"`
	MaximumQueuedBackgroundLearningOperations int32                  `protobuf:"varint,6,opt,name=maximum_queued_background_learning_operations,json=maximumQueuedBackgroundLearningOperations,proto3" json:"maximum_queued_background_learning_operations,omitempty"`
	BackgroundLearningOperationPriority       int32                  `protobuf:"varint,7,opt,name=background_learning_operation_priority,json=backgroundLearningOperationPriority,proto3" json:"background_learning_operation_priority,omitempty"`
}

func (x *PredeclaredPlatformQueueConfiguration) Reset() {
	*x = PredeclaredPlatformQueueConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredeclaredPlatformQueueConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredeclaredPlatformQueueConfiguration) ProtoMessage() {}

func (x *PredeclaredPlatformQueueConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredeclaredPlatformQueueConfiguration.ProtoReflect.Descriptor instead.
func (*PredeclaredPlatformQueueConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_rawDescGZIP(), []int{1}
}

func (x *PredeclaredPlatformQueueConfiguration) GetInstanceNamePrefix() string {
	if x != nil {
		return x.InstanceNamePrefix
	}
	return ""
}

func (x *PredeclaredPlatformQueueConfiguration) GetPlatform() *v2.Platform {
	if x != nil {
		return x.Platform
	}
	return nil
}

func (x *PredeclaredPlatformQueueConfiguration) GetMaximumSizeClass() uint32 {
	if x != nil {
		return x.MaximumSizeClass
	}
	return 0
}

func (x *PredeclaredPlatformQueueConfiguration) GetWorkerInvocationStickinessLimits() []*durationpb.Duration {
	if x != nil {
		return x.WorkerInvocationStickinessLimits
	}
	return nil
}

func (x *PredeclaredPlatformQueueConfiguration) GetMaximumQueuedBackgroundLearningOperations() int32 {
	if x != nil {
		return x.MaximumQueuedBackgroundLearningOperations
	}
	return 0
}

func (x *PredeclaredPlatformQueueConfiguration) GetBackgroundLearningOperationPriority() int32 {
	if x != nil {
		return x.BackgroundLearningOperationPriority
	}
	return 0
}

var File_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto protoreflect.FileDescriptor

var file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_rawDesc = []byte{
	0x0a, 0x37, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x62, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x62, 0x62, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x62, 0x62, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x1a,
	0x36, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x32,
	0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x31, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x31, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x0c, 0x0a, 0x18, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x5f, 0x0a, 0x12, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x68, 0x74, 0x74, 0x70,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x48, 0x74, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x12, 0x61, 0x0a, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x73, 0x12, 0x61, 0x0a, 0x13, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x72, 0x6f, 0x77, 0x73,
	0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x72,
	0x6f, 0x77, 0x73, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x7a, 0x0a, 0x1b, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x19, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x12, 0x3b, 0x0a, 0x1a, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75,
	0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x45, 0x0a, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12, 0x75, 0x0a, 0x1e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x31, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x1a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12,
	0x8b, 0x01, 0x0a, 0x1b, 0x70, 0x72, 0x65, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x64, 0x5f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x18,
	0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72,
	0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x62, 0x62, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x65,
	0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x19, 0x70, 0x72, 0x65, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x64, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x12, 0x64, 0x0a,
	0x12, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x11, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x72, 0x12, 0x6f, 0x0a, 0x18, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x5f, 0x64, 0x72,
	0x61, 0x69, 0x6e, 0x73, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72,
	0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x44, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x72, 0x12, 0x73, 0x0a, 0x1a, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x18, 0x6b, 0x69, 0x6c, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x61, 0x0a, 0x0d, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3c, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x73, 0x0a, 0x18,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x12, 0x6c, 0x0a, 0x26, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6e, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x21, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x51, 0x75, 0x65, 0x75, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4e,
	0x6f, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4a,
	0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x09, 0x10, 0x0a, 0x4a, 0x04, 0x08, 0x0a, 0x10,
	0x0b, 0x4a, 0x04, 0x08, 0x0d, 0x10, 0x0e, 0x4a, 0x04, 0x08, 0x0e, 0x10, 0x0f, 0x22, 0xf5, 0x03,
	0x0a, 0x25, 0x50, 0x72, 0x65, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x64, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x51, 0x75, 0x65, 0x75, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x45, 0x0a, 0x08, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2e, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x6d, 0x61,
	0x78, 0x69, 0x6d, 0x75, 0x6d, 0x53, 0x69, 0x7a, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x68,
	0x0a, 0x23, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x20, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x6e,
	0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x60, 0x0a, 0x2d, 0x6d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x29, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x42, 0x61,
	0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x53, 0x0a, 0x26, 0x62, 0x61,
	0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x23, 0x62, 0x61, 0x63, 0x6b,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4a,
	0x04, 0x08, 0x04, 0x10, 0x05, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2f, 0x62, 0x62,
	0x2d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x62, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_rawDescOnce sync.Once
	file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_rawDescData = file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_rawDesc
)

func file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_rawDescGZIP() []byte {
	file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_rawDescOnce.Do(func() {
		file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_rawDescData)
	})
	return file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_rawDescData
}

var file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_goTypes = []interface{}{
	(*ApplicationConfiguration)(nil),              // 0: buildbarn.configuration.bb_scheduler.ApplicationConfiguration
	(*PredeclaredPlatformQueueConfiguration)(nil), // 1: buildbarn.configuration.bb_scheduler.PredeclaredPlatformQueueConfiguration
	(*http.ServerConfiguration)(nil),              // 2: buildbarn.configuration.http.ServerConfiguration
	(*grpc.ServerConfiguration)(nil),              // 3: buildbarn.configuration.grpc.ServerConfiguration
	(*blobstore.BlobAccessConfiguration)(nil),     // 4: buildbarn.configuration.blobstore.BlobAccessConfiguration
	(*global.Configuration)(nil),                  // 5: buildbarn.configuration.global.Configuration
	(*auth.AuthorizerConfiguration)(nil),          // 6: buildbarn.configuration.auth.AuthorizerConfiguration
	(*scheduler.ActionRouterConfiguration)(nil),   // 7: buildbarn.configuration.scheduler.ActionRouterConfiguration
	(*durationpb.Duration)(nil),                   // 8: google.protobuf.Duration
	(*v2.Platform)(nil),                           // 9: build.bazel.remote.execution.v2.Platform
}
var file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_depIdxs = []int32{
	2,  // 0: buildbarn.configuration.bb_scheduler.ApplicationConfiguration.admin_http_servers:type_name -> buildbarn.configuration.http.ServerConfiguration
	3,  // 1: buildbarn.configuration.bb_scheduler.ApplicationConfiguration.client_grpc_servers:type_name -> buildbarn.configuration.grpc.ServerConfiguration
	3,  // 2: buildbarn.configuration.bb_scheduler.ApplicationConfiguration.worker_grpc_servers:type_name -> buildbarn.configuration.grpc.ServerConfiguration
	4,  // 3: buildbarn.configuration.bb_scheduler.ApplicationConfiguration.content_addressable_storage:type_name -> buildbarn.configuration.blobstore.BlobAccessConfiguration
	5,  // 4: buildbarn.configuration.bb_scheduler.ApplicationConfiguration.global:type_name -> buildbarn.configuration.global.Configuration
	3,  // 5: buildbarn.configuration.bb_scheduler.ApplicationConfiguration.build_queue_state_grpc_servers:type_name -> buildbarn.configuration.grpc.ServerConfiguration
	1,  // 6: buildbarn.configuration.bb_scheduler.ApplicationConfiguration.predeclared_platform_queues:type_name -> buildbarn.configuration.bb_scheduler.PredeclaredPlatformQueueConfiguration
	6,  // 7: buildbarn.configuration.bb_scheduler.ApplicationConfiguration.execute_authorizer:type_name -> buildbarn.configuration.auth.AuthorizerConfiguration
	6,  // 8: buildbarn.configuration.bb_scheduler.ApplicationConfiguration.modify_drains_authorizer:type_name -> buildbarn.configuration.auth.AuthorizerConfiguration
	6,  // 9: buildbarn.configuration.bb_scheduler.ApplicationConfiguration.kill_operations_authorizer:type_name -> buildbarn.configuration.auth.AuthorizerConfiguration
	7,  // 10: buildbarn.configuration.bb_scheduler.ApplicationConfiguration.action_router:type_name -> buildbarn.configuration.scheduler.ActionRouterConfiguration
	4,  // 11: buildbarn.configuration.bb_scheduler.ApplicationConfiguration.initial_size_class_cache:type_name -> buildbarn.configuration.blobstore.BlobAccessConfiguration
	8,  // 12: buildbarn.configuration.bb_scheduler.ApplicationConfiguration.platform_queue_with_no_workers_timeout:type_name -> google.protobuf.Duration
	9,  // 13: buildbarn.configuration.bb_scheduler.PredeclaredPlatformQueueConfiguration.platform:type_name -> build.bazel.remote.execution.v2.Platform
	8,  // 14: buildbarn.configuration.bb_scheduler.PredeclaredPlatformQueueConfiguration.worker_invocation_stickiness_limits:type_name -> google.protobuf.Duration
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_init() }
func file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_init() {
	if File_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationConfiguration); i {
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
		file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredeclaredPlatformQueueConfiguration); i {
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
			RawDescriptor: file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_goTypes,
		DependencyIndexes: file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_depIdxs,
		MessageInfos:      file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_msgTypes,
	}.Build()
	File_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto = out.File
	file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_rawDesc = nil
	file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_goTypes = nil
	file_pkg_proto_configuration_bb_scheduler_bb_scheduler_proto_depIdxs = nil
}
