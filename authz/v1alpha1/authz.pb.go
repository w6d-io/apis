// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.7.1
// source: authz.proto

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

type Provider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IssuerUrl    string `protobuf:"bytes,2,opt,name=issuer_url,json=issuerUrl,proto3" json:"issuer_url,omitempty"`
	ProviderType string `protobuf:"bytes,3,opt,name=provider_type,json=providerType,proto3" json:"provider_type,omitempty"`
}

func (x *Provider) Reset() {
	*x = Provider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provider) ProtoMessage() {}

func (x *Provider) ProtoReflect() protoreflect.Message {
	mi := &file_authz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provider.ProtoReflect.Descriptor instead.
func (*Provider) Descriptor() ([]byte, []int) {
	return file_authz_proto_rawDescGZIP(), []int{0}
}

func (x *Provider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Provider) GetIssuerUrl() string {
	if x != nil {
		return x.IssuerUrl
	}
	return ""
}

func (x *Provider) GetProviderType() string {
	if x != nil {
		return x.ProviderType
	}
	return ""
}

type RoleItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   int64  `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *RoleItem) Reset() {
	*x = RoleItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleItem) ProtoMessage() {}

func (x *RoleItem) ProtoReflect() protoreflect.Message {
	mi := &file_authz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleItem.ProtoReflect.Descriptor instead.
func (*RoleItem) Descriptor() ([]byte, []int) {
	return file_authz_proto_rawDescGZIP(), []int{1}
}

func (x *RoleItem) GetKey() int64 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *RoleItem) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Roles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organizations     []*RoleItem `protobuf:"bytes,1,rep,name=organizations,proto3" json:"organizations,omitempty"`
	Scopes            []*RoleItem `protobuf:"bytes,2,rep,name=scopes,proto3" json:"scopes,omitempty"`
	PrivateProjects   []*RoleItem `protobuf:"bytes,3,rep,name=private_projects,json=privateProjects,proto3" json:"private_projects,omitempty"`
	AffiliateProjects []*RoleItem `protobuf:"bytes,4,rep,name=affiliate_projects,json=affiliateProjects,proto3" json:"affiliate_projects,omitempty"`
}

func (x *Roles) Reset() {
	*x = Roles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Roles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Roles) ProtoMessage() {}

func (x *Roles) ProtoReflect() protoreflect.Message {
	mi := &file_authz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Roles.ProtoReflect.Descriptor instead.
func (*Roles) Descriptor() ([]byte, []int) {
	return file_authz_proto_rawDescGZIP(), []int{2}
}

func (x *Roles) GetOrganizations() []*RoleItem {
	if x != nil {
		return x.Organizations
	}
	return nil
}

func (x *Roles) GetScopes() []*RoleItem {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *Roles) GetPrivateProjects() []*RoleItem {
	if x != nil {
		return x.PrivateProjects
	}
	return nil
}

func (x *Roles) GetAffiliateProjects() []*RoleItem {
	if x != nil {
		return x.AffiliateProjects
	}
	return nil
}

type Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	First string `protobuf:"bytes,1,opt,name=first,proto3" json:"first,omitempty"`
	Last  string `protobuf:"bytes,2,opt,name=last,proto3" json:"last,omitempty"`
}

func (x *Name) Reset() {
	*x = Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name) ProtoMessage() {}

func (x *Name) ProtoReflect() protoreflect.Message {
	mi := &file_authz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name.ProtoReflect.Descriptor instead.
func (*Name) Descriptor() ([]byte, []int) {
	return file_authz_proto_rawDescGZIP(), []int{3}
}

func (x *Name) GetFirst() string {
	if x != nil {
		return x.First
	}
	return ""
}

func (x *Name) GetLast() string {
	if x != nil {
		return x.Last
	}
	return ""
}

type Trait struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      *Name       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email     string      `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Roles     *Roles      `protobuf:"bytes,3,opt,name=roles,proto3" json:"roles,omitempty"`
	Providers []*Provider `protobuf:"bytes,4,rep,name=providers,proto3" json:"providers,omitempty"`
}

func (x *Trait) Reset() {
	*x = Trait{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trait) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trait) ProtoMessage() {}

func (x *Trait) ProtoReflect() protoreflect.Message {
	mi := &file_authz_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trait.ProtoReflect.Descriptor instead.
func (*Trait) Descriptor() ([]byte, []int) {
	return file_authz_proto_rawDescGZIP(), []int{4}
}

func (x *Trait) GetName() *Name {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Trait) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Trait) GetRoles() *Roles {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *Trait) GetProviders() []*Provider {
	if x != nil {
		return x.Providers
	}
	return nil
}

var File_authz_proto protoreflect.FileDescriptor

var file_authz_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x32, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xcb, 0x01, 0x0a, 0x05, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12,
	0x2f, 0x0a, 0x0d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x0d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x21, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x12, 0x61, 0x66, 0x66,
	0x69, 0x6c, 0x69, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x11, 0x61, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x22, 0x30, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6c, 0x61, 0x73, 0x74, 0x22, 0x7f, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x69, 0x74, 0x12, 0x19,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1c, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x27, 0x0a,
	0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x36, 0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authz_proto_rawDescOnce sync.Once
	file_authz_proto_rawDescData = file_authz_proto_rawDesc
)

func file_authz_proto_rawDescGZIP() []byte {
	file_authz_proto_rawDescOnce.Do(func() {
		file_authz_proto_rawDescData = protoimpl.X.CompressGZIP(file_authz_proto_rawDescData)
	})
	return file_authz_proto_rawDescData
}

var file_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_authz_proto_goTypes = []interface{}{
	(*Provider)(nil), // 0: Provider
	(*RoleItem)(nil), // 1: RoleItem
	(*Roles)(nil),    // 2: Roles
	(*Name)(nil),     // 3: Name
	(*Trait)(nil),    // 4: Trait
}
var file_authz_proto_depIdxs = []int32{
	1, // 0: Roles.organizations:type_name -> RoleItem
	1, // 1: Roles.scopes:type_name -> RoleItem
	1, // 2: Roles.private_projects:type_name -> RoleItem
	1, // 3: Roles.affiliate_projects:type_name -> RoleItem
	3, // 4: Trait.name:type_name -> Name
	2, // 5: Trait.roles:type_name -> Roles
	0, // 6: Trait.providers:type_name -> Provider
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_authz_proto_init() }
func file_authz_proto_init() {
	if File_authz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provider); i {
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
		file_authz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleItem); i {
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
		file_authz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Roles); i {
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
		file_authz_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Name); i {
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
		file_authz_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trait); i {
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
			RawDescriptor: file_authz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_authz_proto_goTypes,
		DependencyIndexes: file_authz_proto_depIdxs,
		MessageInfos:      file_authz_proto_msgTypes,
	}.Build()
	File_authz_proto = out.File
	file_authz_proto_rawDesc = nil
	file_authz_proto_goTypes = nil
	file_authz_proto_depIdxs = nil
}
