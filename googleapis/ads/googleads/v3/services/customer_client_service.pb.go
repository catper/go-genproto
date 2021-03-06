// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/customer_client_service.proto

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

// Request message for [CustomerClientService.GetCustomerClient][google.ads.googleads.v3.services.CustomerClientService.GetCustomerClient].
type GetCustomerClientRequest struct {
	// Required. The resource name of the client to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerClientRequest) Reset()         { *m = GetCustomerClientRequest{} }
func (m *GetCustomerClientRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerClientRequest) ProtoMessage()    {}
func (*GetCustomerClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_275643255b459061, []int{0}
}

func (m *GetCustomerClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerClientRequest.Unmarshal(m, b)
}
func (m *GetCustomerClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerClientRequest.Marshal(b, m, deterministic)
}
func (m *GetCustomerClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerClientRequest.Merge(m, src)
}
func (m *GetCustomerClientRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerClientRequest.Size(m)
}
func (m *GetCustomerClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerClientRequest proto.InternalMessageInfo

func (m *GetCustomerClientRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCustomerClientRequest)(nil), "google.ads.googleads.v3.services.GetCustomerClientRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/customer_client_service.proto", fileDescriptor_275643255b459061)
}

var fileDescriptor_275643255b459061 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x25, 0x29, 0x3c, 0x78, 0xe1, 0xbd, 0x85, 0x01, 0xb1, 0x44, 0xc1, 0x52, 0xba, 0x28, 0x5d,
	0xcc, 0x50, 0xb3, 0x10, 0x46, 0x14, 0xa6, 0x15, 0xaa, 0x1b, 0x29, 0x15, 0xba, 0x90, 0x40, 0x98,
	0x26, 0x63, 0x0c, 0x24, 0x99, 0x9a, 0x49, 0xb3, 0x11, 0x17, 0x8a, 0x7f, 0xe0, 0x1f, 0xb8, 0xf4,
	0x53, 0x0a, 0xae, 0xdc, 0xb9, 0x72, 0xe1, 0xca, 0xaf, 0x90, 0x74, 0x32, 0x69, 0x83, 0x0d, 0xdd,
	0x1d, 0x72, 0xce, 0xb9, 0x67, 0xee, 0xb9, 0xd1, 0x4e, 0x3c, 0xc6, 0xbc, 0x80, 0x42, 0xe2, 0x72,
	0x28, 0x60, 0x86, 0x52, 0x13, 0x72, 0x1a, 0xa7, 0xbe, 0x43, 0x39, 0x74, 0x66, 0x3c, 0x61, 0x21,
	0x8d, 0x6d, 0x27, 0xf0, 0x69, 0x94, 0xd8, 0x39, 0x01, 0xa6, 0x31, 0x4b, 0x98, 0xde, 0x10, 0x26,
	0x40, 0x5c, 0x0e, 0x0a, 0x3f, 0x48, 0x4d, 0x20, 0xfd, 0xc6, 0x61, 0x55, 0x42, 0x4c, 0x39, 0x9b,
	0xc5, 0x6b, 0x22, 0xc4, 0x68, 0x63, 0x4f, 0x1a, 0xa7, 0x3e, 0x24, 0x51, 0xc4, 0x12, 0x92, 0xf8,
	0x2c, 0xe2, 0x39, 0xbb, 0xb3, 0xc2, 0x96, 0x6c, 0xfb, 0x2b, 0xc4, 0xb5, 0x4f, 0x03, 0xd7, 0x9e,
	0xd0, 0x1b, 0x92, 0xfa, 0x2c, 0x16, 0x82, 0xe6, 0xa9, 0x56, 0x1f, 0xd0, 0xa4, 0x9f, 0x67, 0xf6,
	0x17, 0xde, 0x11, 0xbd, 0x9d, 0x51, 0x9e, 0xe8, 0x6d, 0xed, 0xbf, 0x7c, 0x96, 0x1d, 0x91, 0x90,
	0xd6, 0x95, 0x86, 0xd2, 0xfe, 0xdb, 0xab, 0x7d, 0x62, 0x75, 0xf4, 0x4f, 0x32, 0x17, 0x24, 0xa4,
	0x07, 0x4f, 0xaa, 0xb6, 0x5d, 0x9e, 0x71, 0x29, 0x36, 0xd6, 0xdf, 0x14, 0x6d, 0xeb, 0x57, 0x80,
	0x8e, 0xc0, 0xa6, 0xa6, 0x40, 0xd5, 0xab, 0x8c, 0x6e, 0xa5, 0xb7, 0xe8, 0x10, 0x94, 0x9d, 0xcd,
	0xf3, 0x0f, 0x5c, 0xde, 0xe4, 0xf1, 0xfd, 0xeb, 0x59, 0x35, 0xf5, 0x6e, 0xd6, 0xfc, 0x5d, 0x89,
	0x39, 0x96, 0xf5, 0x73, 0xd8, 0x29, 0x4e, 0x21, 0xc6, 0x70, 0xd8, 0xb9, 0x37, 0x76, 0xe7, 0xb8,
	0xbe, 0x0c, 0xcd, 0xd1, 0xd4, 0xe7, 0xc0, 0x61, 0x61, 0xef, 0x41, 0xd5, 0x5a, 0x0e, 0x0b, 0x37,
	0x2e, 0xd7, 0x33, 0xd6, 0x96, 0x35, 0xcc, 0x2e, 0x32, 0x54, 0xae, 0xce, 0x72, 0xbf, 0xc7, 0x02,
	0x12, 0x79, 0x80, 0xc5, 0x1e, 0xf4, 0x68, 0xb4, 0xb8, 0x17, 0x5c, 0x26, 0x56, 0xff, 0xa5, 0x47,
	0x12, 0xbc, 0xa8, 0xb5, 0x01, 0xc6, 0xaf, 0x6a, 0x63, 0x20, 0x06, 0x62, 0x97, 0x03, 0x01, 0x33,
	0x34, 0x36, 0x41, 0x1e, 0xcc, 0xe7, 0x52, 0x62, 0x61, 0x97, 0x5b, 0x85, 0xc4, 0x1a, 0x9b, 0x96,
	0x94, 0x7c, 0xab, 0x2d, 0xf1, 0x1d, 0x21, 0xec, 0x72, 0x84, 0x0a, 0x11, 0x42, 0x63, 0x13, 0x21,
	0x29, 0x9b, 0xfc, 0x59, 0xbc, 0xd3, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x63, 0x55, 0x92, 0x93,
	0x4c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerClientServiceClient is the client API for CustomerClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClientServiceClient interface {
	// Returns the requested client in full detail.
	GetCustomerClient(ctx context.Context, in *GetCustomerClientRequest, opts ...grpc.CallOption) (*resources.CustomerClient, error)
}

type customerClientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerClientServiceClient(cc grpc.ClientConnInterface) CustomerClientServiceClient {
	return &customerClientServiceClient{cc}
}

func (c *customerClientServiceClient) GetCustomerClient(ctx context.Context, in *GetCustomerClientRequest, opts ...grpc.CallOption) (*resources.CustomerClient, error) {
	out := new(resources.CustomerClient)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CustomerClientService/GetCustomerClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerClientServiceServer is the server API for CustomerClientService service.
type CustomerClientServiceServer interface {
	// Returns the requested client in full detail.
	GetCustomerClient(context.Context, *GetCustomerClientRequest) (*resources.CustomerClient, error)
}

// UnimplementedCustomerClientServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerClientServiceServer struct {
}

func (*UnimplementedCustomerClientServiceServer) GetCustomerClient(ctx context.Context, req *GetCustomerClientRequest) (*resources.CustomerClient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerClient not implemented")
}

func RegisterCustomerClientServiceServer(s *grpc.Server, srv CustomerClientServiceServer) {
	s.RegisterService(&_CustomerClientService_serviceDesc, srv)
}

func _CustomerClientService_GetCustomerClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerClientServiceServer).GetCustomerClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CustomerClientService/GetCustomerClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerClientServiceServer).GetCustomerClient(ctx, req.(*GetCustomerClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerClientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.CustomerClientService",
	HandlerType: (*CustomerClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerClient",
			Handler:    _CustomerClientService_GetCustomerClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/customer_client_service.proto",
}
