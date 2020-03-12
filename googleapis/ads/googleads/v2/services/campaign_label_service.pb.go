// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/campaign_label_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// Request message for [CampaignLabelService.GetCampaignLabel][google.ads.googleads.v2.services.CampaignLabelService.GetCampaignLabel].
type GetCampaignLabelRequest struct {
	// Required. The resource name of the campaign-label relationship to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignLabelRequest) Reset()         { *m = GetCampaignLabelRequest{} }
func (m *GetCampaignLabelRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignLabelRequest) ProtoMessage()    {}
func (*GetCampaignLabelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c19883c40d100c23, []int{0}
}

func (m *GetCampaignLabelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignLabelRequest.Unmarshal(m, b)
}
func (m *GetCampaignLabelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignLabelRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignLabelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignLabelRequest.Merge(m, src)
}
func (m *GetCampaignLabelRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignLabelRequest.Size(m)
}
func (m *GetCampaignLabelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignLabelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignLabelRequest proto.InternalMessageInfo

func (m *GetCampaignLabelRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CampaignLabelService.MutateCampaignLabels][google.ads.googleads.v2.services.CampaignLabelService.MutateCampaignLabels].
type MutateCampaignLabelsRequest struct {
	// Required. ID of the customer whose campaign-label relationships are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on campaign-label relationships.
	Operations []*CampaignLabelOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateCampaignLabelsRequest) Reset()         { *m = MutateCampaignLabelsRequest{} }
func (m *MutateCampaignLabelsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignLabelsRequest) ProtoMessage()    {}
func (*MutateCampaignLabelsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c19883c40d100c23, []int{1}
}

func (m *MutateCampaignLabelsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignLabelsRequest.Unmarshal(m, b)
}
func (m *MutateCampaignLabelsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignLabelsRequest.Marshal(b, m, deterministic)
}
func (m *MutateCampaignLabelsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignLabelsRequest.Merge(m, src)
}
func (m *MutateCampaignLabelsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignLabelsRequest.Size(m)
}
func (m *MutateCampaignLabelsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignLabelsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignLabelsRequest proto.InternalMessageInfo

func (m *MutateCampaignLabelsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCampaignLabelsRequest) GetOperations() []*CampaignLabelOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCampaignLabelsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCampaignLabelsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on a campaign-label relationship.
type CampaignLabelOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CampaignLabelOperation_Create
	//	*CampaignLabelOperation_Remove
	Operation            isCampaignLabelOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *CampaignLabelOperation) Reset()         { *m = CampaignLabelOperation{} }
func (m *CampaignLabelOperation) String() string { return proto.CompactTextString(m) }
func (*CampaignLabelOperation) ProtoMessage()    {}
func (*CampaignLabelOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_c19883c40d100c23, []int{2}
}

func (m *CampaignLabelOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignLabelOperation.Unmarshal(m, b)
}
func (m *CampaignLabelOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignLabelOperation.Marshal(b, m, deterministic)
}
func (m *CampaignLabelOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignLabelOperation.Merge(m, src)
}
func (m *CampaignLabelOperation) XXX_Size() int {
	return xxx_messageInfo_CampaignLabelOperation.Size(m)
}
func (m *CampaignLabelOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignLabelOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignLabelOperation proto.InternalMessageInfo

type isCampaignLabelOperation_Operation interface {
	isCampaignLabelOperation_Operation()
}

type CampaignLabelOperation_Create struct {
	Create *resources.CampaignLabel `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignLabelOperation_Remove struct {
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*CampaignLabelOperation_Create) isCampaignLabelOperation_Operation() {}

func (*CampaignLabelOperation_Remove) isCampaignLabelOperation_Operation() {}

func (m *CampaignLabelOperation) GetOperation() isCampaignLabelOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CampaignLabelOperation) GetCreate() *resources.CampaignLabel {
	if x, ok := m.GetOperation().(*CampaignLabelOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CampaignLabelOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CampaignLabelOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignLabelOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignLabelOperation_Create)(nil),
		(*CampaignLabelOperation_Remove)(nil),
	}
}

// Response message for a campaign labels mutate.
type MutateCampaignLabelsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCampaignLabelResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *MutateCampaignLabelsResponse) Reset()         { *m = MutateCampaignLabelsResponse{} }
func (m *MutateCampaignLabelsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignLabelsResponse) ProtoMessage()    {}
func (*MutateCampaignLabelsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c19883c40d100c23, []int{3}
}

func (m *MutateCampaignLabelsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignLabelsResponse.Unmarshal(m, b)
}
func (m *MutateCampaignLabelsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignLabelsResponse.Marshal(b, m, deterministic)
}
func (m *MutateCampaignLabelsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignLabelsResponse.Merge(m, src)
}
func (m *MutateCampaignLabelsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignLabelsResponse.Size(m)
}
func (m *MutateCampaignLabelsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignLabelsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignLabelsResponse proto.InternalMessageInfo

func (m *MutateCampaignLabelsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCampaignLabelsResponse) GetResults() []*MutateCampaignLabelResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for a campaign label mutate.
type MutateCampaignLabelResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignLabelResult) Reset()         { *m = MutateCampaignLabelResult{} }
func (m *MutateCampaignLabelResult) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignLabelResult) ProtoMessage()    {}
func (*MutateCampaignLabelResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_c19883c40d100c23, []int{4}
}

func (m *MutateCampaignLabelResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignLabelResult.Unmarshal(m, b)
}
func (m *MutateCampaignLabelResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignLabelResult.Marshal(b, m, deterministic)
}
func (m *MutateCampaignLabelResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignLabelResult.Merge(m, src)
}
func (m *MutateCampaignLabelResult) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignLabelResult.Size(m)
}
func (m *MutateCampaignLabelResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignLabelResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignLabelResult proto.InternalMessageInfo

func (m *MutateCampaignLabelResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignLabelRequest)(nil), "google.ads.googleads.v2.services.GetCampaignLabelRequest")
	proto.RegisterType((*MutateCampaignLabelsRequest)(nil), "google.ads.googleads.v2.services.MutateCampaignLabelsRequest")
	proto.RegisterType((*CampaignLabelOperation)(nil), "google.ads.googleads.v2.services.CampaignLabelOperation")
	proto.RegisterType((*MutateCampaignLabelsResponse)(nil), "google.ads.googleads.v2.services.MutateCampaignLabelsResponse")
	proto.RegisterType((*MutateCampaignLabelResult)(nil), "google.ads.googleads.v2.services.MutateCampaignLabelResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/campaign_label_service.proto", fileDescriptor_c19883c40d100c23)
}

var fileDescriptor_c19883c40d100c23 = []byte{
	// 713 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4d, 0x6b, 0xd4, 0x4e,
	0x18, 0xff, 0x27, 0xfb, 0xa7, 0xda, 0xd9, 0x56, 0x65, 0xac, 0xed, 0x76, 0x5b, 0x70, 0x89, 0x05,
	0x97, 0x45, 0x92, 0x92, 0x82, 0xd4, 0x94, 0x15, 0xb3, 0xc5, 0xb6, 0x8a, 0xda, 0xb2, 0xc5, 0x0a,
	0xb2, 0x12, 0xa6, 0xc9, 0x34, 0x06, 0x92, 0x4c, 0x9c, 0x99, 0x2c, 0x94, 0x52, 0x10, 0xaf, 0x1e,
	0xfd, 0x06, 0x1e, 0xfd, 0x0e, 0x7e, 0x00, 0x7b, 0xf5, 0xd6, 0x93, 0x07, 0x4f, 0x1e, 0xa4, 0x1f,
	0x41, 0xf2, 0x32, 0xbb, 0x9b, 0x75, 0x97, 0xc5, 0xde, 0x9e, 0x3c, 0x2f, 0xbf, 0xe7, 0xf9, 0x3d,
	0x2f, 0x13, 0xd0, 0x74, 0x09, 0x71, 0x7d, 0xac, 0x21, 0x87, 0x69, 0x99, 0x98, 0x48, 0x5d, 0x5d,
	0x63, 0x98, 0x76, 0x3d, 0x1b, 0x33, 0xcd, 0x46, 0x41, 0x84, 0x3c, 0x37, 0xb4, 0x7c, 0x74, 0x88,
	0x7d, 0x2b, 0xd7, 0xab, 0x11, 0x25, 0x9c, 0xc0, 0x5a, 0x16, 0xa3, 0x22, 0x87, 0xa9, 0xbd, 0x70,
	0xb5, 0xab, 0xab, 0x22, 0xbc, 0x7a, 0x7f, 0x5c, 0x02, 0x8a, 0x19, 0x89, 0xe9, 0xdf, 0x19, 0x32,
	0xe4, 0xea, 0xb2, 0x88, 0x8b, 0x3c, 0x0d, 0x85, 0x21, 0xe1, 0x88, 0x7b, 0x24, 0x64, 0xb9, 0x75,
	0x61, 0xc0, 0x6a, 0xfb, 0x1e, 0x0e, 0x79, 0x6e, 0xb8, 0x3d, 0x60, 0x38, 0xf2, 0xb0, 0xef, 0x58,
	0x87, 0xf8, 0x2d, 0xea, 0x7a, 0x84, 0x0e, 0x45, 0xd2, 0xc8, 0xd6, 0x18, 0x47, 0x3c, 0xce, 0x21,
	0x95, 0x4d, 0xb0, 0xb0, 0x8d, 0xf9, 0x66, 0x5e, 0xcb, 0xb3, 0xa4, 0x94, 0x36, 0x7e, 0x17, 0x63,
	0xc6, 0x61, 0x1d, 0xcc, 0x8a, 0x6a, 0xad, 0x10, 0x05, 0xb8, 0x22, 0xd5, 0xa4, 0xfa, 0x74, 0xab,
	0xf4, 0xc3, 0x94, 0xdb, 0x33, 0xc2, 0xf2, 0x02, 0x05, 0x58, 0xb9, 0x90, 0xc0, 0xd2, 0xf3, 0x98,
	0x23, 0x8e, 0x0b, 0x40, 0x4c, 0x20, 0xad, 0x80, 0xb2, 0x1d, 0x33, 0x4e, 0x02, 0x4c, 0x2d, 0xcf,
	0x19, 0xc4, 0x01, 0x42, 0xff, 0xc4, 0x81, 0x6f, 0x00, 0x20, 0x11, 0xa6, 0x19, 0xe3, 0x8a, 0x5c,
	0x2b, 0xd5, 0xcb, 0xfa, 0xba, 0x3a, 0xa9, 0xd5, 0x6a, 0x21, 0xe5, 0xae, 0x00, 0xc8, 0xe1, 0xfb,
	0x80, 0xf0, 0x2e, 0xb8, 0x1e, 0x21, 0xca, 0x3d, 0xe4, 0x5b, 0x47, 0xc8, 0xf3, 0x63, 0x8a, 0x2b,
	0xa5, 0x9a, 0x54, 0xbf, 0xda, 0xbe, 0x96, 0xab, 0xb7, 0x32, 0x2d, 0xbc, 0x03, 0x66, 0xbb, 0xc8,
	0xf7, 0x1c, 0xc4, 0xb1, 0x45, 0x42, 0xff, 0xb8, 0xf2, 0x7f, 0xea, 0x36, 0x23, 0x94, 0xbb, 0xa1,
	0x7f, 0xac, 0x7c, 0x94, 0xc0, 0xfc, 0xe8, 0xcc, 0xf0, 0x29, 0x98, 0xb2, 0x29, 0x46, 0x3c, 0x6b,
	0x58, 0x59, 0x5f, 0x1d, 0xcb, 0xa1, 0xb7, 0x0c, 0x45, 0x12, 0x3b, 0xff, 0xb5, 0x73, 0x04, 0x58,
	0x01, 0x53, 0x14, 0x07, 0xa4, 0x8b, 0x2b, 0x72, 0xd2, 0xb4, 0xc4, 0x92, 0x7d, 0xb7, 0xca, 0x60,
	0xba, 0x47, 0x4e, 0xf9, 0x2a, 0x81, 0xe5, 0xd1, 0x03, 0x60, 0x11, 0x09, 0x19, 0x86, 0x5b, 0xe0,
	0xd6, 0x10, 0x79, 0x0b, 0x53, 0x4a, 0x68, 0xda, 0x82, 0xb2, 0x0e, 0x45, 0x89, 0x34, 0xb2, 0xd5,
	0xfd, 0x74, 0x3f, 0xda, 0x37, 0x8b, 0x6d, 0x79, 0x9c, 0xb8, 0xc3, 0x97, 0xe0, 0x0a, 0xc5, 0x2c,
	0xf6, 0xb9, 0x18, 0xd0, 0xc6, 0xe4, 0x01, 0x8d, 0x28, 0xac, 0x9d, 0x62, 0xb4, 0x05, 0x96, 0xf2,
	0x08, 0x2c, 0x8e, 0xf5, 0x4a, 0xe6, 0x31, 0x62, 0x0f, 0x8b, 0x2b, 0xa8, 0x5f, 0x94, 0xc0, 0x5c,
	0x21, 0x78, 0x3f, 0x4b, 0x0f, 0xbf, 0x49, 0xe0, 0xc6, 0xf0, 0x86, 0xc3, 0x07, 0x93, 0xab, 0x1e,
	0x73, 0x15, 0xd5, 0x7f, 0x9e, 0xa6, 0xb2, 0x73, 0x6e, 0x16, 0x09, 0x7c, 0xf8, 0xfe, 0xf3, 0x93,
	0xac, 0xc3, 0xd5, 0xe4, 0x3d, 0x38, 0x29, 0x58, 0x9a, 0xe2, 0x1c, 0x98, 0xd6, 0xe8, 0x3d, 0x10,
	0xd9, 0x28, 0xb5, 0xc6, 0x29, 0xfc, 0x2d, 0x81, 0xb9, 0x51, 0x63, 0x86, 0xcd, 0x4b, 0x4d, 0x41,
	0xdc, 0x67, 0xf5, 0xe1, 0x65, 0xc3, 0xb3, 0xed, 0x52, 0x5e, 0x9d, 0x9b, 0xf3, 0x03, 0x07, 0x7e,
	0xaf, 0x7f, 0x75, 0x29, 0xd5, 0x75, 0x65, 0x2d, 0xa1, 0xda, 0xe7, 0x76, 0x32, 0xe0, 0xdc, 0x6c,
	0x9c, 0x0e, 0x31, 0x35, 0x82, 0x34, 0x97, 0x21, 0x35, 0xaa, 0x4b, 0x67, 0x66, 0xa5, 0x5f, 0x4f,
	0x2e, 0x45, 0x1e, 0x53, 0x6d, 0x12, 0xb4, 0xde, 0xcb, 0x60, 0xc5, 0x26, 0xc1, 0xc4, 0xda, 0x5b,
	0x8b, 0xa3, 0x16, 0x63, 0x2f, 0x79, 0xfe, 0xf6, 0xa4, 0xd7, 0x3b, 0x79, 0xb8, 0x4b, 0x7c, 0x14,
	0xba, 0x2a, 0xa1, 0xae, 0xe6, 0xe2, 0x30, 0x7d, 0x1c, 0xb5, 0x7e, 0xc2, 0xf1, 0x7f, 0x8a, 0x0d,
	0x21, 0x7c, 0x96, 0x4b, 0xdb, 0xa6, 0xf9, 0x45, 0xae, 0x6d, 0x67, 0x80, 0xa6, 0xc3, 0xd4, 0x4c,
	0x4c, 0xa4, 0x03, 0x5d, 0xcd, 0x13, 0xb3, 0x33, 0xe1, 0xd2, 0x31, 0x1d, 0xd6, 0xe9, 0xb9, 0x74,
	0x0e, 0xf4, 0x8e, 0x70, 0xf9, 0x25, 0xaf, 0x64, 0x7a, 0xc3, 0x30, 0x1d, 0x66, 0x18, 0x3d, 0x27,
	0xc3, 0x38, 0xd0, 0x0d, 0x43, 0xb8, 0x1d, 0x4e, 0xa5, 0x75, 0xae, 0xfd, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x06, 0x7a, 0xde, 0xfd, 0xd0, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CampaignLabelServiceClient is the client API for CampaignLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignLabelServiceClient interface {
	// Returns the requested campaign-label relationship in full detail.
	GetCampaignLabel(ctx context.Context, in *GetCampaignLabelRequest, opts ...grpc.CallOption) (*resources.CampaignLabel, error)
	// Creates and removes campaign-label relationships.
	// Operation statuses are returned.
	MutateCampaignLabels(ctx context.Context, in *MutateCampaignLabelsRequest, opts ...grpc.CallOption) (*MutateCampaignLabelsResponse, error)
}

type campaignLabelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignLabelServiceClient(cc grpc.ClientConnInterface) CampaignLabelServiceClient {
	return &campaignLabelServiceClient{cc}
}

func (c *campaignLabelServiceClient) GetCampaignLabel(ctx context.Context, in *GetCampaignLabelRequest, opts ...grpc.CallOption) (*resources.CampaignLabel, error) {
	out := new(resources.CampaignLabel)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CampaignLabelService/GetCampaignLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignLabelServiceClient) MutateCampaignLabels(ctx context.Context, in *MutateCampaignLabelsRequest, opts ...grpc.CallOption) (*MutateCampaignLabelsResponse, error) {
	out := new(MutateCampaignLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CampaignLabelService/MutateCampaignLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignLabelServiceServer is the server API for CampaignLabelService service.
type CampaignLabelServiceServer interface {
	// Returns the requested campaign-label relationship in full detail.
	GetCampaignLabel(context.Context, *GetCampaignLabelRequest) (*resources.CampaignLabel, error)
	// Creates and removes campaign-label relationships.
	// Operation statuses are returned.
	MutateCampaignLabels(context.Context, *MutateCampaignLabelsRequest) (*MutateCampaignLabelsResponse, error)
}

// UnimplementedCampaignLabelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCampaignLabelServiceServer struct {
}

func (*UnimplementedCampaignLabelServiceServer) GetCampaignLabel(ctx context.Context, req *GetCampaignLabelRequest) (*resources.CampaignLabel, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCampaignLabel not implemented")
}
func (*UnimplementedCampaignLabelServiceServer) MutateCampaignLabels(ctx context.Context, req *MutateCampaignLabelsRequest) (*MutateCampaignLabelsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCampaignLabels not implemented")
}

func RegisterCampaignLabelServiceServer(s *grpc.Server, srv CampaignLabelServiceServer) {
	s.RegisterService(&_CampaignLabelService_serviceDesc, srv)
}

func _CampaignLabelService_GetCampaignLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignLabelServiceServer).GetCampaignLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CampaignLabelService/GetCampaignLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignLabelServiceServer).GetCampaignLabel(ctx, req.(*GetCampaignLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignLabelService_MutateCampaignLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignLabelServiceServer).MutateCampaignLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CampaignLabelService/MutateCampaignLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignLabelServiceServer).MutateCampaignLabels(ctx, req.(*MutateCampaignLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignLabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.CampaignLabelService",
	HandlerType: (*CampaignLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignLabel",
			Handler:    _CampaignLabelService_GetCampaignLabel_Handler,
		},
		{
			MethodName: "MutateCampaignLabels",
			Handler:    _CampaignLabelService_MutateCampaignLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/campaign_label_service.proto",
}
