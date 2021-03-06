// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/campaign_audience_view_service.proto

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

// Request message for [CampaignAudienceViewService.GetCampaignAudienceView][google.ads.googleads.v2.services.CampaignAudienceViewService.GetCampaignAudienceView].
type GetCampaignAudienceViewRequest struct {
	// Required. The resource name of the campaign audience view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignAudienceViewRequest) Reset()         { *m = GetCampaignAudienceViewRequest{} }
func (m *GetCampaignAudienceViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignAudienceViewRequest) ProtoMessage()    {}
func (*GetCampaignAudienceViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8ae39646e976956, []int{0}
}

func (m *GetCampaignAudienceViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignAudienceViewRequest.Unmarshal(m, b)
}
func (m *GetCampaignAudienceViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignAudienceViewRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignAudienceViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignAudienceViewRequest.Merge(m, src)
}
func (m *GetCampaignAudienceViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignAudienceViewRequest.Size(m)
}
func (m *GetCampaignAudienceViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignAudienceViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignAudienceViewRequest proto.InternalMessageInfo

func (m *GetCampaignAudienceViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignAudienceViewRequest)(nil), "google.ads.googleads.v2.services.GetCampaignAudienceViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/campaign_audience_view_service.proto", fileDescriptor_b8ae39646e976956)
}

var fileDescriptor_b8ae39646e976956 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x8b, 0xd4, 0x40,
	0x18, 0x25, 0x39, 0x10, 0x0c, 0xda, 0xa4, 0xb9, 0x23, 0x27, 0x1a, 0x8e, 0x2b, 0x8e, 0x2b, 0x66,
	0x20, 0x16, 0x87, 0x23, 0x8a, 0xb3, 0x22, 0x2b, 0x16, 0x7a, 0x9c, 0x90, 0x42, 0x02, 0x61, 0x2e,
	0xf9, 0x8c, 0x03, 0xc9, 0x4c, 0xcc, 0x24, 0xb9, 0x42, 0x6c, 0x6c, 0xfc, 0x01, 0xfe, 0x83, 0x2b,
	0xfd, 0x29, 0xd7, 0xda, 0x59, 0x59, 0x68, 0xe3, 0xaf, 0x38, 0xb2, 0x33, 0x93, 0xdd, 0x85, 0xcd,
	0x6e, 0xf7, 0xc8, 0x7b, 0xdf, 0x7b, 0xdf, 0x7c, 0x2f, 0xde, 0xab, 0x42, 0xca, 0xa2, 0x04, 0xcc,
	0x72, 0x85, 0x35, 0x1c, 0x50, 0x1f, 0x61, 0x05, 0x4d, 0xcf, 0x33, 0x50, 0x38, 0x63, 0x55, 0xcd,
	0x78, 0x21, 0x52, 0xd6, 0xe5, 0x1c, 0x44, 0x06, 0x69, 0xcf, 0xe1, 0x2a, 0x35, 0x3c, 0xaa, 0x1b,
	0xd9, 0x4a, 0x3f, 0xd4, 0xb3, 0x88, 0xe5, 0x0a, 0x8d, 0x36, 0xa8, 0x8f, 0x90, 0xb5, 0x09, 0x9e,
	0x4f, 0x05, 0x35, 0xa0, 0x64, 0xd7, 0x4c, 0x27, 0xe9, 0x84, 0xe0, 0x81, 0x9d, 0xaf, 0x39, 0x66,
	0x42, 0xc8, 0x96, 0xb5, 0x5c, 0x0a, 0x65, 0xd8, 0xfd, 0x15, 0x36, 0x2b, 0x39, 0x88, 0xd6, 0x10,
	0x8f, 0x56, 0x88, 0x8f, 0x1c, 0xca, 0x3c, 0xbd, 0x84, 0x4f, 0xac, 0xe7, 0xb2, 0xd1, 0x82, 0xa3,
	0x37, 0xde, 0xc3, 0x39, 0xb4, 0x2f, 0x4d, 0x34, 0x35, 0xc9, 0x31, 0x87, 0xab, 0x0b, 0xf8, 0xdc,
	0x81, 0x6a, 0xfd, 0x13, 0xef, 0xbe, 0xdd, 0x31, 0x15, 0xac, 0x82, 0x03, 0x27, 0x74, 0x4e, 0xee,
	0xce, 0xf6, 0xfe, 0x50, 0xf7, 0xe2, 0x9e, 0x65, 0xde, 0xb2, 0x0a, 0xa2, 0x6b, 0xd7, 0x3b, 0xdc,
	0xe4, 0xf4, 0x5e, 0x1f, 0xc1, 0xff, 0xe7, 0x78, 0xfb, 0x13, 0x61, 0xfe, 0x0b, 0xb4, 0xeb, 0x84,
	0x68, 0xfb, 0x9e, 0xc1, 0xd9, 0xa4, 0xc3, 0x78, 0x62, 0xb4, 0x69, 0xfe, 0xe8, 0xdd, 0x6f, 0xba,
	0xfe, 0xc2, 0x6f, 0xbf, 0xfe, 0xfe, 0x70, 0x9f, 0xf8, 0x67, 0x43, 0x3d, 0x5f, 0xd6, 0x98, 0x67,
	0x59, 0xa7, 0x5a, 0x59, 0x41, 0xa3, 0xf0, 0xe9, 0xd8, 0xd7, 0xaa, 0x99, 0xc2, 0xa7, 0x5f, 0x83,
	0xc3, 0x1b, 0x7a, 0xb0, 0x5c, 0xc0, 0xa0, 0x9a, 0x2b, 0x94, 0xc9, 0x6a, 0xf6, 0xdd, 0xf5, 0x8e,
	0x33, 0x59, 0xed, 0x7c, 0xee, 0x2c, 0xdc, 0x72, 0xca, 0xf3, 0xa1, 0xbb, 0x73, 0xe7, 0xc3, 0x6b,
	0xe3, 0x52, 0xc8, 0x92, 0x89, 0x02, 0xc9, 0xa6, 0xc0, 0x05, 0x88, 0x45, 0xb3, 0x78, 0x99, 0x3b,
	0xfd, 0x77, 0x3f, 0xb5, 0xe0, 0xda, 0xdd, 0x9b, 0x53, 0xfa, 0xd3, 0x0d, 0xe7, 0xda, 0x90, 0xe6,
	0x0a, 0x69, 0x38, 0xa0, 0x38, 0x42, 0x26, 0x58, 0xdd, 0x58, 0x49, 0x42, 0x73, 0x95, 0x8c, 0x92,
	0x24, 0x8e, 0x12, 0x2b, 0xf9, 0xef, 0x1e, 0xeb, 0xef, 0x84, 0xd0, 0x5c, 0x11, 0x32, 0x8a, 0x08,
	0x89, 0x23, 0x42, 0xac, 0xec, 0xf2, 0xce, 0x62, 0xcf, 0xc7, 0xb7, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xc1, 0x10, 0x0c, 0xf0, 0x84, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CampaignAudienceViewServiceClient is the client API for CampaignAudienceViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignAudienceViewServiceClient interface {
	// Returns the requested campaign audience view in full detail.
	GetCampaignAudienceView(ctx context.Context, in *GetCampaignAudienceViewRequest, opts ...grpc.CallOption) (*resources.CampaignAudienceView, error)
}

type campaignAudienceViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignAudienceViewServiceClient(cc grpc.ClientConnInterface) CampaignAudienceViewServiceClient {
	return &campaignAudienceViewServiceClient{cc}
}

func (c *campaignAudienceViewServiceClient) GetCampaignAudienceView(ctx context.Context, in *GetCampaignAudienceViewRequest, opts ...grpc.CallOption) (*resources.CampaignAudienceView, error) {
	out := new(resources.CampaignAudienceView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CampaignAudienceViewService/GetCampaignAudienceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignAudienceViewServiceServer is the server API for CampaignAudienceViewService service.
type CampaignAudienceViewServiceServer interface {
	// Returns the requested campaign audience view in full detail.
	GetCampaignAudienceView(context.Context, *GetCampaignAudienceViewRequest) (*resources.CampaignAudienceView, error)
}

// UnimplementedCampaignAudienceViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCampaignAudienceViewServiceServer struct {
}

func (*UnimplementedCampaignAudienceViewServiceServer) GetCampaignAudienceView(ctx context.Context, req *GetCampaignAudienceViewRequest) (*resources.CampaignAudienceView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCampaignAudienceView not implemented")
}

func RegisterCampaignAudienceViewServiceServer(s *grpc.Server, srv CampaignAudienceViewServiceServer) {
	s.RegisterService(&_CampaignAudienceViewService_serviceDesc, srv)
}

func _CampaignAudienceViewService_GetCampaignAudienceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignAudienceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignAudienceViewServiceServer).GetCampaignAudienceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CampaignAudienceViewService/GetCampaignAudienceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignAudienceViewServiceServer).GetCampaignAudienceView(ctx, req.(*GetCampaignAudienceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignAudienceViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.CampaignAudienceViewService",
	HandlerType: (*CampaignAudienceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignAudienceView",
			Handler:    _CampaignAudienceViewService_GetCampaignAudienceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/campaign_audience_view_service.proto",
}
