// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/change_status_service.proto

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

// Request message for '[ChangeStatusService.GetChangeStatus][google.ads.googleads.v1.services.ChangeStatusService.GetChangeStatus]'.
type GetChangeStatusRequest struct {
	// Required. The resource name of the change status to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetChangeStatusRequest) Reset()         { *m = GetChangeStatusRequest{} }
func (m *GetChangeStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetChangeStatusRequest) ProtoMessage()    {}
func (*GetChangeStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c59aa7321bb8d53, []int{0}
}

func (m *GetChangeStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetChangeStatusRequest.Unmarshal(m, b)
}
func (m *GetChangeStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetChangeStatusRequest.Marshal(b, m, deterministic)
}
func (m *GetChangeStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChangeStatusRequest.Merge(m, src)
}
func (m *GetChangeStatusRequest) XXX_Size() int {
	return xxx_messageInfo_GetChangeStatusRequest.Size(m)
}
func (m *GetChangeStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChangeStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetChangeStatusRequest proto.InternalMessageInfo

func (m *GetChangeStatusRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetChangeStatusRequest)(nil), "google.ads.googleads.v1.services.GetChangeStatusRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/change_status_service.proto", fileDescriptor_6c59aa7321bb8d53)
}

var fileDescriptor_6c59aa7321bb8d53 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcf, 0x8a, 0xd3, 0x40,
	0x18, 0x27, 0x29, 0x08, 0x06, 0x45, 0x88, 0xa0, 0x25, 0x0a, 0x96, 0xd2, 0x43, 0xe9, 0x61, 0xc6,
	0x28, 0x82, 0x8c, 0x7a, 0x98, 0x08, 0xd6, 0x93, 0x94, 0x16, 0x7a, 0x90, 0x40, 0x98, 0x26, 0x63,
	0x1a, 0x48, 0x66, 0x6a, 0x66, 0x92, 0x8b, 0x78, 0xf1, 0x15, 0x7c, 0x03, 0x8f, 0x3e, 0x87, 0xa7,
	0x5e, 0xbd, 0xed, 0x69, 0x0f, 0x7b, 0xda, 0x37, 0xd8, 0xd3, 0x2e, 0xc9, 0x64, 0xd2, 0x74, 0xb7,
	0xa5, 0xb7, 0x1f, 0xf9, 0xfd, 0xfb, 0xbe, 0x6f, 0x62, 0xbd, 0x8f, 0x39, 0x8f, 0x53, 0x0a, 0x49,
	0x24, 0xa0, 0x82, 0x15, 0x2a, 0x5d, 0x28, 0x68, 0x5e, 0x26, 0x21, 0x15, 0x30, 0x5c, 0x13, 0x16,
	0xd3, 0x40, 0x48, 0x22, 0x0b, 0x11, 0x34, 0x9f, 0xc1, 0x26, 0xe7, 0x92, 0xdb, 0x03, 0x65, 0x01,
	0x24, 0x12, 0xa0, 0x75, 0x83, 0xd2, 0x05, 0xda, 0xed, 0xbc, 0x39, 0x96, 0x9f, 0x53, 0xc1, 0x8b,
	0xfc, 0x4e, 0x81, 0x0a, 0x76, 0x9e, 0x6b, 0xdb, 0x26, 0x81, 0x84, 0x31, 0x2e, 0x89, 0x4c, 0x38,
	0xd3, 0xec, 0xd3, 0x0e, 0x1b, 0xa6, 0x09, 0x65, 0xb2, 0x21, 0x5e, 0x74, 0x88, 0x6f, 0x09, 0x4d,
	0xa3, 0x60, 0x45, 0xd7, 0xa4, 0x4c, 0x78, 0xae, 0x04, 0x43, 0xcf, 0x7a, 0x32, 0xa5, 0xf2, 0x63,
	0xdd, 0xb8, 0xa8, 0x0b, 0xe7, 0xf4, 0x7b, 0x41, 0x85, 0xb4, 0xc7, 0xd6, 0x43, 0x3d, 0x52, 0xc0,
	0x48, 0x46, 0xfb, 0xc6, 0xc0, 0x18, 0xdf, 0xf7, 0x7a, 0xe7, 0xd8, 0x9c, 0x3f, 0xd0, 0xcc, 0x17,
	0x92, 0xd1, 0x57, 0x57, 0x86, 0xf5, 0xb8, 0x9b, 0xb0, 0x50, 0xbb, 0xda, 0xff, 0x0c, 0xeb, 0xd1,
	0xad, 0x70, 0xfb, 0x2d, 0x38, 0x75, 0x21, 0x70, 0x78, 0x1e, 0x07, 0x1e, 0x75, 0xb6, 0x97, 0x03,
	0x5d, 0xdf, 0xf0, 0xd3, 0x19, 0xde, 0xdf, 0xe0, 0xd7, 0xff, 0x8b, 0xdf, 0xe6, 0x4b, 0x1b, 0x54,
	0xd7, 0xfe, 0xb1, 0xc7, 0x7c, 0x08, 0x0b, 0x21, 0x79, 0x46, 0x73, 0x01, 0x27, 0xcd, 0xf9, 0x55,
	0x08, 0x9c, 0xfc, 0x74, 0x9e, 0x6d, 0x71, 0x7f, 0xd7, 0xd7, 0xa0, 0x4d, 0x22, 0x40, 0xc8, 0x33,
	0xef, 0xda, 0xb0, 0x46, 0x21, 0xcf, 0x4e, 0x6e, 0xe5, 0xf5, 0x0f, 0x5c, 0x68, 0x56, 0x3d, 0xc1,
	0xcc, 0xf8, 0xfa, 0xb9, 0x71, 0xc7, 0x3c, 0x25, 0x2c, 0x06, 0x3c, 0x8f, 0x61, 0x4c, 0x59, 0xfd,
	0x40, 0x70, 0xd7, 0x77, 0xfc, 0x97, 0x7c, 0xa7, 0xc1, 0x1f, 0xb3, 0x37, 0xc5, 0xf8, 0xaf, 0x39,
	0x98, 0xaa, 0x40, 0x1c, 0x09, 0xa0, 0x60, 0x85, 0x96, 0x2e, 0x68, 0x8a, 0xc5, 0x56, 0x4b, 0x7c,
	0x1c, 0x09, 0xbf, 0x95, 0xf8, 0x4b, 0xd7, 0xd7, 0x92, 0x4b, 0x73, 0xa4, 0xbe, 0x23, 0x84, 0x23,
	0x81, 0x50, 0x2b, 0x42, 0x68, 0xe9, 0x22, 0xa4, 0x65, 0xab, 0x7b, 0xf5, 0x9c, 0xaf, 0x6f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x30, 0x5d, 0xab, 0x32, 0x39, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChangeStatusServiceClient is the client API for ChangeStatusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChangeStatusServiceClient interface {
	// Returns the requested change status in full detail.
	GetChangeStatus(ctx context.Context, in *GetChangeStatusRequest, opts ...grpc.CallOption) (*resources.ChangeStatus, error)
}

type changeStatusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChangeStatusServiceClient(cc grpc.ClientConnInterface) ChangeStatusServiceClient {
	return &changeStatusServiceClient{cc}
}

func (c *changeStatusServiceClient) GetChangeStatus(ctx context.Context, in *GetChangeStatusRequest, opts ...grpc.CallOption) (*resources.ChangeStatus, error) {
	out := new(resources.ChangeStatus)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ChangeStatusService/GetChangeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChangeStatusServiceServer is the server API for ChangeStatusService service.
type ChangeStatusServiceServer interface {
	// Returns the requested change status in full detail.
	GetChangeStatus(context.Context, *GetChangeStatusRequest) (*resources.ChangeStatus, error)
}

// UnimplementedChangeStatusServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChangeStatusServiceServer struct {
}

func (*UnimplementedChangeStatusServiceServer) GetChangeStatus(ctx context.Context, req *GetChangeStatusRequest) (*resources.ChangeStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChangeStatus not implemented")
}

func RegisterChangeStatusServiceServer(s *grpc.Server, srv ChangeStatusServiceServer) {
	s.RegisterService(&_ChangeStatusService_serviceDesc, srv)
}

func _ChangeStatusService_GetChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChangeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChangeStatusServiceServer).GetChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ChangeStatusService/GetChangeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChangeStatusServiceServer).GetChangeStatus(ctx, req.(*GetChangeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChangeStatusService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.ChangeStatusService",
	HandlerType: (*ChangeStatusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChangeStatus",
			Handler:    _ChangeStatusService_GetChangeStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/change_status_service.proto",
}
