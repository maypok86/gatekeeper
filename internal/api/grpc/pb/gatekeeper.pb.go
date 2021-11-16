// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: gatekeeper.proto

package apipb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AllowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Ip       string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *AllowRequest) Reset() {
	*x = AllowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatekeeper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowRequest) ProtoMessage() {}

func (x *AllowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gatekeeper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowRequest.ProtoReflect.Descriptor instead.
func (*AllowRequest) Descriptor() ([]byte, []int) {
	return file_gatekeeper_proto_rawDescGZIP(), []int{0}
}

func (x *AllowRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *AllowRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AllowRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type AllowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *AllowResponse) Reset() {
	*x = AllowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatekeeper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowResponse) ProtoMessage() {}

func (x *AllowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gatekeeper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowResponse.ProtoReflect.Descriptor instead.
func (*AllowResponse) Descriptor() ([]byte, []int) {
	return file_gatekeeper_proto_rawDescGZIP(), []int{1}
}

func (x *AllowResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type ResetLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *ResetLoginRequest) Reset() {
	*x = ResetLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatekeeper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetLoginRequest) ProtoMessage() {}

func (x *ResetLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gatekeeper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetLoginRequest.ProtoReflect.Descriptor instead.
func (*ResetLoginRequest) Descriptor() ([]byte, []int) {
	return file_gatekeeper_proto_rawDescGZIP(), []int{2}
}

func (x *ResetLoginRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

type ResetLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ResetLoginResponse) Reset() {
	*x = ResetLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatekeeper_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetLoginResponse) ProtoMessage() {}

func (x *ResetLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gatekeeper_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetLoginResponse.ProtoReflect.Descriptor instead.
func (*ResetLoginResponse) Descriptor() ([]byte, []int) {
	return file_gatekeeper_proto_rawDescGZIP(), []int{3}
}

func (x *ResetLoginResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type ResetIPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *ResetIPRequest) Reset() {
	*x = ResetIPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatekeeper_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetIPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetIPRequest) ProtoMessage() {}

func (x *ResetIPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gatekeeper_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetIPRequest.ProtoReflect.Descriptor instead.
func (*ResetIPRequest) Descriptor() ([]byte, []int) {
	return file_gatekeeper_proto_rawDescGZIP(), []int{4}
}

func (x *ResetIPRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type ResetIPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ResetIPResponse) Reset() {
	*x = ResetIPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatekeeper_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetIPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetIPResponse) ProtoMessage() {}

func (x *ResetIPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gatekeeper_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetIPResponse.ProtoReflect.Descriptor instead.
func (*ResetIPResponse) Descriptor() ([]byte, []int) {
	return file_gatekeeper_proto_rawDescGZIP(), []int{5}
}

func (x *ResetIPResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type WhitelistAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subnet string `protobuf:"bytes,1,opt,name=subnet,proto3" json:"subnet,omitempty"`
}

func (x *WhitelistAddRequest) Reset() {
	*x = WhitelistAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatekeeper_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhitelistAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhitelistAddRequest) ProtoMessage() {}

func (x *WhitelistAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gatekeeper_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhitelistAddRequest.ProtoReflect.Descriptor instead.
func (*WhitelistAddRequest) Descriptor() ([]byte, []int) {
	return file_gatekeeper_proto_rawDescGZIP(), []int{6}
}

func (x *WhitelistAddRequest) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

type WhitelistAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *WhitelistAddResponse) Reset() {
	*x = WhitelistAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatekeeper_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhitelistAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhitelistAddResponse) ProtoMessage() {}

func (x *WhitelistAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gatekeeper_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhitelistAddResponse.ProtoReflect.Descriptor instead.
func (*WhitelistAddResponse) Descriptor() ([]byte, []int) {
	return file_gatekeeper_proto_rawDescGZIP(), []int{7}
}

func (x *WhitelistAddResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type WhitelistRemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subnet string `protobuf:"bytes,1,opt,name=subnet,proto3" json:"subnet,omitempty"`
}

func (x *WhitelistRemoveRequest) Reset() {
	*x = WhitelistRemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatekeeper_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhitelistRemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhitelistRemoveRequest) ProtoMessage() {}

func (x *WhitelistRemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gatekeeper_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhitelistRemoveRequest.ProtoReflect.Descriptor instead.
func (*WhitelistRemoveRequest) Descriptor() ([]byte, []int) {
	return file_gatekeeper_proto_rawDescGZIP(), []int{8}
}

func (x *WhitelistRemoveRequest) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

type WhitelistRemoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *WhitelistRemoveResponse) Reset() {
	*x = WhitelistRemoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatekeeper_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhitelistRemoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhitelistRemoveResponse) ProtoMessage() {}

func (x *WhitelistRemoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gatekeeper_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhitelistRemoveResponse.ProtoReflect.Descriptor instead.
func (*WhitelistRemoveResponse) Descriptor() ([]byte, []int) {
	return file_gatekeeper_proto_rawDescGZIP(), []int{9}
}

func (x *WhitelistRemoveResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type BlacklistAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subnet string `protobuf:"bytes,1,opt,name=subnet,proto3" json:"subnet,omitempty"`
}

func (x *BlacklistAddRequest) Reset() {
	*x = BlacklistAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatekeeper_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlacklistAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlacklistAddRequest) ProtoMessage() {}

func (x *BlacklistAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gatekeeper_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlacklistAddRequest.ProtoReflect.Descriptor instead.
func (*BlacklistAddRequest) Descriptor() ([]byte, []int) {
	return file_gatekeeper_proto_rawDescGZIP(), []int{10}
}

func (x *BlacklistAddRequest) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

type BlacklistAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *BlacklistAddResponse) Reset() {
	*x = BlacklistAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatekeeper_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlacklistAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlacklistAddResponse) ProtoMessage() {}

func (x *BlacklistAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gatekeeper_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlacklistAddResponse.ProtoReflect.Descriptor instead.
func (*BlacklistAddResponse) Descriptor() ([]byte, []int) {
	return file_gatekeeper_proto_rawDescGZIP(), []int{11}
}

func (x *BlacklistAddResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type BlacklistRemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subnet string `protobuf:"bytes,1,opt,name=subnet,proto3" json:"subnet,omitempty"`
}

func (x *BlacklistRemoveRequest) Reset() {
	*x = BlacklistRemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatekeeper_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlacklistRemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlacklistRemoveRequest) ProtoMessage() {}

func (x *BlacklistRemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gatekeeper_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlacklistRemoveRequest.ProtoReflect.Descriptor instead.
func (*BlacklistRemoveRequest) Descriptor() ([]byte, []int) {
	return file_gatekeeper_proto_rawDescGZIP(), []int{12}
}

func (x *BlacklistRemoveRequest) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

type BlacklistRemoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *BlacklistRemoveResponse) Reset() {
	*x = BlacklistRemoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gatekeeper_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlacklistRemoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlacklistRemoveResponse) ProtoMessage() {}

func (x *BlacklistRemoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gatekeeper_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlacklistRemoveResponse.ProtoReflect.Descriptor instead.
func (*BlacklistRemoveResponse) Descriptor() ([]byte, []int) {
	return file_gatekeeper_proto_rawDescGZIP(), []int{13}
}

func (x *BlacklistRemoveResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_gatekeeper_proto protoreflect.FileDescriptor

var file_gatekeeper_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x61, 0x74, 0x65, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x50, 0x0a, 0x0c, 0x41, 0x6c, 0x6c, 0x6f, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x1f, 0x0a, 0x0d, 0x41, 0x6c, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x29, 0x0a, 0x11, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x24, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x20, 0x0a, 0x0e, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x49, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x21, 0x0a,
	0x0f, 0x52, 0x65, 0x73, 0x65, 0x74, 0x49, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b,
	0x22, 0x2d, 0x0a, 0x13, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x22,
	0x26, 0x0a, 0x14, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x30, 0x0a, 0x16, 0x57, 0x68, 0x69, 0x74, 0x65,
	0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x22, 0x29, 0x0a, 0x17, 0x57, 0x68, 0x69,
	0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x02, 0x6f, 0x6b, 0x22, 0x2d, 0x0a, 0x13, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73,
	0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x30, 0x0a, 0x16, 0x42,
	0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x22, 0x29, 0x0a,
	0x17, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0xec, 0x03, 0x0a, 0x11, 0x47, 0x61, 0x74,
	0x65, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30,
	0x0a, 0x05, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3f, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x36, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x65, 0x74, 0x49, 0x50, 0x12, 0x13, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x49, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x49, 0x50, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x57, 0x68, 0x69,
	0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x64, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c,
	0x69, 0x73, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4e, 0x0a, 0x0f, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x45, 0x0a, 0x0c, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x64,
	0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0f, 0x42, 0x6c, 0x61, 0x63, 0x6b,
	0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6c,
	0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x61, 0x70,
	0x69, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gatekeeper_proto_rawDescOnce sync.Once
	file_gatekeeper_proto_rawDescData = file_gatekeeper_proto_rawDesc
)

func file_gatekeeper_proto_rawDescGZIP() []byte {
	file_gatekeeper_proto_rawDescOnce.Do(func() {
		file_gatekeeper_proto_rawDescData = protoimpl.X.CompressGZIP(file_gatekeeper_proto_rawDescData)
	})
	return file_gatekeeper_proto_rawDescData
}

var file_gatekeeper_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_gatekeeper_proto_goTypes = []interface{}{
	(*AllowRequest)(nil),            // 0: api.AllowRequest
	(*AllowResponse)(nil),           // 1: api.AllowResponse
	(*ResetLoginRequest)(nil),       // 2: api.ResetLoginRequest
	(*ResetLoginResponse)(nil),      // 3: api.ResetLoginResponse
	(*ResetIPRequest)(nil),          // 4: api.ResetIPRequest
	(*ResetIPResponse)(nil),         // 5: api.ResetIPResponse
	(*WhitelistAddRequest)(nil),     // 6: api.WhitelistAddRequest
	(*WhitelistAddResponse)(nil),    // 7: api.WhitelistAddResponse
	(*WhitelistRemoveRequest)(nil),  // 8: api.WhitelistRemoveRequest
	(*WhitelistRemoveResponse)(nil), // 9: api.WhitelistRemoveResponse
	(*BlacklistAddRequest)(nil),     // 10: api.BlacklistAddRequest
	(*BlacklistAddResponse)(nil),    // 11: api.BlacklistAddResponse
	(*BlacklistRemoveRequest)(nil),  // 12: api.BlacklistRemoveRequest
	(*BlacklistRemoveResponse)(nil), // 13: api.BlacklistRemoveResponse
}
var file_gatekeeper_proto_depIdxs = []int32{
	0,  // 0: api.GatekeeperService.Allow:input_type -> api.AllowRequest
	2,  // 1: api.GatekeeperService.ResetLogin:input_type -> api.ResetLoginRequest
	4,  // 2: api.GatekeeperService.ResetIP:input_type -> api.ResetIPRequest
	6,  // 3: api.GatekeeperService.WhitelistAdd:input_type -> api.WhitelistAddRequest
	8,  // 4: api.GatekeeperService.WhitelistRemove:input_type -> api.WhitelistRemoveRequest
	10, // 5: api.GatekeeperService.BlacklistAdd:input_type -> api.BlacklistAddRequest
	12, // 6: api.GatekeeperService.BlacklistRemove:input_type -> api.BlacklistRemoveRequest
	1,  // 7: api.GatekeeperService.Allow:output_type -> api.AllowResponse
	3,  // 8: api.GatekeeperService.ResetLogin:output_type -> api.ResetLoginResponse
	5,  // 9: api.GatekeeperService.ResetIP:output_type -> api.ResetIPResponse
	7,  // 10: api.GatekeeperService.WhitelistAdd:output_type -> api.WhitelistAddResponse
	9,  // 11: api.GatekeeperService.WhitelistRemove:output_type -> api.WhitelistRemoveResponse
	11, // 12: api.GatekeeperService.BlacklistAdd:output_type -> api.BlacklistAddResponse
	13, // 13: api.GatekeeperService.BlacklistRemove:output_type -> api.BlacklistRemoveResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_gatekeeper_proto_init() }
func file_gatekeeper_proto_init() {
	if File_gatekeeper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gatekeeper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowRequest); i {
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
		file_gatekeeper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowResponse); i {
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
		file_gatekeeper_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetLoginRequest); i {
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
		file_gatekeeper_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetLoginResponse); i {
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
		file_gatekeeper_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetIPRequest); i {
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
		file_gatekeeper_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetIPResponse); i {
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
		file_gatekeeper_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhitelistAddRequest); i {
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
		file_gatekeeper_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhitelistAddResponse); i {
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
		file_gatekeeper_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhitelistRemoveRequest); i {
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
		file_gatekeeper_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhitelistRemoveResponse); i {
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
		file_gatekeeper_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlacklistAddRequest); i {
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
		file_gatekeeper_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlacklistAddResponse); i {
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
		file_gatekeeper_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlacklistRemoveRequest); i {
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
		file_gatekeeper_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlacklistRemoveResponse); i {
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
			RawDescriptor: file_gatekeeper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gatekeeper_proto_goTypes,
		DependencyIndexes: file_gatekeeper_proto_depIdxs,
		MessageInfos:      file_gatekeeper_proto_msgTypes,
	}.Build()
	File_gatekeeper_proto = out.File
	file_gatekeeper_proto_rawDesc = nil
	file_gatekeeper_proto_goTypes = nil
	file_gatekeeper_proto_depIdxs = nil
}
