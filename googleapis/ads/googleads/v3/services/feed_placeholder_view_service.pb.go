// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/feed_placeholder_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
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

// Request message for [FeedPlaceholderViewService.GetFeedPlaceholderView][google.ads.googleads.v3.services.FeedPlaceholderViewService.GetFeedPlaceholderView].
type GetFeedPlaceholderViewRequest struct {
	// Required. The resource name of the feed placeholder view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedPlaceholderViewRequest) Reset()         { *m = GetFeedPlaceholderViewRequest{} }
func (m *GetFeedPlaceholderViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedPlaceholderViewRequest) ProtoMessage()    {}
func (*GetFeedPlaceholderViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cdfc60d4e9a2d03, []int{0}
}

func (m *GetFeedPlaceholderViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedPlaceholderViewRequest.Unmarshal(m, b)
}
func (m *GetFeedPlaceholderViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedPlaceholderViewRequest.Marshal(b, m, deterministic)
}
func (m *GetFeedPlaceholderViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedPlaceholderViewRequest.Merge(m, src)
}
func (m *GetFeedPlaceholderViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedPlaceholderViewRequest.Size(m)
}
func (m *GetFeedPlaceholderViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedPlaceholderViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedPlaceholderViewRequest proto.InternalMessageInfo

func (m *GetFeedPlaceholderViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFeedPlaceholderViewRequest)(nil), "google.ads.googleads.v3.services.GetFeedPlaceholderViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/feed_placeholder_view_service.proto", fileDescriptor_3cdfc60d4e9a2d03)
}

var fileDescriptor_3cdfc60d4e9a2d03 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x8b, 0xd4, 0x40,
	0x18, 0x25, 0x39, 0x10, 0x0c, 0xda, 0xa4, 0xd0, 0x23, 0x2a, 0xb7, 0x1c, 0x57, 0x1c, 0x57, 0xcc,
	0x80, 0x81, 0x43, 0x46, 0x0e, 0x99, 0x45, 0x5c, 0x6d, 0x8e, 0xe5, 0x84, 0x14, 0x12, 0x08, 0x73,
	0x99, 0xef, 0xb2, 0x03, 0x49, 0x26, 0xce, 0x64, 0xb3, 0x85, 0xd8, 0x58, 0xf8, 0x07, 0xac, 0x6d,
	0x2c, 0xfd, 0x29, 0xdb, 0xda, 0x59, 0x29, 0x58, 0xf9, 0x2b, 0x24, 0x3b, 0x99, 0xec, 0xae, 0x6c,
	0x6e, 0xbb, 0x47, 0xde, 0xfb, 0xde, 0xfb, 0xe6, 0x7b, 0xf1, 0x5e, 0x66, 0x52, 0x66, 0x39, 0x60,
	0xc6, 0x35, 0x36, 0xb0, 0x45, 0x4d, 0x88, 0x35, 0xa8, 0x46, 0xa4, 0xa0, 0xf1, 0x0d, 0x00, 0x4f,
	0xaa, 0x9c, 0xa5, 0x30, 0x93, 0x39, 0x07, 0x95, 0x34, 0x02, 0x16, 0x49, 0x47, 0xa3, 0x4a, 0xc9,
	0x5a, 0xfa, 0x23, 0x33, 0x8a, 0x18, 0xd7, 0xa8, 0x77, 0x41, 0x4d, 0x88, 0xac, 0x4b, 0x70, 0x31,
	0x94, 0xa3, 0x40, 0xcb, 0xb9, 0x1a, 0x0c, 0x32, 0x01, 0xc1, 0x63, 0x3b, 0x5e, 0x09, 0xcc, 0xca,
	0x52, 0xd6, 0xac, 0x16, 0xb2, 0xd4, 0x1d, 0xfb, 0x70, 0x83, 0x4d, 0x73, 0x01, 0x65, 0xdd, 0x11,
	0x47, 0x1b, 0xc4, 0x8d, 0x80, 0x9c, 0x27, 0xd7, 0x30, 0x63, 0x8d, 0x90, 0xca, 0x08, 0x8e, 0xdf,
	0x78, 0x4f, 0x26, 0x50, 0xbf, 0x02, 0xe0, 0xd3, 0x75, 0x70, 0x24, 0x60, 0x71, 0x05, 0xef, 0xe7,
	0xa0, 0x6b, 0xff, 0xd4, 0xbb, 0x6f, 0x37, 0x4c, 0x4a, 0x56, 0xc0, 0xa1, 0x33, 0x72, 0x4e, 0xef,
	0x8e, 0x0f, 0x7e, 0x51, 0xf7, 0xea, 0x9e, 0x65, 0x2e, 0x59, 0x01, 0x4f, 0xbf, 0xba, 0x5e, 0xb0,
	0xc3, 0xe8, 0xad, 0xb9, 0x80, 0xff, 0xdb, 0xf1, 0x1e, 0xec, 0x8e, 0xf2, 0x5f, 0xa0, 0x7d, 0xe7,
	0x43, 0xb7, 0x2e, 0x19, 0x9c, 0x0f, 0x1a, 0xf4, 0xd7, 0x45, 0x3b, 0xc6, 0x8f, 0x2f, 0x7f, 0xd2,
	0xed, 0xd7, 0x7d, 0xfa, 0xf1, 0xe7, 0x8b, 0xfb, 0xcc, 0x3f, 0x6f, 0x8b, 0xf9, 0xb0, 0xc5, 0x5c,
	0xa4, 0x73, 0x5d, 0xcb, 0x02, 0x94, 0xc6, 0x67, 0xab, 0xa6, 0xfe, 0xf3, 0xd2, 0xf8, 0xec, 0x63,
	0xf0, 0x68, 0x49, 0x0f, 0xd7, 0xf1, 0x1d, 0xaa, 0x84, 0x46, 0xa9, 0x2c, 0xc6, 0x9f, 0x5d, 0xef,
	0x24, 0x95, 0xc5, 0xde, 0xb7, 0x8e, 0x8f, 0x86, 0xaf, 0x38, 0x6d, 0x4b, 0x9b, 0x3a, 0xef, 0x5e,
	0x77, 0x26, 0x99, 0xcc, 0x59, 0x99, 0x21, 0xa9, 0x32, 0x9c, 0x41, 0xb9, 0xaa, 0x14, 0xaf, 0x63,
	0x87, 0x7f, 0xea, 0xe7, 0x16, 0x7c, 0x73, 0x0f, 0x26, 0x94, 0x7e, 0x77, 0x47, 0x13, 0x63, 0x48,
	0xb9, 0x46, 0x06, 0xb6, 0x28, 0x0a, 0x51, 0x17, 0xac, 0x97, 0x56, 0x12, 0x53, 0xae, 0xe3, 0x5e,
	0x12, 0x47, 0x61, 0x6c, 0x25, 0x7f, 0xdd, 0x13, 0xf3, 0x9d, 0x10, 0xca, 0x35, 0x21, 0xbd, 0x88,
	0x90, 0x28, 0x24, 0xc4, 0xca, 0xae, 0xef, 0xac, 0xf6, 0x0c, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x7e, 0xa7, 0x70, 0x6c, 0x7b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FeedPlaceholderViewServiceClient is the client API for FeedPlaceholderViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedPlaceholderViewServiceClient interface {
	// Returns the requested feed placeholder view in full detail.
	GetFeedPlaceholderView(ctx context.Context, in *GetFeedPlaceholderViewRequest, opts ...grpc.CallOption) (*resources.FeedPlaceholderView, error)
}

type feedPlaceholderViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedPlaceholderViewServiceClient(cc grpc.ClientConnInterface) FeedPlaceholderViewServiceClient {
	return &feedPlaceholderViewServiceClient{cc}
}

func (c *feedPlaceholderViewServiceClient) GetFeedPlaceholderView(ctx context.Context, in *GetFeedPlaceholderViewRequest, opts ...grpc.CallOption) (*resources.FeedPlaceholderView, error) {
	out := new(resources.FeedPlaceholderView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.FeedPlaceholderViewService/GetFeedPlaceholderView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedPlaceholderViewServiceServer is the server API for FeedPlaceholderViewService service.
type FeedPlaceholderViewServiceServer interface {
	// Returns the requested feed placeholder view in full detail.
	GetFeedPlaceholderView(context.Context, *GetFeedPlaceholderViewRequest) (*resources.FeedPlaceholderView, error)
}

// UnimplementedFeedPlaceholderViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFeedPlaceholderViewServiceServer struct {
}

func (*UnimplementedFeedPlaceholderViewServiceServer) GetFeedPlaceholderView(ctx context.Context, req *GetFeedPlaceholderViewRequest) (*resources.FeedPlaceholderView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedPlaceholderView not implemented")
}

func RegisterFeedPlaceholderViewServiceServer(s *grpc.Server, srv FeedPlaceholderViewServiceServer) {
	s.RegisterService(&_FeedPlaceholderViewService_serviceDesc, srv)
}

func _FeedPlaceholderViewService_GetFeedPlaceholderView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedPlaceholderViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedPlaceholderViewServiceServer).GetFeedPlaceholderView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.FeedPlaceholderViewService/GetFeedPlaceholderView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedPlaceholderViewServiceServer).GetFeedPlaceholderView(ctx, req.(*GetFeedPlaceholderViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedPlaceholderViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.FeedPlaceholderViewService",
	HandlerType: (*FeedPlaceholderViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeedPlaceholderView",
			Handler:    _FeedPlaceholderViewService_GetFeedPlaceholderView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/feed_placeholder_view_service.proto",
}
