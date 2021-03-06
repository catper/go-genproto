// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/ad_type.proto

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

// The possible types of an ad.
type AdTypeEnum_AdType int32

const (
	// No value has been specified.
	AdTypeEnum_UNSPECIFIED AdTypeEnum_AdType = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdTypeEnum_UNKNOWN AdTypeEnum_AdType = 1
	// The ad is a text ad.
	AdTypeEnum_TEXT_AD AdTypeEnum_AdType = 2
	// The ad is an expanded text ad.
	AdTypeEnum_EXPANDED_TEXT_AD AdTypeEnum_AdType = 3
	// The ad is a call only ad.
	AdTypeEnum_CALL_ONLY_AD AdTypeEnum_AdType = 6
	// The ad is an expanded dynamic search ad.
	AdTypeEnum_EXPANDED_DYNAMIC_SEARCH_AD AdTypeEnum_AdType = 7
	// The ad is a hotel ad.
	AdTypeEnum_HOTEL_AD AdTypeEnum_AdType = 8
	// The ad is a Smart Shopping ad.
	AdTypeEnum_SHOPPING_SMART_AD AdTypeEnum_AdType = 9
	// The ad is a standard Shopping ad.
	AdTypeEnum_SHOPPING_PRODUCT_AD AdTypeEnum_AdType = 10
	// The ad is a video ad.
	AdTypeEnum_VIDEO_AD AdTypeEnum_AdType = 12
	// This ad is a Gmail ad.
	AdTypeEnum_GMAIL_AD AdTypeEnum_AdType = 13
	// This ad is an Image ad.
	AdTypeEnum_IMAGE_AD AdTypeEnum_AdType = 14
	// The ad is a responsive search ad.
	AdTypeEnum_RESPONSIVE_SEARCH_AD AdTypeEnum_AdType = 15
	// The ad is a legacy responsive display ad.
	AdTypeEnum_LEGACY_RESPONSIVE_DISPLAY_AD AdTypeEnum_AdType = 16
	// The ad is an app ad.
	AdTypeEnum_APP_AD AdTypeEnum_AdType = 17
	// The ad is a legacy app install ad.
	AdTypeEnum_LEGACY_APP_INSTALL_AD AdTypeEnum_AdType = 18
	// The ad is a responsive display ad.
	AdTypeEnum_RESPONSIVE_DISPLAY_AD AdTypeEnum_AdType = 19
	// The ad is a display upload ad with the HTML5_UPLOAD_AD product type.
	AdTypeEnum_HTML5_UPLOAD_AD AdTypeEnum_AdType = 21
	// The ad is a display upload ad with one of the DYNAMIC_HTML5_* product
	// types.
	AdTypeEnum_DYNAMIC_HTML5_AD AdTypeEnum_AdType = 22
	// The ad is an app engagement ad.
	AdTypeEnum_APP_ENGAGEMENT_AD AdTypeEnum_AdType = 23
	// The ad is a Shopping Comparison Listing ad.
	AdTypeEnum_SHOPPING_COMPARISON_LISTING_AD AdTypeEnum_AdType = 24
)

var AdTypeEnum_AdType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "TEXT_AD",
	3:  "EXPANDED_TEXT_AD",
	6:  "CALL_ONLY_AD",
	7:  "EXPANDED_DYNAMIC_SEARCH_AD",
	8:  "HOTEL_AD",
	9:  "SHOPPING_SMART_AD",
	10: "SHOPPING_PRODUCT_AD",
	12: "VIDEO_AD",
	13: "GMAIL_AD",
	14: "IMAGE_AD",
	15: "RESPONSIVE_SEARCH_AD",
	16: "LEGACY_RESPONSIVE_DISPLAY_AD",
	17: "APP_AD",
	18: "LEGACY_APP_INSTALL_AD",
	19: "RESPONSIVE_DISPLAY_AD",
	21: "HTML5_UPLOAD_AD",
	22: "DYNAMIC_HTML5_AD",
	23: "APP_ENGAGEMENT_AD",
	24: "SHOPPING_COMPARISON_LISTING_AD",
}

var AdTypeEnum_AdType_value = map[string]int32{
	"UNSPECIFIED":                    0,
	"UNKNOWN":                        1,
	"TEXT_AD":                        2,
	"EXPANDED_TEXT_AD":               3,
	"CALL_ONLY_AD":                   6,
	"EXPANDED_DYNAMIC_SEARCH_AD":     7,
	"HOTEL_AD":                       8,
	"SHOPPING_SMART_AD":              9,
	"SHOPPING_PRODUCT_AD":            10,
	"VIDEO_AD":                       12,
	"GMAIL_AD":                       13,
	"IMAGE_AD":                       14,
	"RESPONSIVE_SEARCH_AD":           15,
	"LEGACY_RESPONSIVE_DISPLAY_AD":   16,
	"APP_AD":                         17,
	"LEGACY_APP_INSTALL_AD":          18,
	"RESPONSIVE_DISPLAY_AD":          19,
	"HTML5_UPLOAD_AD":                21,
	"DYNAMIC_HTML5_AD":               22,
	"APP_ENGAGEMENT_AD":              23,
	"SHOPPING_COMPARISON_LISTING_AD": 24,
}

func (x AdTypeEnum_AdType) String() string {
	return proto.EnumName(AdTypeEnum_AdType_name, int32(x))
}

func (AdTypeEnum_AdType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2322810cfdc14ea3, []int{0, 0}
}

// Container for enum describing possible types of an ad.
type AdTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdTypeEnum) Reset()         { *m = AdTypeEnum{} }
func (m *AdTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AdTypeEnum) ProtoMessage()    {}
func (*AdTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_2322810cfdc14ea3, []int{0}
}

func (m *AdTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdTypeEnum.Unmarshal(m, b)
}
func (m *AdTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdTypeEnum.Marshal(b, m, deterministic)
}
func (m *AdTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdTypeEnum.Merge(m, src)
}
func (m *AdTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AdTypeEnum.Size(m)
}
func (m *AdTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.AdTypeEnum_AdType", AdTypeEnum_AdType_name, AdTypeEnum_AdType_value)
	proto.RegisterType((*AdTypeEnum)(nil), "google.ads.googleads.v3.enums.AdTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/ad_type.proto", fileDescriptor_2322810cfdc14ea3)
}

var fileDescriptor_2322810cfdc14ea3 = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x6a, 0xdb, 0x4c,
	0x14, 0xfd, 0x62, 0x83, 0x93, 0x6f, 0xe2, 0xd6, 0xca, 0x38, 0x6e, 0xd2, 0x90, 0x84, 0xc6, 0xdb,
	0x82, 0xb4, 0x10, 0xdd, 0xa8, 0xab, 0x6b, 0xcd, 0x54, 0x1e, 0x2a, 0x8d, 0x06, 0x4b, 0x76, 0xe3,
	0x62, 0x10, 0x6a, 0x65, 0x84, 0x21, 0x96, 0x4c, 0xe4, 0x04, 0xf2, 0x3a, 0x5d, 0xf6, 0x45, 0x0a,
	0x7d, 0x80, 0xae, 0xbb, 0xee, 0x53, 0x94, 0x3b, 0x8a, 0x45, 0x16, 0x6d, 0x37, 0xe2, 0x9e, 0x9f,
	0x39, 0xcc, 0x5c, 0x1d, 0xf2, 0x3a, 0x2f, 0xcb, 0xfc, 0x66, 0x69, 0xa5, 0x59, 0x65, 0xd5, 0x23,
	0x4e, 0xf7, 0xb6, 0xb5, 0x2c, 0xee, 0xd6, 0x95, 0x95, 0x66, 0xc9, 0xf6, 0x61, 0xb3, 0x34, 0x37,
	0xb7, 0xe5, 0xb6, 0xa4, 0x17, 0xb5, 0xc3, 0x4c, 0xb3, 0xca, 0x6c, 0xcc, 0xe6, 0xbd, 0x6d, 0x6a,
	0xf3, 0xd9, 0xf9, 0x2e, 0x6b, 0xb3, 0xb2, 0xd2, 0xa2, 0x28, 0xb7, 0xe9, 0x76, 0x55, 0x16, 0x55,
	0x7d, 0x78, 0xf8, 0xa3, 0x4d, 0x08, 0x64, 0xf1, 0xc3, 0x66, 0xc9, 0x8b, 0xbb, 0xf5, 0xf0, 0x5b,
	0x9b, 0x74, 0x6a, 0x48, 0x7b, 0xe4, 0x70, 0x2a, 0x23, 0xc5, 0x5d, 0xf1, 0x4e, 0x70, 0x66, 0xfc,
	0x47, 0x0f, 0xc9, 0xfe, 0x54, 0xbe, 0x97, 0xe1, 0x07, 0x69, 0xec, 0x21, 0x88, 0xf9, 0x75, 0x9c,
	0x00, 0x33, 0x5a, 0xf4, 0x98, 0x18, 0xfc, 0x5a, 0x81, 0x64, 0x9c, 0x25, 0x3b, 0xb6, 0x4d, 0x0d,
	0xd2, 0x75, 0xc1, 0xf7, 0x93, 0x50, 0xfa, 0x73, 0x64, 0x3a, 0xf4, 0x92, 0x9c, 0x35, 0x3e, 0x36,
	0x97, 0x10, 0x08, 0x37, 0x89, 0x38, 0x4c, 0xdc, 0x31, 0xea, 0xfb, 0xb4, 0x4b, 0x0e, 0xc6, 0x61,
	0xcc, 0x7d, 0x44, 0x07, 0x74, 0x40, 0x8e, 0xa2, 0x71, 0xa8, 0x94, 0x90, 0x5e, 0x12, 0x05, 0x30,
	0xd1, 0xb1, 0xff, 0xd3, 0x13, 0xd2, 0x6f, 0x68, 0x35, 0x09, 0xd9, 0xd4, 0xd5, 0x02, 0xc1, 0xd3,
	0x33, 0xc1, 0x78, 0x88, 0xa8, 0x8b, 0xc8, 0x0b, 0x40, 0xe8, 0xac, 0x67, 0x88, 0x44, 0x00, 0x1e,
	0x47, 0xf4, 0x9c, 0x9e, 0x92, 0xe3, 0x09, 0x8f, 0x54, 0x28, 0x23, 0x31, 0xe3, 0x4f, 0x6e, 0xd0,
	0xa3, 0xaf, 0xc8, 0xb9, 0xcf, 0x3d, 0x70, 0xe7, 0xc9, 0x13, 0x03, 0x13, 0x91, 0xf2, 0x41, 0xbf,
	0xc1, 0xa0, 0x84, 0x74, 0x40, 0x29, 0x9c, 0x8f, 0xe8, 0x4b, 0x32, 0x78, 0x74, 0x23, 0x25, 0x64,
	0x14, 0xe3, 0x7b, 0x81, 0x19, 0x14, 0xa5, 0x3f, 0x27, 0xf4, 0x69, 0x9f, 0xf4, 0xc6, 0x71, 0xe0,
	0xbf, 0x49, 0xa6, 0xca, 0x0f, 0x81, 0x21, 0x39, 0xc0, 0x15, 0xee, 0x36, 0x52, 0x8b, 0xc0, 0x8c,
	0x17, 0xb8, 0x02, 0x4c, 0xe6, 0xd2, 0x03, 0x8f, 0x07, 0x5c, 0xea, 0x97, 0x9e, 0xd0, 0x21, 0xb9,
	0x6c, 0x56, 0xe0, 0x86, 0x81, 0x82, 0x89, 0x88, 0x42, 0x99, 0xf8, 0x22, 0x8a, 0x91, 0x02, 0x66,
	0x9c, 0x8e, 0x7e, 0xee, 0x91, 0xab, 0xcf, 0xe5, 0xda, 0xfc, 0x67, 0x39, 0x46, 0x87, 0xf5, 0xcf,
	0x56, 0xd8, 0x05, 0xb5, 0xf7, 0x71, 0xf4, 0xe8, 0xce, 0xcb, 0x9b, 0xb4, 0xc8, 0xcd, 0xf2, 0x36,
	0xb7, 0xf2, 0x65, 0xa1, 0x9b, 0xb2, 0xeb, 0xe1, 0x66, 0x55, 0xfd, 0xa5, 0x96, 0x6f, 0xf5, 0xf7,
	0x4b, 0xab, 0xed, 0x01, 0x7c, 0x6d, 0x5d, 0x78, 0x75, 0x14, 0x64, 0x95, 0x59, 0x8f, 0x38, 0xcd,
	0x6c, 0x13, 0x7b, 0x56, 0x7d, 0xdf, 0xe9, 0x0b, 0xc8, 0xaa, 0x45, 0xa3, 0x2f, 0x66, 0xf6, 0x42,
	0xeb, 0xbf, 0x5a, 0x57, 0x35, 0xe9, 0x38, 0x90, 0x55, 0x8e, 0xd3, 0x38, 0x1c, 0x67, 0x66, 0x3b,
	0x8e, 0xf6, 0x7c, 0xea, 0xe8, 0x8b, 0xd9, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xc5, 0x68,
	0x27, 0x2e, 0x03, 0x00, 0x00,
}
