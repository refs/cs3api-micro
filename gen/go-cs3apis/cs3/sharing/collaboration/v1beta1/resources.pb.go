// Copyright 2018-2019 CERN
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
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.8.0
// source: cs3/sharing/collaboration/v1beta1/resources.proto

package collaborationv1beta1

import (
	v1beta11 "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/identity/user/v1beta1"
	v1beta1 "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/storage/provider/v1beta1"
	v1beta12 "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/types/v1beta1"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The state of the share.
type ShareState int32

const (
	// The share is no longer valid, for example, the share expired.
	ShareState_SHARE_STATE_INVALID ShareState = 0
	// New shares MUST be created in the SHARE_STATE_PENDING state.
	// This state means the share is pending to be accepted or rejected
	// by the recipient of the share.
	ShareState_SHARE_STATE_PENDING ShareState = 1
	// The recipient of the share has accepted the share.
	ShareState_SHARE_STATE_ACCEPTED ShareState = 2
	// The recipient of the share has rejected the share.
	// Do not means the share is removed, the recipient MAY
	// change the state to accepted or pending.
	ShareState_SHARE_STATE_REJECTED ShareState = 3
)

// Enum value maps for ShareState.
var (
	ShareState_name = map[int32]string{
		0: "SHARE_STATE_INVALID",
		1: "SHARE_STATE_PENDING",
		2: "SHARE_STATE_ACCEPTED",
		3: "SHARE_STATE_REJECTED",
	}
	ShareState_value = map[string]int32{
		"SHARE_STATE_INVALID":  0,
		"SHARE_STATE_PENDING":  1,
		"SHARE_STATE_ACCEPTED": 2,
		"SHARE_STATE_REJECTED": 3,
	}
)

func (x ShareState) Enum() *ShareState {
	p := new(ShareState)
	*p = x
	return p
}

func (x ShareState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ShareState) Descriptor() protoreflect.EnumDescriptor {
	return file_cs3_sharing_collaboration_v1beta1_resources_proto_enumTypes[0].Descriptor()
}

func (ShareState) Type() protoreflect.EnumType {
	return &file_cs3_sharing_collaboration_v1beta1_resources_proto_enumTypes[0]
}

func (x ShareState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ShareState.Descriptor instead.
func (ShareState) EnumDescriptor() ([]byte, []int) {
	return file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDescGZIP(), []int{0}
}

// Shares are relationships between a resource owner
// (usually the authenticated user) who grants permissions to a recipient (grantee)
// on a specified resource (resource_id). UserShares represents both user and groups.
type Share struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// Opaque unique identifier of the share.
	Id *ShareId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// REQUIRED.
	// Unique identifier of the shared resource.
	ResourceId *v1beta1.ResourceId `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// REQUIRED.
	// Permissions for the grantee to use
	// the resource.
	Permissions *SharePermissions `protobuf:"bytes,3,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// REQUIRED.
	// The receiver of the share, like a user, group ...
	Grantee *v1beta1.Grantee `protobuf:"bytes,4,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// REQUIRED.
	// Uniquely identifies the owner of the share
	// (the resource owner at the time of creating the share).
	// In case the ownership of the underlying resource changes
	// the share owner field MAY change to reflect the change of ownsership.
	Owner *v1beta11.UserId `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	// REQUIRED.
	// Uniquely identifies a principal who initiates the share creation.
	// A creator can create shares on behalf of the owner (because of re-sharing,
	// because belonging to special groups, ...).
	// Creator and owner often result in being the same principal.
	Creator *v1beta11.UserId `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	// REQUIRED.
	// Creation time of the share.
	Ctime *v1beta12.Timestamp `protobuf:"bytes,7,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// REQUIRED.
	// Last modification time of the share.
	Mtime *v1beta12.Timestamp `protobuf:"bytes,8,opt,name=mtime,proto3" json:"mtime,omitempty"`
}

func (x *Share) Reset() {
	*x = Share{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Share) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Share) ProtoMessage() {}

func (x *Share) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Share.ProtoReflect.Descriptor instead.
func (*Share) Descriptor() ([]byte, []int) {
	return file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDescGZIP(), []int{0}
}

func (x *Share) GetId() *ShareId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Share) GetResourceId() *v1beta1.ResourceId {
	if x != nil {
		return x.ResourceId
	}
	return nil
}

func (x *Share) GetPermissions() *SharePermissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *Share) GetGrantee() *v1beta1.Grantee {
	if x != nil {
		return x.Grantee
	}
	return nil
}

func (x *Share) GetOwner() *v1beta11.UserId {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *Share) GetCreator() *v1beta11.UserId {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *Share) GetCtime() *v1beta12.Timestamp {
	if x != nil {
		return x.Ctime
	}
	return nil
}

func (x *Share) GetMtime() *v1beta12.Timestamp {
	if x != nil {
		return x.Mtime
	}
	return nil
}

// The permissions for a share.
type SharePermissions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permissions *v1beta1.ResourcePermissions `protobuf:"bytes,1,opt,name=permissions,proto3" json:"permissions,omitempty"` // TODO(labkode): additional permissions for shares like re-sharing
}

func (x *SharePermissions) Reset() {
	*x = SharePermissions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharePermissions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharePermissions) ProtoMessage() {}

func (x *SharePermissions) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharePermissions.ProtoReflect.Descriptor instead.
func (*SharePermissions) Descriptor() ([]byte, []int) {
	return file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDescGZIP(), []int{1}
}

func (x *SharePermissions) GetPermissions() *v1beta1.ResourcePermissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

// A received share is the share that a grantee will receive.
// It expands the original share by adding state to the share,
// a display name from the perspective of the grantee and a
// resource mount path in case the share will be mounted
// in a path in a storage provider.
type ReceivedShare struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	Share *Share `protobuf:"bytes,1,opt,name=share,proto3" json:"share,omitempty"`
	// REQUIRED.
	// The state of the share.
	State ShareState `protobuf:"varint,2,opt,name=state,proto3,enum=cs3.sharing.collaboration.v1beta1.ShareState" json:"state,omitempty"`
}

func (x *ReceivedShare) Reset() {
	*x = ReceivedShare{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceivedShare) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceivedShare) ProtoMessage() {}

func (x *ReceivedShare) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceivedShare.ProtoReflect.Descriptor instead.
func (*ReceivedShare) Descriptor() ([]byte, []int) {
	return file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDescGZIP(), []int{2}
}

func (x *ReceivedShare) GetShare() *Share {
	if x != nil {
		return x.Share
	}
	return nil
}

func (x *ReceivedShare) GetState() ShareState {
	if x != nil {
		return x.State
	}
	return ShareState_SHARE_STATE_INVALID
}

// Uniquely identifies a share in the share provider.
// A share MUST be uniquely identify by four (4) elements:
// 1) The share provider id
// 2) The owner of the share
// 3) The resource id
// 4) The grantee for the share
// This 4-tuple MUST be unique.
// For example, owner Alice shares the resource /home/docs with id
// home:1234 to an user named Bob. The 4-tuple will consist of
// 1) The share provider id = "user"
// 2) The owner of the share = "Alice"
// 3) The resource id = "home:1234"
// 4) The grantee for the share = Grantee("type" = "user", "" => "Bob")
type ShareKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	Owner *v1beta11.UserId `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// REQUIRED.
	ResourceId *v1beta1.ResourceId `protobuf:"bytes,3,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// REQUIRED.
	Grantee *v1beta1.Grantee `protobuf:"bytes,4,opt,name=grantee,proto3" json:"grantee,omitempty"`
}

func (x *ShareKey) Reset() {
	*x = ShareKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareKey) ProtoMessage() {}

func (x *ShareKey) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareKey.ProtoReflect.Descriptor instead.
func (*ShareKey) Descriptor() ([]byte, []int) {
	return file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDescGZIP(), []int{3}
}

func (x *ShareKey) GetOwner() *v1beta11.UserId {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *ShareKey) GetResourceId() *v1beta1.ResourceId {
	if x != nil {
		return x.ResourceId
	}
	return nil
}

func (x *ShareKey) GetGrantee() *v1beta1.Grantee {
	if x != nil {
		return x.Grantee
	}
	return nil
}

// A share id identifies uniquely a // share in the share provider namespace.
// A ShareId MUST be unique inside the share provider.
type ShareId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The internal id used by service implementor to
	// uniquely Collaboration the share in the internal
	// implementation of the service.
	OpaqueId string `protobuf:"bytes,2,opt,name=opaque_id,json=opaqueId,proto3" json:"opaque_id,omitempty"`
}

func (x *ShareId) Reset() {
	*x = ShareId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareId) ProtoMessage() {}

func (x *ShareId) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareId.ProtoReflect.Descriptor instead.
func (*ShareId) Descriptor() ([]byte, []int) {
	return file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDescGZIP(), []int{4}
}

func (x *ShareId) GetOpaqueId() string {
	if x != nil {
		return x.OpaqueId
	}
	return ""
}

// The mechanism to identify a share
// in the share provider namespace.
type ShareReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// One of the specifications MUST be specified.
	//
	// Types that are assignable to Spec:
	//	*ShareReference_Id
	//	*ShareReference_Key
	Spec isShareReference_Spec `protobuf_oneof:"spec"`
}

func (x *ShareReference) Reset() {
	*x = ShareReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareReference) ProtoMessage() {}

func (x *ShareReference) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareReference.ProtoReflect.Descriptor instead.
func (*ShareReference) Descriptor() ([]byte, []int) {
	return file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDescGZIP(), []int{5}
}

func (m *ShareReference) GetSpec() isShareReference_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (x *ShareReference) GetId() *ShareId {
	if x, ok := x.GetSpec().(*ShareReference_Id); ok {
		return x.Id
	}
	return nil
}

func (x *ShareReference) GetKey() *ShareKey {
	if x, ok := x.GetSpec().(*ShareReference_Key); ok {
		return x.Key
	}
	return nil
}

type isShareReference_Spec interface {
	isShareReference_Spec()
}

type ShareReference_Id struct {
	// The id of the share.
	Id *ShareId `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type ShareReference_Key struct {
	// The combination of fields that
	// make the share unique.
	Key *ShareKey `protobuf:"bytes,2,opt,name=key,proto3,oneof"`
}

func (*ShareReference_Id) isShareReference_Spec() {}

func (*ShareReference_Key) isShareReference_Spec() {}

// A share grant specifies the share permissions
// for a grantee.
type ShareGrant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The grantee of the grant.
	Grantee *v1beta1.Grantee `protobuf:"bytes,1,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// REQUIRED.
	// The share permissions for the grant.
	Permissions *SharePermissions `protobuf:"bytes,2,opt,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *ShareGrant) Reset() {
	*x = ShareGrant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareGrant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareGrant) ProtoMessage() {}

func (x *ShareGrant) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareGrant.ProtoReflect.Descriptor instead.
func (*ShareGrant) Descriptor() ([]byte, []int) {
	return file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDescGZIP(), []int{6}
}

func (x *ShareGrant) GetGrantee() *v1beta1.Grantee {
	if x != nil {
		return x.Grantee
	}
	return nil
}

func (x *ShareGrant) GetPermissions() *SharePermissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

var File_cs3_sharing_collaboration_v1beta1_resources_proto protoreflect.FileDescriptor

var file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x73, 0x33, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f,
	0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x21, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x29, 0x63, 0x73, 0x33, 0x2f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2c, 0x63, 0x73, 0x33, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x63, 0x73, 0x33, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84,
	0x04, 0x0a, 0x05, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x3a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x49, 0x64,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x49, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x73, 0x33, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x55, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3f, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x52, 0x07,
	0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x12, 0x3b, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x32, 0x0a,
	0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63,
	0x73, 0x33, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x32, 0x0a, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05,
	0x6d, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x67, 0x0a, 0x10, 0x53, 0x68, 0x61, 0x72, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x53, 0x0a, 0x0b, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x94,
	0x01, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x12, 0x3e, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x12, 0x43, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2d, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xcf, 0x01, 0x0a, 0x08, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4b,
	0x65, 0x79, 0x12, 0x37, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x52, 0x07,
	0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x22, 0x26, 0x0a, 0x07, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x49, 0x64, 0x22,
	0x97, 0x01, 0x0a, 0x0e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x3c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x49, 0x64, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x3f, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x63, 0x73, 0x33, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6c, 0x6c,
	0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x48, 0x00, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x42, 0x06, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0xa4, 0x01, 0x0a, 0x0a, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x73, 0x33, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65,
	0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x12, 0x55, 0x0a, 0x0b, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33,
	0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2a, 0x72, 0x0a, 0x0a, 0x53, 0x68, 0x61, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x17,
	0x0a, 0x13, 0x53, 0x48, 0x41, 0x52, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x48, 0x41, 0x52, 0x45,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x18, 0x0a, 0x14, 0x53, 0x48, 0x41, 0x52, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x48,
	0x41, 0x52, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54,
	0x45, 0x44, 0x10, 0x03, 0x42, 0x9d, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x73, 0x33,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x14, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x43, 0xaa, 0x02, 0x21, 0x43,
	0x73, 0x33, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x61,
	0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31,
	0xca, 0x02, 0x21, 0x43, 0x73, 0x33, 0x5c, 0x53, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x43,
	0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x42,
	0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDescOnce sync.Once
	file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDescData = file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDesc
)

func file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDescGZIP() []byte {
	file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDescOnce.Do(func() {
		file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDescData)
	})
	return file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDescData
}

var file_cs3_sharing_collaboration_v1beta1_resources_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cs3_sharing_collaboration_v1beta1_resources_proto_goTypes = []interface{}{
	(ShareState)(0),                     // 0: cs3.sharing.collaboration.v1beta1.ShareState
	(*Share)(nil),                       // 1: cs3.sharing.collaboration.v1beta1.Share
	(*SharePermissions)(nil),            // 2: cs3.sharing.collaboration.v1beta1.SharePermissions
	(*ReceivedShare)(nil),               // 3: cs3.sharing.collaboration.v1beta1.ReceivedShare
	(*ShareKey)(nil),                    // 4: cs3.sharing.collaboration.v1beta1.ShareKey
	(*ShareId)(nil),                     // 5: cs3.sharing.collaboration.v1beta1.ShareId
	(*ShareReference)(nil),              // 6: cs3.sharing.collaboration.v1beta1.ShareReference
	(*ShareGrant)(nil),                  // 7: cs3.sharing.collaboration.v1beta1.ShareGrant
	(*v1beta1.ResourceId)(nil),          // 8: cs3.storage.provider.v1beta1.ResourceId
	(*v1beta1.Grantee)(nil),             // 9: cs3.storage.provider.v1beta1.Grantee
	(*v1beta11.UserId)(nil),             // 10: cs3.identity.user.v1beta1.UserId
	(*v1beta12.Timestamp)(nil),          // 11: cs3.types.v1beta1.Timestamp
	(*v1beta1.ResourcePermissions)(nil), // 12: cs3.storage.provider.v1beta1.ResourcePermissions
}
var file_cs3_sharing_collaboration_v1beta1_resources_proto_depIdxs = []int32{
	5,  // 0: cs3.sharing.collaboration.v1beta1.Share.id:type_name -> cs3.sharing.collaboration.v1beta1.ShareId
	8,  // 1: cs3.sharing.collaboration.v1beta1.Share.resource_id:type_name -> cs3.storage.provider.v1beta1.ResourceId
	2,  // 2: cs3.sharing.collaboration.v1beta1.Share.permissions:type_name -> cs3.sharing.collaboration.v1beta1.SharePermissions
	9,  // 3: cs3.sharing.collaboration.v1beta1.Share.grantee:type_name -> cs3.storage.provider.v1beta1.Grantee
	10, // 4: cs3.sharing.collaboration.v1beta1.Share.owner:type_name -> cs3.identity.user.v1beta1.UserId
	10, // 5: cs3.sharing.collaboration.v1beta1.Share.creator:type_name -> cs3.identity.user.v1beta1.UserId
	11, // 6: cs3.sharing.collaboration.v1beta1.Share.ctime:type_name -> cs3.types.v1beta1.Timestamp
	11, // 7: cs3.sharing.collaboration.v1beta1.Share.mtime:type_name -> cs3.types.v1beta1.Timestamp
	12, // 8: cs3.sharing.collaboration.v1beta1.SharePermissions.permissions:type_name -> cs3.storage.provider.v1beta1.ResourcePermissions
	1,  // 9: cs3.sharing.collaboration.v1beta1.ReceivedShare.share:type_name -> cs3.sharing.collaboration.v1beta1.Share
	0,  // 10: cs3.sharing.collaboration.v1beta1.ReceivedShare.state:type_name -> cs3.sharing.collaboration.v1beta1.ShareState
	10, // 11: cs3.sharing.collaboration.v1beta1.ShareKey.owner:type_name -> cs3.identity.user.v1beta1.UserId
	8,  // 12: cs3.sharing.collaboration.v1beta1.ShareKey.resource_id:type_name -> cs3.storage.provider.v1beta1.ResourceId
	9,  // 13: cs3.sharing.collaboration.v1beta1.ShareKey.grantee:type_name -> cs3.storage.provider.v1beta1.Grantee
	5,  // 14: cs3.sharing.collaboration.v1beta1.ShareReference.id:type_name -> cs3.sharing.collaboration.v1beta1.ShareId
	4,  // 15: cs3.sharing.collaboration.v1beta1.ShareReference.key:type_name -> cs3.sharing.collaboration.v1beta1.ShareKey
	9,  // 16: cs3.sharing.collaboration.v1beta1.ShareGrant.grantee:type_name -> cs3.storage.provider.v1beta1.Grantee
	2,  // 17: cs3.sharing.collaboration.v1beta1.ShareGrant.permissions:type_name -> cs3.sharing.collaboration.v1beta1.SharePermissions
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_cs3_sharing_collaboration_v1beta1_resources_proto_init() }
func file_cs3_sharing_collaboration_v1beta1_resources_proto_init() {
	if File_cs3_sharing_collaboration_v1beta1_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Share); i {
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
		file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharePermissions); i {
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
		file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceivedShare); i {
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
		file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareKey); i {
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
		file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareId); i {
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
		file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareReference); i {
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
		file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareGrant); i {
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
	file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*ShareReference_Id)(nil),
		(*ShareReference_Key)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cs3_sharing_collaboration_v1beta1_resources_proto_goTypes,
		DependencyIndexes: file_cs3_sharing_collaboration_v1beta1_resources_proto_depIdxs,
		EnumInfos:         file_cs3_sharing_collaboration_v1beta1_resources_proto_enumTypes,
		MessageInfos:      file_cs3_sharing_collaboration_v1beta1_resources_proto_msgTypes,
	}.Build()
	File_cs3_sharing_collaboration_v1beta1_resources_proto = out.File
	file_cs3_sharing_collaboration_v1beta1_resources_proto_rawDesc = nil
	file_cs3_sharing_collaboration_v1beta1_resources_proto_goTypes = nil
	file_cs3_sharing_collaboration_v1beta1_resources_proto_depIdxs = nil
}
