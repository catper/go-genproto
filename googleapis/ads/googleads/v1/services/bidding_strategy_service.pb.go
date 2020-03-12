// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/bidding_strategy_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
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

// Request message for [BiddingStrategyService.GetBiddingStrategy][google.ads.googleads.v1.services.BiddingStrategyService.GetBiddingStrategy].
type GetBiddingStrategyRequest struct {
	// Required. The resource name of the bidding strategy to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBiddingStrategyRequest) Reset()         { *m = GetBiddingStrategyRequest{} }
func (m *GetBiddingStrategyRequest) String() string { return proto.CompactTextString(m) }
func (*GetBiddingStrategyRequest) ProtoMessage()    {}
func (*GetBiddingStrategyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_57d6d2619481cd67, []int{0}
}

func (m *GetBiddingStrategyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBiddingStrategyRequest.Unmarshal(m, b)
}
func (m *GetBiddingStrategyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBiddingStrategyRequest.Marshal(b, m, deterministic)
}
func (m *GetBiddingStrategyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBiddingStrategyRequest.Merge(m, src)
}
func (m *GetBiddingStrategyRequest) XXX_Size() int {
	return xxx_messageInfo_GetBiddingStrategyRequest.Size(m)
}
func (m *GetBiddingStrategyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBiddingStrategyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBiddingStrategyRequest proto.InternalMessageInfo

func (m *GetBiddingStrategyRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [BiddingStrategyService.MutateBiddingStrategies][google.ads.googleads.v1.services.BiddingStrategyService.MutateBiddingStrategies].
type MutateBiddingStrategiesRequest struct {
	// Required. The ID of the customer whose bidding strategies are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual bidding strategies.
	Operations []*BiddingStrategyOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateBiddingStrategiesRequest) Reset()         { *m = MutateBiddingStrategiesRequest{} }
func (m *MutateBiddingStrategiesRequest) String() string { return proto.CompactTextString(m) }
func (*MutateBiddingStrategiesRequest) ProtoMessage()    {}
func (*MutateBiddingStrategiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_57d6d2619481cd67, []int{1}
}

func (m *MutateBiddingStrategiesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateBiddingStrategiesRequest.Unmarshal(m, b)
}
func (m *MutateBiddingStrategiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateBiddingStrategiesRequest.Marshal(b, m, deterministic)
}
func (m *MutateBiddingStrategiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateBiddingStrategiesRequest.Merge(m, src)
}
func (m *MutateBiddingStrategiesRequest) XXX_Size() int {
	return xxx_messageInfo_MutateBiddingStrategiesRequest.Size(m)
}
func (m *MutateBiddingStrategiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateBiddingStrategiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateBiddingStrategiesRequest proto.InternalMessageInfo

func (m *MutateBiddingStrategiesRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateBiddingStrategiesRequest) GetOperations() []*BiddingStrategyOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateBiddingStrategiesRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateBiddingStrategiesRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a bidding strategy.
type BiddingStrategyOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*BiddingStrategyOperation_Create
	//	*BiddingStrategyOperation_Update
	//	*BiddingStrategyOperation_Remove
	Operation            isBiddingStrategyOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *BiddingStrategyOperation) Reset()         { *m = BiddingStrategyOperation{} }
func (m *BiddingStrategyOperation) String() string { return proto.CompactTextString(m) }
func (*BiddingStrategyOperation) ProtoMessage()    {}
func (*BiddingStrategyOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_57d6d2619481cd67, []int{2}
}

func (m *BiddingStrategyOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingStrategyOperation.Unmarshal(m, b)
}
func (m *BiddingStrategyOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingStrategyOperation.Marshal(b, m, deterministic)
}
func (m *BiddingStrategyOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingStrategyOperation.Merge(m, src)
}
func (m *BiddingStrategyOperation) XXX_Size() int {
	return xxx_messageInfo_BiddingStrategyOperation.Size(m)
}
func (m *BiddingStrategyOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingStrategyOperation.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingStrategyOperation proto.InternalMessageInfo

func (m *BiddingStrategyOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isBiddingStrategyOperation_Operation interface {
	isBiddingStrategyOperation_Operation()
}

type BiddingStrategyOperation_Create struct {
	Create *resources.BiddingStrategy `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type BiddingStrategyOperation_Update struct {
	Update *resources.BiddingStrategy `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type BiddingStrategyOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*BiddingStrategyOperation_Create) isBiddingStrategyOperation_Operation() {}

func (*BiddingStrategyOperation_Update) isBiddingStrategyOperation_Operation() {}

func (*BiddingStrategyOperation_Remove) isBiddingStrategyOperation_Operation() {}

func (m *BiddingStrategyOperation) GetOperation() isBiddingStrategyOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *BiddingStrategyOperation) GetCreate() *resources.BiddingStrategy {
	if x, ok := m.GetOperation().(*BiddingStrategyOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *BiddingStrategyOperation) GetUpdate() *resources.BiddingStrategy {
	if x, ok := m.GetOperation().(*BiddingStrategyOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *BiddingStrategyOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*BiddingStrategyOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BiddingStrategyOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BiddingStrategyOperation_Create)(nil),
		(*BiddingStrategyOperation_Update)(nil),
		(*BiddingStrategyOperation_Remove)(nil),
	}
}

// Response message for bidding strategy mutate.
type MutateBiddingStrategiesResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateBiddingStrategyResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *MutateBiddingStrategiesResponse) Reset()         { *m = MutateBiddingStrategiesResponse{} }
func (m *MutateBiddingStrategiesResponse) String() string { return proto.CompactTextString(m) }
func (*MutateBiddingStrategiesResponse) ProtoMessage()    {}
func (*MutateBiddingStrategiesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_57d6d2619481cd67, []int{3}
}

func (m *MutateBiddingStrategiesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateBiddingStrategiesResponse.Unmarshal(m, b)
}
func (m *MutateBiddingStrategiesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateBiddingStrategiesResponse.Marshal(b, m, deterministic)
}
func (m *MutateBiddingStrategiesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateBiddingStrategiesResponse.Merge(m, src)
}
func (m *MutateBiddingStrategiesResponse) XXX_Size() int {
	return xxx_messageInfo_MutateBiddingStrategiesResponse.Size(m)
}
func (m *MutateBiddingStrategiesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateBiddingStrategiesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateBiddingStrategiesResponse proto.InternalMessageInfo

func (m *MutateBiddingStrategiesResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateBiddingStrategiesResponse) GetResults() []*MutateBiddingStrategyResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the bidding strategy mutate.
type MutateBiddingStrategyResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateBiddingStrategyResult) Reset()         { *m = MutateBiddingStrategyResult{} }
func (m *MutateBiddingStrategyResult) String() string { return proto.CompactTextString(m) }
func (*MutateBiddingStrategyResult) ProtoMessage()    {}
func (*MutateBiddingStrategyResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_57d6d2619481cd67, []int{4}
}

func (m *MutateBiddingStrategyResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateBiddingStrategyResult.Unmarshal(m, b)
}
func (m *MutateBiddingStrategyResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateBiddingStrategyResult.Marshal(b, m, deterministic)
}
func (m *MutateBiddingStrategyResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateBiddingStrategyResult.Merge(m, src)
}
func (m *MutateBiddingStrategyResult) XXX_Size() int {
	return xxx_messageInfo_MutateBiddingStrategyResult.Size(m)
}
func (m *MutateBiddingStrategyResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateBiddingStrategyResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateBiddingStrategyResult proto.InternalMessageInfo

func (m *MutateBiddingStrategyResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetBiddingStrategyRequest)(nil), "google.ads.googleads.v1.services.GetBiddingStrategyRequest")
	proto.RegisterType((*MutateBiddingStrategiesRequest)(nil), "google.ads.googleads.v1.services.MutateBiddingStrategiesRequest")
	proto.RegisterType((*BiddingStrategyOperation)(nil), "google.ads.googleads.v1.services.BiddingStrategyOperation")
	proto.RegisterType((*MutateBiddingStrategiesResponse)(nil), "google.ads.googleads.v1.services.MutateBiddingStrategiesResponse")
	proto.RegisterType((*MutateBiddingStrategyResult)(nil), "google.ads.googleads.v1.services.MutateBiddingStrategyResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/bidding_strategy_service.proto", fileDescriptor_57d6d2619481cd67)
}

var fileDescriptor_57d6d2619481cd67 = []byte{
	// 772 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6b, 0xe3, 0x46,
	0x14, 0xaf, 0xe4, 0x90, 0x36, 0xe3, 0xa4, 0x85, 0x29, 0x4d, 0x54, 0xa7, 0x34, 0x46, 0x0d, 0xd4,
	0x98, 0x22, 0x61, 0x87, 0x96, 0x22, 0x13, 0x5a, 0x19, 0xf2, 0x51, 0xda, 0x34, 0x41, 0x81, 0x14,
	0x82, 0x41, 0x8c, 0xa5, 0x89, 0x2a, 0x22, 0x69, 0xd4, 0x99, 0x91, 0xc1, 0x84, 0x5c, 0x72, 0xe9,
	0x1f, 0xd0, 0x5b, 0x8f, 0x3d, 0xf6, 0xcf, 0xe8, 0x31, 0xc7, 0xee, 0x2d, 0xa7, 0x3d, 0xec, 0x69,
	0xef, 0x7b, 0xd9, 0xd3, 0x22, 0x8d, 0xc6, 0x5f, 0xb1, 0xd7, 0x90, 0xdc, 0x9e, 0xe6, 0xfd, 0xe6,
	0xf7, 0x3e, 0x7e, 0x6f, 0x9e, 0xc0, 0x0f, 0x01, 0x21, 0x41, 0x84, 0x4d, 0xe4, 0x33, 0x53, 0x98,
	0xb9, 0x35, 0x68, 0x99, 0x0c, 0xd3, 0x41, 0xe8, 0x61, 0x66, 0xf6, 0x43, 0xdf, 0x0f, 0x93, 0xc0,
	0x65, 0x9c, 0x22, 0x8e, 0x83, 0xa1, 0x5b, 0x7a, 0x8c, 0x94, 0x12, 0x4e, 0x60, 0x5d, 0xdc, 0x32,
	0x90, 0xcf, 0x8c, 0x11, 0x81, 0x31, 0x68, 0x19, 0x92, 0xa0, 0xf6, 0xfd, 0xa2, 0x10, 0x14, 0x33,
	0x92, 0xd1, 0x79, 0x31, 0x04, 0x77, 0xed, 0x0b, 0x79, 0x33, 0x0d, 0x4d, 0x94, 0x24, 0x84, 0x23,
	0x1e, 0x92, 0x84, 0x95, 0xde, 0xad, 0x09, 0xaf, 0x17, 0x85, 0x38, 0xe1, 0xa5, 0x63, 0x67, 0xc2,
	0x71, 0x15, 0xe2, 0xc8, 0x77, 0xfb, 0xf8, 0x77, 0x34, 0x08, 0x09, 0x2d, 0x01, 0x65, 0xce, 0x66,
	0xf1, 0xd5, 0xcf, 0xae, 0x4a, 0x54, 0x8c, 0xd8, 0xf5, 0x0c, 0x37, 0x4d, 0x3d, 0x93, 0x71, 0xc4,
	0xb3, 0x32, 0xa8, 0x7e, 0x00, 0x3e, 0x3f, 0xc2, 0xbc, 0x2b, 0xf2, 0x3d, 0x2f, 0xd3, 0x75, 0xf0,
	0x1f, 0x19, 0x66, 0x1c, 0x36, 0xc0, 0x86, 0xac, 0xc9, 0x4d, 0x50, 0x8c, 0x35, 0xa5, 0xae, 0x34,
	0xd6, 0xba, 0x95, 0x97, 0xb6, 0xea, 0xac, 0x4b, 0xcf, 0xaf, 0x28, 0xc6, 0xfa, 0x1b, 0x05, 0x7c,
	0x79, 0x92, 0x71, 0xc4, 0xf1, 0x34, 0x55, 0x88, 0x99, 0x24, 0xdb, 0x05, 0x55, 0x2f, 0x63, 0x9c,
	0xc4, 0x98, 0xba, 0xa1, 0x3f, 0x49, 0x05, 0xe4, 0xf9, 0x4f, 0x3e, 0x74, 0x01, 0x20, 0x29, 0xa6,
	0xa2, 0x31, 0x9a, 0x5a, 0xaf, 0x34, 0xaa, 0x6d, 0xcb, 0x58, 0xa6, 0x89, 0x31, 0x53, 0xc0, 0xa9,
	0xa4, 0x28, 0x03, 0x8c, 0x29, 0xe1, 0xd7, 0xe0, 0x93, 0x14, 0x51, 0x1e, 0xa2, 0xc8, 0xbd, 0x42,
	0x61, 0x94, 0x51, 0xac, 0x55, 0xea, 0x4a, 0xe3, 0x23, 0xe7, 0xe3, 0xf2, 0xf8, 0x50, 0x9c, 0xc2,
	0xaf, 0xc0, 0xc6, 0x00, 0x45, 0xa1, 0x8f, 0x38, 0x76, 0x49, 0x12, 0x0d, 0xb5, 0x95, 0x02, 0xb6,
	0x2e, 0x0f, 0x4f, 0x93, 0x68, 0xa8, 0xff, 0xad, 0x02, 0x6d, 0x51, 0x6c, 0xd8, 0x01, 0xd5, 0x2c,
	0x2d, 0xee, 0xe7, 0x4a, 0x14, 0xf7, 0xab, 0xed, 0x9a, 0x2c, 0x46, 0x8a, 0x65, 0x1c, 0xe6, 0x62,
	0x9d, 0x20, 0x76, 0xed, 0x00, 0x01, 0xcf, 0x6d, 0xf8, 0x0b, 0x58, 0xf5, 0x28, 0x46, 0x5c, 0x34,
	0xbd, 0xda, 0x6e, 0x2f, 0x6c, 0xc2, 0x68, 0xec, 0x66, 0xbb, 0x70, 0xfc, 0x81, 0x53, 0x72, 0xe4,
	0x6c, 0x82, 0x5b, 0x53, 0x9f, 0xc3, 0x26, 0x38, 0xa0, 0x06, 0x56, 0x29, 0x8e, 0xc9, 0x40, 0xb4,
	0x6e, 0x2d, 0xf7, 0x88, 0xef, 0x6e, 0x15, 0xac, 0x8d, 0x7a, 0xad, 0xff, 0xa7, 0x80, 0x9d, 0x85,
	0x43, 0xc1, 0x52, 0x92, 0x30, 0x0c, 0x0f, 0xc1, 0x67, 0x33, 0x72, 0xb8, 0x98, 0x52, 0x42, 0x0b,
	0xe6, 0x6a, 0x1b, 0xca, 0x3c, 0x69, 0xea, 0x19, 0xe7, 0xc5, 0xe0, 0x3a, 0x9f, 0x4e, 0x0b, 0x75,
	0x90, 0xc3, 0xe1, 0x6f, 0xe0, 0x43, 0x8a, 0x59, 0x16, 0x71, 0x39, 0x34, 0xfb, 0xcb, 0x87, 0x66,
	0x5e, 0x6e, 0x43, 0xa7, 0x60, 0x71, 0x24, 0x9b, 0xde, 0x05, 0xdb, 0xef, 0xc1, 0xe5, 0x53, 0x32,
	0xe7, 0x89, 0x4c, 0xbf, 0x8e, 0xf6, 0x9f, 0x2b, 0x60, 0x73, 0xe6, 0xfa, 0xb9, 0x48, 0x02, 0xfe,
	0xaf, 0x00, 0xf8, 0xf8, 0x01, 0xc2, 0xce, 0xf2, 0xec, 0x17, 0x3e, 0xdb, 0xda, 0x13, 0xc4, 0xd5,
	0x7f, 0x7e, 0xb0, 0xa7, 0x0b, 0xb9, 0x7b, 0xf1, 0xea, 0x2f, 0xf5, 0x5b, 0xb8, 0x97, 0x2f, 0xb6,
	0x9b, 0x29, 0xcf, 0xbe, 0x7c, 0xae, 0xcc, 0x6c, 0xca, 0x4d, 0x37, 0x56, 0xd6, 0x6c, 0xde, 0xc2,
	0xb7, 0x0a, 0xd8, 0x5a, 0x20, 0x3c, 0xfc, 0xf1, 0x69, 0xba, 0x8c, 0x17, 0x49, 0xcd, 0x7e, 0x06,
	0x83, 0x98, 0x3a, 0xfd, 0xf2, 0xc1, 0xde, 0x9c, 0x58, 0x46, 0xdf, 0x8c, 0xf7, 0x43, 0x51, 0x76,
	0x47, 0xff, 0x2e, 0x2f, 0x7b, 0x5c, 0xe7, 0xcd, 0x04, 0x78, 0xbf, 0x79, 0xfb, 0xb8, 0x6a, 0x2b,
	0x2e, 0x22, 0x5a, 0x4a, 0xb3, 0xb6, 0x7d, 0x6f, 0x6b, 0xe3, 0xac, 0x4a, 0x2b, 0x0d, 0x99, 0xe1,
	0x91, 0xb8, 0x7b, 0xa7, 0x82, 0x5d, 0x8f, 0xc4, 0x4b, 0x2b, 0xe8, 0x6e, 0xcf, 0x9f, 0x97, 0xb3,
	0x7c, 0x69, 0x9c, 0x29, 0x97, 0xc7, 0x25, 0x41, 0x40, 0x22, 0x94, 0x04, 0x06, 0xa1, 0x81, 0x19,
	0xe0, 0xa4, 0x58, 0x29, 0xe6, 0x38, 0xe4, 0xe2, 0xbf, 0x60, 0x47, 0x1a, 0xff, 0xa8, 0x95, 0x23,
	0xdb, 0xfe, 0x57, 0xad, 0x1f, 0x09, 0x42, 0xdb, 0x67, 0x86, 0x30, 0x73, 0xeb, 0xa2, 0x65, 0x94,
	0x81, 0xd9, 0xbd, 0x84, 0xf4, 0x6c, 0x9f, 0xf5, 0x46, 0x90, 0xde, 0x45, 0xab, 0x27, 0x21, 0xaf,
	0xd5, 0x5d, 0x71, 0x6e, 0x59, 0xb6, 0xcf, 0x2c, 0x6b, 0x04, 0xb2, 0xac, 0x8b, 0x96, 0x65, 0x49,
	0x58, 0x7f, 0xb5, 0xc8, 0x73, 0xef, 0x5d, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xe8, 0xa9, 0x77,
	0xac, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BiddingStrategyServiceClient is the client API for BiddingStrategyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BiddingStrategyServiceClient interface {
	// Returns the requested bidding strategy in full detail.
	GetBiddingStrategy(ctx context.Context, in *GetBiddingStrategyRequest, opts ...grpc.CallOption) (*resources.BiddingStrategy, error)
	// Creates, updates, or removes bidding strategies. Operation statuses are
	// returned.
	MutateBiddingStrategies(ctx context.Context, in *MutateBiddingStrategiesRequest, opts ...grpc.CallOption) (*MutateBiddingStrategiesResponse, error)
}

type biddingStrategyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBiddingStrategyServiceClient(cc grpc.ClientConnInterface) BiddingStrategyServiceClient {
	return &biddingStrategyServiceClient{cc}
}

func (c *biddingStrategyServiceClient) GetBiddingStrategy(ctx context.Context, in *GetBiddingStrategyRequest, opts ...grpc.CallOption) (*resources.BiddingStrategy, error) {
	out := new(resources.BiddingStrategy)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.BiddingStrategyService/GetBiddingStrategy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *biddingStrategyServiceClient) MutateBiddingStrategies(ctx context.Context, in *MutateBiddingStrategiesRequest, opts ...grpc.CallOption) (*MutateBiddingStrategiesResponse, error) {
	out := new(MutateBiddingStrategiesResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.BiddingStrategyService/MutateBiddingStrategies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BiddingStrategyServiceServer is the server API for BiddingStrategyService service.
type BiddingStrategyServiceServer interface {
	// Returns the requested bidding strategy in full detail.
	GetBiddingStrategy(context.Context, *GetBiddingStrategyRequest) (*resources.BiddingStrategy, error)
	// Creates, updates, or removes bidding strategies. Operation statuses are
	// returned.
	MutateBiddingStrategies(context.Context, *MutateBiddingStrategiesRequest) (*MutateBiddingStrategiesResponse, error)
}

// UnimplementedBiddingStrategyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBiddingStrategyServiceServer struct {
}

func (*UnimplementedBiddingStrategyServiceServer) GetBiddingStrategy(ctx context.Context, req *GetBiddingStrategyRequest) (*resources.BiddingStrategy, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetBiddingStrategy not implemented")
}
func (*UnimplementedBiddingStrategyServiceServer) MutateBiddingStrategies(ctx context.Context, req *MutateBiddingStrategiesRequest) (*MutateBiddingStrategiesResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateBiddingStrategies not implemented")
}

func RegisterBiddingStrategyServiceServer(s *grpc.Server, srv BiddingStrategyServiceServer) {
	s.RegisterService(&_BiddingStrategyService_serviceDesc, srv)
}

func _BiddingStrategyService_GetBiddingStrategy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBiddingStrategyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BiddingStrategyServiceServer).GetBiddingStrategy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.BiddingStrategyService/GetBiddingStrategy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BiddingStrategyServiceServer).GetBiddingStrategy(ctx, req.(*GetBiddingStrategyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BiddingStrategyService_MutateBiddingStrategies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateBiddingStrategiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BiddingStrategyServiceServer).MutateBiddingStrategies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.BiddingStrategyService/MutateBiddingStrategies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BiddingStrategyServiceServer).MutateBiddingStrategies(ctx, req.(*MutateBiddingStrategiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BiddingStrategyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.BiddingStrategyService",
	HandlerType: (*BiddingStrategyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBiddingStrategy",
			Handler:    _BiddingStrategyService_GetBiddingStrategy_Handler,
		},
		{
			MethodName: "MutateBiddingStrategies",
			Handler:    _BiddingStrategyService_MutateBiddingStrategies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/bidding_strategy_service.proto",
}
