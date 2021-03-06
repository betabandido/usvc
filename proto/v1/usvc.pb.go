// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usvc.proto

package v1

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

type Count struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Count) Reset()         { *m = Count{} }
func (m *Count) String() string { return proto.CompactTextString(m) }
func (*Count) ProtoMessage()    {}
func (*Count) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bb0a830d52b622e, []int{0}
}

func (m *Count) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Count.Unmarshal(m, b)
}
func (m *Count) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Count.Marshal(b, m, deterministic)
}
func (m *Count) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Count.Merge(m, src)
}
func (m *Count) XXX_Size() int {
	return xxx_messageInfo_Count.Size(m)
}
func (m *Count) XXX_DiscardUnknown() {
	xxx_messageInfo_Count.DiscardUnknown(m)
}

var xxx_messageInfo_Count proto.InternalMessageInfo

func (m *Count) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type GetCountRequest struct {
	AddDelay             bool     `protobuf:"varint,1,opt,name=addDelay,proto3" json:"addDelay,omitempty"`
	ErrorRate            int32    `protobuf:"varint,2,opt,name=errorRate,proto3" json:"errorRate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCountRequest) Reset()         { *m = GetCountRequest{} }
func (m *GetCountRequest) String() string { return proto.CompactTextString(m) }
func (*GetCountRequest) ProtoMessage()    {}
func (*GetCountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bb0a830d52b622e, []int{1}
}

func (m *GetCountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCountRequest.Unmarshal(m, b)
}
func (m *GetCountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCountRequest.Marshal(b, m, deterministic)
}
func (m *GetCountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCountRequest.Merge(m, src)
}
func (m *GetCountRequest) XXX_Size() int {
	return xxx_messageInfo_GetCountRequest.Size(m)
}
func (m *GetCountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCountRequest proto.InternalMessageInfo

func (m *GetCountRequest) GetAddDelay() bool {
	if m != nil {
		return m.AddDelay
	}
	return false
}

func (m *GetCountRequest) GetErrorRate() int32 {
	if m != nil {
		return m.ErrorRate
	}
	return 0
}

type UpdateCountRequest struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCountRequest) Reset()         { *m = UpdateCountRequest{} }
func (m *UpdateCountRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCountRequest) ProtoMessage()    {}
func (*UpdateCountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bb0a830d52b622e, []int{2}
}

func (m *UpdateCountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCountRequest.Unmarshal(m, b)
}
func (m *UpdateCountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCountRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCountRequest.Merge(m, src)
}
func (m *UpdateCountRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCountRequest.Size(m)
}
func (m *UpdateCountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCountRequest proto.InternalMessageInfo

func (m *UpdateCountRequest) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Count)(nil), "v1.Count")
	proto.RegisterType((*GetCountRequest)(nil), "v1.GetCountRequest")
	proto.RegisterType((*UpdateCountRequest)(nil), "v1.UpdateCountRequest")
}

func init() { proto.RegisterFile("usvc.proto", fileDescriptor_2bb0a830d52b622e) }

var fileDescriptor_2bb0a830d52b622e = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x2e, 0x4b,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x54, 0x92, 0xe5, 0x62, 0x75, 0xce,
	0x2f, 0xcd, 0x2b, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60,
	0xd4, 0x60, 0x0e, 0x82, 0x70, 0x94, 0xbc, 0xb9, 0xf8, 0xdd, 0x53, 0x4b, 0xc0, 0x2a, 0x82, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xa4, 0xb8, 0x38, 0x12, 0x53, 0x52, 0x5c, 0x52, 0x73, 0x12,
	0x2b, 0xc1, 0x6a, 0x39, 0x82, 0xe0, 0x7c, 0x21, 0x19, 0x2e, 0xce, 0xd4, 0xa2, 0xa2, 0xfc, 0xa2,
	0xa0, 0xc4, 0x92, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x84, 0x80, 0x92, 0x16, 0x97,
	0x50, 0x68, 0x41, 0x4a, 0x62, 0x49, 0x2a, 0x8a, 0x79, 0x58, 0x2d, 0x36, 0x2a, 0xe0, 0xe2, 0x01,
	0xab, 0x0a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0xd2, 0xe1, 0xe2, 0x80, 0x39, 0x44, 0x48,
	0x58, 0xaf, 0xcc, 0x50, 0x0f, 0xcd, 0x59, 0x52, 0x9c, 0x20, 0x41, 0xb0, 0x88, 0x12, 0x83, 0x90,
	0x11, 0x17, 0x37, 0x92, 0x4d, 0x42, 0x62, 0x20, 0x39, 0x4c, 0xab, 0x51, 0xf4, 0x24, 0xb1, 0x81,
	0x03, 0xc5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xca, 0x6c, 0x17, 0xb3, 0x22, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CountServiceClient is the client API for CountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CountServiceClient interface {
	GetCount(ctx context.Context, in *GetCountRequest, opts ...grpc.CallOption) (*Count, error)
	UpdateCount(ctx context.Context, in *UpdateCountRequest, opts ...grpc.CallOption) (*Count, error)
}

type countServiceClient struct {
	cc *grpc.ClientConn
}

func NewCountServiceClient(cc *grpc.ClientConn) CountServiceClient {
	return &countServiceClient{cc}
}

func (c *countServiceClient) GetCount(ctx context.Context, in *GetCountRequest, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/v1.CountService/GetCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countServiceClient) UpdateCount(ctx context.Context, in *UpdateCountRequest, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/v1.CountService/UpdateCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountServiceServer is the server API for CountService service.
type CountServiceServer interface {
	GetCount(context.Context, *GetCountRequest) (*Count, error)
	UpdateCount(context.Context, *UpdateCountRequest) (*Count, error)
}

// UnimplementedCountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCountServiceServer struct {
}

func (*UnimplementedCountServiceServer) GetCount(ctx context.Context, req *GetCountRequest) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCount not implemented")
}
func (*UnimplementedCountServiceServer) UpdateCount(ctx context.Context, req *UpdateCountRequest) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCount not implemented")
}

func RegisterCountServiceServer(s *grpc.Server, srv CountServiceServer) {
	s.RegisterService(&_CountService_serviceDesc, srv)
}

func _CountService_GetCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountServiceServer).GetCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CountService/GetCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountServiceServer).GetCount(ctx, req.(*GetCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountService_UpdateCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountServiceServer).UpdateCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CountService/UpdateCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountServiceServer).UpdateCount(ctx, req.(*UpdateCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.CountService",
	HandlerType: (*CountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCount",
			Handler:    _CountService_GetCount_Handler,
		},
		{
			MethodName: "UpdateCount",
			Handler:    _CountService_UpdateCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usvc.proto",
}
