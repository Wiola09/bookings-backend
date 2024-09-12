// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: rpc_create_restriction.proto

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

type CreateRestrictionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestrictionNameSr string `protobuf:"bytes,1,opt,name=restriction_name_sr,json=restrictionNameSr,proto3" json:"restriction_name_sr,omitempty"`
	RestrictionNameEn string `protobuf:"bytes,2,opt,name=restriction_name_en,json=restrictionNameEn,proto3" json:"restriction_name_en,omitempty"`
	RestrictionNameBg string `protobuf:"bytes,3,opt,name=restriction_name_bg,json=restrictionNameBg,proto3" json:"restriction_name_bg,omitempty"`
}

func (x *CreateRestrictionRequest) Reset() {
	*x = CreateRestrictionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_restriction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRestrictionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRestrictionRequest) ProtoMessage() {}

func (x *CreateRestrictionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_restriction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRestrictionRequest.ProtoReflect.Descriptor instead.
func (*CreateRestrictionRequest) Descriptor() ([]byte, []int) {
	return file_rpc_create_restriction_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRestrictionRequest) GetRestrictionNameSr() string {
	if x != nil {
		return x.RestrictionNameSr
	}
	return ""
}

func (x *CreateRestrictionRequest) GetRestrictionNameEn() string {
	if x != nil {
		return x.RestrictionNameEn
	}
	return ""
}

func (x *CreateRestrictionRequest) GetRestrictionNameBg() string {
	if x != nil {
		return x.RestrictionNameBg
	}
	return ""
}

type CreateRestrictionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Restriction *Restriction `protobuf:"bytes,1,opt,name=restriction,proto3" json:"restriction,omitempty"`
}

func (x *CreateRestrictionResponse) Reset() {
	*x = CreateRestrictionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_restriction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRestrictionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRestrictionResponse) ProtoMessage() {}

func (x *CreateRestrictionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_restriction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRestrictionResponse.ProtoReflect.Descriptor instead.
func (*CreateRestrictionResponse) Descriptor() ([]byte, []int) {
	return file_rpc_create_restriction_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRestrictionResponse) GetRestriction() *Restriction {
	if x != nil {
		return x.Restriction
	}
	return nil
}

var File_rpc_create_restriction_proto protoreflect.FileDescriptor

var file_rpc_create_restriction_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x11, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x53, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x45, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x62, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x42, 0x67, 0x22, 0x4e, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x69, 0x6a, 0x61, 0x6e, 0x61, 0x64, 0x6d, 0x69, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x73, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_create_restriction_proto_rawDescOnce sync.Once
	file_rpc_create_restriction_proto_rawDescData = file_rpc_create_restriction_proto_rawDesc
)

func file_rpc_create_restriction_proto_rawDescGZIP() []byte {
	file_rpc_create_restriction_proto_rawDescOnce.Do(func() {
		file_rpc_create_restriction_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_create_restriction_proto_rawDescData)
	})
	return file_rpc_create_restriction_proto_rawDescData
}

var file_rpc_create_restriction_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_create_restriction_proto_goTypes = []any{
	(*CreateRestrictionRequest)(nil),  // 0: pb.CreateRestrictionRequest
	(*CreateRestrictionResponse)(nil), // 1: pb.CreateRestrictionResponse
	(*Restriction)(nil),               // 2: pb.Restriction
}
var file_rpc_create_restriction_proto_depIdxs = []int32{
	2, // 0: pb.CreateRestrictionResponse.restriction:type_name -> pb.Restriction
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_create_restriction_proto_init() }
func file_rpc_create_restriction_proto_init() {
	if File_rpc_create_restriction_proto != nil {
		return
	}
	file_restriction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_create_restriction_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRestrictionRequest); i {
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
		file_rpc_create_restriction_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRestrictionResponse); i {
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
			RawDescriptor: file_rpc_create_restriction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_create_restriction_proto_goTypes,
		DependencyIndexes: file_rpc_create_restriction_proto_depIdxs,
		MessageInfos:      file_rpc_create_restriction_proto_msgTypes,
	}.Build()
	File_rpc_create_restriction_proto = out.File
	file_rpc_create_restriction_proto_rawDesc = nil
	file_rpc_create_restriction_proto_goTypes = nil
	file_rpc_create_restriction_proto_depIdxs = nil
}
