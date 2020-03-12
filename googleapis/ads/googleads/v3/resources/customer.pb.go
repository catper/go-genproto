// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/customer.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	wrappers "github.com/catper/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
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
	// The resource name of the customer.
	// Customer resource names have the form:
	//
	// `customers/{customer_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the customer.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Optional, non-unique descriptive name of the customer.
	DescriptiveName *wrappers.StringValue `protobuf:"bytes,4,opt,name=descriptive_name,json=descriptiveName,proto3" json:"descriptive_name,omitempty"`
	// The currency in which the account operates.
	// A subset of the currency codes from the ISO 4217 standard is
	// supported.
	CurrencyCode *wrappers.StringValue `protobuf:"bytes,5,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// The local timezone ID of the customer.
	TimeZone *wrappers.StringValue `protobuf:"bytes,6,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	// The URL template for constructing a tracking URL out of parameters.
	TrackingUrlTemplate *wrappers.StringValue `protobuf:"bytes,7,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3" json:"tracking_url_template,omitempty"`
	// The URL template for appending params to the final URL
	FinalUrlSuffix *wrappers.StringValue `protobuf:"bytes,11,opt,name=final_url_suffix,json=finalUrlSuffix,proto3" json:"final_url_suffix,omitempty"`
	// Whether auto-tagging is enabled for the customer.
	AutoTaggingEnabled *wrappers.BoolValue `protobuf:"bytes,8,opt,name=auto_tagging_enabled,json=autoTaggingEnabled,proto3" json:"auto_tagging_enabled,omitempty"`
	// Whether the Customer has a Partners program badge. If the Customer is not
	// associated with the Partners program, this will be false. For more
	// information, see https://support.google.com/partners/answer/3125774.
	HasPartnersBadge *wrappers.BoolValue `protobuf:"bytes,9,opt,name=has_partners_badge,json=hasPartnersBadge,proto3" json:"has_partners_badge,omitempty"`
	// Whether the customer is a manager.
	Manager *wrappers.BoolValue `protobuf:"bytes,12,opt,name=manager,proto3" json:"manager,omitempty"`
	// Whether the customer is a test account.
	TestAccount *wrappers.BoolValue `protobuf:"bytes,13,opt,name=test_account,json=testAccount,proto3" json:"test_account,omitempty"`
	// Call reporting setting for a customer.
	CallReportingSetting *CallReportingSetting `protobuf:"bytes,10,opt,name=call_reporting_setting,json=callReportingSetting,proto3" json:"call_reporting_setting,omitempty"`
	// Conversion tracking setting for a customer.
	ConversionTrackingSetting *ConversionTrackingSetting `protobuf:"bytes,14,opt,name=conversion_tracking_setting,json=conversionTrackingSetting,proto3" json:"conversion_tracking_setting,omitempty"`
	// Remarketing setting for a customer.
	RemarketingSetting *RemarketingSetting `protobuf:"bytes,15,opt,name=remarketing_setting,json=remarketingSetting,proto3" json:"remarketing_setting,omitempty"`
	// Reasons why the customer is not eligible to use PaymentMode.CONVERSIONS. If
	// the list is empty, the customer is eligible. This field is read-only.
	PayPerConversionEligibilityFailureReasons []enums.CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason `protobuf:"varint,16,rep,packed,name=pay_per_conversion_eligibility_failure_reasons,json=payPerConversionEligibilityFailureReasons,proto3,enum=google.ads.googleads.v3.enums.CustomerPayPerConversionEligibilityFailureReasonEnum_CustomerPayPerConversionEligibilityFailureReason" json:"pay_per_conversion_eligibility_failure_reasons,omitempty"`
	XXX_NoUnkeyedLiteral                      struct{}                                                                                                      `json:"-"`
	XXX_unrecognized                          []byte                                                                                                        `json:"-"`
	XXX_sizecache                             int32                                                                                                         `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_1221f0b49b42ae61, []int{0}
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
	return fileDescriptor_1221f0b49b42ae61, []int{1}
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
	// The conversion tracking id used for this account. This id is automatically
	// assigned after any conversion tracking feature is used. If the customer
	// doesn't use conversion tracking, this is 0. This field is read-only.
	ConversionTrackingId *wrappers.Int64Value `protobuf:"bytes,1,opt,name=conversion_tracking_id,json=conversionTrackingId,proto3" json:"conversion_tracking_id,omitempty"`
	// The conversion tracking id of the customer's manager. This is set when the
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
	return fileDescriptor_1221f0b49b42ae61, []int{2}
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
	// The Google global site tag.
	GoogleGlobalSiteTag  *wrappers.StringValue `protobuf:"bytes,1,opt,name=google_global_site_tag,json=googleGlobalSiteTag,proto3" json:"google_global_site_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RemarketingSetting) Reset()         { *m = RemarketingSetting{} }
func (m *RemarketingSetting) String() string { return proto.CompactTextString(m) }
func (*RemarketingSetting) ProtoMessage()    {}
func (*RemarketingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_1221f0b49b42ae61, []int{3}
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
	proto.RegisterType((*Customer)(nil), "google.ads.googleads.v3.resources.Customer")
	proto.RegisterType((*CallReportingSetting)(nil), "google.ads.googleads.v3.resources.CallReportingSetting")
	proto.RegisterType((*ConversionTrackingSetting)(nil), "google.ads.googleads.v3.resources.ConversionTrackingSetting")
	proto.RegisterType((*RemarketingSetting)(nil), "google.ads.googleads.v3.resources.RemarketingSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/customer.proto", fileDescriptor_1221f0b49b42ae61)
}

var fileDescriptor_1221f0b49b42ae61 = []byte{
	// 911 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0xc7, 0xb5, 0x1b, 0x68, 0x93, 0xc9, 0x47, 0xa3, 0xe9, 0x52, 0x39, 0x69, 0x55, 0x25, 0x81,
	0x4a, 0x41, 0x48, 0x5e, 0xd4, 0x2d, 0xa0, 0x9a, 0x72, 0xe1, 0x44, 0x69, 0x28, 0x42, 0x68, 0x71,
	0xd2, 0x5c, 0x44, 0x91, 0x46, 0xb3, 0xf6, 0x59, 0x77, 0x94, 0xf1, 0x8c, 0x35, 0x33, 0x0e, 0x84,
	0x8f, 0x3b, 0x9e, 0x04, 0xee, 0x78, 0x12, 0xc4, 0x2b, 0xf0, 0x02, 0xa8, 0x8f, 0xc0, 0x15, 0xf2,
	0xd8, 0xe3, 0x4d, 0xb4, 0xd9, 0x78, 0x7b, 0x95, 0x89, 0xcf, 0xf9, 0xff, 0xfe, 0xc7, 0x93, 0x73,
	0x8e, 0x83, 0x3e, 0x4d, 0xa5, 0x4c, 0x39, 0xf4, 0x69, 0xa2, 0xfb, 0xd5, 0xb1, 0x3c, 0x5d, 0x0c,
	0xfa, 0x0a, 0xb4, 0x2c, 0x54, 0x0c, 0xba, 0x1f, 0x17, 0xda, 0xc8, 0x0c, 0x94, 0x9f, 0x2b, 0x69,
	0x24, 0xde, 0xae, 0xd2, 0x7c, 0x9a, 0x68, 0xbf, 0x51, 0xf8, 0x17, 0x03, 0xbf, 0x51, 0x6c, 0x9e,
	0xce, 0x82, 0x82, 0x28, 0xb2, 0x09, 0x90, 0xe4, 0xf4, 0x92, 0xe4, 0xa0, 0x48, 0x2c, 0xc5, 0x05,
	0x28, 0xcd, 0xa4, 0x20, 0xc0, 0x59, 0xca, 0x46, 0x8c, 0x33, 0x73, 0x49, 0xc6, 0x94, 0xf1, 0x42,
	0x01, 0x51, 0x40, 0xb5, 0x14, 0x95, 0xfd, 0xe6, 0x86, 0x63, 0xe7, 0xac, 0xa9, 0xb1, 0x0e, 0x3d,
	0xae, 0x43, 0xf6, 0xb7, 0x51, 0x31, 0xee, 0xff, 0xa0, 0x68, 0x9e, 0x83, 0xd2, 0x75, 0xfc, 0xd1,
	0x15, 0x29, 0x15, 0x42, 0x1a, 0x6a, 0x98, 0x14, 0x75, 0x74, 0xe7, 0x2f, 0x84, 0x16, 0xf7, 0xeb,
	0xca, 0xf0, 0x87, 0x68, 0xd5, 0xc1, 0x89, 0xa0, 0x19, 0x78, 0x9d, 0xad, 0xce, 0xee, 0x52, 0xb4,
	0xe2, 0x1e, 0x7e, 0x47, 0x33, 0xc0, 0x9f, 0xa0, 0x2e, 0x4b, 0xbc, 0x85, 0xad, 0xce, 0xee, 0xf2,
	0xd3, 0x87, 0xf5, 0x5d, 0xf8, 0xce, 0xdc, 0x7f, 0x25, 0xcc, 0xe7, 0xcf, 0x4e, 0x28, 0x2f, 0x20,
	0xea, 0xb2, 0x04, 0x1f, 0xa2, 0xf5, 0x04, 0x74, 0xac, 0x58, 0x6e, 0xd8, 0x45, 0x0d, 0x7d, 0xcf,
	0x4a, 0x1f, 0x4d, 0x49, 0x8f, 0x8c, 0x62, 0x22, 0xad, 0xb4, 0xf7, 0xae, 0xa8, 0xac, 0x6b, 0x88,
	0x56, 0xe3, 0x42, 0x29, 0x10, 0xf1, 0x25, 0x89, 0x65, 0x02, 0xde, 0xfb, 0x73, 0x50, 0x56, 0x9c,
	0x64, 0x5f, 0x26, 0x80, 0x9f, 0xa3, 0x25, 0xc3, 0x32, 0x20, 0x3f, 0x49, 0x01, 0xde, 0x9d, 0x39,
	0xe4, 0x8b, 0x65, 0xfa, 0xa9, 0x14, 0x80, 0x87, 0xe8, 0x03, 0xa3, 0x68, 0x7c, 0xce, 0x44, 0x4a,
	0x0a, 0xc5, 0x89, 0x81, 0x2c, 0xe7, 0xd4, 0x80, 0x77, 0x77, 0x0e, 0xcc, 0x7d, 0x27, 0x7d, 0xad,
	0xf8, 0x71, 0x2d, 0xc4, 0x2f, 0xd1, 0xfa, 0x98, 0x09, 0xca, 0x2d, 0x4e, 0x17, 0xe3, 0x31, 0xfb,
	0xd1, 0x5b, 0x9e, 0x03, 0xb6, 0x66, 0x55, 0xaf, 0x15, 0x3f, 0xb2, 0x1a, 0xfc, 0x2d, 0xea, 0xd1,
	0xc2, 0x48, 0x62, 0x68, 0x9a, 0x96, 0xd5, 0x81, 0xa0, 0x23, 0x0e, 0x89, 0xb7, 0x68, 0x59, 0x9b,
	0x53, 0xac, 0x3d, 0x29, 0x79, 0x45, 0xc2, 0xa5, 0xee, 0xb8, 0x92, 0x1d, 0x54, 0x2a, 0xfc, 0x35,
	0xc2, 0x6f, 0xa8, 0x26, 0x39, 0x55, 0x46, 0x80, 0xd2, 0x64, 0x44, 0x93, 0x14, 0xbc, 0xa5, 0x56,
	0xd6, 0xfa, 0x1b, 0xaa, 0x87, 0xb5, 0x68, 0xaf, 0xd4, 0xe0, 0x67, 0xe8, 0x6e, 0x46, 0x05, 0x4d,
	0x41, 0x79, 0x2b, 0xad, 0x72, 0x97, 0x8a, 0xbf, 0x42, 0x2b, 0x06, 0xb4, 0x21, 0x34, 0x8e, 0x65,
	0x21, 0x8c, 0xb7, 0xda, 0x2a, 0x5d, 0x2e, 0xf3, 0xc3, 0x2a, 0x1d, 0x67, 0xe8, 0x41, 0x4c, 0x39,
	0x27, 0x0a, 0x72, 0xa9, 0x4c, 0x79, 0x1d, 0x1a, 0x4c, 0xf9, 0xd3, 0x43, 0x16, 0xf4, 0x85, 0xdf,
	0x3a, 0xc5, 0xfe, 0x3e, 0xe5, 0x3c, 0x72, 0xfa, 0xa3, 0x4a, 0x1e, 0xf5, 0xe2, 0x1b, 0x9e, 0xe2,
	0x5f, 0xd0, 0xc3, 0x2b, 0x33, 0xdc, 0x34, 0x88, 0xf3, 0x5c, 0xb3, 0x9e, 0x2f, 0xe6, 0xf1, 0x6c,
	0x28, 0xc7, 0x35, 0xc4, 0x19, 0x6f, 0xc4, 0xb3, 0x42, 0x78, 0x8c, 0xee, 0x2b, 0xc8, 0xa8, 0x3a,
	0x87, 0x6b, 0x6f, 0x7a, 0xcf, 0xba, 0x7e, 0x36, 0x87, 0x6b, 0x34, 0x51, 0x3b, 0x3b, 0xac, 0xa6,
	0x9e, 0xe1, 0x7f, 0x3b, 0xc8, 0x7f, 0xa7, 0x95, 0xa5, 0xbd, 0xf5, 0xad, 0x85, 0xdd, 0xb5, 0xa7,
	0xbf, 0x75, 0x66, 0x16, 0x61, 0x37, 0xa2, 0xef, 0xf6, 0xce, 0x90, 0x5e, 0x0e, 0x41, 0x4d, 0x2e,
	0xe1, 0x60, 0x82, 0x7e, 0x59, 0x91, 0x23, 0x0b, 0x3e, 0x10, 0x45, 0xf6, 0xce, 0xa2, 0xe8, 0xe3,
	0x7c, 0xce, 0x4c, 0x1d, 0xbc, 0x78, 0x1b, 0x3e, 0x47, 0xdb, 0x93, 0x4a, 0xeb, 0x53, 0xce, 0xb4,
	0x1f, 0xcb, 0xac, 0xdf, 0xec, 0xc9, 0x9e, 0xdb, 0xe5, 0xba, 0xff, 0xb3, 0x3b, 0xfe, 0xba, 0xf3,
	0x47, 0x17, 0xf5, 0x6e, 0xea, 0x1e, 0x3c, 0x9c, 0x6a, 0x4b, 0x37, 0xa5, 0x9d, 0xd6, 0xfe, 0xbe,
	0xde, 0x79, 0x6e, 0x4e, 0x01, 0x6d, 0x5b, 0xe2, 0x95, 0xbf, 0xc7, 0x34, 0xbc, 0xdb, 0x0a, 0x7f,
	0x5c, 0x42, 0x26, 0x17, 0x33, 0x65, 0x13, 0xd5, 0x85, 0x5f, 0xb1, 0xa1, 0x71, 0xf9, 0xf5, 0xa8,
	0x57, 0xc2, 0xed, 0xab, 0xaa, 0x77, 0x9d, 0x1e, 0x5a, 0xe5, 0xce, 0x3f, 0x1d, 0xb4, 0x31, 0xb3,
	0xdf, 0xf1, 0xf7, 0xe8, 0xc1, 0x4d, 0x23, 0xc5, 0xdc, 0x55, 0xdd, 0xfa, 0xc1, 0xe9, 0x4d, 0x0f,
	0xcb, 0xab, 0x04, 0x9f, 0xa3, 0x8f, 0x62, 0x25, 0xb5, 0x76, 0x4b, 0x85, 0xcc, 0x30, 0xe8, 0xb6,
	0x1b, 0x6c, 0x59, 0x50, 0xbd, 0x6d, 0xf6, 0x6f, 0x30, 0xdb, 0x49, 0x11, 0x9e, 0x1e, 0xab, 0xf2,
	0xad, 0x2a, 0x2a, 0x49, 0xb9, 0x1c, 0x51, 0x4e, 0x34, 0x33, 0x50, 0xae, 0xec, 0xfa, 0xad, 0x5a,
	0xbe, 0x1f, 0x55, 0xf0, 0xd0, 0x4a, 0x8f, 0x98, 0x81, 0x63, 0x9a, 0xee, 0xfd, 0xd7, 0x41, 0x4f,
	0x62, 0x99, 0xb5, 0x8f, 0xf9, 0xde, 0x6a, 0x33, 0x31, 0x25, 0x7c, 0xd8, 0x39, 0xfd, 0xa6, 0xd6,
	0xa4, 0x92, 0x53, 0x91, 0xfa, 0x52, 0xa5, 0xfd, 0x14, 0x84, 0xb5, 0xee, 0x4f, 0x3a, 0xfe, 0x96,
	0xff, 0x8d, 0xbe, 0x6c, 0x4e, 0xbf, 0x77, 0x17, 0x0e, 0xc3, 0xf0, 0xcf, 0xee, 0xf6, 0x61, 0x85,
	0x0c, 0x13, 0xed, 0x57, 0xc7, 0xf2, 0x74, 0x32, 0xf0, 0x23, 0x97, 0xf9, 0xb7, 0xcb, 0x39, 0x0b,
	0x13, 0x7d, 0xd6, 0xe4, 0x9c, 0x9d, 0x0c, 0xce, 0x9a, 0x9c, 0xb7, 0xdd, 0x27, 0x55, 0x20, 0x08,
	0xc2, 0x44, 0x07, 0x41, 0x93, 0x15, 0x04, 0x27, 0x83, 0x20, 0x68, 0xf2, 0x46, 0x77, 0x6c, 0xb1,
	0x83, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x17, 0xf8, 0x6e, 0x7f, 0xc7, 0x09, 0x00, 0x00,
}
