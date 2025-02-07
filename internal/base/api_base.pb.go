// Built-in base types for API calls. Primarily useful as return types.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: api_base.proto

package base

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

type StringProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *StringProto) Reset() {
	*x = StringProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringProto) ProtoMessage() {}

func (x *StringProto) ProtoReflect() protoreflect.Message {
	mi := &file_api_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringProto.ProtoReflect.Descriptor instead.
func (*StringProto) Descriptor() ([]byte, []int) {
	return file_api_base_proto_rawDescGZIP(), []int{0}
}

func (x *StringProto) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Integer32Proto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Integer32Proto) Reset() {
	*x = Integer32Proto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Integer32Proto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Integer32Proto) ProtoMessage() {}

func (x *Integer32Proto) ProtoReflect() protoreflect.Message {
	mi := &file_api_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Integer32Proto.ProtoReflect.Descriptor instead.
func (*Integer32Proto) Descriptor() ([]byte, []int) {
	return file_api_base_proto_rawDescGZIP(), []int{1}
}

func (x *Integer32Proto) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Integer64Proto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Integer64Proto) Reset() {
	*x = Integer64Proto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Integer64Proto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Integer64Proto) ProtoMessage() {}

func (x *Integer64Proto) ProtoReflect() protoreflect.Message {
	mi := &file_api_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Integer64Proto.ProtoReflect.Descriptor instead.
func (*Integer64Proto) Descriptor() ([]byte, []int) {
	return file_api_base_proto_rawDescGZIP(), []int{2}
}

func (x *Integer64Proto) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type BoolProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BoolProto) Reset() {
	*x = BoolProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolProto) ProtoMessage() {}

func (x *BoolProto) ProtoReflect() protoreflect.Message {
	mi := &file_api_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolProto.ProtoReflect.Descriptor instead.
func (*BoolProto) Descriptor() ([]byte, []int) {
	return file_api_base_proto_rawDescGZIP(), []int{3}
}

func (x *BoolProto) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type DoubleProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DoubleProto) Reset() {
	*x = DoubleProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_base_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleProto) ProtoMessage() {}

func (x *DoubleProto) ProtoReflect() protoreflect.Message {
	mi := &file_api_base_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleProto.ProtoReflect.Descriptor instead.
func (*DoubleProto) Descriptor() ([]byte, []int) {
	return file_api_base_proto_rawDescGZIP(), []int{4}
}

func (x *DoubleProto) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type BytesProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BytesProto) Reset() {
	*x = BytesProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_base_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesProto) ProtoMessage() {}

func (x *BytesProto) ProtoReflect() protoreflect.Message {
	mi := &file_api_base_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesProto.ProtoReflect.Descriptor instead.
func (*BytesProto) Descriptor() ([]byte, []int) {
	return file_api_base_proto_rawDescGZIP(), []int{5}
}

func (x *BytesProto) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type VoidProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VoidProto) Reset() {
	*x = VoidProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_base_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoidProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoidProto) ProtoMessage() {}

func (x *VoidProto) ProtoReflect() protoreflect.Message {
	mi := &file_api_base_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoidProto.ProtoReflect.Descriptor instead.
func (*VoidProto) Descriptor() ([]byte, []int) {
	return file_api_base_proto_rawDescGZIP(), []int{6}
}

var File_api_base_proto protoreflect.FileDescriptor

var file_api_base_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x22, 0x23, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x26, 0x0a, 0x0e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72,
	0x33, 0x32, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x26, 0x0a,
	0x0e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x36, 0x34, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x21, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x0a, 0x0b, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x26, 0x0a,
	0x0a, 0x42, 0x79, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x02, 0x08, 0x01, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x0b, 0x0a, 0x09, 0x56, 0x6f, 0x69, 0x64, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_base_proto_rawDescOnce sync.Once
	file_api_base_proto_rawDescData = file_api_base_proto_rawDesc
)

func file_api_base_proto_rawDescGZIP() []byte {
	file_api_base_proto_rawDescOnce.Do(func() {
		file_api_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_base_proto_rawDescData)
	})
	return file_api_base_proto_rawDescData
}

var file_api_base_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_base_proto_goTypes = []interface{}{
	(*StringProto)(nil),    // 0: appengine.base.StringProto
	(*Integer32Proto)(nil), // 1: appengine.base.Integer32Proto
	(*Integer64Proto)(nil), // 2: appengine.base.Integer64Proto
	(*BoolProto)(nil),      // 3: appengine.base.BoolProto
	(*DoubleProto)(nil),    // 4: appengine.base.DoubleProto
	(*BytesProto)(nil),     // 5: appengine.base.BytesProto
	(*VoidProto)(nil),      // 6: appengine.base.VoidProto
}
var file_api_base_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_base_proto_init() }
func file_api_base_proto_init() {
	if File_api_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringProto); i {
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
		file_api_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Integer32Proto); i {
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
		file_api_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Integer64Proto); i {
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
		file_api_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolProto); i {
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
		file_api_base_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleProto); i {
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
		file_api_base_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BytesProto); i {
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
		file_api_base_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoidProto); i {
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
			RawDescriptor: file_api_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_base_proto_goTypes,
		DependencyIndexes: file_api_base_proto_depIdxs,
		MessageInfos:      file_api_base_proto_msgTypes,
	}.Build()
	File_api_base_proto = out.File
	file_api_base_proto_rawDesc = nil
	file_api_base_proto_goTypes = nil
	file_api_base_proto_depIdxs = nil
}
