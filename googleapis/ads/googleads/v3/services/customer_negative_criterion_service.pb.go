// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/customer_negative_criterion_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
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

// Request message for
// [CustomerNegativeCriterionService.GetCustomerNegativeCriterion][google.ads.googleads.v3.services.CustomerNegativeCriterionService.GetCustomerNegativeCriterion].
type GetCustomerNegativeCriterionRequest struct {
	// Required. The resource name of the criterion to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerNegativeCriterionRequest) Reset()         { *m = GetCustomerNegativeCriterionRequest{} }
func (m *GetCustomerNegativeCriterionRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerNegativeCriterionRequest) ProtoMessage()    {}
func (*GetCustomerNegativeCriterionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_80c7f57f4d9bdab2, []int{0}
}

func (m *GetCustomerNegativeCriterionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerNegativeCriterionRequest.Unmarshal(m, b)
}
func (m *GetCustomerNegativeCriterionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerNegativeCriterionRequest.Marshal(b, m, deterministic)
}
func (m *GetCustomerNegativeCriterionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerNegativeCriterionRequest.Merge(m, src)
}
func (m *GetCustomerNegativeCriterionRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerNegativeCriterionRequest.Size(m)
}
func (m *GetCustomerNegativeCriterionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerNegativeCriterionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerNegativeCriterionRequest proto.InternalMessageInfo

func (m *GetCustomerNegativeCriterionRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [CustomerNegativeCriterionService.MutateCustomerNegativeCriteria][google.ads.googleads.v3.services.CustomerNegativeCriterionService.MutateCustomerNegativeCriteria].
type MutateCustomerNegativeCriteriaRequest struct {
	// Required. The ID of the customer whose criteria are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual criteria.
	Operations []*CustomerNegativeCriterionOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateCustomerNegativeCriteriaRequest) Reset()         { *m = MutateCustomerNegativeCriteriaRequest{} }
func (m *MutateCustomerNegativeCriteriaRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerNegativeCriteriaRequest) ProtoMessage()    {}
func (*MutateCustomerNegativeCriteriaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_80c7f57f4d9bdab2, []int{1}
}

func (m *MutateCustomerNegativeCriteriaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerNegativeCriteriaRequest.Unmarshal(m, b)
}
func (m *MutateCustomerNegativeCriteriaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerNegativeCriteriaRequest.Marshal(b, m, deterministic)
}
func (m *MutateCustomerNegativeCriteriaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerNegativeCriteriaRequest.Merge(m, src)
}
func (m *MutateCustomerNegativeCriteriaRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerNegativeCriteriaRequest.Size(m)
}
func (m *MutateCustomerNegativeCriteriaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerNegativeCriteriaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerNegativeCriteriaRequest proto.InternalMessageInfo

func (m *MutateCustomerNegativeCriteriaRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCustomerNegativeCriteriaRequest) GetOperations() []*CustomerNegativeCriterionOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCustomerNegativeCriteriaRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCustomerNegativeCriteriaRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create or remove) on a customer level negative criterion.
type CustomerNegativeCriterionOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CustomerNegativeCriterionOperation_Create
	//	*CustomerNegativeCriterionOperation_Remove
	Operation            isCustomerNegativeCriterionOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *CustomerNegativeCriterionOperation) Reset()         { *m = CustomerNegativeCriterionOperation{} }
func (m *CustomerNegativeCriterionOperation) String() string { return proto.CompactTextString(m) }
func (*CustomerNegativeCriterionOperation) ProtoMessage()    {}
func (*CustomerNegativeCriterionOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_80c7f57f4d9bdab2, []int{2}
}

func (m *CustomerNegativeCriterionOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerNegativeCriterionOperation.Unmarshal(m, b)
}
func (m *CustomerNegativeCriterionOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerNegativeCriterionOperation.Marshal(b, m, deterministic)
}
func (m *CustomerNegativeCriterionOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerNegativeCriterionOperation.Merge(m, src)
}
func (m *CustomerNegativeCriterionOperation) XXX_Size() int {
	return xxx_messageInfo_CustomerNegativeCriterionOperation.Size(m)
}
func (m *CustomerNegativeCriterionOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerNegativeCriterionOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerNegativeCriterionOperation proto.InternalMessageInfo

type isCustomerNegativeCriterionOperation_Operation interface {
	isCustomerNegativeCriterionOperation_Operation()
}

type CustomerNegativeCriterionOperation_Create struct {
	Create *resources.CustomerNegativeCriterion `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CustomerNegativeCriterionOperation_Remove struct {
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*CustomerNegativeCriterionOperation_Create) isCustomerNegativeCriterionOperation_Operation() {}

func (*CustomerNegativeCriterionOperation_Remove) isCustomerNegativeCriterionOperation_Operation() {}

func (m *CustomerNegativeCriterionOperation) GetOperation() isCustomerNegativeCriterionOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CustomerNegativeCriterionOperation) GetCreate() *resources.CustomerNegativeCriterion {
	if x, ok := m.GetOperation().(*CustomerNegativeCriterionOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CustomerNegativeCriterionOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CustomerNegativeCriterionOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CustomerNegativeCriterionOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CustomerNegativeCriterionOperation_Create)(nil),
		(*CustomerNegativeCriterionOperation_Remove)(nil),
	}
}

// Response message for customer negative criterion mutate.
type MutateCustomerNegativeCriteriaResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCustomerNegativeCriteriaResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *MutateCustomerNegativeCriteriaResponse) Reset() {
	*m = MutateCustomerNegativeCriteriaResponse{}
}
func (m *MutateCustomerNegativeCriteriaResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerNegativeCriteriaResponse) ProtoMessage()    {}
func (*MutateCustomerNegativeCriteriaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_80c7f57f4d9bdab2, []int{3}
}

func (m *MutateCustomerNegativeCriteriaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerNegativeCriteriaResponse.Unmarshal(m, b)
}
func (m *MutateCustomerNegativeCriteriaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerNegativeCriteriaResponse.Marshal(b, m, deterministic)
}
func (m *MutateCustomerNegativeCriteriaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerNegativeCriteriaResponse.Merge(m, src)
}
func (m *MutateCustomerNegativeCriteriaResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerNegativeCriteriaResponse.Size(m)
}
func (m *MutateCustomerNegativeCriteriaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerNegativeCriteriaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerNegativeCriteriaResponse proto.InternalMessageInfo

func (m *MutateCustomerNegativeCriteriaResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCustomerNegativeCriteriaResponse) GetResults() []*MutateCustomerNegativeCriteriaResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the criterion mutate.
type MutateCustomerNegativeCriteriaResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCustomerNegativeCriteriaResult) Reset()         { *m = MutateCustomerNegativeCriteriaResult{} }
func (m *MutateCustomerNegativeCriteriaResult) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerNegativeCriteriaResult) ProtoMessage()    {}
func (*MutateCustomerNegativeCriteriaResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_80c7f57f4d9bdab2, []int{4}
}

func (m *MutateCustomerNegativeCriteriaResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerNegativeCriteriaResult.Unmarshal(m, b)
}
func (m *MutateCustomerNegativeCriteriaResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerNegativeCriteriaResult.Marshal(b, m, deterministic)
}
func (m *MutateCustomerNegativeCriteriaResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerNegativeCriteriaResult.Merge(m, src)
}
func (m *MutateCustomerNegativeCriteriaResult) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerNegativeCriteriaResult.Size(m)
}
func (m *MutateCustomerNegativeCriteriaResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerNegativeCriteriaResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerNegativeCriteriaResult proto.InternalMessageInfo

func (m *MutateCustomerNegativeCriteriaResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCustomerNegativeCriterionRequest)(nil), "google.ads.googleads.v3.services.GetCustomerNegativeCriterionRequest")
	proto.RegisterType((*MutateCustomerNegativeCriteriaRequest)(nil), "google.ads.googleads.v3.services.MutateCustomerNegativeCriteriaRequest")
	proto.RegisterType((*CustomerNegativeCriterionOperation)(nil), "google.ads.googleads.v3.services.CustomerNegativeCriterionOperation")
	proto.RegisterType((*MutateCustomerNegativeCriteriaResponse)(nil), "google.ads.googleads.v3.services.MutateCustomerNegativeCriteriaResponse")
	proto.RegisterType((*MutateCustomerNegativeCriteriaResult)(nil), "google.ads.googleads.v3.services.MutateCustomerNegativeCriteriaResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/customer_negative_criterion_service.proto", fileDescriptor_80c7f57f4d9bdab2)
}

var fileDescriptor_80c7f57f4d9bdab2 = []byte{
	// 720 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xfe, 0xd9, 0xa9, 0xfa, 0xa3, 0x9b, 0x16, 0xa4, 0x45, 0x40, 0x14, 0x2a, 0x88, 0xdc, 0x14,
	0xa2, 0x08, 0xd9, 0x52, 0x72, 0x33, 0xed, 0xc1, 0x29, 0x6d, 0x0a, 0x88, 0xb6, 0x4a, 0xa5, 0x1c,
	0x50, 0x24, 0xb3, 0xb5, 0xb7, 0x61, 0x25, 0xc7, 0x6b, 0x76, 0xd7, 0x91, 0xaa, 0xaa, 0x17, 0x0e,
	0x9c, 0xb8, 0x21, 0xf1, 0x00, 0x70, 0xe3, 0x51, 0x7a, 0x2c, 0xb7, 0x9e, 0x38, 0x70, 0xe2, 0x05,
	0xb8, 0x22, 0xff, 0x59, 0x27, 0xa9, 0x70, 0x5d, 0xa9, 0xdc, 0x26, 0x33, 0xe3, 0x6f, 0xe6, 0xfb,
	0x66, 0x76, 0x02, 0x5e, 0x0c, 0x29, 0x1d, 0x7a, 0xd8, 0x40, 0x2e, 0x37, 0x12, 0x33, 0xb2, 0xc6,
	0x6d, 0x83, 0x63, 0x36, 0x26, 0x0e, 0xe6, 0x86, 0x13, 0x72, 0x41, 0x47, 0x98, 0xd9, 0x3e, 0x1e,
	0x22, 0x41, 0xc6, 0xd8, 0x76, 0x18, 0x11, 0x98, 0x11, 0xea, 0xdb, 0x69, 0x92, 0x1e, 0x30, 0x2a,
	0x28, 0xac, 0x25, 0x00, 0x3a, 0x72, 0xb9, 0x9e, 0x61, 0xe9, 0xe3, 0xb6, 0x2e, 0xb1, 0xaa, 0x1b,
	0x79, 0xd5, 0x18, 0xe6, 0x34, 0x64, 0x05, 0xe5, 0x92, 0x32, 0xd5, 0x65, 0x09, 0x12, 0x10, 0x03,
	0xf9, 0x3e, 0x15, 0x48, 0x10, 0xea, 0xf3, 0x34, 0x7a, 0x6f, 0x2a, 0xea, 0x78, 0x04, 0xfb, 0x22,
	0x0d, 0x3c, 0x9c, 0x0a, 0x1c, 0x12, 0xec, 0xb9, 0xf6, 0x01, 0x7e, 0x8b, 0xc6, 0x84, 0xb2, 0x0b,
	0x5f, 0xb2, 0xc0, 0x31, 0xb8, 0x40, 0x22, 0x4c, 0x21, 0xb5, 0x5d, 0xb0, 0xd2, 0xc5, 0x62, 0x23,
	0x6d, 0x6c, 0x27, 0xed, 0x6b, 0x43, 0xb6, 0xd5, 0xc3, 0xef, 0x42, 0xcc, 0x05, 0x6c, 0x80, 0x25,
	0x49, 0xc3, 0xf6, 0xd1, 0x08, 0x57, 0x94, 0x9a, 0xd2, 0x58, 0xe8, 0x94, 0x7e, 0x58, 0x6a, 0x6f,
	0x51, 0x46, 0x76, 0xd0, 0x08, 0x6b, 0x1f, 0x54, 0xb0, 0xfa, 0x2a, 0x14, 0x48, 0xe0, 0x1c, 0x50,
	0x24, 0x31, 0xeb, 0xa0, 0x9c, 0x09, 0x42, 0xdc, 0x69, 0x44, 0x20, 0xfd, 0xcf, 0x5d, 0x48, 0x00,
	0xa0, 0x01, 0x66, 0x89, 0x0e, 0x15, 0xb5, 0x56, 0x6a, 0x94, 0x5b, 0xcf, 0xf4, 0xa2, 0x69, 0xe8,
	0xb9, 0x8c, 0x76, 0x25, 0x58, 0x5a, 0x6a, 0x02, 0x0e, 0x1f, 0x83, 0x5b, 0x01, 0x62, 0x82, 0x20,
	0xcf, 0x3e, 0x44, 0xc4, 0x0b, 0x19, 0xae, 0x94, 0x6a, 0x4a, 0xe3, 0x46, 0xef, 0x66, 0xea, 0xde,
	0x4a, 0xbc, 0x70, 0x05, 0x2c, 0x8d, 0x91, 0x47, 0x5c, 0x24, 0xb0, 0x4d, 0x7d, 0xef, 0xa8, 0x32,
	0x17, 0xa7, 0x2d, 0x4a, 0xe7, 0xae, 0xef, 0x1d, 0x69, 0x5f, 0x15, 0xa0, 0x15, 0x77, 0x01, 0xfb,
	0x60, 0xde, 0x61, 0x18, 0x89, 0x44, 0xd2, 0x72, 0x6b, 0x2d, 0x97, 0x5b, 0xb6, 0x47, 0xf9, 0xe4,
	0xb6, 0xff, 0xeb, 0xa5, 0x68, 0xb0, 0x02, 0xe6, 0x19, 0x1e, 0xd1, 0x31, 0xae, 0xa8, 0x91, 0xb0,
	0x51, 0x24, 0xf9, 0xdd, 0x29, 0x83, 0x85, 0x8c, 0xb4, 0x76, 0xa6, 0x80, 0x47, 0x45, 0xe3, 0xe2,
	0x01, 0xf5, 0x39, 0x86, 0x5b, 0xe0, 0xce, 0x05, 0x79, 0x6c, 0xcc, 0x18, 0x65, 0xb1, 0x48, 0xe5,
	0x16, 0x94, 0x8d, 0xb3, 0xc0, 0xd1, 0xf7, 0xe3, 0x1d, 0xeb, 0xdd, 0x9e, 0x15, 0x6e, 0x33, 0x4a,
	0x87, 0x6f, 0xc0, 0xff, 0x0c, 0xf3, 0xd0, 0x13, 0x72, 0x9c, 0x5b, 0xc5, 0xe3, 0x2c, 0x6c, 0x31,
	0xf4, 0x44, 0x4f, 0xc2, 0x6a, 0x2f, 0x41, 0xfd, 0x2a, 0x1f, 0x44, 0x73, 0xfc, 0xcb, 0x56, 0xcf,
	0x2e, 0x74, 0xeb, 0x6c, 0x0e, 0xd4, 0x72, 0x05, 0xdf, 0x4f, 0x1a, 0x84, 0xbf, 0x15, 0xb0, 0x7c,
	0xd9, 0x3b, 0x82, 0x9b, 0xc5, 0x1c, 0xaf, 0xf0, 0x0e, 0xab, 0xd7, 0xda, 0x0e, 0xad, 0x77, 0x6e,
	0xcd, 0x12, 0x7e, 0xff, 0xfd, 0xe7, 0x27, 0x75, 0x0d, 0x9a, 0xd1, 0x99, 0x3a, 0x9e, 0x89, 0xac,
	0xcb, 0x27, 0xc8, 0x8d, 0x66, 0x76, 0xb7, 0x2e, 0xca, 0x68, 0x34, 0x4f, 0xe0, 0x67, 0x15, 0x3c,
	0xb8, 0x5c, 0x6c, 0xd8, 0xbd, 0xfe, 0x7c, 0x13, 0xf6, 0xdb, 0xff, 0x60, 0x51, 0xe2, 0x5d, 0xd6,
	0xd0, 0xb9, 0x75, 0x77, 0xea, 0xf8, 0x3c, 0x99, 0x5c, 0x81, 0x58, 0x92, 0x8e, 0xb6, 0x1e, 0x49,
	0x32, 0xd1, 0xe0, 0x78, 0x2a, 0x79, 0xbd, 0x79, 0x92, 0xab, 0x88, 0x39, 0x8a, 0xeb, 0x9b, 0x4a,
	0xb3, 0x7a, 0xff, 0xd4, 0xaa, 0x4c, 0x7a, 0x4c, 0xad, 0x80, 0x70, 0xdd, 0xa1, 0xa3, 0xce, 0x47,
	0x15, 0xd4, 0x1d, 0x3a, 0x2a, 0xe4, 0xd3, 0x59, 0x2d, 0x5a, 0xbd, 0xbd, 0xe8, 0x8c, 0xef, 0x29,
	0xaf, 0xb7, 0x53, 0xa8, 0x21, 0xf5, 0x90, 0x3f, 0xd4, 0x29, 0x1b, 0x1a, 0x43, 0xec, 0xc7, 0x47,
	0xde, 0x98, 0x14, 0xcf, 0xff, 0x2f, 0x7c, 0x2a, 0x8d, 0x2f, 0x6a, 0xa9, 0x6b, 0x59, 0xdf, 0xd4,
	0x5a, 0x37, 0x01, 0xb4, 0x5c, 0xae, 0x27, 0x66, 0x64, 0xf5, 0xdb, 0x7a, 0x5a, 0x98, 0x9f, 0xca,
	0x94, 0x81, 0xe5, 0xf2, 0x41, 0x96, 0x32, 0xe8, 0xb7, 0x07, 0x32, 0xe5, 0x97, 0x5a, 0x4f, 0xfc,
	0xa6, 0x69, 0xb9, 0xdc, 0x34, 0xb3, 0x24, 0xd3, 0xec, 0xb7, 0x4d, 0x53, 0xa6, 0x1d, 0xcc, 0xc7,
	0x7d, 0xb6, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x4d, 0x94, 0x2b, 0xb2, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerNegativeCriterionServiceClient is the client API for CustomerNegativeCriterionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerNegativeCriterionServiceClient interface {
	// Returns the requested criterion in full detail.
	GetCustomerNegativeCriterion(ctx context.Context, in *GetCustomerNegativeCriterionRequest, opts ...grpc.CallOption) (*resources.CustomerNegativeCriterion, error)
	// Creates or removes criteria. Operation statuses are returned.
	MutateCustomerNegativeCriteria(ctx context.Context, in *MutateCustomerNegativeCriteriaRequest, opts ...grpc.CallOption) (*MutateCustomerNegativeCriteriaResponse, error)
}

type customerNegativeCriterionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerNegativeCriterionServiceClient(cc grpc.ClientConnInterface) CustomerNegativeCriterionServiceClient {
	return &customerNegativeCriterionServiceClient{cc}
}

func (c *customerNegativeCriterionServiceClient) GetCustomerNegativeCriterion(ctx context.Context, in *GetCustomerNegativeCriterionRequest, opts ...grpc.CallOption) (*resources.CustomerNegativeCriterion, error) {
	out := new(resources.CustomerNegativeCriterion)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CustomerNegativeCriterionService/GetCustomerNegativeCriterion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerNegativeCriterionServiceClient) MutateCustomerNegativeCriteria(ctx context.Context, in *MutateCustomerNegativeCriteriaRequest, opts ...grpc.CallOption) (*MutateCustomerNegativeCriteriaResponse, error) {
	out := new(MutateCustomerNegativeCriteriaResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CustomerNegativeCriterionService/MutateCustomerNegativeCriteria", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerNegativeCriterionServiceServer is the server API for CustomerNegativeCriterionService service.
type CustomerNegativeCriterionServiceServer interface {
	// Returns the requested criterion in full detail.
	GetCustomerNegativeCriterion(context.Context, *GetCustomerNegativeCriterionRequest) (*resources.CustomerNegativeCriterion, error)
	// Creates or removes criteria. Operation statuses are returned.
	MutateCustomerNegativeCriteria(context.Context, *MutateCustomerNegativeCriteriaRequest) (*MutateCustomerNegativeCriteriaResponse, error)
}

// UnimplementedCustomerNegativeCriterionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerNegativeCriterionServiceServer struct {
}

func (*UnimplementedCustomerNegativeCriterionServiceServer) GetCustomerNegativeCriterion(ctx context.Context, req *GetCustomerNegativeCriterionRequest) (*resources.CustomerNegativeCriterion, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCustomerNegativeCriterion not implemented")
}
func (*UnimplementedCustomerNegativeCriterionServiceServer) MutateCustomerNegativeCriteria(ctx context.Context, req *MutateCustomerNegativeCriteriaRequest) (*MutateCustomerNegativeCriteriaResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCustomerNegativeCriteria not implemented")
}

func RegisterCustomerNegativeCriterionServiceServer(s *grpc.Server, srv CustomerNegativeCriterionServiceServer) {
	s.RegisterService(&_CustomerNegativeCriterionService_serviceDesc, srv)
}

func _CustomerNegativeCriterionService_GetCustomerNegativeCriterion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerNegativeCriterionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerNegativeCriterionServiceServer).GetCustomerNegativeCriterion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CustomerNegativeCriterionService/GetCustomerNegativeCriterion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerNegativeCriterionServiceServer).GetCustomerNegativeCriterion(ctx, req.(*GetCustomerNegativeCriterionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerNegativeCriterionService_MutateCustomerNegativeCriteria_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerNegativeCriteriaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerNegativeCriterionServiceServer).MutateCustomerNegativeCriteria(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CustomerNegativeCriterionService/MutateCustomerNegativeCriteria",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerNegativeCriterionServiceServer).MutateCustomerNegativeCriteria(ctx, req.(*MutateCustomerNegativeCriteriaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerNegativeCriterionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.CustomerNegativeCriterionService",
	HandlerType: (*CustomerNegativeCriterionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerNegativeCriterion",
			Handler:    _CustomerNegativeCriterionService_GetCustomerNegativeCriterion_Handler,
		},
		{
			MethodName: "MutateCustomerNegativeCriteria",
			Handler:    _CustomerNegativeCriterionService_MutateCustomerNegativeCriteria_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/customer_negative_criterion_service.proto",
}
