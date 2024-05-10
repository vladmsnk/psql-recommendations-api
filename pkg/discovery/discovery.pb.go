// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: discovery/discovery.proto

package pb

import (
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

type RegisterInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	Config       []byte `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *RegisterInstanceRequest) Reset() {
	*x = RegisterInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_discovery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterInstanceRequest) ProtoMessage() {}

func (x *RegisterInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_discovery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterInstanceRequest.ProtoReflect.Descriptor instead.
func (*RegisterInstanceRequest) Descriptor() ([]byte, []int) {
	return file_discovery_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterInstanceRequest) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *RegisterInstanceRequest) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

type RegisterInstanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerId  string `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	InstanceName string `protobuf:"bytes,2,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	Host         string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Port         int64  `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *RegisterInstanceResponse) Reset() {
	*x = RegisterInstanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_discovery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterInstanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterInstanceResponse) ProtoMessage() {}

func (x *RegisterInstanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_discovery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterInstanceResponse.ProtoReflect.Descriptor instead.
func (*RegisterInstanceResponse) Descriptor() ([]byte, []int) {
	return file_discovery_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterInstanceResponse) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

func (x *RegisterInstanceResponse) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *RegisterInstanceResponse) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RegisterInstanceResponse) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_discovery_discovery_proto protoreflect.FileDescriptor

var file_discovery_discovery_proto_rawDesc = []byte{
	0x0a, 0x19, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x22, 0x56, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x8a,
	0x01, 0x0a, 0x18, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x68, 0x0a, 0x09, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x5b, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x22, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_discovery_discovery_proto_rawDescOnce sync.Once
	file_discovery_discovery_proto_rawDescData = file_discovery_discovery_proto_rawDesc
)

func file_discovery_discovery_proto_rawDescGZIP() []byte {
	file_discovery_discovery_proto_rawDescOnce.Do(func() {
		file_discovery_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_discovery_discovery_proto_rawDescData)
	})
	return file_discovery_discovery_proto_rawDescData
}

var file_discovery_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_discovery_discovery_proto_goTypes = []interface{}{
	(*RegisterInstanceRequest)(nil),  // 0: discovery.RegisterInstanceRequest
	(*RegisterInstanceResponse)(nil), // 1: discovery.RegisterInstanceResponse
}
var file_discovery_discovery_proto_depIdxs = []int32{
	0, // 0: discovery.Discovery.RegisterInstance:input_type -> discovery.RegisterInstanceRequest
	1, // 1: discovery.Discovery.RegisterInstance:output_type -> discovery.RegisterInstanceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_discovery_discovery_proto_init() }
func file_discovery_discovery_proto_init() {
	if File_discovery_discovery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_discovery_discovery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterInstanceRequest); i {
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
		file_discovery_discovery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterInstanceResponse); i {
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
			RawDescriptor: file_discovery_discovery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_discovery_discovery_proto_goTypes,
		DependencyIndexes: file_discovery_discovery_proto_depIdxs,
		MessageInfos:      file_discovery_discovery_proto_msgTypes,
	}.Build()
	File_discovery_discovery_proto = out.File
	file_discovery_discovery_proto_rawDesc = nil
	file_discovery_discovery_proto_goTypes = nil
	file_discovery_discovery_proto_depIdxs = nil
}
