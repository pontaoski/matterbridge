// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: chat/v1/emotes.proto

package v1

import (
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

type CreateEmotePackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackName string `protobuf:"bytes,1,opt,name=pack_name,json=packName,proto3" json:"pack_name,omitempty"`
}

func (x *CreateEmotePackRequest) Reset() {
	*x = CreateEmotePackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_emotes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmotePackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmotePackRequest) ProtoMessage() {}

func (x *CreateEmotePackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_emotes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmotePackRequest.ProtoReflect.Descriptor instead.
func (*CreateEmotePackRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_emotes_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEmotePackRequest) GetPackName() string {
	if x != nil {
		return x.PackName
	}
	return ""
}

type CreateEmotePackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackId uint64 `protobuf:"varint,1,opt,name=pack_id,json=packId,proto3" json:"pack_id,omitempty"`
}

func (x *CreateEmotePackResponse) Reset() {
	*x = CreateEmotePackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_emotes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmotePackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmotePackResponse) ProtoMessage() {}

func (x *CreateEmotePackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_emotes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmotePackResponse.ProtoReflect.Descriptor instead.
func (*CreateEmotePackResponse) Descriptor() ([]byte, []int) {
	return file_chat_v1_emotes_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEmotePackResponse) GetPackId() uint64 {
	if x != nil {
		return x.PackId
	}
	return 0
}

type GetEmotePacksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEmotePacksRequest) Reset() {
	*x = GetEmotePacksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_emotes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmotePacksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmotePacksRequest) ProtoMessage() {}

func (x *GetEmotePacksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_emotes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmotePacksRequest.ProtoReflect.Descriptor instead.
func (*GetEmotePacksRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_emotes_proto_rawDescGZIP(), []int{2}
}

type GetEmotePacksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Packs []*GetEmotePacksResponse_EmotePack `protobuf:"bytes,1,rep,name=packs,proto3" json:"packs,omitempty"`
}

func (x *GetEmotePacksResponse) Reset() {
	*x = GetEmotePacksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_emotes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmotePacksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmotePacksResponse) ProtoMessage() {}

func (x *GetEmotePacksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_emotes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmotePacksResponse.ProtoReflect.Descriptor instead.
func (*GetEmotePacksResponse) Descriptor() ([]byte, []int) {
	return file_chat_v1_emotes_proto_rawDescGZIP(), []int{3}
}

func (x *GetEmotePacksResponse) GetPacks() []*GetEmotePacksResponse_EmotePack {
	if x != nil {
		return x.Packs
	}
	return nil
}

type GetEmotePackEmotesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackId uint64 `protobuf:"varint,1,opt,name=pack_id,json=packId,proto3" json:"pack_id,omitempty"`
}

func (x *GetEmotePackEmotesRequest) Reset() {
	*x = GetEmotePackEmotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_emotes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmotePackEmotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmotePackEmotesRequest) ProtoMessage() {}

func (x *GetEmotePackEmotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_emotes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmotePackEmotesRequest.ProtoReflect.Descriptor instead.
func (*GetEmotePackEmotesRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_emotes_proto_rawDescGZIP(), []int{4}
}

func (x *GetEmotePackEmotesRequest) GetPackId() uint64 {
	if x != nil {
		return x.PackId
	}
	return 0
}

type GetEmotePackEmotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emotes []*GetEmotePackEmotesResponse_Emote `protobuf:"bytes,1,rep,name=emotes,proto3" json:"emotes,omitempty"`
}

func (x *GetEmotePackEmotesResponse) Reset() {
	*x = GetEmotePackEmotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_emotes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmotePackEmotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmotePackEmotesResponse) ProtoMessage() {}

func (x *GetEmotePackEmotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_emotes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmotePackEmotesResponse.ProtoReflect.Descriptor instead.
func (*GetEmotePackEmotesResponse) Descriptor() ([]byte, []int) {
	return file_chat_v1_emotes_proto_rawDescGZIP(), []int{5}
}

func (x *GetEmotePackEmotesResponse) GetEmotes() []*GetEmotePackEmotesResponse_Emote {
	if x != nil {
		return x.Emotes
	}
	return nil
}

type AddEmoteToPackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackId  uint64 `protobuf:"varint,1,opt,name=pack_id,json=packId,proto3" json:"pack_id,omitempty"`
	ImageId string `protobuf:"bytes,2,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AddEmoteToPackRequest) Reset() {
	*x = AddEmoteToPackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_emotes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEmoteToPackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEmoteToPackRequest) ProtoMessage() {}

func (x *AddEmoteToPackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_emotes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEmoteToPackRequest.ProtoReflect.Descriptor instead.
func (*AddEmoteToPackRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_emotes_proto_rawDescGZIP(), []int{6}
}

func (x *AddEmoteToPackRequest) GetPackId() uint64 {
	if x != nil {
		return x.PackId
	}
	return 0
}

func (x *AddEmoteToPackRequest) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *AddEmoteToPackRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteEmoteFromPackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackId  uint64 `protobuf:"varint,1,opt,name=pack_id,json=packId,proto3" json:"pack_id,omitempty"`
	ImageId string `protobuf:"bytes,2,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
}

func (x *DeleteEmoteFromPackRequest) Reset() {
	*x = DeleteEmoteFromPackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_emotes_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEmoteFromPackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEmoteFromPackRequest) ProtoMessage() {}

func (x *DeleteEmoteFromPackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_emotes_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEmoteFromPackRequest.ProtoReflect.Descriptor instead.
func (*DeleteEmoteFromPackRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_emotes_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteEmoteFromPackRequest) GetPackId() uint64 {
	if x != nil {
		return x.PackId
	}
	return 0
}

func (x *DeleteEmoteFromPackRequest) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

type DeleteEmotePackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackId uint64 `protobuf:"varint,1,opt,name=pack_id,json=packId,proto3" json:"pack_id,omitempty"`
}

func (x *DeleteEmotePackRequest) Reset() {
	*x = DeleteEmotePackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_emotes_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEmotePackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEmotePackRequest) ProtoMessage() {}

func (x *DeleteEmotePackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_emotes_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEmotePackRequest.ProtoReflect.Descriptor instead.
func (*DeleteEmotePackRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_emotes_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteEmotePackRequest) GetPackId() uint64 {
	if x != nil {
		return x.PackId
	}
	return 0
}

type DequipEmotePackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackId uint64 `protobuf:"varint,1,opt,name=pack_id,json=packId,proto3" json:"pack_id,omitempty"`
}

func (x *DequipEmotePackRequest) Reset() {
	*x = DequipEmotePackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_emotes_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DequipEmotePackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DequipEmotePackRequest) ProtoMessage() {}

func (x *DequipEmotePackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_emotes_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DequipEmotePackRequest.ProtoReflect.Descriptor instead.
func (*DequipEmotePackRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_emotes_proto_rawDescGZIP(), []int{9}
}

func (x *DequipEmotePackRequest) GetPackId() uint64 {
	if x != nil {
		return x.PackId
	}
	return 0
}

type GetEmotePacksResponse_EmotePack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackId    uint64 `protobuf:"varint,1,opt,name=pack_id,json=packId,proto3" json:"pack_id,omitempty"`
	PackOwner uint64 `protobuf:"varint,2,opt,name=pack_owner,json=packOwner,proto3" json:"pack_owner,omitempty"`
	PackName  string `protobuf:"bytes,3,opt,name=pack_name,json=packName,proto3" json:"pack_name,omitempty"`
}

func (x *GetEmotePacksResponse_EmotePack) Reset() {
	*x = GetEmotePacksResponse_EmotePack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_emotes_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmotePacksResponse_EmotePack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmotePacksResponse_EmotePack) ProtoMessage() {}

func (x *GetEmotePacksResponse_EmotePack) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_emotes_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmotePacksResponse_EmotePack.ProtoReflect.Descriptor instead.
func (*GetEmotePacksResponse_EmotePack) Descriptor() ([]byte, []int) {
	return file_chat_v1_emotes_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetEmotePacksResponse_EmotePack) GetPackId() uint64 {
	if x != nil {
		return x.PackId
	}
	return 0
}

func (x *GetEmotePacksResponse_EmotePack) GetPackOwner() uint64 {
	if x != nil {
		return x.PackOwner
	}
	return 0
}

func (x *GetEmotePacksResponse_EmotePack) GetPackName() string {
	if x != nil {
		return x.PackName
	}
	return ""
}

type GetEmotePackEmotesResponse_Emote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageId string `protobuf:"bytes,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetEmotePackEmotesResponse_Emote) Reset() {
	*x = GetEmotePackEmotesResponse_Emote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_emotes_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmotePackEmotesResponse_Emote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmotePackEmotesResponse_Emote) ProtoMessage() {}

func (x *GetEmotePackEmotesResponse_Emote) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_emotes_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmotePackEmotesResponse_Emote.ProtoReflect.Descriptor instead.
func (*GetEmotePackEmotesResponse_Emote) Descriptor() ([]byte, []int) {
	return file_chat_v1_emotes_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetEmotePackEmotesResponse_Emote) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *GetEmotePackEmotesResponse_Emote) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_chat_v1_emotes_proto protoreflect.FileDescriptor

var file_chat_v1_emotes_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x35, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x36, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x07, 0x70, 0x61,
	0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52,
	0x06, 0x70, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x45, 0x6d,
	0x6f, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xca, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x05, 0x70, 0x61, 0x63,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x05, 0x70, 0x61, 0x63,
	0x6b, 0x73, 0x1a, 0x68, 0x0a, 0x09, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x12,
	0x1b, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x02, 0x30, 0x01, 0x52, 0x06, 0x70, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0a,
	0x70, 0x61, 0x63, 0x6b, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x02, 0x30, 0x01, 0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x45, 0x6d, 0x6f, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x07, 0x70, 0x61, 0x63,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x06,
	0x70, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x22, 0xa0, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x45, 0x6d,
	0x6f, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x6f, 0x74,
	0x65, 0x50, 0x61, 0x63, 0x6b, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x06, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x73, 0x1a, 0x36, 0x0a, 0x05, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x63, 0x0a, 0x15, 0x41, 0x64, 0x64,
	0x45, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x06, 0x70, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x54,
	0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x07,
	0x70, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30,
	0x01, 0x52, 0x06, 0x70, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d,
	0x6f, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x02, 0x30, 0x01, 0x52, 0x06, 0x70, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x16, 0x44,
	0x65, 0x71, 0x75, 0x69, 0x70, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x06, 0x70, 0x61, 0x63, 0x6b,
	0x49, 0x64, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2d, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_v1_emotes_proto_rawDescOnce sync.Once
	file_chat_v1_emotes_proto_rawDescData = file_chat_v1_emotes_proto_rawDesc
)

func file_chat_v1_emotes_proto_rawDescGZIP() []byte {
	file_chat_v1_emotes_proto_rawDescOnce.Do(func() {
		file_chat_v1_emotes_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_v1_emotes_proto_rawDescData)
	})
	return file_chat_v1_emotes_proto_rawDescData
}

var file_chat_v1_emotes_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_chat_v1_emotes_proto_goTypes = []interface{}{
	(*CreateEmotePackRequest)(nil),           // 0: protocol.chat.v1.CreateEmotePackRequest
	(*CreateEmotePackResponse)(nil),          // 1: protocol.chat.v1.CreateEmotePackResponse
	(*GetEmotePacksRequest)(nil),             // 2: protocol.chat.v1.GetEmotePacksRequest
	(*GetEmotePacksResponse)(nil),            // 3: protocol.chat.v1.GetEmotePacksResponse
	(*GetEmotePackEmotesRequest)(nil),        // 4: protocol.chat.v1.GetEmotePackEmotesRequest
	(*GetEmotePackEmotesResponse)(nil),       // 5: protocol.chat.v1.GetEmotePackEmotesResponse
	(*AddEmoteToPackRequest)(nil),            // 6: protocol.chat.v1.AddEmoteToPackRequest
	(*DeleteEmoteFromPackRequest)(nil),       // 7: protocol.chat.v1.DeleteEmoteFromPackRequest
	(*DeleteEmotePackRequest)(nil),           // 8: protocol.chat.v1.DeleteEmotePackRequest
	(*DequipEmotePackRequest)(nil),           // 9: protocol.chat.v1.DequipEmotePackRequest
	(*GetEmotePacksResponse_EmotePack)(nil),  // 10: protocol.chat.v1.GetEmotePacksResponse.EmotePack
	(*GetEmotePackEmotesResponse_Emote)(nil), // 11: protocol.chat.v1.GetEmotePackEmotesResponse.Emote
}
var file_chat_v1_emotes_proto_depIdxs = []int32{
	10, // 0: protocol.chat.v1.GetEmotePacksResponse.packs:type_name -> protocol.chat.v1.GetEmotePacksResponse.EmotePack
	11, // 1: protocol.chat.v1.GetEmotePackEmotesResponse.emotes:type_name -> protocol.chat.v1.GetEmotePackEmotesResponse.Emote
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_chat_v1_emotes_proto_init() }
func file_chat_v1_emotes_proto_init() {
	if File_chat_v1_emotes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_v1_emotes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmotePackRequest); i {
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
		file_chat_v1_emotes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmotePackResponse); i {
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
		file_chat_v1_emotes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmotePacksRequest); i {
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
		file_chat_v1_emotes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmotePacksResponse); i {
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
		file_chat_v1_emotes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmotePackEmotesRequest); i {
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
		file_chat_v1_emotes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmotePackEmotesResponse); i {
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
		file_chat_v1_emotes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEmoteToPackRequest); i {
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
		file_chat_v1_emotes_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEmoteFromPackRequest); i {
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
		file_chat_v1_emotes_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEmotePackRequest); i {
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
		file_chat_v1_emotes_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DequipEmotePackRequest); i {
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
		file_chat_v1_emotes_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmotePacksResponse_EmotePack); i {
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
		file_chat_v1_emotes_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmotePackEmotesResponse_Emote); i {
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
			RawDescriptor: file_chat_v1_emotes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_v1_emotes_proto_goTypes,
		DependencyIndexes: file_chat_v1_emotes_proto_depIdxs,
		MessageInfos:      file_chat_v1_emotes_proto_msgTypes,
	}.Build()
	File_chat_v1_emotes_proto = out.File
	file_chat_v1_emotes_proto_rawDesc = nil
	file_chat_v1_emotes_proto_goTypes = nil
	file_chat_v1_emotes_proto_depIdxs = nil
}