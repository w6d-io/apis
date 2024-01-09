// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.7.1
// source: iam.proto

package v1alpha1

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

type Mode int32

const (
	Mode_Admin  Mode = 0
	Mode_Public Mode = 1
	Mode_Trait  Mode = 2
)

// Enum value maps for Mode.
var (
	Mode_name = map[int32]string{
		0: "Admin",
		1: "Public",
		2: "Trait",
	}
	Mode_value = map[string]int32{
		"Admin":  0,
		"Public": 1,
		"Trait":  2,
	}
)

func (x Mode) Enum() *Mode {
	p := new(Mode)
	*p = x
	return p
}

func (x Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_iam_proto_enumTypes[0].Descriptor()
}

func (Mode) Type() protoreflect.EnumType {
	return &file_iam_proto_enumTypes[0]
}

func (x Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Mode.Descriptor instead.
func (Mode) EnumDescriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{0}
}

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PermType string `protobuf:"bytes,2,opt,name=perm_type,json=permType,proto3" json:"perm_type,omitempty"`
	Resource string `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	Value    string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Mode     Mode   `protobuf:"varint,5,opt,name=mode,proto3,enum=permission.Mode" json:"mode,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Input) GetPermType() string {
	if x != nil {
		return x.PermType
	}
	return ""
}

func (x *Input) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *Input) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Input) GetMode() Mode {
	if x != nil {
		return x.Mode
	}
	return Mode_Admin
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iam_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_iam_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_iam_proto_rawDescGZIP(), []int{1}
}

var File_iam_proto protoreflect.FileDescriptor

var file_iam_proto_rawDesc = []byte{
	0x0a, 0x09, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x8c, 0x01, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x65, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x24, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2a,
	0x28, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x54, 0x72, 0x61, 0x69, 0x74, 0x10, 0x02, 0x32, 0xb7, 0x01, 0x0a, 0x03, 0x49, 0x61,
	0x6d, 0x12, 0x37, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x10, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x11,
	0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x11, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x11,
	0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x77, 0x36, 0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_iam_proto_rawDescOnce sync.Once
	file_iam_proto_rawDescData = file_iam_proto_rawDesc
)

func file_iam_proto_rawDescGZIP() []byte {
	file_iam_proto_rawDescOnce.Do(func() {
		file_iam_proto_rawDescData = protoimpl.X.CompressGZIP(file_iam_proto_rawDescData)
	})
	return file_iam_proto_rawDescData
}

var file_iam_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_iam_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_iam_proto_goTypes = []interface{}{
	(Mode)(0),     // 0: permission.Mode
	(*Input)(nil), // 1: permission.Input
	(*Reply)(nil), // 2: permission.Reply
}
var file_iam_proto_depIdxs = []int32{
	0, // 0: permission.Input.mode:type_name -> permission.Mode
	1, // 1: permission.Iam.AddPermission:input_type -> permission.Input
	1, // 2: permission.Iam.RemovePermission:input_type -> permission.Input
	1, // 3: permission.Iam.ReplacePermission:input_type -> permission.Input
	2, // 4: permission.Iam.AddPermission:output_type -> permission.Reply
	2, // 5: permission.Iam.RemovePermission:output_type -> permission.Reply
	2, // 6: permission.Iam.ReplacePermission:output_type -> permission.Reply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_iam_proto_init() }
func file_iam_proto_init() {
	if File_iam_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iam_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
		file_iam_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
			RawDescriptor: file_iam_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iam_proto_goTypes,
		DependencyIndexes: file_iam_proto_depIdxs,
		EnumInfos:         file_iam_proto_enumTypes,
		MessageInfos:      file_iam_proto_msgTypes,
	}.Build()
	File_iam_proto = out.File
	file_iam_proto_rawDesc = nil
	file_iam_proto_goTypes = nil
	file_iam_proto_depIdxs = nil
}
