// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: pb/bridge.proto

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
		mi := &file_pb_bridge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReq) ProtoMessage() {}

func (x *PingReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_bridge_proto_msgTypes[0]
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
	return file_pb_bridge_proto_rawDescGZIP(), []int{0}
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
		mi := &file_pb_bridge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRsp) ProtoMessage() {}

func (x *PingRsp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_bridge_proto_msgTypes[1]
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
	return file_pb_bridge_proto_rawDescGZIP(), []int{1}
}

func (x *PingRsp) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_pb_bridge_proto protoreflect.FileDescriptor

var file_pb_bridge_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x62, 0x2f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x22, 0x27, 0x0a, 0x07, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0x27, 0x0a, 0x07, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x73, 0x70, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x37, 0x0a, 0x09, 0x52,
	0x70, 0x63, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67,
	0x12, 0x0f, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x1a, 0x0f, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x73, 0x70, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_bridge_proto_rawDescOnce sync.Once
	file_pb_bridge_proto_rawDescData = file_pb_bridge_proto_rawDesc
)

func file_pb_bridge_proto_rawDescGZIP() []byte {
	file_pb_bridge_proto_rawDescOnce.Do(func() {
		file_pb_bridge_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_bridge_proto_rawDescData)
	})
	return file_pb_bridge_proto_rawDescData
}

var file_pb_bridge_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_bridge_proto_goTypes = []interface{}{
	(*PingReq)(nil), // 0: bridge.PingReq
	(*PingRsp)(nil), // 1: bridge.PingRsp
}
var file_pb_bridge_proto_depIdxs = []int32{
	0, // 0: bridge.RpcBridge.Ping:input_type -> bridge.PingReq
	1, // 1: bridge.RpcBridge.Ping:output_type -> bridge.PingRsp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_bridge_proto_init() }
func file_pb_bridge_proto_init() {
	if File_pb_bridge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_bridge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pb_bridge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_bridge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_bridge_proto_goTypes,
		DependencyIndexes: file_pb_bridge_proto_depIdxs,
		MessageInfos:      file_pb_bridge_proto_msgTypes,
	}.Build()
	File_pb_bridge_proto = out.File
	file_pb_bridge_proto_rawDesc = nil
	file_pb_bridge_proto_goTypes = nil
	file_pb_bridge_proto_depIdxs = nil
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

// RpcBridgeServer is the server API for RpcBridge service.
type RpcBridgeServer interface {
	Ping(context.Context, *PingReq) (*PingRsp, error)
}

// UnimplementedRpcBridgeServer can be embedded to have forward compatible implementations.
type UnimplementedRpcBridgeServer struct {
}

func (*UnimplementedRpcBridgeServer) Ping(context.Context, *PingReq) (*PingRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
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

var _RpcBridge_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bridge.RpcBridge",
	HandlerType: (*RpcBridgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _RpcBridge_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/bridge.proto",
}
