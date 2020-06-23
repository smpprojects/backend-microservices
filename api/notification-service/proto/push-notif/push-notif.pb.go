// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/notification-service/proto/push-notif/push-notif.proto

package pushnotif

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_notif_dbe0c72e593a62d2, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Response struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_notif_dbe0c72e593a62d2, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "pushnotif.Request")
	proto.RegisterType((*Response)(nil), "pushnotif.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PushNotifHandlerClient is the client API for PushNotifHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PushNotifHandlerClient interface {
	Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type pushNotifHandlerClient struct {
	cc *grpc.ClientConn
}

func NewPushNotifHandlerClient(cc *grpc.ClientConn) PushNotifHandlerClient {
	return &pushNotifHandlerClient{cc}
}

func (c *pushNotifHandlerClient) Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pushnotif.PushNotifHandler/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushNotifHandlerServer is the server API for PushNotifHandler service.
type PushNotifHandlerServer interface {
	Hello(context.Context, *Request) (*Response, error)
}

func RegisterPushNotifHandlerServer(s *grpc.Server, srv PushNotifHandlerServer) {
	s.RegisterService(&_PushNotifHandler_serviceDesc, srv)
}

func _PushNotifHandler_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushNotifHandlerServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pushnotif.PushNotifHandler/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushNotifHandlerServer).Hello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _PushNotifHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pushnotif.PushNotifHandler",
	HandlerType: (*PushNotifHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _PushNotifHandler_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/notification-service/proto/push-notif/push-notif.proto",
}

func init() {
	proto.RegisterFile("api/notification-service/proto/push-notif/push-notif.proto", fileDescriptor_push_notif_dbe0c72e593a62d2)
}

var fileDescriptor_push_notif_dbe0c72e593a62d2 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4a, 0x2c, 0xc8, 0xd4,
	0xcf, 0xcb, 0x2f, 0xc9, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2d, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x28, 0x2d, 0xce, 0xd0,
	0x05, 0xcb, 0x23, 0x31, 0xf5, 0xc0, 0x72, 0x42, 0x9c, 0x20, 0x11, 0xb0, 0x80, 0x92, 0x32, 0x17,
	0x7b, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x04, 0x17, 0xbb, 0x6f, 0x6a, 0x71, 0x71,
	0x62, 0x7a, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0xab, 0xa4, 0xc2, 0xc5, 0x11,
	0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x8a, 0x5b, 0x95, 0x91, 0x0b, 0x97, 0x40, 0x40, 0x69,
	0x71, 0x86, 0x1f, 0xc8, 0x5c, 0x8f, 0xc4, 0xbc, 0x94, 0x9c, 0xd4, 0x22, 0x21, 0x03, 0x2e, 0x56,
	0x8f, 0xd4, 0x9c, 0x9c, 0x7c, 0x21, 0x21, 0x3d, 0xb8, 0x9d, 0x7a, 0x50, 0x0b, 0xa5, 0x84, 0x51,
	0xc4, 0x20, 0xe6, 0x27, 0xb1, 0x81, 0x9d, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x6d,
	0x54, 0x30, 0xe0, 0x00, 0x00, 0x00,
}
