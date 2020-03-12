// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/feed_item_target_service.proto

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

// Request message for [FeedItemTargetService.GetFeedItemTarget][google.ads.googleads.v3.services.FeedItemTargetService.GetFeedItemTarget].
type GetFeedItemTargetRequest struct {
	// Required. The resource name of the feed item targets to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedItemTargetRequest) Reset()         { *m = GetFeedItemTargetRequest{} }
func (m *GetFeedItemTargetRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedItemTargetRequest) ProtoMessage()    {}
func (*GetFeedItemTargetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aae52ecf2ca0ddf4, []int{0}
}

func (m *GetFeedItemTargetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedItemTargetRequest.Unmarshal(m, b)
}
func (m *GetFeedItemTargetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedItemTargetRequest.Marshal(b, m, deterministic)
}
func (m *GetFeedItemTargetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedItemTargetRequest.Merge(m, src)
}
func (m *GetFeedItemTargetRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedItemTargetRequest.Size(m)
}
func (m *GetFeedItemTargetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedItemTargetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedItemTargetRequest proto.InternalMessageInfo

func (m *GetFeedItemTargetRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [FeedItemTargetService.MutateFeedItemTargets][google.ads.googleads.v3.services.FeedItemTargetService.MutateFeedItemTargets].
type MutateFeedItemTargetsRequest struct {
	// Required. The ID of the customer whose feed item targets are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual feed item targets.
	Operations           []*FeedItemTargetOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MutateFeedItemTargetsRequest) Reset()         { *m = MutateFeedItemTargetsRequest{} }
func (m *MutateFeedItemTargetsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateFeedItemTargetsRequest) ProtoMessage()    {}
func (*MutateFeedItemTargetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aae52ecf2ca0ddf4, []int{1}
}

func (m *MutateFeedItemTargetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedItemTargetsRequest.Unmarshal(m, b)
}
func (m *MutateFeedItemTargetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedItemTargetsRequest.Marshal(b, m, deterministic)
}
func (m *MutateFeedItemTargetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedItemTargetsRequest.Merge(m, src)
}
func (m *MutateFeedItemTargetsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateFeedItemTargetsRequest.Size(m)
}
func (m *MutateFeedItemTargetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedItemTargetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedItemTargetsRequest proto.InternalMessageInfo

func (m *MutateFeedItemTargetsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateFeedItemTargetsRequest) GetOperations() []*FeedItemTargetOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// A single operation (create, remove) on an feed item target.
type FeedItemTargetOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*FeedItemTargetOperation_Create
	//	*FeedItemTargetOperation_Remove
	Operation            isFeedItemTargetOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *FeedItemTargetOperation) Reset()         { *m = FeedItemTargetOperation{} }
func (m *FeedItemTargetOperation) String() string { return proto.CompactTextString(m) }
func (*FeedItemTargetOperation) ProtoMessage()    {}
func (*FeedItemTargetOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_aae52ecf2ca0ddf4, []int{2}
}

func (m *FeedItemTargetOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemTargetOperation.Unmarshal(m, b)
}
func (m *FeedItemTargetOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemTargetOperation.Marshal(b, m, deterministic)
}
func (m *FeedItemTargetOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemTargetOperation.Merge(m, src)
}
func (m *FeedItemTargetOperation) XXX_Size() int {
	return xxx_messageInfo_FeedItemTargetOperation.Size(m)
}
func (m *FeedItemTargetOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemTargetOperation.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemTargetOperation proto.InternalMessageInfo

type isFeedItemTargetOperation_Operation interface {
	isFeedItemTargetOperation_Operation()
}

type FeedItemTargetOperation_Create struct {
	Create *resources.FeedItemTarget `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type FeedItemTargetOperation_Remove struct {
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*FeedItemTargetOperation_Create) isFeedItemTargetOperation_Operation() {}

func (*FeedItemTargetOperation_Remove) isFeedItemTargetOperation_Operation() {}

func (m *FeedItemTargetOperation) GetOperation() isFeedItemTargetOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *FeedItemTargetOperation) GetCreate() *resources.FeedItemTarget {
	if x, ok := m.GetOperation().(*FeedItemTargetOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *FeedItemTargetOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*FeedItemTargetOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FeedItemTargetOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FeedItemTargetOperation_Create)(nil),
		(*FeedItemTargetOperation_Remove)(nil),
	}
}

// Response message for an feed item target mutate.
type MutateFeedItemTargetsResponse struct {
	// All results for the mutate.
	Results              []*MutateFeedItemTargetResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *MutateFeedItemTargetsResponse) Reset()         { *m = MutateFeedItemTargetsResponse{} }
func (m *MutateFeedItemTargetsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateFeedItemTargetsResponse) ProtoMessage()    {}
func (*MutateFeedItemTargetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aae52ecf2ca0ddf4, []int{3}
}

func (m *MutateFeedItemTargetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedItemTargetsResponse.Unmarshal(m, b)
}
func (m *MutateFeedItemTargetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedItemTargetsResponse.Marshal(b, m, deterministic)
}
func (m *MutateFeedItemTargetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedItemTargetsResponse.Merge(m, src)
}
func (m *MutateFeedItemTargetsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateFeedItemTargetsResponse.Size(m)
}
func (m *MutateFeedItemTargetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedItemTargetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedItemTargetsResponse proto.InternalMessageInfo

func (m *MutateFeedItemTargetsResponse) GetResults() []*MutateFeedItemTargetResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the feed item target mutate.
type MutateFeedItemTargetResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateFeedItemTargetResult) Reset()         { *m = MutateFeedItemTargetResult{} }
func (m *MutateFeedItemTargetResult) String() string { return proto.CompactTextString(m) }
func (*MutateFeedItemTargetResult) ProtoMessage()    {}
func (*MutateFeedItemTargetResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_aae52ecf2ca0ddf4, []int{4}
}

func (m *MutateFeedItemTargetResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedItemTargetResult.Unmarshal(m, b)
}
func (m *MutateFeedItemTargetResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedItemTargetResult.Marshal(b, m, deterministic)
}
func (m *MutateFeedItemTargetResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedItemTargetResult.Merge(m, src)
}
func (m *MutateFeedItemTargetResult) XXX_Size() int {
	return xxx_messageInfo_MutateFeedItemTargetResult.Size(m)
}
func (m *MutateFeedItemTargetResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedItemTargetResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedItemTargetResult proto.InternalMessageInfo

func (m *MutateFeedItemTargetResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFeedItemTargetRequest)(nil), "google.ads.googleads.v3.services.GetFeedItemTargetRequest")
	proto.RegisterType((*MutateFeedItemTargetsRequest)(nil), "google.ads.googleads.v3.services.MutateFeedItemTargetsRequest")
	proto.RegisterType((*FeedItemTargetOperation)(nil), "google.ads.googleads.v3.services.FeedItemTargetOperation")
	proto.RegisterType((*MutateFeedItemTargetsResponse)(nil), "google.ads.googleads.v3.services.MutateFeedItemTargetsResponse")
	proto.RegisterType((*MutateFeedItemTargetResult)(nil), "google.ads.googleads.v3.services.MutateFeedItemTargetResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/feed_item_target_service.proto", fileDescriptor_aae52ecf2ca0ddf4)
}

var fileDescriptor_aae52ecf2ca0ddf4 = []byte{
	// 623 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xb1, 0x6f, 0xd3, 0x4e,
	0x14, 0xfe, 0xd9, 0x91, 0xfa, 0x53, 0x2f, 0x30, 0x70, 0x52, 0xa9, 0x65, 0x8a, 0x88, 0x4c, 0x87,
	0x2a, 0x42, 0xb6, 0x5a, 0x33, 0x50, 0x43, 0xa9, 0x1c, 0x21, 0xda, 0x08, 0x01, 0x55, 0x40, 0x11,
	0x42, 0x11, 0x96, 0x6b, 0xbf, 0x1a, 0x4b, 0xb1, 0x2f, 0xdc, 0x5d, 0xc2, 0x50, 0x75, 0x80, 0x9d,
	0x89, 0xff, 0x00, 0x31, 0xf1, 0xa7, 0x54, 0x62, 0x62, 0xeb, 0xc4, 0xc0, 0xc4, 0xc8, 0xc6, 0x86,
	0xec, 0xf3, 0x25, 0x76, 0x89, 0x15, 0xd1, 0xed, 0xd9, 0xef, 0x7b, 0xdf, 0xf7, 0xbe, 0x7b, 0xef,
	0x0e, 0xed, 0x46, 0x84, 0x44, 0x43, 0xb0, 0xfc, 0x90, 0x59, 0x22, 0xcc, 0xa2, 0x89, 0x6d, 0x31,
	0xa0, 0x93, 0x38, 0x00, 0x66, 0x1d, 0x01, 0x84, 0x5e, 0xcc, 0x21, 0xf1, 0xb8, 0x4f, 0x23, 0xe0,
	0x5e, 0x91, 0x31, 0x47, 0x94, 0x70, 0x82, 0x5b, 0xa2, 0xca, 0xf4, 0x43, 0x66, 0x4e, 0x09, 0xcc,
	0x89, 0x6d, 0x4a, 0x02, 0xfd, 0x4e, 0x9d, 0x04, 0x05, 0x46, 0xc6, 0x74, 0x9e, 0x86, 0xe0, 0xd6,
	0xd7, 0x64, 0xe5, 0x28, 0xb6, 0xfc, 0x34, 0x25, 0xdc, 0xe7, 0x31, 0x49, 0x59, 0x91, 0x5d, 0x2d,
	0x65, 0x83, 0x61, 0x0c, 0xa9, 0x2c, 0xbb, 0x51, 0x4a, 0x1c, 0xc5, 0x30, 0x0c, 0xbd, 0x43, 0x78,
	0xed, 0x4f, 0x62, 0x42, 0x05, 0xc0, 0x78, 0x80, 0xb4, 0x3d, 0xe0, 0x0f, 0x01, 0xc2, 0x2e, 0x87,
	0xe4, 0x79, 0x2e, 0xd9, 0x83, 0x37, 0x63, 0x60, 0x1c, 0x6f, 0xa0, 0xcb, 0xb2, 0x2f, 0x2f, 0xf5,
	0x13, 0xd0, 0x94, 0x96, 0xb2, 0xb1, 0xdc, 0x69, 0x7c, 0x77, 0xd5, 0xde, 0x25, 0x99, 0x79, 0xe2,
	0x27, 0x60, 0x7c, 0x56, 0xd0, 0xda, 0xe3, 0x31, 0xf7, 0x39, 0x54, 0x99, 0x98, 0xa4, 0x5a, 0x47,
	0xcd, 0x60, 0xcc, 0x38, 0x49, 0x80, 0x7a, 0x71, 0x58, 0x26, 0x42, 0xf2, 0x7f, 0x37, 0xc4, 0xaf,
	0x10, 0x22, 0x23, 0xa0, 0xc2, 0x9a, 0xa6, 0xb6, 0x1a, 0x1b, 0xcd, 0xad, 0x6d, 0x73, 0xd1, 0xa9,
	0x9a, 0x55, 0xcd, 0xa7, 0x92, 0xa1, 0xe0, 0x9f, 0x31, 0x1a, 0x1f, 0x14, 0xb4, 0x5a, 0x03, 0xc6,
	0x8f, 0xd0, 0x52, 0x40, 0xc1, 0xe7, 0xc2, 0x65, 0x73, 0x6b, 0xb3, 0x56, 0x77, 0x3a, 0xab, 0x73,
	0xc2, 0xfb, 0xff, 0xf5, 0x0a, 0x0a, 0xac, 0xa1, 0x25, 0x0a, 0x09, 0x99, 0x80, 0xa6, 0x66, 0x4e,
	0xb3, 0x8c, 0xf8, 0xee, 0x34, 0xd1, 0xf2, 0xb4, 0x21, 0xe3, 0x2d, 0xba, 0x5e, 0x73, 0x6a, 0x6c,
	0x44, 0x52, 0x06, 0xb8, 0x8f, 0xfe, 0xa7, 0xc0, 0xc6, 0x43, 0x2e, 0x4f, 0xe3, 0xde, 0xe2, 0xd3,
	0x98, 0xc7, 0xd8, 0xcb, 0x49, 0x7a, 0x92, 0xcc, 0x70, 0x91, 0x5e, 0x0f, 0xc3, 0x37, 0xe7, 0xce,
	0xbd, 0x3a, 0xf2, 0xad, 0xdf, 0x0d, 0xb4, 0x52, 0xad, 0x7e, 0x26, 0x3a, 0xc0, 0x5f, 0x15, 0x74,
	0xe5, 0xaf, 0x9d, 0xc2, 0xce, 0xe2, 0xce, 0xeb, 0x16, 0x51, 0xff, 0xf7, 0x59, 0x18, 0xdd, 0x33,
	0xb7, 0x6a, 0xe2, 0xfd, 0xb7, 0x1f, 0x1f, 0x55, 0x1b, 0x6f, 0x66, 0xb7, 0xed, 0xb8, 0x92, 0xd9,
	0x91, 0x1b, 0xc8, 0xac, 0x76, 0x7e, 0xfd, 0x4a, 0x93, 0xb0, 0xda, 0x27, 0xf8, 0x97, 0x82, 0x56,
	0xe6, 0x8e, 0x09, 0xdf, 0xbf, 0xd8, 0x34, 0xe4, 0xad, 0xd0, 0x77, 0x2f, 0x5c, 0x2f, 0xf6, 0xc3,
	0x78, 0x71, 0xe6, 0x5e, 0x2d, 0xdd, 0xab, 0x5b, 0xb3, 0x5d, 0xcf, 0xed, 0x6e, 0x1b, 0xb7, 0x33,
	0xbb, 0x33, 0x7f, 0xc7, 0x25, 0xf0, 0x4e, 0xfb, 0xe4, 0xbc, 0x5b, 0x27, 0xc9, 0xd5, 0x1c, 0xa5,
	0xad, 0x5f, 0x3b, 0x75, 0xb5, 0x59, 0x47, 0x45, 0x34, 0x8a, 0x99, 0x19, 0x90, 0xa4, 0xf3, 0x4e,
	0x45, 0xeb, 0x01, 0x49, 0x16, 0x76, 0xdf, 0xd1, 0xe7, 0x6e, 0xc8, 0x41, 0xf6, 0xf2, 0x1c, 0x28,
	0x2f, 0xf7, 0x8b, 0xfa, 0x88, 0x0c, 0xfd, 0x34, 0x32, 0x09, 0x8d, 0xac, 0x08, 0xd2, 0xfc, 0x5d,
	0xb2, 0x66, 0x8a, 0xf5, 0xef, 0xf1, 0x5d, 0x19, 0x7c, 0x52, 0x1b, 0x7b, 0xae, 0xfb, 0x45, 0x6d,
	0xed, 0x09, 0x42, 0x37, 0x64, 0xa6, 0x08, 0xb3, 0xa8, 0x6f, 0x9b, 0x85, 0x30, 0x3b, 0x95, 0x90,
	0x81, 0x1b, 0xb2, 0xc1, 0x14, 0x32, 0xe8, 0xdb, 0x03, 0x09, 0xf9, 0xa9, 0xae, 0x8b, 0xff, 0x8e,
	0xe3, 0x86, 0xcc, 0x71, 0xa6, 0x20, 0xc7, 0xe9, 0xdb, 0x8e, 0x23, 0x61, 0x87, 0x4b, 0x79, 0x9f,
	0xf6, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x99, 0x51, 0xef, 0x36, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FeedItemTargetServiceClient is the client API for FeedItemTargetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedItemTargetServiceClient interface {
	// Returns the requested feed item targets in full detail.
	GetFeedItemTarget(ctx context.Context, in *GetFeedItemTargetRequest, opts ...grpc.CallOption) (*resources.FeedItemTarget, error)
	// Creates or removes feed item targets. Operation statuses are returned.
	MutateFeedItemTargets(ctx context.Context, in *MutateFeedItemTargetsRequest, opts ...grpc.CallOption) (*MutateFeedItemTargetsResponse, error)
}

type feedItemTargetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedItemTargetServiceClient(cc grpc.ClientConnInterface) FeedItemTargetServiceClient {
	return &feedItemTargetServiceClient{cc}
}

func (c *feedItemTargetServiceClient) GetFeedItemTarget(ctx context.Context, in *GetFeedItemTargetRequest, opts ...grpc.CallOption) (*resources.FeedItemTarget, error) {
	out := new(resources.FeedItemTarget)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.FeedItemTargetService/GetFeedItemTarget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedItemTargetServiceClient) MutateFeedItemTargets(ctx context.Context, in *MutateFeedItemTargetsRequest, opts ...grpc.CallOption) (*MutateFeedItemTargetsResponse, error) {
	out := new(MutateFeedItemTargetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.FeedItemTargetService/MutateFeedItemTargets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedItemTargetServiceServer is the server API for FeedItemTargetService service.
type FeedItemTargetServiceServer interface {
	// Returns the requested feed item targets in full detail.
	GetFeedItemTarget(context.Context, *GetFeedItemTargetRequest) (*resources.FeedItemTarget, error)
	// Creates or removes feed item targets. Operation statuses are returned.
	MutateFeedItemTargets(context.Context, *MutateFeedItemTargetsRequest) (*MutateFeedItemTargetsResponse, error)
}

// UnimplementedFeedItemTargetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFeedItemTargetServiceServer struct {
}

func (*UnimplementedFeedItemTargetServiceServer) GetFeedItemTarget(ctx context.Context, req *GetFeedItemTargetRequest) (*resources.FeedItemTarget, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedItemTarget not implemented")
}
func (*UnimplementedFeedItemTargetServiceServer) MutateFeedItemTargets(ctx context.Context, req *MutateFeedItemTargetsRequest) (*MutateFeedItemTargetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateFeedItemTargets not implemented")
}

func RegisterFeedItemTargetServiceServer(s *grpc.Server, srv FeedItemTargetServiceServer) {
	s.RegisterService(&_FeedItemTargetService_serviceDesc, srv)
}

func _FeedItemTargetService_GetFeedItemTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedItemTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedItemTargetServiceServer).GetFeedItemTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.FeedItemTargetService/GetFeedItemTarget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedItemTargetServiceServer).GetFeedItemTarget(ctx, req.(*GetFeedItemTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedItemTargetService_MutateFeedItemTargets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedItemTargetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedItemTargetServiceServer).MutateFeedItemTargets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.FeedItemTargetService/MutateFeedItemTargets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedItemTargetServiceServer).MutateFeedItemTargets(ctx, req.(*MutateFeedItemTargetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedItemTargetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.FeedItemTargetService",
	HandlerType: (*FeedItemTargetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeedItemTarget",
			Handler:    _FeedItemTargetService_GetFeedItemTarget_Handler,
		},
		{
			MethodName: "MutateFeedItemTargets",
			Handler:    _FeedItemTargetService_MutateFeedItemTargets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/feed_item_target_service.proto",
}
