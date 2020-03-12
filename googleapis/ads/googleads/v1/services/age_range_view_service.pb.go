// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/age_range_view_service.proto

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

// Request message for [AgeRangeViewService.GetAgeRangeView][google.ads.googleads.v1.services.AgeRangeViewService.GetAgeRangeView].
type GetAgeRangeViewRequest struct {
	// Required. The resource name of the age range view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAgeRangeViewRequest) Reset()         { *m = GetAgeRangeViewRequest{} }
func (m *GetAgeRangeViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetAgeRangeViewRequest) ProtoMessage()    {}
func (*GetAgeRangeViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec3750cf7941374c, []int{0}
}

func (m *GetAgeRangeViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAgeRangeViewRequest.Unmarshal(m, b)
}
func (m *GetAgeRangeViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAgeRangeViewRequest.Marshal(b, m, deterministic)
}
func (m *GetAgeRangeViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAgeRangeViewRequest.Merge(m, src)
}
func (m *GetAgeRangeViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetAgeRangeViewRequest.Size(m)
}
func (m *GetAgeRangeViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAgeRangeViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAgeRangeViewRequest proto.InternalMessageInfo

func (m *GetAgeRangeViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAgeRangeViewRequest)(nil), "google.ads.googleads.v1.services.GetAgeRangeViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/age_range_view_service.proto", fileDescriptor_ec3750cf7941374c)
}

var fileDescriptor_ec3750cf7941374c = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x8b, 0xd4, 0x40,
	0x18, 0x25, 0x39, 0x10, 0x0c, 0x8a, 0x10, 0x41, 0x97, 0x28, 0xb8, 0x1c, 0x57, 0x1c, 0x57, 0xcc,
	0x10, 0x05, 0x91, 0x91, 0x2b, 0x26, 0x4d, 0xac, 0xe4, 0x58, 0x21, 0x85, 0x04, 0xc2, 0x5c, 0xf2,
	0x39, 0x0e, 0x24, 0x33, 0xeb, 0x4c, 0x36, 0x57, 0x88, 0x8d, 0x7f, 0xc1, 0x7f, 0x60, 0xe9, 0xff,
	0xb0, 0xb9, 0xd6, 0xce, 0xca, 0xc2, 0xca, 0x9f, 0x60, 0xa3, 0x64, 0x27, 0x93, 0xcd, 0xea, 0x2e,
	0xdb, 0x3d, 0xf2, 0xde, 0xfb, 0xde, 0xf7, 0xbd, 0x49, 0x70, 0xce, 0x95, 0xe2, 0x35, 0x60, 0x56,
	0x19, 0x6c, 0x61, 0x8f, 0xba, 0x18, 0x1b, 0xd0, 0x9d, 0x28, 0xc1, 0x60, 0xc6, 0xa1, 0xd0, 0x4c,
	0x72, 0x28, 0x3a, 0x01, 0x57, 0xc5, 0xf0, 0x1d, 0x2d, 0xb5, 0x6a, 0x55, 0x38, 0xb7, 0x1e, 0xc4,
	0x2a, 0x83, 0x46, 0x3b, 0xea, 0x62, 0xe4, 0xec, 0xd1, 0xd3, 0x7d, 0x01, 0x1a, 0x8c, 0x5a, 0xe9,
	0xff, 0x13, 0xec, 0xe4, 0xe8, 0xa1, 0xf3, 0x2d, 0x05, 0x66, 0x52, 0xaa, 0x96, 0xb5, 0x42, 0x49,
	0x33, 0xb0, 0xf7, 0x27, 0x6c, 0x59, 0x0b, 0x90, 0xed, 0x40, 0x3c, 0x9a, 0x10, 0x6f, 0x04, 0xd4,
	0x55, 0x71, 0x09, 0x6f, 0x59, 0x27, 0x94, 0xb6, 0x82, 0xe3, 0x24, 0xb8, 0x97, 0x42, 0x4b, 0x39,
	0x2c, 0xfa, 0xc4, 0x4c, 0xc0, 0xd5, 0x02, 0xde, 0xad, 0xc0, 0xb4, 0xe1, 0x69, 0x70, 0xdb, 0xed,
	0x54, 0x48, 0xd6, 0xc0, 0xcc, 0x9b, 0x7b, 0xa7, 0x37, 0x93, 0xa3, 0x1f, 0xd4, 0x5f, 0xdc, 0x72,
	0xcc, 0x4b, 0xd6, 0xc0, 0xe3, 0xdf, 0x5e, 0x70, 0x77, 0x3a, 0xe1, 0x95, 0x3d, 0x36, 0xfc, 0xea,
	0x05, 0x77, 0xfe, 0x19, 0x1e, 0x3e, 0x43, 0x87, 0x2a, 0x42, 0xbb, 0xf7, 0x89, 0xf0, 0x5e, 0xe7,
	0x58, 0x1d, 0x9a, 0xfa, 0x8e, 0xd3, 0xef, 0x74, 0xfb, 0x82, 0x8f, 0xdf, 0x7e, 0x7e, 0xf2, 0xe3,
	0x10, 0xf7, 0x75, 0xbf, 0xdf, 0x62, 0xce, 0xcb, 0x95, 0x69, 0x55, 0x03, 0xda, 0xe0, 0xb3, 0xbe,
	0xff, 0x71, 0x88, 0xc1, 0x67, 0x1f, 0xa2, 0x07, 0xd7, 0x74, 0xb6, 0x09, 0x1c, 0xd0, 0x52, 0x18,
	0x54, 0xaa, 0x26, 0xf9, 0xe3, 0x05, 0x27, 0xa5, 0x6a, 0x0e, 0x9e, 0x95, 0xcc, 0x76, 0x54, 0x74,
	0xd1, 0xbf, 0xc1, 0x85, 0xf7, 0xfa, 0xc5, 0xe0, 0xe6, 0xaa, 0x66, 0x92, 0x23, 0xa5, 0x39, 0xe6,
	0x20, 0xd7, 0x2f, 0x84, 0x37, 0x79, 0xfb, 0xff, 0xca, 0xe7, 0x0e, 0x7c, 0xf6, 0x8f, 0x52, 0x4a,
	0xbf, 0xf8, 0xf3, 0xd4, 0x0e, 0xa4, 0x95, 0x41, 0x16, 0xf6, 0x28, 0x8b, 0xd1, 0x10, 0x6c, 0xae,
	0x9d, 0x24, 0xa7, 0x95, 0xc9, 0x47, 0x49, 0x9e, 0xc5, 0xb9, 0x93, 0xfc, 0xf2, 0x4f, 0xec, 0x77,
	0x42, 0x68, 0x65, 0x08, 0x19, 0x45, 0x84, 0x64, 0x31, 0x21, 0x4e, 0x76, 0x79, 0x63, 0xbd, 0xe7,
	0x93, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xde, 0x48, 0xdd, 0x92, 0x3c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AgeRangeViewServiceClient is the client API for AgeRangeViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgeRangeViewServiceClient interface {
	// Returns the requested age range view in full detail.
	GetAgeRangeView(ctx context.Context, in *GetAgeRangeViewRequest, opts ...grpc.CallOption) (*resources.AgeRangeView, error)
}

type ageRangeViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgeRangeViewServiceClient(cc grpc.ClientConnInterface) AgeRangeViewServiceClient {
	return &ageRangeViewServiceClient{cc}
}

func (c *ageRangeViewServiceClient) GetAgeRangeView(ctx context.Context, in *GetAgeRangeViewRequest, opts ...grpc.CallOption) (*resources.AgeRangeView, error) {
	out := new(resources.AgeRangeView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AgeRangeViewService/GetAgeRangeView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgeRangeViewServiceServer is the server API for AgeRangeViewService service.
type AgeRangeViewServiceServer interface {
	// Returns the requested age range view in full detail.
	GetAgeRangeView(context.Context, *GetAgeRangeViewRequest) (*resources.AgeRangeView, error)
}

// UnimplementedAgeRangeViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAgeRangeViewServiceServer struct {
}

func (*UnimplementedAgeRangeViewServiceServer) GetAgeRangeView(ctx context.Context, req *GetAgeRangeViewRequest) (*resources.AgeRangeView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgeRangeView not implemented")
}

func RegisterAgeRangeViewServiceServer(s *grpc.Server, srv AgeRangeViewServiceServer) {
	s.RegisterService(&_AgeRangeViewService_serviceDesc, srv)
}

func _AgeRangeViewService_GetAgeRangeView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgeRangeViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgeRangeViewServiceServer).GetAgeRangeView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AgeRangeViewService/GetAgeRangeView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgeRangeViewServiceServer).GetAgeRangeView(ctx, req.(*GetAgeRangeViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgeRangeViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.AgeRangeViewService",
	HandlerType: (*AgeRangeViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAgeRangeView",
			Handler:    _AgeRangeViewService_GetAgeRangeView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/age_range_view_service.proto",
}
