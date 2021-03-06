// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/placement_type.proto

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

// Possible placement types for a feed mapping.
type PlacementTypeEnum_PlacementType int32

const (
	// Not specified.
	PlacementTypeEnum_UNSPECIFIED PlacementTypeEnum_PlacementType = 0
	// Used for return value only. Represents value unknown in this version.
	PlacementTypeEnum_UNKNOWN PlacementTypeEnum_PlacementType = 1
	// Websites(e.g. 'www.flowers4sale.com').
	PlacementTypeEnum_WEBSITE PlacementTypeEnum_PlacementType = 2
	// Mobile application categories(e.g. 'Games').
	PlacementTypeEnum_MOBILE_APP_CATEGORY PlacementTypeEnum_PlacementType = 3
	// mobile applications(e.g. 'mobileapp::2-com.whatsthewordanswers').
	PlacementTypeEnum_MOBILE_APPLICATION PlacementTypeEnum_PlacementType = 4
	// YouTube videos(e.g. 'youtube.com/video/wtLJPvx7-ys').
	PlacementTypeEnum_YOUTUBE_VIDEO PlacementTypeEnum_PlacementType = 5
	// YouTube channels(e.g. 'youtube.com::L8ZULXASCc1I_oaOT0NaOQ').
	PlacementTypeEnum_YOUTUBE_CHANNEL PlacementTypeEnum_PlacementType = 6
)

var PlacementTypeEnum_PlacementType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "WEBSITE",
	3: "MOBILE_APP_CATEGORY",
	4: "MOBILE_APPLICATION",
	5: "YOUTUBE_VIDEO",
	6: "YOUTUBE_CHANNEL",
}

var PlacementTypeEnum_PlacementType_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"WEBSITE":             2,
	"MOBILE_APP_CATEGORY": 3,
	"MOBILE_APPLICATION":  4,
	"YOUTUBE_VIDEO":       5,
	"YOUTUBE_CHANNEL":     6,
}

func (x PlacementTypeEnum_PlacementType) String() string {
	return proto.EnumName(PlacementTypeEnum_PlacementType_name, int32(x))
}

func (PlacementTypeEnum_PlacementType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d0c6cf34c3732264, []int{0, 0}
}

// Container for enum describing possible placement types.
type PlacementTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlacementTypeEnum) Reset()         { *m = PlacementTypeEnum{} }
func (m *PlacementTypeEnum) String() string { return proto.CompactTextString(m) }
func (*PlacementTypeEnum) ProtoMessage()    {}
func (*PlacementTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0c6cf34c3732264, []int{0}
}

func (m *PlacementTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlacementTypeEnum.Unmarshal(m, b)
}
func (m *PlacementTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlacementTypeEnum.Marshal(b, m, deterministic)
}
func (m *PlacementTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlacementTypeEnum.Merge(m, src)
}
func (m *PlacementTypeEnum) XXX_Size() int {
	return xxx_messageInfo_PlacementTypeEnum.Size(m)
}
func (m *PlacementTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PlacementTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PlacementTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.PlacementTypeEnum_PlacementType", PlacementTypeEnum_PlacementType_name, PlacementTypeEnum_PlacementType_value)
	proto.RegisterType((*PlacementTypeEnum)(nil), "google.ads.googleads.v2.enums.PlacementTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/placement_type.proto", fileDescriptor_d0c6cf34c3732264)
}

var fileDescriptor_d0c6cf34c3732264 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x5d, 0x6a, 0xab, 0x40,
	0x14, 0xbe, 0x9a, 0x7b, 0x73, 0x61, 0x42, 0x88, 0x99, 0xc0, 0xbd, 0x50, 0x9a, 0x87, 0x64, 0x01,
	0x23, 0xd8, 0xb7, 0xe9, 0xd3, 0x68, 0xa6, 0xa9, 0x34, 0x55, 0x69, 0xd4, 0x90, 0x22, 0x88, 0x8d,
	0x22, 0x81, 0x38, 0x23, 0x19, 0x13, 0xc8, 0x3a, 0xba, 0x83, 0xbe, 0xb5, 0x4b, 0xe9, 0x46, 0x0a,
	0x5d, 0x45, 0x51, 0x6b, 0x4a, 0x1e, 0xda, 0x97, 0xe1, 0x9b, 0xf3, 0xfd, 0x70, 0xce, 0x07, 0xb4,
	0x94, 0xf3, 0x74, 0x93, 0xa8, 0x51, 0x2c, 0xd4, 0x1a, 0x96, 0x68, 0xaf, 0xa9, 0x09, 0xdb, 0x65,
	0x42, 0xcd, 0x37, 0xd1, 0x2a, 0xc9, 0x12, 0x56, 0x84, 0xc5, 0x21, 0x4f, 0x50, 0xbe, 0xe5, 0x05,
	0x87, 0xc3, 0x5a, 0x88, 0xa2, 0x58, 0xa0, 0xa3, 0x07, 0xed, 0x35, 0x54, 0x79, 0xce, 0xce, 0x9b,
	0xc8, 0x7c, 0xad, 0x46, 0x8c, 0xf1, 0x22, 0x2a, 0xd6, 0x9c, 0x89, 0xda, 0x3c, 0x7e, 0x96, 0x40,
	0xdf, 0x69, 0x52, 0xdd, 0x43, 0x9e, 0x50, 0xb6, 0xcb, 0xc6, 0x8f, 0x12, 0xe8, 0x9e, 0x4c, 0x61,
	0x0f, 0x74, 0x3c, 0x6b, 0xee, 0x50, 0xc3, 0xbc, 0x32, 0xe9, 0x44, 0xf9, 0x05, 0x3b, 0xe0, 0xaf,
	0x67, 0xdd, 0x58, 0xf6, 0xc2, 0x52, 0xa4, 0xf2, 0xb3, 0xa0, 0xfa, 0xdc, 0x74, 0xa9, 0x22, 0xc3,
	0xff, 0x60, 0x70, 0x6b, 0xeb, 0xe6, 0x8c, 0x86, 0xc4, 0x71, 0x42, 0x83, 0xb8, 0x74, 0x6a, 0xdf,
	0x2d, 0x95, 0x16, 0xfc, 0x07, 0xe0, 0x17, 0x31, 0x33, 0x0d, 0xe2, 0x9a, 0xb6, 0xa5, 0xfc, 0x86,
	0x7d, 0xd0, 0x5d, 0xda, 0x9e, 0xeb, 0xe9, 0x34, 0xf4, 0xcd, 0x09, 0xb5, 0x95, 0x3f, 0x70, 0x00,
	0x7a, 0xcd, 0xc8, 0xb8, 0x26, 0x96, 0x45, 0x67, 0x4a, 0x5b, 0x7f, 0x93, 0xc0, 0x68, 0xc5, 0x33,
	0xf4, 0xe3, 0xbd, 0x3a, 0x3c, 0x59, 0xdc, 0x29, 0xaf, 0x74, 0xa4, 0x7b, 0xfd, 0xd3, 0x94, 0xf2,
	0x4d, 0xc4, 0x52, 0xc4, 0xb7, 0xa9, 0x9a, 0x26, 0xac, 0xea, 0xa0, 0x29, 0x3a, 0x5f, 0x8b, 0x6f,
	0x7a, 0xbf, 0xac, 0xde, 0x27, 0xb9, 0x35, 0x25, 0xe4, 0x45, 0x1e, 0x4e, 0xeb, 0x28, 0x12, 0x0b,
	0x54, 0xc3, 0x12, 0xf9, 0x1a, 0x2a, 0xab, 0x13, 0xaf, 0x0d, 0x1f, 0x90, 0x58, 0x04, 0x47, 0x3e,
	0xf0, 0xb5, 0xa0, 0xe2, 0xdf, 0xe5, 0x51, 0x3d, 0xc4, 0x98, 0xc4, 0x02, 0xe3, 0xa3, 0x02, 0x63,
	0x5f, 0xc3, 0xb8, 0xd2, 0x3c, 0xb4, 0xab, 0xc5, 0x2e, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x66,
	0xda, 0xd4, 0x4a, 0x0f, 0x02, 0x00, 0x00,
}
