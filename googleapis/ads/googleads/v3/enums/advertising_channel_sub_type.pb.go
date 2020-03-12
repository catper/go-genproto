// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/advertising_channel_sub_type.proto

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

// Enum describing the different channel subtypes.
type AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType int32

const (
	// Not specified.
	AdvertisingChannelSubTypeEnum_UNSPECIFIED AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 0
	// Used as a return value only. Represents value unknown in this version.
	AdvertisingChannelSubTypeEnum_UNKNOWN AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 1
	// Mobile app campaigns for Search.
	AdvertisingChannelSubTypeEnum_SEARCH_MOBILE_APP AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 2
	// Mobile app campaigns for Display.
	AdvertisingChannelSubTypeEnum_DISPLAY_MOBILE_APP AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 3
	// AdWords express campaigns for search.
	AdvertisingChannelSubTypeEnum_SEARCH_EXPRESS AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 4
	// AdWords Express campaigns for display.
	AdvertisingChannelSubTypeEnum_DISPLAY_EXPRESS AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 5
	// Smart Shopping campaigns.
	AdvertisingChannelSubTypeEnum_SHOPPING_SMART_ADS AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 6
	// Gmail Ad campaigns.
	AdvertisingChannelSubTypeEnum_DISPLAY_GMAIL_AD AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 7
	// Smart display campaigns.
	AdvertisingChannelSubTypeEnum_DISPLAY_SMART_CAMPAIGN AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 8
	// Video Outstream campaigns.
	AdvertisingChannelSubTypeEnum_VIDEO_OUTSTREAM AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 9
	// Video TrueView for Action campaigns.
	AdvertisingChannelSubTypeEnum_VIDEO_ACTION AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 10
	// Video campaigns with non-skippable video ads.
	AdvertisingChannelSubTypeEnum_VIDEO_NON_SKIPPABLE AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 11
	// App Campaign that allows you to easily promote your Android or iOS app
	// across Google's top properties including Search, Play, YouTube, and the
	// Google Display Network.
	AdvertisingChannelSubTypeEnum_APP_CAMPAIGN AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 12
	// App Campaign for engagement, focused on driving re-engagement with the
	// app across several of Google’s top properties including Search, YouTube,
	// and the Google Display Network.
	AdvertisingChannelSubTypeEnum_APP_CAMPAIGN_FOR_ENGAGEMENT AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 13
	// Shopping Comparison Listing campaigns.
	AdvertisingChannelSubTypeEnum_SHOPPING_COMPARISON_LISTING_ADS AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 15
)

var AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "SEARCH_MOBILE_APP",
	3:  "DISPLAY_MOBILE_APP",
	4:  "SEARCH_EXPRESS",
	5:  "DISPLAY_EXPRESS",
	6:  "SHOPPING_SMART_ADS",
	7:  "DISPLAY_GMAIL_AD",
	8:  "DISPLAY_SMART_CAMPAIGN",
	9:  "VIDEO_OUTSTREAM",
	10: "VIDEO_ACTION",
	11: "VIDEO_NON_SKIPPABLE",
	12: "APP_CAMPAIGN",
	13: "APP_CAMPAIGN_FOR_ENGAGEMENT",
	15: "SHOPPING_COMPARISON_LISTING_ADS",
}

var AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_value = map[string]int32{
	"UNSPECIFIED":                     0,
	"UNKNOWN":                         1,
	"SEARCH_MOBILE_APP":               2,
	"DISPLAY_MOBILE_APP":              3,
	"SEARCH_EXPRESS":                  4,
	"DISPLAY_EXPRESS":                 5,
	"SHOPPING_SMART_ADS":              6,
	"DISPLAY_GMAIL_AD":                7,
	"DISPLAY_SMART_CAMPAIGN":          8,
	"VIDEO_OUTSTREAM":                 9,
	"VIDEO_ACTION":                    10,
	"VIDEO_NON_SKIPPABLE":             11,
	"APP_CAMPAIGN":                    12,
	"APP_CAMPAIGN_FOR_ENGAGEMENT":     13,
	"SHOPPING_COMPARISON_LISTING_ADS": 15,
}

func (x AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType) String() string {
	return proto.EnumName(AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_name, int32(x))
}

func (AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_007989ef93506f12, []int{0, 0}
}

// An immutable specialization of an Advertising Channel.
type AdvertisingChannelSubTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdvertisingChannelSubTypeEnum) Reset()         { *m = AdvertisingChannelSubTypeEnum{} }
func (m *AdvertisingChannelSubTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AdvertisingChannelSubTypeEnum) ProtoMessage()    {}
func (*AdvertisingChannelSubTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_007989ef93506f12, []int{0}
}

func (m *AdvertisingChannelSubTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdvertisingChannelSubTypeEnum.Unmarshal(m, b)
}
func (m *AdvertisingChannelSubTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdvertisingChannelSubTypeEnum.Marshal(b, m, deterministic)
}
func (m *AdvertisingChannelSubTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdvertisingChannelSubTypeEnum.Merge(m, src)
}
func (m *AdvertisingChannelSubTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AdvertisingChannelSubTypeEnum.Size(m)
}
func (m *AdvertisingChannelSubTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdvertisingChannelSubTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdvertisingChannelSubTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType", AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_name, AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_value)
	proto.RegisterType((*AdvertisingChannelSubTypeEnum)(nil), "google.ads.googleads.v3.enums.AdvertisingChannelSubTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/advertising_channel_sub_type.proto", fileDescriptor_007989ef93506f12)
}

var fileDescriptor_007989ef93506f12 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x69, 0x02, 0x2d, 0x6c, 0x0a, 0x59, 0xb6, 0x50, 0x44, 0x21, 0x54, 0x2d, 0x77, 0xfb,
	0xe0, 0x9b, 0xb9, 0x30, 0x71, 0xb6, 0xee, 0xaa, 0xf1, 0x7a, 0xe5, 0x75, 0xc2, 0x1f, 0x45, 0x5a,
	0x39, 0xb5, 0x65, 0x22, 0x25, 0x76, 0x94, 0x75, 0x22, 0xf5, 0x41, 0x78, 0x01, 0x8e, 0x3c, 0x0a,
	0x8f, 0xc2, 0xb5, 0x2f, 0x80, 0x6c, 0xd7, 0x69, 0x2f, 0xe1, 0x62, 0x8d, 0xbe, 0xf9, 0xed, 0x37,
	0x63, 0x7d, 0x83, 0x3e, 0xa7, 0x79, 0x9e, 0xce, 0x13, 0x33, 0x8a, 0xb5, 0x59, 0x97, 0x65, 0xb5,
	0xb1, 0xcc, 0x24, 0x5b, 0x2f, 0xb4, 0x19, 0xc5, 0x9b, 0x64, 0x55, 0xcc, 0xf4, 0x2c, 0x4b, 0xd5,
	0xf5, 0x8f, 0x28, 0xcb, 0x92, 0xb9, 0xd2, 0xeb, 0xa9, 0x2a, 0x6e, 0x96, 0x89, 0xb1, 0x5c, 0xe5,
	0x45, 0x4e, 0x7a, 0xf5, 0x33, 0x23, 0x8a, 0xb5, 0xb1, 0x75, 0x30, 0x36, 0x96, 0x51, 0x39, 0x9c,
	0xbc, 0x6f, 0x06, 0x2c, 0x67, 0x66, 0x94, 0x65, 0x79, 0x11, 0x15, 0xb3, 0x3c, 0xd3, 0xf5, 0xe3,
	0xf3, 0x9f, 0x6d, 0xd4, 0x83, 0xfb, 0x19, 0x4e, 0x3d, 0x42, 0xae, 0xa7, 0xe1, 0xcd, 0x32, 0xa1,
	0xd9, 0x7a, 0x71, 0x7e, 0xdb, 0x42, 0x6f, 0x77, 0x12, 0xa4, 0x8b, 0x3a, 0x23, 0x2e, 0x05, 0x75,
	0xd8, 0x05, 0xa3, 0x03, 0xfc, 0x88, 0x74, 0xd0, 0xc1, 0x88, 0x5f, 0x71, 0xff, 0x0b, 0xc7, 0x7b,
	0xe4, 0x35, 0x7a, 0x29, 0x29, 0x04, 0xce, 0xa5, 0xf2, 0xfc, 0x3e, 0x1b, 0x52, 0x05, 0x42, 0xe0,
	0x16, 0x39, 0x46, 0x64, 0xc0, 0xa4, 0x18, 0xc2, 0xb7, 0x87, 0x7a, 0x9b, 0x10, 0xf4, 0xe2, 0x0e,
	0xa7, 0x5f, 0x45, 0x40, 0xa5, 0xc4, 0x8f, 0xc9, 0x11, 0xea, 0x36, 0x6c, 0x23, 0x3e, 0x29, 0x0d,
	0xe4, 0xa5, 0x2f, 0x04, 0xe3, 0xae, 0x92, 0x1e, 0x04, 0xa1, 0x82, 0x81, 0xc4, 0xfb, 0xe4, 0x15,
	0xc2, 0x0d, 0xec, 0x7a, 0xc0, 0x86, 0x0a, 0x06, 0xf8, 0x80, 0x9c, 0xa0, 0xe3, 0x46, 0xad, 0x61,
	0x07, 0x3c, 0x01, 0xcc, 0xe5, 0xf8, 0x69, 0x69, 0x3f, 0x66, 0x03, 0xea, 0x2b, 0x7f, 0x14, 0xca,
	0x30, 0xa0, 0xe0, 0xe1, 0x67, 0x04, 0xa3, 0xc3, 0x5a, 0x04, 0x27, 0x64, 0x3e, 0xc7, 0x88, 0xbc,
	0x41, 0x47, 0xb5, 0xc2, 0x7d, 0xae, 0xe4, 0x15, 0x13, 0x02, 0xfa, 0x43, 0x8a, 0x3b, 0x25, 0x0a,
	0x42, 0xdc, 0x3b, 0x1e, 0x92, 0x53, 0xf4, 0xee, 0xa1, 0xa2, 0x2e, 0xfc, 0x40, 0x51, 0xee, 0x82,
	0x4b, 0x3d, 0xca, 0x43, 0xfc, 0x9c, 0x7c, 0x44, 0xa7, 0xdb, 0xe5, 0x1d, 0xdf, 0x13, 0x10, 0x30,
	0xe9, 0x73, 0x35, 0x64, 0x32, 0x2c, 0xa5, 0xf2, 0x4f, 0xba, 0xfd, 0xdb, 0x3d, 0x74, 0x76, 0x9d,
	0x2f, 0x8c, 0xff, 0x66, 0xdb, 0xff, 0xb0, 0x33, 0x18, 0x51, 0xa6, 0x2b, 0xf6, 0xbe, 0xf7, 0xef,
	0x0c, 0xd2, 0x7c, 0x1e, 0x65, 0xa9, 0x91, 0xaf, 0x52, 0x33, 0x4d, 0xb2, 0x2a, 0xfb, 0xe6, 0xdc,
	0x96, 0x33, 0xbd, 0xe3, 0xfa, 0x3e, 0x55, 0xdf, 0x5f, 0xad, 0xb6, 0x0b, 0xf0, 0xbb, 0xd5, 0x73,
	0x6b, 0x2b, 0x88, 0xb5, 0x51, 0x97, 0x65, 0x35, 0xb6, 0x8c, 0xf2, 0x4c, 0xf4, 0x9f, 0xa6, 0x3f,
	0x81, 0x58, 0x4f, 0xb6, 0xfd, 0xc9, 0xd8, 0x9a, 0x54, 0xfd, 0xbf, 0xad, 0xb3, 0x5a, 0xb4, 0x6d,
	0x88, 0xb5, 0x6d, 0x6f, 0x09, 0xdb, 0x1e, 0x5b, 0xb6, 0x5d, 0x31, 0xd3, 0xfd, 0x6a, 0x31, 0xeb,
	0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0xfd, 0xc9, 0x77, 0x15, 0x03, 0x00, 0x00,
}
