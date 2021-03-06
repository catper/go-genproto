// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/click_view_service.proto

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

// Request message for [ClickViewService.GetClickView][google.ads.googleads.v3.services.ClickViewService.GetClickView].
type GetClickViewRequest struct {
	// Required. The resource name of the click view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClickViewRequest) Reset()         { *m = GetClickViewRequest{} }
func (m *GetClickViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetClickViewRequest) ProtoMessage()    {}
func (*GetClickViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_894a9152e8b8f104, []int{0}
}

func (m *GetClickViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClickViewRequest.Unmarshal(m, b)
}
func (m *GetClickViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClickViewRequest.Marshal(b, m, deterministic)
}
func (m *GetClickViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClickViewRequest.Merge(m, src)
}
func (m *GetClickViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetClickViewRequest.Size(m)
}
func (m *GetClickViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClickViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetClickViewRequest proto.InternalMessageInfo

func (m *GetClickViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetClickViewRequest)(nil), "google.ads.googleads.v3.services.GetClickViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/click_view_service.proto", fileDescriptor_894a9152e8b8f104)
}

var fileDescriptor_894a9152e8b8f104 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcf, 0x8a, 0xd3, 0x40,
	0x18, 0x27, 0x29, 0x08, 0x86, 0x0a, 0x12, 0x11, 0x4b, 0x14, 0x2c, 0xa5, 0x87, 0x52, 0xca, 0x0c,
	0x34, 0x78, 0x70, 0x44, 0x64, 0xaa, 0x50, 0x4f, 0x52, 0x2a, 0xe4, 0x20, 0x81, 0x30, 0x4d, 0x3e,
	0xe3, 0x60, 0x92, 0xa9, 0x99, 0x34, 0x3d, 0x88, 0x17, 0x5f, 0xc1, 0x37, 0xf0, 0xe8, 0x1b, 0xf8,
	0x0a, 0xbd, 0xee, 0x6d, 0x4f, 0x7b, 0xd8, 0xd3, 0xee, 0x2b, 0xec, 0x61, 0x49, 0x27, 0x93, 0xb6,
	0xcb, 0x96, 0xde, 0x7e, 0xcc, 0xef, 0xcf, 0xf7, 0x7d, 0xbf, 0xc4, 0x7a, 0x1d, 0x0b, 0x11, 0x27,
	0x80, 0x59, 0x24, 0xb1, 0x82, 0x15, 0x2a, 0x5d, 0x2c, 0x21, 0x2f, 0x79, 0x08, 0x12, 0x87, 0x09,
	0x0f, 0xbf, 0x07, 0x25, 0x87, 0x75, 0x50, 0xbf, 0xa1, 0x65, 0x2e, 0x0a, 0x61, 0x77, 0x95, 0x1e,
	0xb1, 0x48, 0xa2, 0xc6, 0x8a, 0x4a, 0x17, 0x69, 0xab, 0x33, 0x3e, 0x16, 0x9e, 0x83, 0x14, 0xab,
	0xfc, 0x30, 0x5d, 0xa5, 0x3a, 0x2f, 0xb4, 0x67, 0xc9, 0x31, 0xcb, 0x32, 0x51, 0xb0, 0x82, 0x8b,
	0x4c, 0xd6, 0xec, 0xb3, 0x3d, 0x36, 0x4c, 0x38, 0x64, 0x45, 0x4d, 0xbc, 0xdc, 0x23, 0xbe, 0x72,
	0x48, 0xa2, 0x60, 0x01, 0xdf, 0x58, 0xc9, 0x45, 0xae, 0x04, 0xbd, 0x77, 0xd6, 0x93, 0x29, 0x14,
	0xef, 0xab, 0x71, 0x1e, 0x87, 0xf5, 0x1c, 0x7e, 0xac, 0x40, 0x16, 0xf6, 0xc0, 0x7a, 0xa4, 0x97,
	0x09, 0x32, 0x96, 0x42, 0xc7, 0xe8, 0x1a, 0x83, 0x87, 0x93, 0xd6, 0x05, 0x35, 0xe7, 0x6d, 0xcd,
	0x7c, 0x62, 0x29, 0x8c, 0xaf, 0x0d, 0xeb, 0x71, 0x63, 0xff, 0xac, 0x4e, 0xb4, 0xff, 0x1b, 0x56,
	0x7b, 0x3f, 0xd6, 0x7e, 0x85, 0x4e, 0xb5, 0x82, 0xee, 0x59, 0xc3, 0x19, 0x1d, 0xb5, 0x35, 0x55,
	0xa1, 0xc6, 0xd4, 0xfb, 0x70, 0x4e, 0x0f, 0xb7, 0xfe, 0x7d, 0x76, 0xf9, 0xc7, 0x44, 0xf6, 0xa8,
	0xea, 0xf6, 0xe7, 0x01, 0xf3, 0x36, 0x5c, 0xc9, 0x42, 0xa4, 0x90, 0x4b, 0x3c, 0x54, 0x65, 0x57,
	0x09, 0x12, 0x0f, 0x7f, 0x39, 0xcf, 0x37, 0xb4, 0xb3, 0x1b, 0x55, 0xa3, 0x25, 0x97, 0x28, 0x14,
	0xe9, 0xe4, 0xc6, 0xb0, 0xfa, 0xa1, 0x48, 0x4f, 0x5e, 0x33, 0x79, 0x7a, 0xb7, 0x93, 0x59, 0x55,
	0xf7, 0xcc, 0xf8, 0xf2, 0xb1, 0xb6, 0xc6, 0x22, 0x61, 0x59, 0x8c, 0x44, 0x1e, 0xe3, 0x18, 0xb2,
	0xed, 0xc7, 0xc0, 0xbb, 0x61, 0xc7, 0x7f, 0xbc, 0x37, 0x1a, 0xfc, 0x35, 0x5b, 0x53, 0x4a, 0xff,
	0x99, 0xdd, 0xa9, 0x0a, 0xa4, 0x91, 0x44, 0x0a, 0x56, 0xc8, 0x73, 0x51, 0x3d, 0x58, 0x6e, 0xb4,
	0xc4, 0xa7, 0x91, 0xf4, 0x1b, 0x89, 0xef, 0xb9, 0xbe, 0x96, 0x5c, 0x99, 0x7d, 0xf5, 0x4e, 0x08,
	0x8d, 0x24, 0x21, 0x8d, 0x88, 0x10, 0xcf, 0x25, 0x44, 0xcb, 0x16, 0x0f, 0xb6, 0x7b, 0xba, 0xb7,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xc9, 0xf7, 0x89, 0x1f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ClickViewServiceClient is the client API for ClickViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClickViewServiceClient interface {
	// Returns the requested click view in full detail.
	GetClickView(ctx context.Context, in *GetClickViewRequest, opts ...grpc.CallOption) (*resources.ClickView, error)
}

type clickViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClickViewServiceClient(cc grpc.ClientConnInterface) ClickViewServiceClient {
	return &clickViewServiceClient{cc}
}

func (c *clickViewServiceClient) GetClickView(ctx context.Context, in *GetClickViewRequest, opts ...grpc.CallOption) (*resources.ClickView, error) {
	out := new(resources.ClickView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.ClickViewService/GetClickView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClickViewServiceServer is the server API for ClickViewService service.
type ClickViewServiceServer interface {
	// Returns the requested click view in full detail.
	GetClickView(context.Context, *GetClickViewRequest) (*resources.ClickView, error)
}

// UnimplementedClickViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClickViewServiceServer struct {
}

func (*UnimplementedClickViewServiceServer) GetClickView(ctx context.Context, req *GetClickViewRequest) (*resources.ClickView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClickView not implemented")
}

func RegisterClickViewServiceServer(s *grpc.Server, srv ClickViewServiceServer) {
	s.RegisterService(&_ClickViewService_serviceDesc, srv)
}

func _ClickViewService_GetClickView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClickViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClickViewServiceServer).GetClickView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.ClickViewService/GetClickView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClickViewServiceServer).GetClickView(ctx, req.(*GetClickViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClickViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.ClickViewService",
	HandlerType: (*ClickViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClickView",
			Handler:    _ClickViewService_GetClickView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/click_view_service.proto",
}
