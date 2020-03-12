// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/geographic_view_service.proto

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

// Request message for [GeographicViewService.GetGeographicView][google.ads.googleads.v2.services.GeographicViewService.GetGeographicView].
type GetGeographicViewRequest struct {
	// Required. The resource name of the geographic view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGeographicViewRequest) Reset()         { *m = GetGeographicViewRequest{} }
func (m *GetGeographicViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetGeographicViewRequest) ProtoMessage()    {}
func (*GetGeographicViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51498bd55db0eccc, []int{0}
}

func (m *GetGeographicViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGeographicViewRequest.Unmarshal(m, b)
}
func (m *GetGeographicViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGeographicViewRequest.Marshal(b, m, deterministic)
}
func (m *GetGeographicViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGeographicViewRequest.Merge(m, src)
}
func (m *GetGeographicViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetGeographicViewRequest.Size(m)
}
func (m *GetGeographicViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGeographicViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGeographicViewRequest proto.InternalMessageInfo

func (m *GetGeographicViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetGeographicViewRequest)(nil), "google.ads.googleads.v2.services.GetGeographicViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/geographic_view_service.proto", fileDescriptor_51498bd55db0eccc)
}

var fileDescriptor_51498bd55db0eccc = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3f, 0x8b, 0xd4, 0x40,
	0x1c, 0x25, 0x39, 0x10, 0x0c, 0x5a, 0x18, 0x10, 0x97, 0x28, 0xb8, 0x1c, 0x57, 0x1c, 0x57, 0xcc,
	0x70, 0xb9, 0x42, 0x18, 0x51, 0x98, 0x45, 0x88, 0x36, 0x72, 0x9c, 0x90, 0x42, 0x02, 0x61, 0x2e,
	0xf9, 0x39, 0x37, 0x90, 0x64, 0xe2, 0x4c, 0x36, 0x5b, 0x88, 0x85, 0xe2, 0x37, 0xf0, 0x1b, 0x58,
	0xfa, 0x51, 0x16, 0xac, 0xec, 0xac, 0x2c, 0xac, 0xfc, 0x14, 0x92, 0x9d, 0x4c, 0x76, 0x83, 0x1b,
	0xb6, 0x7b, 0xcc, 0x7b, 0xef, 0xf7, 0x7e, 0x7f, 0xc6, 0x7b, 0xce, 0xa5, 0xe4, 0x05, 0x60, 0x96,
	0x6b, 0x6c, 0x60, 0x87, 0xda, 0x10, 0x6b, 0x50, 0xad, 0xc8, 0x40, 0x63, 0x0e, 0x92, 0x2b, 0x56,
	0xdf, 0x88, 0x2c, 0x6d, 0x05, 0xac, 0xd2, 0x9e, 0x40, 0xb5, 0x92, 0x8d, 0xf4, 0xe7, 0xc6, 0x84,
	0x58, 0xae, 0xd1, 0xe0, 0x47, 0x6d, 0x88, 0xac, 0x3f, 0x78, 0x32, 0x95, 0xa0, 0x40, 0xcb, 0xa5,
	0xda, 0x13, 0x61, 0x4a, 0x07, 0x8f, 0xac, 0xb1, 0x16, 0x98, 0x55, 0x95, 0x6c, 0x58, 0x23, 0x64,
	0xa5, 0x7b, 0xf6, 0xc1, 0x0e, 0x9b, 0x15, 0x02, 0xaa, 0xa6, 0x27, 0x1e, 0xef, 0x10, 0xef, 0x04,
	0x14, 0x79, 0x7a, 0x0d, 0x37, 0xac, 0x15, 0x52, 0x19, 0xc1, 0xf1, 0x0b, 0x6f, 0x16, 0x41, 0x13,
	0x0d, 0x99, 0xb1, 0x80, 0xd5, 0x15, 0xbc, 0x5f, 0x82, 0x6e, 0xfc, 0x53, 0xef, 0xae, 0x6d, 0x2b,
	0xad, 0x58, 0x09, 0x33, 0x67, 0xee, 0x9c, 0xde, 0x5e, 0x1c, 0xfd, 0xa6, 0xee, 0xd5, 0x1d, 0xcb,
	0xbc, 0x66, 0x25, 0x84, 0x5f, 0x5c, 0xef, 0xfe, 0xb8, 0xc6, 0x1b, 0x33, 0xb1, 0xff, 0xc3, 0xf1,
	0xee, 0xfd, 0x17, 0xe0, 0x13, 0x74, 0x68, 0x53, 0x68, 0xaa, 0xab, 0xe0, 0x7c, 0xd2, 0x3b, 0xec,
	0x10, 0x8d, 0x9d, 0xc7, 0xaf, 0x7e, 0xd1, 0xf1, 0x24, 0x9f, 0x7f, 0xfe, 0xf9, 0xea, 0x5e, 0xf8,
	0xe7, 0xdd, 0xe6, 0x3f, 0x8c, 0x98, 0x67, 0xd9, 0x52, 0x37, 0xb2, 0x04, 0xa5, 0xf1, 0xd9, 0xce,
	0x29, 0xba, 0x32, 0x1a, 0x9f, 0x7d, 0x0c, 0x1e, 0xae, 0xe9, 0x6c, 0x1b, 0xda, 0xa3, 0x5a, 0x68,
	0x94, 0xc9, 0x72, 0xf1, 0xc9, 0xf5, 0x4e, 0x32, 0x59, 0x1e, 0x1c, 0x6e, 0x11, 0xec, 0x5d, 0xd6,
	0x65, 0x77, 0x91, 0x4b, 0xe7, 0xed, 0xcb, 0xde, 0xcf, 0x65, 0xc1, 0x2a, 0x8e, 0xa4, 0xe2, 0x98,
	0x43, 0xb5, 0xb9, 0x17, 0xde, 0x26, 0x4e, 0xff, 0xd2, 0xa7, 0x16, 0x7c, 0x73, 0x8f, 0x22, 0x4a,
	0xbf, 0xbb, 0xf3, 0xc8, 0x14, 0xa4, 0xb9, 0x46, 0x06, 0x76, 0x28, 0x0e, 0x51, 0x1f, 0xac, 0xd7,
	0x56, 0x92, 0xd0, 0x5c, 0x27, 0x83, 0x24, 0x89, 0xc3, 0xc4, 0x4a, 0xfe, 0xba, 0x27, 0xe6, 0x9d,
	0x10, 0x9a, 0x6b, 0x42, 0x06, 0x11, 0x21, 0x71, 0x48, 0x88, 0x95, 0x5d, 0xdf, 0xda, 0xf4, 0x79,
	0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x1a, 0xda, 0xd5, 0x4c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GeographicViewServiceClient is the client API for GeographicViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeographicViewServiceClient interface {
	// Returns the requested geographic view in full detail.
	GetGeographicView(ctx context.Context, in *GetGeographicViewRequest, opts ...grpc.CallOption) (*resources.GeographicView, error)
}

type geographicViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGeographicViewServiceClient(cc grpc.ClientConnInterface) GeographicViewServiceClient {
	return &geographicViewServiceClient{cc}
}

func (c *geographicViewServiceClient) GetGeographicView(ctx context.Context, in *GetGeographicViewRequest, opts ...grpc.CallOption) (*resources.GeographicView, error) {
	out := new(resources.GeographicView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.GeographicViewService/GetGeographicView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeographicViewServiceServer is the server API for GeographicViewService service.
type GeographicViewServiceServer interface {
	// Returns the requested geographic view in full detail.
	GetGeographicView(context.Context, *GetGeographicViewRequest) (*resources.GeographicView, error)
}

// UnimplementedGeographicViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGeographicViewServiceServer struct {
}

func (*UnimplementedGeographicViewServiceServer) GetGeographicView(ctx context.Context, req *GetGeographicViewRequest) (*resources.GeographicView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGeographicView not implemented")
}

func RegisterGeographicViewServiceServer(s *grpc.Server, srv GeographicViewServiceServer) {
	s.RegisterService(&_GeographicViewService_serviceDesc, srv)
}

func _GeographicViewService_GetGeographicView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGeographicViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeographicViewServiceServer).GetGeographicView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.GeographicViewService/GetGeographicView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeographicViewServiceServer).GetGeographicView(ctx, req.(*GetGeographicViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GeographicViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.GeographicViewService",
	HandlerType: (*GeographicViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGeographicView",
			Handler:    _GeographicViewService_GetGeographicView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/geographic_view_service.proto",
}
