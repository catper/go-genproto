// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/topic_constant_service.proto

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

// Request message for [TopicConstantService.GetTopicConstant][google.ads.googleads.v2.services.TopicConstantService.GetTopicConstant].
type GetTopicConstantRequest struct {
	// Required. Resource name of the Topic to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopicConstantRequest) Reset()         { *m = GetTopicConstantRequest{} }
func (m *GetTopicConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopicConstantRequest) ProtoMessage()    {}
func (*GetTopicConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fa68112b91a9250, []int{0}
}

func (m *GetTopicConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopicConstantRequest.Unmarshal(m, b)
}
func (m *GetTopicConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopicConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetTopicConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopicConstantRequest.Merge(m, src)
}
func (m *GetTopicConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopicConstantRequest.Size(m)
}
func (m *GetTopicConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopicConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopicConstantRequest proto.InternalMessageInfo

func (m *GetTopicConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetTopicConstantRequest)(nil), "google.ads.googleads.v2.services.GetTopicConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/topic_constant_service.proto", fileDescriptor_8fa68112b91a9250)
}

var fileDescriptor_8fa68112b91a9250 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0x25, 0x29, 0x08, 0x06, 0x05, 0x09, 0x42, 0x6b, 0x14, 0x2c, 0xa5, 0x48, 0xf1, 0x30, 0x23,
	0x11, 0x04, 0x47, 0x7b, 0x98, 0xf6, 0x50, 0x4f, 0x52, 0xaa, 0xf4, 0x20, 0x81, 0x30, 0x4d, 0xc6,
	0x38, 0x90, 0xcc, 0xc4, 0xcc, 0x34, 0x17, 0x11, 0xc4, 0xbf, 0xe0, 0x3f, 0xf0, 0xe8, 0x7f, 0xf0,
	0x0f, 0xf4, 0xea, 0xcd, 0xd3, 0x1e, 0xf6, 0xb4, 0xd7, 0xfd, 0x03, 0x4b, 0x3a, 0x99, 0xb4, 0xd9,
	0xdd, 0xd0, 0xdb, 0x63, 0xde, 0x7b, 0xdf, 0xfb, 0xbe, 0x97, 0x38, 0xd3, 0x44, 0x88, 0x24, 0xa5,
	0x90, 0xc4, 0x12, 0x6a, 0x58, 0xa1, 0xd2, 0x87, 0x92, 0x16, 0x25, 0x8b, 0xa8, 0x84, 0x4a, 0xe4,
	0x2c, 0x0a, 0x23, 0xc1, 0xa5, 0x22, 0x5c, 0x85, 0xf5, 0x3b, 0xc8, 0x0b, 0xa1, 0x84, 0x3b, 0xd4,
	0x1e, 0x40, 0x62, 0x09, 0x1a, 0x3b, 0x28, 0x7d, 0x60, 0xec, 0xde, 0xab, 0xae, 0x80, 0x82, 0x4a,
	0xb1, 0x2d, 0x6e, 0x26, 0xe8, 0xc9, 0xde, 0x13, 0xe3, 0xcb, 0x19, 0x24, 0x9c, 0x0b, 0x45, 0x14,
	0x13, 0x5c, 0xd6, 0x6c, 0xff, 0x88, 0x8d, 0x52, 0x46, 0x1b, 0xdb, 0xd3, 0x23, 0xe2, 0x33, 0xa3,
	0x69, 0x1c, 0x6e, 0xe8, 0x17, 0x52, 0x32, 0x51, 0x68, 0xc1, 0x68, 0xee, 0xf4, 0x17, 0x54, 0x7d,
	0xac, 0x22, 0xe7, 0x75, 0xe2, 0x8a, 0x7e, 0xdd, 0x52, 0xa9, 0xdc, 0x89, 0x73, 0xdf, 0x2c, 0x15,
	0x72, 0x92, 0xd1, 0x81, 0x35, 0xb4, 0x26, 0x77, 0x67, 0xbd, 0x33, 0x6c, 0xaf, 0xee, 0x19, 0xe6,
	0x3d, 0xc9, 0xa8, 0x7f, 0x69, 0x39, 0x0f, 0x5b, 0x23, 0x3e, 0xe8, 0x73, 0xdd, 0xbf, 0x96, 0xf3,
	0xe0, 0xfa, 0x78, 0xf7, 0x35, 0x38, 0xd5, 0x12, 0xe8, 0x58, 0xc9, 0x7b, 0xd1, 0x69, 0x6d, 0xea,
	0x03, 0x2d, 0xe3, 0xe8, 0xed, 0x7f, 0xdc, 0xbe, 0xe2, 0xe7, 0xbf, 0xf3, 0x5f, 0xf6, 0x33, 0x77,
	0x5c, 0x75, 0xfe, 0xad, 0xc5, 0x4c, 0xd5, 0xb1, 0x53, 0xc2, 0xe7, 0xdf, 0xbd, 0xc7, 0x3b, 0x3c,
	0x38, 0xc4, 0xd4, 0x28, 0x67, 0x12, 0x44, 0x22, 0x9b, 0xfd, 0xb0, 0x9d, 0x71, 0x24, 0xb2, 0x93,
	0xd7, 0xcc, 0x1e, 0xdd, 0xd6, 0xcd, 0xb2, 0xaa, 0x7f, 0x69, 0x7d, 0x7a, 0x57, 0xdb, 0x13, 0x91,
	0x12, 0x9e, 0x00, 0x51, 0x24, 0x30, 0xa1, 0x7c, 0xff, 0x71, 0xe0, 0x21, 0xb0, 0xfb, 0x87, 0x7c,
	0x63, 0xc0, 0x6f, 0xbb, 0xb7, 0xc0, 0xf8, 0x8f, 0x3d, 0x5c, 0xe8, 0x81, 0x38, 0x96, 0x40, 0xc3,
	0x0a, 0xad, 0x7d, 0x50, 0x07, 0xcb, 0x9d, 0x91, 0x04, 0x38, 0x96, 0x41, 0x23, 0x09, 0xd6, 0x7e,
	0x60, 0x24, 0x17, 0xf6, 0x58, 0xbf, 0x23, 0x84, 0x63, 0x89, 0x50, 0x23, 0x42, 0x68, 0xed, 0x23,
	0x64, 0x64, 0x9b, 0x3b, 0xfb, 0x3d, 0x5f, 0x5e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xa9, 0xae,
	0xfa, 0x37, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TopicConstantServiceClient is the client API for TopicConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TopicConstantServiceClient interface {
	// Returns the requested topic constant in full detail.
	GetTopicConstant(ctx context.Context, in *GetTopicConstantRequest, opts ...grpc.CallOption) (*resources.TopicConstant, error)
}

type topicConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTopicConstantServiceClient(cc grpc.ClientConnInterface) TopicConstantServiceClient {
	return &topicConstantServiceClient{cc}
}

func (c *topicConstantServiceClient) GetTopicConstant(ctx context.Context, in *GetTopicConstantRequest, opts ...grpc.CallOption) (*resources.TopicConstant, error) {
	out := new(resources.TopicConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.TopicConstantService/GetTopicConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicConstantServiceServer is the server API for TopicConstantService service.
type TopicConstantServiceServer interface {
	// Returns the requested topic constant in full detail.
	GetTopicConstant(context.Context, *GetTopicConstantRequest) (*resources.TopicConstant, error)
}

// UnimplementedTopicConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTopicConstantServiceServer struct {
}

func (*UnimplementedTopicConstantServiceServer) GetTopicConstant(ctx context.Context, req *GetTopicConstantRequest) (*resources.TopicConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopicConstant not implemented")
}

func RegisterTopicConstantServiceServer(s *grpc.Server, srv TopicConstantServiceServer) {
	s.RegisterService(&_TopicConstantService_serviceDesc, srv)
}

func _TopicConstantService_GetTopicConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicConstantServiceServer).GetTopicConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.TopicConstantService/GetTopicConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicConstantServiceServer).GetTopicConstant(ctx, req.(*GetTopicConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TopicConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.TopicConstantService",
	HandlerType: (*TopicConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTopicConstant",
			Handler:    _TopicConstantService_GetTopicConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/topic_constant_service.proto",
}
