// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/ad_group_ad_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
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

// Request message for [AdGroupAdService.GetAdGroupAd][google.ads.googleads.v1.services.AdGroupAdService.GetAdGroupAd].
type GetAdGroupAdRequest struct {
	// Required. The resource name of the ad to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupAdRequest) Reset()         { *m = GetAdGroupAdRequest{} }
func (m *GetAdGroupAdRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupAdRequest) ProtoMessage()    {}
func (*GetAdGroupAdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3faf2e58edc49841, []int{0}
}

func (m *GetAdGroupAdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupAdRequest.Unmarshal(m, b)
}
func (m *GetAdGroupAdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupAdRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupAdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupAdRequest.Merge(m, src)
}
func (m *GetAdGroupAdRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupAdRequest.Size(m)
}
func (m *GetAdGroupAdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupAdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupAdRequest proto.InternalMessageInfo

func (m *GetAdGroupAdRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [AdGroupAdService.MutateAdGroupAds][google.ads.googleads.v1.services.AdGroupAdService.MutateAdGroupAds].
type MutateAdGroupAdsRequest struct {
	// Required. The ID of the customer whose ads are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual ads.
	Operations []*AdGroupAdOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateAdGroupAdsRequest) Reset()         { *m = MutateAdGroupAdsRequest{} }
func (m *MutateAdGroupAdsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupAdsRequest) ProtoMessage()    {}
func (*MutateAdGroupAdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3faf2e58edc49841, []int{1}
}

func (m *MutateAdGroupAdsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupAdsRequest.Unmarshal(m, b)
}
func (m *MutateAdGroupAdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupAdsRequest.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupAdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupAdsRequest.Merge(m, src)
}
func (m *MutateAdGroupAdsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupAdsRequest.Size(m)
}
func (m *MutateAdGroupAdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupAdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupAdsRequest proto.InternalMessageInfo

func (m *MutateAdGroupAdsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdGroupAdsRequest) GetOperations() []*AdGroupAdOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateAdGroupAdsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateAdGroupAdsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on an ad group ad.
type AdGroupAdOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// Configuration for how policies are validated.
	PolicyValidationParameter *common.PolicyValidationParameter `protobuf:"bytes,5,opt,name=policy_validation_parameter,json=policyValidationParameter,proto3" json:"policy_validation_parameter,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdGroupAdOperation_Create
	//	*AdGroupAdOperation_Update
	//	*AdGroupAdOperation_Remove
	Operation            isAdGroupAdOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *AdGroupAdOperation) Reset()         { *m = AdGroupAdOperation{} }
func (m *AdGroupAdOperation) String() string { return proto.CompactTextString(m) }
func (*AdGroupAdOperation) ProtoMessage()    {}
func (*AdGroupAdOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_3faf2e58edc49841, []int{2}
}

func (m *AdGroupAdOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAdOperation.Unmarshal(m, b)
}
func (m *AdGroupAdOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAdOperation.Marshal(b, m, deterministic)
}
func (m *AdGroupAdOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAdOperation.Merge(m, src)
}
func (m *AdGroupAdOperation) XXX_Size() int {
	return xxx_messageInfo_AdGroupAdOperation.Size(m)
}
func (m *AdGroupAdOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAdOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAdOperation proto.InternalMessageInfo

func (m *AdGroupAdOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func (m *AdGroupAdOperation) GetPolicyValidationParameter() *common.PolicyValidationParameter {
	if m != nil {
		return m.PolicyValidationParameter
	}
	return nil
}

type isAdGroupAdOperation_Operation interface {
	isAdGroupAdOperation_Operation()
}

type AdGroupAdOperation_Create struct {
	Create *resources.AdGroupAd `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupAdOperation_Update struct {
	Update *resources.AdGroupAd `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type AdGroupAdOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*AdGroupAdOperation_Create) isAdGroupAdOperation_Operation() {}

func (*AdGroupAdOperation_Update) isAdGroupAdOperation_Operation() {}

func (*AdGroupAdOperation_Remove) isAdGroupAdOperation_Operation() {}

func (m *AdGroupAdOperation) GetOperation() isAdGroupAdOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdGroupAdOperation) GetCreate() *resources.AdGroupAd {
	if x, ok := m.GetOperation().(*AdGroupAdOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdGroupAdOperation) GetUpdate() *resources.AdGroupAd {
	if x, ok := m.GetOperation().(*AdGroupAdOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *AdGroupAdOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdGroupAdOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupAdOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupAdOperation_Create)(nil),
		(*AdGroupAdOperation_Update)(nil),
		(*AdGroupAdOperation_Remove)(nil),
	}
}

// Response message for an ad group ad mutate.
type MutateAdGroupAdsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateAdGroupAdResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MutateAdGroupAdsResponse) Reset()         { *m = MutateAdGroupAdsResponse{} }
func (m *MutateAdGroupAdsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupAdsResponse) ProtoMessage()    {}
func (*MutateAdGroupAdsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3faf2e58edc49841, []int{3}
}

func (m *MutateAdGroupAdsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupAdsResponse.Unmarshal(m, b)
}
func (m *MutateAdGroupAdsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupAdsResponse.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupAdsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupAdsResponse.Merge(m, src)
}
func (m *MutateAdGroupAdsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupAdsResponse.Size(m)
}
func (m *MutateAdGroupAdsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupAdsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupAdsResponse proto.InternalMessageInfo

func (m *MutateAdGroupAdsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateAdGroupAdsResponse) GetResults() []*MutateAdGroupAdResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the ad mutate.
type MutateAdGroupAdResult struct {
	// The resource name returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupAdResult) Reset()         { *m = MutateAdGroupAdResult{} }
func (m *MutateAdGroupAdResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupAdResult) ProtoMessage()    {}
func (*MutateAdGroupAdResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_3faf2e58edc49841, []int{4}
}

func (m *MutateAdGroupAdResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupAdResult.Unmarshal(m, b)
}
func (m *MutateAdGroupAdResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupAdResult.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupAdResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupAdResult.Merge(m, src)
}
func (m *MutateAdGroupAdResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupAdResult.Size(m)
}
func (m *MutateAdGroupAdResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupAdResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupAdResult proto.InternalMessageInfo

func (m *MutateAdGroupAdResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupAdRequest)(nil), "google.ads.googleads.v1.services.GetAdGroupAdRequest")
	proto.RegisterType((*MutateAdGroupAdsRequest)(nil), "google.ads.googleads.v1.services.MutateAdGroupAdsRequest")
	proto.RegisterType((*AdGroupAdOperation)(nil), "google.ads.googleads.v1.services.AdGroupAdOperation")
	proto.RegisterType((*MutateAdGroupAdsResponse)(nil), "google.ads.googleads.v1.services.MutateAdGroupAdsResponse")
	proto.RegisterType((*MutateAdGroupAdResult)(nil), "google.ads.googleads.v1.services.MutateAdGroupAdResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/ad_group_ad_service.proto", fileDescriptor_3faf2e58edc49841)
}

var fileDescriptor_3faf2e58edc49841 = []byte{
	// 818 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x6b, 0xe3, 0x46,
	0x14, 0xaf, 0xec, 0x36, 0x6d, 0xc6, 0x49, 0x1b, 0x26, 0xa4, 0x51, 0x9d, 0x42, 0x8d, 0x1a, 0xa8,
	0x71, 0x83, 0x84, 0x9d, 0x94, 0x12, 0xa5, 0xa1, 0xc8, 0xb4, 0x71, 0x7a, 0x48, 0xe3, 0x2a, 0x60,
	0x68, 0x31, 0x88, 0x89, 0x34, 0x71, 0x45, 0x24, 0x8d, 0x3a, 0x33, 0x32, 0x98, 0x90, 0x4b, 0xfb,
	0x11, 0xfa, 0x0d, 0xf6, 0xb8, 0xf7, 0x3d, 0xec, 0x57, 0x08, 0x7b, 0xdb, 0x5b, 0x0e, 0xcb, 0x1e,
	0xf6, 0xb0, 0xec, 0x67, 0xd8, 0xc3, 0x22, 0x8d, 0x46, 0xfe, 0x93, 0x18, 0x93, 0xdc, 0x9e, 0xe6,
	0xfd, 0x7e, 0xbf, 0xf7, 0xde, 0xbc, 0xf7, 0x46, 0xc0, 0x1c, 0x10, 0x32, 0x08, 0xb0, 0x81, 0x3c,
	0x66, 0x08, 0x33, 0xb5, 0x86, 0x4d, 0x83, 0x61, 0x3a, 0xf4, 0x5d, 0xcc, 0x0c, 0xe4, 0x39, 0x03,
	0x4a, 0x92, 0xd8, 0x41, 0x9e, 0x93, 0x1f, 0xea, 0x31, 0x25, 0x9c, 0xc0, 0x9a, 0x20, 0xe8, 0xc8,
	0x63, 0x7a, 0xc1, 0xd5, 0x87, 0x4d, 0x5d, 0x72, 0xab, 0xdf, 0xcf, 0x53, 0x77, 0x49, 0x18, 0x92,
	0xc8, 0x88, 0x49, 0xe0, 0xbb, 0x23, 0x21, 0x57, 0xdd, 0x9d, 0x07, 0xa6, 0x98, 0x91, 0x84, 0xce,
	0xe4, 0x92, 0x93, 0xbe, 0x96, 0xa4, 0xd8, 0x37, 0x50, 0x14, 0x11, 0x8e, 0xb8, 0x4f, 0x22, 0x96,
	0x7b, 0x37, 0x27, 0xbc, 0x6e, 0xe0, 0xe3, 0x88, 0xe7, 0x8e, 0x6f, 0x26, 0x1c, 0x17, 0x3e, 0x0e,
	0x3c, 0xe7, 0x1c, 0xff, 0x8d, 0x86, 0x3e, 0xa1, 0x39, 0x20, 0xaf, 0xcd, 0xc8, 0xbe, 0xce, 0x93,
	0x8b, 0x1c, 0x15, 0x22, 0x76, 0x39, 0xa3, 0x4d, 0x63, 0xd7, 0x60, 0x1c, 0xf1, 0x24, 0x0f, 0xaa,
	0xfd, 0x0c, 0xd6, 0x3b, 0x98, 0x5b, 0x5e, 0x27, 0xcd, 0xd4, 0xf2, 0x6c, 0xfc, 0x4f, 0x82, 0x19,
	0x87, 0x75, 0xb0, 0x2a, 0x0b, 0x71, 0x22, 0x14, 0x62, 0x55, 0xa9, 0x29, 0xf5, 0xe5, 0x76, 0xf9,
	0xb5, 0x55, 0xb2, 0x57, 0xa4, 0xe7, 0x77, 0x14, 0x62, 0xed, 0xad, 0x02, 0x36, 0x4f, 0x12, 0x8e,
	0x38, 0x2e, 0x44, 0x98, 0x54, 0xd9, 0x06, 0x15, 0x37, 0x61, 0x9c, 0x84, 0x98, 0x3a, 0xbe, 0x37,
	0xa9, 0x01, 0xe4, 0xf9, 0x6f, 0x1e, 0xfc, 0x13, 0x00, 0x12, 0x63, 0x2a, 0xee, 0x42, 0x2d, 0xd5,
	0xca, 0xf5, 0x4a, 0x6b, 0x4f, 0x5f, 0xd4, 0x2e, 0xbd, 0x08, 0x77, 0x2a, 0xc9, 0xb9, 0xf4, 0x58,
	0x0c, 0x7e, 0x07, 0xbe, 0x88, 0x11, 0xe5, 0x3e, 0x0a, 0x9c, 0x0b, 0xe4, 0x07, 0x09, 0xc5, 0x6a,
	0xb9, 0xa6, 0xd4, 0x3f, 0xb3, 0x3f, 0xcf, 0x8f, 0x8f, 0xc4, 0x29, 0xfc, 0x16, 0xac, 0x0e, 0x51,
	0xe0, 0x7b, 0x88, 0x63, 0x87, 0x44, 0xc1, 0x48, 0xfd, 0x38, 0x83, 0xad, 0xc8, 0xc3, 0xd3, 0x28,
	0x18, 0x69, 0xff, 0x95, 0x01, 0xbc, 0x1b, 0x15, 0x1e, 0x80, 0x4a, 0x12, 0x67, 0xcc, 0xf4, 0xc2,
	0x33, 0x66, 0xa5, 0x55, 0x95, 0x05, 0xc8, 0x9e, 0xe8, 0x47, 0x69, 0x4f, 0x4e, 0x10, 0xbb, 0xb4,
	0x81, 0x80, 0xa7, 0x36, 0x1c, 0x81, 0x2d, 0x31, 0x57, 0x4e, 0x1e, 0xca, 0x27, 0x91, 0x13, 0x23,
	0x8a, 0x42, 0xcc, 0x31, 0x55, 0x3f, 0xc9, 0xc4, 0xf6, 0xe7, 0xde, 0x86, 0x18, 0x4d, 0xbd, 0x9b,
	0x49, 0xf4, 0x0a, 0x85, 0xae, 0x14, 0xb0, 0xbf, 0x8a, 0xe7, 0xb9, 0xe0, 0x11, 0x58, 0x72, 0x29,
	0x46, 0x5c, 0x34, 0xb7, 0xd2, 0xda, 0x99, 0x1b, 0xa5, 0x98, 0xe9, 0xf1, 0xa5, 0x1f, 0x7f, 0x64,
	0xe7, 0xec, 0x54, 0x47, 0x14, 0xa4, 0x96, 0x1e, 0xa7, 0x23, 0xd8, 0x50, 0x05, 0x4b, 0x14, 0x87,
	0x64, 0x28, 0x7a, 0xb4, 0x9c, 0x7a, 0xc4, 0x77, 0xbb, 0x02, 0x96, 0x8b, 0xa6, 0x6a, 0xcf, 0x14,
	0xa0, 0xde, 0x1d, 0x38, 0x16, 0x93, 0x88, 0xa5, 0xb9, 0x6c, 0xcc, 0x34, 0xdc, 0xc1, 0x94, 0x12,
	0x9a, 0x49, 0x56, 0x5a, 0x50, 0xa6, 0x46, 0x63, 0x57, 0x3f, 0xcb, 0xf6, 0xc0, 0x5e, 0x9f, 0x1e,
	0x85, 0x5f, 0x53, 0x38, 0xfc, 0x03, 0x7c, 0x4a, 0x31, 0x4b, 0x02, 0x2e, 0x07, 0xf2, 0xc7, 0xc5,
	0x03, 0x39, 0x93, 0x94, 0x9d, 0xf1, 0x6d, 0xa9, 0xa3, 0xfd, 0x04, 0x36, 0xee, 0x45, 0xa4, 0xb3,
	0x77, 0xcf, 0xae, 0x4d, 0xaf, 0x59, 0xeb, 0x45, 0x19, 0xac, 0x15, 0xc4, 0x33, 0x11, 0x12, 0x3e,
	0x57, 0xc0, 0xca, 0xe4, 0xf6, 0xc2, 0x1f, 0x16, 0x67, 0x79, 0xcf, 0xb6, 0x57, 0x1f, 0xd4, 0x31,
	0xed, 0x97, 0x5b, 0x6b, 0x3a, 0xe1, 0x7f, 0x5f, 0xbe, 0xf9, 0xbf, 0xa4, 0xc3, 0x9d, 0xf4, 0xf9,
	0xbb, 0x9a, 0xf2, 0x1c, 0xca, 0x35, 0x67, 0x46, 0xc3, 0x40, 0x45, 0xbb, 0x8c, 0xc6, 0x35, 0x7c,
	0xa5, 0x80, 0xb5, 0xd9, 0x36, 0xc2, 0xfd, 0x07, 0xdf, 0xb2, 0x7c, 0x6b, 0xaa, 0xe6, 0x63, 0xa8,
	0x62, 0x6a, 0xb4, 0xb3, 0x5b, 0xeb, 0xcb, 0x89, 0x87, 0x6a, 0x67, 0xfc, 0x82, 0x64, 0xa5, 0xed,
	0x69, 0x46, 0xf6, 0x1b, 0x28, 0x6a, 0xb9, 0x9a, 0x00, 0x1f, 0x36, 0xae, 0x27, 0x2a, 0x33, 0xc3,
	0x2c, 0x86, 0xa9, 0x34, 0xaa, 0x5b, 0x37, 0x96, 0x3a, 0xce, 0x23, 0xb7, 0x62, 0x9f, 0xa5, 0x3b,
	0xdb, 0x7e, 0xaf, 0x80, 0x6d, 0x97, 0x84, 0x0b, 0x73, 0x6e, 0x6f, 0xcc, 0x36, 0xbd, 0x9b, 0xbe,
	0x27, 0x5d, 0xe5, 0xaf, 0xe3, 0x9c, 0x3a, 0x20, 0x01, 0x8a, 0x06, 0x3a, 0xa1, 0x03, 0x63, 0x80,
	0xa3, 0xec, 0xb5, 0x31, 0xc6, 0xc1, 0xe6, 0xff, 0x2a, 0x0f, 0xa4, 0xf1, 0xa4, 0x54, 0xee, 0x58,
	0xd6, 0xd3, 0x52, 0xad, 0x23, 0x04, 0x2d, 0x8f, 0xe9, 0xc2, 0x4c, 0xad, 0x5e, 0x53, 0xcf, 0x03,
	0xb3, 0x1b, 0x09, 0xe9, 0x5b, 0x1e, 0xeb, 0x17, 0x90, 0x7e, 0xaf, 0xd9, 0x97, 0x90, 0x77, 0xa5,
	0x6d, 0x71, 0x6e, 0x9a, 0xe9, 0x65, 0x98, 0x05, 0xc8, 0x34, 0x7b, 0x4d, 0xd3, 0x94, 0xb0, 0xf3,
	0xa5, 0x2c, 0xcf, 0xdd, 0x0f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xf5, 0x92, 0xb8, 0xd1, 0x07,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdGroupAdServiceClient is the client API for AdGroupAdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupAdServiceClient interface {
	// Returns the requested ad in full detail.
	GetAdGroupAd(ctx context.Context, in *GetAdGroupAdRequest, opts ...grpc.CallOption) (*resources.AdGroupAd, error)
	// Creates, updates, or removes ads. Operation statuses are returned.
	MutateAdGroupAds(ctx context.Context, in *MutateAdGroupAdsRequest, opts ...grpc.CallOption) (*MutateAdGroupAdsResponse, error)
}

type adGroupAdServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupAdServiceClient(cc grpc.ClientConnInterface) AdGroupAdServiceClient {
	return &adGroupAdServiceClient{cc}
}

func (c *adGroupAdServiceClient) GetAdGroupAd(ctx context.Context, in *GetAdGroupAdRequest, opts ...grpc.CallOption) (*resources.AdGroupAd, error) {
	out := new(resources.AdGroupAd)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AdGroupAdService/GetAdGroupAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupAdServiceClient) MutateAdGroupAds(ctx context.Context, in *MutateAdGroupAdsRequest, opts ...grpc.CallOption) (*MutateAdGroupAdsResponse, error) {
	out := new(MutateAdGroupAdsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AdGroupAdService/MutateAdGroupAds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupAdServiceServer is the server API for AdGroupAdService service.
type AdGroupAdServiceServer interface {
	// Returns the requested ad in full detail.
	GetAdGroupAd(context.Context, *GetAdGroupAdRequest) (*resources.AdGroupAd, error)
	// Creates, updates, or removes ads. Operation statuses are returned.
	MutateAdGroupAds(context.Context, *MutateAdGroupAdsRequest) (*MutateAdGroupAdsResponse, error)
}

// UnimplementedAdGroupAdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupAdServiceServer struct {
}

func (*UnimplementedAdGroupAdServiceServer) GetAdGroupAd(ctx context.Context, req *GetAdGroupAdRequest) (*resources.AdGroupAd, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetAdGroupAd not implemented")
}
func (*UnimplementedAdGroupAdServiceServer) MutateAdGroupAds(ctx context.Context, req *MutateAdGroupAdsRequest) (*MutateAdGroupAdsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateAdGroupAds not implemented")
}

func RegisterAdGroupAdServiceServer(s *grpc.Server, srv AdGroupAdServiceServer) {
	s.RegisterService(&_AdGroupAdService_serviceDesc, srv)
}

func _AdGroupAdService_GetAdGroupAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAdServiceServer).GetAdGroupAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AdGroupAdService/GetAdGroupAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAdServiceServer).GetAdGroupAd(ctx, req.(*GetAdGroupAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupAdService_MutateAdGroupAds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupAdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAdServiceServer).MutateAdGroupAds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AdGroupAdService/MutateAdGroupAds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAdServiceServer).MutateAdGroupAds(ctx, req.(*MutateAdGroupAdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupAdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.AdGroupAdService",
	HandlerType: (*AdGroupAdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupAd",
			Handler:    _AdGroupAdService_GetAdGroupAd_Handler,
		},
		{
			MethodName: "MutateAdGroupAds",
			Handler:    _AdGroupAdService_MutateAdGroupAds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/ad_group_ad_service.proto",
}
