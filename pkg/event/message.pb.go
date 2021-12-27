// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: onebot-plus/pkg/event/message.proto

package event

import (
	model "github.com/dezhishen/onebot-sdk/pkg/model"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_onebot_plus_pkg_event_message_proto protoreflect.FileDescriptor

var file_onebot_plus_pkg_event_message_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6f, 0x6e, 0x65, 0x62, 0x6f, 0x74, 0x2d, 0x70, 0x6c, 0x75, 0x73, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x28, 0x6f, 0x6e,
	0x65, 0x62, 0x6f, 0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x55, 0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x73, 0x67, 0x12, 0x40, 0x0a, 0x06, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x12, 0x1e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x47, 0x52,
	0x50, 0x43, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x51, 0x0a, 0x0f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x73, 0x67, 0x12, 0x3e, 0x0a,
	0x06, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x47, 0x52, 0x50, 0x43, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_onebot_plus_pkg_event_message_proto_goTypes = []interface{}{
	(*model.EventMessagePrivateGRPC)(nil), // 0: model.EventMessagePrivateGRPC
	(*model.EventMessageGroupGRPC)(nil),   // 1: model.EventMessageGroupGRPC
	(*emptypb.Empty)(nil),                 // 2: google.protobuf.Empty
}
var file_onebot_plus_pkg_event_message_proto_depIdxs = []int32{
	0, // 0: event.event_private_msg.Handle:input_type -> model.EventMessagePrivateGRPC
	1, // 1: event.event_group_msg.Handle:input_type -> model.EventMessageGroupGRPC
	2, // 2: event.event_private_msg.Handle:output_type -> google.protobuf.Empty
	2, // 3: event.event_group_msg.Handle:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_onebot_plus_pkg_event_message_proto_init() }
func file_onebot_plus_pkg_event_message_proto_init() {
	if File_onebot_plus_pkg_event_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_onebot_plus_pkg_event_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_onebot_plus_pkg_event_message_proto_goTypes,
		DependencyIndexes: file_onebot_plus_pkg_event_message_proto_depIdxs,
	}.Build()
	File_onebot_plus_pkg_event_message_proto = out.File
	file_onebot_plus_pkg_event_message_proto_rawDesc = nil
	file_onebot_plus_pkg_event_message_proto_goTypes = nil
	file_onebot_plus_pkg_event_message_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventPrivateMsgClient is the client API for EventPrivateMsg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventPrivateMsgClient interface {
	//发送消息
	Handle(ctx context.Context, in *model.EventMessagePrivateGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type eventPrivateMsgClient struct {
	cc grpc.ClientConnInterface
}

func NewEventPrivateMsgClient(cc grpc.ClientConnInterface) EventPrivateMsgClient {
	return &eventPrivateMsgClient{cc}
}

func (c *eventPrivateMsgClient) Handle(ctx context.Context, in *model.EventMessagePrivateGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/event.event_private_msg/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventPrivateMsgServer is the server API for EventPrivateMsg service.
type EventPrivateMsgServer interface {
	//发送消息
	Handle(context.Context, *model.EventMessagePrivateGRPC) (*emptypb.Empty, error)
}

// UnimplementedEventPrivateMsgServer can be embedded to have forward compatible implementations.
type UnimplementedEventPrivateMsgServer struct {
}

func (*UnimplementedEventPrivateMsgServer) Handle(context.Context, *model.EventMessagePrivateGRPC) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterEventPrivateMsgServer(s *grpc.Server, srv EventPrivateMsgServer) {
	s.RegisterService(&_EventPrivateMsg_serviceDesc, srv)
}

func _EventPrivateMsg_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EventMessagePrivateGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventPrivateMsgServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.event_private_msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventPrivateMsgServer).Handle(ctx, req.(*model.EventMessagePrivateGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventPrivateMsg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "event.event_private_msg",
	HandlerType: (*EventPrivateMsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _EventPrivateMsg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "onebot-plus/pkg/event/message.proto",
}

// EventGroupMsgClient is the client API for EventGroupMsg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventGroupMsgClient interface {
	Handle(ctx context.Context, in *model.EventMessageGroupGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type eventGroupMsgClient struct {
	cc grpc.ClientConnInterface
}

func NewEventGroupMsgClient(cc grpc.ClientConnInterface) EventGroupMsgClient {
	return &eventGroupMsgClient{cc}
}

func (c *eventGroupMsgClient) Handle(ctx context.Context, in *model.EventMessageGroupGRPC, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/event.event_group_msg/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventGroupMsgServer is the server API for EventGroupMsg service.
type EventGroupMsgServer interface {
	Handle(context.Context, *model.EventMessageGroupGRPC) (*emptypb.Empty, error)
}

// UnimplementedEventGroupMsgServer can be embedded to have forward compatible implementations.
type UnimplementedEventGroupMsgServer struct {
}

func (*UnimplementedEventGroupMsgServer) Handle(context.Context, *model.EventMessageGroupGRPC) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterEventGroupMsgServer(s *grpc.Server, srv EventGroupMsgServer) {
	s.RegisterService(&_EventGroupMsg_serviceDesc, srv)
}

func _EventGroupMsg_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.EventMessageGroupGRPC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventGroupMsgServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.event_group_msg/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventGroupMsgServer).Handle(ctx, req.(*model.EventMessageGroupGRPC))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventGroupMsg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "event.event_group_msg",
	HandlerType: (*EventGroupMsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _EventGroupMsg_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "onebot-plus/pkg/event/message.proto",
}
