// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/content_label_type.proto

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

// Enum listing the content label types supported by ContentLabel criterion.
type ContentLabelTypeEnum_ContentLabelType int32

const (
	// Not specified.
	ContentLabelTypeEnum_UNSPECIFIED ContentLabelTypeEnum_ContentLabelType = 0
	// Used for return value only. Represents value unknown in this version.
	ContentLabelTypeEnum_UNKNOWN ContentLabelTypeEnum_ContentLabelType = 1
	// Sexually suggestive content.
	ContentLabelTypeEnum_SEXUALLY_SUGGESTIVE ContentLabelTypeEnum_ContentLabelType = 2
	// Below the fold placement.
	ContentLabelTypeEnum_BELOW_THE_FOLD ContentLabelTypeEnum_ContentLabelType = 3
	// Parked domain.
	ContentLabelTypeEnum_PARKED_DOMAIN ContentLabelTypeEnum_ContentLabelType = 4
	// Game.
	ContentLabelTypeEnum_GAME ContentLabelTypeEnum_ContentLabelType = 5
	// Juvenile, gross & bizarre content.
	ContentLabelTypeEnum_JUVENILE ContentLabelTypeEnum_ContentLabelType = 6
	// Profanity & rough language.
	ContentLabelTypeEnum_PROFANITY ContentLabelTypeEnum_ContentLabelType = 7
	// Death & tragedy.
	ContentLabelTypeEnum_TRAGEDY ContentLabelTypeEnum_ContentLabelType = 8
	// Video.
	ContentLabelTypeEnum_VIDEO ContentLabelTypeEnum_ContentLabelType = 9
	// Content rating: G.
	ContentLabelTypeEnum_VIDEO_RATING_DV_G ContentLabelTypeEnum_ContentLabelType = 10
	// Content rating: PG.
	ContentLabelTypeEnum_VIDEO_RATING_DV_PG ContentLabelTypeEnum_ContentLabelType = 11
	// Content rating: T.
	ContentLabelTypeEnum_VIDEO_RATING_DV_T ContentLabelTypeEnum_ContentLabelType = 12
	// Content rating: MA.
	ContentLabelTypeEnum_VIDEO_RATING_DV_MA ContentLabelTypeEnum_ContentLabelType = 13
	// Content rating: not yet rated.
	ContentLabelTypeEnum_VIDEO_NOT_YET_RATED ContentLabelTypeEnum_ContentLabelType = 14
	// Embedded video.
	ContentLabelTypeEnum_EMBEDDED_VIDEO ContentLabelTypeEnum_ContentLabelType = 15
	// Live streaming video.
	ContentLabelTypeEnum_LIVE_STREAMING_VIDEO ContentLabelTypeEnum_ContentLabelType = 16
)

var ContentLabelTypeEnum_ContentLabelType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "SEXUALLY_SUGGESTIVE",
	3:  "BELOW_THE_FOLD",
	4:  "PARKED_DOMAIN",
	5:  "GAME",
	6:  "JUVENILE",
	7:  "PROFANITY",
	8:  "TRAGEDY",
	9:  "VIDEO",
	10: "VIDEO_RATING_DV_G",
	11: "VIDEO_RATING_DV_PG",
	12: "VIDEO_RATING_DV_T",
	13: "VIDEO_RATING_DV_MA",
	14: "VIDEO_NOT_YET_RATED",
	15: "EMBEDDED_VIDEO",
	16: "LIVE_STREAMING_VIDEO",
}

var ContentLabelTypeEnum_ContentLabelType_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"SEXUALLY_SUGGESTIVE":  2,
	"BELOW_THE_FOLD":       3,
	"PARKED_DOMAIN":        4,
	"GAME":                 5,
	"JUVENILE":             6,
	"PROFANITY":            7,
	"TRAGEDY":              8,
	"VIDEO":                9,
	"VIDEO_RATING_DV_G":    10,
	"VIDEO_RATING_DV_PG":   11,
	"VIDEO_RATING_DV_T":    12,
	"VIDEO_RATING_DV_MA":   13,
	"VIDEO_NOT_YET_RATED":  14,
	"EMBEDDED_VIDEO":       15,
	"LIVE_STREAMING_VIDEO": 16,
}

func (x ContentLabelTypeEnum_ContentLabelType) String() string {
	return proto.EnumName(ContentLabelTypeEnum_ContentLabelType_name, int32(x))
}

func (ContentLabelTypeEnum_ContentLabelType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bcaccae1ff7e8076, []int{0, 0}
}

// Container for enum describing content label types in ContentLabel.
type ContentLabelTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentLabelTypeEnum) Reset()         { *m = ContentLabelTypeEnum{} }
func (m *ContentLabelTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ContentLabelTypeEnum) ProtoMessage()    {}
func (*ContentLabelTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcaccae1ff7e8076, []int{0}
}

func (m *ContentLabelTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentLabelTypeEnum.Unmarshal(m, b)
}
func (m *ContentLabelTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentLabelTypeEnum.Marshal(b, m, deterministic)
}
func (m *ContentLabelTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentLabelTypeEnum.Merge(m, src)
}
func (m *ContentLabelTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ContentLabelTypeEnum.Size(m)
}
func (m *ContentLabelTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentLabelTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ContentLabelTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.ContentLabelTypeEnum_ContentLabelType", ContentLabelTypeEnum_ContentLabelType_name, ContentLabelTypeEnum_ContentLabelType_value)
	proto.RegisterType((*ContentLabelTypeEnum)(nil), "google.ads.googleads.v1.enums.ContentLabelTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/content_label_type.proto", fileDescriptor_bcaccae1ff7e8076)
}

var fileDescriptor_bcaccae1ff7e8076 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xdd, 0x6e, 0x9b, 0x3c,
	0x18, 0xfe, 0x42, 0xff, 0x12, 0xa7, 0x69, 0x5d, 0x7f, 0xed, 0x36, 0x4d, 0xeb, 0x41, 0x7b, 0x01,
	0xa0, 0x68, 0xd2, 0x0e, 0xd8, 0x91, 0x53, 0xbf, 0x61, 0xac, 0x04, 0x50, 0x30, 0x74, 0x99, 0x22,
	0x59, 0xb4, 0x20, 0x14, 0x29, 0xc1, 0xa8, 0xd0, 0x4a, 0xbd, 0x9d, 0x1d, 0xee, 0x52, 0x76, 0xbc,
	0xab, 0xa8, 0xb4, 0x7b, 0x98, 0x8c, 0x97, 0x1c, 0x44, 0xdd, 0x4e, 0xd0, 0xa3, 0xf7, 0xf9, 0xc1,
	0xf6, 0xf3, 0xa2, 0x0f, 0x85, 0x94, 0xc5, 0x32, 0xb7, 0xd2, 0xac, 0xb6, 0x34, 0x54, 0xe8, 0x71,
	0x68, 0xe5, 0xe5, 0xc3, 0xaa, 0xb6, 0xee, 0x64, 0xd9, 0xe4, 0x65, 0x23, 0x96, 0xe9, 0x6d, 0xbe,
	0x14, 0xcd, 0x53, 0x95, 0x9b, 0xd5, 0xbd, 0x6c, 0x24, 0x39, 0xd7, 0x62, 0x33, 0xcd, 0x6a, 0x73,
	0xe3, 0x33, 0x1f, 0x87, 0x66, 0xeb, 0x7b, 0xfb, 0x6e, 0x1d, 0x5b, 0x2d, 0xac, 0xb4, 0x2c, 0x65,
	0x93, 0x36, 0x0b, 0x59, 0xd6, 0xda, 0x7c, 0xf9, 0xcb, 0x40, 0xa7, 0x57, 0x3a, 0xd9, 0x53, 0xc1,
	0xfc, 0xa9, 0xca, 0xa1, 0x7c, 0x58, 0x5d, 0xfe, 0x34, 0x10, 0xde, 0x26, 0xc8, 0x31, 0xea, 0xc7,
	0x7e, 0x14, 0xc2, 0x95, 0x3b, 0x76, 0x81, 0xe1, 0xff, 0x48, 0x1f, 0x1d, 0xc4, 0xfe, 0xb5, 0x1f,
	0xdc, 0xf8, 0xb8, 0x43, 0x5e, 0xa3, 0xff, 0x23, 0xf8, 0x12, 0x53, 0xcf, 0x9b, 0x89, 0x28, 0x76,
	0x1c, 0x88, 0xb8, 0x9b, 0x00, 0x36, 0x08, 0x41, 0x47, 0x23, 0xf0, 0x82, 0x1b, 0xc1, 0x3f, 0x81,
	0x18, 0x07, 0x1e, 0xc3, 0x3b, 0xe4, 0x04, 0x0d, 0x42, 0x3a, 0xbd, 0x06, 0x26, 0x58, 0x30, 0xa1,
	0xae, 0x8f, 0x77, 0x49, 0x17, 0xed, 0x3a, 0x74, 0x02, 0x78, 0x8f, 0x1c, 0xa2, 0xee, 0xe7, 0x38,
	0x01, 0xdf, 0xf5, 0x00, 0xef, 0x93, 0x01, 0xea, 0x85, 0xd3, 0x60, 0x4c, 0x7d, 0x97, 0xcf, 0xf0,
	0x81, 0xfa, 0x27, 0x9f, 0x52, 0x07, 0xd8, 0x0c, 0x77, 0x49, 0x0f, 0xed, 0x25, 0x2e, 0x83, 0x00,
	0xf7, 0xc8, 0x19, 0x3a, 0x69, 0xa1, 0x98, 0x52, 0xee, 0xfa, 0x8e, 0x60, 0x89, 0x70, 0x30, 0x22,
	0xaf, 0x10, 0xd9, 0x1e, 0x87, 0x0e, 0xee, 0xbf, 0x24, 0xe7, 0xf8, 0xf0, 0x25, 0xf9, 0x84, 0xe2,
	0x81, 0xba, 0x9c, 0x9e, 0xfb, 0x01, 0x17, 0x33, 0xe0, 0x8a, 0x07, 0x86, 0x8f, 0xd4, 0xe5, 0x60,
	0x32, 0x02, 0xc6, 0x80, 0x09, 0x7d, 0x94, 0x63, 0xf2, 0x06, 0x9d, 0x7a, 0x6e, 0x02, 0x22, 0xe2,
	0x53, 0xa0, 0x13, 0x15, 0xa3, 0x19, 0x3c, 0x7a, 0xee, 0xa0, 0x8b, 0x3b, 0xb9, 0x32, 0xff, 0xd9,
	0xd9, 0xe8, 0x6c, 0xfb, 0xe5, 0x43, 0x55, 0x56, 0xd8, 0xf9, 0x3a, 0xfa, 0xe3, 0x2b, 0xe4, 0x32,
	0x2d, 0x0b, 0x53, 0xde, 0x17, 0x56, 0x91, 0x97, 0x6d, 0x95, 0xeb, 0x9d, 0xa9, 0x16, 0xf5, 0x5f,
	0x56, 0xe8, 0x63, 0xfb, 0xfd, 0x66, 0xec, 0x38, 0x94, 0x7e, 0x37, 0xce, 0x1d, 0x1d, 0x45, 0xb3,
	0xda, 0xd4, 0x50, 0xa1, 0x64, 0x68, 0xaa, 0xfa, 0xeb, 0x1f, 0x6b, 0x7e, 0x4e, 0xb3, 0x7a, 0xbe,
	0xe1, 0xe7, 0xc9, 0x70, 0xde, 0xf2, 0xcf, 0xc6, 0x85, 0x1e, 0xda, 0x36, 0xcd, 0x6a, 0xdb, 0xde,
	0x28, 0x6c, 0x3b, 0x19, 0xda, 0x76, 0xab, 0xb9, 0xdd, 0x6f, 0x0f, 0xf6, 0xfe, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x20, 0x89, 0xed, 0x7b, 0xda, 0x02, 0x00, 0x00,
}
