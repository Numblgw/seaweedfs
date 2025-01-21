// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.3
// source: s3.proto

package s3_pb

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

type S3ConfigureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S3ConfigurationFileContent []byte `protobuf:"bytes,1,opt,name=s3_configuration_file_content,json=s3ConfigurationFileContent,proto3" json:"s3_configuration_file_content,omitempty"`
}

func (x *S3ConfigureRequest) Reset() {
	*x = S3ConfigureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3ConfigureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3ConfigureRequest) ProtoMessage() {}

func (x *S3ConfigureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_s3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3ConfigureRequest.ProtoReflect.Descriptor instead.
func (*S3ConfigureRequest) Descriptor() ([]byte, []int) {
	return file_s3_proto_rawDescGZIP(), []int{0}
}

func (x *S3ConfigureRequest) GetS3ConfigurationFileContent() []byte {
	if x != nil {
		return x.S3ConfigurationFileContent
	}
	return nil
}

type S3ConfigureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *S3ConfigureResponse) Reset() {
	*x = S3ConfigureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3ConfigureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3ConfigureResponse) ProtoMessage() {}

func (x *S3ConfigureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_s3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3ConfigureResponse.ProtoReflect.Descriptor instead.
func (*S3ConfigureResponse) Descriptor() ([]byte, []int) {
	return file_s3_proto_rawDescGZIP(), []int{1}
}

type S3CircuitBreakerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Global  *S3CircuitBreakerOptions            `protobuf:"bytes,1,opt,name=global,proto3" json:"global,omitempty"`
	Buckets map[string]*S3CircuitBreakerOptions `protobuf:"bytes,2,rep,name=buckets,proto3" json:"buckets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *S3CircuitBreakerConfig) Reset() {
	*x = S3CircuitBreakerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3CircuitBreakerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3CircuitBreakerConfig) ProtoMessage() {}

func (x *S3CircuitBreakerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_s3_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3CircuitBreakerConfig.ProtoReflect.Descriptor instead.
func (*S3CircuitBreakerConfig) Descriptor() ([]byte, []int) {
	return file_s3_proto_rawDescGZIP(), []int{2}
}

func (x *S3CircuitBreakerConfig) GetGlobal() *S3CircuitBreakerOptions {
	if x != nil {
		return x.Global
	}
	return nil
}

func (x *S3CircuitBreakerConfig) GetBuckets() map[string]*S3CircuitBreakerOptions {
	if x != nil {
		return x.Buckets
	}
	return nil
}

type S3CircuitBreakerOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool             `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Actions map[string]int64 `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *S3CircuitBreakerOptions) Reset() {
	*x = S3CircuitBreakerOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3CircuitBreakerOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3CircuitBreakerOptions) ProtoMessage() {}

func (x *S3CircuitBreakerOptions) ProtoReflect() protoreflect.Message {
	mi := &file_s3_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3CircuitBreakerOptions.ProtoReflect.Descriptor instead.
func (*S3CircuitBreakerOptions) Descriptor() ([]byte, []int) {
	return file_s3_proto_rawDescGZIP(), []int{3}
}

func (x *S3CircuitBreakerOptions) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *S3CircuitBreakerOptions) GetActions() map[string]int64 {
	if x != nil {
		return x.Actions
	}
	return nil
}

var File_s3_proto protoreflect.FileDescriptor

var file_s3_proto_rawDesc = []byte{
	0x0a, 0x08, 0x73, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x22, 0x57, 0x0a, 0x12, 0x53, 0x33, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41,
	0x0a, 0x1d, 0x73, 0x33, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x1a, 0x73, 0x33, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x33, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x87, 0x02, 0x0a, 0x16, 0x53, 0x33, 0x43,
	0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x3d, 0x0a, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x62, 0x2e, 0x53, 0x33, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65, 0x61,
	0x6b, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x06, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x12, 0x4b, 0x0a, 0x07, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x62, 0x2e, 0x53, 0x33, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65, 0x61,
	0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x1a,
	0x61, 0x0a, 0x0c, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x3b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2e,
	0x53, 0x33, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xbd, 0x01, 0x0a, 0x17, 0x53, 0x33, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74,
	0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x4c, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x33, 0x43, 0x69, 0x72, 0x63, 0x75,
	0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x32, 0x5f, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x77, 0x65, 0x65, 0x64, 0x53, 0x33, 0x12,
	0x52, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x12, 0x20, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x33, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x33,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x49, 0x0a, 0x10, 0x73, 0x65, 0x61, 0x77, 0x65, 0x65, 0x64, 0x66, 0x73,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x07, 0x53, 0x33, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x61,
	0x77, 0x65, 0x65, 0x64, 0x66, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x77, 0x65, 0x65, 0x64, 0x66, 0x73,
	0x2f, 0x77, 0x65, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x33, 0x5f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_s3_proto_rawDescOnce sync.Once
	file_s3_proto_rawDescData = file_s3_proto_rawDesc
)

func file_s3_proto_rawDescGZIP() []byte {
	file_s3_proto_rawDescOnce.Do(func() {
		file_s3_proto_rawDescData = protoimpl.X.CompressGZIP(file_s3_proto_rawDescData)
	})
	return file_s3_proto_rawDescData
}

var file_s3_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_s3_proto_goTypes = []any{
	(*S3ConfigureRequest)(nil),      // 0: messaging_pb.S3ConfigureRequest
	(*S3ConfigureResponse)(nil),     // 1: messaging_pb.S3ConfigureResponse
	(*S3CircuitBreakerConfig)(nil),  // 2: messaging_pb.S3CircuitBreakerConfig
	(*S3CircuitBreakerOptions)(nil), // 3: messaging_pb.S3CircuitBreakerOptions
	nil,                             // 4: messaging_pb.S3CircuitBreakerConfig.BucketsEntry
	nil,                             // 5: messaging_pb.S3CircuitBreakerOptions.ActionsEntry
}
var file_s3_proto_depIdxs = []int32{
	3, // 0: messaging_pb.S3CircuitBreakerConfig.global:type_name -> messaging_pb.S3CircuitBreakerOptions
	4, // 1: messaging_pb.S3CircuitBreakerConfig.buckets:type_name -> messaging_pb.S3CircuitBreakerConfig.BucketsEntry
	5, // 2: messaging_pb.S3CircuitBreakerOptions.actions:type_name -> messaging_pb.S3CircuitBreakerOptions.ActionsEntry
	3, // 3: messaging_pb.S3CircuitBreakerConfig.BucketsEntry.value:type_name -> messaging_pb.S3CircuitBreakerOptions
	0, // 4: messaging_pb.SeaweedS3.Configure:input_type -> messaging_pb.S3ConfigureRequest
	1, // 5: messaging_pb.SeaweedS3.Configure:output_type -> messaging_pb.S3ConfigureResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_s3_proto_init() }
func file_s3_proto_init() {
	if File_s3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_s3_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*S3ConfigureRequest); i {
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
		file_s3_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*S3ConfigureResponse); i {
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
		file_s3_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*S3CircuitBreakerConfig); i {
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
		file_s3_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*S3CircuitBreakerOptions); i {
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
			RawDescriptor: file_s3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_s3_proto_goTypes,
		DependencyIndexes: file_s3_proto_depIdxs,
		MessageInfos:      file_s3_proto_msgTypes,
	}.Build()
	File_s3_proto = out.File
	file_s3_proto_rawDesc = nil
	file_s3_proto_goTypes = nil
	file_s3_proto_depIdxs = nil
}
