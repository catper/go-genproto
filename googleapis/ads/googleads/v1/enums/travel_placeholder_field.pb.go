// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/travel_placeholder_field.proto

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

// Possible values for Travel placeholder fields.
type TravelPlaceholderFieldEnum_TravelPlaceholderField int32

const (
	// Not specified.
	TravelPlaceholderFieldEnum_UNSPECIFIED TravelPlaceholderFieldEnum_TravelPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	TravelPlaceholderFieldEnum_UNKNOWN TravelPlaceholderFieldEnum_TravelPlaceholderField = 1
	// Data Type: STRING. Required. Destination id. Example: PAR, LON.
	// For feed items that only have destination id, destination id must be a
	// unique key. For feed items that have both destination id and origin id,
	// then the combination must be a unique key.
	TravelPlaceholderFieldEnum_DESTINATION_ID TravelPlaceholderFieldEnum_TravelPlaceholderField = 2
	// Data Type: STRING. Origin id. Example: PAR, LON.
	// Combination of DESTINATION_ID and ORIGIN_ID must be
	// unique per offer.
	TravelPlaceholderFieldEnum_ORIGIN_ID TravelPlaceholderFieldEnum_TravelPlaceholderField = 3
	// Data Type: STRING. Required. Main headline with name to be shown in
	// dynamic ad.
	TravelPlaceholderFieldEnum_TITLE TravelPlaceholderFieldEnum_TravelPlaceholderField = 4
	// Data Type: STRING. The destination name. Shorter names are recommended.
	TravelPlaceholderFieldEnum_DESTINATION_NAME TravelPlaceholderFieldEnum_TravelPlaceholderField = 5
	// Data Type: STRING. Origin name. Shorter names are recommended.
	TravelPlaceholderFieldEnum_ORIGIN_NAME TravelPlaceholderFieldEnum_TravelPlaceholderField = 6
	// Data Type: STRING. Price to be shown in the ad. Highly recommended for
	// dynamic ads.
	// Example: "100.00 USD"
	TravelPlaceholderFieldEnum_PRICE TravelPlaceholderFieldEnum_TravelPlaceholderField = 7
	// Data Type: STRING. Formatted price to be shown in the ad.
	// Example: "Starting at $100.00 USD", "$80 - $100"
	TravelPlaceholderFieldEnum_FORMATTED_PRICE TravelPlaceholderFieldEnum_TravelPlaceholderField = 8
	// Data Type: STRING. Sale price to be shown in the ad.
	// Example: "80.00 USD"
	TravelPlaceholderFieldEnum_SALE_PRICE TravelPlaceholderFieldEnum_TravelPlaceholderField = 9
	// Data Type: STRING. Formatted sale price to be shown in the ad.
	// Example: "On sale for $80.00", "$60 - $80"
	TravelPlaceholderFieldEnum_FORMATTED_SALE_PRICE TravelPlaceholderFieldEnum_TravelPlaceholderField = 10
	// Data Type: URL. Image to be displayed in the ad.
	TravelPlaceholderFieldEnum_IMAGE_URL TravelPlaceholderFieldEnum_TravelPlaceholderField = 11
	// Data Type: STRING. Category of travel offer used to group like items
	// together for recommendation engine.
	TravelPlaceholderFieldEnum_CATEGORY TravelPlaceholderFieldEnum_TravelPlaceholderField = 12
	// Data Type: STRING_LIST. Keywords used for product retrieval.
	TravelPlaceholderFieldEnum_CONTEXTUAL_KEYWORDS TravelPlaceholderFieldEnum_TravelPlaceholderField = 13
	// Data Type: STRING. Address of travel offer, including postal code.
	TravelPlaceholderFieldEnum_DESTINATION_ADDRESS TravelPlaceholderFieldEnum_TravelPlaceholderField = 14
	// Data Type: URL_LIST. Required. Final URLs to be used in ad, when using
	// Upgraded URLs; the more specific the better (e.g. the individual URL of a
	// specific travel offer and its location).
	TravelPlaceholderFieldEnum_FINAL_URL TravelPlaceholderFieldEnum_TravelPlaceholderField = 15
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	TravelPlaceholderFieldEnum_FINAL_MOBILE_URLS TravelPlaceholderFieldEnum_TravelPlaceholderField = 16
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	TravelPlaceholderFieldEnum_TRACKING_URL TravelPlaceholderFieldEnum_TravelPlaceholderField = 17
	// Data Type: STRING. Android app link. Must be formatted as:
	// android-app://{package_id}/{scheme}/{host_path}.
	// The components are defined as follows:
	// package_id: app ID as specified in Google Play.
	// scheme: the scheme to pass to the application. Can be HTTP, or a custom
	//   scheme.
	// host_path: identifies the specific content within your application.
	TravelPlaceholderFieldEnum_ANDROID_APP_LINK TravelPlaceholderFieldEnum_TravelPlaceholderField = 18
	// Data Type: STRING_LIST. List of recommended destination IDs to show
	// together with this item.
	TravelPlaceholderFieldEnum_SIMILAR_DESTINATION_IDS TravelPlaceholderFieldEnum_TravelPlaceholderField = 19
	// Data Type: STRING. iOS app link.
	TravelPlaceholderFieldEnum_IOS_APP_LINK TravelPlaceholderFieldEnum_TravelPlaceholderField = 20
	// Data Type: INT64. iOS app store ID.
	TravelPlaceholderFieldEnum_IOS_APP_STORE_ID TravelPlaceholderFieldEnum_TravelPlaceholderField = 21
)

var TravelPlaceholderFieldEnum_TravelPlaceholderField_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "DESTINATION_ID",
	3:  "ORIGIN_ID",
	4:  "TITLE",
	5:  "DESTINATION_NAME",
	6:  "ORIGIN_NAME",
	7:  "PRICE",
	8:  "FORMATTED_PRICE",
	9:  "SALE_PRICE",
	10: "FORMATTED_SALE_PRICE",
	11: "IMAGE_URL",
	12: "CATEGORY",
	13: "CONTEXTUAL_KEYWORDS",
	14: "DESTINATION_ADDRESS",
	15: "FINAL_URL",
	16: "FINAL_MOBILE_URLS",
	17: "TRACKING_URL",
	18: "ANDROID_APP_LINK",
	19: "SIMILAR_DESTINATION_IDS",
	20: "IOS_APP_LINK",
	21: "IOS_APP_STORE_ID",
}

var TravelPlaceholderFieldEnum_TravelPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":             0,
	"UNKNOWN":                 1,
	"DESTINATION_ID":          2,
	"ORIGIN_ID":               3,
	"TITLE":                   4,
	"DESTINATION_NAME":        5,
	"ORIGIN_NAME":             6,
	"PRICE":                   7,
	"FORMATTED_PRICE":         8,
	"SALE_PRICE":              9,
	"FORMATTED_SALE_PRICE":    10,
	"IMAGE_URL":               11,
	"CATEGORY":                12,
	"CONTEXTUAL_KEYWORDS":     13,
	"DESTINATION_ADDRESS":     14,
	"FINAL_URL":               15,
	"FINAL_MOBILE_URLS":       16,
	"TRACKING_URL":            17,
	"ANDROID_APP_LINK":        18,
	"SIMILAR_DESTINATION_IDS": 19,
	"IOS_APP_LINK":            20,
	"IOS_APP_STORE_ID":        21,
}

func (x TravelPlaceholderFieldEnum_TravelPlaceholderField) String() string {
	return proto.EnumName(TravelPlaceholderFieldEnum_TravelPlaceholderField_name, int32(x))
}

func (TravelPlaceholderFieldEnum_TravelPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ca7a3e1388fd357c, []int{0, 0}
}

// Values for Travel placeholder fields.
// For more information about dynamic remarketing feeds, see
// https://support.google.com/google-ads/answer/6053288.
type TravelPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TravelPlaceholderFieldEnum) Reset()         { *m = TravelPlaceholderFieldEnum{} }
func (m *TravelPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*TravelPlaceholderFieldEnum) ProtoMessage()    {}
func (*TravelPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca7a3e1388fd357c, []int{0}
}

func (m *TravelPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TravelPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *TravelPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TravelPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (m *TravelPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TravelPlaceholderFieldEnum.Merge(m, src)
}
func (m *TravelPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_TravelPlaceholderFieldEnum.Size(m)
}
func (m *TravelPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_TravelPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_TravelPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.TravelPlaceholderFieldEnum_TravelPlaceholderField", TravelPlaceholderFieldEnum_TravelPlaceholderField_name, TravelPlaceholderFieldEnum_TravelPlaceholderField_value)
	proto.RegisterType((*TravelPlaceholderFieldEnum)(nil), "google.ads.googleads.v1.enums.TravelPlaceholderFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/travel_placeholder_field.proto", fileDescriptor_ca7a3e1388fd357c)
}

var fileDescriptor_ca7a3e1388fd357c = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x66, 0x2d, 0xfb, 0x73, 0xbb, 0xd6, 0x73, 0x3b, 0x86, 0x36, 0x76, 0xb1, 0x3d, 0x40, 0xa2,
	0x8a, 0xbb, 0xc0, 0x8d, 0xdb, 0xb8, 0x91, 0xd5, 0xd4, 0x89, 0x6c, 0xb7, 0x63, 0xa8, 0x52, 0x14,
	0x96, 0x10, 0x2a, 0xa5, 0x49, 0xd5, 0x74, 0x7d, 0x20, 0x2e, 0x79, 0x01, 0xde, 0x81, 0x07, 0xe0,
	0x21, 0xb8, 0xe2, 0x11, 0x90, 0x9d, 0xfe, 0x21, 0x0d, 0x6e, 0xa2, 0x73, 0xce, 0xf7, 0xe3, 0x13,
	0xfb, 0x03, 0xef, 0x93, 0x3c, 0x4f, 0xd2, 0xd8, 0x0c, 0xa3, 0xc2, 0x2c, 0x4b, 0x55, 0xad, 0x3a,
	0x66, 0x9c, 0x3d, 0xcd, 0x0a, 0x73, 0xb9, 0x08, 0x57, 0x71, 0x1a, 0xcc, 0xd3, 0xf0, 0x31, 0xfe,
	0x92, 0xa7, 0x51, 0xbc, 0x08, 0x3e, 0x4f, 0xe3, 0x34, 0x32, 0xe6, 0x8b, 0x7c, 0x99, 0xa3, 0x9b,
	0x52, 0x62, 0x84, 0x51, 0x61, 0x6c, 0xd5, 0xc6, 0xaa, 0x63, 0x68, 0xf5, 0xd5, 0x9b, 0x8d, 0xf9,
	0x7c, 0x6a, 0x86, 0x59, 0x96, 0x2f, 0xc3, 0xe5, 0x34, 0xcf, 0x8a, 0x52, 0x7c, 0xf7, 0xb3, 0x0a,
	0xae, 0xa4, 0xf6, 0xf7, 0x77, 0xf6, 0x7d, 0xe5, 0x4e, 0xb2, 0xa7, 0xd9, 0xdd, 0xf7, 0x2a, 0x78,
	0xf5, 0x3c, 0x8c, 0x9a, 0xa0, 0x36, 0x62, 0xc2, 0x27, 0x3d, 0xda, 0xa7, 0xc4, 0x86, 0x2f, 0x50,
	0x0d, 0x1c, 0x8f, 0xd8, 0x80, 0x79, 0xf7, 0x0c, 0x1e, 0x20, 0x04, 0x1a, 0x36, 0x11, 0x92, 0x32,
	0x2c, 0xa9, 0xc7, 0x02, 0x6a, 0xc3, 0x0a, 0x3a, 0x03, 0xa7, 0x1e, 0xa7, 0x0e, 0xd5, 0x6d, 0x15,
	0x9d, 0x82, 0x43, 0x49, 0xa5, 0x4b, 0xe0, 0x4b, 0xd4, 0x06, 0x70, 0x9f, 0xcd, 0xf0, 0x90, 0xc0,
	0x43, 0x75, 0xc2, 0x9a, 0xaf, 0x07, 0x47, 0x4a, 0xe1, 0x73, 0xda, 0x23, 0xf0, 0x18, 0xb5, 0x40,
	0xb3, 0xef, 0xf1, 0x21, 0x96, 0x92, 0xd8, 0x41, 0x39, 0x3c, 0x41, 0x0d, 0x00, 0x04, 0x76, 0xc9,
	0xba, 0x3f, 0x45, 0xaf, 0x41, 0x7b, 0x47, 0xda, 0x43, 0x80, 0x5a, 0x85, 0x0e, 0xb1, 0x43, 0x82,
	0x11, 0x77, 0x61, 0x0d, 0xd5, 0xc1, 0x49, 0x0f, 0x4b, 0xe2, 0x78, 0xfc, 0x01, 0xd6, 0xd1, 0x25,
	0x68, 0xf5, 0x3c, 0x26, 0xc9, 0x07, 0x39, 0xc2, 0x6e, 0x30, 0x20, 0x0f, 0xf7, 0x1e, 0xb7, 0x05,
	0x3c, 0x53, 0xc0, 0xfe, 0x9a, 0xd8, 0xb6, 0x39, 0x11, 0x02, 0x36, 0x94, 0x5d, 0x9f, 0x32, 0xec,
	0x6a, 0xbb, 0x26, 0xba, 0x00, 0xe7, 0x65, 0x3b, 0xf4, 0xba, 0xd4, 0xd5, 0x87, 0x08, 0x08, 0x11,
	0x04, 0x75, 0xc9, 0x71, 0x6f, 0x40, 0x99, 0xa3, 0x89, 0xe7, 0xea, 0xbf, 0x31, 0xb3, 0xb9, 0x47,
	0xed, 0x00, 0xfb, 0x7e, 0xe0, 0x52, 0x36, 0x80, 0x08, 0x5d, 0x83, 0x4b, 0x41, 0x87, 0xd4, 0xc5,
	0x3c, 0xf8, 0xfb, 0x0e, 0x05, 0x6c, 0x29, 0x13, 0xea, 0x89, 0x1d, 0xbd, 0xad, 0x4c, 0x36, 0x13,
	0x21, 0x3d, 0x4e, 0xd4, 0xed, 0x5e, 0x74, 0x7f, 0x1f, 0x80, 0xdb, 0xc7, 0x7c, 0x66, 0xfc, 0x37,
	0x1c, 0xdd, 0xeb, 0xe7, 0x1f, 0xd7, 0x57, 0xd9, 0xf0, 0x0f, 0x3e, 0x76, 0xd7, 0xea, 0x24, 0x4f,
	0xc3, 0x2c, 0x31, 0xf2, 0x45, 0x62, 0x26, 0x71, 0xa6, 0x93, 0xb3, 0x09, 0xea, 0x7c, 0x5a, 0xfc,
	0x23, 0xb7, 0xef, 0xf4, 0xf7, 0x6b, 0xa5, 0xea, 0x60, 0xfc, 0xad, 0x72, 0xe3, 0x94, 0x56, 0x38,
	0x2a, 0x8c, 0xb2, 0x54, 0xd5, 0xb8, 0x63, 0xa8, 0x9c, 0x15, 0x3f, 0x36, 0xf8, 0x04, 0x47, 0xc5,
	0x64, 0x8b, 0x4f, 0xc6, 0x9d, 0x89, 0xc6, 0x7f, 0x55, 0x6e, 0xcb, 0xa1, 0x65, 0xe1, 0xa8, 0xb0,
	0xac, 0x2d, 0xc3, 0xb2, 0xc6, 0x1d, 0xcb, 0xd2, 0x9c, 0x4f, 0x47, 0x7a, 0xb1, 0xb7, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xb1, 0xed, 0x26, 0x00, 0x4f, 0x03, 0x00, 0x00,
}
