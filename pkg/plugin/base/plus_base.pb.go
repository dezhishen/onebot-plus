// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: plus_base.proto

package base

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_plus_base_proto protoreflect.FileDescriptor

var file_plus_base_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x6c, 0x75, 0x73, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x95, 0x03, 0x0a, 0x14, 0x4f, 0x6e, 0x65, 0x62, 0x6f, 0x74, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x42, 0x61, 0x73, 0x65, 0x47, 0x52, 0x50, 0x43, 0x12, 0x3a, 0x0a,
	0x02, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x04,
	0x48, 0x65, 0x6c, 0x70, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x49, 0x6e,
	0x69, 0x74, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x0a, 0x42, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x45, 0x78, 0x69, 0x74, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x37, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x7a, 0x68, 0x69,
	0x73, 0x68, 0x65, 0x6e, 0x2f, 0x6f, 0x6e, 0x65, 0x62, 0x6f, 0x74, 0x2d, 0x70, 0x6c, 0x75, 0x73,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x3b, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_plus_base_proto_goTypes = []interface{}{
	(*emptypb.Empty)(nil),          // 0: google.protobuf.Empty
	(*wrapperspb.UInt32Value)(nil), // 1: google.protobuf.UInt32Value
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
}
var file_plus_base_proto_depIdxs = []int32{
	0, // 0: base.OnebotPluginBaseGRPC.Id:input_type -> google.protobuf.Empty
	0, // 1: base.OnebotPluginBaseGRPC.Name:input_type -> google.protobuf.Empty
	0, // 2: base.OnebotPluginBaseGRPC.Description:input_type -> google.protobuf.Empty
	0, // 3: base.OnebotPluginBaseGRPC.Help:input_type -> google.protobuf.Empty
	1, // 4: base.OnebotPluginBaseGRPC.Init:input_type -> google.protobuf.UInt32Value
	1, // 5: base.OnebotPluginBaseGRPC.BeforeExit:input_type -> google.protobuf.UInt32Value
	2, // 6: base.OnebotPluginBaseGRPC.Id:output_type -> google.protobuf.StringValue
	2, // 7: base.OnebotPluginBaseGRPC.Name:output_type -> google.protobuf.StringValue
	2, // 8: base.OnebotPluginBaseGRPC.Description:output_type -> google.protobuf.StringValue
	2, // 9: base.OnebotPluginBaseGRPC.Help:output_type -> google.protobuf.StringValue
	0, // 10: base.OnebotPluginBaseGRPC.Init:output_type -> google.protobuf.Empty
	0, // 11: base.OnebotPluginBaseGRPC.BeforeExit:output_type -> google.protobuf.Empty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_plus_base_proto_init() }
func file_plus_base_proto_init() {
	if File_plus_base_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plus_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plus_base_proto_goTypes,
		DependencyIndexes: file_plus_base_proto_depIdxs,
	}.Build()
	File_plus_base_proto = out.File
	file_plus_base_proto_rawDesc = nil
	file_plus_base_proto_goTypes = nil
	file_plus_base_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OnebotPluginBaseGRPCClient is the client API for OnebotPluginBaseGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OnebotPluginBaseGRPCClient interface {
	Id(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	Name(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	Description(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	Help(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	Init(ctx context.Context, in *wrapperspb.UInt32Value, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BeforeExit(ctx context.Context, in *wrapperspb.UInt32Value, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type onebotPluginBaseGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewOnebotPluginBaseGRPCClient(cc grpc.ClientConnInterface) OnebotPluginBaseGRPCClient {
	return &onebotPluginBaseGRPCClient{cc}
}

func (c *onebotPluginBaseGRPCClient) Id(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/base.OnebotPluginBaseGRPC/Id", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotPluginBaseGRPCClient) Name(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/base.OnebotPluginBaseGRPC/Name", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotPluginBaseGRPCClient) Description(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/base.OnebotPluginBaseGRPC/Description", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotPluginBaseGRPCClient) Help(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/base.OnebotPluginBaseGRPC/Help", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotPluginBaseGRPCClient) Init(ctx context.Context, in *wrapperspb.UInt32Value, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/base.OnebotPluginBaseGRPC/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onebotPluginBaseGRPCClient) BeforeExit(ctx context.Context, in *wrapperspb.UInt32Value, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/base.OnebotPluginBaseGRPC/BeforeExit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnebotPluginBaseGRPCServer is the server API for OnebotPluginBaseGRPC service.
type OnebotPluginBaseGRPCServer interface {
	Id(context.Context, *emptypb.Empty) (*wrapperspb.StringValue, error)
	Name(context.Context, *emptypb.Empty) (*wrapperspb.StringValue, error)
	Description(context.Context, *emptypb.Empty) (*wrapperspb.StringValue, error)
	Help(context.Context, *emptypb.Empty) (*wrapperspb.StringValue, error)
	Init(context.Context, *wrapperspb.UInt32Value) (*emptypb.Empty, error)
	BeforeExit(context.Context, *wrapperspb.UInt32Value) (*emptypb.Empty, error)
}

// UnimplementedOnebotPluginBaseGRPCServer can be embedded to have forward compatible implementations.
type UnimplementedOnebotPluginBaseGRPCServer struct {
}

func (*UnimplementedOnebotPluginBaseGRPCServer) Id(context.Context, *emptypb.Empty) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Id not implemented")
}
func (*UnimplementedOnebotPluginBaseGRPCServer) Name(context.Context, *emptypb.Empty) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Name not implemented")
}
func (*UnimplementedOnebotPluginBaseGRPCServer) Description(context.Context, *emptypb.Empty) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Description not implemented")
}
func (*UnimplementedOnebotPluginBaseGRPCServer) Help(context.Context, *emptypb.Empty) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Help not implemented")
}
func (*UnimplementedOnebotPluginBaseGRPCServer) Init(context.Context, *wrapperspb.UInt32Value) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (*UnimplementedOnebotPluginBaseGRPCServer) BeforeExit(context.Context, *wrapperspb.UInt32Value) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeforeExit not implemented")
}

func RegisterOnebotPluginBaseGRPCServer(s *grpc.Server, srv OnebotPluginBaseGRPCServer) {
	s.RegisterService(&_OnebotPluginBaseGRPC_serviceDesc, srv)
}

func _OnebotPluginBaseGRPC_Id_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotPluginBaseGRPCServer).Id(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.OnebotPluginBaseGRPC/Id",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotPluginBaseGRPCServer).Id(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotPluginBaseGRPC_Name_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotPluginBaseGRPCServer).Name(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.OnebotPluginBaseGRPC/Name",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotPluginBaseGRPCServer).Name(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotPluginBaseGRPC_Description_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotPluginBaseGRPCServer).Description(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.OnebotPluginBaseGRPC/Description",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotPluginBaseGRPCServer).Description(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotPluginBaseGRPC_Help_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotPluginBaseGRPCServer).Help(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.OnebotPluginBaseGRPC/Help",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotPluginBaseGRPCServer).Help(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotPluginBaseGRPC_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.UInt32Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotPluginBaseGRPCServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.OnebotPluginBaseGRPC/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotPluginBaseGRPCServer).Init(ctx, req.(*wrapperspb.UInt32Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnebotPluginBaseGRPC_BeforeExit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.UInt32Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnebotPluginBaseGRPCServer).BeforeExit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.OnebotPluginBaseGRPC/BeforeExit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnebotPluginBaseGRPCServer).BeforeExit(ctx, req.(*wrapperspb.UInt32Value))
	}
	return interceptor(ctx, in, info, handler)
}

var _OnebotPluginBaseGRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "base.OnebotPluginBaseGRPC",
	HandlerType: (*OnebotPluginBaseGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Id",
			Handler:    _OnebotPluginBaseGRPC_Id_Handler,
		},
		{
			MethodName: "Name",
			Handler:    _OnebotPluginBaseGRPC_Name_Handler,
		},
		{
			MethodName: "Description",
			Handler:    _OnebotPluginBaseGRPC_Description_Handler,
		},
		{
			MethodName: "Help",
			Handler:    _OnebotPluginBaseGRPC_Help_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _OnebotPluginBaseGRPC_Init_Handler,
		},
		{
			MethodName: "BeforeExit",
			Handler:    _OnebotPluginBaseGRPC_BeforeExit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plus_base.proto",
}
