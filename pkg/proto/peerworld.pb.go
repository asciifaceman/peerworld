// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peerworld.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	peerworld.proto

It has these top-level messages:
	IntroductionMessage
*/
package main

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

type IntroductionMessage struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	Port int32  `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
}

func (m *IntroductionMessage) Reset()                    { *m = IntroductionMessage{} }
func (m *IntroductionMessage) String() string            { return proto.CompactTextString(m) }
func (*IntroductionMessage) ProtoMessage()               {}
func (*IntroductionMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IntroductionMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IntroductionMessage) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *IntroductionMessage) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*IntroductionMessage)(nil), "main.IntroductionMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Neighborly service

type NeighborlyClient interface {
	Introduce(ctx context.Context, in *IntroductionMessage, opts ...grpc.CallOption) (*IntroductionMessage, error)
}

type neighborlyClient struct {
	cc *grpc.ClientConn
}

func NewNeighborlyClient(cc *grpc.ClientConn) NeighborlyClient {
	return &neighborlyClient{cc}
}

func (c *neighborlyClient) Introduce(ctx context.Context, in *IntroductionMessage, opts ...grpc.CallOption) (*IntroductionMessage, error) {
	out := new(IntroductionMessage)
	err := grpc.Invoke(ctx, "/main.Neighborly/Introduce", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Neighborly service

type NeighborlyServer interface {
	Introduce(context.Context, *IntroductionMessage) (*IntroductionMessage, error)
}

func RegisterNeighborlyServer(s *grpc.Server, srv NeighborlyServer) {
	s.RegisterService(&_Neighborly_serviceDesc, srv)
}

func _Neighborly_Introduce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntroductionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NeighborlyServer).Introduce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Neighborly/Introduce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NeighborlyServer).Introduce(ctx, req.(*IntroductionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Neighborly_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.Neighborly",
	HandlerType: (*NeighborlyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Introduce",
			Handler:    _Neighborly_Introduce_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peerworld.proto",
}

func init() { proto.RegisterFile("peerworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x48, 0x4d, 0x2d,
	0x2a, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc,
	0xcc, 0x53, 0x0a, 0xe4, 0x12, 0xf6, 0xcc, 0x2b, 0x29, 0xca, 0x4f, 0x29, 0x4d, 0x2e, 0xc9, 0xcc,
	0xcf, 0xf3, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x41, 0x62, 0x19, 0xf9, 0xc5, 0x25, 0x12,
	0x4c, 0x10, 0x31, 0x10, 0x1b, 0x24, 0x56, 0x90, 0x5f, 0x54, 0x22, 0xc1, 0xac, 0xc0, 0xa8, 0xc1,
	0x1a, 0x04, 0x66, 0x1b, 0x05, 0x72, 0x71, 0xf9, 0xa5, 0x66, 0xa6, 0x67, 0x24, 0xe5, 0x17, 0xe5,
	0x54, 0x0a, 0x39, 0x73, 0x71, 0xc2, 0x2c, 0x48, 0x15, 0x92, 0xd4, 0x03, 0x59, 0xaa, 0x87, 0xc5,
	0x46, 0x29, 0xdc, 0x52, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x27, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x03, 0x0d, 0x94, 0xa0, 0xc5, 0x00, 0x00, 0x00,
}
