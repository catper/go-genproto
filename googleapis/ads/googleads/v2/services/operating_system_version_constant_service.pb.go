// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/operating_system_version_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request message for
// [OperatingSystemVersionConstantService.GetOperatingSystemVersionConstant][google.ads.googleads.v2.services.OperatingSystemVersionConstantService.GetOperatingSystemVersionConstant].
type GetOperatingSystemVersionConstantRequest struct {
	// Required. Resource name of the OS version to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOperatingSystemVersionConstantRequest) Reset() {
	*m = GetOperatingSystemVersionConstantRequest{}
}
func (m *GetOperatingSystemVersionConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetOperatingSystemVersionConstantRequest) ProtoMessage()    {}
func (*GetOperatingSystemVersionConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_02bc5000ca7d978c, []int{0}
}

func (m *GetOperatingSystemVersionConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Unmarshal(m, b)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Merge(m, src)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Size(m)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperatingSystemVersionConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperatingSystemVersionConstantRequest proto.InternalMessageInfo

func (m *GetOperatingSystemVersionConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetOperatingSystemVersionConstantRequest)(nil), "google.ads.googleads.v2.services.GetOperatingSystemVersionConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/operating_system_version_constant_service.proto", fileDescriptor_02bc5000ca7d978c)
}

var fileDescriptor_02bc5000ca7d978c = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x69, 0x16, 0x04, 0x83, 0x5e, 0x72, 0x71, 0xa9, 0x82, 0x75, 0x59, 0xa1, 0xec, 0x61,
	0x06, 0x22, 0x22, 0x8c, 0x78, 0x98, 0x7a, 0xa8, 0x0a, 0x6a, 0xd9, 0x95, 0x1e, 0xa4, 0x10, 0x66,
	0x93, 0xd7, 0x38, 0xd0, 0xcc, 0x5b, 0xe7, 0x9d, 0x0d, 0x88, 0x78, 0xd0, 0x6f, 0x20, 0x7e, 0x03,
	0x8f, 0x7e, 0x12, 0xd9, 0xab, 0x37, 0x4f, 0x1e, 0x3c, 0xf9, 0x29, 0x96, 0x74, 0x32, 0xd9, 0xed,
	0xa1, 0x7f, 0x6e, 0x0f, 0x7d, 0x9f, 0xfe, 0xde, 0x27, 0xcf, 0x9b, 0xc4, 0x93, 0x12, 0xb1, 0x9c,
	0x03, 0x57, 0x05, 0x71, 0x2f, 0x1b, 0x55, 0xa7, 0x9c, 0xc0, 0xd6, 0x3a, 0x07, 0xe2, 0xb8, 0x00,
	0xab, 0x9c, 0x36, 0x65, 0x46, 0x1f, 0xc9, 0x41, 0x95, 0xd5, 0x60, 0x49, 0xa3, 0xc9, 0x72, 0x34,
	0xe4, 0x94, 0x71, 0x59, 0x6b, 0x65, 0x0b, 0x8b, 0x0e, 0x93, 0x81, 0xc7, 0x30, 0x55, 0x10, 0xeb,
	0x88, 0xac, 0x4e, 0x59, 0x20, 0xf6, 0x9f, 0xaf, 0xdb, 0x69, 0x81, 0xf0, 0xcc, 0xee, 0xb4, 0xd4,
	0x2f, 0xeb, 0xdf, 0x09, 0xa8, 0x85, 0xe6, 0xca, 0x18, 0x74, 0xca, 0x69, 0x34, 0xd4, 0x4e, 0x6f,
	0x5d, 0x99, 0xe6, 0x73, 0x0d, 0xdd, 0xdf, 0xee, 0x5e, 0x19, 0xbc, 0xd3, 0x30, 0x2f, 0xb2, 0x53,
	0x78, 0xaf, 0x6a, 0x8d, 0xd6, 0x1b, 0x0e, 0xde, 0xc4, 0xc3, 0x31, 0xb8, 0xd7, 0x21, 0xc5, 0xc9,
	0x32, 0xc4, 0xd4, 0x67, 0x78, 0xda, 0x46, 0x38, 0x86, 0x0f, 0x67, 0x40, 0x2e, 0x19, 0xc6, 0x37,
	0x43, 0xf0, 0xcc, 0xa8, 0x0a, 0xf6, 0x7b, 0x83, 0xde, 0xf0, 0xfa, 0x68, 0xef, 0xaf, 0x8c, 0x8e,
	0x6f, 0x84, 0xc9, 0x2b, 0x55, 0x41, 0xfa, 0x2b, 0x8a, 0xef, 0x6f, 0x66, 0x9e, 0xf8, 0x8e, 0x92,
	0x2f, 0x51, 0x7c, 0x6f, 0x6b, 0x80, 0xe4, 0x05, 0xdb, 0xd6, 0x35, 0xdb, 0xf5, 0x29, 0xfa, 0x72,
	0x2d, 0xab, 0xbb, 0x0a, 0xdb, 0x4c, 0x3a, 0x78, 0xf9, 0x47, 0xae, 0x36, 0xf1, 0xf5, 0xf7, 0xbf,
	0xef, 0xd1, 0xa3, 0xe4, 0x61, 0x73, 0xdb, 0x4f, 0x2b, 0x93, 0x27, 0xb8, 0x11, 0x45, 0xfc, 0xe8,
	0x73, 0xff, 0xf6, 0xb9, 0xdc, 0xbf, 0x0c, 0xd2, 0xaa, 0x85, 0x26, 0x96, 0x63, 0x35, 0xfa, 0x16,
	0xc5, 0x87, 0x39, 0x56, 0x5b, 0x0b, 0x18, 0x1d, 0xed, 0x54, 0xf8, 0xa4, 0xb9, 0xfa, 0xa4, 0xf7,
	0xf6, 0x59, 0xcb, 0x2b, 0x71, 0xae, 0x4c, 0xc9, 0xd0, 0x96, 0xbc, 0x04, 0xb3, 0x7c, 0x27, 0xf8,
	0x65, 0x82, 0xf5, 0x5f, 0xcb, 0xe3, 0x20, 0x7e, 0x44, 0x7b, 0x63, 0x29, 0x7f, 0x46, 0x83, 0xb1,
	0x07, 0xca, 0x82, 0x98, 0x97, 0x8d, 0x9a, 0xa6, 0xac, 0x5d, 0x4c, 0xe7, 0xc1, 0x32, 0x93, 0x05,
	0xcd, 0x3a, 0xcb, 0x6c, 0x9a, 0xce, 0x82, 0xe5, 0x7f, 0x74, 0xe8, 0x7f, 0x17, 0x42, 0x16, 0x24,
	0x44, 0x67, 0x12, 0x62, 0x9a, 0x0a, 0x11, 0x6c, 0xa7, 0xd7, 0x96, 0x39, 0x1f, 0x5c, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x6a, 0xf4, 0x10, 0x0a, 0xd4, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OperatingSystemVersionConstantServiceClient is the client API for OperatingSystemVersionConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperatingSystemVersionConstantServiceClient interface {
	// Returns the requested OS version constant in full detail.
	GetOperatingSystemVersionConstant(ctx context.Context, in *GetOperatingSystemVersionConstantRequest, opts ...grpc.CallOption) (*resources.OperatingSystemVersionConstant, error)
}

type operatingSystemVersionConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOperatingSystemVersionConstantServiceClient(cc grpc.ClientConnInterface) OperatingSystemVersionConstantServiceClient {
	return &operatingSystemVersionConstantServiceClient{cc}
}

func (c *operatingSystemVersionConstantServiceClient) GetOperatingSystemVersionConstant(ctx context.Context, in *GetOperatingSystemVersionConstantRequest, opts ...grpc.CallOption) (*resources.OperatingSystemVersionConstant, error) {
	out := new(resources.OperatingSystemVersionConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.OperatingSystemVersionConstantService/GetOperatingSystemVersionConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatingSystemVersionConstantServiceServer is the server API for OperatingSystemVersionConstantService service.
type OperatingSystemVersionConstantServiceServer interface {
	// Returns the requested OS version constant in full detail.
	GetOperatingSystemVersionConstant(context.Context, *GetOperatingSystemVersionConstantRequest) (*resources.OperatingSystemVersionConstant, error)
}

// UnimplementedOperatingSystemVersionConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOperatingSystemVersionConstantServiceServer struct {
}

func (*UnimplementedOperatingSystemVersionConstantServiceServer) GetOperatingSystemVersionConstant(ctx context.Context, req *GetOperatingSystemVersionConstantRequest) (*resources.OperatingSystemVersionConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperatingSystemVersionConstant not implemented")
}

func RegisterOperatingSystemVersionConstantServiceServer(s *grpc.Server, srv OperatingSystemVersionConstantServiceServer) {
	s.RegisterService(&_OperatingSystemVersionConstantService_serviceDesc, srv)
}

func _OperatingSystemVersionConstantService_GetOperatingSystemVersionConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperatingSystemVersionConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatingSystemVersionConstantServiceServer).GetOperatingSystemVersionConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.OperatingSystemVersionConstantService/GetOperatingSystemVersionConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatingSystemVersionConstantServiceServer).GetOperatingSystemVersionConstant(ctx, req.(*GetOperatingSystemVersionConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OperatingSystemVersionConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.OperatingSystemVersionConstantService",
	HandlerType: (*OperatingSystemVersionConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOperatingSystemVersionConstant",
			Handler:    _OperatingSystemVersionConstantService_GetOperatingSystemVersionConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/operating_system_version_constant_service.proto",
}
