// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/price_placeholder_field.proto

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

// Possible values for Price placeholder fields.
type PricePlaceholderFieldEnum_PricePlaceholderField int32

const (
	// Not specified.
	PricePlaceholderFieldEnum_UNSPECIFIED PricePlaceholderFieldEnum_PricePlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	PricePlaceholderFieldEnum_UNKNOWN PricePlaceholderFieldEnum_PricePlaceholderField = 1
	// Data Type: STRING. The type of your price feed. Must match one of the
	// predefined price feed type exactly.
	PricePlaceholderFieldEnum_TYPE PricePlaceholderFieldEnum_PricePlaceholderField = 2
	// Data Type: STRING. The qualifier of each price. Must match one of the
	// predefined price qualifiers exactly.
	PricePlaceholderFieldEnum_PRICE_QUALIFIER PricePlaceholderFieldEnum_PricePlaceholderField = 3
	// Data Type: URL. Tracking template for the price feed when using Upgraded
	// URLs.
	PricePlaceholderFieldEnum_TRACKING_TEMPLATE PricePlaceholderFieldEnum_PricePlaceholderField = 4
	// Data Type: STRING. Language of the price feed. Must match one of the
	// available available locale codes exactly.
	PricePlaceholderFieldEnum_LANGUAGE PricePlaceholderFieldEnum_PricePlaceholderField = 5
	// Data Type: STRING. Final URL suffix for the price feed when using
	// parallel tracking.
	PricePlaceholderFieldEnum_FINAL_URL_SUFFIX PricePlaceholderFieldEnum_PricePlaceholderField = 6
	// Data Type: STRING. The header of item 1 of the table.
	PricePlaceholderFieldEnum_ITEM_1_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 100
	// Data Type: STRING. The description of item 1 of the table.
	PricePlaceholderFieldEnum_ITEM_1_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 101
	// Data Type: MONEY. The price (money with currency) of item 1 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_1_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 102
	// Data Type: STRING. The price unit of item 1 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_1_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 103
	// Data Type: URL_LIST. The final URLs of item 1 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_1_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 104
	// Data Type: URL_LIST. The final mobile URLs of item 1 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_1_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 105
	// Data Type: STRING. The header of item 2 of the table.
	PricePlaceholderFieldEnum_ITEM_2_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 200
	// Data Type: STRING. The description of item 2 of the table.
	PricePlaceholderFieldEnum_ITEM_2_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 201
	// Data Type: MONEY. The price (money with currency) of item 2 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_2_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 202
	// Data Type: STRING. The price unit of item 2 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_2_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 203
	// Data Type: URL_LIST. The final URLs of item 2 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_2_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 204
	// Data Type: URL_LIST. The final mobile URLs of item 2 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_2_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 205
	// Data Type: STRING. The header of item 3 of the table.
	PricePlaceholderFieldEnum_ITEM_3_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 300
	// Data Type: STRING. The description of item 3 of the table.
	PricePlaceholderFieldEnum_ITEM_3_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 301
	// Data Type: MONEY. The price (money with currency) of item 3 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_3_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 302
	// Data Type: STRING. The price unit of item 3 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_3_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 303
	// Data Type: URL_LIST. The final URLs of item 3 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_3_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 304
	// Data Type: URL_LIST. The final mobile URLs of item 3 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_3_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 305
	// Data Type: STRING. The header of item 4 of the table.
	PricePlaceholderFieldEnum_ITEM_4_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 400
	// Data Type: STRING. The description of item 4 of the table.
	PricePlaceholderFieldEnum_ITEM_4_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 401
	// Data Type: MONEY. The price (money with currency) of item 4 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_4_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 402
	// Data Type: STRING. The price unit of item 4 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_4_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 403
	// Data Type: URL_LIST. The final URLs of item 4 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_4_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 404
	// Data Type: URL_LIST. The final mobile URLs of item 4 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_4_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 405
	// Data Type: STRING. The header of item 5 of the table.
	PricePlaceholderFieldEnum_ITEM_5_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 500
	// Data Type: STRING. The description of item 5 of the table.
	PricePlaceholderFieldEnum_ITEM_5_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 501
	// Data Type: MONEY. The price (money with currency) of item 5 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_5_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 502
	// Data Type: STRING. The price unit of item 5 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_5_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 503
	// Data Type: URL_LIST. The final URLs of item 5 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_5_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 504
	// Data Type: URL_LIST. The final mobile URLs of item 5 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_5_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 505
	// Data Type: STRING. The header of item 6 of the table.
	PricePlaceholderFieldEnum_ITEM_6_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 600
	// Data Type: STRING. The description of item 6 of the table.
	PricePlaceholderFieldEnum_ITEM_6_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 601
	// Data Type: MONEY. The price (money with currency) of item 6 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_6_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 602
	// Data Type: STRING. The price unit of item 6 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_6_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 603
	// Data Type: URL_LIST. The final URLs of item 6 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_6_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 604
	// Data Type: URL_LIST. The final mobile URLs of item 6 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_6_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 605
	// Data Type: STRING. The header of item 7 of the table.
	PricePlaceholderFieldEnum_ITEM_7_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 700
	// Data Type: STRING. The description of item 7 of the table.
	PricePlaceholderFieldEnum_ITEM_7_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 701
	// Data Type: MONEY. The price (money with currency) of item 7 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_7_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 702
	// Data Type: STRING. The price unit of item 7 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_7_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 703
	// Data Type: URL_LIST. The final URLs of item 7 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_7_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 704
	// Data Type: URL_LIST. The final mobile URLs of item 7 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_7_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 705
	// Data Type: STRING. The header of item 8 of the table.
	PricePlaceholderFieldEnum_ITEM_8_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 800
	// Data Type: STRING. The description of item 8 of the table.
	PricePlaceholderFieldEnum_ITEM_8_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 801
	// Data Type: MONEY. The price (money with currency) of item 8 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_8_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 802
	// Data Type: STRING. The price unit of item 8 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_8_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 803
	// Data Type: URL_LIST. The final URLs of item 8 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_8_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 804
	// Data Type: URL_LIST. The final mobile URLs of item 8 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_8_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 805
)

var PricePlaceholderFieldEnum_PricePlaceholderField_name = map[int32]string{
	0:   "UNSPECIFIED",
	1:   "UNKNOWN",
	2:   "TYPE",
	3:   "PRICE_QUALIFIER",
	4:   "TRACKING_TEMPLATE",
	5:   "LANGUAGE",
	6:   "FINAL_URL_SUFFIX",
	100: "ITEM_1_HEADER",
	101: "ITEM_1_DESCRIPTION",
	102: "ITEM_1_PRICE",
	103: "ITEM_1_UNIT",
	104: "ITEM_1_FINAL_URLS",
	105: "ITEM_1_FINAL_MOBILE_URLS",
	200: "ITEM_2_HEADER",
	201: "ITEM_2_DESCRIPTION",
	202: "ITEM_2_PRICE",
	203: "ITEM_2_UNIT",
	204: "ITEM_2_FINAL_URLS",
	205: "ITEM_2_FINAL_MOBILE_URLS",
	300: "ITEM_3_HEADER",
	301: "ITEM_3_DESCRIPTION",
	302: "ITEM_3_PRICE",
	303: "ITEM_3_UNIT",
	304: "ITEM_3_FINAL_URLS",
	305: "ITEM_3_FINAL_MOBILE_URLS",
	400: "ITEM_4_HEADER",
	401: "ITEM_4_DESCRIPTION",
	402: "ITEM_4_PRICE",
	403: "ITEM_4_UNIT",
	404: "ITEM_4_FINAL_URLS",
	405: "ITEM_4_FINAL_MOBILE_URLS",
	500: "ITEM_5_HEADER",
	501: "ITEM_5_DESCRIPTION",
	502: "ITEM_5_PRICE",
	503: "ITEM_5_UNIT",
	504: "ITEM_5_FINAL_URLS",
	505: "ITEM_5_FINAL_MOBILE_URLS",
	600: "ITEM_6_HEADER",
	601: "ITEM_6_DESCRIPTION",
	602: "ITEM_6_PRICE",
	603: "ITEM_6_UNIT",
	604: "ITEM_6_FINAL_URLS",
	605: "ITEM_6_FINAL_MOBILE_URLS",
	700: "ITEM_7_HEADER",
	701: "ITEM_7_DESCRIPTION",
	702: "ITEM_7_PRICE",
	703: "ITEM_7_UNIT",
	704: "ITEM_7_FINAL_URLS",
	705: "ITEM_7_FINAL_MOBILE_URLS",
	800: "ITEM_8_HEADER",
	801: "ITEM_8_DESCRIPTION",
	802: "ITEM_8_PRICE",
	803: "ITEM_8_UNIT",
	804: "ITEM_8_FINAL_URLS",
	805: "ITEM_8_FINAL_MOBILE_URLS",
}

var PricePlaceholderFieldEnum_PricePlaceholderField_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"TYPE":                     2,
	"PRICE_QUALIFIER":          3,
	"TRACKING_TEMPLATE":        4,
	"LANGUAGE":                 5,
	"FINAL_URL_SUFFIX":         6,
	"ITEM_1_HEADER":            100,
	"ITEM_1_DESCRIPTION":       101,
	"ITEM_1_PRICE":             102,
	"ITEM_1_UNIT":              103,
	"ITEM_1_FINAL_URLS":        104,
	"ITEM_1_FINAL_MOBILE_URLS": 105,
	"ITEM_2_HEADER":            200,
	"ITEM_2_DESCRIPTION":       201,
	"ITEM_2_PRICE":             202,
	"ITEM_2_UNIT":              203,
	"ITEM_2_FINAL_URLS":        204,
	"ITEM_2_FINAL_MOBILE_URLS": 205,
	"ITEM_3_HEADER":            300,
	"ITEM_3_DESCRIPTION":       301,
	"ITEM_3_PRICE":             302,
	"ITEM_3_UNIT":              303,
	"ITEM_3_FINAL_URLS":        304,
	"ITEM_3_FINAL_MOBILE_URLS": 305,
	"ITEM_4_HEADER":            400,
	"ITEM_4_DESCRIPTION":       401,
	"ITEM_4_PRICE":             402,
	"ITEM_4_UNIT":              403,
	"ITEM_4_FINAL_URLS":        404,
	"ITEM_4_FINAL_MOBILE_URLS": 405,
	"ITEM_5_HEADER":            500,
	"ITEM_5_DESCRIPTION":       501,
	"ITEM_5_PRICE":             502,
	"ITEM_5_UNIT":              503,
	"ITEM_5_FINAL_URLS":        504,
	"ITEM_5_FINAL_MOBILE_URLS": 505,
	"ITEM_6_HEADER":            600,
	"ITEM_6_DESCRIPTION":       601,
	"ITEM_6_PRICE":             602,
	"ITEM_6_UNIT":              603,
	"ITEM_6_FINAL_URLS":        604,
	"ITEM_6_FINAL_MOBILE_URLS": 605,
	"ITEM_7_HEADER":            700,
	"ITEM_7_DESCRIPTION":       701,
	"ITEM_7_PRICE":             702,
	"ITEM_7_UNIT":              703,
	"ITEM_7_FINAL_URLS":        704,
	"ITEM_7_FINAL_MOBILE_URLS": 705,
	"ITEM_8_HEADER":            800,
	"ITEM_8_DESCRIPTION":       801,
	"ITEM_8_PRICE":             802,
	"ITEM_8_UNIT":              803,
	"ITEM_8_FINAL_URLS":        804,
	"ITEM_8_FINAL_MOBILE_URLS": 805,
}

func (x PricePlaceholderFieldEnum_PricePlaceholderField) String() string {
	return proto.EnumName(PricePlaceholderFieldEnum_PricePlaceholderField_name, int32(x))
}

func (PricePlaceholderFieldEnum_PricePlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4bcb1bf31a53d818, []int{0, 0}
}

// Values for Price placeholder fields.
type PricePlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PricePlaceholderFieldEnum) Reset()         { *m = PricePlaceholderFieldEnum{} }
func (m *PricePlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*PricePlaceholderFieldEnum) ProtoMessage()    {}
func (*PricePlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bcb1bf31a53d818, []int{0}
}

func (m *PricePlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PricePlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *PricePlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PricePlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (m *PricePlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PricePlaceholderFieldEnum.Merge(m, src)
}
func (m *PricePlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_PricePlaceholderFieldEnum.Size(m)
}
func (m *PricePlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PricePlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PricePlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.PricePlaceholderFieldEnum_PricePlaceholderField", PricePlaceholderFieldEnum_PricePlaceholderField_name, PricePlaceholderFieldEnum_PricePlaceholderField_value)
	proto.RegisterType((*PricePlaceholderFieldEnum)(nil), "google.ads.googleads.v3.enums.PricePlaceholderFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/price_placeholder_field.proto", fileDescriptor_4bcb1bf31a53d818)
}

var fileDescriptor_4bcb1bf31a53d818 = []byte{
	// 728 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd4, 0x49, 0x6f, 0xd3, 0x50,
	0x10, 0x07, 0x70, 0xbc, 0xc4, 0xb4, 0xaf, 0x45, 0x7d, 0x7d, 0xd0, 0x02, 0x55, 0x7b, 0x68, 0x3f,
	0x80, 0x23, 0xe2, 0x6c, 0x72, 0x4f, 0x4e, 0xea, 0x04, 0xab, 0xa9, 0x6b, 0xb2, 0x94, 0x45, 0x91,
	0xac, 0x50, 0xbb, 0x69, 0xa4, 0x34, 0x8e, 0xe2, 0xb6, 0x9f, 0x83, 0xf5, 0x0e, 0x05, 0x2e, 0x88,
	0xf5, 0xce, 0x7a, 0x83, 0x02, 0x07, 0x6e, 0xac, 0x5f, 0x80, 0x4b, 0x2f, 0xac, 0x37, 0xe4, 0xf8,
	0x79, 0x12, 0x4b, 0x86, 0x8b, 0x35, 0x9a, 0x79, 0x33, 0xfe, 0x9d, 0xfe, 0x68, 0xb1, 0xe9, 0x38,
	0xcd, 0xb6, 0x1d, 0x6f, 0x58, 0x6e, 0xdc, 0x2f, 0xbd, 0x6a, 0x57, 0x8a, 0xdb, 0x9d, 0x9d, 0x2d,
	0x37, 0xde, 0xed, 0xb5, 0xd6, 0x6d, 0xb3, 0xdb, 0x6e, 0xac, 0xdb, 0x9b, 0x4e, 0xdb, 0xb2, 0x7b,
	0xe6, 0x46, 0xcb, 0x6e, 0x5b, 0x62, 0xb7, 0xe7, 0x6c, 0x3b, 0x64, 0xce, 0xdf, 0x10, 0x1b, 0x96,
	0x2b, 0xc2, 0xb2, 0xb8, 0x2b, 0x89, 0xfd, 0xe5, 0x99, 0xd9, 0xe0, 0x76, 0xb7, 0x15, 0x6f, 0x74,
	0x3a, 0xce, 0x76, 0x63, 0xbb, 0xe5, 0x74, 0x5c, 0x7f, 0x79, 0xe1, 0x60, 0x14, 0x9d, 0x34, 0xbc,
	0xf3, 0xc6, 0xe0, 0x7a, 0xc1, 0x3b, 0xae, 0x76, 0x76, 0xb6, 0x16, 0xde, 0x8f, 0xa2, 0xa9, 0xc8,
	0x29, 0x99, 0x40, 0x63, 0x35, 0xbd, 0x62, 0xa8, 0x79, 0xad, 0xa0, 0xa9, 0x4b, 0xf8, 0x10, 0x19,
	0x43, 0x87, 0x6b, 0xfa, 0xb2, 0xbe, 0x7a, 0x56, 0xc7, 0x0c, 0x19, 0x41, 0x7c, 0xf5, 0xbc, 0xa1,
	0x62, 0x96, 0x1c, 0x45, 0x13, 0x46, 0x59, 0xcb, 0xab, 0xe6, 0x99, 0x9a, 0x52, 0xf2, 0xde, 0x96,
	0x31, 0x47, 0xa6, 0xd0, 0x64, 0xb5, 0xac, 0xe4, 0x97, 0x35, 0xbd, 0x68, 0x56, 0xd5, 0x15, 0xa3,
	0xa4, 0x54, 0x55, 0xcc, 0x93, 0x71, 0x34, 0x52, 0x52, 0xf4, 0x62, 0x4d, 0x29, 0xaa, 0x38, 0x46,
	0x8e, 0x21, 0x5c, 0xd0, 0x74, 0xa5, 0x64, 0xd6, 0xca, 0x25, 0xb3, 0x52, 0x2b, 0x14, 0xb4, 0x73,
	0x58, 0x20, 0x93, 0xe8, 0x88, 0x56, 0x55, 0x57, 0xcc, 0x53, 0xe6, 0x69, 0x55, 0x59, 0x52, 0xcb,
	0xd8, 0x22, 0xd3, 0x88, 0xd0, 0xd6, 0x92, 0x5a, 0xc9, 0x97, 0x35, 0xa3, 0xaa, 0xad, 0xea, 0xd8,
	0x26, 0x18, 0x8d, 0xd3, 0x7e, 0x5f, 0x80, 0x37, 0x3c, 0x34, 0xed, 0xd4, 0x74, 0xad, 0x8a, 0x9b,
	0x1e, 0x84, 0x36, 0xe0, 0x57, 0x15, 0xbc, 0x49, 0x66, 0xd1, 0x89, 0x50, 0x7b, 0x65, 0x35, 0xa7,
	0x95, 0x54, 0x7f, 0xda, 0x22, 0x84, 0x12, 0x12, 0x01, 0xe1, 0x15, 0x43, 0x8e, 0x53, 0x43, 0x22,
	0x64, 0x78, 0xcd, 0x90, 0x49, 0x8a, 0x48, 0x50, 0xc4, 0x3e, 0x43, 0x30, 0x55, 0x24, 0x7c, 0xc5,
	0x1b, 0x86, 0x4c, 0x53, 0x46, 0x62, 0x98, 0xf1, 0x96, 0x21, 0x73, 0xd4, 0x91, 0x88, 0x70, 0xbc,
	0x63, 0x00, 0x22, 0x05, 0x90, 0x7b, 0x2c, 0x40, 0xa4, 0x10, 0xe4, 0x3e, 0x0b, 0x10, 0x89, 0x42,
	0x1e, 0xb0, 0x00, 0x91, 0x7c, 0xc8, 0x43, 0x16, 0x20, 0xd2, 0x30, 0xe4, 0x11, 0x0b, 0x10, 0x29,
	0x02, 0xf2, 0x98, 0x05, 0x48, 0x32, 0x80, 0x5c, 0xe2, 0x00, 0x92, 0x0c, 0x41, 0x2e, 0x73, 0x00,
	0x49, 0x52, 0xc8, 0x15, 0x0e, 0x20, 0x49, 0x1f, 0x72, 0x95, 0x03, 0x48, 0x72, 0x18, 0x72, 0x8d,
	0x03, 0x48, 0x32, 0x02, 0x72, 0x9d, 0x03, 0x48, 0x2a, 0x80, 0x7c, 0x1f, 0x40, 0x52, 0x21, 0xc8,
	0x8f, 0x01, 0x24, 0x45, 0x21, 0x3f, 0x07, 0x90, 0x94, 0x0f, 0xf9, 0x35, 0x80, 0xa4, 0x86, 0x21,
	0xbf, 0x07, 0x90, 0x54, 0x04, 0xe4, 0xcf, 0x00, 0x92, 0x0e, 0x20, 0x1f, 0x78, 0x80, 0xa4, 0x43,
	0x90, 0x8f, 0x3c, 0x40, 0xd2, 0x14, 0xf2, 0x89, 0x07, 0x48, 0xda, 0x87, 0x7c, 0xe6, 0x01, 0x92,
	0x1e, 0x86, 0x7c, 0xe1, 0x01, 0x92, 0x8e, 0x80, 0x7c, 0xe5, 0x01, 0x92, 0x09, 0x20, 0x4f, 0x62,
	0x00, 0xc9, 0x84, 0x20, 0x4f, 0x63, 0x00, 0xc9, 0x50, 0xc8, 0xb3, 0x18, 0x40, 0x32, 0x3e, 0xe4,
	0x79, 0x0c, 0x20, 0x99, 0x61, 0xc8, 0x8b, 0x18, 0x40, 0x32, 0x11, 0x90, 0x97, 0x31, 0x80, 0x64,
	0x03, 0xc8, 0x0d, 0x01, 0x20, 0xd9, 0x10, 0xe4, 0xa6, 0x00, 0x90, 0x2c, 0x85, 0xec, 0x09, 0x00,
	0xc9, 0xfa, 0x90, 0x5b, 0x02, 0x40, 0xb2, 0xc3, 0x90, 0xdb, 0x02, 0x40, 0xb2, 0x11, 0x90, 0x3b,
	0x42, 0xee, 0x80, 0x41, 0xf3, 0xeb, 0xce, 0x96, 0xf8, 0xdf, 0xd4, 0xcc, 0xcd, 0x44, 0xc6, 0x9e,
	0xe1, 0x65, 0xa6, 0xc1, 0x5c, 0xc8, 0xd1, 0xe5, 0xa6, 0xd3, 0x6e, 0x74, 0x9a, 0xa2, 0xd3, 0x6b,
	0xc6, 0x9b, 0x76, 0xa7, 0x9f, 0xa8, 0x41, 0x7e, 0x77, 0x5b, 0xee, 0x3f, 0xe2, 0x7c, 0xb1, 0xff,
	0xdd, 0x63, 0xb9, 0xa2, 0xa2, 0xdc, 0x65, 0xe7, 0x8a, 0xfe, 0x29, 0xc5, 0x72, 0x45, 0xbf, 0xf4,
	0xaa, 0x35, 0x49, 0xf4, 0x02, 0xd8, 0xdd, 0x0f, 0xe6, 0x75, 0xc5, 0x72, 0xeb, 0x30, 0xaf, 0xaf,
	0x49, 0xf5, 0xfe, 0xfc, 0x1b, 0x3b, 0xef, 0x37, 0x65, 0x59, 0xb1, 0x5c, 0x59, 0x86, 0x17, 0xb2,
	0xbc, 0x26, 0xc9, 0x72, 0xff, 0xcd, 0x45, 0xa1, 0x0f, 0x93, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x7b, 0x39, 0x81, 0xe3, 0x66, 0x06, 0x00, 0x00,
}
