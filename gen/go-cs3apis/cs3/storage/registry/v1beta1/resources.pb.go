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
// source: cs3/storage/registry/v1beta1/resources.proto

package registryv1beta1

import (
	v1beta1 "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/types/v1beta1"
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

// The information for the storage provider.
type ProviderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OPTIONAL.
	// Opaque information (containing storage-specific information).
	// For example, additional metadata attached to the resource.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The storage provider id that will become part of the
	// resource id.
	// For example, if the provider_id is "home", resources obtained
	// from this storage provider will have a resource id like "home:1234".
	ProviderId string `protobuf:"bytes,2,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	// REQUIRED.
	// The public path prefix this storage provider handles.
	// In Unix literature, the mount path.
	// For example, if the provider_path is "/home", resources obtained
	// from this storage provier will have a resource path lik "/home/docs".
	ProviderPath string `protobuf:"bytes,3,opt,name=provider_path,json=providerPath,proto3" json:"provider_path,omitempty"`
	// REQUIRED.
	// The address where the storage provider can be reached.
	// For example, tcp://localhost:1099.
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	// OPTIONAL.
	// Information to describe the functionalities
	// offered by the storage provider. Meant to be read
	// by humans.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// REQUIRED.
	// List of available methods.
	Features *ProviderInfo_Features `protobuf:"bytes,6,opt,name=features,proto3" json:"features,omitempty"`
}

func (x *ProviderInfo) Reset() {
	*x = ProviderInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_storage_registry_v1beta1_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderInfo) ProtoMessage() {}

func (x *ProviderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_storage_registry_v1beta1_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderInfo.ProtoReflect.Descriptor instead.
func (*ProviderInfo) Descriptor() ([]byte, []int) {
	return file_cs3_storage_registry_v1beta1_resources_proto_rawDescGZIP(), []int{0}
}

func (x *ProviderInfo) GetOpaque() *v1beta1.Opaque {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *ProviderInfo) GetProviderId() string {
	if x != nil {
		return x.ProviderId
	}
	return ""
}

func (x *ProviderInfo) GetProviderPath() string {
	if x != nil {
		return x.ProviderPath
	}
	return ""
}

func (x *ProviderInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ProviderInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProviderInfo) GetFeatures() *ProviderInfo_Features {
	if x != nil {
		return x.Features
	}
	return nil
}

// REQUIRED.
// Represents the list of features available
// on this storage provider. If the feature is not supported,
// the related service methods MUST return CODE_UNIMPLEMENTED.
type ProviderInfo_Features struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recycle      bool `protobuf:"varint,1,opt,name=recycle,proto3" json:"recycle,omitempty"`
	FileVersions bool `protobuf:"varint,2,opt,name=file_versions,json=fileVersions,proto3" json:"file_versions,omitempty"`
}

func (x *ProviderInfo_Features) Reset() {
	*x = ProviderInfo_Features{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs3_storage_registry_v1beta1_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderInfo_Features) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderInfo_Features) ProtoMessage() {}

func (x *ProviderInfo_Features) ProtoReflect() protoreflect.Message {
	mi := &file_cs3_storage_registry_v1beta1_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderInfo_Features.ProtoReflect.Descriptor instead.
func (*ProviderInfo_Features) Descriptor() ([]byte, []int) {
	return file_cs3_storage_registry_v1beta1_resources_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProviderInfo_Features) GetRecycle() bool {
	if x != nil {
		return x.Recycle
	}
	return false
}

func (x *ProviderInfo_Features) GetFileVersions() bool {
	if x != nil {
		return x.FileVersions
	}
	return false
}

var File_cs3_storage_registry_v1beta1_resources_proto protoreflect.FileDescriptor

var file_cs3_storage_registry_v1beta1_resources_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x63, 0x73, 0x33, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c,
	0x63, 0x73, 0x33, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1d, 0x63, 0x73,
	0x33, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x02, 0x0a, 0x0c,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x06,
	0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63,
	0x73, 0x33, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x4f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x4f, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x1a, 0x49, 0x0a, 0x08, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x72, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x66, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x89, 0x01,
	0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x73, 0x33, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x42, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x52, 0xaa, 0x02, 0x1c, 0x43, 0x73,
	0x33, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1c, 0x43, 0x73, 0x33,
	0x5c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x5c, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cs3_storage_registry_v1beta1_resources_proto_rawDescOnce sync.Once
	file_cs3_storage_registry_v1beta1_resources_proto_rawDescData = file_cs3_storage_registry_v1beta1_resources_proto_rawDesc
)

func file_cs3_storage_registry_v1beta1_resources_proto_rawDescGZIP() []byte {
	file_cs3_storage_registry_v1beta1_resources_proto_rawDescOnce.Do(func() {
		file_cs3_storage_registry_v1beta1_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_cs3_storage_registry_v1beta1_resources_proto_rawDescData)
	})
	return file_cs3_storage_registry_v1beta1_resources_proto_rawDescData
}

var file_cs3_storage_registry_v1beta1_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cs3_storage_registry_v1beta1_resources_proto_goTypes = []interface{}{
	(*ProviderInfo)(nil),          // 0: cs3.storage.registry.v1beta1.ProviderInfo
	(*ProviderInfo_Features)(nil), // 1: cs3.storage.registry.v1beta1.ProviderInfo.Features
	(*v1beta1.Opaque)(nil),        // 2: cs3.types.v1beta1.Opaque
}
var file_cs3_storage_registry_v1beta1_resources_proto_depIdxs = []int32{
	2, // 0: cs3.storage.registry.v1beta1.ProviderInfo.opaque:type_name -> cs3.types.v1beta1.Opaque
	1, // 1: cs3.storage.registry.v1beta1.ProviderInfo.features:type_name -> cs3.storage.registry.v1beta1.ProviderInfo.Features
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cs3_storage_registry_v1beta1_resources_proto_init() }
func file_cs3_storage_registry_v1beta1_resources_proto_init() {
	if File_cs3_storage_registry_v1beta1_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cs3_storage_registry_v1beta1_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderInfo); i {
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
		file_cs3_storage_registry_v1beta1_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderInfo_Features); i {
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
			RawDescriptor: file_cs3_storage_registry_v1beta1_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cs3_storage_registry_v1beta1_resources_proto_goTypes,
		DependencyIndexes: file_cs3_storage_registry_v1beta1_resources_proto_depIdxs,
		MessageInfos:      file_cs3_storage_registry_v1beta1_resources_proto_msgTypes,
	}.Build()
	File_cs3_storage_registry_v1beta1_resources_proto = out.File
	file_cs3_storage_registry_v1beta1_resources_proto_rawDesc = nil
	file_cs3_storage_registry_v1beta1_resources_proto_goTypes = nil
	file_cs3_storage_registry_v1beta1_resources_proto_depIdxs = nil
}
