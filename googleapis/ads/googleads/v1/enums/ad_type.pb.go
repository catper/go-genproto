// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/ad_type.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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
}

var AdTypeEnum_AdType_value = map[string]int32{
	"UNSPECIFIED":                  0,
	"UNKNOWN":                      1,
	"TEXT_AD":                      2,
	"EXPANDED_TEXT_AD":             3,
	"CALL_ONLY_AD":                 6,
	"EXPANDED_DYNAMIC_SEARCH_AD":   7,
	"HOTEL_AD":                     8,
	"SHOPPING_SMART_AD":            9,
	"SHOPPING_PRODUCT_AD":          10,
	"VIDEO_AD":                     12,
	"GMAIL_AD":                     13,
	"IMAGE_AD":                     14,
	"RESPONSIVE_SEARCH_AD":         15,
	"LEGACY_RESPONSIVE_DISPLAY_AD": 16,
	"APP_AD":                       17,
	"LEGACY_APP_INSTALL_AD":        18,
	"RESPONSIVE_DISPLAY_AD":        19,
	"HTML5_UPLOAD_AD":              21,
	"DYNAMIC_HTML5_AD":             22,
	"APP_ENGAGEMENT_AD":            23,
}

func (x AdTypeEnum_AdType) String() string {
	return proto.EnumName(AdTypeEnum_AdType_name, int32(x))
}

func (AdTypeEnum_AdType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c765db6f43f63180, []int{0, 0}
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
	return fileDescriptor_c765db6f43f63180, []int{0}
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
	proto.RegisterEnum("google.ads.googleads.v1.enums.AdTypeEnum_AdType", AdTypeEnum_AdType_name, AdTypeEnum_AdType_value)
	proto.RegisterType((*AdTypeEnum)(nil), "google.ads.googleads.v1.enums.AdTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/ad_type.proto", fileDescriptor_c765db6f43f63180)
}

var fileDescriptor_c765db6f43f63180 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xa6, 0x89, 0x94, 0x96, 0x49, 0x20, 0xee, 0xa6, 0xa1, 0x50, 0xb5, 0x88, 0xf6, 0x8a, 0x64,
	0x2b, 0x42, 0x5c, 0xcc, 0x69, 0xe2, 0x5d, 0x1c, 0x0b, 0x7b, 0xbd, 0x8a, 0x9d, 0xd0, 0xa0, 0x48,
	0x96, 0xc1, 0x91, 0x15, 0xa9, 0xb1, 0xa3, 0x3a, 0xad, 0xd4, 0xd7, 0x81, 0x1b, 0x47, 0x1e, 0x83,
	0xa7, 0xe0, 0xcc, 0x53, 0xa0, 0x59, 0x37, 0x51, 0x0f, 0xc0, 0xc5, 0x9a, 0xef, 0xc7, 0x9f, 0x76,
	0x7e, 0xe0, 0x75, 0x5e, 0x96, 0xf9, 0xd5, 0xc2, 0x4a, 0xb3, 0xca, 0xaa, 0x4b, 0xaa, 0x6e, 0x07,
	0xd6, 0xa2, 0xb8, 0x59, 0x55, 0x56, 0x9a, 0x25, 0x9b, 0xbb, 0xf5, 0xc2, 0x5c, 0x5f, 0x97, 0x9b,
	0x92, 0x9d, 0xd5, 0x0e, 0x33, 0xcd, 0x2a, 0x73, 0x67, 0x36, 0x6f, 0x07, 0xa6, 0x36, 0x9f, 0x9c,
	0x6e, 0xb3, 0xd6, 0x4b, 0x2b, 0x2d, 0x8a, 0x72, 0x93, 0x6e, 0x96, 0x65, 0x51, 0xd5, 0x3f, 0x5f,
	0xfc, 0x68, 0x02, 0x60, 0x16, 0xdf, 0xad, 0x17, 0xa2, 0xb8, 0x59, 0x5d, 0x7c, 0x6b, 0x42, 0xab,
	0x86, 0xac, 0x0b, 0xed, 0x89, 0x8c, 0x94, 0x70, 0xbc, 0xf7, 0x9e, 0xe0, 0xc6, 0x23, 0xd6, 0x86,
	0xfd, 0x89, 0xfc, 0x20, 0xc3, 0x8f, 0xd2, 0xd8, 0x23, 0x10, 0x8b, 0xcb, 0x38, 0x41, 0x6e, 0x34,
	0xd8, 0x11, 0x18, 0xe2, 0x52, 0xa1, 0xe4, 0x82, 0x27, 0x5b, 0xb6, 0xc9, 0x0c, 0xe8, 0x38, 0xe8,
	0xfb, 0x49, 0x28, 0xfd, 0x19, 0x31, 0x2d, 0xf6, 0x12, 0x4e, 0x76, 0x3e, 0x3e, 0x93, 0x18, 0x78,
	0x4e, 0x12, 0x09, 0x1c, 0x3b, 0x23, 0xd2, 0xf7, 0x59, 0x07, 0x0e, 0x46, 0x61, 0x2c, 0x7c, 0x42,
	0x07, 0xac, 0x0f, 0x87, 0xd1, 0x28, 0x54, 0xca, 0x93, 0x6e, 0x12, 0x05, 0x38, 0xd6, 0xb1, 0x8f,
	0xd9, 0x31, 0xf4, 0x76, 0xb4, 0x1a, 0x87, 0x7c, 0xe2, 0x68, 0x01, 0xe8, 0xef, 0xa9, 0xc7, 0x45,
	0x48, 0xa8, 0x43, 0xc8, 0x0d, 0xd0, 0xd3, 0x59, 0x4f, 0x08, 0x79, 0x01, 0xba, 0x82, 0xd0, 0x53,
	0xf6, 0x1c, 0x8e, 0xc6, 0x22, 0x52, 0xa1, 0x8c, 0xbc, 0xa9, 0x78, 0xf0, 0x82, 0x2e, 0x7b, 0x05,
	0xa7, 0xbe, 0x70, 0xd1, 0x99, 0x25, 0x0f, 0x0c, 0xdc, 0x8b, 0x94, 0x8f, 0xba, 0x07, 0x83, 0x01,
	0xb4, 0x50, 0x29, 0xaa, 0x0f, 0xd9, 0x0b, 0xe8, 0xdf, 0xbb, 0x89, 0xf2, 0x64, 0x14, 0x53, 0xbf,
	0xc8, 0x0d, 0x46, 0xd2, 0xdf, 0x13, 0x7a, 0xac, 0x07, 0xdd, 0x51, 0x1c, 0xf8, 0x6f, 0x93, 0x89,
	0xf2, 0x43, 0xe4, 0x44, 0xf6, 0x69, 0x84, 0xdb, 0x89, 0xd4, 0x22, 0x72, 0xe3, 0x19, 0x8d, 0x80,
	0x92, 0x85, 0x74, 0xd1, 0x15, 0x81, 0x90, 0xba, 0xd3, 0xe3, 0xe1, 0xaf, 0x3d, 0x38, 0xff, 0x52,
	0xae, 0xcc, 0xff, 0x2e, 0x7e, 0xd8, 0xae, 0x17, 0xa9, 0x68, 0xcf, 0x6a, 0xef, 0xd3, 0xf0, 0xde,
	0x9d, 0x97, 0x57, 0x69, 0x91, 0x9b, 0xe5, 0x75, 0x6e, 0xe5, 0x8b, 0x42, 0x5f, 0xc1, 0xf6, 0xc6,
	0xd6, 0xcb, 0xea, 0x1f, 0x27, 0xf7, 0x4e, 0x7f, 0xbf, 0x36, 0x9a, 0x2e, 0xe2, 0xf7, 0xc6, 0x99,
	0x5b, 0x47, 0x61, 0x56, 0x99, 0x75, 0x49, 0xd5, 0x74, 0x60, 0xd2, 0x0d, 0x55, 0x3f, 0xb7, 0xfa,
	0x1c, 0xb3, 0x6a, 0xbe, 0xd3, 0xe7, 0xd3, 0xc1, 0x5c, 0xeb, 0xbf, 0x1b, 0xe7, 0x35, 0x69, 0xdb,
	0x98, 0x55, 0xb6, 0xbd, 0x73, 0xd8, 0xf6, 0x74, 0x60, 0xdb, 0xda, 0xf3, 0xb9, 0xa5, 0x1f, 0xf6,
	0xe6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0x88, 0xc1, 0xa2, 0x0a, 0x03, 0x00, 0x00,
}
