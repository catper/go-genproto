// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/external_conversion_source.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
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

// The external conversion source that is associated with a ConversionAction.
type ExternalConversionSourceEnum_ExternalConversionSource int32

const (
	// Not specified.
	ExternalConversionSourceEnum_UNSPECIFIED ExternalConversionSourceEnum_ExternalConversionSource = 0
	// Represents value unknown in this version.
	ExternalConversionSourceEnum_UNKNOWN ExternalConversionSourceEnum_ExternalConversionSource = 1
	// Conversion that occurs when a user navigates to a particular webpage
	// after viewing an ad; Displayed in Google Ads UI as 'Website'.
	ExternalConversionSourceEnum_WEBPAGE ExternalConversionSourceEnum_ExternalConversionSource = 2
	// Conversion that comes from linked Google Analytics goal or transaction;
	// Displayed in Google Ads UI as 'Analytics'.
	ExternalConversionSourceEnum_ANALYTICS ExternalConversionSourceEnum_ExternalConversionSource = 3
	// Website conversion that is uploaded through ConversionUploadService;
	// Displayed in Google Ads UI as 'Import from clicks'.
	ExternalConversionSourceEnum_UPLOAD ExternalConversionSourceEnum_ExternalConversionSource = 4
	// Conversion that occurs when a user clicks on a call extension directly on
	// an ad; Displayed in Google Ads UI as 'Calls from ads'.
	ExternalConversionSourceEnum_AD_CALL_METRICS ExternalConversionSourceEnum_ExternalConversionSource = 5
	// Conversion that occurs when a user calls a dynamically-generated phone
	// number (by installed javascript) from an advertiser's website after
	// clicking on an ad; Displayed in Google Ads UI as 'Calls from website'.
	ExternalConversionSourceEnum_WEBSITE_CALL_METRICS ExternalConversionSourceEnum_ExternalConversionSource = 6
	// Conversion that occurs when a user visits an advertiser's retail store
	// after clicking on a Google ad;
	// Displayed in Google Ads UI as 'Store visits'.
	ExternalConversionSourceEnum_STORE_VISITS ExternalConversionSourceEnum_ExternalConversionSource = 7
	// Conversion that occurs when a user takes an in-app action such as a
	// purchase in an Android app;
	// Displayed in Google Ads UI as 'Android in-app action'.
	ExternalConversionSourceEnum_ANDROID_IN_APP ExternalConversionSourceEnum_ExternalConversionSource = 8
	// Conversion that occurs when a user takes an in-app action such as a
	// purchase in an iOS app;
	// Displayed in Google Ads UI as 'iOS in-app action'.
	ExternalConversionSourceEnum_IOS_IN_APP ExternalConversionSourceEnum_ExternalConversionSource = 9
	// Conversion that occurs when a user opens an iOS app for the first time;
	// Displayed in Google Ads UI as 'iOS app install (first open)'.
	ExternalConversionSourceEnum_IOS_FIRST_OPEN ExternalConversionSourceEnum_ExternalConversionSource = 10
	// Legacy app conversions that do not have an AppPlatform provided;
	// Displayed in Google Ads UI as 'Mobile app'.
	ExternalConversionSourceEnum_APP_UNSPECIFIED ExternalConversionSourceEnum_ExternalConversionSource = 11
	// Conversion that occurs when a user opens an Android app for the first
	// time; Displayed in Google Ads UI as 'Android app install (first open)'.
	ExternalConversionSourceEnum_ANDROID_FIRST_OPEN ExternalConversionSourceEnum_ExternalConversionSource = 12
	// Call conversion that is uploaded through ConversionUploadService;
	// Displayed in Google Ads UI as 'Import from calls'.
	ExternalConversionSourceEnum_UPLOAD_CALLS ExternalConversionSourceEnum_ExternalConversionSource = 13
	// Conversion that comes from a linked Firebase event;
	// Displayed in Google Ads UI as 'Firebase'.
	ExternalConversionSourceEnum_FIREBASE ExternalConversionSourceEnum_ExternalConversionSource = 14
	// Conversion that occurs when a user clicks on a mobile phone number;
	// Displayed in Google Ads UI as 'Phone number clicks'.
	ExternalConversionSourceEnum_CLICK_TO_CALL ExternalConversionSourceEnum_ExternalConversionSource = 15
	// Conversion that comes from Salesforce;
	// Displayed in Google Ads UI as 'Salesforce.com'.
	ExternalConversionSourceEnum_SALESFORCE ExternalConversionSourceEnum_ExternalConversionSource = 16
	// Conversion that comes from in-store purchases recorded by CRM;
	// Displayed in Google Ads UI as 'Store sales (data partner)'.
	ExternalConversionSourceEnum_STORE_SALES_CRM ExternalConversionSourceEnum_ExternalConversionSource = 17
	// Conversion that comes from in-store purchases from payment network;
	// Displayed in Google Ads UI as 'Store sales (payment network)'.
	ExternalConversionSourceEnum_STORE_SALES_PAYMENT_NETWORK ExternalConversionSourceEnum_ExternalConversionSource = 18
	// Codeless Google Play conversion;
	// Displayed in Google Ads UI as 'Google Play'.
	ExternalConversionSourceEnum_GOOGLE_PLAY ExternalConversionSourceEnum_ExternalConversionSource = 19
	// Conversion that comes from a linked third-party app analytics event;
	// Displayed in Google Ads UI as 'Third-party app analytics'.
	ExternalConversionSourceEnum_THIRD_PARTY_APP_ANALYTICS ExternalConversionSourceEnum_ExternalConversionSource = 20
	// Conversion that is controlled by Google Attribution.
	ExternalConversionSourceEnum_GOOGLE_ATTRIBUTION ExternalConversionSourceEnum_ExternalConversionSource = 21
	// Store Sales conversion based on first-party or third-party merchant data
	// uploads. Displayed in Google Ads UI as 'Store sales (direct)'.
	ExternalConversionSourceEnum_STORE_SALES_DIRECT ExternalConversionSourceEnum_ExternalConversionSource = 22
)

var ExternalConversionSourceEnum_ExternalConversionSource_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "WEBPAGE",
	3:  "ANALYTICS",
	4:  "UPLOAD",
	5:  "AD_CALL_METRICS",
	6:  "WEBSITE_CALL_METRICS",
	7:  "STORE_VISITS",
	8:  "ANDROID_IN_APP",
	9:  "IOS_IN_APP",
	10: "IOS_FIRST_OPEN",
	11: "APP_UNSPECIFIED",
	12: "ANDROID_FIRST_OPEN",
	13: "UPLOAD_CALLS",
	14: "FIREBASE",
	15: "CLICK_TO_CALL",
	16: "SALESFORCE",
	17: "STORE_SALES_CRM",
	18: "STORE_SALES_PAYMENT_NETWORK",
	19: "GOOGLE_PLAY",
	20: "THIRD_PARTY_APP_ANALYTICS",
	21: "GOOGLE_ATTRIBUTION",
	22: "STORE_SALES_DIRECT",
}

var ExternalConversionSourceEnum_ExternalConversionSource_value = map[string]int32{
	"UNSPECIFIED":                 0,
	"UNKNOWN":                     1,
	"WEBPAGE":                     2,
	"ANALYTICS":                   3,
	"UPLOAD":                      4,
	"AD_CALL_METRICS":             5,
	"WEBSITE_CALL_METRICS":        6,
	"STORE_VISITS":                7,
	"ANDROID_IN_APP":              8,
	"IOS_IN_APP":                  9,
	"IOS_FIRST_OPEN":              10,
	"APP_UNSPECIFIED":             11,
	"ANDROID_FIRST_OPEN":          12,
	"UPLOAD_CALLS":                13,
	"FIREBASE":                    14,
	"CLICK_TO_CALL":               15,
	"SALESFORCE":                  16,
	"STORE_SALES_CRM":             17,
	"STORE_SALES_PAYMENT_NETWORK": 18,
	"GOOGLE_PLAY":                 19,
	"THIRD_PARTY_APP_ANALYTICS":   20,
	"GOOGLE_ATTRIBUTION":          21,
	"STORE_SALES_DIRECT":          22,
}

func (x ExternalConversionSourceEnum_ExternalConversionSource) String() string {
	return proto.EnumName(ExternalConversionSourceEnum_ExternalConversionSource_name, int32(x))
}

func (ExternalConversionSourceEnum_ExternalConversionSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b7b39ef7737c2983, []int{0, 0}
}

// Container for enum describing the external conversion source that is
// associated with a ConversionAction.
type ExternalConversionSourceEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExternalConversionSourceEnum) Reset()         { *m = ExternalConversionSourceEnum{} }
func (m *ExternalConversionSourceEnum) String() string { return proto.CompactTextString(m) }
func (*ExternalConversionSourceEnum) ProtoMessage()    {}
func (*ExternalConversionSourceEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7b39ef7737c2983, []int{0}
}

func (m *ExternalConversionSourceEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalConversionSourceEnum.Unmarshal(m, b)
}
func (m *ExternalConversionSourceEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalConversionSourceEnum.Marshal(b, m, deterministic)
}
func (m *ExternalConversionSourceEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalConversionSourceEnum.Merge(m, src)
}
func (m *ExternalConversionSourceEnum) XXX_Size() int {
	return xxx_messageInfo_ExternalConversionSourceEnum.Size(m)
}
func (m *ExternalConversionSourceEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalConversionSourceEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalConversionSourceEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.ExternalConversionSourceEnum_ExternalConversionSource", ExternalConversionSourceEnum_ExternalConversionSource_name, ExternalConversionSourceEnum_ExternalConversionSource_value)
	proto.RegisterType((*ExternalConversionSourceEnum)(nil), "google.ads.googleads.v3.enums.ExternalConversionSourceEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/external_conversion_source.proto", fileDescriptor_b7b39ef7737c2983)
}

var fileDescriptor_b7b39ef7737c2983 = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x66, 0xdd, 0xd8, 0x8f, 0xf7, 0xe7, 0x79, 0x63, 0x1a, 0xb0, 0x0a, 0x6d, 0x0f, 0x90, 0x5c,
	0xe4, 0x2e, 0x48, 0x48, 0x4e, 0xe2, 0x16, 0x6b, 0x59, 0x6c, 0xd9, 0x6e, 0xab, 0xa2, 0x4a, 0x56,
	0x58, 0xa3, 0xa8, 0xd2, 0x1a, 0x57, 0x4d, 0x5b, 0x71, 0xc9, 0xb3, 0x70, 0xc9, 0xa3, 0xf0, 0x12,
	0xdc, 0x73, 0xc9, 0x13, 0x20, 0x27, 0x6d, 0x29, 0x17, 0xe5, 0xc6, 0x3a, 0x3f, 0xdf, 0x39, 0xdf,
	0x77, 0xac, 0x0f, 0x7c, 0xc8, 0x8d, 0xc9, 0x9f, 0x33, 0x37, 0x1d, 0x96, 0x6e, 0x1d, 0xda, 0x68,
	0xe1, 0xb9, 0x59, 0x31, 0x1f, 0x97, 0x6e, 0xf6, 0x65, 0x96, 0x4d, 0x8b, 0xf4, 0x59, 0x3f, 0x99,
	0x62, 0x91, 0x4d, 0xcb, 0x91, 0x29, 0x74, 0x69, 0xe6, 0xd3, 0xa7, 0xcc, 0x99, 0x4c, 0xcd, 0xcc,
	0xa0, 0x66, 0x3d, 0xe4, 0xa4, 0xc3, 0xd2, 0x59, 0xcf, 0x3b, 0x0b, 0xcf, 0xa9, 0xe6, 0xdf, 0xdc,
	0xae, 0xd6, 0x4f, 0x46, 0x6e, 0x5a, 0x14, 0x66, 0x96, 0xce, 0x46, 0xa6, 0x28, 0xeb, 0xe1, 0xfb,
	0xaf, 0x7b, 0xe0, 0x96, 0x2c, 0x19, 0xc2, 0x35, 0x81, 0xac, 0xf6, 0x93, 0x62, 0x3e, 0xbe, 0xff,
	0xb9, 0x0b, 0x6e, 0xb6, 0x01, 0xd0, 0x39, 0x38, 0xee, 0x24, 0x92, 0x93, 0x90, 0xb6, 0x28, 0x89,
	0xe0, 0x0b, 0x74, 0x0c, 0x0e, 0x3a, 0xc9, 0x43, 0xc2, 0x7a, 0x09, 0xdc, 0xb1, 0x49, 0x8f, 0x04,
	0x1c, 0xb7, 0x09, 0x6c, 0xa0, 0x53, 0x70, 0x84, 0x13, 0x1c, 0xf7, 0x15, 0x0d, 0x25, 0xdc, 0x45,
	0x00, 0xec, 0x77, 0x78, 0xcc, 0x70, 0x04, 0xf7, 0xd0, 0x25, 0x38, 0xc7, 0x91, 0x0e, 0x71, 0x1c,
	0xeb, 0x47, 0xa2, 0x84, 0x05, 0xbc, 0x44, 0x37, 0xe0, 0xaa, 0x47, 0x02, 0x49, 0x15, 0xf9, 0xb7,
	0xb3, 0x8f, 0x20, 0x38, 0x91, 0x8a, 0x09, 0xa2, 0xbb, 0x54, 0x52, 0x25, 0xe1, 0x01, 0x42, 0xe0,
	0x0c, 0x27, 0x91, 0x60, 0x34, 0xd2, 0x34, 0xd1, 0x98, 0x73, 0x78, 0x88, 0xce, 0x00, 0xa0, 0x4c,
	0xae, 0xf2, 0x23, 0x8b, 0xb1, 0x79, 0x8b, 0x0a, 0xa9, 0x34, 0xe3, 0x24, 0x81, 0xa0, 0x22, 0xe6,
	0x5c, 0x6f, 0x9e, 0x70, 0x8c, 0xae, 0x01, 0x5a, 0x2d, 0xdb, 0x00, 0x9f, 0x58, 0xda, 0x5a, 0x71,
	0xa5, 0x47, 0xc2, 0x53, 0x74, 0x02, 0x0e, 0x5b, 0x54, 0x90, 0x00, 0x4b, 0x02, 0xcf, 0xd0, 0x05,
	0x38, 0x0d, 0x63, 0x1a, 0x3e, 0x68, 0xc5, 0x2a, 0x04, 0x3c, 0xb7, 0x1a, 0x24, 0x8e, 0x89, 0x6c,
	0x31, 0x11, 0x12, 0x08, 0x2d, 0x5f, 0xad, 0xbc, 0xaa, 0xea, 0x50, 0x3c, 0xc2, 0x0b, 0xf4, 0x0e,
	0xbc, 0xdd, 0x2c, 0x72, 0xdc, 0x7f, 0x24, 0x89, 0xd2, 0x09, 0x51, 0x3d, 0x26, 0x1e, 0x20, 0xb2,
	0x9f, 0xdc, 0x66, 0xac, 0x1d, 0x13, 0xcd, 0x63, 0xdc, 0x87, 0x97, 0xa8, 0x09, 0x5e, 0xab, 0x8f,
	0x54, 0x44, 0x9a, 0x63, 0xa1, 0xfa, 0xf6, 0x3e, 0xfd, 0xf7, 0x6b, 0xaf, 0xec, 0x01, 0x4b, 0x3c,
	0x56, 0x4a, 0xd0, 0xa0, 0xa3, 0x28, 0x4b, 0xe0, 0x2b, 0x5b, 0xdf, 0x24, 0x8a, 0xa8, 0x20, 0xa1,
	0x82, 0xd7, 0xc1, 0xef, 0x1d, 0x70, 0xf7, 0x64, 0xc6, 0xce, 0x7f, 0x6d, 0x14, 0x34, 0xb7, 0x99,
	0x80, 0x5b, 0x1f, 0xf1, 0x9d, 0x4f, 0xc1, 0x72, 0x3e, 0x37, 0xcf, 0x69, 0x91, 0x3b, 0x66, 0x9a,
	0xbb, 0x79, 0x56, 0x54, 0x2e, 0x5b, 0xd9, 0x7a, 0x32, 0x2a, 0xb7, 0xb8, 0xfc, 0x7d, 0xf5, 0x7e,
	0x6b, 0xec, 0xb6, 0x31, 0xfe, 0xde, 0x68, 0xb6, 0xeb, 0x55, 0x78, 0x58, 0x3a, 0x75, 0x68, 0xa3,
	0xae, 0xe7, 0x58, 0x47, 0x96, 0x3f, 0x56, 0xfd, 0x01, 0x1e, 0x96, 0x83, 0x75, 0x7f, 0xd0, 0xf5,
	0x06, 0x55, 0xff, 0x57, 0xe3, 0xae, 0x2e, 0xfa, 0x3e, 0x1e, 0x96, 0xbe, 0xbf, 0x46, 0xf8, 0x7e,
	0xd7, 0xf3, 0xfd, 0x0a, 0xf3, 0x79, 0xbf, 0x12, 0xe6, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x93,
	0x89, 0xa5, 0x1b, 0x7d, 0x03, 0x00, 0x00,
}
