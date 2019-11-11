// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/grpc.proto

package grpc

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

type SortRequest struct {
	TableToSort          []int32  `protobuf:"varint,1,rep,packed,name=tableToSort,proto3" json:"tableToSort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SortRequest) Reset()         { *m = SortRequest{} }
func (m *SortRequest) String() string { return proto.CompactTextString(m) }
func (*SortRequest) ProtoMessage()    {}
func (*SortRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df467ee3cbdf4524, []int{0}
}

func (m *SortRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SortRequest.Unmarshal(m, b)
}
func (m *SortRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SortRequest.Marshal(b, m, deterministic)
}
func (m *SortRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SortRequest.Merge(m, src)
}
func (m *SortRequest) XXX_Size() int {
	return xxx_messageInfo_SortRequest.Size(m)
}
func (m *SortRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SortRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SortRequest proto.InternalMessageInfo

func (m *SortRequest) GetTableToSort() []int32 {
	if m != nil {
		return m.TableToSort
	}
	return nil
}

type Response struct {
	Sorted               bool     `protobuf:"varint,1,opt,name=sorted,proto3" json:"sorted,omitempty"`
	SortedTable          []int32  `protobuf:"varint,2,rep,packed,name=sortedTable,proto3" json:"sortedTable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_df467ee3cbdf4524, []int{1}
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

func (m *Response) GetSorted() bool {
	if m != nil {
		return m.Sorted
	}
	return false
}

func (m *Response) GetSortedTable() []int32 {
	if m != nil {
		return m.SortedTable
	}
	return nil
}

func init() {
	proto.RegisterType((*SortRequest)(nil), "grpc.SortRequest")
	proto.RegisterType((*Response)(nil), "grpc.Response")
}

func init() { proto.RegisterFile("proto/grpc.proto", fileDescriptor_df467ee3cbdf4524) }

var fileDescriptor_df467ee3cbdf4524 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x03, 0x33, 0x85, 0x58, 0x40, 0x6c, 0x25, 0x7d, 0x2e,
	0xee, 0xe0, 0xfc, 0xa2, 0x92, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x05, 0x2e, 0xee,
	0x92, 0xc4, 0xa4, 0x9c, 0xd4, 0x90, 0x7c, 0x90, 0xa8, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0x6b, 0x10,
	0xb2, 0x90, 0x92, 0x0b, 0x17, 0x47, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x18,
	0x17, 0x5b, 0x71, 0x7e, 0x51, 0x49, 0x6a, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x94,
	0x07, 0x32, 0x05, 0xc2, 0x0a, 0x01, 0x69, 0x94, 0x60, 0x82, 0x98, 0x82, 0x24, 0x64, 0x64, 0xcb,
	0xc5, 0x07, 0x32, 0x2d, 0x33, 0x2f, 0x3d, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x48, 0x9b,
	0x8b, 0x05, 0x24, 0x22, 0x24, 0xa8, 0x07, 0x76, 0x23, 0x92, 0xa3, 0xa4, 0xf8, 0x20, 0x42, 0x30,
	0x6b, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0x5e, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xcc,
	0x30, 0x5f, 0xd6, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SortingServiceClient is the client API for SortingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SortingServiceClient interface {
	Sort(ctx context.Context, in *SortRequest, opts ...grpc.CallOption) (*Response, error)
}

type sortingServiceClient struct {
	cc *grpc.ClientConn
}

func NewSortingServiceClient(cc *grpc.ClientConn) SortingServiceClient {
	return &sortingServiceClient{cc}
}

func (c *sortingServiceClient) Sort(ctx context.Context, in *SortRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpc.SortingService/Sort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SortingServiceServer is the server API for SortingService service.
type SortingServiceServer interface {
	Sort(context.Context, *SortRequest) (*Response, error)
}

// UnimplementedSortingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSortingServiceServer struct {
}

func (*UnimplementedSortingServiceServer) Sort(ctx context.Context, req *SortRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sort not implemented")
}

func RegisterSortingServiceServer(s *grpc.Server, srv SortingServiceServer) {
	s.RegisterService(&_SortingService_serviceDesc, srv)
}

func _SortingService_Sort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SortingServiceServer).Sort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.SortingService/Sort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SortingServiceServer).Sort(ctx, req.(*SortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SortingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.SortingService",
	HandlerType: (*SortingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sort",
			Handler:    _SortingService_Sort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/grpc.proto",
}
