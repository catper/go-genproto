// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/customer.proto

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

// A customer.
type Customer struct {
	// Immutable. The resource name of the customer.
	// Customer resource names have the form:
	//
	// `customers/{customer_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the customer.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Optional, non-unique descriptive name of the customer.
	DescriptiveName *wrappers.StringValue `protobuf:"bytes,4,opt,name=descriptive_name,json=descriptiveName,proto3" json:"descriptive_name,omitempty"`
	// Immutable. The currency in which the account operates.
	// A subset of the currency codes from the ISO 4217 standard is
	// supported.
	CurrencyCode *wrappers.StringValue `protobuf:"bytes,5,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// Immutable. The local timezone ID of the customer.
	TimeZone *wrappers.StringValue `protobuf:"bytes,6,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	// The URL template for constructing a tracking URL out of parameters.
	TrackingUrlTemplate *wrappers.StringValue `protobuf:"bytes,7,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3" json:"tracking_url_template,omitempty"`
	// The URL template for appending params to the final URL
	FinalUrlSuffix *wrappers.StringValue `protobuf:"bytes,11,opt,name=final_url_suffix,json=finalUrlSuffix,proto3" json:"final_url_suffix,omitempty"`
	// Whether auto-tagging is enabled for the customer.
	AutoTaggingEnabled *wrappers.BoolValue `protobuf:"bytes,8,opt,name=auto_tagging_enabled,json=autoTaggingEnabled,proto3" json:"auto_tagging_enabled,omitempty"`
	// Output only. Whether the Customer has a Partners program badge. If the Customer is not
	// associated with the Partners program, this will be false. For more
	// information, see https://support.google.com/partners/answer/3125774.
	HasPartnersBadge *wrappers.BoolValue `protobuf:"bytes,9,opt,name=has_partners_badge,json=hasPartnersBadge,proto3" json:"has_partners_badge,omitempty"`
	// Output only. Whether the customer is a manager.
	Manager *wrappers.BoolValue `protobuf:"bytes,12,opt,name=manager,proto3" json:"manager,omitempty"`
	// Output only. Whether the customer is a test account.
	TestAccount *wrappers.BoolValue `protobuf:"bytes,13,opt,name=test_account,json=testAccount,proto3" json:"test_account,omitempty"`
	// Call reporting setting for a customer.
	CallReportingSetting *CallReportingSetting `protobuf:"bytes,10,opt,name=call_reporting_setting,json=callReportingSetting,proto3" json:"call_reporting_setting,omitempty"`
	// Output only. Conversion tracking setting for a customer.
	ConversionTrackingSetting *ConversionTrackingSetting `protobuf:"bytes,14,opt,name=conversion_tracking_setting,json=conversionTrackingSetting,proto3" json:"conversion_tracking_setting,omitempty"`
	// Output only. Remarketing setting for a customer.
	RemarketingSetting *RemarketingSetting `protobuf:"bytes,15,opt,name=remarketing_setting,json=remarketingSetting,proto3" json:"remarketing_setting,omitempty"`
	// Output only. Reasons why the customer is not eligible to use PaymentMode.CONVERSIONS. If
	// the list is empty, the customer is eligible. This field is read-only.
	PayPerConversionEligibilityFailureReasons []enums.CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason `protobuf:"varint,16,rep,packed,name=pay_per_conversion_eligibility_failure_reasons,json=payPerConversionEligibilityFailureReasons,proto3,enum=google.ads.googleads.v1.enums.CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason" json:"pay_per_conversion_eligibility_failure_reasons,omitempty"`
	XXX_NoUnkeyedLiteral                      struct{}                                                                                                      `json:"-"`
	XXX_unrecognized                          []byte                                                                                                        `json:"-"`
	XXX_sizecache                             int32                                                                                                         `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f95889c65daf179, []int{0}
}

func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (m *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(m, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Customer) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Customer) GetDescriptiveName() *wrappers.StringValue {
	if m != nil {
		return m.DescriptiveName
	}
	return nil
}

func (m *Customer) GetCurrencyCode() *wrappers.StringValue {
	if m != nil {
		return m.CurrencyCode
	}
	return nil
}

func (m *Customer) GetTimeZone() *wrappers.StringValue {
	if m != nil {
		return m.TimeZone
	}
	return nil
}

func (m *Customer) GetTrackingUrlTemplate() *wrappers.StringValue {
	if m != nil {
		return m.TrackingUrlTemplate
	}
	return nil
}

func (m *Customer) GetFinalUrlSuffix() *wrappers.StringValue {
	if m != nil {
		return m.FinalUrlSuffix
	}
	return nil
}

func (m *Customer) GetAutoTaggingEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.AutoTaggingEnabled
	}
	return nil
}

func (m *Customer) GetHasPartnersBadge() *wrappers.BoolValue {
	if m != nil {
		return m.HasPartnersBadge
	}
	return nil
}

func (m *Customer) GetManager() *wrappers.BoolValue {
	if m != nil {
		return m.Manager
	}
	return nil
}

func (m *Customer) GetTestAccount() *wrappers.BoolValue {
	if m != nil {
		return m.TestAccount
	}
	return nil
}

func (m *Customer) GetCallReportingSetting() *CallReportingSetting {
	if m != nil {
		return m.CallReportingSetting
	}
	return nil
}

func (m *Customer) GetConversionTrackingSetting() *ConversionTrackingSetting {
	if m != nil {
		return m.ConversionTrackingSetting
	}
	return nil
}

func (m *Customer) GetRemarketingSetting() *RemarketingSetting {
	if m != nil {
		return m.RemarketingSetting
	}
	return nil
}

func (m *Customer) GetPayPerConversionEligibilityFailureReasons() []enums.CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason {
	if m != nil {
		return m.PayPerConversionEligibilityFailureReasons
	}
	return nil
}

// Call reporting setting for a customer.
type CallReportingSetting struct {
	// Enable reporting of phone call events by redirecting them via Google
	// System.
	CallReportingEnabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=call_reporting_enabled,json=callReportingEnabled,proto3" json:"call_reporting_enabled,omitempty"`
	// Whether to enable call conversion reporting.
	CallConversionReportingEnabled *wrappers.BoolValue `protobuf:"bytes,2,opt,name=call_conversion_reporting_enabled,json=callConversionReportingEnabled,proto3" json:"call_conversion_reporting_enabled,omitempty"`
	// Customer-level call conversion action to attribute a call conversion to.
	// If not set a default conversion action is used. Only in effect when
	// call_conversion_reporting_enabled is set to true.
	CallConversionAction *wrappers.StringValue `protobuf:"bytes,9,opt,name=call_conversion_action,json=callConversionAction,proto3" json:"call_conversion_action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CallReportingSetting) Reset()         { *m = CallReportingSetting{} }
func (m *CallReportingSetting) String() string { return proto.CompactTextString(m) }
func (*CallReportingSetting) ProtoMessage()    {}
func (*CallReportingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f95889c65daf179, []int{1}
}

func (m *CallReportingSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallReportingSetting.Unmarshal(m, b)
}
func (m *CallReportingSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallReportingSetting.Marshal(b, m, deterministic)
}
func (m *CallReportingSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallReportingSetting.Merge(m, src)
}
func (m *CallReportingSetting) XXX_Size() int {
	return xxx_messageInfo_CallReportingSetting.Size(m)
}
func (m *CallReportingSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_CallReportingSetting.DiscardUnknown(m)
}

var xxx_messageInfo_CallReportingSetting proto.InternalMessageInfo

func (m *CallReportingSetting) GetCallReportingEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.CallReportingEnabled
	}
	return nil
}

func (m *CallReportingSetting) GetCallConversionReportingEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.CallConversionReportingEnabled
	}
	return nil
}

func (m *CallReportingSetting) GetCallConversionAction() *wrappers.StringValue {
	if m != nil {
		return m.CallConversionAction
	}
	return nil
}

// A collection of customer-wide settings related to Google Ads Conversion
// Tracking.
type ConversionTrackingSetting struct {
	// Output only. The conversion tracking id used for this account. This id is automatically
	// assigned after any conversion tracking feature is used. If the customer
	// doesn't use conversion tracking, this is 0. This field is read-only.
	ConversionTrackingId *wrappers.Int64Value `protobuf:"bytes,1,opt,name=conversion_tracking_id,json=conversionTrackingId,proto3" json:"conversion_tracking_id,omitempty"`
	// Output only. The conversion tracking id of the customer's manager. This is set when the
	// customer is opted into cross account conversion tracking, and it overrides
	// conversion_tracking_id. This field can only be managed through the Google
	// Ads UI. This field is read-only.
	CrossAccountConversionTrackingId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=cross_account_conversion_tracking_id,json=crossAccountConversionTrackingId,proto3" json:"cross_account_conversion_tracking_id,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}             `json:"-"`
	XXX_unrecognized                 []byte               `json:"-"`
	XXX_sizecache                    int32                `json:"-"`
}

func (m *ConversionTrackingSetting) Reset()         { *m = ConversionTrackingSetting{} }
func (m *ConversionTrackingSetting) String() string { return proto.CompactTextString(m) }
func (*ConversionTrackingSetting) ProtoMessage()    {}
func (*ConversionTrackingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f95889c65daf179, []int{2}
}

func (m *ConversionTrackingSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionTrackingSetting.Unmarshal(m, b)
}
func (m *ConversionTrackingSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionTrackingSetting.Marshal(b, m, deterministic)
}
func (m *ConversionTrackingSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionTrackingSetting.Merge(m, src)
}
func (m *ConversionTrackingSetting) XXX_Size() int {
	return xxx_messageInfo_ConversionTrackingSetting.Size(m)
}
func (m *ConversionTrackingSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionTrackingSetting.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionTrackingSetting proto.InternalMessageInfo

func (m *ConversionTrackingSetting) GetConversionTrackingId() *wrappers.Int64Value {
	if m != nil {
		return m.ConversionTrackingId
	}
	return nil
}

func (m *ConversionTrackingSetting) GetCrossAccountConversionTrackingId() *wrappers.Int64Value {
	if m != nil {
		return m.CrossAccountConversionTrackingId
	}
	return nil
}

// Remarketing setting for a customer.
type RemarketingSetting struct {
	// Output only. The Google global site tag.
	GoogleGlobalSiteTag  *wrappers.StringValue `protobuf:"bytes,1,opt,name=google_global_site_tag,json=googleGlobalSiteTag,proto3" json:"google_global_site_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RemarketingSetting) Reset()         { *m = RemarketingSetting{} }
func (m *RemarketingSetting) String() string { return proto.CompactTextString(m) }
func (*RemarketingSetting) ProtoMessage()    {}
func (*RemarketingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f95889c65daf179, []int{3}
}

func (m *RemarketingSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemarketingSetting.Unmarshal(m, b)
}
func (m *RemarketingSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemarketingSetting.Marshal(b, m, deterministic)
}
func (m *RemarketingSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemarketingSetting.Merge(m, src)
}
func (m *RemarketingSetting) XXX_Size() int {
	return xxx_messageInfo_RemarketingSetting.Size(m)
}
func (m *RemarketingSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_RemarketingSetting.DiscardUnknown(m)
}

var xxx_messageInfo_RemarketingSetting proto.InternalMessageInfo

func (m *RemarketingSetting) GetGoogleGlobalSiteTag() *wrappers.StringValue {
	if m != nil {
		return m.GoogleGlobalSiteTag
	}
	return nil
}

func init() {
	proto.RegisterType((*Customer)(nil), "google.ads.googleads.v1.resources.Customer")
	proto.RegisterType((*CallReportingSetting)(nil), "google.ads.googleads.v1.resources.CallReportingSetting")
	proto.RegisterType((*ConversionTrackingSetting)(nil), "google.ads.googleads.v1.resources.ConversionTrackingSetting")
	proto.RegisterType((*RemarketingSetting)(nil), "google.ads.googleads.v1.resources.RemarketingSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/customer.proto", fileDescriptor_3f95889c65daf179)
}

var fileDescriptor_3f95889c65daf179 = []byte{
	// 974 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xdd, 0x4e, 0xdc, 0x46,
	0x14, 0xc7, 0xe5, 0xdd, 0x92, 0xc0, 0xf0, 0x11, 0x34, 0xd0, 0xc8, 0x90, 0x28, 0x05, 0xda, 0x48,
	0xa0, 0x4a, 0xde, 0x90, 0x7e, 0x29, 0xdb, 0xa8, 0x92, 0x17, 0x01, 0x4a, 0xd5, 0x46, 0xab, 0x85,
	0x70, 0x81, 0x90, 0x46, 0xb3, 0xf6, 0x59, 0x33, 0x62, 0x3c, 0x63, 0xcd, 0x8c, 0x69, 0x69, 0xd5,
	0x8b, 0x56, 0x7d, 0x91, 0xf6, 0xb2, 0x8f, 0xd2, 0xeb, 0x3e, 0x40, 0x6e, 0x7a, 0x13, 0xa9, 0x2f,
	0xc0, 0x55, 0xe5, 0x8f, 0xf1, 0x2e, 0x59, 0x36, 0xeb, 0x5c, 0x31, 0x78, 0xce, 0xff, 0xf7, 0x3f,
	0x67, 0xec, 0x73, 0x66, 0xd1, 0x93, 0x48, 0xca, 0x88, 0x43, 0x8b, 0x86, 0xba, 0x55, 0x2c, 0xb3,
	0xd5, 0xe5, 0x6e, 0x4b, 0x81, 0x96, 0xa9, 0x0a, 0x40, 0xb7, 0x82, 0x54, 0x1b, 0x19, 0x83, 0xf2,
	0x12, 0x25, 0x8d, 0xc4, 0x9b, 0x45, 0x98, 0x47, 0x43, 0xed, 0x55, 0x0a, 0xef, 0x72, 0xd7, 0xab,
	0x14, 0xeb, 0xa7, 0x93, 0xa0, 0x20, 0xd2, 0x78, 0x08, 0x24, 0x09, 0xbd, 0x22, 0x09, 0x28, 0x12,
	0x48, 0x71, 0x09, 0x4a, 0x33, 0x29, 0x08, 0x70, 0x16, 0xb1, 0x3e, 0xe3, 0xcc, 0x5c, 0x91, 0x01,
	0x65, 0x3c, 0x55, 0x40, 0x14, 0x50, 0x2d, 0x45, 0x61, 0xbf, 0xfe, 0x91, 0x65, 0x27, 0xac, 0x35,
	0x60, 0xc0, 0x43, 0xd2, 0x87, 0x73, 0x7a, 0xc9, 0x64, 0x99, 0xdf, 0xfa, 0xda, 0x48, 0x80, 0x4d,
	0xa9, 0xdc, 0x7a, 0x54, 0x6e, 0xe5, 0xff, 0xf5, 0xd3, 0x41, 0xeb, 0x07, 0x45, 0x93, 0x04, 0x94,
	0x2e, 0xf7, 0x1f, 0x8e, 0x48, 0xa9, 0x10, 0xd2, 0x50, 0xc3, 0xa4, 0x28, 0x77, 0xb7, 0xfe, 0x98,
	0x47, 0xb3, 0x7b, 0x65, 0xea, 0xf8, 0x25, 0x5a, 0xb4, 0x70, 0x22, 0x68, 0x0c, 0xae, 0xb3, 0xe1,
	0x6c, 0xcf, 0x75, 0x76, 0x5e, 0xfb, 0x33, 0xd7, 0xfe, 0xc7, 0x68, 0x73, 0x78, 0x32, 0xe5, 0x2a,
	0x61, 0xda, 0x0b, 0x64, 0xdc, 0xb2, 0x84, 0xde, 0x82, 0xd5, 0xbf, 0xa4, 0x31, 0xe0, 0x27, 0xa8,
	0xc1, 0x42, 0xb7, 0xb9, 0xe1, 0x6c, 0xcf, 0x3f, 0x7d, 0x50, 0x6a, 0x3c, 0x9b, 0xa7, 0xf7, 0x42,
	0x98, 0x2f, 0x3f, 0x3f, 0xa1, 0x3c, 0x85, 0x4e, 0xf3, 0xb5, 0xdf, 0xec, 0x35, 0x58, 0x88, 0x0f,
	0xd1, 0x72, 0x08, 0x3a, 0x50, 0x2c, 0x31, 0xec, 0xb2, 0x4c, 0xe2, 0x83, 0x5c, 0xff, 0x70, 0x4c,
	0x7f, 0x64, 0x14, 0x13, 0x51, 0x0e, 0xe8, 0xdd, 0x1b, 0x51, 0xe5, 0xd6, 0x07, 0x68, 0x31, 0x48,
	0x95, 0x02, 0x11, 0x5c, 0x91, 0x40, 0x86, 0xe0, 0xce, 0x4c, 0xa7, 0x64, 0x69, 0xcc, 0xf4, 0x16,
	0xac, 0x6e, 0x4f, 0x86, 0x80, 0xbf, 0x41, 0x73, 0x86, 0xc5, 0x40, 0x7e, 0x92, 0x02, 0xdc, 0x3b,
	0x75, 0x19, 0xb3, 0x99, 0xe6, 0x54, 0x0a, 0xc0, 0x5d, 0xf4, 0xa1, 0x51, 0x34, 0xb8, 0x60, 0x22,
	0x22, 0xa9, 0xe2, 0xc4, 0x40, 0x9c, 0x70, 0x6a, 0xc0, 0xbd, 0x5b, 0xa3, 0xaa, 0x15, 0x2b, 0x7d,
	0xa5, 0xf8, 0x71, 0x29, 0xc4, 0x07, 0x68, 0x79, 0xc0, 0x04, 0xe5, 0x39, 0x4e, 0xa7, 0x83, 0x01,
	0xfb, 0xd1, 0x9d, 0xaf, 0x01, 0x5b, 0xca, 0x55, 0xaf, 0x14, 0x3f, 0xca, 0x35, 0xf8, 0x3b, 0xb4,
	0x4a, 0x53, 0x23, 0x89, 0xa1, 0x51, 0x94, 0x65, 0x07, 0x82, 0xf6, 0x39, 0x84, 0xee, 0x6c, 0xce,
	0x5a, 0x1f, 0x63, 0x75, 0xa4, 0xe4, 0x05, 0x09, 0x67, 0xba, 0xe3, 0x42, 0xb6, 0x5f, 0xa8, 0xf0,
	0xf7, 0x08, 0x9f, 0x53, 0x4d, 0x12, 0xaa, 0x8c, 0x00, 0xa5, 0x49, 0x9f, 0x86, 0x11, 0xb8, 0x73,
	0xd3, 0x58, 0xc5, 0x9b, 0x5f, 0x3e, 0xa7, 0xba, 0x5b, 0x2a, 0x3b, 0x99, 0x10, 0x3f, 0x43, 0x77,
	0x63, 0x2a, 0x68, 0x04, 0xca, 0x5d, 0xa8, 0xc7, 0xb0, 0xf1, 0xb8, 0x83, 0x16, 0x0c, 0x68, 0x43,
	0x68, 0x10, 0xc8, 0x54, 0x18, 0x77, 0xb1, 0x9e, 0x7e, 0x3e, 0x13, 0xf9, 0x85, 0x06, 0xc7, 0xe8,
	0x7e, 0x40, 0x39, 0x27, 0x0a, 0x12, 0xa9, 0x4c, 0x76, 0x3a, 0x1a, 0x4c, 0xf6, 0xd7, 0x45, 0x39,
	0xed, 0x2b, 0x6f, 0xea, 0xbc, 0xf0, 0xf6, 0x28, 0xe7, 0x3d, 0xab, 0x3f, 0x2a, 0xe4, 0xbd, 0xd5,
	0xe0, 0x96, 0xa7, 0xf8, 0x57, 0x07, 0x3d, 0x18, 0x19, 0x17, 0xd5, 0x07, 0x63, 0x4d, 0x97, 0x72,
	0xd3, 0xe7, 0x75, 0x4c, 0x2b, 0xca, 0x71, 0x09, 0x29, 0x3d, 0x8a, 0x22, 0xd7, 0x82, 0x49, 0xfb,
	0xf8, 0x02, 0xad, 0x28, 0x88, 0xa9, 0xba, 0x80, 0x1b, 0xf5, 0xde, 0xcb, 0xad, 0xbf, 0xa8, 0x61,
	0xdd, 0x1b, 0xaa, 0x6f, 0x78, 0x62, 0x35, 0xb6, 0x81, 0xff, 0x73, 0x90, 0xf7, 0x5e, 0x73, 0x52,
	0xbb, 0xcb, 0x1b, 0xcd, 0xed, 0xa5, 0xa7, 0xbf, 0x3b, 0x13, 0x33, 0xc9, 0xc7, 0xb0, 0x67, 0x27,
	0x51, 0x97, 0x5e, 0x75, 0x41, 0x0d, 0x8f, 0x63, 0x7f, 0x88, 0x3e, 0x28, 0xc8, 0xbd, 0x1c, 0xbc,
	0x2f, 0xd2, 0xf8, 0xbd, 0x45, 0x45, 0x61, 0x3b, 0x49, 0xcd, 0x70, 0xdd, 0x7e, 0xfe, 0xc6, 0x7f,
	0x56, 0x63, 0x7c, 0xe2, 0x55, 0x7b, 0x8b, 0xe8, 0xd6, 0xcf, 0x76, 0xf9, 0xcb, 0xd6, 0x3f, 0x0d,
	0xb4, 0x7a, 0xdb, 0xd7, 0x84, 0xbb, 0x63, 0x9f, 0xa9, 0x6d, 0x62, 0x67, 0x6a, 0x13, 0xdf, 0xfc,
	0x12, 0x6d, 0x1b, 0x03, 0xda, 0xcc, 0x89, 0x23, 0x2f, 0x65, 0x1c, 0xde, 0x98, 0x0a, 0x7f, 0x94,
	0x41, 0x86, 0x07, 0x33, 0x66, 0xf3, 0x9b, 0x53, 0x66, 0x3e, 0xe2, 0x43, 0x83, 0xec, 0x5e, 0x2a,
	0x47, 0xc6, 0xbb, 0x67, 0xac, 0x77, 0xed, 0x7f, 0x8a, 0x76, 0x26, 0x1f, 0x67, 0x05, 0xf5, 0x73,
	0x66, 0x51, 0xeb, 0xdb, 0x4f, 0xb7, 0xfe, 0x75, 0xd0, 0xda, 0xc4, 0x7e, 0xc1, 0x27, 0xe8, 0xfe,
	0x6d, 0x2d, 0xc9, 0xec, 0xd9, 0x4e, 0xbf, 0xcf, 0x56, 0xc7, 0x9b, 0xed, 0x45, 0x88, 0x13, 0xf4,
	0x49, 0xa0, 0xa4, 0xd6, 0x76, 0x3e, 0x91, 0x09, 0x2e, 0x8d, 0x9a, 0x2e, 0x1b, 0x39, 0xad, 0x1c,
	0x5c, 0x7b, 0xb7, 0x38, 0x6e, 0x71, 0x84, 0xc7, 0x7b, 0x33, 0xab, 0xaf, 0x40, 0x93, 0x88, 0xcb,
	0x3e, 0xe5, 0x44, 0x33, 0x03, 0xd9, 0x65, 0x50, 0xd6, 0x37, 0xf5, 0x96, 0x6b, 0xf6, 0x56, 0x8a,
	0x88, 0xc3, 0x5c, 0x7f, 0xc4, 0x0c, 0x1c, 0xd3, 0xa8, 0x73, 0xed, 0xa0, 0xc7, 0x81, 0x8c, 0xa7,
	0x0f, 0x8c, 0xce, 0x62, 0xd5, 0x76, 0x99, 0x43, 0xd7, 0x39, 0xfd, 0xb6, 0xd4, 0x44, 0x92, 0x53,
	0x11, 0x79, 0x52, 0x45, 0xad, 0x08, 0x44, 0xee, 0xdf, 0x1a, 0xbe, 0xe2, 0x77, 0xfc, 0xaa, 0xfb,
	0xba, 0x5a, 0xfd, 0xd9, 0x68, 0x1e, 0xfa, 0xfe, 0x5f, 0x8d, 0xcd, 0xc3, 0x02, 0xe9, 0x87, 0xda,
	0x2b, 0x96, 0xd9, 0xea, 0x64, 0xd7, 0xeb, 0xd9, 0xc8, 0xbf, 0x6d, 0xcc, 0x99, 0x1f, 0xea, 0xb3,
	0x2a, 0xe6, 0xec, 0x64, 0xf7, 0xac, 0x8a, 0x79, 0xd3, 0x78, 0x5c, 0x6c, 0xb4, 0xdb, 0x7e, 0xa8,
	0xdb, 0xed, 0x2a, 0xaa, 0xdd, 0x3e, 0xd9, 0x6d, 0xb7, 0xab, 0xb8, 0xfe, 0x9d, 0x3c, 0xd9, 0xcf,
	0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x17, 0xe5, 0x84, 0x8e, 0x81, 0x0a, 0x00, 0x00,
}
