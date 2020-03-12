// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/shopping_performance_view_service.proto

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

// Request message for
// [ShoppingPerformanceViewService.GetShoppingPerformanceView][google.ads.googleads.v3.services.ShoppingPerformanceViewService.GetShoppingPerformanceView].
type GetShoppingPerformanceViewRequest struct {
	// Required. The resource name of the Shopping performance view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShoppingPerformanceViewRequest) Reset()         { *m = GetShoppingPerformanceViewRequest{} }
func (m *GetShoppingPerformanceViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetShoppingPerformanceViewRequest) ProtoMessage()    {}
func (*GetShoppingPerformanceViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f8e4c8a85feddc3, []int{0}
}

func (m *GetShoppingPerformanceViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShoppingPerformanceViewRequest.Unmarshal(m, b)
}
func (m *GetShoppingPerformanceViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShoppingPerformanceViewRequest.Marshal(b, m, deterministic)
}
func (m *GetShoppingPerformanceViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShoppingPerformanceViewRequest.Merge(m, src)
}
func (m *GetShoppingPerformanceViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetShoppingPerformanceViewRequest.Size(m)
}
func (m *GetShoppingPerformanceViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShoppingPerformanceViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShoppingPerformanceViewRequest proto.InternalMessageInfo

func (m *GetShoppingPerformanceViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetShoppingPerformanceViewRequest)(nil), "google.ads.googleads.v3.services.GetShoppingPerformanceViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/shopping_performance_view_service.proto", fileDescriptor_9f8e4c8a85feddc3)
}

var fileDescriptor_9f8e4c8a85feddc3 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xb1, 0xae, 0xd3, 0x30,
	0x14, 0x55, 0xf2, 0x24, 0x24, 0x22, 0x58, 0xb2, 0xf0, 0x14, 0x10, 0x94, 0xc7, 0x1b, 0x9e, 0x18,
	0x6c, 0x89, 0x0c, 0x08, 0x23, 0x06, 0x97, 0xa1, 0x2c, 0x40, 0xf5, 0x2a, 0x65, 0x40, 0x91, 0x22,
	0x37, 0xb9, 0x4d, 0x2d, 0x25, 0x76, 0xb0, 0xd3, 0x74, 0x40, 0x2c, 0x8c, 0xac, 0xfc, 0x01, 0x23,
	0x03, 0x1f, 0xd2, 0x95, 0x8d, 0x89, 0x81, 0x89, 0x81, 0x6f, 0x40, 0xa9, 0xe3, 0xb4, 0x45, 0x4a,
	0xbb, 0x1d, 0xe5, 0x9c, 0x9c, 0x73, 0x7d, 0xcf, 0xf5, 0x5e, 0xe5, 0x52, 0xe6, 0x05, 0x60, 0x96,
	0x69, 0x6c, 0x60, 0x8b, 0x9a, 0x10, 0x6b, 0x50, 0x0d, 0x4f, 0x41, 0x63, 0xbd, 0x94, 0x55, 0xc5,
	0x45, 0x9e, 0x54, 0xa0, 0x16, 0x52, 0x95, 0x4c, 0xa4, 0x90, 0x34, 0x1c, 0xd6, 0x49, 0x27, 0x41,
	0x95, 0x92, 0xb5, 0xf4, 0x47, 0xe6, 0x77, 0xc4, 0x32, 0x8d, 0x7a, 0x27, 0xd4, 0x84, 0xc8, 0x3a,
	0x05, 0x74, 0x28, 0x4b, 0x81, 0x96, 0x2b, 0x75, 0x34, 0xcc, 0x84, 0x04, 0xf7, 0xac, 0x45, 0xc5,
	0x31, 0x13, 0x42, 0xd6, 0xac, 0xe6, 0x52, 0xe8, 0x8e, 0xbd, 0xb3, 0xc7, 0xa6, 0x05, 0x07, 0x51,
	0x77, 0xc4, 0x83, 0x3d, 0x62, 0xc1, 0xa1, 0xc8, 0x92, 0x39, 0x2c, 0x59, 0xc3, 0xa5, 0x32, 0x82,
	0x8b, 0xd7, 0xde, 0xc3, 0x09, 0xd4, 0xb3, 0x2e, 0x7d, 0xba, 0x0b, 0x8f, 0x38, 0xac, 0xaf, 0xe1,
	0xfd, 0x0a, 0x74, 0xed, 0x5f, 0x79, 0xb7, 0xed, 0xa4, 0x89, 0x60, 0x25, 0x9c, 0x3b, 0x23, 0xe7,
	0xea, 0xe6, 0xf8, 0xec, 0x17, 0x75, 0xaf, 0x6f, 0x59, 0xe6, 0x0d, 0x2b, 0xe1, 0xc9, 0x77, 0xd7,
	0xbb, 0x3f, 0x60, 0x36, 0x33, 0xdb, 0xf0, 0xff, 0x3a, 0x5e, 0x30, 0x1c, 0xe9, 0xbf, 0x44, 0xa7,
	0xd6, 0x89, 0x4e, 0x0e, 0x1c, 0x90, 0x41, 0x93, 0x7e, 0xe3, 0x68, 0xc0, 0xe2, 0xe2, 0xed, 0x4f,
	0x7a, 0xf8, 0xda, 0x4f, 0x3f, 0x7e, 0x7f, 0x71, 0x9f, 0xf9, 0x4f, 0xdb, 0xc2, 0x3e, 0x1c, 0x30,
	0x2f, 0xd2, 0x95, 0xae, 0x65, 0x09, 0x4a, 0xe3, 0xc7, 0x7d, 0x83, 0xff, 0xf9, 0x7d, 0x0c, 0xee,
	0x6e, 0xe8, 0xf9, 0x6e, 0x86, 0x0e, 0x55, 0x5c, 0xa3, 0x54, 0x96, 0xe3, 0xcf, 0xae, 0x77, 0x99,
	0xca, 0xf2, 0xe4, 0xa3, 0xc7, 0x8f, 0x8e, 0xaf, 0x75, 0xda, 0xb6, 0x39, 0x75, 0xde, 0x75, 0x67,
	0x8d, 0x72, 0x59, 0x30, 0x91, 0x23, 0xa9, 0x72, 0x9c, 0x83, 0xd8, 0x76, 0x8d, 0x77, 0xd1, 0xc3,
	0x57, 0xff, 0xdc, 0x82, 0xaf, 0xee, 0xd9, 0x84, 0xd2, 0x6f, 0xee, 0x68, 0x62, 0x0c, 0x69, 0xa6,
	0x91, 0x81, 0x2d, 0x8a, 0x42, 0xd4, 0x05, 0xeb, 0x8d, 0x95, 0xc4, 0x34, 0xd3, 0x71, 0x2f, 0x89,
	0xa3, 0x30, 0xb6, 0x92, 0x3f, 0xee, 0xa5, 0xf9, 0x4e, 0x08, 0xcd, 0x34, 0x21, 0xbd, 0x88, 0x90,
	0x28, 0x24, 0xc4, 0xca, 0xe6, 0x37, 0xb6, 0x73, 0x86, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x43,
	0xdb, 0x03, 0x47, 0x9c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShoppingPerformanceViewServiceClient is the client API for ShoppingPerformanceViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShoppingPerformanceViewServiceClient interface {
	// Returns the requested Shopping performance view in full detail.
	GetShoppingPerformanceView(ctx context.Context, in *GetShoppingPerformanceViewRequest, opts ...grpc.CallOption) (*resources.ShoppingPerformanceView, error)
}

type shoppingPerformanceViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShoppingPerformanceViewServiceClient(cc grpc.ClientConnInterface) ShoppingPerformanceViewServiceClient {
	return &shoppingPerformanceViewServiceClient{cc}
}

func (c *shoppingPerformanceViewServiceClient) GetShoppingPerformanceView(ctx context.Context, in *GetShoppingPerformanceViewRequest, opts ...grpc.CallOption) (*resources.ShoppingPerformanceView, error) {
	out := new(resources.ShoppingPerformanceView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.ShoppingPerformanceViewService/GetShoppingPerformanceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShoppingPerformanceViewServiceServer is the server API for ShoppingPerformanceViewService service.
type ShoppingPerformanceViewServiceServer interface {
	// Returns the requested Shopping performance view in full detail.
	GetShoppingPerformanceView(context.Context, *GetShoppingPerformanceViewRequest) (*resources.ShoppingPerformanceView, error)
}

// UnimplementedShoppingPerformanceViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShoppingPerformanceViewServiceServer struct {
}

func (*UnimplementedShoppingPerformanceViewServiceServer) GetShoppingPerformanceView(ctx context.Context, req *GetShoppingPerformanceViewRequest) (*resources.ShoppingPerformanceView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShoppingPerformanceView not implemented")
}

func RegisterShoppingPerformanceViewServiceServer(s *grpc.Server, srv ShoppingPerformanceViewServiceServer) {
	s.RegisterService(&_ShoppingPerformanceViewService_serviceDesc, srv)
}

func _ShoppingPerformanceViewService_GetShoppingPerformanceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShoppingPerformanceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingPerformanceViewServiceServer).GetShoppingPerformanceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.ShoppingPerformanceViewService/GetShoppingPerformanceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingPerformanceViewServiceServer).GetShoppingPerformanceView(ctx, req.(*GetShoppingPerformanceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShoppingPerformanceViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.ShoppingPerformanceViewService",
	HandlerType: (*ShoppingPerformanceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShoppingPerformanceView",
			Handler:    _ShoppingPerformanceViewService_GetShoppingPerformanceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/shopping_performance_view_service.proto",
}
