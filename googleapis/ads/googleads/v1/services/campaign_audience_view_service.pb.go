// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/campaign_audience_view_service.proto

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

// Request message for [CampaignAudienceViewService.GetCampaignAudienceView][google.ads.googleads.v1.services.CampaignAudienceViewService.GetCampaignAudienceView].
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
	return fileDescriptor_ae1606ecaa0347f1, []int{0}
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
	proto.RegisterType((*GetCampaignAudienceViewRequest)(nil), "google.ads.googleads.v1.services.GetCampaignAudienceViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/campaign_audience_view_service.proto", fileDescriptor_ae1606ecaa0347f1)
}

var fileDescriptor_ae1606ecaa0347f1 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x8b, 0xd4, 0x40,
	0x18, 0x25, 0x39, 0x10, 0x0c, 0xda, 0xa4, 0xb9, 0x23, 0x27, 0x1a, 0x8e, 0x2b, 0x8e, 0x2b, 0x66,
	0x88, 0x16, 0x87, 0x23, 0x8a, 0xb3, 0x22, 0x2b, 0x16, 0x7a, 0x9c, 0x90, 0x42, 0x02, 0x61, 0x2e,
	0xf9, 0x8c, 0x03, 0xc9, 0x4c, 0xcc, 0x24, 0xb9, 0x42, 0x6c, 0x6c, 0xfc, 0x01, 0xfe, 0x83, 0x2b,
	0xfd, 0x29, 0xd7, 0xda, 0x59, 0x59, 0x68, 0xe3, 0xaf, 0x90, 0xec, 0xcc, 0x64, 0x77, 0x61, 0xb3,
	0xdb, 0x3d, 0xf2, 0xde, 0xf7, 0xde, 0x37, 0xdf, 0x8b, 0xf7, 0xb2, 0x90, 0xb2, 0x28, 0x01, 0xb3,
	0x5c, 0x61, 0x0d, 0x07, 0xd4, 0x47, 0x58, 0x41, 0xd3, 0xf3, 0x0c, 0x14, 0xce, 0x58, 0x55, 0x33,
	0x5e, 0x88, 0x94, 0x75, 0x39, 0x07, 0x91, 0x41, 0xda, 0x73, 0xb8, 0x4a, 0x0d, 0x8f, 0xea, 0x46,
	0xb6, 0xd2, 0x0f, 0xf5, 0x2c, 0x62, 0xb9, 0x42, 0xa3, 0x0d, 0xea, 0x23, 0x64, 0x6d, 0x82, 0x67,
	0x53, 0x41, 0x0d, 0x28, 0xd9, 0x35, 0xd3, 0x49, 0x3a, 0x21, 0xb8, 0x67, 0xe7, 0x6b, 0x8e, 0x99,
	0x10, 0xb2, 0x65, 0x2d, 0x97, 0x42, 0x19, 0x76, 0x7f, 0x85, 0xcd, 0x4a, 0x0e, 0xa2, 0x35, 0xc4,
	0x83, 0x15, 0xe2, 0x03, 0x87, 0x32, 0x4f, 0x2f, 0xe1, 0x23, 0xeb, 0xb9, 0x6c, 0xb4, 0xe0, 0xe8,
	0xb5, 0x77, 0x7f, 0x0e, 0xed, 0x0b, 0x13, 0x4d, 0x4d, 0x72, 0xcc, 0xe1, 0xea, 0x02, 0x3e, 0x75,
	0xa0, 0x5a, 0xff, 0xc4, 0xbb, 0x6b, 0x77, 0x4c, 0x05, 0xab, 0xe0, 0xc0, 0x09, 0x9d, 0x93, 0xdb,
	0xb3, 0xbd, 0xdf, 0xd4, 0xbd, 0xb8, 0x63, 0x99, 0x37, 0xac, 0x82, 0x87, 0xd7, 0xae, 0x77, 0xb8,
	0xc9, 0xe9, 0x9d, 0x3e, 0x82, 0xff, 0xd7, 0xf1, 0xf6, 0x27, 0xc2, 0xfc, 0xe7, 0x68, 0xd7, 0x09,
	0xd1, 0xf6, 0x3d, 0x83, 0xb3, 0x49, 0x87, 0xf1, 0xc4, 0x68, 0xd3, 0xfc, 0xd1, 0xdb, 0x5f, 0x74,
	0xfd, 0x85, 0x5f, 0x7f, 0xfe, 0xf9, 0xee, 0x3e, 0xf6, 0xcf, 0x86, 0x7a, 0x3e, 0xaf, 0x31, 0x4f,
	0xb3, 0x4e, 0xb5, 0xb2, 0x82, 0x46, 0xe1, 0xd3, 0xb1, 0xaf, 0x55, 0x33, 0x85, 0x4f, 0xbf, 0x04,
	0x87, 0x37, 0xf4, 0x60, 0xb9, 0x80, 0x41, 0x35, 0x57, 0x28, 0x93, 0xd5, 0xec, 0x9b, 0xeb, 0x1d,
	0x67, 0xb2, 0xda, 0xf9, 0xdc, 0x59, 0xb8, 0xe5, 0x94, 0xe7, 0x43, 0x77, 0xe7, 0xce, 0xfb, 0x57,
	0xc6, 0xa5, 0x90, 0x25, 0x13, 0x05, 0x92, 0x4d, 0x81, 0x0b, 0x10, 0x8b, 0x66, 0xf1, 0x32, 0x77,
	0xfa, 0xef, 0x7e, 0x62, 0xc1, 0xb5, 0xbb, 0x37, 0xa7, 0xf4, 0x87, 0x1b, 0xce, 0xb5, 0x21, 0xcd,
	0x15, 0xd2, 0x70, 0x40, 0x71, 0x84, 0x4c, 0xb0, 0xba, 0xb1, 0x92, 0x84, 0xe6, 0x2a, 0x19, 0x25,
	0x49, 0x1c, 0x25, 0x56, 0xf2, 0xcf, 0x3d, 0xd6, 0xdf, 0x09, 0xa1, 0xb9, 0x22, 0x64, 0x14, 0x11,
	0x12, 0x47, 0x84, 0x58, 0xd9, 0xe5, 0xad, 0xc5, 0x9e, 0x8f, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xc5, 0x1d, 0x92, 0xc3, 0x84, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignAudienceViewService/GetCampaignAudienceView", in, out, opts...)
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
		FullMethod: "/google.ads.googleads.v1.services.CampaignAudienceViewService/GetCampaignAudienceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignAudienceViewServiceServer).GetCampaignAudienceView(ctx, req.(*GetCampaignAudienceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignAudienceViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CampaignAudienceViewService",
	HandlerType: (*CampaignAudienceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignAudienceView",
			Handler:    _CampaignAudienceViewService_GetCampaignAudienceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/campaign_audience_view_service.proto",
}
