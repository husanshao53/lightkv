// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: bridge.proto

package bridge

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *PingReq) Reset() {
	*x = PingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReq) ProtoMessage() {}

func (x *PingReq) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReq.ProtoReflect.Descriptor instead.
func (*PingReq) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{0}
}

func (x *PingReq) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type PingRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *PingRsp) Reset() {
	*x = PingRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRsp) ProtoMessage() {}

func (x *PingRsp) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRsp.ProtoReflect.Descriptor instead.
func (*PingRsp) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{1}
}

func (x *PingRsp) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{2}
}

func (x *GetReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetRsp) Reset() {
	*x = GetRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRsp) ProtoMessage() {}

func (x *GetRsp) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRsp.ProtoReflect.Descriptor instead.
func (*GetRsp) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{3}
}

func (x *GetRsp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetRsp) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type PutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Expire int64  `protobuf:"varint,3,opt,name=expire,proto3" json:"expire,omitempty"`
}

func (x *PutReq) Reset() {
	*x = PutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutReq) ProtoMessage() {}

func (x *PutReq) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutReq.ProtoReflect.Descriptor instead.
func (*PutReq) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{4}
}

func (x *PutReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutReq) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *PutReq) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

type PutRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PutRsp) Reset() {
	*x = PutRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRsp) ProtoMessage() {}

func (x *PutRsp) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRsp.ProtoReflect.Descriptor instead.
func (*PutRsp) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{5}
}

func (x *PutRsp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutRsp) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type DelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DelReq) Reset() {
	*x = DelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelReq) ProtoMessage() {}

func (x *DelReq) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelReq.ProtoReflect.Descriptor instead.
func (*DelReq) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{6}
}

func (x *DelReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DelRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DelRsp) Reset() {
	*x = DelRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelRsp) ProtoMessage() {}

func (x *DelRsp) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelRsp.ProtoReflect.Descriptor instead.
func (*DelRsp) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{7}
}

func (x *DelRsp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type PublishReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *PublishReq) Reset() {
	*x = PublishReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishReq) ProtoMessage() {}

func (x *PublishReq) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishReq.ProtoReflect.Descriptor instead.
func (*PublishReq) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{8}
}

func (x *PublishReq) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type PublishRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Type  int32  `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *PublishRsp) Reset() {
	*x = PublishRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRsp) ProtoMessage() {}

func (x *PublishRsp) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRsp.ProtoReflect.Descriptor instead.
func (*PublishRsp) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{9}
}

func (x *PublishRsp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PublishRsp) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *PublishRsp) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type WatchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *WatchReq) Reset() {
	*x = WatchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchReq) ProtoMessage() {}

func (x *WatchReq) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchReq.ProtoReflect.Descriptor instead.
func (*WatchReq) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{10}
}

func (x *WatchReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type WatchRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *WatchRsp) Reset() {
	*x = WatchRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRsp) ProtoMessage() {}

func (x *WatchRsp) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRsp.ProtoReflect.Descriptor instead.
func (*WatchRsp) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{11}
}

func (x *WatchRsp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_bridge_proto protoreflect.FileDescriptor

var file_bridge_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x22, 0x27, 0x0a, 0x07, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x27, 0x0a, 0x07, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x1a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x30, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x52, 0x73, 0x70, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x48, 0x0a, 0x06, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x22, 0x30, 0x0a, 0x06, 0x50, 0x75, 0x74, 0x52, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x1a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x1a,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x52, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x2a, 0x0a, 0x0a, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x48, 0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x52, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x1c, 0x0a, 0x08, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x1c,
	0x0a, 0x08, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x32, 0xd1, 0x02, 0x0a,
	0x09, 0x52, 0x70, 0x63, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x2e,
	0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e,
	0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x27, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x0e, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e,
	0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e,
	0x50, 0x75, 0x74, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x03, 0x44, 0x65, 0x6c, 0x12,
	0x0e, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a,
	0x0e, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x52, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x37, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x12, 0x2e, 0x62,
	0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x1a, 0x12, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x52, 0x73, 0x70, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x30, 0x0a, 0x08, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67,
	0x65, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0a,
	0x55, 0x6e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x2e, 0x62, 0x72, 0x69,
	0x64, 0x67, 0x65, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x62,
	0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x73, 0x70, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bridge_proto_rawDescOnce sync.Once
	file_bridge_proto_rawDescData = file_bridge_proto_rawDesc
)

func file_bridge_proto_rawDescGZIP() []byte {
	file_bridge_proto_rawDescOnce.Do(func() {
		file_bridge_proto_rawDescData = protoimpl.X.CompressGZIP(file_bridge_proto_rawDescData)
	})
	return file_bridge_proto_rawDescData
}

var file_bridge_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_bridge_proto_goTypes = []interface{}{
	(*PingReq)(nil),    // 0: bridge.PingReq
	(*PingRsp)(nil),    // 1: bridge.PingRsp
	(*GetReq)(nil),     // 2: bridge.GetReq
	(*GetRsp)(nil),     // 3: bridge.GetRsp
	(*PutReq)(nil),     // 4: bridge.PutReq
	(*PutRsp)(nil),     // 5: bridge.PutRsp
	(*DelReq)(nil),     // 6: bridge.DelReq
	(*DelRsp)(nil),     // 7: bridge.DelRsp
	(*PublishReq)(nil), // 8: bridge.PublishReq
	(*PublishRsp)(nil), // 9: bridge.PublishRsp
	(*WatchReq)(nil),   // 10: bridge.WatchReq
	(*WatchRsp)(nil),   // 11: bridge.WatchRsp
}
var file_bridge_proto_depIdxs = []int32{
	0,  // 0: bridge.RpcBridge.Ping:input_type -> bridge.PingReq
	2,  // 1: bridge.RpcBridge.Get:input_type -> bridge.GetReq
	4,  // 2: bridge.RpcBridge.Put:input_type -> bridge.PutReq
	6,  // 3: bridge.RpcBridge.Del:input_type -> bridge.DelReq
	8,  // 4: bridge.RpcBridge.Publish:input_type -> bridge.PublishReq
	10, // 5: bridge.RpcBridge.WatchKey:input_type -> bridge.WatchReq
	10, // 6: bridge.RpcBridge.UnWatchKey:input_type -> bridge.WatchReq
	1,  // 7: bridge.RpcBridge.Ping:output_type -> bridge.PingRsp
	3,  // 8: bridge.RpcBridge.Get:output_type -> bridge.GetRsp
	5,  // 9: bridge.RpcBridge.Put:output_type -> bridge.PutRsp
	7,  // 10: bridge.RpcBridge.Del:output_type -> bridge.DelRsp
	9,  // 11: bridge.RpcBridge.Publish:output_type -> bridge.PublishRsp
	11, // 12: bridge.RpcBridge.WatchKey:output_type -> bridge.WatchRsp
	11, // 13: bridge.RpcBridge.UnWatchKey:output_type -> bridge.WatchRsp
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_bridge_proto_init() }
func file_bridge_proto_init() {
	if File_bridge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bridge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReq); i {
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
		file_bridge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRsp); i {
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
		file_bridge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReq); i {
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
		file_bridge_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRsp); i {
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
		file_bridge_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutReq); i {
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
		file_bridge_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutRsp); i {
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
		file_bridge_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelReq); i {
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
		file_bridge_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelRsp); i {
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
		file_bridge_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishReq); i {
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
		file_bridge_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRsp); i {
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
		file_bridge_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchReq); i {
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
		file_bridge_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRsp); i {
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
			RawDescriptor: file_bridge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bridge_proto_goTypes,
		DependencyIndexes: file_bridge_proto_depIdxs,
		MessageInfos:      file_bridge_proto_msgTypes,
	}.Build()
	File_bridge_proto = out.File
	file_bridge_proto_rawDesc = nil
	file_bridge_proto_goTypes = nil
	file_bridge_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RpcBridgeClient is the client API for RpcBridge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RpcBridgeClient interface {
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRsp, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRsp, error)
	Put(ctx context.Context, in *PutReq, opts ...grpc.CallOption) (*PutRsp, error)
	Del(ctx context.Context, in *DelReq, opts ...grpc.CallOption) (*DelRsp, error)
	Publish(ctx context.Context, opts ...grpc.CallOption) (RpcBridge_PublishClient, error)
	WatchKey(ctx context.Context, in *WatchReq, opts ...grpc.CallOption) (*WatchRsp, error)
	UnWatchKey(ctx context.Context, in *WatchReq, opts ...grpc.CallOption) (*WatchRsp, error)
}

type rpcBridgeClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcBridgeClient(cc grpc.ClientConnInterface) RpcBridgeClient {
	return &rpcBridgeClient{cc}
}

func (c *rpcBridgeClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRsp, error) {
	out := new(PingRsp)
	err := c.cc.Invoke(ctx, "/bridge.RpcBridge/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcBridgeClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRsp, error) {
	out := new(GetRsp)
	err := c.cc.Invoke(ctx, "/bridge.RpcBridge/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcBridgeClient) Put(ctx context.Context, in *PutReq, opts ...grpc.CallOption) (*PutRsp, error) {
	out := new(PutRsp)
	err := c.cc.Invoke(ctx, "/bridge.RpcBridge/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcBridgeClient) Del(ctx context.Context, in *DelReq, opts ...grpc.CallOption) (*DelRsp, error) {
	out := new(DelRsp)
	err := c.cc.Invoke(ctx, "/bridge.RpcBridge/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcBridgeClient) Publish(ctx context.Context, opts ...grpc.CallOption) (RpcBridge_PublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RpcBridge_serviceDesc.Streams[0], "/bridge.RpcBridge/Publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcBridgePublishClient{stream}
	return x, nil
}

type RpcBridge_PublishClient interface {
	Send(*PublishReq) error
	Recv() (*PublishRsp, error)
	grpc.ClientStream
}

type rpcBridgePublishClient struct {
	grpc.ClientStream
}

func (x *rpcBridgePublishClient) Send(m *PublishReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rpcBridgePublishClient) Recv() (*PublishRsp, error) {
	m := new(PublishRsp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rpcBridgeClient) WatchKey(ctx context.Context, in *WatchReq, opts ...grpc.CallOption) (*WatchRsp, error) {
	out := new(WatchRsp)
	err := c.cc.Invoke(ctx, "/bridge.RpcBridge/WatchKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcBridgeClient) UnWatchKey(ctx context.Context, in *WatchReq, opts ...grpc.CallOption) (*WatchRsp, error) {
	out := new(WatchRsp)
	err := c.cc.Invoke(ctx, "/bridge.RpcBridge/UnWatchKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcBridgeServer is the server API for RpcBridge service.
type RpcBridgeServer interface {
	Ping(context.Context, *PingReq) (*PingRsp, error)
	Get(context.Context, *GetReq) (*GetRsp, error)
	Put(context.Context, *PutReq) (*PutRsp, error)
	Del(context.Context, *DelReq) (*DelRsp, error)
	Publish(RpcBridge_PublishServer) error
	WatchKey(context.Context, *WatchReq) (*WatchRsp, error)
	UnWatchKey(context.Context, *WatchReq) (*WatchRsp, error)
}

// UnimplementedRpcBridgeServer can be embedded to have forward compatible implementations.
type UnimplementedRpcBridgeServer struct {
}

func (*UnimplementedRpcBridgeServer) Ping(context.Context, *PingReq) (*PingRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedRpcBridgeServer) Get(context.Context, *GetReq) (*GetRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRpcBridgeServer) Put(context.Context, *PutReq) (*PutRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedRpcBridgeServer) Del(context.Context, *DelReq) (*DelRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (*UnimplementedRpcBridgeServer) Publish(RpcBridge_PublishServer) error {
	return status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (*UnimplementedRpcBridgeServer) WatchKey(context.Context, *WatchReq) (*WatchRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WatchKey not implemented")
}
func (*UnimplementedRpcBridgeServer) UnWatchKey(context.Context, *WatchReq) (*WatchRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnWatchKey not implemented")
}

func RegisterRpcBridgeServer(s *grpc.Server, srv RpcBridgeServer) {
	s.RegisterService(&_RpcBridge_serviceDesc, srv)
}

func _RpcBridge_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcBridgeServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bridge.RpcBridge/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcBridgeServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcBridge_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcBridgeServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bridge.RpcBridge/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcBridgeServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcBridge_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcBridgeServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bridge.RpcBridge/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcBridgeServer).Put(ctx, req.(*PutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcBridge_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcBridgeServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bridge.RpcBridge/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcBridgeServer).Del(ctx, req.(*DelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcBridge_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RpcBridgeServer).Publish(&rpcBridgePublishServer{stream})
}

type RpcBridge_PublishServer interface {
	Send(*PublishRsp) error
	Recv() (*PublishReq, error)
	grpc.ServerStream
}

type rpcBridgePublishServer struct {
	grpc.ServerStream
}

func (x *rpcBridgePublishServer) Send(m *PublishRsp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rpcBridgePublishServer) Recv() (*PublishReq, error) {
	m := new(PublishReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RpcBridge_WatchKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcBridgeServer).WatchKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bridge.RpcBridge/WatchKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcBridgeServer).WatchKey(ctx, req.(*WatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcBridge_UnWatchKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcBridgeServer).UnWatchKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bridge.RpcBridge/UnWatchKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcBridgeServer).UnWatchKey(ctx, req.(*WatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RpcBridge_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bridge.RpcBridge",
	HandlerType: (*RpcBridgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _RpcBridge_Ping_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RpcBridge_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _RpcBridge_Put_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _RpcBridge_Del_Handler,
		},
		{
			MethodName: "WatchKey",
			Handler:    _RpcBridge_WatchKey_Handler,
		},
		{
			MethodName: "UnWatchKey",
			Handler:    _RpcBridge_UnWatchKey_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Publish",
			Handler:       _RpcBridge_Publish_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bridge.proto",
}
