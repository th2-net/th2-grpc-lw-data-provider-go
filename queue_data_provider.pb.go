//
// Copyright 2022-2025 Exactpro (Exactpro Systems Limited)
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
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: th2_grpc_lw_data_provider/queue_data_provider.proto

package th2_grpc_lw_data_provider_go

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type MessageGroupsQueueSearchRequest struct {
	state          protoimpl.MessageState                        `protogen:"open.v1"`
	StartTimestamp *timestamppb.Timestamp                        `protobuf:"bytes,1,opt,name=start_timestamp,json=startTimestamp,proto3" json:"start_timestamp,omitempty"`
	EndTimestamp   *timestamppb.Timestamp                        `protobuf:"bytes,2,opt,name=end_timestamp,json=endTimestamp,proto3" json:"end_timestamp,omitempty"`
	MessageGroup   []*MessageGroupsQueueSearchRequest_BookGroups `protobuf:"bytes,7,rep,name=message_group,json=messageGroup,proto3" json:"message_group,omitempty"`
	SyncInterval   *durationpb.Duration                          `protobuf:"bytes,4,opt,name=sync_interval,json=syncInterval,proto3" json:"sync_interval,omitempty"`
	ExternalQueue  string                                        `protobuf:"bytes,5,opt,name=external_queue,json=externalQueue,proto3" json:"external_queue,omitempty"`
	KeepAlive      bool                                          `protobuf:"varint,6,opt,name=keep_alive,json=keepAlive,proto3" json:"keep_alive,omitempty"`
	// *
	// Enables sending raw batches directly to external_queue. Sends raw directly first then sends raw to pins
	SendRawDirectly bool `protobuf:"varint,8,opt,name=send_raw_directly,json=sendRawDirectly,proto3" json:"send_raw_directly,omitempty"`
	// *
	// If enabled the message won't be sent to codecs
	RawOnly       bool `protobuf:"varint,9,opt,name=raw_only,json=rawOnly,proto3" json:"raw_only,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageGroupsQueueSearchRequest) Reset() {
	*x = MessageGroupsQueueSearchRequest{}
	mi := &file_th2_grpc_lw_data_provider_queue_data_provider_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageGroupsQueueSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageGroupsQueueSearchRequest) ProtoMessage() {}

func (x *MessageGroupsQueueSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_th2_grpc_lw_data_provider_queue_data_provider_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageGroupsQueueSearchRequest.ProtoReflect.Descriptor instead.
func (*MessageGroupsQueueSearchRequest) Descriptor() ([]byte, []int) {
	return file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDescGZIP(), []int{0}
}

func (x *MessageGroupsQueueSearchRequest) GetStartTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTimestamp
	}
	return nil
}

func (x *MessageGroupsQueueSearchRequest) GetEndTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTimestamp
	}
	return nil
}

func (x *MessageGroupsQueueSearchRequest) GetMessageGroup() []*MessageGroupsQueueSearchRequest_BookGroups {
	if x != nil {
		return x.MessageGroup
	}
	return nil
}

func (x *MessageGroupsQueueSearchRequest) GetSyncInterval() *durationpb.Duration {
	if x != nil {
		return x.SyncInterval
	}
	return nil
}

func (x *MessageGroupsQueueSearchRequest) GetExternalQueue() string {
	if x != nil {
		return x.ExternalQueue
	}
	return ""
}

func (x *MessageGroupsQueueSearchRequest) GetKeepAlive() bool {
	if x != nil {
		return x.KeepAlive
	}
	return false
}

func (x *MessageGroupsQueueSearchRequest) GetSendRawDirectly() bool {
	if x != nil {
		return x.SendRawDirectly
	}
	return false
}

func (x *MessageGroupsQueueSearchRequest) GetRawOnly() bool {
	if x != nil {
		return x.RawOnly
	}
	return false
}

type MessageLoadedStatistic struct {
	state         protoimpl.MessageState              `protogen:"open.v1"`
	Stat          []*MessageLoadedStatistic_GroupStat `protobuf:"bytes,1,rep,name=stat,proto3" json:"stat,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageLoadedStatistic) Reset() {
	*x = MessageLoadedStatistic{}
	mi := &file_th2_grpc_lw_data_provider_queue_data_provider_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageLoadedStatistic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageLoadedStatistic) ProtoMessage() {}

func (x *MessageLoadedStatistic) ProtoReflect() protoreflect.Message {
	mi := &file_th2_grpc_lw_data_provider_queue_data_provider_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageLoadedStatistic.ProtoReflect.Descriptor instead.
func (*MessageLoadedStatistic) Descriptor() ([]byte, []int) {
	return file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDescGZIP(), []int{1}
}

func (x *MessageLoadedStatistic) GetStat() []*MessageLoadedStatistic_GroupStat {
	if x != nil {
		return x.Stat
	}
	return nil
}

type EventQueueSearchRequest struct {
	state          protoimpl.MessageState                `protogen:"open.v1"`
	StartTimestamp *timestamppb.Timestamp                `protobuf:"bytes,1,opt,name=start_timestamp,json=startTimestamp,proto3" json:"start_timestamp,omitempty"`
	EndTimestamp   *timestamppb.Timestamp                `protobuf:"bytes,2,opt,name=end_timestamp,json=endTimestamp,proto3" json:"end_timestamp,omitempty"`
	EventScopes    []*EventQueueSearchRequest_BookScopes `protobuf:"bytes,7,rep,name=event_scopes,json=eventScopes,proto3" json:"event_scopes,omitempty"`
	SyncInterval   *durationpb.Duration                  `protobuf:"bytes,4,opt,name=sync_interval,json=syncInterval,proto3" json:"sync_interval,omitempty"`
	ExternalQueue  string                                `protobuf:"bytes,5,opt,name=external_queue,json=externalQueue,proto3" json:"external_queue,omitempty"`
	KeepAlive      bool                                  `protobuf:"varint,6,opt,name=keep_alive,json=keepAlive,proto3" json:"keep_alive,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *EventQueueSearchRequest) Reset() {
	*x = EventQueueSearchRequest{}
	mi := &file_th2_grpc_lw_data_provider_queue_data_provider_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventQueueSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventQueueSearchRequest) ProtoMessage() {}

func (x *EventQueueSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_th2_grpc_lw_data_provider_queue_data_provider_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventQueueSearchRequest.ProtoReflect.Descriptor instead.
func (*EventQueueSearchRequest) Descriptor() ([]byte, []int) {
	return file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDescGZIP(), []int{2}
}

func (x *EventQueueSearchRequest) GetStartTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTimestamp
	}
	return nil
}

func (x *EventQueueSearchRequest) GetEndTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTimestamp
	}
	return nil
}

func (x *EventQueueSearchRequest) GetEventScopes() []*EventQueueSearchRequest_BookScopes {
	if x != nil {
		return x.EventScopes
	}
	return nil
}

func (x *EventQueueSearchRequest) GetSyncInterval() *durationpb.Duration {
	if x != nil {
		return x.SyncInterval
	}
	return nil
}

func (x *EventQueueSearchRequest) GetExternalQueue() string {
	if x != nil {
		return x.ExternalQueue
	}
	return ""
}

func (x *EventQueueSearchRequest) GetKeepAlive() bool {
	if x != nil {
		return x.KeepAlive
	}
	return false
}

type EventLoadedStatistic struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	Stat          []*EventLoadedStatistic_ScopeStat `protobuf:"bytes,1,rep,name=stat,proto3" json:"stat,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventLoadedStatistic) Reset() {
	*x = EventLoadedStatistic{}
	mi := &file_th2_grpc_lw_data_provider_queue_data_provider_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventLoadedStatistic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventLoadedStatistic) ProtoMessage() {}

func (x *EventLoadedStatistic) ProtoReflect() protoreflect.Message {
	mi := &file_th2_grpc_lw_data_provider_queue_data_provider_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventLoadedStatistic.ProtoReflect.Descriptor instead.
func (*EventLoadedStatistic) Descriptor() ([]byte, []int) {
	return file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDescGZIP(), []int{3}
}

func (x *EventLoadedStatistic) GetStat() []*EventLoadedStatistic_ScopeStat {
	if x != nil {
		return x.Stat
	}
	return nil
}

type MessageGroupsQueueSearchRequest_BookGroups struct {
	state         protoimpl.MessageState              `protogen:"open.v1"`
	BookId        *BookId                             `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Group         []*MessageGroupsSearchRequest_Group `protobuf:"bytes,2,rep,name=group,proto3" json:"group,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageGroupsQueueSearchRequest_BookGroups) Reset() {
	*x = MessageGroupsQueueSearchRequest_BookGroups{}
	mi := &file_th2_grpc_lw_data_provider_queue_data_provider_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageGroupsQueueSearchRequest_BookGroups) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageGroupsQueueSearchRequest_BookGroups) ProtoMessage() {}

func (x *MessageGroupsQueueSearchRequest_BookGroups) ProtoReflect() protoreflect.Message {
	mi := &file_th2_grpc_lw_data_provider_queue_data_provider_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageGroupsQueueSearchRequest_BookGroups.ProtoReflect.Descriptor instead.
func (*MessageGroupsQueueSearchRequest_BookGroups) Descriptor() ([]byte, []int) {
	return file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MessageGroupsQueueSearchRequest_BookGroups) GetBookId() *BookId {
	if x != nil {
		return x.BookId
	}
	return nil
}

func (x *MessageGroupsQueueSearchRequest_BookGroups) GetGroup() []*MessageGroupsSearchRequest_Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type MessageLoadedStatistic_GroupStat struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	BookId        *BookId                           `protobuf:"bytes,3,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Group         *MessageGroupsSearchRequest_Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Count         uint64                            `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageLoadedStatistic_GroupStat) Reset() {
	*x = MessageLoadedStatistic_GroupStat{}
	mi := &file_th2_grpc_lw_data_provider_queue_data_provider_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageLoadedStatistic_GroupStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageLoadedStatistic_GroupStat) ProtoMessage() {}

func (x *MessageLoadedStatistic_GroupStat) ProtoReflect() protoreflect.Message {
	mi := &file_th2_grpc_lw_data_provider_queue_data_provider_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageLoadedStatistic_GroupStat.ProtoReflect.Descriptor instead.
func (*MessageLoadedStatistic_GroupStat) Descriptor() ([]byte, []int) {
	return file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDescGZIP(), []int{1, 0}
}

func (x *MessageLoadedStatistic_GroupStat) GetBookId() *BookId {
	if x != nil {
		return x.BookId
	}
	return nil
}

func (x *MessageLoadedStatistic_GroupStat) GetGroup() *MessageGroupsSearchRequest_Group {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *MessageLoadedStatistic_GroupStat) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type EventQueueSearchRequest_BookScopes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookId        *BookId                `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Scope         []*EventScope          `protobuf:"bytes,2,rep,name=scope,proto3" json:"scope,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventQueueSearchRequest_BookScopes) Reset() {
	*x = EventQueueSearchRequest_BookScopes{}
	mi := &file_th2_grpc_lw_data_provider_queue_data_provider_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventQueueSearchRequest_BookScopes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventQueueSearchRequest_BookScopes) ProtoMessage() {}

func (x *EventQueueSearchRequest_BookScopes) ProtoReflect() protoreflect.Message {
	mi := &file_th2_grpc_lw_data_provider_queue_data_provider_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventQueueSearchRequest_BookScopes.ProtoReflect.Descriptor instead.
func (*EventQueueSearchRequest_BookScopes) Descriptor() ([]byte, []int) {
	return file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDescGZIP(), []int{2, 0}
}

func (x *EventQueueSearchRequest_BookScopes) GetBookId() *BookId {
	if x != nil {
		return x.BookId
	}
	return nil
}

func (x *EventQueueSearchRequest_BookScopes) GetScope() []*EventScope {
	if x != nil {
		return x.Scope
	}
	return nil
}

type EventLoadedStatistic_ScopeStat struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookId        *BookId                `protobuf:"bytes,3,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Scope         *EventScope            `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope,omitempty"`
	Count         uint64                 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventLoadedStatistic_ScopeStat) Reset() {
	*x = EventLoadedStatistic_ScopeStat{}
	mi := &file_th2_grpc_lw_data_provider_queue_data_provider_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventLoadedStatistic_ScopeStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventLoadedStatistic_ScopeStat) ProtoMessage() {}

func (x *EventLoadedStatistic_ScopeStat) ProtoReflect() protoreflect.Message {
	mi := &file_th2_grpc_lw_data_provider_queue_data_provider_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventLoadedStatistic_ScopeStat.ProtoReflect.Descriptor instead.
func (*EventLoadedStatistic_ScopeStat) Descriptor() ([]byte, []int) {
	return file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDescGZIP(), []int{3, 0}
}

func (x *EventLoadedStatistic_ScopeStat) GetBookId() *BookId {
	if x != nil {
		return x.BookId
	}
	return nil
}

func (x *EventLoadedStatistic_ScopeStat) GetScope() *EventScope {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *EventLoadedStatistic_ScopeStat) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_th2_grpc_lw_data_provider_queue_data_provider_proto protoreflect.FileDescriptor

var file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDesc = string([]byte{
	0x0a, 0x33, 0x74, 0x68, 0x32, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x77, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x74, 0x68, 0x32, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6c, 0x77, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x74, 0x68,
	0x32, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x6c, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5,
	0x04, 0x0a, 0x1f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x43, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3f, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x65, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x40, 0x2e, 0x74, 0x68, 0x32, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x6c, 0x77, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x52, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x3e, 0x0a, 0x0d, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12,
	0x25, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x61,
	0x6c, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6b, 0x65, 0x65, 0x70,
	0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x72, 0x61,
	0x77, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6c, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x52, 0x61, 0x77, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6c,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x61, 0x77, 0x4f, 0x6e, 0x6c, 0x79, 0x1a, 0x91, 0x01, 0x0a,
	0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x35, 0x0a, 0x07, 0x62,
	0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74,
	0x68, 0x32, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x6c, 0x77, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b,
	0x49, 0x64, 0x12, 0x4c, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x74, 0x68, 0x32, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6c, 0x77, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x8d, 0x02, 0x0a, 0x16, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x12, 0x4a, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x36, 0x2e, 0x74, 0x68, 0x32, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x6c, 0x77, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x6f,
	0x61, 0x64, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x52, 0x04, 0x73, 0x74, 0x61, 0x74, 0x1a, 0xa6, 0x01,
	0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x12, 0x35, 0x0a, 0x07, 0x62,
	0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74,
	0x68, 0x32, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x6c, 0x77, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b,
	0x49, 0x64, 0x12, 0x4c, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x74, 0x68, 0x32, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6c, 0x77, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xff, 0x03, 0x0a, 0x17, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x43, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3f, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x5b, 0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x74, 0x68, 0x32, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x6c, 0x77, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x52, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x6b, 0x65, 0x65, 0x70, 0x5f, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x1a, 0x7b, 0x0a, 0x0a, 0x42,
	0x6f, 0x6f, 0x6b, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x07, 0x62, 0x6f, 0x6f,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x68, 0x32,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6c,
	0x77, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64,
	0x12, 0x36, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x74, 0x68, 0x32, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x6c, 0x77, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22, 0xf3, 0x01, 0x0a, 0x14, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x12, 0x48, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x34, 0x2e, 0x74, 0x68, 0x32, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x6c, 0x77, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x61, 0x64,
	0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x52, 0x04, 0x73, 0x74, 0x61, 0x74, 0x1a, 0x90, 0x01, 0x0a, 0x09,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x53, 0x74, 0x61, 0x74, 0x12, 0x35, 0x0a, 0x07, 0x62, 0x6f, 0x6f,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x68, 0x32,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6c,
	0x77, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64,
	0x12, 0x36, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x74, 0x68, 0x32, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x6c, 0x77, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xfa,
	0x01, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x75, 0x65, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x7a, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x35, 0x2e, 0x74, 0x68,
	0x32, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x6c, 0x77, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x74, 0x68, 0x32, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6c, 0x77, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x12, 0x69, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x2d, 0x2e, 0x74, 0x68, 0x32, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x6c, 0x77, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x74, 0x68, 0x32, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x6c, 0x77, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x61, 0x64,
	0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x42, 0x5a, 0x0a, 0x25, 0x63,
	0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x63, 0x74, 0x70, 0x72, 0x6f, 0x2e, 0x74, 0x68, 0x32, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6c, 0x77, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x32, 0x2d, 0x6e, 0x65, 0x74, 0x2f, 0x74, 0x68, 0x32, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x6c, 0x77, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDescOnce sync.Once
	file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDescData []byte
)

func file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDescGZIP() []byte {
	file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDescOnce.Do(func() {
		file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDesc), len(file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDesc)))
	})
	return file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDescData
}

var file_th2_grpc_lw_data_provider_queue_data_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_th2_grpc_lw_data_provider_queue_data_provider_proto_goTypes = []any{
	(*MessageGroupsQueueSearchRequest)(nil),            // 0: th2.data_provider.lw.MessageGroupsQueueSearchRequest
	(*MessageLoadedStatistic)(nil),                     // 1: th2.data_provider.lw.MessageLoadedStatistic
	(*EventQueueSearchRequest)(nil),                    // 2: th2.data_provider.lw.EventQueueSearchRequest
	(*EventLoadedStatistic)(nil),                       // 3: th2.data_provider.lw.EventLoadedStatistic
	(*MessageGroupsQueueSearchRequest_BookGroups)(nil), // 4: th2.data_provider.lw.MessageGroupsQueueSearchRequest.BookGroups
	(*MessageLoadedStatistic_GroupStat)(nil),           // 5: th2.data_provider.lw.MessageLoadedStatistic.GroupStat
	(*EventQueueSearchRequest_BookScopes)(nil),         // 6: th2.data_provider.lw.EventQueueSearchRequest.BookScopes
	(*EventLoadedStatistic_ScopeStat)(nil),             // 7: th2.data_provider.lw.EventLoadedStatistic.ScopeStat
	(*timestamppb.Timestamp)(nil),                      // 8: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),                        // 9: google.protobuf.Duration
	(*BookId)(nil),                                     // 10: th2.data_provider.lw.BookId
	(*MessageGroupsSearchRequest_Group)(nil),           // 11: th2.data_provider.lw.MessageGroupsSearchRequest.Group
	(*EventScope)(nil),                                 // 12: th2.data_provider.lw.EventScope
}
var file_th2_grpc_lw_data_provider_queue_data_provider_proto_depIdxs = []int32{
	8,  // 0: th2.data_provider.lw.MessageGroupsQueueSearchRequest.start_timestamp:type_name -> google.protobuf.Timestamp
	8,  // 1: th2.data_provider.lw.MessageGroupsQueueSearchRequest.end_timestamp:type_name -> google.protobuf.Timestamp
	4,  // 2: th2.data_provider.lw.MessageGroupsQueueSearchRequest.message_group:type_name -> th2.data_provider.lw.MessageGroupsQueueSearchRequest.BookGroups
	9,  // 3: th2.data_provider.lw.MessageGroupsQueueSearchRequest.sync_interval:type_name -> google.protobuf.Duration
	5,  // 4: th2.data_provider.lw.MessageLoadedStatistic.stat:type_name -> th2.data_provider.lw.MessageLoadedStatistic.GroupStat
	8,  // 5: th2.data_provider.lw.EventQueueSearchRequest.start_timestamp:type_name -> google.protobuf.Timestamp
	8,  // 6: th2.data_provider.lw.EventQueueSearchRequest.end_timestamp:type_name -> google.protobuf.Timestamp
	6,  // 7: th2.data_provider.lw.EventQueueSearchRequest.event_scopes:type_name -> th2.data_provider.lw.EventQueueSearchRequest.BookScopes
	9,  // 8: th2.data_provider.lw.EventQueueSearchRequest.sync_interval:type_name -> google.protobuf.Duration
	7,  // 9: th2.data_provider.lw.EventLoadedStatistic.stat:type_name -> th2.data_provider.lw.EventLoadedStatistic.ScopeStat
	10, // 10: th2.data_provider.lw.MessageGroupsQueueSearchRequest.BookGroups.book_id:type_name -> th2.data_provider.lw.BookId
	11, // 11: th2.data_provider.lw.MessageGroupsQueueSearchRequest.BookGroups.group:type_name -> th2.data_provider.lw.MessageGroupsSearchRequest.Group
	10, // 12: th2.data_provider.lw.MessageLoadedStatistic.GroupStat.book_id:type_name -> th2.data_provider.lw.BookId
	11, // 13: th2.data_provider.lw.MessageLoadedStatistic.GroupStat.group:type_name -> th2.data_provider.lw.MessageGroupsSearchRequest.Group
	10, // 14: th2.data_provider.lw.EventQueueSearchRequest.BookScopes.book_id:type_name -> th2.data_provider.lw.BookId
	12, // 15: th2.data_provider.lw.EventQueueSearchRequest.BookScopes.scope:type_name -> th2.data_provider.lw.EventScope
	10, // 16: th2.data_provider.lw.EventLoadedStatistic.ScopeStat.book_id:type_name -> th2.data_provider.lw.BookId
	12, // 17: th2.data_provider.lw.EventLoadedStatistic.ScopeStat.scope:type_name -> th2.data_provider.lw.EventScope
	0,  // 18: th2.data_provider.lw.QueueDataProvider.SearchMessageGroups:input_type -> th2.data_provider.lw.MessageGroupsQueueSearchRequest
	2,  // 19: th2.data_provider.lw.QueueDataProvider.SearchEvents:input_type -> th2.data_provider.lw.EventQueueSearchRequest
	1,  // 20: th2.data_provider.lw.QueueDataProvider.SearchMessageGroups:output_type -> th2.data_provider.lw.MessageLoadedStatistic
	3,  // 21: th2.data_provider.lw.QueueDataProvider.SearchEvents:output_type -> th2.data_provider.lw.EventLoadedStatistic
	20, // [20:22] is the sub-list for method output_type
	18, // [18:20] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_th2_grpc_lw_data_provider_queue_data_provider_proto_init() }
func file_th2_grpc_lw_data_provider_queue_data_provider_proto_init() {
	if File_th2_grpc_lw_data_provider_queue_data_provider_proto != nil {
		return
	}
	file_th2_grpc_lw_data_provider_lw_data_provider_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDesc), len(file_th2_grpc_lw_data_provider_queue_data_provider_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_th2_grpc_lw_data_provider_queue_data_provider_proto_goTypes,
		DependencyIndexes: file_th2_grpc_lw_data_provider_queue_data_provider_proto_depIdxs,
		MessageInfos:      file_th2_grpc_lw_data_provider_queue_data_provider_proto_msgTypes,
	}.Build()
	File_th2_grpc_lw_data_provider_queue_data_provider_proto = out.File
	file_th2_grpc_lw_data_provider_queue_data_provider_proto_goTypes = nil
	file_th2_grpc_lw_data_provider_queue_data_provider_proto_depIdxs = nil
}
