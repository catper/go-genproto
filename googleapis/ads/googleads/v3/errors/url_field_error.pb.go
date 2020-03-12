// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/url_field_error.proto

package errors

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

// Enum describing possible url field errors.
type UrlFieldErrorEnum_UrlFieldError int32

const (
	// Enum unspecified.
	UrlFieldErrorEnum_UNSPECIFIED UrlFieldErrorEnum_UrlFieldError = 0
	// The received error code is not known in this version.
	UrlFieldErrorEnum_UNKNOWN UrlFieldErrorEnum_UrlFieldError = 1
	// The tracking url template is invalid.
	UrlFieldErrorEnum_INVALID_TRACKING_URL_TEMPLATE UrlFieldErrorEnum_UrlFieldError = 2
	// The tracking url template contains invalid tag.
	UrlFieldErrorEnum_INVALID_TAG_IN_TRACKING_URL_TEMPLATE UrlFieldErrorEnum_UrlFieldError = 3
	// The tracking url template must contain at least one tag (e.g. {lpurl}),
	// This applies only to tracking url template associated with website ads or
	// product ads.
	UrlFieldErrorEnum_MISSING_TRACKING_URL_TEMPLATE_TAG UrlFieldErrorEnum_UrlFieldError = 4
	// The tracking url template must start with a valid protocol (or lpurl
	// tag).
	UrlFieldErrorEnum_MISSING_PROTOCOL_IN_TRACKING_URL_TEMPLATE UrlFieldErrorEnum_UrlFieldError = 5
	// The tracking url template starts with an invalid protocol.
	UrlFieldErrorEnum_INVALID_PROTOCOL_IN_TRACKING_URL_TEMPLATE UrlFieldErrorEnum_UrlFieldError = 6
	// The tracking url template contains illegal characters.
	UrlFieldErrorEnum_MALFORMED_TRACKING_URL_TEMPLATE UrlFieldErrorEnum_UrlFieldError = 7
	// The tracking url template must contain a host name (or lpurl tag).
	UrlFieldErrorEnum_MISSING_HOST_IN_TRACKING_URL_TEMPLATE UrlFieldErrorEnum_UrlFieldError = 8
	// The tracking url template has an invalid or missing top level domain
	// extension.
	UrlFieldErrorEnum_INVALID_TLD_IN_TRACKING_URL_TEMPLATE UrlFieldErrorEnum_UrlFieldError = 9
	// The tracking url template contains nested occurrences of the same
	// conditional tag (i.e. {ifmobile:{ifmobile:x}}).
	UrlFieldErrorEnum_REDUNDANT_NESTED_TRACKING_URL_TEMPLATE_TAG UrlFieldErrorEnum_UrlFieldError = 10
	// The final url is invalid.
	UrlFieldErrorEnum_INVALID_FINAL_URL UrlFieldErrorEnum_UrlFieldError = 11
	// The final url contains invalid tag.
	UrlFieldErrorEnum_INVALID_TAG_IN_FINAL_URL UrlFieldErrorEnum_UrlFieldError = 12
	// The final url contains nested occurrences of the same conditional tag
	// (i.e. {ifmobile:{ifmobile:x}}).
	UrlFieldErrorEnum_REDUNDANT_NESTED_FINAL_URL_TAG UrlFieldErrorEnum_UrlFieldError = 13
	// The final url must start with a valid protocol.
	UrlFieldErrorEnum_MISSING_PROTOCOL_IN_FINAL_URL UrlFieldErrorEnum_UrlFieldError = 14
	// The final url starts with an invalid protocol.
	UrlFieldErrorEnum_INVALID_PROTOCOL_IN_FINAL_URL UrlFieldErrorEnum_UrlFieldError = 15
	// The final url contains illegal characters.
	UrlFieldErrorEnum_MALFORMED_FINAL_URL UrlFieldErrorEnum_UrlFieldError = 16
	// The final url must contain a host name.
	UrlFieldErrorEnum_MISSING_HOST_IN_FINAL_URL UrlFieldErrorEnum_UrlFieldError = 17
	// The tracking url template has an invalid or missing top level domain
	// extension.
	UrlFieldErrorEnum_INVALID_TLD_IN_FINAL_URL UrlFieldErrorEnum_UrlFieldError = 18
	// The final mobile url is invalid.
	UrlFieldErrorEnum_INVALID_FINAL_MOBILE_URL UrlFieldErrorEnum_UrlFieldError = 19
	// The final mobile url contains invalid tag.
	UrlFieldErrorEnum_INVALID_TAG_IN_FINAL_MOBILE_URL UrlFieldErrorEnum_UrlFieldError = 20
	// The final mobile url contains nested occurrences of the same conditional
	// tag (i.e. {ifmobile:{ifmobile:x}}).
	UrlFieldErrorEnum_REDUNDANT_NESTED_FINAL_MOBILE_URL_TAG UrlFieldErrorEnum_UrlFieldError = 21
	// The final mobile url must start with a valid protocol.
	UrlFieldErrorEnum_MISSING_PROTOCOL_IN_FINAL_MOBILE_URL UrlFieldErrorEnum_UrlFieldError = 22
	// The final mobile url starts with an invalid protocol.
	UrlFieldErrorEnum_INVALID_PROTOCOL_IN_FINAL_MOBILE_URL UrlFieldErrorEnum_UrlFieldError = 23
	// The final mobile url contains illegal characters.
	UrlFieldErrorEnum_MALFORMED_FINAL_MOBILE_URL UrlFieldErrorEnum_UrlFieldError = 24
	// The final mobile url must contain a host name.
	UrlFieldErrorEnum_MISSING_HOST_IN_FINAL_MOBILE_URL UrlFieldErrorEnum_UrlFieldError = 25
	// The tracking url template has an invalid or missing top level domain
	// extension.
	UrlFieldErrorEnum_INVALID_TLD_IN_FINAL_MOBILE_URL UrlFieldErrorEnum_UrlFieldError = 26
	// The final app url is invalid.
	UrlFieldErrorEnum_INVALID_FINAL_APP_URL UrlFieldErrorEnum_UrlFieldError = 27
	// The final app url contains invalid tag.
	UrlFieldErrorEnum_INVALID_TAG_IN_FINAL_APP_URL UrlFieldErrorEnum_UrlFieldError = 28
	// The final app url contains nested occurrences of the same conditional tag
	// (i.e. {ifmobile:{ifmobile:x}}).
	UrlFieldErrorEnum_REDUNDANT_NESTED_FINAL_APP_URL_TAG UrlFieldErrorEnum_UrlFieldError = 29
	// More than one app url found for the same OS type.
	UrlFieldErrorEnum_MULTIPLE_APP_URLS_FOR_OSTYPE UrlFieldErrorEnum_UrlFieldError = 30
	// The OS type given for an app url is not valid.
	UrlFieldErrorEnum_INVALID_OSTYPE UrlFieldErrorEnum_UrlFieldError = 31
	// The protocol given for an app url is not valid. (E.g. "android-app://")
	UrlFieldErrorEnum_INVALID_PROTOCOL_FOR_APP_URL UrlFieldErrorEnum_UrlFieldError = 32
	// The package id (app id) given for an app url is not valid.
	UrlFieldErrorEnum_INVALID_PACKAGE_ID_FOR_APP_URL UrlFieldErrorEnum_UrlFieldError = 33
	// The number of url custom parameters for an resource exceeds the maximum
	// limit allowed.
	UrlFieldErrorEnum_URL_CUSTOM_PARAMETERS_COUNT_EXCEEDS_LIMIT UrlFieldErrorEnum_UrlFieldError = 34
	// An invalid character appears in the parameter key.
	UrlFieldErrorEnum_INVALID_CHARACTERS_IN_URL_CUSTOM_PARAMETER_KEY UrlFieldErrorEnum_UrlFieldError = 39
	// An invalid character appears in the parameter value.
	UrlFieldErrorEnum_INVALID_CHARACTERS_IN_URL_CUSTOM_PARAMETER_VALUE UrlFieldErrorEnum_UrlFieldError = 40
	// The url custom parameter value fails url tag validation.
	UrlFieldErrorEnum_INVALID_TAG_IN_URL_CUSTOM_PARAMETER_VALUE UrlFieldErrorEnum_UrlFieldError = 41
	// The custom parameter contains nested occurrences of the same conditional
	// tag (i.e. {ifmobile:{ifmobile:x}}).
	UrlFieldErrorEnum_REDUNDANT_NESTED_URL_CUSTOM_PARAMETER_TAG UrlFieldErrorEnum_UrlFieldError = 42
	// The protocol (http:// or https://) is missing.
	UrlFieldErrorEnum_MISSING_PROTOCOL UrlFieldErrorEnum_UrlFieldError = 43
	// Unsupported protocol in URL. Only http and https are supported.
	UrlFieldErrorEnum_INVALID_PROTOCOL UrlFieldErrorEnum_UrlFieldError = 52
	// The url is invalid.
	UrlFieldErrorEnum_INVALID_URL UrlFieldErrorEnum_UrlFieldError = 44
	// Destination Url is deprecated.
	UrlFieldErrorEnum_DESTINATION_URL_DEPRECATED UrlFieldErrorEnum_UrlFieldError = 45
	// The url contains invalid tag.
	UrlFieldErrorEnum_INVALID_TAG_IN_URL UrlFieldErrorEnum_UrlFieldError = 46
	// The url must contain at least one tag (e.g. {lpurl}), This applies only
	// to urls associated with website ads or product ads.
	UrlFieldErrorEnum_MISSING_URL_TAG UrlFieldErrorEnum_UrlFieldError = 47
	// Duplicate url id.
	UrlFieldErrorEnum_DUPLICATE_URL_ID UrlFieldErrorEnum_UrlFieldError = 48
	// Invalid url id.
	UrlFieldErrorEnum_INVALID_URL_ID UrlFieldErrorEnum_UrlFieldError = 49
	// The final url suffix cannot begin with '?' or '&' characters and must be
	// a valid query string.
	UrlFieldErrorEnum_FINAL_URL_SUFFIX_MALFORMED UrlFieldErrorEnum_UrlFieldError = 50
	// The final url suffix cannot contain {lpurl} related or {ignore} tags.
	UrlFieldErrorEnum_INVALID_TAG_IN_FINAL_URL_SUFFIX UrlFieldErrorEnum_UrlFieldError = 51
	// The top level domain is invalid, e.g, not a public top level domain
	// listed in publicsuffix.org.
	UrlFieldErrorEnum_INVALID_TOP_LEVEL_DOMAIN UrlFieldErrorEnum_UrlFieldError = 53
	// Malformed top level domain in URL.
	UrlFieldErrorEnum_MALFORMED_TOP_LEVEL_DOMAIN UrlFieldErrorEnum_UrlFieldError = 54
	// Malformed URL.
	UrlFieldErrorEnum_MALFORMED_URL UrlFieldErrorEnum_UrlFieldError = 55
	// No host found in URL.
	UrlFieldErrorEnum_MISSING_HOST UrlFieldErrorEnum_UrlFieldError = 56
	// Custom parameter value cannot be null.
	UrlFieldErrorEnum_NULL_CUSTOM_PARAMETER_VALUE UrlFieldErrorEnum_UrlFieldError = 57
)

var UrlFieldErrorEnum_UrlFieldError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "INVALID_TRACKING_URL_TEMPLATE",
	3:  "INVALID_TAG_IN_TRACKING_URL_TEMPLATE",
	4:  "MISSING_TRACKING_URL_TEMPLATE_TAG",
	5:  "MISSING_PROTOCOL_IN_TRACKING_URL_TEMPLATE",
	6:  "INVALID_PROTOCOL_IN_TRACKING_URL_TEMPLATE",
	7:  "MALFORMED_TRACKING_URL_TEMPLATE",
	8:  "MISSING_HOST_IN_TRACKING_URL_TEMPLATE",
	9:  "INVALID_TLD_IN_TRACKING_URL_TEMPLATE",
	10: "REDUNDANT_NESTED_TRACKING_URL_TEMPLATE_TAG",
	11: "INVALID_FINAL_URL",
	12: "INVALID_TAG_IN_FINAL_URL",
	13: "REDUNDANT_NESTED_FINAL_URL_TAG",
	14: "MISSING_PROTOCOL_IN_FINAL_URL",
	15: "INVALID_PROTOCOL_IN_FINAL_URL",
	16: "MALFORMED_FINAL_URL",
	17: "MISSING_HOST_IN_FINAL_URL",
	18: "INVALID_TLD_IN_FINAL_URL",
	19: "INVALID_FINAL_MOBILE_URL",
	20: "INVALID_TAG_IN_FINAL_MOBILE_URL",
	21: "REDUNDANT_NESTED_FINAL_MOBILE_URL_TAG",
	22: "MISSING_PROTOCOL_IN_FINAL_MOBILE_URL",
	23: "INVALID_PROTOCOL_IN_FINAL_MOBILE_URL",
	24: "MALFORMED_FINAL_MOBILE_URL",
	25: "MISSING_HOST_IN_FINAL_MOBILE_URL",
	26: "INVALID_TLD_IN_FINAL_MOBILE_URL",
	27: "INVALID_FINAL_APP_URL",
	28: "INVALID_TAG_IN_FINAL_APP_URL",
	29: "REDUNDANT_NESTED_FINAL_APP_URL_TAG",
	30: "MULTIPLE_APP_URLS_FOR_OSTYPE",
	31: "INVALID_OSTYPE",
	32: "INVALID_PROTOCOL_FOR_APP_URL",
	33: "INVALID_PACKAGE_ID_FOR_APP_URL",
	34: "URL_CUSTOM_PARAMETERS_COUNT_EXCEEDS_LIMIT",
	39: "INVALID_CHARACTERS_IN_URL_CUSTOM_PARAMETER_KEY",
	40: "INVALID_CHARACTERS_IN_URL_CUSTOM_PARAMETER_VALUE",
	41: "INVALID_TAG_IN_URL_CUSTOM_PARAMETER_VALUE",
	42: "REDUNDANT_NESTED_URL_CUSTOM_PARAMETER_TAG",
	43: "MISSING_PROTOCOL",
	52: "INVALID_PROTOCOL",
	44: "INVALID_URL",
	45: "DESTINATION_URL_DEPRECATED",
	46: "INVALID_TAG_IN_URL",
	47: "MISSING_URL_TAG",
	48: "DUPLICATE_URL_ID",
	49: "INVALID_URL_ID",
	50: "FINAL_URL_SUFFIX_MALFORMED",
	51: "INVALID_TAG_IN_FINAL_URL_SUFFIX",
	53: "INVALID_TOP_LEVEL_DOMAIN",
	54: "MALFORMED_TOP_LEVEL_DOMAIN",
	55: "MALFORMED_URL",
	56: "MISSING_HOST",
	57: "NULL_CUSTOM_PARAMETER_VALUE",
}

var UrlFieldErrorEnum_UrlFieldError_value = map[string]int32{
	"UNSPECIFIED":                                      0,
	"UNKNOWN":                                          1,
	"INVALID_TRACKING_URL_TEMPLATE":                    2,
	"INVALID_TAG_IN_TRACKING_URL_TEMPLATE":             3,
	"MISSING_TRACKING_URL_TEMPLATE_TAG":                4,
	"MISSING_PROTOCOL_IN_TRACKING_URL_TEMPLATE":        5,
	"INVALID_PROTOCOL_IN_TRACKING_URL_TEMPLATE":        6,
	"MALFORMED_TRACKING_URL_TEMPLATE":                  7,
	"MISSING_HOST_IN_TRACKING_URL_TEMPLATE":            8,
	"INVALID_TLD_IN_TRACKING_URL_TEMPLATE":             9,
	"REDUNDANT_NESTED_TRACKING_URL_TEMPLATE_TAG":       10,
	"INVALID_FINAL_URL":                                11,
	"INVALID_TAG_IN_FINAL_URL":                         12,
	"REDUNDANT_NESTED_FINAL_URL_TAG":                   13,
	"MISSING_PROTOCOL_IN_FINAL_URL":                    14,
	"INVALID_PROTOCOL_IN_FINAL_URL":                    15,
	"MALFORMED_FINAL_URL":                              16,
	"MISSING_HOST_IN_FINAL_URL":                        17,
	"INVALID_TLD_IN_FINAL_URL":                         18,
	"INVALID_FINAL_MOBILE_URL":                         19,
	"INVALID_TAG_IN_FINAL_MOBILE_URL":                  20,
	"REDUNDANT_NESTED_FINAL_MOBILE_URL_TAG":            21,
	"MISSING_PROTOCOL_IN_FINAL_MOBILE_URL":             22,
	"INVALID_PROTOCOL_IN_FINAL_MOBILE_URL":             23,
	"MALFORMED_FINAL_MOBILE_URL":                       24,
	"MISSING_HOST_IN_FINAL_MOBILE_URL":                 25,
	"INVALID_TLD_IN_FINAL_MOBILE_URL":                  26,
	"INVALID_FINAL_APP_URL":                            27,
	"INVALID_TAG_IN_FINAL_APP_URL":                     28,
	"REDUNDANT_NESTED_FINAL_APP_URL_TAG":               29,
	"MULTIPLE_APP_URLS_FOR_OSTYPE":                     30,
	"INVALID_OSTYPE":                                   31,
	"INVALID_PROTOCOL_FOR_APP_URL":                     32,
	"INVALID_PACKAGE_ID_FOR_APP_URL":                   33,
	"URL_CUSTOM_PARAMETERS_COUNT_EXCEEDS_LIMIT":        34,
	"INVALID_CHARACTERS_IN_URL_CUSTOM_PARAMETER_KEY":   39,
	"INVALID_CHARACTERS_IN_URL_CUSTOM_PARAMETER_VALUE": 40,
	"INVALID_TAG_IN_URL_CUSTOM_PARAMETER_VALUE":        41,
	"REDUNDANT_NESTED_URL_CUSTOM_PARAMETER_TAG":        42,
	"MISSING_PROTOCOL":                                 43,
	"INVALID_PROTOCOL":                                 52,
	"INVALID_URL":                                      44,
	"DESTINATION_URL_DEPRECATED":                       45,
	"INVALID_TAG_IN_URL":                               46,
	"MISSING_URL_TAG":                                  47,
	"DUPLICATE_URL_ID":                                 48,
	"INVALID_URL_ID":                                   49,
	"FINAL_URL_SUFFIX_MALFORMED":                       50,
	"INVALID_TAG_IN_FINAL_URL_SUFFIX":                  51,
	"INVALID_TOP_LEVEL_DOMAIN":                         53,
	"MALFORMED_TOP_LEVEL_DOMAIN":                       54,
	"MALFORMED_URL":                                    55,
	"MISSING_HOST":                                     56,
	"NULL_CUSTOM_PARAMETER_VALUE":                      57,
}

func (x UrlFieldErrorEnum_UrlFieldError) String() string {
	return proto.EnumName(UrlFieldErrorEnum_UrlFieldError_name, int32(x))
}

func (UrlFieldErrorEnum_UrlFieldError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fda0f9682a3a1a06, []int{0, 0}
}

// Container for enum describing possible url field errors.
type UrlFieldErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UrlFieldErrorEnum) Reset()         { *m = UrlFieldErrorEnum{} }
func (m *UrlFieldErrorEnum) String() string { return proto.CompactTextString(m) }
func (*UrlFieldErrorEnum) ProtoMessage()    {}
func (*UrlFieldErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_fda0f9682a3a1a06, []int{0}
}

func (m *UrlFieldErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UrlFieldErrorEnum.Unmarshal(m, b)
}
func (m *UrlFieldErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UrlFieldErrorEnum.Marshal(b, m, deterministic)
}
func (m *UrlFieldErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UrlFieldErrorEnum.Merge(m, src)
}
func (m *UrlFieldErrorEnum) XXX_Size() int {
	return xxx_messageInfo_UrlFieldErrorEnum.Size(m)
}
func (m *UrlFieldErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UrlFieldErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UrlFieldErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.UrlFieldErrorEnum_UrlFieldError", UrlFieldErrorEnum_UrlFieldError_name, UrlFieldErrorEnum_UrlFieldError_value)
	proto.RegisterType((*UrlFieldErrorEnum)(nil), "google.ads.googleads.v3.errors.UrlFieldErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/url_field_error.proto", fileDescriptor_fda0f9682a3a1a06)
}

var fileDescriptor_fda0f9682a3a1a06 = []byte{
	// 865 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xc7, 0x97, 0x74, 0x6b, 0xb7, 0x93, 0x26, 0x61, 0x4e, 0x9a, 0xb6, 0x49, 0x13, 0x27, 0xf1,
	0xda, 0xad, 0xe9, 0x56, 0xb9, 0xab, 0xb3, 0x2f, 0xef, 0x8a, 0x91, 0x68, 0x97, 0x08, 0x45, 0x09,
	0x12, 0xe5, 0xb5, 0x43, 0x00, 0xc2, 0x9b, 0x33, 0x23, 0x80, 0x6b, 0x05, 0x56, 0xda, 0xe7, 0x19,
	0x76, 0xb9, 0x07, 0xd8, 0x43, 0xec, 0x51, 0x06, 0xec, 0x1d, 0x06, 0x4a, 0x96, 0x25, 0x3b, 0x56,
	0x86, 0x5d, 0x99, 0x38, 0xe7, 0x77, 0x3e, 0xf8, 0x3f, 0x22, 0x69, 0x38, 0x1e, 0xc4, 0xf1, 0x60,
	0x78, 0xde, 0xe8, 0xf5, 0x93, 0x46, 0xb6, 0x34, 0xab, 0xf7, 0xcd, 0xc6, 0xf9, 0x78, 0x1c, 0x8f,
	0x93, 0xc6, 0xbb, 0xf1, 0x50, 0xff, 0x7a, 0x71, 0x3e, 0xec, 0xeb, 0xd4, 0x60, 0x5d, 0x8e, 0xe3,
	0xab, 0x18, 0x6b, 0x19, 0x6a, 0xf5, 0xfa, 0x89, 0x35, 0x8d, 0xb2, 0xde, 0x37, 0xad, 0x2c, 0x6a,
	0x67, 0x37, 0xcf, 0x7a, 0x79, 0xd1, 0xe8, 0x8d, 0x46, 0xf1, 0x55, 0xef, 0xea, 0x22, 0x1e, 0x25,
	0x59, 0x74, 0xfd, 0xcf, 0x35, 0xd8, 0x88, 0xc6, 0xc3, 0xb6, 0x49, 0xcb, 0x4c, 0x00, 0x1b, 0xbd,
	0x7b, 0x5b, 0xff, 0x6d, 0x0d, 0x56, 0x67, 0xac, 0xb8, 0x0e, 0x2b, 0x91, 0x0c, 0x7d, 0x66, 0xf3,
	0x36, 0x67, 0x0e, 0xf9, 0x00, 0x57, 0xe0, 0x4e, 0x24, 0x4f, 0xa5, 0xf7, 0xa3, 0x24, 0x4b, 0x78,
	0x08, 0x7b, 0x5c, 0x76, 0xa9, 0xe0, 0x8e, 0x56, 0x01, 0xb5, 0x4f, 0xb9, 0xec, 0xe8, 0x28, 0x10,
	0x5a, 0x31, 0xd7, 0x17, 0x54, 0x31, 0xb2, 0x8c, 0x4f, 0xe1, 0xf1, 0x14, 0xa1, 0x1d, 0xcd, 0x65,
	0x05, 0x79, 0x0b, 0x9f, 0xc0, 0xa1, 0xcb, 0xc3, 0xd0, 0x78, 0x16, 0x22, 0x26, 0x9e, 0x7c, 0x88,
	0xcf, 0xe1, 0x28, 0xc7, 0xfc, 0xc0, 0x53, 0x9e, 0xed, 0x89, 0xea, 0xac, 0x1f, 0x19, 0x3c, 0xaf,
	0xff, 0xdf, 0xf8, 0x6d, 0xfc, 0x14, 0xf6, 0x5d, 0x2a, 0xda, 0x5e, 0xe0, 0xb2, 0xaa, 0x3d, 0xdd,
	0xc1, 0x23, 0x78, 0x92, 0xb7, 0xf0, 0xca, 0x0b, 0x55, 0x75, 0xbe, 0x8f, 0x67, 0xb6, 0x2f, 0x9c,
	0x6a, 0xf2, 0x13, 0xb4, 0xe0, 0x59, 0xc0, 0x9c, 0x48, 0x3a, 0x54, 0x2a, 0x2d, 0x59, 0xa8, 0xaa,
	0x1a, 0x48, 0x75, 0x00, 0xdc, 0x82, 0x8d, 0x3c, 0x73, 0x9b, 0x4b, 0x2a, 0x0c, 0x43, 0x56, 0x70,
	0x17, 0x1e, 0xce, 0xe9, 0x5d, 0x78, 0xef, 0x62, 0x1d, 0x6a, 0xd7, 0x8a, 0x4c, 0xfd, 0x69, 0xe2,
	0x55, 0x33, 0xd4, 0x45, 0x02, 0x17, 0x69, 0xd6, 0xca, 0x73, 0x5f, 0x8c, 0xac, 0xe3, 0x03, 0xd8,
	0x2c, 0x84, 0x2c, 0x1c, 0x04, 0xf7, 0x60, 0x7b, 0x5e, 0xbc, 0xc2, 0xbd, 0x31, 0xd3, 0x7f, 0x26,
	0x58, 0xe1, 0xc5, 0xb2, 0x37, 0x33, 0xbb, 0xde, 0x09, 0x17, 0x2c, 0xf5, 0x6e, 0x9a, 0xe1, 0x2d,
	0xdc, 0x7b, 0x09, 0xba, 0x67, 0x86, 0x57, 0x21, 0x41, 0x81, 0xa5, 0x4a, 0x6c, 0x99, 0xe1, 0x55,
	0x2b, 0x51, 0x4a, 0x7a, 0xbf, 0x3c, 0xe6, 0x1b, 0xc9, 0x07, 0x58, 0x83, 0x9d, 0x79, 0x5d, 0x4a,
	0xfe, 0x87, 0xf8, 0x18, 0x0e, 0x16, 0xcb, 0x53, 0xa2, 0xb6, 0x67, 0x76, 0x5a, 0x56, 0xa9, 0x04,
	0xed, 0xe0, 0x36, 0x6c, 0xcd, 0x8a, 0x45, 0x7d, 0x3f, 0x75, 0x3d, 0xc2, 0x03, 0xd8, 0x5d, 0xa8,
	0x54, 0x4e, 0xec, 0xe2, 0x67, 0x50, 0xaf, 0x90, 0x69, 0xc2, 0xa4, 0x1a, 0xed, 0x99, 0x4c, 0x6e,
	0x24, 0x14, 0xf7, 0x05, 0xcb, 0x3d, 0xa1, 0x6e, 0x7b, 0x81, 0xf6, 0x42, 0xf5, 0xc6, 0x67, 0xa4,
	0x86, 0x08, 0x6b, 0x79, 0xad, 0x89, 0x6d, 0xbf, 0x5c, 0x7f, 0xaa, 0x97, 0x09, 0xca, 0xeb, 0x1f,
	0x98, 0x2f, 0x75, 0x4a, 0x50, 0xfb, 0x94, 0x76, 0x98, 0x36, 0xfb, 0x28, 0x31, 0x87, 0xe6, 0x6c,
	0x9b, 0x46, 0xec, 0x28, 0x54, 0x9e, 0xab, 0x7d, 0x1a, 0x50, 0x97, 0x29, 0x16, 0x84, 0xda, 0xf6,
	0x22, 0xa9, 0x34, 0x7b, 0x6d, 0x33, 0xe6, 0x84, 0x5a, 0x70, 0x97, 0x2b, 0x52, 0xc7, 0x97, 0x60,
	0xe5, 0x29, 0xed, 0x57, 0x34, 0xa0, 0x76, 0xca, 0x72, 0xa9, 0x17, 0x25, 0xd1, 0xa7, 0xec, 0x0d,
	0xf9, 0x1c, 0x8f, 0xe1, 0xc5, 0xff, 0x88, 0xe9, 0x52, 0x11, 0x31, 0xf2, 0xb4, 0x7c, 0xe9, 0x4c,
	0xe4, 0xbd, 0x01, 0x3f, 0x32, 0xf8, 0x35, 0xad, 0x17, 0x06, 0x18, 0xc9, 0x9f, 0xe1, 0x3d, 0x20,
	0xf3, 0x9f, 0x25, 0xf9, 0xc2, 0x58, 0xe7, 0x25, 0x25, 0xc7, 0xe6, 0xfe, 0xce, 0xad, 0x46, 0xb3,
	0x2f, 0xcd, 0xf7, 0xe7, 0xb0, 0x50, 0x71, 0x49, 0x15, 0xf7, 0xb2, 0xbe, 0x1c, 0xe6, 0x07, 0xcc,
	0xa6, 0x8a, 0x39, 0xe4, 0x39, 0xde, 0x07, 0xbc, 0xde, 0x3a, 0xb1, 0x70, 0x13, 0xd6, 0xf3, 0xa2,
	0xf9, 0xf0, 0x1b, 0xa6, 0xa6, 0x13, 0xf9, 0x82, 0x9b, 0xe0, 0xd4, 0xcc, 0x1d, 0xf2, 0xa2, 0x3c,
	0xf0, 0x89, 0xed, 0x2b, 0x53, 0xb6, 0xb8, 0x67, 0xc2, 0xa8, 0xdd, 0xe6, 0xaf, 0xf5, 0xf4, 0x1c,
	0x90, 0x97, 0x95, 0x47, 0xb7, 0xc0, 0x49, 0x73, 0xe6, 0x6e, 0xf0, 0x7c, 0x2d, 0x58, 0x97, 0x09,
	0xed, 0x78, 0x2e, 0xe5, 0x92, 0x7c, 0x3d, 0x7b, 0xb2, 0xae, 0xf9, 0xbf, 0xc1, 0x0d, 0x58, 0x2d,
	0xfc, 0x66, 0x53, 0xdf, 0x22, 0x81, 0xbb, 0xe5, 0xc3, 0x46, 0xbe, 0xc3, 0x7d, 0x78, 0x24, 0x23,
	0x51, 0x39, 0xab, 0xef, 0x4f, 0xfe, 0x59, 0x82, 0xfa, 0x2f, 0xf1, 0x5b, 0xeb, 0xe6, 0xd7, 0xf7,
	0x04, 0x67, 0x9e, 0x51, 0xdf, 0xbc, 0xb9, 0xfe, 0xd2, 0x4f, 0xce, 0x24, 0x6a, 0x10, 0x0f, 0x7b,
	0xa3, 0x81, 0x15, 0x8f, 0x07, 0x8d, 0xc1, 0xf9, 0x28, 0x7d, 0x91, 0xf3, 0x97, 0xff, 0xf2, 0x22,
	0xa9, 0xfa, 0x23, 0xf0, 0x43, 0xf6, 0xf3, 0xfb, 0xf2, 0xad, 0x0e, 0xa5, 0x7f, 0x2c, 0xd7, 0x3a,
	0x59, 0x32, 0xda, 0x4f, 0xac, 0x6c, 0x69, 0x56, 0xdd, 0xa6, 0x95, 0x96, 0x4c, 0xfe, 0xca, 0x81,
	0x33, 0xda, 0x4f, 0xce, 0xa6, 0xc0, 0x59, 0xb7, 0x79, 0x96, 0x01, 0x7f, 0x2f, 0xd7, 0x33, 0x6b,
	0xab, 0x45, 0xfb, 0x49, 0xab, 0x35, 0x45, 0x5a, 0xad, 0x6e, 0xb3, 0xd5, 0xca, 0xa0, 0x9f, 0x6f,
	0xa7, 0xdd, 0x35, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xcf, 0x6b, 0x07, 0xa5, 0x08, 0x00,
	0x00,
}
