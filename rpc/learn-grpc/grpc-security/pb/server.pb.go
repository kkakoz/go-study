// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type SimpleRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleRequest) Reset()         { *m = SimpleRequest{} }
func (m *SimpleRequest) String() string { return proto.CompactTextString(m) }
func (*SimpleRequest) ProtoMessage()    {}
func (*SimpleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *SimpleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleRequest.Unmarshal(m, b)
}
func (m *SimpleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleRequest.Marshal(b, m, deterministic)
}
func (m *SimpleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleRequest.Merge(m, src)
}
func (m *SimpleRequest) XXX_Size() int {
	return xxx_messageInfo_SimpleRequest.Size(m)
}
func (m *SimpleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleRequest proto.InternalMessageInfo

func (m *SimpleRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type SimpleResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleResponse) Reset()         { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()    {}
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *SimpleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleResponse.Unmarshal(m, b)
}
func (m *SimpleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleResponse.Marshal(b, m, deterministic)
}
func (m *SimpleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleResponse.Merge(m, src)
}
func (m *SimpleResponse) XXX_Size() int {
	return xxx_messageInfo_SimpleResponse.Size(m)
}
func (m *SimpleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleResponse proto.InternalMessageInfo

func (m *SimpleResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*SimpleRequest)(nil), "SimpleRequest")
	proto.RegisterType((*SimpleResponse)(nil), "SimpleResponse")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe6, 0xe2, 0x0d, 0xce, 0xcc, 0x2d,
	0xc8, 0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c,
	0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0xd4, 0xb8, 0xf8, 0x60, 0x8a,
	0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0xa1,
	0xca, 0x20, 0x1c, 0x23, 0x1b, 0x2e, 0x3e, 0x97, 0xd4, 0xc4, 0x94, 0x9c, 0xcc, 0xbc, 0xd4, 0x60,
	0xb0, 0x25, 0x42, 0x5a, 0x5c, 0xac, 0x41, 0xf9, 0xa5, 0x25, 0xa9, 0x42, 0x7c, 0x7a, 0x28, 0xd6,
	0x48, 0xf1, 0xeb, 0xa1, 0x9a, 0xa8, 0xc4, 0xe0, 0xc4, 0x1e, 0xc5, 0xaa, 0xa7, 0x6f, 0x5d, 0x90,
	0x94, 0xc4, 0x06, 0x76, 0x9a, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xff, 0xfa, 0x36, 0x46, 0xaa,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeadlineServerClient is the client API for DeadlineServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeadlineServerClient interface {
	Route(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
}

type deadlineServerClient struct {
	cc *grpc.ClientConn
}

func NewDeadlineServerClient(cc *grpc.ClientConn) DeadlineServerClient {
	return &deadlineServerClient{cc}
}

func (c *deadlineServerClient) Route(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/DeadlineServer/Route", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeadlineServerServer is the server API for DeadlineServer service.
type DeadlineServerServer interface {
	Route(context.Context, *SimpleRequest) (*SimpleResponse, error)
}

func RegisterDeadlineServerServer(s *grpc.Server, srv DeadlineServerServer) {
	s.RegisterService(&_DeadlineServer_serviceDesc, srv)
}

func _DeadlineServer_Route_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeadlineServerServer).Route(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DeadlineServer/Route",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeadlineServerServer).Route(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeadlineServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DeadlineServer",
	HandlerType: (*DeadlineServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Route",
			Handler:    _DeadlineServer_Route_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
