// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/billing_setup.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A billing setup, which associates a payments account and an advertiser. A
// billing setup is specific to one advertiser.
type BillingSetup struct {
	// Immutable. The resource name of the billing setup.
	// BillingSetup resource names have the form:
	//
	// `customers/{customer_id}/billingSetups/{billing_setup_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the billing setup.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. The status of the billing setup.
	Status enums.BillingSetupStatusEnum_BillingSetupStatus `protobuf:"varint,3,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.BillingSetupStatusEnum_BillingSetupStatus" json:"status,omitempty"`
	// Immutable. The resource name of the payments account associated with this billing
	// setup. Payments resource names have the form:
	//
	// `customers/{customer_id}/paymentsAccounts/{payments_account_id}`
	// When setting up billing, this is used to signup with an existing payments
	// account (and then payments_account_info should not be set).
	// When getting a billing setup, this and payments_account_info will be
	// populated.
	PaymentsAccount *wrappers.StringValue `protobuf:"bytes,11,opt,name=payments_account,json=paymentsAccount,proto3" json:"payments_account,omitempty"`
	// Immutable. The payments account information associated with this billing setup.
	// When setting up billing, this is used to signup with a new payments account
	// (and then payments_account should not be set).
	// When getting a billing setup, this and payments_account will be
	// populated.
	PaymentsAccountInfo *BillingSetup_PaymentsAccountInfo `protobuf:"bytes,12,opt,name=payments_account_info,json=paymentsAccountInfo,proto3" json:"payments_account_info,omitempty"`
	// When creating a new billing setup, this is when the setup should take
	// effect. NOW is the only acceptable start time if the customer doesn't have
	// any approved setups.
	//
	// When fetching an existing billing setup, this is the requested start time.
	// However, if the setup was approved (see status) after the requested start
	// time, then this is the approval time.
	//
	// Types that are valid to be assigned to StartTime:
	//	*BillingSetup_StartDateTime
	//	*BillingSetup_StartTimeType
	StartTime isBillingSetup_StartTime `protobuf_oneof:"start_time"`
	// When the billing setup ends / ended. This is either FOREVER or the start
	// time of the next scheduled billing setup.
	//
	// Types that are valid to be assigned to EndTime:
	//	*BillingSetup_EndDateTime
	//	*BillingSetup_EndTimeType
	EndTime              isBillingSetup_EndTime `protobuf_oneof:"end_time"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *BillingSetup) Reset()         { *m = BillingSetup{} }
func (m *BillingSetup) String() string { return proto.CompactTextString(m) }
func (*BillingSetup) ProtoMessage()    {}
func (*BillingSetup) Descriptor() ([]byte, []int) {
	return fileDescriptor_4372b9605c7b101a, []int{0}
}

func (m *BillingSetup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillingSetup.Unmarshal(m, b)
}
func (m *BillingSetup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillingSetup.Marshal(b, m, deterministic)
}
func (m *BillingSetup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillingSetup.Merge(m, src)
}
func (m *BillingSetup) XXX_Size() int {
	return xxx_messageInfo_BillingSetup.Size(m)
}
func (m *BillingSetup) XXX_DiscardUnknown() {
	xxx_messageInfo_BillingSetup.DiscardUnknown(m)
}

var xxx_messageInfo_BillingSetup proto.InternalMessageInfo

func (m *BillingSetup) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *BillingSetup) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *BillingSetup) GetStatus() enums.BillingSetupStatusEnum_BillingSetupStatus {
	if m != nil {
		return m.Status
	}
	return enums.BillingSetupStatusEnum_UNSPECIFIED
}

func (m *BillingSetup) GetPaymentsAccount() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsAccount
	}
	return nil
}

func (m *BillingSetup) GetPaymentsAccountInfo() *BillingSetup_PaymentsAccountInfo {
	if m != nil {
		return m.PaymentsAccountInfo
	}
	return nil
}

type isBillingSetup_StartTime interface {
	isBillingSetup_StartTime()
}

type BillingSetup_StartDateTime struct {
	StartDateTime *wrappers.StringValue `protobuf:"bytes,9,opt,name=start_date_time,json=startDateTime,proto3,oneof"`
}

type BillingSetup_StartTimeType struct {
	StartTimeType enums.TimeTypeEnum_TimeType `protobuf:"varint,10,opt,name=start_time_type,json=startTimeType,proto3,enum=google.ads.googleads.v1.enums.TimeTypeEnum_TimeType,oneof"`
}

func (*BillingSetup_StartDateTime) isBillingSetup_StartTime() {}

func (*BillingSetup_StartTimeType) isBillingSetup_StartTime() {}

func (m *BillingSetup) GetStartTime() isBillingSetup_StartTime {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *BillingSetup) GetStartDateTime() *wrappers.StringValue {
	if x, ok := m.GetStartTime().(*BillingSetup_StartDateTime); ok {
		return x.StartDateTime
	}
	return nil
}

func (m *BillingSetup) GetStartTimeType() enums.TimeTypeEnum_TimeType {
	if x, ok := m.GetStartTime().(*BillingSetup_StartTimeType); ok {
		return x.StartTimeType
	}
	return enums.TimeTypeEnum_UNSPECIFIED
}

type isBillingSetup_EndTime interface {
	isBillingSetup_EndTime()
}

type BillingSetup_EndDateTime struct {
	EndDateTime *wrappers.StringValue `protobuf:"bytes,13,opt,name=end_date_time,json=endDateTime,proto3,oneof"`
}

type BillingSetup_EndTimeType struct {
	EndTimeType enums.TimeTypeEnum_TimeType `protobuf:"varint,14,opt,name=end_time_type,json=endTimeType,proto3,enum=google.ads.googleads.v1.enums.TimeTypeEnum_TimeType,oneof"`
}

func (*BillingSetup_EndDateTime) isBillingSetup_EndTime() {}

func (*BillingSetup_EndTimeType) isBillingSetup_EndTime() {}

func (m *BillingSetup) GetEndTime() isBillingSetup_EndTime {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *BillingSetup) GetEndDateTime() *wrappers.StringValue {
	if x, ok := m.GetEndTime().(*BillingSetup_EndDateTime); ok {
		return x.EndDateTime
	}
	return nil
}

func (m *BillingSetup) GetEndTimeType() enums.TimeTypeEnum_TimeType {
	if x, ok := m.GetEndTime().(*BillingSetup_EndTimeType); ok {
		return x.EndTimeType
	}
	return enums.TimeTypeEnum_UNSPECIFIED
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BillingSetup) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BillingSetup_StartDateTime)(nil),
		(*BillingSetup_StartTimeType)(nil),
		(*BillingSetup_EndDateTime)(nil),
		(*BillingSetup_EndTimeType)(nil),
	}
}

// Container of payments account information for this billing.
type BillingSetup_PaymentsAccountInfo struct {
	// Output only. A 16 digit id used to identify the payments account associated with the
	// billing setup.
	//
	// This must be passed as a string with dashes, e.g. "1234-5678-9012-3456".
	PaymentsAccountId *wrappers.StringValue `protobuf:"bytes,1,opt,name=payments_account_id,json=paymentsAccountId,proto3" json:"payments_account_id,omitempty"`
	// Immutable. The name of the payments account associated with the billing setup.
	//
	// This enables the user to specify a meaningful name for a payments account
	// to aid in reconciling monthly invoices.
	//
	// This name will be printed in the monthly invoices.
	PaymentsAccountName *wrappers.StringValue `protobuf:"bytes,2,opt,name=payments_account_name,json=paymentsAccountName,proto3" json:"payments_account_name,omitempty"`
	// Immutable. A 12 digit id used to identify the payments profile associated with the
	// billing setup.
	//
	// This must be passed in as a string with dashes, e.g. "1234-5678-9012".
	PaymentsProfileId *wrappers.StringValue `protobuf:"bytes,3,opt,name=payments_profile_id,json=paymentsProfileId,proto3" json:"payments_profile_id,omitempty"`
	// Output only. The name of the payments profile associated with the billing setup.
	PaymentsProfileName *wrappers.StringValue `protobuf:"bytes,4,opt,name=payments_profile_name,json=paymentsProfileName,proto3" json:"payments_profile_name,omitempty"`
	// Output only. A secondary payments profile id present in uncommon situations, e.g.
	// when a sequential liability agreement has been arranged.
	SecondaryPaymentsProfileId *wrappers.StringValue `protobuf:"bytes,5,opt,name=secondary_payments_profile_id,json=secondaryPaymentsProfileId,proto3" json:"secondary_payments_profile_id,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}              `json:"-"`
	XXX_unrecognized           []byte                `json:"-"`
	XXX_sizecache              int32                 `json:"-"`
}

func (m *BillingSetup_PaymentsAccountInfo) Reset()         { *m = BillingSetup_PaymentsAccountInfo{} }
func (m *BillingSetup_PaymentsAccountInfo) String() string { return proto.CompactTextString(m) }
func (*BillingSetup_PaymentsAccountInfo) ProtoMessage()    {}
func (*BillingSetup_PaymentsAccountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4372b9605c7b101a, []int{0, 0}
}

func (m *BillingSetup_PaymentsAccountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillingSetup_PaymentsAccountInfo.Unmarshal(m, b)
}
func (m *BillingSetup_PaymentsAccountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillingSetup_PaymentsAccountInfo.Marshal(b, m, deterministic)
}
func (m *BillingSetup_PaymentsAccountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillingSetup_PaymentsAccountInfo.Merge(m, src)
}
func (m *BillingSetup_PaymentsAccountInfo) XXX_Size() int {
	return xxx_messageInfo_BillingSetup_PaymentsAccountInfo.Size(m)
}
func (m *BillingSetup_PaymentsAccountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BillingSetup_PaymentsAccountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BillingSetup_PaymentsAccountInfo proto.InternalMessageInfo

func (m *BillingSetup_PaymentsAccountInfo) GetPaymentsAccountId() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsAccountId
	}
	return nil
}

func (m *BillingSetup_PaymentsAccountInfo) GetPaymentsAccountName() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsAccountName
	}
	return nil
}

func (m *BillingSetup_PaymentsAccountInfo) GetPaymentsProfileId() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsProfileId
	}
	return nil
}

func (m *BillingSetup_PaymentsAccountInfo) GetPaymentsProfileName() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsProfileName
	}
	return nil
}

func (m *BillingSetup_PaymentsAccountInfo) GetSecondaryPaymentsProfileId() *wrappers.StringValue {
	if m != nil {
		return m.SecondaryPaymentsProfileId
	}
	return nil
}

func init() {
	proto.RegisterType((*BillingSetup)(nil), "google.ads.googleads.v1.resources.BillingSetup")
	proto.RegisterType((*BillingSetup_PaymentsAccountInfo)(nil), "google.ads.googleads.v1.resources.BillingSetup.PaymentsAccountInfo")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/billing_setup.proto", fileDescriptor_4372b9605c7b101a)
}

var fileDescriptor_4372b9605c7b101a = []byte{
	// 740 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xbe, 0x49, 0x6e, 0xab, 0xdb, 0x49, 0xd2, 0xde, 0x3a, 0xba, 0x52, 0x6e, 0x28, 0xd0, 0x22,
	0x55, 0x44, 0x48, 0x1d, 0x37, 0xa5, 0x20, 0x64, 0x56, 0x0e, 0xa0, 0xb6, 0x20, 0xa1, 0xe0, 0x96,
	0x2c, 0x20, 0xc2, 0x9a, 0x78, 0x26, 0xc6, 0x52, 0x3c, 0x63, 0x79, 0xc6, 0x45, 0x51, 0xe9, 0x9b,
	0xb0, 0x62, 0xc9, 0xa3, 0xb0, 0xe1, 0x15, 0xba, 0xee, 0x23, 0xb0, 0x40, 0xc8, 0x63, 0x8f, 0xf3,
	0xdb, 0xd6, 0x12, 0xbb, 0xe3, 0x39, 0xdf, 0xf9, 0xce, 0x77, 0x7e, 0xc6, 0x03, 0x1e, 0xb9, 0x8c,
	0xb9, 0x43, 0xa2, 0x23, 0xcc, 0xf5, 0xc4, 0x8c, 0xad, 0xd3, 0x96, 0x1e, 0x12, 0xce, 0xa2, 0xd0,
	0x21, 0x5c, 0xef, 0x7b, 0xc3, 0xa1, 0x47, 0x5d, 0x9b, 0x13, 0x11, 0x05, 0x30, 0x08, 0x99, 0x60,
	0xda, 0x56, 0x82, 0x85, 0x08, 0x73, 0x98, 0x85, 0xc1, 0xd3, 0x16, 0xcc, 0xc2, 0x1a, 0x4f, 0xae,
	0x62, 0x26, 0x34, 0xf2, 0x67, 0x58, 0x6d, 0x2e, 0x90, 0x88, 0x78, 0x42, 0xde, 0xd8, 0xb9, 0x3e,
	0x52, 0x78, 0x3e, 0xb1, 0xc5, 0x28, 0x20, 0x29, 0xfc, 0xae, 0x82, 0x07, 0x9e, 0x3e, 0xf0, 0xc8,
	0x10, 0xdb, 0x7d, 0xf2, 0x11, 0x9d, 0x7a, 0x2c, 0x4c, 0x01, 0xff, 0x4f, 0x00, 0x94, 0xbe, 0xd4,
	0x75, 0x27, 0x75, 0xc9, 0xaf, 0x7e, 0x34, 0xd0, 0x3f, 0x85, 0x28, 0x08, 0x48, 0xa8, 0xa4, 0x6c,
	0x4c, 0x84, 0x22, 0x4a, 0x99, 0x40, 0xc2, 0x63, 0x34, 0xf5, 0xde, 0xfb, 0x52, 0x06, 0x95, 0x76,
	0x52, 0xc7, 0x71, 0x5c, 0x86, 0x66, 0x81, 0xaa, 0x4a, 0x60, 0x53, 0xe4, 0x93, 0x7a, 0x61, 0xb3,
	0xd0, 0x5c, 0x69, 0xef, 0x5c, 0x98, 0x4b, 0x3f, 0xcd, 0xfb, 0x60, 0x7b, 0xdc, 0xaa, 0xd4, 0x0a,
	0x3c, 0x0e, 0x1d, 0xe6, 0xeb, 0x93, 0x2c, 0x56, 0x45, 0x71, 0xbc, 0x46, 0x3e, 0xd1, 0x76, 0x41,
	0xd1, 0xc3, 0xf5, 0xe2, 0x66, 0xa1, 0x59, 0xde, 0xbb, 0x95, 0xc6, 0x41, 0xa5, 0x17, 0x1e, 0x51,
	0xf1, 0x78, 0xbf, 0x8b, 0x86, 0x11, 0x69, 0x97, 0x2e, 0xcc, 0x92, 0x55, 0xf4, 0xb0, 0x46, 0xc0,
	0x72, 0xd2, 0xcf, 0x7a, 0x69, 0xb3, 0xd0, 0x5c, 0xdd, 0x3b, 0x84, 0x57, 0x4d, 0x4b, 0x36, 0x14,
	0x4e, 0x26, 0x3f, 0x96, 0x81, 0x2f, 0x68, 0xe4, 0x2f, 0x38, 0x4e, 0x52, 0xa4, 0xe4, 0xda, 0x08,
	0xfc, 0x1b, 0xa0, 0x91, 0x4f, 0xa8, 0xe0, 0x36, 0x72, 0x1c, 0x16, 0x51, 0x51, 0x2f, 0x4b, 0x99,
	0x1b, 0x73, 0x32, 0x8f, 0x45, 0xe8, 0x51, 0x37, 0xd1, 0xb9, 0x2b, 0xbb, 0xf1, 0x00, 0x34, 0xaf,
	0xec, 0x46, 0x27, 0xa5, 0x35, 0x13, 0x56, 0x6b, 0x2d, 0x98, 0x3e, 0xd0, 0x3e, 0x83, 0xff, 0x66,
	0x53, 0xdb, 0x1e, 0x1d, 0xb0, 0x7a, 0x45, 0xe6, 0x7f, 0x06, 0x6f, 0x5c, 0xcf, 0xa9, 0xea, 0xe0,
	0x4c, 0xc2, 0x23, 0x3a, 0x60, 0x71, 0xad, 0x4b, 0x56, 0x2d, 0x98, 0xf7, 0x68, 0xaf, 0xc0, 0x1a,
	0x17, 0x28, 0x14, 0x36, 0x46, 0x82, 0xd8, 0xf1, 0x3a, 0xd6, 0x57, 0x72, 0xd4, 0x1d, 0x13, 0x1e,
	0xfe, 0x65, 0x55, 0x65, 0xec, 0x73, 0x24, 0xc8, 0x89, 0xe7, 0x13, 0xcd, 0x51, 0x64, 0xd9, 0x5a,
	0xd7, 0x81, 0x9c, 0xda, 0xfe, 0x0d, 0x53, 0x8b, 0xa3, 0x4f, 0x46, 0x01, 0x91, 0xb3, 0x52, 0x1f,
	0xd3, 0x49, 0xd4, 0xa9, 0x76, 0x00, 0xaa, 0x84, 0xe2, 0x09, 0xbd, 0xd5, 0x7c, 0x7a, 0x4b, 0x87,
	0x05, 0xab, 0x4c, 0x28, 0xce, 0xd4, 0xda, 0x09, 0xd1, 0x58, 0xeb, 0xea, 0x9f, 0x69, 0x55, 0x09,
	0xd4, 0x59, 0xe3, 0x47, 0x09, 0xd4, 0x16, 0x4c, 0x43, 0x7b, 0x03, 0x6a, 0xf3, 0x13, 0xc7, 0xf2,
	0x7e, 0xe5, 0xa9, 0xc3, 0x5a, 0x9f, 0x1d, 0x24, 0xd6, 0xde, 0x2e, 0x58, 0x22, 0x79, 0x69, 0x8b,
	0x39, 0x87, 0x39, 0xb7, 0x1d, 0xf2, 0xbe, 0x4e, 0x2a, 0x0d, 0x42, 0x36, 0xf0, 0x86, 0x24, 0x56,
	0x5a, 0xca, 0x4b, 0x9a, 0x29, 0xed, 0x24, 0xc1, 0x33, 0x4a, 0x15, 0xa5, 0x54, 0xfa, 0x77, 0xde,
	0xf2, 0x6b, 0x33, 0xa4, 0x52, 0x29, 0x06, 0xb7, 0x39, 0x71, 0x18, 0xc5, 0x28, 0x1c, 0xd9, 0x8b,
	0x34, 0x2f, 0xe5, 0xa5, 0x6f, 0x64, 0x3c, 0x9d, 0x59, 0xf1, 0xc6, 0x87, 0x4b, 0xf3, 0x7d, 0xce,
	0x3f, 0x9f, 0xb6, 0xe7, 0x44, 0x5c, 0x30, 0x9f, 0x84, 0x5c, 0x3f, 0x53, 0xe6, 0xb9, 0x7a, 0x2a,
	0x24, 0x84, 0xeb, 0x67, 0x53, 0x2f, 0xc7, 0x79, 0xbb, 0x02, 0xc0, 0xf8, 0x02, 0xb5, 0x01, 0xf8,
	0x47, 0x2d, 0x68, 0xfb, 0x57, 0x01, 0x6c, 0x3b, 0xcc, 0xbf, 0xf9, 0x67, 0xd0, 0x5e, 0x9f, 0x54,
	0xd1, 0x89, 0xcb, 0xec, 0x14, 0xde, 0xbd, 0x4c, 0xe3, 0x5c, 0x36, 0x44, 0xd4, 0x85, 0x2c, 0x74,
	0x75, 0x97, 0x50, 0xd9, 0x04, 0x7d, 0x5c, 0xc4, 0x35, 0x2f, 0xe7, 0xd3, 0xcc, 0xfa, 0x5a, 0x2c,
	0x1d, 0x98, 0xe6, 0xb7, 0xe2, 0xd6, 0x41, 0x42, 0x69, 0x62, 0x0e, 0x13, 0x33, 0xb6, 0xba, 0x2d,
	0x68, 0x29, 0xe4, 0x77, 0x85, 0xe9, 0x99, 0x98, 0xf7, 0x32, 0x4c, 0xaf, 0xdb, 0xea, 0x65, 0x98,
	0xcb, 0xe2, 0x76, 0xe2, 0x30, 0x0c, 0x13, 0x73, 0xc3, 0xc8, 0x50, 0x86, 0xd1, 0x6d, 0x19, 0x46,
	0x86, 0xeb, 0x2f, 0x4b, 0xb1, 0x0f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xdc, 0x0c, 0xb0,
	0xe5, 0x07, 0x00, 0x00,
}
