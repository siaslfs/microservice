// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app.proto

// 定义包名

package app

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 定义 Res 消息结构
type Params struct {
	Keys                 string   `protobuf:"bytes,1,opt,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{0}
}

func (m *Params) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Params.Unmarshal(m, b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Params.Marshal(b, m, deterministic)
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return xxx_messageInfo_Params.Size(m)
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetKeys() string {
	if m != nil {
		return m.Keys
	}
	return ""
}

type Result struct {
	Code                 int32    `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Data                 string   `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{1}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Result) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "app.Params")
	proto.RegisterType((*Result)(nil), "app.Result")
}

func init() { proto.RegisterFile("app.proto", fileDescriptor_e0f9056a14b86d47) }

var fileDescriptor_e0f9056a14b86d47 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0x28, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0x28, 0x50, 0x92, 0xe1, 0x62, 0x0b, 0x48,
	0x2c, 0x4a, 0xcc, 0x2d, 0x16, 0x12, 0xe2, 0x62, 0xc9, 0x4e, 0xad, 0x2c, 0x96, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x0c, 0xb8, 0xd8, 0x82, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0x40,
	0xb2, 0xce, 0xf9, 0x29, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x60, 0x36, 0x48, 0xcc,
	0x25, 0xb1, 0x24, 0x11, 0xa6, 0x03, 0xc4, 0x36, 0x32, 0xe1, 0x62, 0x0b, 0x4f, 0xcc, 0x2c, 0x49,
	0x2d, 0x12, 0xd2, 0xe2, 0xe2, 0x71, 0x4f, 0x2d, 0x71, 0x2c, 0x28, 0x70, 0xce, 0xcf, 0x4b, 0xcb,
	0x4c, 0x17, 0xe2, 0xd6, 0x03, 0x59, 0x0d, 0xb1, 0x4c, 0x0a, 0xc2, 0x81, 0x98, 0xad, 0xc4, 0x90,
	0xc4, 0x06, 0x76, 0x91, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x05, 0xf5, 0x45, 0x9e, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WaiterClient is the client API for Waiter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WaiterClient interface {
	// 定义接口 (结构体可以复用)
	// 方法 (请求消息结构体) returns (返回消息结构体) {}
	GetAppConfig(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Result, error)
}

type waiterClient struct {
	cc *grpc.ClientConn
}

func NewWaiterClient(cc *grpc.ClientConn) WaiterClient {
	return &waiterClient{cc}
}

func (c *waiterClient) GetAppConfig(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/app.Waiter/GetAppConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WaiterServer is the server API for Waiter service.
type WaiterServer interface {
	// 定义接口 (结构体可以复用)
	// 方法 (请求消息结构体) returns (返回消息结构体) {}
	GetAppConfig(context.Context, *Params) (*Result, error)
}

// UnimplementedWaiterServer can be embedded to have forward compatible implementations.
type UnimplementedWaiterServer struct {
}

func (*UnimplementedWaiterServer) GetAppConfig(ctx context.Context, req *Params) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppConfig not implemented")
}

func RegisterWaiterServer(s *grpc.Server, srv WaiterServer) {
	s.RegisterService(&_Waiter_serviceDesc, srv)
}

func _Waiter_GetAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaiterServer).GetAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.Waiter/GetAppConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaiterServer).GetAppConfig(ctx, req.(*Params))
	}
	return interceptor(ctx, in, info, handler)
}

var _Waiter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "app.Waiter",
	HandlerType: (*WaiterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAppConfig",
			Handler:    _Waiter_GetAppConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}
