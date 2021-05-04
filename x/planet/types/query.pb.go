// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: planet/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("planet/query.proto", fileDescriptor_2b55d081339b1f1e) }

var fileDescriptor_2b55d081339b1f1e = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xb1, 0x0e, 0x82, 0x30,
	0x10, 0x40, 0x61, 0x50, 0x13, 0x46, 0x46, 0x62, 0xba, 0xb9, 0x38, 0x70, 0x41, 0x3f, 0xc0, 0xc4,
	0x3f, 0x70, 0x75, 0x6b, 0xc9, 0xa5, 0x36, 0x81, 0x5e, 0xa5, 0x87, 0x91, 0xbf, 0xf0, 0xb3, 0x1c,
	0x19, 0x1d, 0x0d, 0xfc, 0x88, 0x81, 0xe2, 0x74, 0x97, 0xdc, 0xbb, 0x97, 0x97, 0xa4, 0xae, 0x92,
	0x16, 0x19, 0xee, 0x2d, 0x36, 0x5d, 0xee, 0x1a, 0x62, 0x4a, 0x53, 0x46, 0xcf, 0x79, 0x38, 0x2c,
	0x23, 0xdb, 0x6a, 0x22, 0x5d, 0x21, 0x48, 0x67, 0x40, 0x5a, 0x4b, 0x2c, 0xd9, 0x90, 0xf5, 0xe1,
	0x23, 0xdb, 0x97, 0xe4, 0x6b, 0xf2, 0xa0, 0xa4, 0xc7, 0xa0, 0x82, 0x47, 0xa1, 0x90, 0x65, 0x01,
	0x4e, 0x6a, 0x63, 0x67, 0x38, 0xb0, 0x87, 0x4d, 0xb2, 0xba, 0x4c, 0xc4, 0xf9, 0xf4, 0x1e, 0x44,
	0xdc, 0x0f, 0x22, 0xfe, 0x0e, 0x22, 0x7e, 0x8d, 0x22, 0xea, 0x47, 0x11, 0x7d, 0x46, 0x11, 0x5d,
	0x77, 0xda, 0xf0, 0xad, 0x55, 0x79, 0x49, 0x35, 0x4c, 0x2d, 0xb0, 0x44, 0x3e, 0xff, 0x0b, 0x77,
	0x0e, 0xbd, 0x5a, 0xcf, 0xc2, 0xe3, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x71, 0x86, 0xf8, 0x3c, 0xc4,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.planet.planet.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "planet/query.proto",
}
