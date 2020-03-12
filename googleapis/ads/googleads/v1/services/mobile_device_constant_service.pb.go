// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/mobile_device_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
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

// Request message for [MobileDeviceConstantService.GetMobileDeviceConstant][google.ads.googleads.v1.services.MobileDeviceConstantService.GetMobileDeviceConstant].
type GetMobileDeviceConstantRequest struct {
	// Required. Resource name of the mobile device to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMobileDeviceConstantRequest) Reset()         { *m = GetMobileDeviceConstantRequest{} }
func (m *GetMobileDeviceConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetMobileDeviceConstantRequest) ProtoMessage()    {}
func (*GetMobileDeviceConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dacc361447bd918, []int{0}
}

func (m *GetMobileDeviceConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMobileDeviceConstantRequest.Unmarshal(m, b)
}
func (m *GetMobileDeviceConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMobileDeviceConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetMobileDeviceConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMobileDeviceConstantRequest.Merge(m, src)
}
func (m *GetMobileDeviceConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetMobileDeviceConstantRequest.Size(m)
}
func (m *GetMobileDeviceConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMobileDeviceConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMobileDeviceConstantRequest proto.InternalMessageInfo

func (m *GetMobileDeviceConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMobileDeviceConstantRequest)(nil), "google.ads.googleads.v1.services.GetMobileDeviceConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/mobile_device_constant_service.proto", fileDescriptor_6dacc361447bd918)
}

var fileDescriptor_6dacc361447bd918 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4f, 0x8b, 0xd3, 0x40,
	0x1c, 0x25, 0x29, 0x08, 0x06, 0xbd, 0xe4, 0xd2, 0x92, 0x8a, 0x86, 0xd2, 0x43, 0x51, 0x9c, 0x21,
	0x7a, 0x10, 0x46, 0x14, 0xa7, 0x55, 0x2a, 0x82, 0x52, 0x2a, 0xf4, 0x20, 0x81, 0x30, 0x4d, 0xc6,
	0x38, 0x90, 0xcc, 0xd4, 0x4c, 0x9a, 0x8b, 0x78, 0xf1, 0xe2, 0x07, 0x10, 0xfc, 0x00, 0x7b, 0xdc,
	0x8f, 0xd2, 0xeb, 0x9e, 0x76, 0x4f, 0x7b, 0xd8, 0xd3, 0x7e, 0x8a, 0x25, 0x99, 0x4c, 0xda, 0x42,
	0xd3, 0xde, 0x5e, 0xf2, 0xde, 0xef, 0xbd, 0xdf, 0x9f, 0xb1, 0x3e, 0xc4, 0x42, 0xc4, 0x09, 0x85,
	0x24, 0x92, 0x50, 0xc1, 0x12, 0x15, 0x1e, 0x94, 0x34, 0x2b, 0x58, 0x48, 0x25, 0x4c, 0xc5, 0x92,
	0x25, 0x34, 0x88, 0x68, 0xf9, 0x19, 0x84, 0x82, 0xcb, 0x9c, 0xf0, 0x3c, 0xa8, 0x79, 0xb0, 0xca,
	0x44, 0x2e, 0x6c, 0x57, 0xd5, 0x02, 0x12, 0x49, 0xd0, 0xd8, 0x80, 0xc2, 0x03, 0xda, 0xc6, 0x79,
	0xdb, 0x16, 0x94, 0x51, 0x29, 0xd6, 0x59, 0x7b, 0x92, 0x4a, 0x70, 0x1e, 0xe9, 0xfa, 0x15, 0x83,
	0x84, 0x73, 0x91, 0x93, 0x9c, 0x09, 0x2e, 0x6b, 0xb6, 0xbb, 0xc3, 0x86, 0x09, 0xa3, 0x4d, 0xd9,
	0x93, 0x1d, 0xe2, 0x3b, 0xa3, 0x49, 0x14, 0x2c, 0xe9, 0x0f, 0x52, 0x30, 0x91, 0x29, 0xc1, 0xe0,
	0x93, 0xf5, 0x78, 0x4a, 0xf3, 0xcf, 0x55, 0xf4, 0xfb, 0x2a, 0x79, 0x52, 0x07, 0xcf, 0xe9, 0xcf,
	0x35, 0x95, 0xb9, 0x3d, 0xb2, 0x1e, 0xea, 0x1e, 0x03, 0x4e, 0x52, 0xda, 0x33, 0x5c, 0x63, 0x74,
	0x7f, 0xdc, 0xb9, 0xc6, 0xe6, 0xfc, 0x81, 0x66, 0xbe, 0x90, 0x94, 0xbe, 0xf8, 0x6f, 0x5a, 0xfd,
	0x43, 0x4e, 0x5f, 0xd5, 0x12, 0xec, 0x4b, 0xc3, 0xea, 0xb6, 0x84, 0xd9, 0xef, 0xc0, 0xa9, 0x15,
	0x82, 0xe3, 0x7d, 0x3a, 0xaf, 0x5a, 0x1d, 0x9a, 0x15, 0x83, 0x43, 0xf5, 0x83, 0xc9, 0x15, 0xde,
	0x9f, 0xf0, 0xcf, 0xc5, 0xcd, 0x3f, 0xf3, 0xb9, 0xfd, 0xac, 0x3c, 0xcf, 0xaf, 0x3d, 0xe6, 0x4d,
	0x7a, 0xc0, 0x40, 0xc2, 0xa7, 0xbf, 0x9d, 0xfe, 0x06, 0xf7, 0xb6, 0xa1, 0x35, 0x5a, 0x31, 0x09,
	0x42, 0x91, 0x8e, 0xff, 0x9a, 0xd6, 0x30, 0x14, 0xe9, 0xc9, 0x11, 0xc7, 0xee, 0x91, 0xf5, 0xcd,
	0xca, 0x7b, 0xcd, 0x8c, 0x6f, 0x1f, 0x6b, 0x97, 0x58, 0x24, 0x84, 0xc7, 0x40, 0x64, 0x31, 0x8c,
	0x29, 0xaf, 0xae, 0x09, 0xb7, 0xb9, 0xed, 0x2f, 0xfa, 0xb5, 0x06, 0x67, 0x66, 0x67, 0x8a, 0xf1,
	0xb9, 0xe9, 0x4e, 0x95, 0x21, 0x8e, 0x24, 0x50, 0xb0, 0x44, 0x0b, 0x0f, 0xd4, 0xc1, 0x72, 0xa3,
	0x25, 0x3e, 0x8e, 0xa4, 0xdf, 0x48, 0xfc, 0x85, 0xe7, 0x6b, 0xc9, 0xad, 0x39, 0x54, 0xff, 0x11,
	0xc2, 0x91, 0x44, 0xa8, 0x11, 0x21, 0xb4, 0xf0, 0x10, 0xd2, 0xb2, 0xe5, 0xbd, 0xaa, 0xcf, 0x97,
	0x77, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x10, 0x89, 0x4f, 0x78, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MobileDeviceConstantServiceClient is the client API for MobileDeviceConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MobileDeviceConstantServiceClient interface {
	// Returns the requested mobile device constant in full detail.
	GetMobileDeviceConstant(ctx context.Context, in *GetMobileDeviceConstantRequest, opts ...grpc.CallOption) (*resources.MobileDeviceConstant, error)
}

type mobileDeviceConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMobileDeviceConstantServiceClient(cc grpc.ClientConnInterface) MobileDeviceConstantServiceClient {
	return &mobileDeviceConstantServiceClient{cc}
}

func (c *mobileDeviceConstantServiceClient) GetMobileDeviceConstant(ctx context.Context, in *GetMobileDeviceConstantRequest, opts ...grpc.CallOption) (*resources.MobileDeviceConstant, error) {
	out := new(resources.MobileDeviceConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.MobileDeviceConstantService/GetMobileDeviceConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MobileDeviceConstantServiceServer is the server API for MobileDeviceConstantService service.
type MobileDeviceConstantServiceServer interface {
	// Returns the requested mobile device constant in full detail.
	GetMobileDeviceConstant(context.Context, *GetMobileDeviceConstantRequest) (*resources.MobileDeviceConstant, error)
}

// UnimplementedMobileDeviceConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMobileDeviceConstantServiceServer struct {
}

func (*UnimplementedMobileDeviceConstantServiceServer) GetMobileDeviceConstant(ctx context.Context, req *GetMobileDeviceConstantRequest) (*resources.MobileDeviceConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMobileDeviceConstant not implemented")
}

func RegisterMobileDeviceConstantServiceServer(s *grpc.Server, srv MobileDeviceConstantServiceServer) {
	s.RegisterService(&_MobileDeviceConstantService_serviceDesc, srv)
}

func _MobileDeviceConstantService_GetMobileDeviceConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMobileDeviceConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileDeviceConstantServiceServer).GetMobileDeviceConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.MobileDeviceConstantService/GetMobileDeviceConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileDeviceConstantServiceServer).GetMobileDeviceConstant(ctx, req.(*GetMobileDeviceConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MobileDeviceConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.MobileDeviceConstantService",
	HandlerType: (*MobileDeviceConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMobileDeviceConstant",
			Handler:    _MobileDeviceConstantService_GetMobileDeviceConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/mobile_device_constant_service.proto",
}
