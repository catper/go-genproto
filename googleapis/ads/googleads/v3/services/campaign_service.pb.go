// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/campaign_service.proto

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

// Request message for [CampaignService.GetCampaign][google.ads.googleads.v3.services.CampaignService.GetCampaign].
type GetCampaignRequest struct {
	// Required. The resource name of the campaign to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignRequest) Reset()         { *m = GetCampaignRequest{} }
func (m *GetCampaignRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignRequest) ProtoMessage()    {}
func (*GetCampaignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6454f925b0d531, []int{0}
}

func (m *GetCampaignRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignRequest.Unmarshal(m, b)
}
func (m *GetCampaignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignRequest.Merge(m, src)
}
func (m *GetCampaignRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignRequest.Size(m)
}
func (m *GetCampaignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignRequest proto.InternalMessageInfo

func (m *GetCampaignRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CampaignService.MutateCampaigns][google.ads.googleads.v3.services.CampaignService.MutateCampaigns].
type MutateCampaignsRequest struct {
	// Required. The ID of the customer whose campaigns are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual campaigns.
	Operations []*CampaignOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateCampaignsRequest) Reset()         { *m = MutateCampaignsRequest{} }
func (m *MutateCampaignsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignsRequest) ProtoMessage()    {}
func (*MutateCampaignsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6454f925b0d531, []int{1}
}

func (m *MutateCampaignsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignsRequest.Unmarshal(m, b)
}
func (m *MutateCampaignsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignsRequest.Marshal(b, m, deterministic)
}
func (m *MutateCampaignsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignsRequest.Merge(m, src)
}
func (m *MutateCampaignsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignsRequest.Size(m)
}
func (m *MutateCampaignsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignsRequest proto.InternalMessageInfo

func (m *MutateCampaignsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCampaignsRequest) GetOperations() []*CampaignOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCampaignsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCampaignsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a campaign.
type CampaignOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CampaignOperation_Create
	//	*CampaignOperation_Update
	//	*CampaignOperation_Remove
	Operation            isCampaignOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CampaignOperation) Reset()         { *m = CampaignOperation{} }
func (m *CampaignOperation) String() string { return proto.CompactTextString(m) }
func (*CampaignOperation) ProtoMessage()    {}
func (*CampaignOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6454f925b0d531, []int{2}
}

func (m *CampaignOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignOperation.Unmarshal(m, b)
}
func (m *CampaignOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignOperation.Marshal(b, m, deterministic)
}
func (m *CampaignOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignOperation.Merge(m, src)
}
func (m *CampaignOperation) XXX_Size() int {
	return xxx_messageInfo_CampaignOperation.Size(m)
}
func (m *CampaignOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignOperation proto.InternalMessageInfo

func (m *CampaignOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCampaignOperation_Operation interface {
	isCampaignOperation_Operation()
}

type CampaignOperation_Create struct {
	Create *resources.Campaign `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignOperation_Update struct {
	Update *resources.Campaign `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type CampaignOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CampaignOperation_Create) isCampaignOperation_Operation() {}

func (*CampaignOperation_Update) isCampaignOperation_Operation() {}

func (*CampaignOperation_Remove) isCampaignOperation_Operation() {}

func (m *CampaignOperation) GetOperation() isCampaignOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CampaignOperation) GetCreate() *resources.Campaign {
	if x, ok := m.GetOperation().(*CampaignOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CampaignOperation) GetUpdate() *resources.Campaign {
	if x, ok := m.GetOperation().(*CampaignOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *CampaignOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CampaignOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignOperation_Create)(nil),
		(*CampaignOperation_Update)(nil),
		(*CampaignOperation_Remove)(nil),
	}
}

// Response message for campaign mutate.
type MutateCampaignsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCampaignResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MutateCampaignsResponse) Reset()         { *m = MutateCampaignsResponse{} }
func (m *MutateCampaignsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignsResponse) ProtoMessage()    {}
func (*MutateCampaignsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6454f925b0d531, []int{3}
}

func (m *MutateCampaignsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignsResponse.Unmarshal(m, b)
}
func (m *MutateCampaignsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignsResponse.Marshal(b, m, deterministic)
}
func (m *MutateCampaignsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignsResponse.Merge(m, src)
}
func (m *MutateCampaignsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignsResponse.Size(m)
}
func (m *MutateCampaignsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignsResponse proto.InternalMessageInfo

func (m *MutateCampaignsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCampaignsResponse) GetResults() []*MutateCampaignResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the campaign mutate.
type MutateCampaignResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignResult) Reset()         { *m = MutateCampaignResult{} }
func (m *MutateCampaignResult) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignResult) ProtoMessage()    {}
func (*MutateCampaignResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae6454f925b0d531, []int{4}
}

func (m *MutateCampaignResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignResult.Unmarshal(m, b)
}
func (m *MutateCampaignResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignResult.Marshal(b, m, deterministic)
}
func (m *MutateCampaignResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignResult.Merge(m, src)
}
func (m *MutateCampaignResult) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignResult.Size(m)
}
func (m *MutateCampaignResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignResult proto.InternalMessageInfo

func (m *MutateCampaignResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignRequest)(nil), "google.ads.googleads.v3.services.GetCampaignRequest")
	proto.RegisterType((*MutateCampaignsRequest)(nil), "google.ads.googleads.v3.services.MutateCampaignsRequest")
	proto.RegisterType((*CampaignOperation)(nil), "google.ads.googleads.v3.services.CampaignOperation")
	proto.RegisterType((*MutateCampaignsResponse)(nil), "google.ads.googleads.v3.services.MutateCampaignsResponse")
	proto.RegisterType((*MutateCampaignResult)(nil), "google.ads.googleads.v3.services.MutateCampaignResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/campaign_service.proto", fileDescriptor_ae6454f925b0d531)
}

var fileDescriptor_ae6454f925b0d531 = []byte{
	// 750 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4d, 0x6b, 0xd4, 0x4e,
	0x18, 0xff, 0x27, 0xfb, 0xa7, 0xda, 0x49, 0x6b, 0x71, 0xac, 0xed, 0xb2, 0x0a, 0x2e, 0xb1, 0xe0,
	0xb2, 0xd5, 0x44, 0x36, 0xe2, 0x4b, 0x4a, 0x85, 0x6c, 0xe9, 0x8b, 0x87, 0xda, 0x92, 0x42, 0x11,
	0x59, 0x58, 0xa6, 0xc9, 0x74, 0x0d, 0x4d, 0x32, 0x71, 0x66, 0xb2, 0x50, 0x4a, 0x2f, 0x82, 0x9f,
	0xc0, 0x6f, 0xe0, 0xd1, 0xab, 0xf8, 0x09, 0xbc, 0xf5, 0xea, 0xad, 0x27, 0x0f, 0xe2, 0xc1, 0xaf,
	0xe0, 0x45, 0x92, 0xc9, 0xec, 0x5b, 0x5b, 0x96, 0xf6, 0xf6, 0xec, 0x3c, 0xbf, 0xdf, 0xef, 0x79,
	0xcf, 0x82, 0x67, 0x1d, 0x42, 0x3a, 0x21, 0x36, 0x91, 0xcf, 0x4c, 0x61, 0x66, 0x56, 0xd7, 0x32,
	0x19, 0xa6, 0xdd, 0xc0, 0xc3, 0xcc, 0xf4, 0x50, 0x94, 0xa0, 0xa0, 0x13, 0xb7, 0x8b, 0x17, 0x23,
	0xa1, 0x84, 0x13, 0x58, 0x15, 0x68, 0x03, 0xf9, 0xcc, 0xe8, 0x11, 0x8d, 0xae, 0x65, 0x48, 0x62,
	0xe5, 0xf1, 0x45, 0xd2, 0x14, 0x33, 0x92, 0xd2, 0x41, 0x6d, 0xa1, 0x59, 0xb9, 0x2b, 0x19, 0x49,
	0x60, 0xa2, 0x38, 0x26, 0x1c, 0xf1, 0x80, 0xc4, 0xac, 0xf0, 0xce, 0x0f, 0x78, 0xbd, 0x30, 0xc0,
	0x31, 0x2f, 0x1c, 0xf7, 0x06, 0x1c, 0xfb, 0x01, 0x0e, 0xfd, 0xf6, 0x1e, 0x7e, 0x87, 0xba, 0x01,
	0xa1, 0x05, 0xa0, 0xc8, 0xd5, 0xcc, 0x7f, 0xed, 0xa5, 0xfb, 0x05, 0x2a, 0x42, 0xec, 0x60, 0x44,
	0x9b, 0x26, 0x9e, 0xc9, 0x38, 0xe2, 0x69, 0x11, 0x54, 0x7f, 0x09, 0xe0, 0x3a, 0xe6, 0x2b, 0x45,
	0x9e, 0x2e, 0x7e, 0x9f, 0x62, 0xc6, 0x61, 0x0d, 0x4c, 0xcb, 0x22, 0xda, 0x31, 0x8a, 0x70, 0x59,
	0xa9, 0x2a, 0xb5, 0xc9, 0x66, 0xe9, 0xa7, 0xa3, 0xba, 0x53, 0xd2, 0xf3, 0x1a, 0x45, 0x58, 0xff,
	0xad, 0x80, 0xb9, 0xcd, 0x94, 0x23, 0x8e, 0xa5, 0x06, 0x93, 0x22, 0x0b, 0x40, 0xf3, 0x52, 0xc6,
	0x49, 0x84, 0x69, 0x3b, 0xf0, 0x07, 0x25, 0x80, 0x7c, 0x7f, 0xe5, 0xc3, 0x37, 0x00, 0x90, 0x04,
	0x53, 0xd1, 0x89, 0xb2, 0x5a, 0x2d, 0xd5, 0xb4, 0x86, 0x65, 0x8c, 0x6b, 0xbe, 0x21, 0xa3, 0x6d,
	0x49, 0x6e, 0xa1, 0xdc, 0xd7, 0x82, 0x0f, 0xc0, 0x4c, 0x82, 0x28, 0x0f, 0x50, 0xd8, 0xde, 0x47,
	0x41, 0x98, 0x52, 0x5c, 0x2e, 0x55, 0x95, 0xda, 0x75, 0xf7, 0x46, 0xf1, 0xbc, 0x26, 0x5e, 0xe1,
	0x7d, 0x30, 0xdd, 0x45, 0x61, 0xe0, 0x23, 0x8e, 0xdb, 0x24, 0x0e, 0x0f, 0xcb, 0xff, 0xe7, 0xb0,
	0x29, 0xf9, 0xb8, 0x15, 0x87, 0x87, 0xfa, 0x47, 0x15, 0xdc, 0x3c, 0x13, 0x14, 0x2e, 0x01, 0x2d,
	0x4d, 0x72, 0x62, 0xd6, 0xec, 0x9c, 0xa8, 0x35, 0x2a, 0x32, 0x7d, 0x39, 0x0f, 0x63, 0x2d, 0x9b,
	0xc7, 0x26, 0x62, 0x07, 0x2e, 0x10, 0xf0, 0xcc, 0x86, 0xab, 0x60, 0xc2, 0xa3, 0x18, 0x71, 0xd1,
	0x5e, 0xad, 0xb1, 0x78, 0x61, 0xd9, 0xbd, 0x8d, 0xea, 0xd5, 0xbd, 0xf1, 0x9f, 0x5b, 0x90, 0x33,
	0x19, 0x21, 0x5a, 0x56, 0xaf, 0x24, 0x23, 0xc8, 0xb0, 0x0c, 0x26, 0x28, 0x8e, 0x48, 0x57, 0x74,
	0x69, 0x32, 0xf3, 0x88, 0xdf, 0x4d, 0x0d, 0x4c, 0xf6, 0xda, 0xaa, 0x7f, 0x55, 0xc0, 0xfc, 0x99,
	0x81, 0xb3, 0x84, 0xc4, 0x0c, 0xc3, 0x35, 0x70, 0x7b, 0xa4, 0xe3, 0x6d, 0x4c, 0x29, 0xa1, 0xb9,
	0xa2, 0xd6, 0x80, 0x32, 0x31, 0x9a, 0x78, 0xc6, 0x4e, 0xbe, 0x85, 0xee, 0xad, 0xe1, 0x59, 0xac,
	0x66, 0x70, 0xb8, 0x0d, 0xae, 0x51, 0xcc, 0xd2, 0x90, 0xcb, 0x85, 0x78, 0x3a, 0x7e, 0x21, 0x86,
	0x73, 0x72, 0x73, 0xba, 0x2b, 0x65, 0xf4, 0x25, 0x30, 0x7b, 0x1e, 0x20, 0x1b, 0xfd, 0x39, 0x8b,
	0x3e, 0xbc, 0xe3, 0x8d, 0xef, 0x25, 0x30, 0x23, 0x79, 0x3b, 0x22, 0x1e, 0xfc, 0xa6, 0x00, 0x6d,
	0xe0, 0x70, 0xe0, 0x93, 0xf1, 0x19, 0x9e, 0xbd, 0xb3, 0xca, 0x65, 0x46, 0xa5, 0xaf, 0x9c, 0x3a,
	0xc3, 0xc9, 0x7e, 0xf8, 0xf1, 0xeb, 0x93, 0xfa, 0x08, 0x2e, 0x66, 0xdf, 0x9c, 0xa3, 0x21, 0xcf,
	0xb2, 0x3c, 0x30, 0x66, 0xd6, 0x7b, 0x1f, 0x21, 0x66, 0xd6, 0x8f, 0xe1, 0xa9, 0x02, 0x66, 0x46,
	0xc6, 0x07, 0x9f, 0x5f, 0xb6, 0xbb, 0xf2, 0xc4, 0x2b, 0x2f, 0xae, 0xc0, 0x14, 0xbb, 0xa2, 0xbb,
	0xa7, 0xce, 0xdc, 0xc0, 0xe7, 0xe1, 0x61, 0xff, 0x70, 0xf3, 0xb2, 0x2c, 0xdd, 0xc8, 0xca, 0xea,
	0xd7, 0x71, 0x34, 0x00, 0x5e, 0xae, 0x1f, 0xf7, 0xab, 0xb2, 0xa3, 0x3c, 0x82, 0xad, 0xd4, 0x2b,
	0x77, 0x4e, 0x9c, 0x72, 0x3f, 0x8b, 0xc2, 0x4a, 0x02, 0x66, 0x78, 0x24, 0x6a, 0xfe, 0x55, 0xc0,
	0x82, 0x47, 0xa2, 0xb1, 0x19, 0x37, 0x67, 0x47, 0x66, 0xbd, 0x9d, 0x5d, 0xf1, 0xb6, 0xf2, 0x76,
	0xa3, 0x60, 0x76, 0x48, 0x88, 0xe2, 0x8e, 0x41, 0x68, 0xc7, 0xec, 0xe0, 0x38, 0xbf, 0x71, 0xb3,
	0x1f, 0xeb, 0xe2, 0x7f, 0x9a, 0x25, 0x69, 0x7c, 0x56, 0x4b, 0xeb, 0x8e, 0xf3, 0x45, 0xad, 0xae,
	0x0b, 0x41, 0xc7, 0x67, 0x86, 0x30, 0x33, 0x6b, 0xd7, 0x32, 0x8a, 0xc0, 0xec, 0x44, 0x42, 0x5a,
	0x8e, 0xcf, 0x5a, 0x3d, 0x48, 0x6b, 0xd7, 0x6a, 0x49, 0xc8, 0x1f, 0x75, 0x41, 0xbc, 0xdb, 0xb6,
	0xe3, 0x33, 0xdb, 0xee, 0x81, 0x6c, 0x7b, 0xd7, 0xb2, 0x6d, 0x09, 0xdb, 0x9b, 0xc8, 0xf3, 0xb4,
	0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xae, 0xd1, 0x44, 0x10, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CampaignServiceClient is the client API for CampaignService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignServiceClient interface {
	// Returns the requested campaign in full detail.
	GetCampaign(ctx context.Context, in *GetCampaignRequest, opts ...grpc.CallOption) (*resources.Campaign, error)
	// Creates, updates, or removes campaigns. Operation statuses are returned.
	MutateCampaigns(ctx context.Context, in *MutateCampaignsRequest, opts ...grpc.CallOption) (*MutateCampaignsResponse, error)
}

type campaignServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignServiceClient(cc grpc.ClientConnInterface) CampaignServiceClient {
	return &campaignServiceClient{cc}
}

func (c *campaignServiceClient) GetCampaign(ctx context.Context, in *GetCampaignRequest, opts ...grpc.CallOption) (*resources.Campaign, error) {
	out := new(resources.Campaign)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CampaignService/GetCampaign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignServiceClient) MutateCampaigns(ctx context.Context, in *MutateCampaignsRequest, opts ...grpc.CallOption) (*MutateCampaignsResponse, error) {
	out := new(MutateCampaignsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CampaignService/MutateCampaigns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignServiceServer is the server API for CampaignService service.
type CampaignServiceServer interface {
	// Returns the requested campaign in full detail.
	GetCampaign(context.Context, *GetCampaignRequest) (*resources.Campaign, error)
	// Creates, updates, or removes campaigns. Operation statuses are returned.
	MutateCampaigns(context.Context, *MutateCampaignsRequest) (*MutateCampaignsResponse, error)
}

// UnimplementedCampaignServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCampaignServiceServer struct {
}

func (*UnimplementedCampaignServiceServer) GetCampaign(ctx context.Context, req *GetCampaignRequest) (*resources.Campaign, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCampaign not implemented")
}
func (*UnimplementedCampaignServiceServer) MutateCampaigns(ctx context.Context, req *MutateCampaignsRequest) (*MutateCampaignsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCampaigns not implemented")
}

func RegisterCampaignServiceServer(s *grpc.Server, srv CampaignServiceServer) {
	s.RegisterService(&_CampaignService_serviceDesc, srv)
}

func _CampaignService_GetCampaign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignServiceServer).GetCampaign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CampaignService/GetCampaign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignServiceServer).GetCampaign(ctx, req.(*GetCampaignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignService_MutateCampaigns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignServiceServer).MutateCampaigns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CampaignService/MutateCampaigns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignServiceServer).MutateCampaigns(ctx, req.(*MutateCampaignsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.CampaignService",
	HandlerType: (*CampaignServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaign",
			Handler:    _CampaignService_GetCampaign_Handler,
		},
		{
			MethodName: "MutateCampaigns",
			Handler:    _CampaignService_MutateCampaigns_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/campaign_service.proto",
}
