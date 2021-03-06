// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/ad_schedule_view_service.proto

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

// Request message for [AdScheduleViewService.GetAdScheduleView][google.ads.googleads.v1.services.AdScheduleViewService.GetAdScheduleView].
type GetAdScheduleViewRequest struct {
	// Required. The resource name of the ad schedule view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdScheduleViewRequest) Reset()         { *m = GetAdScheduleViewRequest{} }
func (m *GetAdScheduleViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdScheduleViewRequest) ProtoMessage()    {}
func (*GetAdScheduleViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6e0ec29d39ff4cc, []int{0}
}

func (m *GetAdScheduleViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdScheduleViewRequest.Unmarshal(m, b)
}
func (m *GetAdScheduleViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdScheduleViewRequest.Marshal(b, m, deterministic)
}
func (m *GetAdScheduleViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdScheduleViewRequest.Merge(m, src)
}
func (m *GetAdScheduleViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdScheduleViewRequest.Size(m)
}
func (m *GetAdScheduleViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdScheduleViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdScheduleViewRequest proto.InternalMessageInfo

func (m *GetAdScheduleViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdScheduleViewRequest)(nil), "google.ads.googleads.v1.services.GetAdScheduleViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/ad_schedule_view_service.proto", fileDescriptor_b6e0ec29d39ff4cc)
}

var fileDescriptor_b6e0ec29d39ff4cc = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3f, 0x8b, 0xd4, 0x40,
	0x14, 0x27, 0x39, 0x10, 0x0c, 0x5a, 0x18, 0x10, 0x97, 0x28, 0xb8, 0x1c, 0x57, 0x1c, 0x57, 0xcc,
	0x10, 0xaf, 0x91, 0x11, 0x91, 0x59, 0x84, 0xd5, 0x46, 0x8e, 0x3b, 0x48, 0x21, 0x81, 0x30, 0x97,
	0x79, 0xe6, 0x06, 0x92, 0xcc, 0x9a, 0x97, 0xe4, 0x0a, 0xb1, 0x50, 0xfc, 0x06, 0x7e, 0x03, 0x4b,
	0x3f, 0xca, 0x81, 0x95, 0x9d, 0x95, 0x85, 0x95, 0x9f, 0x42, 0xb2, 0x93, 0xc9, 0x6e, 0x70, 0xc3,
	0x76, 0x3f, 0xf2, 0xfb, 0xf7, 0xde, 0x9b, 0x78, 0x2f, 0x32, 0xad, 0xb3, 0x1c, 0xa8, 0x90, 0x48,
	0x0d, 0xec, 0x50, 0x1b, 0x52, 0x84, 0xaa, 0x55, 0x29, 0x20, 0x15, 0x32, 0xc1, 0xf4, 0x0a, 0x64,
	0x93, 0x43, 0xd2, 0x2a, 0xb8, 0x4e, 0x7a, 0x86, 0xac, 0x2a, 0x5d, 0x6b, 0x7f, 0x6e, 0x5c, 0x44,
	0x48, 0x24, 0x43, 0x00, 0x69, 0x43, 0x62, 0x03, 0x82, 0xa7, 0x53, 0x15, 0x15, 0xa0, 0x6e, 0xaa,
	0x5d, 0x1d, 0x26, 0x3b, 0x78, 0x64, 0x9d, 0x2b, 0x45, 0x45, 0x59, 0xea, 0x5a, 0xd4, 0x4a, 0x97,
	0xd8, 0xb3, 0x0f, 0xb6, 0xd8, 0x34, 0x57, 0x50, 0xd6, 0x3d, 0xf1, 0x78, 0x8b, 0x78, 0xa7, 0x20,
	0x97, 0xc9, 0x25, 0x5c, 0x89, 0x56, 0xe9, 0xca, 0x08, 0x0e, 0x5f, 0x7a, 0xb3, 0x25, 0xd4, 0x5c,
	0x5e, 0xf4, 0x9d, 0x91, 0x82, 0xeb, 0x73, 0x78, 0xdf, 0x00, 0xd6, 0xfe, 0xb1, 0x77, 0xd7, 0xce,
	0x95, 0x94, 0xa2, 0x80, 0x99, 0x33, 0x77, 0x8e, 0x6f, 0x2f, 0x0e, 0x7e, 0x73, 0xf7, 0xfc, 0x8e,
	0x65, 0xde, 0x88, 0x02, 0x9e, 0x7c, 0x71, 0xbd, 0xfb, 0xe3, 0x8c, 0x0b, 0xb3, 0xb2, 0xff, 0xc3,
	0xf1, 0xee, 0xfd, 0x57, 0xe0, 0x33, 0xb2, 0xef, 0x54, 0x64, 0x6a, 0xaa, 0x20, 0x9c, 0xf4, 0x0e,
	0x47, 0x24, 0x63, 0xe7, 0xe1, 0xeb, 0x5f, 0x7c, 0xbc, 0xc9, 0xe7, 0x9f, 0x7f, 0xbe, 0xba, 0xa7,
	0x7e, 0xd8, 0x9d, 0xfe, 0xc3, 0x88, 0x79, 0x9e, 0x36, 0x58, 0xeb, 0x02, 0x2a, 0xa4, 0x27, 0x54,
	0x8c, 0x62, 0x90, 0x9e, 0x7c, 0x0c, 0x1e, 0xde, 0xf0, 0xd9, 0xa6, 0xb4, 0x47, 0x2b, 0x85, 0x24,
	0xd5, 0xc5, 0xe2, 0x93, 0xeb, 0x1d, 0xa5, 0xba, 0xd8, 0xbb, 0xdc, 0x22, 0xd8, 0x79, 0xac, 0xb3,
	0xee, 0x45, 0xce, 0x9c, 0xb7, 0xaf, 0x7a, 0x7f, 0xa6, 0x73, 0x51, 0x66, 0x44, 0x57, 0x19, 0xcd,
	0xa0, 0x5c, 0xbf, 0x17, 0xdd, 0x34, 0x4e, 0xff, 0xa7, 0xcf, 0x2c, 0xf8, 0xe6, 0x1e, 0x2c, 0x39,
	0xff, 0xee, 0xce, 0x97, 0x26, 0x90, 0x4b, 0x24, 0x06, 0x76, 0x28, 0x0a, 0x49, 0x5f, 0x8c, 0x37,
	0x56, 0x12, 0x73, 0x89, 0xf1, 0x20, 0x89, 0xa3, 0x30, 0xb6, 0x92, 0xbf, 0xee, 0x91, 0xf9, 0xce,
	0x18, 0x97, 0xc8, 0xd8, 0x20, 0x62, 0x2c, 0x0a, 0x19, 0xb3, 0xb2, 0xcb, 0x5b, 0xeb, 0x39, 0x4f,
	0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xec, 0xd4, 0x24, 0x4e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdScheduleViewServiceClient is the client API for AdScheduleViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdScheduleViewServiceClient interface {
	// Returns the requested ad schedule view in full detail.
	GetAdScheduleView(ctx context.Context, in *GetAdScheduleViewRequest, opts ...grpc.CallOption) (*resources.AdScheduleView, error)
}

type adScheduleViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdScheduleViewServiceClient(cc grpc.ClientConnInterface) AdScheduleViewServiceClient {
	return &adScheduleViewServiceClient{cc}
}

func (c *adScheduleViewServiceClient) GetAdScheduleView(ctx context.Context, in *GetAdScheduleViewRequest, opts ...grpc.CallOption) (*resources.AdScheduleView, error) {
	out := new(resources.AdScheduleView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AdScheduleViewService/GetAdScheduleView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdScheduleViewServiceServer is the server API for AdScheduleViewService service.
type AdScheduleViewServiceServer interface {
	// Returns the requested ad schedule view in full detail.
	GetAdScheduleView(context.Context, *GetAdScheduleViewRequest) (*resources.AdScheduleView, error)
}

// UnimplementedAdScheduleViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdScheduleViewServiceServer struct {
}

func (*UnimplementedAdScheduleViewServiceServer) GetAdScheduleView(ctx context.Context, req *GetAdScheduleViewRequest) (*resources.AdScheduleView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdScheduleView not implemented")
}

func RegisterAdScheduleViewServiceServer(s *grpc.Server, srv AdScheduleViewServiceServer) {
	s.RegisterService(&_AdScheduleViewService_serviceDesc, srv)
}

func _AdScheduleViewService_GetAdScheduleView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdScheduleViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdScheduleViewServiceServer).GetAdScheduleView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AdScheduleViewService/GetAdScheduleView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdScheduleViewServiceServer).GetAdScheduleView(ctx, req.(*GetAdScheduleViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdScheduleViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.AdScheduleViewService",
	HandlerType: (*AdScheduleViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdScheduleView",
			Handler:    _AdScheduleViewService_GetAdScheduleView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/ad_schedule_view_service.proto",
}
