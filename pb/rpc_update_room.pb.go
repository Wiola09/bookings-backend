// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: rpc_update_room.proto

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

type UpdateRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId             int32   `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RoomNameSr         *string `protobuf:"bytes,2,opt,name=room_name_sr,json=roomNameSr,proto3,oneof" json:"room_name_sr,omitempty"`
	RoomNameEn         *string `protobuf:"bytes,3,opt,name=room_name_en,json=roomNameEn,proto3,oneof" json:"room_name_en,omitempty"`
	RoomNameBg         *string `protobuf:"bytes,4,opt,name=room_name_bg,json=roomNameBg,proto3,oneof" json:"room_name_bg,omitempty"`
	RoomShortdesSr     *string `protobuf:"bytes,5,opt,name=room_shortdes_sr,json=roomShortdesSr,proto3,oneof" json:"room_shortdes_sr,omitempty"`
	RoomShortdesEn     *string `protobuf:"bytes,6,opt,name=room_shortdes_en,json=roomShortdesEn,proto3,oneof" json:"room_shortdes_en,omitempty"`
	RoomShortdescBg    *string `protobuf:"bytes,7,opt,name=room_shortdesc_bg,json=roomShortdescBg,proto3,oneof" json:"room_shortdesc_bg,omitempty"`
	RoomDesSr          *string `protobuf:"bytes,8,opt,name=room_des_sr,json=roomDesSr,proto3,oneof" json:"room_des_sr,omitempty"`
	RoomDesEn          *string `protobuf:"bytes,9,opt,name=room_des_en,json=roomDesEn,proto3,oneof" json:"room_des_en,omitempty"`
	RoomDescBg         *string `protobuf:"bytes,10,opt,name=room_desc_bg,json=roomDescBg,proto3,oneof" json:"room_desc_bg,omitempty"`
	RoomPicturesFolder *string `protobuf:"bytes,11,opt,name=room_pictures_folder,json=roomPicturesFolder,proto3,oneof" json:"room_pictures_folder,omitempty"`
	RoomGuestNumber    *int32  `protobuf:"varint,12,opt,name=room_guest_number,json=roomGuestNumber,proto3,oneof" json:"room_guest_number,omitempty"`
	RoomPriceEn        *int32  `protobuf:"varint,13,opt,name=room_price_en,json=roomPriceEn,proto3,oneof" json:"room_price_en,omitempty"`
}

func (x *UpdateRoomRequest) Reset() {
	*x = UpdateRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_update_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoomRequest) ProtoMessage() {}

func (x *UpdateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_update_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoomRequest.ProtoReflect.Descriptor instead.
func (*UpdateRoomRequest) Descriptor() ([]byte, []int) {
	return file_rpc_update_room_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateRoomRequest) GetRoomId() int32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *UpdateRoomRequest) GetRoomNameSr() string {
	if x != nil && x.RoomNameSr != nil {
		return *x.RoomNameSr
	}
	return ""
}

func (x *UpdateRoomRequest) GetRoomNameEn() string {
	if x != nil && x.RoomNameEn != nil {
		return *x.RoomNameEn
	}
	return ""
}

func (x *UpdateRoomRequest) GetRoomNameBg() string {
	if x != nil && x.RoomNameBg != nil {
		return *x.RoomNameBg
	}
	return ""
}

func (x *UpdateRoomRequest) GetRoomShortdesSr() string {
	if x != nil && x.RoomShortdesSr != nil {
		return *x.RoomShortdesSr
	}
	return ""
}

func (x *UpdateRoomRequest) GetRoomShortdesEn() string {
	if x != nil && x.RoomShortdesEn != nil {
		return *x.RoomShortdesEn
	}
	return ""
}

func (x *UpdateRoomRequest) GetRoomShortdescBg() string {
	if x != nil && x.RoomShortdescBg != nil {
		return *x.RoomShortdescBg
	}
	return ""
}

func (x *UpdateRoomRequest) GetRoomDesSr() string {
	if x != nil && x.RoomDesSr != nil {
		return *x.RoomDesSr
	}
	return ""
}

func (x *UpdateRoomRequest) GetRoomDesEn() string {
	if x != nil && x.RoomDesEn != nil {
		return *x.RoomDesEn
	}
	return ""
}

func (x *UpdateRoomRequest) GetRoomDescBg() string {
	if x != nil && x.RoomDescBg != nil {
		return *x.RoomDescBg
	}
	return ""
}

func (x *UpdateRoomRequest) GetRoomPicturesFolder() string {
	if x != nil && x.RoomPicturesFolder != nil {
		return *x.RoomPicturesFolder
	}
	return ""
}

func (x *UpdateRoomRequest) GetRoomGuestNumber() int32 {
	if x != nil && x.RoomGuestNumber != nil {
		return *x.RoomGuestNumber
	}
	return 0
}

func (x *UpdateRoomRequest) GetRoomPriceEn() int32 {
	if x != nil && x.RoomPriceEn != nil {
		return *x.RoomPriceEn
	}
	return 0
}

type UpdateRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room *Room `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *UpdateRoomResponse) Reset() {
	*x = UpdateRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_update_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoomResponse) ProtoMessage() {}

func (x *UpdateRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_update_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoomResponse.ProtoReflect.Descriptor instead.
func (*UpdateRoomResponse) Descriptor() ([]byte, []int) {
	return file_rpc_update_room_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRoomResponse) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

var File_rpc_update_room_proto protoreflect.FileDescriptor

var file_rpc_update_room_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6f,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0a, 0x72, 0x6f, 0x6f,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x06, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0c, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x73, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a,
	0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x72, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a,
	0x0c, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x45,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0c, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x62, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x72, 0x6f,
	0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x67, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10, 0x72,
	0x6f, 0x6f, 0x6d, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x65, 0x73, 0x5f, 0x73, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0e, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x64, 0x65, 0x73, 0x53, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0e, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x11, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x62, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0f, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x64, 0x65, 0x73, 0x63, 0x42, 0x67, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0b, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x64, 0x65, 0x73, 0x5f, 0x73, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x06, 0x52, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x73, 0x53, 0x72, 0x88, 0x01, 0x01, 0x12,
	0x23, 0x0a, 0x0b, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x64, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x73, 0x45,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0c, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x5f, 0x62, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x0a, 0x72, 0x6f,
	0x6f, 0x6d, 0x44, 0x65, 0x73, 0x63, 0x42, 0x67, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x14, 0x72,
	0x6f, 0x6f, 0x6d, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x5f, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x12, 0x72, 0x6f, 0x6f,
	0x6d, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x88,
	0x01, 0x01, 0x12, 0x2f, 0x0a, 0x11, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x67, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x48, 0x0a, 0x52,
	0x0f, 0x72, 0x6f, 0x6f, 0x6d, 0x47, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0d, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x5f, 0x65, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x48, 0x0b, 0x52, 0x0b, 0x72, 0x6f,
	0x6f, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d,
	0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x72, 0x42, 0x0f, 0x0a,
	0x0d, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x62, 0x67, 0x42,
	0x13, 0x0a, 0x11, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x65,
	0x73, 0x5f, 0x73, 0x72, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x64, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x62, 0x67, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x64, 0x65, 0x73, 0x5f, 0x73, 0x72, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x64, 0x65, 0x73, 0x5f, 0x65, 0x6e, 0x42,
	0x0f, 0x0a, 0x0d, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x62, 0x67,
	0x42, 0x17, 0x0a, 0x15, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x67, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x65,
	0x6e, 0x22, 0x32, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6a, 0x61, 0x6e, 0x61, 0x64, 0x6d, 0x69, 0x2f, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_update_room_proto_rawDescOnce sync.Once
	file_rpc_update_room_proto_rawDescData = file_rpc_update_room_proto_rawDesc
)

func file_rpc_update_room_proto_rawDescGZIP() []byte {
	file_rpc_update_room_proto_rawDescOnce.Do(func() {
		file_rpc_update_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_update_room_proto_rawDescData)
	})
	return file_rpc_update_room_proto_rawDescData
}

var file_rpc_update_room_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_update_room_proto_goTypes = []any{
	(*UpdateRoomRequest)(nil),  // 0: pb.UpdateRoomRequest
	(*UpdateRoomResponse)(nil), // 1: pb.UpdateRoomResponse
	(*Room)(nil),               // 2: pb.Room
}
var file_rpc_update_room_proto_depIdxs = []int32{
	2, // 0: pb.UpdateRoomResponse.room:type_name -> pb.Room
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_update_room_proto_init() }
func file_rpc_update_room_proto_init() {
	if File_rpc_update_room_proto != nil {
		return
	}
	file_room_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_update_room_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRoomRequest); i {
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
		file_rpc_update_room_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRoomResponse); i {
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
	file_rpc_update_room_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_update_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_update_room_proto_goTypes,
		DependencyIndexes: file_rpc_update_room_proto_depIdxs,
		MessageInfos:      file_rpc_update_room_proto_msgTypes,
	}.Build()
	File_rpc_update_room_proto = out.File
	file_rpc_update_room_proto_rawDesc = nil
	file_rpc_update_room_proto_goTypes = nil
	file_rpc_update_room_proto_depIdxs = nil
}
