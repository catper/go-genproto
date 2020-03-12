// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/campaign_feed_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for [CampaignFeedService.GetCampaignFeed][google.ads.googleads.v3.services.CampaignFeedService.GetCampaignFeed].
type GetCampaignFeedRequest struct {
	// Required. The resource name of the campaign feed to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignFeedRequest) Reset()         { *m = GetCampaignFeedRequest{} }
func (m *GetCampaignFeedRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignFeedRequest) ProtoMessage()    {}
func (*GetCampaignFeedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f28fd8f5bcce23c1, []int{0}
}

func (m *GetCampaignFeedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignFeedRequest.Unmarshal(m, b)
}
func (m *GetCampaignFeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignFeedRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignFeedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignFeedRequest.Merge(m, src)
}
func (m *GetCampaignFeedRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignFeedRequest.Size(m)
}
func (m *GetCampaignFeedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignFeedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignFeedRequest proto.InternalMessageInfo

func (m *GetCampaignFeedRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CampaignFeedService.MutateCampaignFeeds][google.ads.googleads.v3.services.CampaignFeedService.MutateCampaignFeeds].
type MutateCampaignFeedsRequest struct {
	// Required. The ID of the customer whose campaign feeds are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual campaign feeds.
	Operations []*CampaignFeedOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignFeedsRequest) Reset()         { *m = MutateCampaignFeedsRequest{} }
func (m *MutateCampaignFeedsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignFeedsRequest) ProtoMessage()    {}
func (*MutateCampaignFeedsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f28fd8f5bcce23c1, []int{1}
}

func (m *MutateCampaignFeedsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignFeedsRequest.Unmarshal(m, b)
}
func (m *MutateCampaignFeedsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignFeedsRequest.Marshal(b, m, deterministic)
}
func (m *MutateCampaignFeedsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignFeedsRequest.Merge(m, src)
}
func (m *MutateCampaignFeedsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignFeedsRequest.Size(m)
}
func (m *MutateCampaignFeedsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignFeedsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignFeedsRequest proto.InternalMessageInfo

func (m *MutateCampaignFeedsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCampaignFeedsRequest) GetOperations() []*CampaignFeedOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCampaignFeedsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCampaignFeedsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a campaign feed.
type CampaignFeedOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CampaignFeedOperation_Create
	//	*CampaignFeedOperation_Update
	//	*CampaignFeedOperation_Remove
	Operation            isCampaignFeedOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *CampaignFeedOperation) Reset()         { *m = CampaignFeedOperation{} }
func (m *CampaignFeedOperation) String() string { return proto.CompactTextString(m) }
func (*CampaignFeedOperation) ProtoMessage()    {}
func (*CampaignFeedOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f28fd8f5bcce23c1, []int{2}
}

func (m *CampaignFeedOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignFeedOperation.Unmarshal(m, b)
}
func (m *CampaignFeedOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignFeedOperation.Marshal(b, m, deterministic)
}
func (m *CampaignFeedOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignFeedOperation.Merge(m, src)
}
func (m *CampaignFeedOperation) XXX_Size() int {
	return xxx_messageInfo_CampaignFeedOperation.Size(m)
}
func (m *CampaignFeedOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignFeedOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignFeedOperation proto.InternalMessageInfo

func (m *CampaignFeedOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCampaignFeedOperation_Operation interface {
	isCampaignFeedOperation_Operation()
}

type CampaignFeedOperation_Create struct {
	Create *resources.CampaignFeed `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignFeedOperation_Update struct {
	Update *resources.CampaignFeed `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type CampaignFeedOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CampaignFeedOperation_Create) isCampaignFeedOperation_Operation() {}

func (*CampaignFeedOperation_Update) isCampaignFeedOperation_Operation() {}

func (*CampaignFeedOperation_Remove) isCampaignFeedOperation_Operation() {}

func (m *CampaignFeedOperation) GetOperation() isCampaignFeedOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CampaignFeedOperation) GetCreate() *resources.CampaignFeed {
	if x, ok := m.GetOperation().(*CampaignFeedOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CampaignFeedOperation) GetUpdate() *resources.CampaignFeed {
	if x, ok := m.GetOperation().(*CampaignFeedOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *CampaignFeedOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CampaignFeedOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignFeedOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignFeedOperation_Create)(nil),
		(*CampaignFeedOperation_Update)(nil),
		(*CampaignFeedOperation_Remove)(nil),
	}
}

// Response message for a campaign feed mutate.
type MutateCampaignFeedsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCampaignFeedResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MutateCampaignFeedsResponse) Reset()         { *m = MutateCampaignFeedsResponse{} }
func (m *MutateCampaignFeedsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignFeedsResponse) ProtoMessage()    {}
func (*MutateCampaignFeedsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f28fd8f5bcce23c1, []int{3}
}

func (m *MutateCampaignFeedsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignFeedsResponse.Unmarshal(m, b)
}
func (m *MutateCampaignFeedsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignFeedsResponse.Marshal(b, m, deterministic)
}
func (m *MutateCampaignFeedsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignFeedsResponse.Merge(m, src)
}
func (m *MutateCampaignFeedsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignFeedsResponse.Size(m)
}
func (m *MutateCampaignFeedsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignFeedsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignFeedsResponse proto.InternalMessageInfo

func (m *MutateCampaignFeedsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCampaignFeedsResponse) GetResults() []*MutateCampaignFeedResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the campaign feed mutate.
type MutateCampaignFeedResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignFeedResult) Reset()         { *m = MutateCampaignFeedResult{} }
func (m *MutateCampaignFeedResult) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignFeedResult) ProtoMessage()    {}
func (*MutateCampaignFeedResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_f28fd8f5bcce23c1, []int{4}
}

func (m *MutateCampaignFeedResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignFeedResult.Unmarshal(m, b)
}
func (m *MutateCampaignFeedResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignFeedResult.Marshal(b, m, deterministic)
}
func (m *MutateCampaignFeedResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignFeedResult.Merge(m, src)
}
func (m *MutateCampaignFeedResult) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignFeedResult.Size(m)
}
func (m *MutateCampaignFeedResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignFeedResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignFeedResult proto.InternalMessageInfo

func (m *MutateCampaignFeedResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignFeedRequest)(nil), "google.ads.googleads.v3.services.GetCampaignFeedRequest")
	proto.RegisterType((*MutateCampaignFeedsRequest)(nil), "google.ads.googleads.v3.services.MutateCampaignFeedsRequest")
	proto.RegisterType((*CampaignFeedOperation)(nil), "google.ads.googleads.v3.services.CampaignFeedOperation")
	proto.RegisterType((*MutateCampaignFeedsResponse)(nil), "google.ads.googleads.v3.services.MutateCampaignFeedsResponse")
	proto.RegisterType((*MutateCampaignFeedResult)(nil), "google.ads.googleads.v3.services.MutateCampaignFeedResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/campaign_feed_service.proto", fileDescriptor_f28fd8f5bcce23c1)
}

var fileDescriptor_f28fd8f5bcce23c1 = []byte{
	// 761 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6a, 0xdb, 0x4a,
	0x14, 0xbe, 0x92, 0x2f, 0xb9, 0x37, 0xe3, 0xe4, 0x06, 0x26, 0x24, 0x57, 0x38, 0x85, 0x1a, 0x35,
	0x50, 0x63, 0x8a, 0x44, 0x6d, 0x4a, 0x8a, 0x92, 0x50, 0xe4, 0x52, 0x3b, 0x59, 0xa4, 0x09, 0x4a,
	0xf1, 0xa2, 0x18, 0xc4, 0x44, 0x1a, 0xbb, 0x22, 0x92, 0x46, 0x9d, 0x91, 0x0c, 0x21, 0x64, 0xd3,
	0x37, 0x28, 0x7d, 0x83, 0x2e, 0xfb, 0x0a, 0x5d, 0x77, 0x93, 0x6d, 0x77, 0x59, 0x75, 0x51, 0x28,
	0x64, 0xd1, 0x57, 0x68, 0x91, 0x46, 0x63, 0xcb, 0x89, 0x8d, 0x69, 0x76, 0x47, 0x73, 0xbe, 0xf3,
	0x9d, 0xff, 0x23, 0xb0, 0x33, 0x20, 0x64, 0xe0, 0x63, 0x1d, 0xb9, 0x4c, 0xe7, 0x62, 0x2a, 0x0d,
	0x9b, 0x3a, 0xc3, 0x74, 0xe8, 0x39, 0x98, 0xe9, 0x0e, 0x0a, 0x22, 0xe4, 0x0d, 0x42, 0xbb, 0x8f,
	0xb1, 0x6b, 0xe7, 0xcf, 0x5a, 0x44, 0x49, 0x4c, 0x60, 0x95, 0x9b, 0x68, 0xc8, 0x65, 0xda, 0xc8,
	0x5a, 0x1b, 0x36, 0x35, 0x61, 0x5d, 0x79, 0x32, 0x8b, 0x9f, 0x62, 0x46, 0x12, 0x7a, 0xcb, 0x01,
	0x27, 0xae, 0xdc, 0x13, 0x66, 0x91, 0xa7, 0xa3, 0x30, 0x24, 0x31, 0x8a, 0x3d, 0x12, 0xb2, 0x5c,
	0xfb, 0x7f, 0x41, 0xeb, 0xf8, 0x1e, 0x0e, 0xe3, 0x5c, 0x71, 0xbf, 0xa0, 0xe8, 0x7b, 0xd8, 0x77,
	0xed, 0x13, 0xfc, 0x06, 0x0d, 0x3d, 0x42, 0x73, 0x40, 0x1e, 0xb0, 0x9e, 0x7d, 0x9d, 0x24, 0xfd,
	0x1c, 0x15, 0x20, 0x76, 0x7a, 0x83, 0x9b, 0x46, 0x8e, 0xce, 0x62, 0x14, 0x27, 0xb9, 0x53, 0xb5,
	0x05, 0xd6, 0x3b, 0x38, 0x7e, 0x9e, 0x07, 0xdb, 0xc6, 0xd8, 0xb5, 0xf0, 0xdb, 0x04, 0xb3, 0x18,
	0xd6, 0xc0, 0xb2, 0xc8, 0xc6, 0x0e, 0x51, 0x80, 0x15, 0xa9, 0x2a, 0xd5, 0x16, 0x5b, 0xa5, 0x6f,
	0xa6, 0x6c, 0x2d, 0x09, 0xcd, 0x4b, 0x14, 0x60, 0xf5, 0xa7, 0x04, 0x2a, 0x07, 0x49, 0x8c, 0x62,
	0x5c, 0xe4, 0x61, 0x82, 0x68, 0x13, 0x94, 0x9d, 0x84, 0xc5, 0x24, 0xc0, 0xd4, 0xf6, 0xdc, 0x22,
	0x0d, 0x10, 0xef, 0xfb, 0x2e, 0xec, 0x01, 0x40, 0x22, 0x4c, 0x79, 0x45, 0x14, 0xb9, 0x5a, 0xaa,
	0x95, 0x1b, 0x5b, 0xda, 0xbc, 0x4e, 0x68, 0x45, 0x8f, 0x87, 0xc2, 0x3e, 0x67, 0x1f, 0xf3, 0xc1,
	0x87, 0x60, 0x25, 0x42, 0x34, 0xf6, 0x90, 0x6f, 0xf7, 0x91, 0xe7, 0x27, 0x14, 0x2b, 0xa5, 0xaa,
	0x54, 0xfb, 0xd7, 0xfa, 0x2f, 0x7f, 0x6e, 0xf3, 0x57, 0xf8, 0x00, 0x2c, 0x0f, 0x91, 0xef, 0xb9,
	0x28, 0xc6, 0x36, 0x09, 0xfd, 0x33, 0xe5, 0xef, 0x0c, 0xb6, 0x24, 0x1e, 0x0f, 0x43, 0xff, 0x4c,
	0x7d, 0x2f, 0x83, 0xb5, 0xa9, 0x8e, 0xe1, 0x36, 0x28, 0x27, 0x51, 0x66, 0x9c, 0x16, 0x3f, 0x33,
	0x2e, 0x37, 0x2a, 0x22, 0x0d, 0xd1, 0x1f, 0xad, 0x9d, 0xf6, 0xe7, 0x00, 0xb1, 0x53, 0x0b, 0x70,
	0x78, 0x2a, 0xc3, 0x7d, 0xb0, 0xe0, 0x50, 0x8c, 0x62, 0x5e, 0xea, 0x72, 0x43, 0x9f, 0x99, 0xfe,
	0x68, 0xcc, 0x26, 0xf2, 0xdf, 0xfb, 0xcb, 0xca, 0x09, 0x52, 0x2a, 0x4e, 0xac, 0xc8, 0x77, 0xa6,
	0xe2, 0x04, 0x50, 0x01, 0x0b, 0x14, 0x07, 0x64, 0xc8, 0x2b, 0xb6, 0x98, 0x6a, 0xf8, 0x77, 0xab,
	0x0c, 0x16, 0x47, 0x25, 0x56, 0x3f, 0x4b, 0x60, 0x63, 0xea, 0x10, 0xb0, 0x88, 0x84, 0x0c, 0xc3,
	0x36, 0x58, 0xbb, 0xd1, 0x01, 0x1b, 0x53, 0x4a, 0x68, 0xc6, 0x5a, 0x6e, 0x40, 0x11, 0x20, 0x8d,
	0x1c, 0xed, 0x38, 0x9b, 0x50, 0x6b, 0x75, 0xb2, 0x37, 0x2f, 0x52, 0x38, 0x7c, 0x05, 0xfe, 0xa1,
	0x98, 0x25, 0x7e, 0x2c, 0x86, 0xc4, 0x98, 0x3f, 0x24, 0xb7, 0xe3, 0xb2, 0x32, 0x0a, 0x4b, 0x50,
	0xa9, 0xcf, 0x80, 0x32, 0x0b, 0x94, 0x8e, 0xc4, 0x94, 0x45, 0x98, 0xdc, 0x81, 0xc6, 0x8f, 0x12,
	0x58, 0x2d, 0xda, 0x1e, 0x73, 0xdf, 0xf0, 0x8b, 0x04, 0x56, 0x6e, 0x2c, 0x18, 0x7c, 0x3a, 0x3f,
	0xe2, 0xe9, 0x3b, 0x59, 0xf9, 0xd3, 0x36, 0xaa, 0x9d, 0x2b, 0x73, 0x32, 0xf8, 0x77, 0x5f, 0xbf,
	0x7f, 0x90, 0x1f, 0x43, 0x3d, 0x3d, 0x56, 0xe7, 0x13, 0x9a, 0x5d, 0xb1, 0x8c, 0x4c, 0xaf, 0x8f,
	0xae, 0x57, 0xd6, 0x43, 0xbd, 0x7e, 0x01, 0xaf, 0x25, 0xb0, 0x3a, 0xa5, 0xbd, 0x70, 0xe7, 0x2e,
	0xd5, 0x17, 0xa7, 0xa1, 0xb2, 0x7b, 0x47, 0x6b, 0x3e, 0x53, 0x6a, 0xf7, 0xca, 0x5c, 0x2f, 0x9c,
	0x96, 0x47, 0xe3, 0x85, 0xcf, 0xd2, 0xdc, 0x52, 0x1b, 0x69, 0x9a, 0xe3, 0xbc, 0xce, 0x0b, 0xe0,
	0xdd, 0xfa, 0xc5, 0x64, 0x96, 0x46, 0x90, 0x79, 0x32, 0xa4, 0x7a, 0x65, 0xe3, 0xd2, 0x54, 0xc6,
	0xd1, 0xe4, 0x52, 0xe4, 0x31, 0xcd, 0x21, 0x41, 0xeb, 0x97, 0x04, 0x36, 0x1d, 0x12, 0xcc, 0x8d,
	0xbc, 0xa5, 0x4c, 0x99, 0x87, 0xa3, 0xf4, 0x02, 0x1c, 0x49, 0xaf, 0xf7, 0x72, 0xeb, 0x01, 0xf1,
	0x51, 0x38, 0xd0, 0x08, 0x1d, 0xe8, 0x03, 0x1c, 0x66, 0xf7, 0x41, 0x1f, 0xfb, 0x9b, 0xfd, 0xff,
	0xda, 0x16, 0xc2, 0x47, 0xb9, 0xd4, 0x31, 0xcd, 0x4f, 0x72, 0xb5, 0xc3, 0x09, 0x4d, 0x97, 0x69,
	0x5c, 0x4c, 0xa5, 0x6e, 0x53, 0xcb, 0x1d, 0xb3, 0x4b, 0x01, 0xe9, 0x99, 0x2e, 0xeb, 0x8d, 0x20,
	0xbd, 0x6e, 0xb3, 0x27, 0x20, 0xd7, 0xf2, 0x26, 0x7f, 0x37, 0x0c, 0xd3, 0x65, 0x86, 0x31, 0x02,
	0x19, 0x46, 0xb7, 0x69, 0x18, 0x02, 0x76, 0xb2, 0x90, 0xc5, 0xd9, 0xfc, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x53, 0x2c, 0xb0, 0x93, 0x66, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CampaignFeedServiceClient is the client API for CampaignFeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignFeedServiceClient interface {
	// Returns the requested campaign feed in full detail.
	GetCampaignFeed(ctx context.Context, in *GetCampaignFeedRequest, opts ...grpc.CallOption) (*resources.CampaignFeed, error)
	// Creates, updates, or removes campaign feeds. Operation statuses are
	// returned.
	MutateCampaignFeeds(ctx context.Context, in *MutateCampaignFeedsRequest, opts ...grpc.CallOption) (*MutateCampaignFeedsResponse, error)
}

type campaignFeedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignFeedServiceClient(cc grpc.ClientConnInterface) CampaignFeedServiceClient {
	return &campaignFeedServiceClient{cc}
}

func (c *campaignFeedServiceClient) GetCampaignFeed(ctx context.Context, in *GetCampaignFeedRequest, opts ...grpc.CallOption) (*resources.CampaignFeed, error) {
	out := new(resources.CampaignFeed)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CampaignFeedService/GetCampaignFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignFeedServiceClient) MutateCampaignFeeds(ctx context.Context, in *MutateCampaignFeedsRequest, opts ...grpc.CallOption) (*MutateCampaignFeedsResponse, error) {
	out := new(MutateCampaignFeedsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CampaignFeedService/MutateCampaignFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignFeedServiceServer is the server API for CampaignFeedService service.
type CampaignFeedServiceServer interface {
	// Returns the requested campaign feed in full detail.
	GetCampaignFeed(context.Context, *GetCampaignFeedRequest) (*resources.CampaignFeed, error)
	// Creates, updates, or removes campaign feeds. Operation statuses are
	// returned.
	MutateCampaignFeeds(context.Context, *MutateCampaignFeedsRequest) (*MutateCampaignFeedsResponse, error)
}

// UnimplementedCampaignFeedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCampaignFeedServiceServer struct {
}

func (*UnimplementedCampaignFeedServiceServer) GetCampaignFeed(ctx context.Context, req *GetCampaignFeedRequest) (*resources.CampaignFeed, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCampaignFeed not implemented")
}
func (*UnimplementedCampaignFeedServiceServer) MutateCampaignFeeds(ctx context.Context, req *MutateCampaignFeedsRequest) (*MutateCampaignFeedsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCampaignFeeds not implemented")
}

func RegisterCampaignFeedServiceServer(s *grpc.Server, srv CampaignFeedServiceServer) {
	s.RegisterService(&_CampaignFeedService_serviceDesc, srv)
}

func _CampaignFeedService_GetCampaignFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignFeedServiceServer).GetCampaignFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CampaignFeedService/GetCampaignFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignFeedServiceServer).GetCampaignFeed(ctx, req.(*GetCampaignFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignFeedService_MutateCampaignFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignFeedServiceServer).MutateCampaignFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CampaignFeedService/MutateCampaignFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignFeedServiceServer).MutateCampaignFeeds(ctx, req.(*MutateCampaignFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignFeedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.CampaignFeedService",
	HandlerType: (*CampaignFeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignFeed",
			Handler:    _CampaignFeedService_GetCampaignFeed_Handler,
		},
		{
			MethodName: "MutateCampaignFeeds",
			Handler:    _CampaignFeedService_MutateCampaignFeeds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/campaign_feed_service.proto",
}
