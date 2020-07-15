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
// source: cs3/preferences/v1beta1/preferences_api.proto

package preferencesv1beta1

import (
	context "context"
	v1beta1 "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/rpc/v1beta1"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SetKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// REQUIRED.
	// The value associated with the key.
	Val string `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *SetKeyRequest) Reset() {
	*x = SetKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_preferences_v1beta1_preferences_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetKeyRequest) ProtoMessage() {}

func (x *SetKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_preferences_v1beta1_preferences_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetKeyRequest.ProtoReflect.Descriptor instead.
func (*SetKeyRequest) Descriptor() ([]byte, []int) {
	return file_cs3_preferences_v1beta1_preferences_api_proto_rawDescGZIP(), []int{0}
}

func (x *SetKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetKeyRequest) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type SetKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The response status.
	Status *v1beta1.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetKeyResponse) Reset() {
	*x = SetKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_preferences_v1beta1_preferences_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetKeyResponse) ProtoMessage() {}

func (x *SetKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_preferences_v1beta1_preferences_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetKeyResponse.ProtoReflect.Descriptor instead.
func (*SetKeyResponse) Descriptor() ([]byte, []int) {
	return file_cs3_preferences_v1beta1_preferences_api_proto_rawDescGZIP(), []int{1}
}

func (x *SetKeyResponse) GetStatus() *v1beta1.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type GetKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetKeyRequest) Reset() {
	*x = GetKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_preferences_v1beta1_preferences_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyRequest) ProtoMessage() {}

func (x *GetKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_preferences_v1beta1_preferences_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyRequest.ProtoReflect.Descriptor instead.
func (*GetKeyRequest) Descriptor() ([]byte, []int) {
	return file_cs3_preferences_v1beta1_preferences_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	// The response status.
	Status *v1beta1.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// REQUIRED.
	// The value associated with the key.
	Val string `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *GetKeyResponse) Reset() {
	*x = GetKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_preferences_v1beta1_preferences_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyResponse) ProtoMessage() {}

func (x *GetKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_preferences_v1beta1_preferences_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyResponse.ProtoReflect.Descriptor instead.
func (*GetKeyResponse) Descriptor() ([]byte, []int) {
	return file_cs3_preferences_v1beta1_preferences_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetKeyResponse) GetStatus() *v1beta1.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetKeyResponse) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

var File_cs3_preferences_v1beta1_preferences_api_proto protoreflect.FileDescriptor

var file_cs3_preferences_v1beta1_preferences_api_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x73, 0x33, 0x2f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x63, 0x73, 0x33, 0x2e, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x63, 0x73, 0x33, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x41, 0x0a, 0x0e, 0x53,
	0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x63, 0x73, 0x33, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x21,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x22, 0x53, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x32, 0xc6, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x41, 0x50, 0x49, 0x12, 0x59, 0x0a, 0x06, 0x53, 0x65, 0x74,
	0x4b, 0x65, 0x79, 0x12, 0x26, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x73,
	0x33, 0x2e, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x26,
	0x2e, 0x63, 0x73, 0x33, 0x2e, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x70, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x82, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42,
	0x13, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x41, 0x70, 0x69, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x12, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x50, 0x58,
	0xaa, 0x02, 0x17, 0x43, 0x73, 0x33, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x17, 0x43, 0x73, 0x33,
	0x5c, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x42,
	0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cs3_preferences_v1beta1_preferences_api_proto_rawDescOnce sync.Once
	file_cs3_preferences_v1beta1_preferences_api_proto_rawDescData = file_cs3_preferences_v1beta1_preferences_api_proto_rawDesc
)

func file_cs3_preferences_v1beta1_preferences_api_proto_rawDescGZIP() []byte {
	file_cs3_preferences_v1beta1_preferences_api_proto_rawDescOnce.Do(func() {
		file_cs3_preferences_v1beta1_preferences_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_cs3_preferences_v1beta1_preferences_api_proto_rawDescData)
	})
	return file_cs3_preferences_v1beta1_preferences_api_proto_rawDescData
}

var file_cs3_preferences_v1beta1_preferences_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cs3_preferences_v1beta1_preferences_api_proto_goTypes = []interface{}{
	(*SetKeyRequest)(nil),  // 0: cs3.preferences.v1beta1.SetKeyRequest
	(*SetKeyResponse)(nil), // 1: cs3.preferences.v1beta1.SetKeyResponse
	(*GetKeyRequest)(nil),  // 2: cs3.preferences.v1beta1.GetKeyRequest
	(*GetKeyResponse)(nil), // 3: cs3.preferences.v1beta1.GetKeyResponse
	(*v1beta1.Status)(nil), // 4: cs3.rpc.v1beta1.Status
}
var file_cs3_preferences_v1beta1_preferences_api_proto_depIdxs = []int32{
	4, // 0: cs3.preferences.v1beta1.SetKeyResponse.status:type_name -> cs3.rpc.v1beta1.Status
	4, // 1: cs3.preferences.v1beta1.GetKeyResponse.status:type_name -> cs3.rpc.v1beta1.Status
	0, // 2: cs3.preferences.v1beta1.PreferencesAPI.SetKey:input_type -> cs3.preferences.v1beta1.SetKeyRequest
	2, // 3: cs3.preferences.v1beta1.PreferencesAPI.GetKey:input_type -> cs3.preferences.v1beta1.GetKeyRequest
	1, // 4: cs3.preferences.v1beta1.PreferencesAPI.SetKey:output_type -> cs3.preferences.v1beta1.SetKeyResponse
	3, // 5: cs3.preferences.v1beta1.PreferencesAPI.GetKey:output_type -> cs3.preferences.v1beta1.GetKeyResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cs3_preferences_v1beta1_preferences_api_proto_init() }
func file_cs3_preferences_v1beta1_preferences_api_proto_init() {
	if File_cs3_preferences_v1beta1_preferences_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cs3_preferences_v1beta1_preferences_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetKeyRequest); i {
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
		file_cs3_preferences_v1beta1_preferences_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetKeyResponse); i {
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
		file_cs3_preferences_v1beta1_preferences_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyRequest); i {
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
		file_cs3_preferences_v1beta1_preferences_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyResponse); i {
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
			RawDescriptor: file_cs3_preferences_v1beta1_preferences_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cs3_preferences_v1beta1_preferences_api_proto_goTypes,
		DependencyIndexes: file_cs3_preferences_v1beta1_preferences_api_proto_depIdxs,
		MessageInfos:      file_cs3_preferences_v1beta1_preferences_api_proto_msgTypes,
	}.Build()
	File_cs3_preferences_v1beta1_preferences_api_proto = out.File
	file_cs3_preferences_v1beta1_preferences_api_proto_rawDesc = nil
	file_cs3_preferences_v1beta1_preferences_api_proto_goTypes = nil
	file_cs3_preferences_v1beta1_preferences_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PreferencesAPIClient is the client API for PreferencesAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PreferencesAPIClient interface {
	// Maps the key-value pair.
	SetKey(ctx context.Context, in *SetKeyRequest, opts ...grpc.CallOption) (*SetKeyResponse, error)
	// Returns the value associated with the
	// requested key.
	GetKey(ctx context.Context, in *GetKeyRequest, opts ...grpc.CallOption) (*GetKeyResponse, error)
}

type preferencesAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewPreferencesAPIClient(cc grpc.ClientConnInterface) PreferencesAPIClient {
	return &preferencesAPIClient{cc}
}

func (c *preferencesAPIClient) SetKey(ctx context.Context, in *SetKeyRequest, opts ...grpc.CallOption) (*SetKeyResponse, error) {
	out := new(SetKeyResponse)
	err := c.cc.Invoke(ctx, "/cs3.preferences.v1beta1.PreferencesAPI/SetKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *preferencesAPIClient) GetKey(ctx context.Context, in *GetKeyRequest, opts ...grpc.CallOption) (*GetKeyResponse, error) {
	out := new(GetKeyResponse)
	err := c.cc.Invoke(ctx, "/cs3.preferences.v1beta1.PreferencesAPI/GetKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PreferencesAPIServer is the server API for PreferencesAPI service.
type PreferencesAPIServer interface {
	// Maps the key-value pair.
	SetKey(context.Context, *SetKeyRequest) (*SetKeyResponse, error)
	// Returns the value associated with the
	// requested key.
	GetKey(context.Context, *GetKeyRequest) (*GetKeyResponse, error)
}

// UnimplementedPreferencesAPIServer can be embedded to have forward compatible implementations.
type UnimplementedPreferencesAPIServer struct {
}

func (*UnimplementedPreferencesAPIServer) SetKey(context.Context, *SetKeyRequest) (*SetKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetKey not implemented")
}
func (*UnimplementedPreferencesAPIServer) GetKey(context.Context, *GetKeyRequest) (*GetKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKey not implemented")
}

func RegisterPreferencesAPIServer(s *grpc.Server, srv PreferencesAPIServer) {
	s.RegisterService(&_PreferencesAPI_serviceDesc, srv)
}

func _PreferencesAPI_SetKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PreferencesAPIServer).SetKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.preferences.v1beta1.PreferencesAPI/SetKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PreferencesAPIServer).SetKey(ctx, req.(*SetKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PreferencesAPI_GetKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PreferencesAPIServer).GetKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.preferences.v1beta1.PreferencesAPI/GetKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PreferencesAPIServer).GetKey(ctx, req.(*GetKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PreferencesAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.preferences.v1beta1.PreferencesAPI",
	HandlerType: (*PreferencesAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetKey",
			Handler:    _PreferencesAPI_SetKey_Handler,
		},
		{
			MethodName: "GetKey",
			Handler:    _PreferencesAPI_GetKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/preferences/v1beta1/preferences_api.proto",
}