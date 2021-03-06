// Code generated by protoc-gen-go. DO NOT EDIT.
// source: injure.proto

package injure

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Request struct {
	Methods              []string `protobuf:"bytes,1,rep,name=methods,proto3" json:"methods,omitempty"`
	Errno                uint32   `protobuf:"varint,2,opt,name=errno,proto3" json:"errno,omitempty"`
	Random               bool     `protobuf:"varint,3,opt,name=random,proto3" json:"random,omitempty"`
	Pct                  uint32   `protobuf:"varint,4,opt,name=pct,proto3" json:"pct,omitempty"`
	Path                 string   `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	Delay                uint32   `protobuf:"varint,6,opt,name=delay,proto3" json:"delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_534a76e04d5e385a, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

func (m *Request) GetErrno() uint32 {
	if m != nil {
		return m.Errno
	}
	return 0
}

func (m *Request) GetRandom() bool {
	if m != nil {
		return m.Random
	}
	return false
}

func (m *Request) GetPct() uint32 {
	if m != nil {
		return m.Pct
	}
	return 0
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Request) GetDelay() uint32 {
	if m != nil {
		return m.Delay
	}
	return 0
}

type Response struct {
	Methods              []string `protobuf:"bytes,1,rep,name=methods,proto3" json:"methods,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_534a76e04d5e385a, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

type InjectedResponse struct {
	Injected             bool     `protobuf:"varint,1,opt,name=injected,proto3" json:"injected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InjectedResponse) Reset()         { *m = InjectedResponse{} }
func (m *InjectedResponse) String() string { return proto.CompactTextString(m) }
func (*InjectedResponse) ProtoMessage()    {}
func (*InjectedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_534a76e04d5e385a, []int{2}
}

func (m *InjectedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InjectedResponse.Unmarshal(m, b)
}
func (m *InjectedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InjectedResponse.Marshal(b, m, deterministic)
}
func (m *InjectedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InjectedResponse.Merge(m, src)
}
func (m *InjectedResponse) XXX_Size() int {
	return xxx_messageInfo_InjectedResponse.Size(m)
}
func (m *InjectedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InjectedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InjectedResponse proto.InternalMessageInfo

func (m *InjectedResponse) GetInjected() bool {
	if m != nil {
		return m.Injected
	}
	return false
}

func init() {
	proto.RegisterType((*Request)(nil), "injure.Request")
	proto.RegisterType((*Response)(nil), "injure.Response")
	proto.RegisterType((*InjectedResponse)(nil), "injure.InjectedResponse")
}

func init() { proto.RegisterFile("injure.proto", fileDescriptor_534a76e04d5e385a) }

var fileDescriptor_534a76e04d5e385a = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xd7, 0xfd, 0xe9, 0xba, 0x57, 0x87, 0x23, 0xc8, 0x08, 0xf3, 0x52, 0x82, 0x87, 0x9e,
	0x32, 0x50, 0x06, 0xe2, 0x61, 0xe0, 0x41, 0x61, 0x07, 0x2f, 0xf1, 0x13, 0x74, 0xeb, 0xeb, 0xfe,
	0xd0, 0x36, 0x35, 0x4d, 0x85, 0x7d, 0x08, 0xbf, 0xa9, 0x1f, 0x42, 0x9a, 0x34, 0x13, 0x84, 0x0a,
	0xbb, 0xe5, 0x79, 0x79, 0x9e, 0x27, 0x6f, 0x7e, 0x81, 0xcb, 0x7d, 0x7e, 0xa8, 0x14, 0xf2, 0x42,
	0x49, 0x2d, 0x89, 0x6f, 0xd5, 0xec, 0x66, 0x2b, 0xe5, 0x36, 0xc5, 0xb9, 0x99, 0xae, 0xab, 0xf7,
	0x39, 0x66, 0x85, 0x3e, 0x5a, 0x13, 0xfb, 0xf2, 0x60, 0x28, 0xf0, 0xa3, 0xc2, 0x52, 0x13, 0x0a,
	0xc3, 0x0c, 0xf5, 0x4e, 0x26, 0x25, 0xf5, 0xc2, 0x5e, 0x34, 0x12, 0x4e, 0x92, 0x6b, 0x18, 0xa0,
	0x52, 0xb9, 0xa4, 0xdd, 0xd0, 0x8b, 0xc6, 0xc2, 0x0a, 0x32, 0x05, 0x5f, 0xc5, 0x79, 0x22, 0x33,
	0xda, 0x0b, 0xbd, 0x28, 0x10, 0x8d, 0x22, 0x13, 0xe8, 0x15, 0x1b, 0x4d, 0xfb, 0xc6, 0x5b, 0x1f,
	0x09, 0x81, 0x7e, 0x11, 0xeb, 0x1d, 0x1d, 0x84, 0x5e, 0x34, 0x12, 0xe6, 0x5c, 0x77, 0x26, 0x98,
	0xc6, 0x47, 0xea, 0xdb, 0x4e, 0x23, 0xd8, 0x2d, 0x04, 0x02, 0xcb, 0x42, 0xe6, 0x25, 0xb6, 0xef,
	0xc3, 0x38, 0x4c, 0x56, 0xf9, 0x01, 0x37, 0x1a, 0x93, 0x93, 0x7b, 0x06, 0xc1, 0xbe, 0x99, 0x51,
	0xcf, 0xec, 0x73, 0xd2, 0x77, 0xdf, 0x5d, 0xf0, 0x57, 0x86, 0x06, 0x59, 0xc0, 0xf0, 0xb5, 0x79,
	0xd5, 0x94, 0x5b, 0x32, 0xdc, 0x91, 0xe1, 0xcf, 0x35, 0x99, 0xd9, 0x84, 0x37, 0x1c, 0x5d, 0x37,
	0xeb, 0x90, 0x25, 0x80, 0xc0, 0x8d, 0xfc, 0x44, 0xf5, 0x94, 0xa6, 0xad, 0xc9, 0x96, 0x39, 0xeb,
	0x90, 0x47, 0x18, 0x37, 0x79, 0x7b, 0x3b, 0xb9, 0xfa, 0xbd, 0xc4, 0xd0, 0xff, 0x27, 0xbb, 0x80,
	0xe0, 0x0d, 0xf5, 0x4b, 0x5c, 0xa5, 0xfa, 0x9c, 0xd8, 0x03, 0x5c, 0xb8, 0x58, 0xbd, 0xf3, 0x19,
	0xc9, 0x25, 0x04, 0x0e, 0x6f, 0xeb, 0x53, 0xa9, 0xab, 0xfb, 0xfb, 0x11, 0xac, 0xb3, 0xf6, 0x8d,
	0xf7, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x81, 0x47, 0x82, 0x90, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InjureClient is the client API for Injure service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InjureClient interface {
	Methods(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Response, error)
	RecoverAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	RecoverMethod(ctx context.Context, in *Request, opts ...grpc.CallOption) (*empty.Empty, error)
	SetFault(ctx context.Context, in *Request, opts ...grpc.CallOption) (*empty.Empty, error)
	SetFaultAll(ctx context.Context, in *Request, opts ...grpc.CallOption) (*empty.Empty, error)
	Injected(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*InjectedResponse, error)
}

type injureClient struct {
	cc *grpc.ClientConn
}

func NewInjureClient(cc *grpc.ClientConn) InjureClient {
	return &injureClient{cc}
}

func (c *injureClient) Methods(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/injure.Injure/Methods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injureClient) RecoverAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/injure.Injure/RecoverAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injureClient) RecoverMethod(ctx context.Context, in *Request, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/injure.Injure/RecoverMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injureClient) SetFault(ctx context.Context, in *Request, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/injure.Injure/SetFault", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injureClient) SetFaultAll(ctx context.Context, in *Request, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/injure.Injure/SetFaultAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injureClient) Injected(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*InjectedResponse, error) {
	out := new(InjectedResponse)
	err := c.cc.Invoke(ctx, "/injure.Injure/Injected", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InjureServer is the server API for Injure service.
type InjureServer interface {
	Methods(context.Context, *empty.Empty) (*Response, error)
	RecoverAll(context.Context, *empty.Empty) (*empty.Empty, error)
	RecoverMethod(context.Context, *Request) (*empty.Empty, error)
	SetFault(context.Context, *Request) (*empty.Empty, error)
	SetFaultAll(context.Context, *Request) (*empty.Empty, error)
	Injected(context.Context, *empty.Empty) (*InjectedResponse, error)
}

// UnimplementedInjureServer can be embedded to have forward compatible implementations.
type UnimplementedInjureServer struct {
}

func (*UnimplementedInjureServer) Methods(ctx context.Context, req *empty.Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Methods not implemented")
}
func (*UnimplementedInjureServer) RecoverAll(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverAll not implemented")
}
func (*UnimplementedInjureServer) RecoverMethod(ctx context.Context, req *Request) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverMethod not implemented")
}
func (*UnimplementedInjureServer) SetFault(ctx context.Context, req *Request) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFault not implemented")
}
func (*UnimplementedInjureServer) SetFaultAll(ctx context.Context, req *Request) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFaultAll not implemented")
}
func (*UnimplementedInjureServer) Injected(ctx context.Context, req *empty.Empty) (*InjectedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Injected not implemented")
}

func RegisterInjureServer(s *grpc.Server, srv InjureServer) {
	s.RegisterService(&_Injure_serviceDesc, srv)
}

func _Injure_Methods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjureServer).Methods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injure.Injure/Methods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjureServer).Methods(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Injure_RecoverAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjureServer).RecoverAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injure.Injure/RecoverAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjureServer).RecoverAll(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Injure_RecoverMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjureServer).RecoverMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injure.Injure/RecoverMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjureServer).RecoverMethod(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Injure_SetFault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjureServer).SetFault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injure.Injure/SetFault",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjureServer).SetFault(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Injure_SetFaultAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjureServer).SetFaultAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injure.Injure/SetFaultAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjureServer).SetFaultAll(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Injure_Injected_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjureServer).Injected(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injure.Injure/Injected",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjureServer).Injected(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Injure_serviceDesc = grpc.ServiceDesc{
	ServiceName: "injure.Injure",
	HandlerType: (*InjureServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Methods",
			Handler:    _Injure_Methods_Handler,
		},
		{
			MethodName: "RecoverAll",
			Handler:    _Injure_RecoverAll_Handler,
		},
		{
			MethodName: "RecoverMethod",
			Handler:    _Injure_RecoverMethod_Handler,
		},
		{
			MethodName: "SetFault",
			Handler:    _Injure_SetFault_Handler,
		},
		{
			MethodName: "SetFaultAll",
			Handler:    _Injure_SetFaultAll_Handler,
		},
		{
			MethodName: "Injected",
			Handler:    _Injure_Injected_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "injure.proto",
}
